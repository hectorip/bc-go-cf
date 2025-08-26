# Go Scheduler Preemption Demo

This module demonstrates the preemptive nature of the Go scheduler, showing how it prevents goroutines from monopolizing CPU time.

## Background

Prior to Go 1.14, the Go scheduler was cooperative, meaning goroutines had to voluntarily yield control (through function calls, channel operations, etc.). Since Go 1.14, the scheduler is preemptive - it can forcibly interrupt long-running goroutines to ensure fairness.

## Run the Demo

```bash
cd preemption_demo
go run preemption.go
```

## What the Demo Shows

### 1. **Simulating Non-Preemptive Behavior**
Shows what happens when goroutines run tight loops without yield points (simulating pre-1.14 behavior).

### 2. **Preemptive Scheduler**
Demonstrates how the modern scheduler ensures all goroutines get CPU time, even when running CPU-intensive code.

### 3. **Tight Loop Preemption**
Tests extreme cases with infinite loops - all goroutines should make progress thanks to preemption.

### 4. **Fairness with Many Goroutines**
Shows how the scheduler maintains fairness when 100+ goroutines compete for CPU time.

### 5. **Preemption Points Visualization**
Visually shows when the scheduler switches between goroutines by printing their IDs.

## Key Concepts

### Preemption Triggers
The Go scheduler can preempt goroutines:
- After ~10ms of continuous execution
- At function call boundaries (stack growth checks)
- During garbage collection
- Via asynchronous preemption signals (Go 1.14+)

### Why Preemption Matters
- **Fairness**: Prevents CPU starvation
- **Responsiveness**: Ensures all goroutines make progress
- **Latency**: Reduces tail latencies in concurrent systems
- **Robustness**: System remains responsive even with badly behaved code

## Experiments to Try

1. **Vary GOMAXPROCS**:
   ```bash
   GOMAXPROCS=1 go run preemption.go
   GOMAXPROCS=8 go run preemption.go
   ```

2. **Check with older Go versions** (if available):
   Compare behavior between Go versions < 1.14 and >= 1.14

3. **Monitor with runtime metrics**:
   Add `runtime.NumGoroutine()` calls to see goroutine counts

## Technical Details

The scheduler uses **asynchronous preemption** via signals:
- On Unix: SIGURG signal
- On Windows: SuspendThread/ResumeThread

When a goroutine runs too long:
1. Timer expires (10ms)
2. Signal sent to running thread
3. Goroutine suspended at safe point
4. Another goroutine scheduled
5. Original goroutine resumed later

## Performance Impact

Preemption has minimal overhead:
- ~1% CPU overhead for signal handling
- Improves overall system responsiveness
- Better tail latency characteristics
- More predictable performance under load