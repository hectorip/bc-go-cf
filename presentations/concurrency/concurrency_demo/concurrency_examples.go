package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Ejemplos de Concurrencia en Go ===\n")

	// 1. Goroutines básicas
	fmt.Println("1. GOROUTINES BÁSICAS")
	fmt.Println("---------------------")
	demoBasicGoroutines()
	time.Sleep(2 * time.Second)

	// 2. Channels unbuffered
	fmt.Println("\n2. CHANNELS UNBUFFERED (Sincronización)")
	fmt.Println("----------------------------------------")
	demoUnbufferedChannels()

	// 3. Channels buffered
	fmt.Println("\n3. CHANNELS BUFFERED (Desacoplamiento)")
	fmt.Println("---------------------------------------")
	demoBufferedChannels()

	// 4. Select statement
	fmt.Println("\n4. SELECT STATEMENT (Multiplexación)")
	fmt.Println("-------------------------------------")
	demoSelect()

	// 5. Pipeline pattern
	fmt.Println("\n5. PIPELINE PATTERN")
	fmt.Println("-------------------")
	demoPipeline()

	// 6. Fan-out/Fan-in pattern
	fmt.Println("\n6. FAN-OUT/FAN-IN PATTERN")
	fmt.Println("--------------------------")
	demoFanOutFanIn()

	// 7. Worker Pool pattern
	fmt.Println("\n7. WORKER POOL PATTERN")
	fmt.Println("----------------------")
	demoWorkerPool()

	// 8. Context para cancelación
	fmt.Println("\n8. CONTEXT PARA CANCELACIÓN")
	fmt.Println("----------------------------")
	demoContext()

	// 9. Rate Limiting
	fmt.Println("\n9. RATE LIMITING")
	fmt.Println("----------------")
	demoRateLimiting()

	fmt.Println("\n=== Fin de los ejemplos ===")
}

// 1. Demostración de goroutines básicas
func demoBasicGoroutines() {
	// Goroutine con función nombrada
	go sayHello("Mundo")

	// Goroutine con función anónima
	go func(msg string) {
		fmt.Printf("Mensaje desde goroutine anónima: %s\n", msg)
	}("Hola Go!")

	// Múltiples goroutines
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d ejecutándose\n", id)
		}(i)
	}
}

func sayHello(name string) {
	fmt.Printf("Hola %s desde una goroutine!\n", name)
}

// 2. Channels unbuffered para sincronización
func demoUnbufferedChannels() {
	// Canal sin buffer - sincronización directa
	ch := make(chan string)

	// Goroutine productora
	go func() {
		fmt.Println("  Enviando mensaje...")
		ch <- "Mensaje sincronizado"
		fmt.Println("  Mensaje enviado")
	}()

	// Recepción sincronizada
	fmt.Println("  Esperando mensaje...")
	msg := <-ch
	fmt.Printf("  Recibido: %s\n", msg)

	// Usando channels para esperar finalización
	done := make(chan bool)
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  Trabajo completado")
		done <- true
	}()
	<-done // Esperar hasta que termine
}

// 3. Channels buffered para desacoplamiento
func demoBufferedChannels() {
	// Canal con buffer de tamaño 3
	ch := make(chan int, 3)

	// Podemos enviar sin bloquear (hasta llenar el buffer)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Printf("  Buffer lleno con %d elementos\n", len(ch))

	// Recibir valores
	fmt.Printf("  Recibido: %d\n", <-ch)
	fmt.Printf("  Recibido: %d\n", <-ch)

	// Ahora hay espacio en el buffer
	ch <- 4
	fmt.Printf("  Elementos en buffer: %d\n", len(ch))

	// Cerrar el channel
	close(ch)

	// Recibir valores restantes después de cerrar
	for val := range ch {
		fmt.Printf("  Valor restante: %d\n", val)
	}
}

// 4. Select para multiplexación
func demoSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Mensaje del canal 1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Mensaje del canal 2"
	}()

	// Select con timeout
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("  Recibido: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("  Recibido: %s\n", msg2)
		case <-time.After(150 * time.Millisecond):
			fmt.Println("  Timeout! No hay mensajes disponibles")
		}
	}

	// Select no bloqueante con default
	select {
	case msg := <-ch1:
		fmt.Printf("  Mensaje: %s\n", msg)
	default:
		fmt.Println("  No hay mensajes (non-blocking)")
	}
}

