---
title: Go + Sistema Operativo — Biblioteca Estándar para Interacción con el SO
info: "Una guía completa de los paquetes estándar de Go para trabajar con el sistema operativo"
---

# Go + Sistema Operativo

## Biblioteca Estándar para Interacción con el SO

---

## Objetivo de esta presentación

La biblioteca estándar de Go proporciona un conjunto rico y bien diseñado de paquetes para interactuar con el sistema operativo. 

Esta presentación explora:
- Las capacidades principales
- Diferencias clave entre paquetes similares
- Mejores prácticas para su uso efectivo

---
layout: section
---

# Categorías de Paquetes

---

## Organización Funcional

Los paquetes de Go para el SO se organizan en categorías funcionales claras:

* **Archivos y Sistema de Archivos**
* **Procesos y Señales**
* **Usuario y Entorno**
* **Red**
* **Logging**
* **Runtime**
* **Binarios**
* **Empaquetado**

---

# Sistema de Archivos: La Base

---

## Múltiples Enfoques

Go ofrece múltiples paquetes para trabajar con archivos, cada uno con un propósito específico.

La elección correcta depende de si necesitas:
- **Operaciones reales del SO**
- **Abstracción portable**

---

## `os` — Operaciones del Sistema Operativo Real

El paquete `os` es la interfaz directa con el sistema operativo. 

Proporciona acceso completo a:
- Archivos y directorios
- Variables de entorno
- Procesos

---

## Capacidades de `os`

* **Crear o modificar archivos:** `os.Create()`, `os.OpenFile()`
* **Gestionar el entorno:** `os.Setenv()`, `os.Getenv()`
* **Operaciones de directorio:** `os.Mkdir()`, `os.RemoveAll()`
* **Información de archivos:** `os.Stat()` → `FileInfo`

---

## Ejemplo: `os` en acción

```go
// Crear archivo con permisos específicos
f, err := os.OpenFile("data.log", 
    os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
defer f.Close()

// Manipular variables de entorno
os.Setenv("APP_MODE", "production")
home := os.Getenv("HOME")

// Información de archivo
info, _ := os.Stat("config.json")
fmt.Printf("Tamaño: %d bytes\n", info.Size())
```

---

## `io/fs` — Sistema de Archivos Abstracto

Introducido en Go 1.16, `io/fs` define **interfaces** para sistemas de archivos.

Permite código que funciona con cualquier implementación:
- Archivos reales
- Archivos embebidos
- En memoria
- Remotos

---

## Características de `io/fs`

* **Interface `FS`:** Operación básica `Open(name string)`
* **Solo lectura:** Diseñado para seguridad y portabilidad
* **Interfaces adicionales:** `ReadDirFS`, `StatFS`, `GlobFS`
* **Walk genérico:** `fs.WalkDir()` funciona con cualquier `FS`

---

## Ejemplo: Abstracción con `io/fs`

```go
// Función que acepta cualquier sistema de archivos
func loadConfig(fsys fs.FS) ([]byte, error) {
    return fs.ReadFile(fsys, "config/app.yaml")
}

// Usar con archivos reales
config, _ := loadConfig(os.DirFS("/etc/myapp"))

// Usar con archivos embebidos
//go:embed config/*
var embeddedFS embed.FS
config, _ = loadConfig(embeddedFS)
```

---

## `os` vs `io/fs`

**Diferencia clave:**
- `io/fs` es una **abstracción**
- `os` son **operaciones concretas**

**Cuándo usar cada uno:**
- Usa `io/fs` para código testeable y portable
- Usa `os` cuando necesitas control total sobre el sistema

---

## `path/filepath` vs `path`

Estos dos paquetes manejan rutas pero con propósitos **completamente diferentes**.

---

## `path/filepath` — Rutas del Sistema Operativo

* **Separador según SO:** `\` en Windows, `/` en Unix
* **Operaciones conscientes del SO:** `filepath.Join()`
* **Rutas absolutas:** `filepath.Abs()`
* **Patrones glob:** `filepath.Glob("*.go")`
* **Recorrido eficiente:** `filepath.WalkDir()`

---

## Ejemplo: `filepath` multiplataforma

```go
// Construye ruta correcta según el SO
logPath := filepath.Join("var", "log", "app.log")
// Windows: var\log\app.log
// Unix: var/log/app.log

// Obtener ruta absoluta
absPath, _ := filepath.Abs("./config")

