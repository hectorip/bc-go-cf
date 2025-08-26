package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=== Go Scheduler Preemption Demo ===\n")
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("NumCPU: %d\n\n", runtime.NumCPU())

	// Demo 1: Without preemption (Go < 1.14 behavior simulation)
	fmt.Println("1. SIMULATING NON-PREEMPTIVE BEHAVIOR")
	fmt.Println("--------------------------------------")
	demoWithoutPreemption()
	
	fmt.Println("\n2. PREEMPTIVE SCHEDULER (Go >= 1.14)")
	fmt.Println("-------------------------------------")
	demoWithPreemption()

	fmt.Println("\n3. TIGHT LOOP PREEMPTION")
	fmt.Println("------------------------")
	demoTightLoopPreemption()

	fmt.Println("\n4. FAIRNESS WITH MANY GOROUTINES")
	fmt.Println("---------------------------------")
	demoFairness()

	fmt.Println("\n5. PREEMPTION POINTS VISUALIZATION")
	fmt.Println("-----------------------------------")
	demoPreemptionPoints()

	fmt.Println("\n=== Demo Complete ===")
}

// Demo 1: Simulating non-preemptive behavior
func demoWithoutPreemption() {
	fmt.Println("Starting CPU-intensive goroutines (no voluntary yields)...")
	
	var wg sync.WaitGroup
	done := make(chan bool)
	
	// Start a monitoring goroutine
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				fmt.Print(".")
			case <-done:
				fmt.Println("\nMonitor stopped")
				return
			}
		}
	}()
	
	// Start CPU-intensive goroutines
	numGoroutines := runtime.GOMAXPROCS(0)
	wg.Add(numGoroutines)
	
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			
			// Tight loop without any function calls or channel operations
			// In Go < 1.14, this would monopolize the CPU
			start := time.Now()
			iterations := int64(0)
			
			// Run for 1 second
			for time.Since(start) < 1*time.Second {
				iterations++
				// Pure computation, no yield points
				_ = iterations * iterations
			}
			
			fmt.Printf("\nGoroutine %d completed %d iterations", id, iterations)
		}(i)
	}
	
	wg.Wait()
	done <- true
	time.Sleep(100 * time.Millisecond)
}

// Demo 2: Modern preemptive scheduler
func demoWithPreemption() {
	fmt.Println("Starting multiple long-running goroutines...")
	fmt.Println("The scheduler will preempt them to ensure fairness\n")
	
	var (
		counters = make([]int64, 4)
		wg       sync.WaitGroup
	)
	
	// Start time
	start := time.Now()
	duration := 2 * time.Second
	
	// Launch goroutines with heavy computation
	for i := 0; i < len(counters); i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for time.Since(start) < duration {
				// Increment counter in tight loop
				atomic.AddInt64(&counters[id], 1)
				
				// Heavy computation to consume CPU
				sum := 0
				for j := 0; j < 1000; j++ {
					sum += j * j
				}
				_ = sum
			}
		}(i)
	}
	
	// Monitor progress
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		
		for time.Since(start) < duration {
			<-ticker.C
			fmt.Printf("Progress at %v:\n", time.Since(start).Round(100*time.Millisecond))
			for i, count := range counters {
				fmt.Printf("  Goroutine %d: %d iterations\n", i, atomic.LoadInt64(&count))
			}
			fmt.Println()
		}
	}()
	
	wg.Wait()
	
	// Final results
	fmt.Println("Final iteration counts (should be relatively balanced):")
	total := int64(0)
	for i, count := range counters {
		fmt.Printf("  Goroutine %d: %d iterations\n", i, count)
		total += count
	}
	fmt.Printf("  Total: %d iterations\n", total)
	
	// Calculate fairness
	min, max := counters[0], counters[0]
	for _, count := range counters[1:] {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	fairness := float64(min) / float64(max) * 100
	fmt.Printf("  Fairness ratio: %.1f%% (closer to 100%% = better)\n", fairness)
}

