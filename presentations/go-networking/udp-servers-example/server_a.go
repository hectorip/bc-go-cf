package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type ServerA struct {
	conn       *net.UDPConn
	serverBAddr *net.UDPAddr
	stopChan   chan bool
	messages   chan Message
}

func NewServerA(listenPort int, serverBPort int) (*ServerA, error) {
	listenAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		return nil, err
	}

	serverBAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", serverBPort))
	if err != nil {
		return nil, err
	}

	return &ServerA{
		conn:        conn,
		serverBAddr: serverBAddr,
		stopChan:    make(chan bool),
		messages:    make(chan Message, 100),
	}, nil
}

func (s *ServerA) Start() {
	log.Printf("[Server A] Iniciado en %s", s.conn.LocalAddr())
	
	go s.receiveMessages()
	go s.sendHeartbeats()
	go s.processMessages()
}

func (s *ServerA) receiveMessages() {
	buffer := make([]byte, 1024)
	for {
		select {
		case <-s.stopChan:
			return
		default:
			s.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
			n, addr, err := s.conn.ReadFromUDP(buffer)
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}
				log.Printf("[Server A] Error recibiendo: %v", err)
				continue
			}

			msg := Message{
				Type:    "response",
				Content: string(buffer[:n]),
				From:    addr.String(),
				Time:    time.Now(),
			}
			s.messages <- msg
		}
	}
}

func (s *ServerA) sendHeartbeats() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	
	counter := 0
	for {
		select {
		case <-s.stopChan:
			return
		case <-ticker.C:
			counter++
			msg := fmt.Sprintf("HEARTBEAT:%d:ServerA:Status=OK", counter)
			_, err := s.conn.WriteToUDP([]byte(msg), s.serverBAddr)
			if err != nil {
				log.Printf("[Server A] Error enviando heartbeat: %v", err)
			} else {
				log.Printf("[Server A] Heartbeat #%d enviado a Server B", counter)
			}
		}
	}
}

func (s *ServerA) processMessages() {
	for msg := range s.messages {
		log.Printf("[Server A] Respuesta de %s: %s", msg.From, msg.Content)
		
		if msg.Content == "PING" {
			log.Printf("[Server A] Recibido PING, enviando PONG...")
			s.SendMessage("PONG")
		}
	}
}

func (s *ServerA) SendMessage(content string) error {
	msg := fmt.Sprintf("MSG:%s:ServerA:%d", content, time.Now().Unix())
	_, err := s.conn.WriteToUDP([]byte(msg), s.serverBAddr)
	if err != nil {
		return err
	}
	log.Printf("[Server A] Mensaje enviado: %s", content)
	return nil
}

func (s *ServerA) Stop() {
	close(s.stopChan)
	s.conn.Close()
	close(s.messages)
	log.Println("[Server A] Detenido")
}