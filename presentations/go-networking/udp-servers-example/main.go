package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Message struct {
	Type    string
	Content string
	From    string
	Time    time.Time
}

func main() {
	var (
		portA    = flag.Int("port-a", 5001, "Puerto para Server A")
		portB    = flag.Int("port-b", 5002, "Puerto para Server B")
		mode     = flag.String("mode", "both", "Modo de ejecución: 'a', 'b', o 'both'")
		interact = flag.Bool("interactive", false, "Modo interactivo para enviar mensajes manualmente")
	)
	flag.Parse()

	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("=== Sistema de Comunicación UDP ===")
	log.Printf("Configuración: Server A puerto %d, Server B puerto %d", *portA, *portB)

	var wg sync.WaitGroup
	var serverA *ServerA
	var serverB *ServerB
	var err error

	switch *mode {
	case "a":
		log.Println("Iniciando solo Server A...")
		serverA, err = NewServerA(*portA, *portB)
		if err != nil {
			log.Fatalf("Error creando Server A: %v", err)
		}
		serverA.Start()
		wg.Add(1)

	case "b":
		log.Println("Iniciando solo Server B...")
		serverB, err = NewServerB(*portB, *portA)
		if err != nil {
			log.Fatalf("Error creando Server B: %v", err)
		}
		serverB.Start()
		wg.Add(1)

	case "both":
		log.Println("Iniciando ambos servidores...")
		
		serverA, err = NewServerA(*portA, *portB)
		if err != nil {
			log.Fatalf("Error creando Server A: %v", err)
		}
		
		serverB, err = NewServerB(*portB, *portA)
		if err != nil {
			log.Fatalf("Error creando Server B: %v", err)
		}
		
		serverA.Start()
		serverB.Start()
		wg.Add(2)
		
		time.Sleep(1 * time.Second)
		log.Println("=== Sistema iniciado correctamente ===")
		log.Println("Los servidores están intercambiando mensajes UDP...")
		log.Println("")

	default:
		log.Fatalf("Modo inválido: %s. Use 'a', 'b', o 'both'", *mode)
	}

	if *interact && serverA != nil {
		go interactiveMode(serverA)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("\n=== Señal de interrupción recibida, cerrando servidores... ===")
		
		if serverA != nil {
			serverA.Stop()
			wg.Done()
		}
		if serverB != nil {
			serverB.Stop()
			wg.Done()
		}
	}()

	wg.Wait()
	log.Println("=== Sistema detenido correctamente ===")
}

func interactiveMode(serverA *ServerA) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n=== MODO INTERACTIVO ===")
	fmt.Println("Comandos disponibles:")
	fmt.Println("  send <mensaje> - Enviar mensaje a Server B")
	fmt.Println("  help          - Mostrar esta ayuda")
	fmt.Println("  exit          - Salir del programa")
	fmt.Println("========================\n")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)
		
		if len(parts) == 0 {
			continue
		}

		command := strings.ToLower(parts[0])
		
		switch command {
		case "send":
			if len(parts) < 2 {
				fmt.Println("Uso: send <mensaje>")
				continue
			}
			err := serverA.SendMessage(parts[1])
			if err != nil {
				fmt.Printf("Error enviando mensaje: %v\n", err)
			} else {
				fmt.Println("Mensaje enviado exitosamente")
			}

		case "help":
			fmt.Println("\nComandos disponibles:")
			fmt.Println("  send <mensaje> - Enviar mensaje a Server B")
			fmt.Println("  help          - Mostrar esta ayuda")
			fmt.Println("  exit          - Salir del programa\n")

		case "exit":
			fmt.Println("Saliendo...")
			os.Exit(0)

		default:
			fmt.Printf("Comando desconocido: %s. Use 'help' para ver comandos disponibles.\n", command)
		}
	}
}