# Ejemplos de Concurrencia en Go

Este programa demuestra las características principales de concurrencia en Go con ejemplos prácticos y ejecutables.

## Ejecutar el programa

```bash
go run concurrency_examples.go
```

## Contenido

El programa incluye 9 secciones con ejemplos de:

### 1. **Goroutines Básicas**
- Lanzamiento de goroutines con `go`
- Funciones nombradas y anónimas
- Ejecución concurrente múltiple

### 2. **Channels Unbuffered**
- Sincronización directa (rendezvous)
- Bloqueo hasta que emisor y receptor estén listos
- Uso para coordinación

### 3. **Channels Buffered**
- Desacoplamiento temporal
- Capacidad definida
- Operaciones no bloqueantes hasta llenar buffer

### 4. **Select Statement**
- Multiplexación de canales
- Timeouts
- Operaciones no bloqueantes con `default`

### 5. **Pipeline Pattern**
- Procesamiento en etapas
- Composición de funciones
- Flujo de datos unidireccional

### 6. **Fan-Out/Fan-In**
- Distribución de trabajo (fan-out)
- Consolidación de resultados (fan-in)
- Paralelización de procesamiento

### 7. **Worker Pool**
- Límite de concurrencia
- Pool fijo de workers
- Cola de trabajos

### 8. **Context Package**
- Cancelación propagada
- Timeouts y deadlines
- Control del ciclo de vida

### 9. **Rate Limiting**
- Control de frecuencia con ticker
- Token bucket para ráfagas
- Limitación de throughput

## Conceptos Clave Ilustrados

### Sintaxis Básica

```go
// Goroutine
go funcionCualquiera()

// Channel creation
ch := make(chan int)        // unbuffered
ch := make(chan int, 10)    // buffered

// Channel operations
ch <- valor     // enviar
valor := <-ch   // recibir
close(ch)       // cerrar

// Select
select {
case v := <-ch1:
    // manejar ch1
case ch2 <- v:
    // enviar a ch2
case <-time.After(1 * time.Second):
    // timeout
default:
    // no bloquear
}
```

### Patrones Comunes

**Pipeline**: Conectar etapas de procesamiento
```go
números → cuadrados → dobles → resultado
```

**Fan-Out/Fan-In**: Paralelización y consolidación
```go
         ┌→ worker1 →┐
entrada →├→ worker2 →├→ salida
         └→ worker3 →┘
```

**Worker Pool**: Concurrencia limitada
```go
jobs → [worker1, worker2, worker3] → results
```

## Notas Importantes

- Las goroutines son extremadamente ligeras (~2KB stack inicial)
- Los channels son la forma idiomática de comunicación
- `select` permite composición elegante de operaciones concurrentes
- `context` es esencial para gestión de lifecycle en sistemas complejos
- Siempre usar `go run -race` para detectar race conditions

## Experimentar

Modifica los valores de:
- Número de workers en el pool
- Tamaños de buffer en channels
- Timeouts en context
- Delays en rate limiting

Para ver cómo afectan el comportamiento del programa.