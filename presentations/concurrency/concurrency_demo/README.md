# Concurrency Demo

Ejemplos completos de concurrencia en Go, demostrando todos los patrones principales.

## Ejecutar

```bash
cd concurrency_demo
go run concurrency_examples.go
```

## Contenido

- Goroutines básicas
- Channels (buffered y unbuffered)
- Select statement
- Pipeline pattern
- Fan-out/Fan-in
- Worker pool
- Context para cancelación
- Rate limiting

## Detectar Race Conditions

```bash
go run -race concurrency_examples.go
```