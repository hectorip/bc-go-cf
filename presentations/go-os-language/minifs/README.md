# Mini Sistema de Archivos en Go

Un sistema de archivos virtual en memoria implementado en Go, ideal para aprender conceptos de sistemas operativos y Go.

## Características

- ✅ Operaciones CRUD completas (crear, leer, actualizar, eliminar)
- ✅ Estructura de árbol de directorios
- ✅ Operaciones concurrentes seguras con mutex
- ✅ Metadatos (permisos, timestamps, tamaño)
- ✅ Navegación de rutas estilo Unix
- ✅ Recorrido recursivo del árbol
- ✅ Renombrado y movimiento de archivos
- ✅ Cálculo de tamaño de directorios

## Instalación

```bash
go get github.com/hectorip/minifs
```

## Uso Rápido

```go
package main

import (
    "fmt"
    "log"
    "github.com/hectorip/minifs"
)

func main() {
    // Crear sistema de archivos
    fs := minifs.NewFileSystem()
    
    // Crear directorios
    fs.MkdirAll("/home/user/documents", 0755)
    
    // Crear archivo
    content := []byte("Hola, mundo!")
    fs.WriteFile("/home/user/documents/saludo.txt", content)
    
    // Leer archivo
    data, err := fs.ReadFile("/home/user/documents/saludo.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Contenido: %s\n", data)
    
    // Listar directorio
    files, _ := fs.ListDir("/home/user/documents")
    for _, file := range files {
        fmt.Printf("%s (%d bytes)\n", file.Name, file.Size)
    }
}
```

## API Completa

### Crear Sistema de Archivos
```go
fs := minifs.NewFileSystem()
```

### Operaciones con Directorios
```go
// Crear directorio
fs.CreateDir("/path/to/dir", 0755)

// Crear directorios recursivamente
fs.MkdirAll("/path/to/deep/dir", 0755)

// Listar contenido
files, err := fs.ListDir("/path")

// Eliminar directorio vacío
fs.Remove("/path/to/dir")

// Eliminar recursivamente
fs.RemoveAll("/path/to/dir")
```

### Operaciones con Archivos
```go
// Escribir archivo
fs.WriteFile("/path/file.txt", []byte("contenido"))

// Leer archivo
data, err := fs.ReadFile("/path/file.txt")

// Añadir al archivo
fs.AppendFile("/path/file.txt", []byte(" más contenido"))

// Eliminar archivo
fs.Remove("/path/file.txt")
```

### Información y Metadatos
```go
// Verificar existencia
exists := fs.Exists("/path")

// Obtener información
info, err := fs.Stat("/path")
fmt.Printf("Nombre: %s, Tamaño: %d, Es directorio: %v\n", 
    info.Name, info.Size, info.IsDir)

// Calcular tamaño total
size, err := fs.Size("/path")
```

### Operaciones Avanzadas
```go
// Renombrar o mover
fs.Rename("/old/path", "/new/path")

// Recorrer árbol
fs.Walk("/", func(path string, info minifs.FileInfo) error {
    fmt.Printf("%s (%d bytes)\n", path, info.Size)
    return nil
})
```

## Ejecutar Tests

```bash
# Tests unitarios
go test ./...

# Tests con cobertura
go test -cover ./...

# Tests con race detector
go test -race ./...

# Benchmarks
go test -bench=. ./...
```

## Ejecutar Ejemplo

```bash
cd minifs/example
go run main.go
```

## Estructura del Proyecto

```
minifs/
├── minifs.go           # Implementación principal
├── minifs_test.go      # Tests unitarios y benchmarks
├── go.mod              # Módulo de Go
├── README.md           # Esta documentación
└── example/
    └── main.go         # Ejemplo de uso completo
```

## Casos de Uso

- **Testing**: Sistema de archivos en memoria para pruebas
- **Educación**: Aprender conceptos de sistemas de archivos
- **Prototipado**: Desarrollo rápido sin E/O real
- **Sandboxing**: Entorno aislado para operaciones de archivos
- **Caché**: Almacenamiento temporal en memoria

## Rendimiento

Benchmarks en MacBook Pro M1:

```
BenchmarkWriteFile         500000      2341 ns/op
BenchmarkReadFile         2000000       743 ns/op
BenchmarkListDir           300000      4127 ns/op
BenchmarkConcurrentWrites  200000      8234 ns/op
BenchmarkWalk              100000     15234 ns/op
```

## Limitaciones

- Todo se almacena en memoria (no persistente)
- Sin soporte para enlaces simbólicos
- Sin sistema de permisos de usuario/grupo
- Sin límites de cuota o espacio

## Contribuir

¡Las contribuciones son bienvenidas! Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## Licencia

MIT

## Autor

Héctor Iván Patricio Moreno