// Demo 3: Tight loop preemption
func demoTightLoopPreemption() {
	fmt.Println("Testing preemption in extremely tight loops...")
	
	const numGoroutines = 10
	progress := make([]int32, numGoroutines)
	stopFlag := int32(0)
	
	// Start goroutines with infinite loops
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			// Extremely tight loop - just incrementing
			for atomic.LoadInt32(&stopFlag) == 0 {
				atomic.AddInt32(&progress[id], 1)
			}
		}(i)
	}
	
	// Let them run for a bit
	time.Sleep(100 * time.Millisecond)
	
	// Check progress - all should have made some progress
	fmt.Println("Progress after 100ms (all should be > 0):")
	allProgressed := true
	for i, p := range progress {
		value := atomic.LoadInt32(&p)
		fmt.Printf("  Goroutine %d: %d\n", i, value)
		if value == 0 {
			allProgressed = false
		}
	}
	
	// Stop all goroutines
	atomic.StoreInt32(&stopFlag, 1)
	time.Sleep(10 * time.Millisecond)
	
	if allProgressed {
		fmt.Println("✓ All goroutines made progress (preemption working!)")
	} else {
		fmt.Println("✗ Some goroutines didn't progress (preemption issue)")
	}
}

// Demo 4: Fairness with many goroutines
func demoFairness() {
	fmt.Println("Testing fairness with many goroutines competing...")
	
	const (
		numGoroutines = 100
		runDuration   = 1 * time.Second
	)
	
	counters := make([]int64, numGoroutines)
	var wg sync.WaitGroup
	start := time.Now()
	
	// Launch many competing goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for time.Since(start) < runDuration {
				// Do some work
				counters[id]++
				
				// Simulate varying workloads
				if id%10 == 0 {
					// Some goroutines do more work
					for j := 0; j < 100; j++ {
						_ = j * j
					}
				}
			}
		}(i)
	}
	
	wg.Wait()
	
	// Analyze distribution
	var total, min, max int64
	min = counters[0]
	max = counters[0]
	
	for _, count := range counters {
		total += count
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	
	average := total / numGoroutines
	
	fmt.Printf("Results for %d goroutines:\n", numGoroutines)
	fmt.Printf("  Total iterations: %d\n", total)
	fmt.Printf("  Average: %d\n", average)
	fmt.Printf("  Min: %d\n", min)
	fmt.Printf("  Max: %d\n", max)
	fmt.Printf("  Spread: %d (smaller is better)\n", max-min)
	
	// Count goroutines within 20% of average
	withinRange := 0
	for _, count := range counters {
		if float64(count) >= float64(average)*0.8 && 
		   float64(count) <= float64(average)*1.2 {
			withinRange++
		}
	}
	
	fmt.Printf("  Goroutines within 20%% of average: %d/%d (%.1f%%)\n", 
		withinRange, numGoroutines, float64(withinRange)/float64(numGoroutines)*100)
}

// Demo 5: Visualizing preemption points
func demoPreemptionPoints() {
	fmt.Println("Visualizing when preemption occurs...")
	fmt.Println("Each goroutine prints its ID when scheduled:\n")
	
	const numGoroutines = 4
	runFlag := int32(1)
	
	// Use a smaller GOMAXPROCS to make preemption more visible
	oldMaxProcs := runtime.GOMAXPROCS(2)
	defer runtime.GOMAXPROCS(oldMaxProcs)
	
	var wg sync.WaitGroup
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			lastPrint := time.Now()
			iterations := 0
			
			for atomic.LoadInt32(&runFlag) == 1 {
				iterations++
				
				// Print ID periodically to show when running
				if time.Since(lastPrint) > 50*time.Millisecond {
					fmt.Printf("%d", id)
					lastPrint = time.Now()
				}
				
				// CPU-intensive work
				sum := 0
				for j := 0; j < 10000; j++ {
					sum += j * j
				}
				_ = sum
			}
			
			fmt.Printf("\n[Goroutine %d completed %d iterations]\n", id, iterations)
		}(i)
	}
	
	// Let them run for a while
	time.Sleep(1 * time.Second)
	
	// Signal stop
	atomic.StoreInt32(&runFlag, 0)
	wg.Wait()
	
	fmt.Println("\nPattern shows scheduler preempting and switching between goroutines")
	fmt.Println("Mixed numbers indicate concurrent execution and preemption")
}