// Buscar archivos Go
goFiles, _ := filepath.Glob("src/*.go")
```

---

## Recorrido de directorios con `filepath`

```go
// Caminar directorio eficientemente
filepath.WalkDir("/src", func(path string, d fs.DirEntry, err error) error {
    if d.IsDir() && d.Name() == ".git" {
        return filepath.SkipDir // Optimización
    }
    if strings.HasSuffix(path, ".go") {
        fmt.Println("Go file:", path)
    }
    return nil
})
```

---

## `path` — Rutas Lógicas

* **Siempre usa `/`:** Independiente del SO
* **Para URLs e imports:** No para archivos locales
* **Operaciones de texto:** `path.Clean()`, `path.Base()`, `path.Dir()`

---

## Ejemplo: `path` para URLs

```go
// Para URLs o rutas lógicas
importPath := path.Join("github.com", "user", "repo")
// Siempre: github.com/user/repo (nunca usa \)

// Limpiar rutas de URLs
cleaned := path.Clean("/a/b/../c/") // "/a/c"

// Obtener componentes
dir := path.Dir("/users/home/file.txt")  // "/users/home"
base := path.Base("/users/home/file.txt") // "file.txt"
```

---

## Regla Simple

**Usa `filepath` para archivos locales**
**Usa `path` para URLs/imports**

---

## `embed` — Archivos Dentro del Binario

Go 1.16+ permite **incluir archivos estáticos en el ejecutable compilado**.

Elimina dependencias externas en producción.

---

## Ventajas de `embed`

* **Distribución simple:** Un solo binario contiene todo
* **Sin problemas de rutas:** Los archivos siempre están disponibles
* **Inmutable:** No pueden ser modificados accidentalmente
* **Implementa `fs.FS`:** Compatible con todo el ecosistema

---

## Ejemplo: Embediendo archivos

```go
package main
import "embed"

//go:embed templates/*.html static/css static/js
var assets embed.FS

//go:embed config.yaml
var configData []byte

func main() {
    // Usar con http.FileServer
    http.Handle("/static/", http.FileServer(http.FS(assets)))
    
    // O leer directamente
    tmplData, _ := assets.ReadFile("templates/index.html")
}
```

---

## `testing/fstest` — Tests sin I/O

Proporciona `MapFS`: un sistema de archivos en memoria perfecto para testing.

---

## Ejemplo: Testing con `fstest`

```go
func TestFileProcessor(t *testing.T) {
    // Sistema de archivos de prueba
    testFS := fstest.MapFS{
        "config.json": &fstest.MapFile{
            Data: []byte(`{"debug": true}`),
        },
        "data/users.csv": &fstest.MapFile{
            Data: []byte("id,name\n1,Alice"),
            ModTime: time.Now(),
        },
    }
    
    // Probar función que acepta fs.FS
    result, err := ProcessFiles(testFS)
    if err != nil {
        t.Fatal(err)
    }
}
```

---

# Procesos y Control del Sistema

---

## Herramientas de Control

Go proporciona control completo sobre:
- Procesos externos
- Señales del sistema operativo

Fundamental para herramientas de sistema y servicios.

---

## `os/exec` — Ejecutando Comandos

La forma idiomática de ejecutar programas externos desde Go.

---

## Características de `os/exec`

* **Constructor `Command()`:** Argumentos seguros
* **Pipes automáticos:** Sin archivos temporales
* **Control de contexto:** Timeout y cancelación
* **Streaming:** Salida en tiempo real

---

## Ejemplo: Ejecución simple

```go
// Captura de salida
out, err := exec.Command("git", "status", "--short").Output()
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(out))

// Ejecución con error checking
cmd := exec.Command("ls", "-la")
if err := cmd.Run(); err != nil {
    log.Fatal(err)
}
```

---

## Comando con timeout

```go
// Comando con límite de tiempo
ctx, cancel := context.WithTimeout(
    context.Background(), 
    5*time.Second,
)
defer cancel()

cmd := exec.CommandContext(ctx, "slow-process")
if err := cmd.Run(); err != nil {
    // Timeout o error
    log.Printf("Error: %v", err)
}
```

---

## Pipeline con pipes

```go
// Pipeline complejo
cmd := exec.Command("grep", "error")
cmd.Stdin = strings.NewReader(logData)

var stderr bytes.Buffer
cmd.Stderr = &stderr

