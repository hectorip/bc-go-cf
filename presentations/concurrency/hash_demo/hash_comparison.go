package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

// findHashWithZeros busca un hash que tenga n ceros al principio
func findHashWithZeros(start, end int, zeros int, wg *sync.WaitGroup, result chan<- string) {
	if wg != nil {
		defer wg.Done()
	}
	
	target := strings.Repeat("0", zeros)
	
	for i := start; i < end; i++ {
		data := fmt.Sprintf("data_%d", i)
		hash := sha256.Sum256([]byte(data))
		hashStr := hex.EncodeToString(hash[:])
		
		if strings.HasPrefix(hashStr, target) {
			select {
			case result <- fmt.Sprintf("Encontrado: %s -> %s", data, hashStr):
			default:
			}
			if wg != nil {
				return
			}
		}
	}
}

// Versión secuencial
func sequentialSearch(iterations int, zeros int) (string, time.Duration) {
	start := time.Now()
	result := make(chan string, 1)
	
	findHashWithZeros(0, iterations, zeros, nil, result)
	
	elapsed := time.Since(start)
	
	select {
	case res := <-result:
		return res, elapsed
	default:
		return "No encontrado", elapsed
	}
}

// Versión concurrente
func concurrentSearch(iterations int, zeros int, numGoroutines int) (string, time.Duration) {
	start := time.Now()
	
	var wg sync.WaitGroup
	result := make(chan string, numGoroutines)
	
	chunkSize := iterations / numGoroutines
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		startIdx := i * chunkSize
		endIdx := startIdx + chunkSize
		if i == numGoroutines-1 {
			endIdx = iterations
		}
		go findHashWithZeros(startIdx, endIdx, zeros, &wg, result)
	}
	
	wg.Wait()
	elapsed := time.Since(start)
	
	select {
	case res := <-result:
		return res, elapsed
	default:
		return "No encontrado", elapsed
	}
}

// Versión concurrente con overhead adicional (comunicación excesiva)
func concurrentWithOverhead(iterations int, zeros int, numGoroutines int) (string, time.Duration) {
	start := time.Now()
	
	var wg sync.WaitGroup
	result := make(chan string, iterations) // Canal muy grande (overhead de memoria)
	progress := make(chan int, iterations) // Canal adicional para "progreso"
	
	chunkSize := iterations / numGoroutines
	
	// Goroutine adicional innecesaria para "monitorear" progreso
	go func() {
		count := 0
		for range progress {
			count++
			// Overhead adicional: operación innecesaria
			_ = fmt.Sprintf("Progreso: %d", count)
		}
	}()
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		startIdx := i * chunkSize
		endIdx := startIdx + chunkSize
		if i == numGoroutines-1 {
			endIdx = iterations
		}
		
		go func(start, end int) {
			defer wg.Done()
			target := strings.Repeat("0", zeros)
			
			for j := start; j < end; j++ {
				// Overhead: comunicación excesiva
				select {
				case progress <- j:
				default:
				}
				
				data := fmt.Sprintf("data_%d", j)
				hash := sha256.Sum256([]byte(data))
				hashStr := hex.EncodeToString(hash[:])
				
				if strings.HasPrefix(hashStr, target) {
					select {
					case result <- fmt.Sprintf("Encontrado: %s -> %s", data, hashStr):
					default:
					}
					return
				}
			}
		}(startIdx, endIdx)
	}
	
	wg.Wait()
	close(progress)
	elapsed := time.Since(start)
	
	select {
	case res := <-result:
		return res, elapsed
	default:
		return "No encontrado", elapsed
	}
}

func main() {
	fmt.Println("=== Comparación de Concurrencia vs Secuencialidad ===")
	fmt.Printf("CPUs disponibles: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n\n", runtime.GOMAXPROCS(0))
	
	iterations := 1000000
	zeros := 5 // Buscar hashes que empiecen con 5 ceros
	
	fmt.Printf("Buscando hash con %d ceros al principio en %d iteraciones\n\n", zeros, iterations)
	
	// Test 1: Versión secuencial
	fmt.Println("1. VERSIÓN SECUENCIAL:")
	result1, time1 := sequentialSearch(iterations, zeros)
	fmt.Printf("   Resultado: %s\n", result1)
	fmt.Printf("   Tiempo: %v\n\n", time1)
	
	// Test 2: Concurrente con pocas goroutines
	fmt.Println("2. VERSIÓN CONCURRENTE (2 goroutines):")
	result2, time2 := concurrentSearch(iterations, zeros, 2)
	fmt.Printf("   Resultado: %s\n", result2)
	fmt.Printf("   Tiempo: %v\n", time2)
	fmt.Printf("   Speedup vs secuencial: %.2fx\n\n", float64(time1)/float64(time2))
	
	// Test 3: Concurrente con número óptimo de goroutines
	numCPU := runtime.NumCPU()
	fmt.Printf("3. VERSIÓN CONCURRENTE (%d goroutines - igual a CPUs):\n", numCPU)
	result3, time3 := concurrentSearch(iterations, zeros, numCPU)
	fmt.Printf("   Resultado: %s\n", result3)
	fmt.Printf("   Tiempo: %v\n", time3)
	fmt.Printf("   Speedup vs secuencial: %.2fx\n\n", float64(time1)/float64(time3))
	
	// Test 4: Concurrente con demasiadas goroutines
	fmt.Println("4. VERSIÓN CONCURRENTE (1000 goroutines - excesivas):")
	result4, time4 := concurrentSearch(iterations, zeros, 1000)
	fmt.Printf("   Resultado: %s\n", result4)
	fmt.Printf("   Tiempo: %v\n", time4)
	fmt.Printf("   Speedup vs secuencial: %.2fx\n\n", float64(time1)/float64(time4))
	
	// Test 5: Concurrente con overhead adicional
	fmt.Println("5. VERSIÓN CONCURRENTE CON OVERHEAD (100 goroutines + comunicación excesiva):")
	result5, time5 := concurrentWithOverhead(iterations, zeros, 100)
	fmt.Printf("   Resultado: %s\n", result5)
	fmt.Printf("   Tiempo: %v\n", time5)
	fmt.Printf("   Speedup vs secuencial: %.2fx\n\n", float64(time1)/float64(time5))
	
	// Resumen
	fmt.Println("=== RESUMEN ===")
	fmt.Printf("Secuencial:                      %v (baseline)\n", time1)
	fmt.Printf("Concurrente (2 goroutines):      %v (%.2fx)\n", time2, float64(time1)/float64(time2))
	fmt.Printf("Concurrente (%d goroutines):      %v (%.2fx)\n", numCPU, time3, float64(time1)/float64(time3))
	fmt.Printf("Concurrente (1000 goroutines):   %v (%.2fx)\n", time4, float64(time1)/float64(time4))
	fmt.Printf("Concurrente con overhead:        %v (%.2fx)\n", time5, float64(time1)/float64(time5))
	
	fmt.Println("\n=== OBSERVACIONES ===")
	fmt.Println("1. La concurrencia NO siempre es más rápida que la secuencialidad")
	fmt.Println("2. Demasiadas goroutines pueden degradar el rendimiento")
	fmt.Println("3. El overhead de sincronización y comunicación puede superar los beneficios")
	fmt.Println("4. Para tareas CPU-intensivas, el número óptimo de goroutines ≈ número de CPUs")
}