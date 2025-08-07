---
theme: default
background: https://images.unsplash.com/photo-1555066931-4365d14bab8c?q=80&w=2070
title: Go como Lenguaje para Sistemas Operativos
info: |
  ## Go para Programación de Sistemas
  Una exploración de las características que hacen a Go ideal para desarrollo de sistemas operativos.
  
  Aprende más en [golang.org](https://golang.org)
class: text-center
drawings:
  persist: false
transition: slide-left
mdc: true
---

# Go como Lenguaje para Sistemas Operativos

Explorando las características que hacen a Go ideal para programación de sistemas

<div @click="$slidev.nav.next" class="mt-12 py-1" hover:bg="white op-10">
  Presiona Espacio para continuar <carbon:arrow-right />
</div>

---
transition: fade-out
---

# ¿Por qué Go para Sistemas?

Go fue diseñado con la programación de sistemas en mente

<v-clicks>

- 🚀 **Rendimiento** - Compilado a código máquina nativo
- 🔧 **Control de memoria** - Recolector de basura eficiente y predecible
- 🎯 **Concurrencia nativa** - Goroutines y canales integrados en el lenguaje
- 📦 **Binarios estáticos** - Distribución simple sin dependencias
- 🛡️ **Seguridad de tipos** - Sistema de tipos fuerte y estático
- ⚡ **Compilación rápida** - Tiempos de compilación extremadamente rápidos

</v-clicks>

---
layout: two-cols
---

# Características Clave

<v-click>

## <span v-mark.highlight.yellow="1">Gestión de Memoria</span>

Go ofrece un equilibrio único entre control y automatización

</v-click>

<v-click>

- Recolector de basura concurrente
- Control sobre asignación de memoria
- Escape analysis para optimización
- Pools de memoria reutilizables

</v-click>

::right::

<v-click>

```go
// Pool de objetos reutilizables
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 4096)
    },
}

func processData() {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf)
    
    // Usar el buffer...
}
```

</v-click>

---

# Concurrencia: El Superpoder de Go

<span v-mark.underline.orange>Goroutines</span> y <span v-mark.circle.blue="2">Canales</span> hacen la concurrencia simple y segura

<v-clicks>

```go
func main() {
    // Canal para comunicación
    results := make(chan int, 10)
    
    // Lanzar múltiples goroutines
    for i := 0; i < 10; i++ {
        go worker(i, results)
    }
    
    // Recolectar resultados
    for i := 0; i < 10; i++ {
        fmt.Println(<-results)
    }
}

func worker(id int, results chan<- int) {
    // Trabajo pesado aquí...
    results <- id * 2
}
```

</v-clicks>

---
layout: image-right
image: https://images.unsplash.com/photo-1518432031352-d6fc5c10da5a?q=80&w=1974
---

# Interfaz con el Sistema

Go facilita la interacción con el sistema operativo

<v-click>

## <span v-mark.box.green>Syscalls directas</span>

```go
import "syscall"

fd, err := syscall.Open("/dev/null", 
    syscall.O_RDWR, 0)
if err != nil {
    panic(err)
}
defer syscall.Close(fd)
```

</v-click>

<v-click>

## Integración con C

```go
// #include <stdio.h>
import "C"

func main() {
    C.printf(C.CString("Hola desde C\n"))
}
```

</v-click>

---

# Proyectos de Sistemas en Go

Ejemplos exitosos de sistemas construidos con Go

<div class="grid grid-cols-2 gap-4 mt-8">

<v-click>

<div class="p-4 border rounded">

### <span v-mark.highlight.yellow>Docker</span>
Sistema de contenedores que revolucionó el deployment

- Motor de contenedores
- Gestión de namespaces
- Control de cgroups

</div>

</v-click>

<v-click>

<div class="p-4 border rounded">

### <span v-mark.highlight.blue>Kubernetes</span>
Orquestador de contenedores

- Gestión de clusters
- Scheduling de procesos
- Control de recursos

</div>

</v-click>
</div>

---

# Proyectos de Sistemas en Go (cont.)


<div class="grid grid-cols-2 gap-4 mt-8">

<v-click>

<div class="p-4 border rounded">

### <span v-mark.highlight.green>TinyGo</span>
Compilador Go para sistemas embebidos

- Microcontroladores
- WebAssembly
- Dispositivos IoT

</div>

</v-click>

<v-click>

<div class="p-4 border rounded">

### <span v-mark.highlight.orange>gVisor</span>
Kernel de aplicación en espacio de usuario

- Sandbox de seguridad
- Compatibilidad con Linux
- Aislamiento de procesos

</div>

</v-click>

</div>

---

# Gestión de Recursos del Sistema

````md magic-move {lines: true}
```go
// Paso 1: Estructura básica
type ResourceManager struct {
    cpu    int
    memory int64
}
```

```go
// Paso 2: Añadir métodos de control
type ResourceManager struct {
    cpu    int
    memory int64
    mutex  sync.RWMutex
}

func (rm *ResourceManager) SetLimits(cpu int, mem int64) {
    rm.mutex.Lock()
    defer rm.mutex.Unlock()
    rm.cpu = cpu
    rm.memory = mem
}
```

```go
// Paso 3: Implementar monitoreo
type ResourceManager struct {
    cpu     int
    memory  int64
    mutex   sync.RWMutex
    monitor chan Stats
}

func (rm *ResourceManager) StartMonitoring(ctx context.Context) {
    ticker := time.NewTicker(1 * time.Second)
    go func() {
        for {
            select {
            case <-ticker.C:
                rm.monitor <- rm.getCurrentStats()
            case <-ctx.Done():
                return
            }
        }
    }()
}
```
````

---

# Manejo de Señales del Sistema

Go proporciona manejo elegante de señales POSIX

<v-clicks>

```go
func main() {
    // Canal para señales
    sigChan := make(chan os.Signal, 1)
    
    // Registrar señales a capturar
    signal.Notify(sigChan, 
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGHUP)
    
    // Goroutine para manejar señales
    go func() {
        for sig := range sigChan {
            switch sig {
            case syscall.SIGINT, syscall.SIGTERM:
                fmt.Println("Apagado graceful...")
                cleanup()
                os.Exit(0)
            case syscall.SIGHUP:
                fmt.Println("Recargando configuración...")
                reloadConfig()
            }
        }
    }()
    
    // Continuar con la lógica principal...
}
```

</v-clicks>

---
layout: two-cols
---

# Ventajas vs Otros Lenguajes

## <span v-mark.underline.green="1">Go vs C/C++</span>

<v-clicks at="2">

✅ **Gestión de memoria automática**
- Sin punteros colgantes
- Sin fugas de memoria manuales

✅ **Concurrencia más simple**
- Goroutines vs threads POSIX
- Canales vs mutexes/semáforos

✅ **Tooling integrado**
- Formatter, linter, test runner

</v-clicks>

::right::

## <span v-mark.underline.blue="4">Go vs Rust</span>

<v-clicks at="5">

✅ **Curva de aprendizaje suave**
- Sintaxis simple y clara
- Menos conceptos complejos

✅ **Compilación más rápida**
- Segundos vs minutos

✅ **GC vs Borrow Checker**
- Trade-off: simplicidad vs control

</v-clicks>

---

# Casos de Uso Ideales

<div class="grid grid-cols-2 gap-6 mt-8">

<v-click>

### <span v-mark.box.orange>✅ Perfecto para</span>

- Servicios de red y APIs
- Herramientas de línea de comandos
- Sistemas distribuidos
- Middleware y proxies
- Servicios de infraestructura
- Aplicaciones cloud-native

</v-click>

<v-click>

### <span v-mark.box.red="2">⚠️ Considerar alternativas</span>

- Kernels de SO (necesita runtime)
- Drivers de dispositivos
- Sistemas hard real-time
- Aplicaciones con GC crítico
- Sistemas embebidos muy limitados

</v-click>

</div>

---

# Optimizaciones para Sistemas

Técnicas avanzadas para maximizar el rendimiento

<v-clicks>

## 1. Uso de unsafe para operaciones críticas

```go
import "unsafe"

// Conversión rápida sin copia
func BytesToString(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}
```

## 2. Control del GC

```go
// Deshabilitar GC temporalmente
debug.SetGCPercent(-1)
defer debug.SetGCPercent(100)

// Forzar recolección
runtime.GC()
```

</v-clicks>

---

# Herramientas de Desarrollo

Go incluye herramientas poderosas para desarrollo de sistemas

<div class="grid grid-cols-3 gap-4 mt-6">

<v-click>

<div class="text-center p-4">
<div class="text-4xl mb-2">🔍</div>

### pprof
Profiling de CPU y memoria

```bash
go tool pprof cpu.prof
```
</div>

</v-click>

<v-click>

<div class="text-center p-4">
<div class="text-4xl mb-2">🏃</div>

### race detector
Detección de condiciones de carrera

```bash
go run -race main.go
```
</div>

</v-click>

<v-click>

<div class="text-center p-4">
<div class="text-4xl mb-2">📊</div>

### trace
Visualización de ejecución

```bash
go tool trace trace.out
```
</div>

</v-click>

</div>

---

# Ejemplo: Mini Sistema de Archivos

```go {all|3-8|10-19|21-28} 
package main

type FileSystem struct {
    root *Node
    mu   sync.RWMutex
}

type Node struct {
    name     string
    isDir    bool
    children map[string]*Node
    content  []byte
    parent   *Node
    mode     os.FileMode
    modTime  time.Time
}

func (fs *FileSystem) CreateFile(path string, content []byte) error {
    fs.mu.Lock()
    defer fs.mu.Unlock()
    
    dir, name := filepath.Split(path)
    parent := fs.navigateTo(dir)
    
    parent.children[name] = &Node{
        name:    name,
        content: content,
        parent:  parent,
        modTime: time.Now(),
    }
    return nil
}
```

---
layout: center
class: text-center
---

# Ecosistema y Comunidad

<div class="grid grid-cols-2 gap-8 mt-8">

<v-click>

### <span v-mark.highlight.blue>Bibliotecas de Sistemas</span>

- **golang.org/x/sys** - Syscalls multiplataforma
- **github.com/shirou/gopsutil** - Info del sistema
- **github.com/containerd/containerd** - Runtime de contenedores
- **github.com/coreos/go-systemd** - Integración systemd

</v-click>

<v-click>

### <span v-mark.highlight.green="2">Recursos de Aprendizaje</span>

- The Go Programming Language (libro)
- Go by Example
- Effective Go
- Blog oficial de Go

</v-click>

</div>

---

# Futuro de Go en Sistemas

<v-clicks>

## 🚀 Desarrollos en progreso

- **GC mejorado** - Latencias aún más bajas
- **WASM support** - Go en el navegador y edge computing
- **TinyGo** - Expansión a más microcontroladores

</v-clicks>

---
layout: center
class: text-center
---

# Conclusiones

<div class="mt-8">

<v-click>

## <span v-mark.circle.orange>Go es una excelente opción para programación de sistemas modernos</span>

</v-click>

<v-click>

### Combina:

Rendimiento más cercano a C

Simplicidad de lenguajes de alto nivel

Concurrencia de primera clase

Tooling excepcional

Comunidad activa

</v-click>
</div>