output, err := cmd.Output()
if err != nil {
    log.Printf("grep stderr: %s", stderr.String())
}

fmt.Println("Matches:", string(output))
```

---

## Streaming en tiempo real

```go
cmd := exec.Command("tail", "-f", "/var/log/app.log")
stdout, _ := cmd.StdoutPipe()
cmd.Start()

scanner := bufio.NewScanner(stdout)
for scanner.Scan() {
    line := scanner.Text()
    processLine(line)
}
```

---

## Seguridad en `os/exec`

**IMPORTANTE:** Siempre usa `Command()` con argumentos separados.

```go
// ✅ CORRECTO
exec.Command("rm", userInput)

// ❌ PELIGROSO - inyección de comandos
exec.Command("sh", "-c", "rm " + userInput)
```

---

## `os/signal` — Manejo de Señales

Las señales son el mecanismo de IPC más básico en Unix.

---

## Casos de uso comunes

* **Shutdown graceful:** SIGINT (Ctrl+C) y SIGTERM
* **Recarga de config:** SIGHUP
* **Rotación de logs:** SIGUSR1/SIGUSR2
* **Operaciones custom:** Cualquier señal

---

## Patrón moderno con context

```go
// Go 1.16+
ctx, stop := signal.NotifyContext(
    context.Background(), 
    os.Interrupt,    // SIGINT (Ctrl+C)
    syscall.SIGTERM, // Kill elegante
)
defer stop()

// Esperar señal
<-ctx.Done()
fmt.Println("Señal recibida, cerrando...")
```

---

## Servidor con graceful shutdown

```go
server := &http.Server{Addr: ":8080"}

// Goroutine para manejar señales
go func() {
    <-ctx.Done()
    
    // Dar 10 segundos para cerrar conexiones
    shutdownCtx, cancel := context.WithTimeout(
        context.Background(), 
        10*time.Second,
    )
    defer cancel()
    
    server.Shutdown(shutdownCtx)
}()

server.ListenAndServe()
```

---

## Múltiples señales

```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, 
    syscall.SIGHUP, 
    syscall.SIGUSR1,
)

go func() {
    for sig := range sigChan {
        switch sig {
        case syscall.SIGHUP:
            reloadConfig()
        case syscall.SIGUSR1:
            rotateLogs()
        }
    }
}()
```

---

## `syscall` — Bajo Nivel

Acceso directo a las llamadas del sistema.

**⚠️ DEPRECADO:** Está congelado. Usa `golang.org/x/sys` para código nuevo.

---

## Cuándo usar `syscall`

* **Control de grupos de procesos**
* **Atributos de proceso avanzados**
* **Operaciones no portables**

---

## Ejemplo: Grupos de procesos

```go
// Crear proceso en nuevo grupo
cmd := exec.Command("start-service.sh")
cmd.SysProcAttr = &syscall.SysProcAttr{
    Setpgid: true,  // Nuevo grupo
    Pgid:    0,     // Usar PID como PGID
}
cmd.Start()

// Matar todo el grupo
syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM)
```

---

## Límites de recursos

```go
// Cambiar límites (Unix)
var rLimit syscall.Rlimit
syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)

// Aumentar límite de archivos abiertos
rLimit.Cur = 4096  
syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
```

---

# Usuario, Entorno y Tiempo

---

## `os/user` — Información de Usuarios

Información sobre usuarios del SO de forma portable.

---

## Capacidades de `os/user`

* **Usuario actual:** UID, GID, username, home
* **Búsqueda de usuarios:** Por nombre o ID
* **Grupos:** Información de grupos
* **Portable:** Unix y Windows

---

## Ejemplo: Usuario actual

```go
current, err := user.Current()
if err == nil {
    fmt.Printf("Usuario: %s\n", current.Username)
    fmt.Printf("Home: %s\n", current.HomeDir)
    fmt.Printf("UID: %s\n", current.Uid)
    fmt.Printf("GID: %s\n", current.Gid)
}
```

---

## Buscar usuarios y expandir paths

```go
// Buscar otro usuario
u, err := user.Lookup("postgres")
if err == nil {
    fmt.Printf("PostgreSQL home: %s\n", u.HomeDir)
}

