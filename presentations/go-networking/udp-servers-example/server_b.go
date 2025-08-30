package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type ServerB struct {
	conn       *net.UDPConn
	serverAAddr *net.UDPAddr
	stopChan   chan bool
	messages   chan Message
	stats      Statistics
}

type Statistics struct {
	HeartbeatsReceived int
	MessagesReceived   int
	MessagesProcessed  int
	LastHeartbeat      time.Time
}

func NewServerB(listenPort int, serverAPort int) (*ServerB, error) {
	listenAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		return nil, err
	}

	serverAAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", serverAPort))
	if err != nil {
		return nil, err
	}

	return &ServerB{
		conn:        conn,
		serverAAddr: serverAAddr,
		stopChan:    make(chan bool),
		messages:    make(chan Message, 100),
		stats:       Statistics{},
	}, nil
}

func (s *ServerB) Start() {
	log.Printf("[Server B] Iniciado en %s", s.conn.LocalAddr())
	
	go s.receiveMessages()
	go s.processMessages()
	go s.periodicTasks()
}

func (s *ServerB) receiveMessages() {
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
				log.Printf("[Server B] Error recibiendo: %v", err)
				continue
			}

			content := string(buffer[:n])
			msgType := s.parseMessageType(content)
			
			msg := Message{
				Type:    msgType,
				Content: content,
				From:    addr.String(),
				Time:    time.Now(),
			}
			s.messages <- msg
			s.stats.MessagesReceived++
		}
	}
}

func (s *ServerB) parseMessageType(content string) string {
	if strings.HasPrefix(content, "HEARTBEAT:") {
		return "heartbeat"
	} else if strings.HasPrefix(content, "MSG:") {
		return "message"
	}
	return "unknown"
}

func (s *ServerB) processMessages() {
	for msg := range s.messages {
		switch msg.Type {
		case "heartbeat":
			s.handleHeartbeat(msg)
		case "message":
			s.handleMessage(msg)
		default:
			log.Printf("[Server B] Mensaje desconocido de %s: %s", msg.From, msg.Content)
		}
		s.stats.MessagesProcessed++
	}
}

func (s *ServerB) handleHeartbeat(msg Message) {
	parts := strings.Split(msg.Content, ":")
	if len(parts) >= 4 {
		log.Printf("[Server B] Heartbeat #%s de %s - Status: %s", parts[1], parts[2], parts[3])
		s.stats.HeartbeatsReceived++
		s.stats.LastHeartbeat = msg.Time
		
		response := fmt.Sprintf("ACK:HEARTBEAT:%s", parts[1])
		s.sendResponse(response, msg.From)
	}
}

func (s *ServerB) handleMessage(msg Message) {
	parts := strings.Split(msg.Content, ":")
	if len(parts) >= 3 {
		content := parts[1]
		sender := parts[2]
		log.Printf("[Server B] Mensaje de %s: %s", sender, content)
		
		if content == "PONG" {
			log.Printf("[Server B] PONG recibido, conexión bidireccional confirmada")
		}
		
		response := fmt.Sprintf("RECEIVED:%s:OK", content)
		s.sendResponse(response, msg.From)
	}
}

func (s *ServerB) sendResponse(content string, to string) {
	addr, err := net.ResolveUDPAddr("udp", to)
	if err != nil {
		log.Printf("[Server B] Error resolviendo dirección %s: %v", to, err)
		return
	}
	
	_, err = s.conn.WriteToUDP([]byte(content), addr)
	if err != nil {
		log.Printf("[Server B] Error enviando respuesta: %v", err)
	}
}

func (s *ServerB) periodicTasks() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	
	pingTicker := time.NewTicker(15 * time.Second)
	defer pingTicker.Stop()
	
	for {
		select {
		case <-s.stopChan:
			return
		case <-ticker.C:
			s.printStatistics()
			s.checkHeartbeatTimeout()
		case <-pingTicker.C:
			s.sendPing()
		}
	}
}

func (s *ServerB) printStatistics() {
	log.Printf("[Server B] === ESTADÍSTICAS ===")
	log.Printf("[Server B] Heartbeats recibidos: %d", s.stats.HeartbeatsReceived)
	log.Printf("[Server B] Mensajes recibidos: %d", s.stats.MessagesReceived)
	log.Printf("[Server B] Mensajes procesados: %d", s.stats.MessagesProcessed)
	if !s.stats.LastHeartbeat.IsZero() {
		log.Printf("[Server B] Último heartbeat hace: %.0f segundos", time.Since(s.stats.LastHeartbeat).Seconds())
	}
	log.Printf("[Server B] ===================")
}

func (s *ServerB) checkHeartbeatTimeout() {
	if !s.stats.LastHeartbeat.IsZero() && time.Since(s.stats.LastHeartbeat) > 10*time.Second {
		log.Printf("[Server B] ⚠️  ALERTA: No se han recibido heartbeats en más de 10 segundos")
	}
}

func (s *ServerB) sendPing() {
	_, err := s.conn.WriteToUDP([]byte("PING"), s.serverAAddr)
	if err != nil {
		log.Printf("[Server B] Error enviando PING: %v", err)
	} else {
		log.Printf("[Server B] PING enviado a Server A")
	}
}

func (s *ServerB) Stop() {
	close(s.stopChan)
	s.conn.Close()
	close(s.messages)
	log.Println("[Server B] Detenido")
}