// 5. Pipeline Pattern
func demoPipeline() {
	// Etapa 1: Generar números
	numbers := generateNumbers(5)

	// Etapa 2: Elevar al cuadrado
	squares := squareNumbers(numbers)

	// Etapa 3: Multiplicar por 2
	doubled := doubleNumbers(squares)

	// Consumir resultados
	for result := range doubled {
		fmt.Printf("  Resultado pipeline: %d\n", result)
	}
}

func generateNumbers(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func squareNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func doubleNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// 6. Fan-out/Fan-in Pattern
func demoFanOutFanIn() {
	// Canal de entrada
	in := generateNumbers(10)

	// Fan-out: distribuir trabajo a 3 workers
	c1 := processWork(in, 1)
	c2 := processWork(in, 2)
	c3 := processWork(in, 3)

	// Fan-in: consolidar resultados
	results := fanIn(c1, c2, c3)

	// Consumir resultados consolidados
	for result := range results {
		fmt.Println(result)
	}
}

func processWork(in <-chan int, workerID int) <-chan string {
	out := make(chan string)
	go func() {
		for n := range in {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			out <- fmt.Sprintf("  Worker %d procesó: %d", workerID, n)
		}
		close(out)
	}()
	return out
}

func fanIn(channels ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// Función para leer de un canal
	output := func(c <-chan string) {
		defer wg.Done()
		for msg := range c {
			out <- msg
		}
	}

	// Iniciar una goroutine por cada canal
	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	// Cerrar el canal de salida cuando todos terminen
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// 7. Worker Pool Pattern
func demoWorkerPool() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	// Iniciar workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Enviar trabajos
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Esperar que terminen los workers
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recolectar resultados
	for result := range results {
		fmt.Println(result)
	}
}

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simular trabajo
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		results <- fmt.Sprintf("  Worker %d completó job %d", id, job)
	}
}

// 8. Context para cancelación
func demoContext() {
	// Crear context con cancelación
	ctx, cancel := context.WithCancel(context.Background())

	// Iniciar worker con context
	go longRunningTask(ctx, 1)
	go longRunningTask(ctx, 2)

	// Esperar un poco y luego cancelar
	time.Sleep(500 * time.Millisecond)
	fmt.Println("  Cancelando todas las tareas...")
	cancel()

	// Dar tiempo para que se procese la cancelación
	time.Sleep(100 * time.Millisecond)

	// Context con timeout
	fmt.Println("\n  Context con timeout:")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel2()

	go taskWithTimeout(ctx2)
	
	// Esperar que termine o timeout
	<-ctx2.Done()
	if ctx2.Err() == context.DeadlineExceeded {
		fmt.Println("  Timeout alcanzado!")
	}
}

func longRunningTask(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("  Task %d: Recibida señal de cancelación\n", id)
			return
		default:
			fmt.Printf("  Task %d: Trabajando...\n", id)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func taskWithTimeout(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  Tarea completada")
	case <-ctx.Done():
		fmt.Println("  Tarea cancelada por timeout")
	}
}

// 9. Rate Limiting
func demoRateLimiting() {
	// Simulación de requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Rate limiter simple con ticker
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()

	fmt.Println("  Rate limiting (1 request cada 200ms):")
	for req := range requests {
		<-limiter.C
		fmt.Printf("  Procesando request %d en %s\n", req, time.Now().Format("15:04:05.000"))
	}

	// Token bucket para permitir ráfagas
	fmt.Println("\n  Token bucket (permite ráfagas):")
	demoTokenBucket()
}

func demoTokenBucket() {
	// Bucket con capacidad de 3 tokens
	tokens := make(chan struct{}, 3)
	
	// Llenar el bucket inicialmente
	for i := 0; i < cap(tokens); i++ {
		tokens <- struct{}{}
	}

	// Reponer tokens cada 500ms
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case tokens <- struct{}{}:
				fmt.Println("    Token añadido al bucket")
			default:
				// Bucket lleno
			}
		}
	}()

	// Procesar requests consumiendo tokens
	for i := 1; i <= 5; i++ {
		<-tokens // Consumir un token
		go func(id int) {
			fmt.Printf("  Request %d procesado (token consumido)\n", id)
		}(i)
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)
}