// Expandir ~ en paths
func expandHome(path string) string {
    if strings.HasPrefix(path, "~/") {
        if u, _ := user.Current(); u != nil {
            return filepath.Join(u.HomeDir, path[2:])
        }
    }
    return path
}
```

---

## `time` — Temporización

Fundamental para timeouts, scheduling, y performance.

---

## Conceptos clave de `time`

* **`time.Time`:** Momento específico con zona horaria
* **`time.Duration`:** Intervalo (int64 nanosegundos)
* **Timers vs Tickers:** Una vez vs repetitivo
* **Monotonic clocks:** Medición precisa

---

## Timeout pattern

```go
select {
case result := <-doWork():
    processResult(result)
    
case <-time.After(5 * time.Second):
    return ErrTimeout
}
```

---

## Tareas periódicas

```go
ticker := time.NewTicker(1 * time.Minute)
defer ticker.Stop()

for {
    select {
    case <-ticker.C:
        collectMetrics()
        
    case <-ctx.Done():
        return
    }
}
```

---

## Medición precisa

```go
// ✅ Usa monotonic clock
start := time.Now()
expensiveOperation()
elapsed := time.Since(start)

// Formateo de tiempos
const layout = "2006-01-02 15:04:05"
t, _ := time.Parse(layout, "2024-03-15 14:30:00")
formatted := t.Format(time.RFC3339)
```

---

# Networking

---

## Stack de red en Go

Go tiene soporte de red excepcional:
- Sockets raw
- TCP/UDP
- Unix sockets
- HTTP/2 automático

---

## `net` — Networking de Bajo Nivel

Primitivas de red portables.

---

## Capacidades de `net`

* **Listeners y Dialers:** Cliente/servidor
* **Múltiples protocolos:** TCP, UDP, Unix
* **DNS integrado:** Con cache del SO
* **Interfaces de red:** Enumerar y consultar

---

## Servidor TCP

```go
listener, err := net.Listen("tcp", ":8080")
if err != nil {
    log.Fatal(err)
}
defer listener.Close()

for {
    conn, err := listener.Accept()
    if err != nil {
        continue
    }
    go handleConnection(conn)
}
```

---

## Cliente con timeout

```go
// Conexión con límite de tiempo
conn, err := net.DialTimeout(
    "tcp", 
    "api.example.com:443", 
    5*time.Second,
)
if err != nil {
    log.Fatal(err)
}
defer conn.Close()
```

---

## Unix domain sockets

```go
// IPC local rápido
listener, err := net.Listen("unix", "/tmp/app.sock")
if err != nil {
    log.Fatal(err)
}
defer listener.Close()

// Cliente
conn, _ := net.Dial("unix", "/tmp/app.sock")
```

---

## Interfaces y DNS

```go
// Enumerar interfaces de red
interfaces, _ := net.Interfaces()
for _, iface := range interfaces {
    addrs, _ := iface.Addrs()
    fmt.Printf("%s: %v\n", iface.Name, addrs)
}

// Resolver DNS
ips, _ := net.LookupIP("example.com")
mx, _ := net.LookupMX("example.com")
```

---

## `net/netip` — IPs Modernas

Go 1.18+ introduce tipos **value-based** más eficientes.

---

## Ventajas de `netip`

* **Comparable:** Keys en maps
* **Inmutable:** Thread-safe sin locks
* **Eficiente:** 24 bytes vs slice
* **IPv6 zones:** Soporte completo
* **Validación:** Zero values = inválido

---

## Ejemplo: IPs comparables

```go
// Parsing seguro
addr, err := netip.ParseAddr("192.168.1.1")
if !addr.IsValid() {
    return errors.New("IP inválida")
}

// Comparación directa
if addr == netip.MustParseAddr("192.168.1.1") {
    fmt.Println("IPs idénticas")
}

// Uso en maps
connections := make(map[netip.AddrPort]net.Conn)
addrPort := netip.MustParseAddrPort("[::1]:8080")
connections[addrPort] = conn
```

---

## Subredes con `netip`

```go
// Prefijos y subredes
prefix := netip.MustParsePrefix("10.0.0.0/8")

if prefix.Contains(addr) {
    fmt.Println("La IP está en la subred")
}

// Iterar sobre rango
for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
    // Procesar cada IP
}
```

---

## `net/http` — HTTP Production-Ready

Servidor HTTP completo sin dependencias externas.

---

## Características del servidor

* **Multiplexer integrado:** Routing básico
* **HTTP/2 automático:** Con TLS
* **Graceful shutdown:** Go 1.8+
* **Context integration:** Por request

---

## Servidor configurado

```go
mux := http.NewServeMux()
mux.HandleFunc("/health", healthCheck)
mux.Handle("/api/", http.StripPrefix("/api", apiHandler))

server := &http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
    IdleTimeout:  60 * time.Second,
}

server.ListenAndServe()
```

---

## Cliente HTTP avanzado

```go
client := &http.Client{
    Timeout: 30 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}

// Request con context
ctx, cancel := context.WithTimeout(
    context.Background(), 
    5*time.Second,
)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := client.Do(req)
```

---

# Logging y Observabilidad

---

## Tres niveles de logging

1. **`log`** — Básico
2. **`log/slog`** — Estructurado
3. **`log/syslog`** — Sistema

---

## `log` — Simple y Directo

Minimalista pero suficiente para muchas aplicaciones.

---

## Características de `log`

* **Thread-safe:** Uso concurrente
* **Formato configurable:** Timestamps, archivo, línea
* **Output flexible:** `io.Writer`
* **Fatal/Panic:** Errores críticos

---

## Configuración de `log`

```go
// Configuración básica
log.SetFlags(log.LstdFlags | log.Lshortfile)
log.SetPrefix("[APP] ")

// Logger personalizado
errorLog := log.New(
    os.Stderr, 
    "ERROR: ", 
    log.LstdFlags|log.Llongfile,
)

// Output múltiple
logFile, _ := os.OpenFile("app.log", 
    os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
log.SetOutput(io.MultiWriter(os.Stdout, logFile))
```

---

## `log/slog` — Estructurado Moderno

Go 1.21+ responde a la necesidad de logging estructurado.

---

## Ventajas de `slog`

* **Niveles nativos:** Debug, Info, Warn, Error
* **Campos estructurados:** Key-value tipados
* **Handlers pluggables:** JSON, Text, custom
* **Performance:** Zero-allocation
* **Context-aware:** Integración natural

---

## Configuración de `slog`

```go
// Handler JSON para producción
jsonHandler := slog.NewJSONHandler(os.Stdout, 
    &slog.HandlerOptions{
        Level:     slog.LevelInfo,
        AddSource: true,
    },
)

logger := slog.New(jsonHandler)
slog.SetDefault(logger)
```

---

## Logging estructurado

```go
// Campos estructurados
slog.Info("request processed",
    slog.String("method", "GET"),
    slog.String("path", "/api/users"),
    slog.Duration("latency", 250*time.Millisecond),
    slog.Int("status", 200),
)

// Logger con contexto
userLogger := logger.With(
    slog.String("user_id", userID),
    slog.String("session", sessionID),
)
```

---

## Grupos de atributos

```go
slog.Info("database query",
    slog.Group("db",
        slog.String("query", "SELECT * FROM users"),
        slog.Duration("duration", 50*time.Millisecond),
    ),
    slog.Group("cache",
        slog.Bool("hit", false),
    ),
)
```

---

## `log/syslog` — Unix/Linux

Integración con la infraestructura de logging del sistema.

---

## Uso de syslog

```go
// Conectar a syslog local
writer, err := syslog.New(
    syslog.LOG_INFO|syslog.LOG_DAEMON, 
    "myapp",
)
defer writer.Close()

// Diferentes niveles
writer.Info("Application started")
writer.Warning("Config not found")
writer.Err("Database connection failed")

// Syslog remoto
remote, _ := syslog.Dial("tcp", "logserver:514",
    syslog.LOG_WARNING|syslog.LOG_DAEMON, "myapp")
```

---

# Runtime y Métricas

---

## Introspección del programa

Go expone información detallada sobre el runtime.

---

## `runtime` — Control del Runtime

Información y control sobre el runtime de Go.

---

## Información del sistema

```go
// Información del entorno
fmt.Printf("OS: %s\n", runtime.GOOS)
fmt.Printf("Arch: %s\n", runtime.GOARCH)
fmt.Printf("CPUs: %d\n", runtime.NumCPU())
fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
```

---

## Control de memoria

```go
var m runtime.MemStats
runtime.ReadMemStats(&m) // ⚠️ Stop-the-world

fmt.Printf("Alloc: %v MB\n", m.Alloc/1024/1024)
fmt.Printf("TotalAlloc: %v MB\n", m.TotalAlloc/1024/1024)
fmt.Printf("Sys: %v MB\n", m.Sys/1024/1024)
fmt.Printf("NumGC: %v\n", m.NumGC)

// Forzar GC
runtime.GC()
```

---

## Stack traces

```go
// Stack trace completo
buf := make([]byte, 1<<16)
stackSize := runtime.Stack(buf, true)
fmt.Printf("%s\n", buf[:stackSize])

// Información de caller
_, file, line, ok := runtime.Caller(1)
if ok {
    fmt.Printf("Called from %s:%d\n", file, line)
}
```

---

## `runtime/debug` — Build Info

Información de compilación y control de debugging.

---

## Build information

```go
info, ok := debug.ReadBuildInfo()
if ok {
    fmt.Printf("Go Version: %s\n", info.GoVersion)
    fmt.Printf("Module: %s\n", info.Main.Path)
    
    for _, dep := range info.Deps {
        fmt.Printf("Dep: %s@%s\n", dep.Path, dep.Version)
    }
    
    for _, setting := range info.Settings {
        if setting.Key == "vcs.revision" {
            fmt.Printf("Git: %s\n", setting.Value)
        }
    }
}
```

---

## Control de GC

```go
// Ajustar agresividad del GC
debug.SetGCPercent(50) // Default: 100

// Límite de memoria (Go 1.19+)
debug.SetMemoryLimit(1 << 30) // 1GB

// Stack trace en panic
defer func() {
    if r := recover(); r != nil {
        fmt.Fprintf(os.Stderr, "Panic: %v\n%s\n", 
            r, debug.Stack())
    }
}()
```

---

## `runtime/metrics` — Métricas Estables

Go 1.16+ ofrece métricas sin stop-the-world.

---

## Ventajas de metrics

* **No detiene el mundo**
* **API estable**
* **Extensible**

---

## Leer métricas

```go
// Métricas disponibles
descs := metrics.All()
for _, desc := range descs {
    fmt.Printf("%s: %s\n", desc.Name, desc.Description)
}

// Leer específicas
samples := []metrics.Sample{
    {Name: "/memory/classes/heap/free:bytes"},
    {Name: "/gc/cycles/total:gc-cycles"},
    {Name: "/sched/goroutines:goroutines"},
}

metrics.Read(samples)
for _, sample := range samples {
    fmt.Printf("%s = %v\n", sample.Name, sample.Value)
}
```

---

# Formatos Binarios

---

## `plugin` — Carga Dinámica

Cargar `.so` en tiempo de ejecución (Linux/macOS).

---

## Crear un plugin

```go
// plugin.go - compilar con:
// go build -buildmode=plugin
package main

func Greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

var Version = "1.0.0"
```

---

## Cargar plugin

```go
// Aplicación principal
p, err := plugin.Open("greeting.so")
if err != nil {
    panic(err)
}

// Buscar función
greetSym, _ := p.Lookup("Greet")
greetFunc := greetSym.(func(string) string)
result := greetFunc("World")

// Buscar variable
versionSym, _ := p.Lookup("Version")
version := *versionSym.(*string)
```

---

## `debug/elf` — Linux/Unix

```go
file, _ := elf.Open("/usr/bin/ls")
defer file.Close()

// Arquitectura
fmt.Printf("Class: %s\n", file.Class)
fmt.Printf("Machine: %s\n", file.Machine)

// Símbolos
symbols, _ := file.Symbols()
for _, sym := range symbols {
    if strings.Contains(sym.Name, "main") {
        fmt.Printf("Found: %s at 0x%x\n", 
            sym.Name, sym.Value)
    }
}
```

---

## `debug/pe` — Windows

```go
file, _ := pe.Open("app.exe")
defer file.Close()

// Imports
imports, _ := file.ImportedSymbols()
for _, imp := range imports {
    fmt.Println("Import:", imp)
}

// Tipo de archivo
if file.Characteristics&pe.IMAGE_FILE_DLL != 0 {
    fmt.Println("Es una DLL")
}
```

---

## `debug/macho` — macOS

```go
file, _ := macho.Open("/usr/bin/go")
defer file.Close()

// CPU info
fmt.Printf("CPU: %s\n", file.Cpu)

// Load commands
for _, load := range file.Loads {
    fmt.Printf("Load command: %T\n", load)
}
```

---

## `debug/dwarf` — Debug Info

```go
// Abrir ELF con DWARF
elfFile, _ := elf.Open("app")
dwarf, _ := elfFile.DWARF()

// Leer debug info
reader := dwarf.Reader()
for {
    entry, _ := reader.Next()
    if entry == nil {
        break
    }
    
    if entry.Tag == dwarf.TagCompileUnit {
        // Procesar unidad de compilación
    }
}
```

---

# Empaquetado y Compresión

---

## `archive/tar` — Formato Unix

TAR mantiene permisos y metadata.

---

## Crear TAR

```go
buf := new(bytes.Buffer)
tw := tar.NewWriter(buf)
defer tw.Close()

// Agregar archivo
header := &tar.Header{
    Name:     "dir/file.txt",
    Mode:     0644,
    Size:     int64(len(data)),
    ModTime:  time.Now(),
    Typeflag: tar.TypeReg,
}
tw.WriteHeader(header)
tw.Write(data)
```

---

## Leer TAR

```go
tr := tar.NewReader(reader)
for {
    header, err := tr.Next()
    if err == io.EOF {
        break
    }
    
    fmt.Printf("File: %s (%d bytes)\n", 
        header.Name, header.Size)
    
    if header.Typeflag == tar.TypeReg {
        data, _ := io.ReadAll(tr)
        // Procesar archivo
    }
}
```

---

## TAR.GZ

```go
// Crear TAR.GZ
gzFile, _ := os.Create("backup.tar.gz")
gzWriter := gzip.NewWriter(gzFile)
tarWriter := tar.NewWriter(gzWriter)

// Agregar archivos...

tarWriter.Close()
gzWriter.Close()
gzFile.Close()
```

---

## `archive/zip` — Multiplataforma

ZIP soporta compresión por archivo.

---

## Crear ZIP

```go
buf := new(bytes.Buffer)
zw := zip.NewWriter(buf)
defer zw.Close()

// Archivo simple
fw, _ := zw.Create("file.txt")
fw.Write([]byte("contenido"))

// Con compresión específica
header := &zip.FileHeader{
    Name:   "compressed.txt",
    Method: zip.Deflate,
}
fw, _ = zw.CreateHeader(header)
fw.Write([]byte("datos"))
```

---

## Leer ZIP

```go
reader, _ := zip.NewReader(
    bytes.NewReader(buf.Bytes()), 
    int64(buf.Len()),
)

for _, file := range reader.File {
    fmt.Printf("%s: %d bytes\n", 
        file.Name, 
        file.UncompressedSize64,
    )
    
    rc, _ := file.Open()
    data, _ := io.ReadAll(rc)
    rc.Close()
}
```

---

## `compress/gzip` — Streams

GZIP para datos en tránsito.

---

## Comprimir/Descomprimir

```go
// Comprimir
var buf bytes.Buffer
gw := gzip.NewWriter(&buf)
gw.Write([]byte("datos"))
gw.Close()

// Descomprimir
gr, _ := gzip.NewReader(&buf)
data, _ := io.ReadAll(gr)
gr.Close()
```

---

## HTTP con compresión

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Detectar soporte
    if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
        w.Header().Set("Content-Encoding", "gzip")
        gz := gzip.NewWriter(w)
        defer gz.Close()
        // Usar gz como writer
    }
    
    w.Write([]byte("respuesta"))
})
```

---

## Niveles de compresión

```go
// Rápido
gw, _ := gzip.NewWriterLevel(&buf, gzip.BestSpeed)

// Máxima compresión
gw, _ = gzip.NewWriterLevel(&buf, gzip.BestCompression)

// Balance (default)
gw, _ = gzip.NewWriterLevel(&buf, gzip.DefaultCompression)
```

---

## Otros formatos

* **`compress/bzip2`:** Solo lectura
* **`compress/flate`:** DEFLATE raw
* **`compress/lzw`:** LZW (GIF, PDF)
* **`compress/zlib`:** Similar a gzip

---

# Utilidades de Apoyo

---

## `context` — Cancelación

Fundamental para control de timeouts y cancelación.

---

## Patrones de context

```go
// Timeout
ctx, cancel := context.WithTimeout(
    context.Background(), 
    30*time.Second,
)
defer cancel()

// Cancelación manual
ctx, cancel := context.WithCancel(context.Background())
go func() {
    <-signalChan
    cancel()
}()
```

---

## Propagación de context

```go
// Verificar cancelación
for {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        // Continuar
    }
}

// Propagar en funciones
func processData(ctx context.Context, data []byte) error {
    result, err := callAPI(ctx, data)
    if err != nil {
        return err
    }
    return saveResult(ctx, result)
}
```

---

## `io` — Utilidades I/O

```go
// Copiar eficientemente
written, _ := io.Copy(dest, source)

// Limitar lectura
limited := io.LimitReader(reader, 1024*1024)

// Escribir a múltiples destinos
multi := io.MultiWriter(os.Stdout, logFile)

// Inspeccionar mientras se lee
tee := io.TeeReader(response.Body, os.Stdout)
```

---

## `bufio` — I/O con Buffer

```go
// Leer líneas
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    processLine(line)
}

// Writer con buffer
bw := bufio.NewWriter(file)
bw.WriteString(data)
bw.Flush() // ¡No olvidar!

// Peek sin consumir
br := bufio.NewReader(conn)
peek, _ := br.Peek(4)
```

---

## `crypto/x509` — Certificados

```go
// Cargar certificados del sistema
systemPool, _ := x509.SystemCertPool()

// Agregar CA custom
caCert, _ := os.ReadFile("ca.crt")
systemPool.AppendCertsFromPEM(caCert)

// Cliente HTTPS con CAs
client := &http.Client{
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{
            RootCAs: systemPool,
        },
    },
}
```

---

# Paquetes No Estándar Esenciales

---

## `golang.org/x/sys`

Syscalls modernos y mantenidos.

```go
import "golang.org/x/sys/unix"

// File descriptors avanzados
fd, _ := unix.Open("/dev/null", unix.O_RDWR, 0)

// Señales extendidas
unix.Kill(pid, unix.SIGURG)

// Linux específico
unix.Prctl(unix.PR_SET_DUMPABLE, 1, 0, 0, 0)
```

---

## `golang.org/x/term`

Terminal y TTY.

```go
// Leer password sin eco
fmt.Print("Password: ")
password, _ := term.ReadPassword(
    int(os.Stdin.Fd()),
)

// Verificar si es terminal
if term.IsTerminal(int(os.Stdout.Fd())) {
    // Usar colores ANSI
}
```

---

## `fsnotify/fsnotify`

File watching multiplataforma.

```go
watcher, _ := fsnotify.NewWatcher()
defer watcher.Close()

watcher.Add("/path/to/watch")

for {
    select {
    case event := <-watcher.Events:
        if event.Op&fsnotify.Write == fsnotify.Write {
            fmt.Println("Modified:", event.Name)
        }
    case err := <-watcher.Errors:
        log.Println("Error:", err)
    }
}
```

---

# Resumen

---

## Principios Clave

1. **Usa la abstracción correcta**
2. **Prefiere la stdlib**
3. **Contexts everywhere**
4. **Maneja señales**
5. **Estructura tus logs**

---

## Decisiones Comunes

| Necesidad | Usa | No uses |
|-----------|-----|---------|
| Rutas locales | `filepath` | `path` |
| URLs/imports | `path` | `filepath` |
| IPs en maps | `net/netip` | `net.IP` |
| Logging estructurado | `slog` | `log` |
| Syscalls nuevos | `x/sys` | `syscall` |
| Archivos de test | `fstest` | Reales |

---

## Patrón: Graceful Shutdown

```go
ctx, stop := signal.NotifyContext(
    context.Background(), 
    os.Interrupt, 
    syscall.SIGTERM,
)
defer stop()

server := startServer()
<-ctx.Done()

shutdownCtx, _ := context.WithTimeout(
    context.Background(), 
    10*time.Second,
)
server.Shutdown(shutdownCtx)
```

---

## Patrón: Resource Cleanup

```go
// Siempre defer Close()
file, err := os.Open(name)
if err != nil {
    return err
}
defer file.Close()

// Error wrapping
if err != nil {
    return fmt.Errorf("processing %s: %w", filename, err)
}
```

---
layout: end
---

# ¡Gracias!

---

## Recursos Adicionales

* **Documentación:** https://pkg.go.dev/std
* **Go by Example:** https://gobyexample.com
* **Effective Go:** https://go.dev/doc/effective_go
* **Go Blog:** https://go.dev/blog

---

## Práctica Recomendada

Construye herramientas CLI:

1. Monitor de archivos con `fsnotify`
2. Servidor HTTP con graceful shutdown
3. Backup con `tar` y `gzip`
4. Analizador de binarios con `debug/elf`

La biblioteca estándar de Go es tu mejor aliada.