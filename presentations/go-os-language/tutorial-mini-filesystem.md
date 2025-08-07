# Tutorial: Construyendo un Mini Sistema de Archivos en Go

## Introducción

En este tutorial, implementaremos un sistema de archivos simple en memoria usando Go. Este proyecto es ideal para entender cómo funcionan los sistemas de archivos a nivel conceptual y cómo Go maneja estructuras de datos complejas, concurrencia y operaciones de E/S.

## ¿Qué construiremos?

Un sistema de archivos virtual en memoria que permite:
- Crear archivos y directorios
- Leer y escribir contenido
- Navegar por la estructura de directorios
- Listar contenidos
- Eliminar archivos y directorios
- Operaciones concurrentes seguras
- Metadatos básicos (permisos, timestamps, tamaño)

## Conceptos clave

### 1. Estructura de árbol
Los sistemas de archivos son estructuras de árbol donde:
- El **nodo raíz** es el directorio principal (`/`)
- Los **nodos internos** son directorios
- Los **nodos hoja** son archivos

### 2. Inodos (Nodes)
Cada elemento del sistema de archivos tiene:
- **Nombre**: identificador del archivo/directorio
- **Tipo**: archivo o directorio
- **Contenido**: datos (solo archivos)
- **Hijos**: subdirectorios y archivos (solo directorios)
- **Metadatos**: permisos, timestamps, tamaño

### 3. Concurrencia
Múltiples operaciones pueden ocurrir simultáneamente, necesitamos:
- **Mutex de lectura/escritura** para sincronización
- **Operaciones atómicas** para mantener consistencia

## Implementación paso a paso

### Paso 1: Estructura base

```go
package minifs

import (
    "errors"
    "os"
    "path/filepath"
    "strings"
    "sync"
    "time"
)

// NodeType representa el tipo de nodo
type NodeType int

const (
    FileNode NodeType = iota
    DirNode
)

// Node representa un archivo o directorio
type Node struct {
    name     string
    nodeType NodeType
    content  []byte
    children map[string]*Node
    parent   *Node
    
    // Metadatos
    mode     os.FileMode
    modTime  time.Time
    size     int64
    mu       sync.RWMutex
}

// FileSystem representa nuestro sistema de archivos
type FileSystem struct {
    root *Node
    mu   sync.RWMutex
}
```

### Paso 2: Crear el sistema de archivos

```go
// NewFileSystem crea un nuevo sistema de archivos con raíz
func NewFileSystem() *FileSystem {
    root := &Node{
        name:     "/",
        nodeType: DirNode,
        children: make(map[string]*Node),
        mode:     0755,
        modTime:  time.Now(),
    }
    
    return &FileSystem{
        root: root,
    }
}
```

### Paso 3: Navegación de rutas

```go
// parsePath divide una ruta en sus componentes
func (fs *FileSystem) parsePath(path string) []string {
    path = filepath.Clean(path)
    if path == "/" || path == "." {
        return []string{}
    }
    
    path = strings.TrimPrefix(path, "/")
    return strings.Split(path, "/")
}

// navigateTo navega hasta el directorio especificado
func (fs *FileSystem) navigateTo(path string) (*Node, error) {
    parts := fs.parsePath(path)
    current := fs.root
    
    for _, part := range parts {
        current.mu.RLock()
        child, exists := current.children[part]
        current.mu.RUnlock()
        
        if !exists {
            return nil, errors.New("ruta no encontrada: " + path)
        }
        
        if child.nodeType != DirNode {
            return nil, errors.New("no es un directorio: " + part)
        }
        
        current = child
    }
    
    return current, nil
}
```

### Paso 4: Crear directorios

```go
// CreateDir crea un nuevo directorio
func (fs *FileSystem) CreateDir(path string, mode os.FileMode) error {
    fs.mu.Lock()
    defer fs.mu.Unlock()
    
    dir, name := filepath.Split(path)
    if name == "" {
        return errors.New("nombre de directorio vacío")
    }
    
    parent, err := fs.navigateTo(dir)
    if err != nil {
        return err
    }
    
    parent.mu.Lock()
    defer parent.mu.Unlock()
    
    if _, exists := parent.children[name]; exists {
        return errors.New("el directorio ya existe: " + name)
    }
    
    newDir := &Node{
        name:     name,
        nodeType: DirNode,
        children: make(map[string]*Node),
        parent:   parent,
        mode:     mode,
        modTime:  time.Now(),
    }
    
    parent.children[name] = newDir
    parent.modTime = time.Now()
    
    return nil
}

// MkdirAll crea un directorio y todos sus padres si no existen
func (fs *FileSystem) MkdirAll(path string, mode os.FileMode) error {
    parts := fs.parsePath(path)
    current := ""
    
    for _, part := range parts {
        current = filepath.Join(current, part)
        if err := fs.CreateDir(current, mode); err != nil {
            // Si el directorio ya existe, continuamos
            if !strings.Contains(err.Error(), "ya existe") {
                return err
            }
        }
    }
    
    return nil
}
```

### Paso 5: Crear y escribir archivos

```go
// CreateFile crea un nuevo archivo con contenido
func (fs *FileSystem) CreateFile(path string, content []byte, mode os.FileMode) error {
    fs.mu.Lock()
    defer fs.mu.Unlock()
    
    dir, name := filepath.Split(path)
    if name == "" {
        return errors.New("nombre de archivo vacío")
    }
    
    parent, err := fs.navigateTo(dir)
    if err != nil {
        return err
    }
    
    parent.mu.Lock()
    defer parent.mu.Unlock()
    
    if existing, exists := parent.children[name]; exists {
        if existing.nodeType == DirNode {
            return errors.New("ya existe un directorio con ese nombre: " + name)
        }
        // Sobrescribir archivo existente
        existing.mu.Lock()
        existing.content = content
        existing.size = int64(len(content))
        existing.modTime = time.Now()
        existing.mu.Unlock()
        return nil
    }
    
    newFile := &Node{
        name:     name,
        nodeType: FileNode,
        content:  content,
        parent:   parent,
        mode:     mode,
        modTime:  time.Now(),
        size:     int64(len(content)),
    }
    
    parent.children[name] = newFile
    parent.modTime = time.Now()
    
    return nil
}

// WriteFile escribe contenido en un archivo (lo crea si no existe)
func (fs *FileSystem) WriteFile(path string, content []byte) error {
    return fs.CreateFile(path, content, 0644)
}
```

### Paso 6: Leer archivos

```go
// ReadFile lee el contenido de un archivo
func (fs *FileSystem) ReadFile(path string) ([]byte, error) {
    fs.mu.RLock()
    defer fs.mu.RUnlock()
    
    dir, name := filepath.Split(path)
    
    parent, err := fs.navigateTo(dir)
    if err != nil {
        return nil, err
    }
    
    parent.mu.RLock()
    node, exists := parent.children[name]
    parent.mu.RUnlock()
    
    if !exists {
        return nil, errors.New("archivo no encontrado: " + path)
    }
    
    if node.nodeType != FileNode {
        return nil, errors.New("no es un archivo: " + path)
    }
    
    node.mu.RLock()
    defer node.mu.RUnlock()
    
    // Retornar una copia del contenido
    content := make([]byte, len(node.content))
    copy(content, node.content)
    
    return content, nil
}
```

### Paso 7: Listar directorios

```go
// FileInfo representa información de un archivo/directorio
type FileInfo struct {
    Name    string
    Size    int64
    Mode    os.FileMode
    ModTime time.Time
    IsDir   bool
}

// ListDir lista el contenido de un directorio
func (fs *FileSystem) ListDir(path string) ([]FileInfo, error) {
    fs.mu.RLock()
    defer fs.mu.RUnlock()
    
    dir, err := fs.navigateTo(path)
    if err != nil {
        return nil, err
    }
    
    dir.mu.RLock()
    defer dir.mu.RUnlock()
    
    var files []FileInfo
    for _, child := range dir.children {
        info := FileInfo{
            Name:    child.name,
            Size:    child.size,
            Mode:    child.mode,
            ModTime: child.modTime,
            IsDir:   child.nodeType == DirNode,
        }
        files = append(files, info)
    }
    
    return files, nil
}
```

### Paso 8: Eliminar archivos y directorios

```go
// Remove elimina un archivo o directorio vacío
func (fs *FileSystem) Remove(path string) error {
    fs.mu.Lock()
    defer fs.mu.Unlock()
    
    dir, name := filepath.Split(path)
    if name == "" {
        return errors.New("no se puede eliminar la raíz")
    }
    
    parent, err := fs.navigateTo(dir)
    if err != nil {
        return err
    }
    
    parent.mu.Lock()
    defer parent.mu.Unlock()
    
    node, exists := parent.children[name]
    if !exists {
        return errors.New("no existe: " + path)
    }
    
    // Si es directorio, verificar que esté vacío
    if node.nodeType == DirNode && len(node.children) > 0 {
        return errors.New("directorio no vacío: " + path)
    }
    
    delete(parent.children, name)
    parent.modTime = time.Now()
    
    return nil
}

// RemoveAll elimina un archivo o directorio y todo su contenido
func (fs *FileSystem) RemoveAll(path string) error {
    fs.mu.Lock()
    defer fs.mu.Unlock()
    
    if path == "/" || path == "" {
        return errors.New("no se puede eliminar la raíz")
    }
    
    dir, name := filepath.Split(path)
    
    parent, err := fs.navigateTo(dir)
    if err != nil {
        return err
    }
    
    parent.mu.Lock()
    defer parent.mu.Unlock()
    
    if _, exists := parent.children[name]; !exists {
        return errors.New("no existe: " + path)
    }
    
    delete(parent.children, name)
    parent.modTime = time.Now()
    
    return nil
}
```

### Paso 9: Operaciones adicionales

```go
// Exists verifica si una ruta existe
func (fs *FileSystem) Exists(path string) bool {
    fs.mu.RLock()
    defer fs.mu.RUnlock()
    
    if path == "/" || path == "" {
        return true
    }
    
    dir, name := filepath.Split(path)
    
    parent, err := fs.navigateTo(dir)
    if err != nil {
        return false
    }
    
    parent.mu.RLock()
    _, exists := parent.children[name]
    parent.mu.RUnlock()
    
    return exists
}

// Stat obtiene información de un archivo/directorio
func (fs *FileSystem) Stat(path string) (FileInfo, error) {
    fs.mu.RLock()
    defer fs.mu.RUnlock()
    
    if path == "/" || path == "" {
        return FileInfo{
            Name:    "/",
            Mode:    fs.root.mode,
            ModTime: fs.root.modTime,
            IsDir:   true,
        }, nil
    }
    
    dir, name := filepath.Split(path)
    
    parent, err := fs.navigateTo(dir)
    if err != nil {
        return FileInfo{}, err
    }
    
    parent.mu.RLock()
    node, exists := parent.children[name]
    parent.mu.RUnlock()
    
    if !exists {
        return FileInfo{}, errors.New("no existe: " + path)
    }
    
    node.mu.RLock()
    defer node.mu.RUnlock()
    
    return FileInfo{
        Name:    node.name,
        Size:    node.size,
        Mode:    node.mode,
        ModTime: node.modTime,
        IsDir:   node.nodeType == DirNode,
    }, nil
}

// Walk recorre el árbol de archivos
func (fs *FileSystem) Walk(path string, walkFn func(path string, info FileInfo) error) error {
    fs.mu.RLock()
    defer fs.mu.RUnlock()
    
    return fs.walkRecursive(path, fs.root, walkFn)
}

func (fs *FileSystem) walkRecursive(path string, node *Node, walkFn func(string, FileInfo) error) error {
    node.mu.RLock()
    info := FileInfo{
        Name:    node.name,
        Size:    node.size,
        Mode:    node.mode,
        ModTime: node.modTime,
        IsDir:   node.nodeType == DirNode,
    }
    
    children := make([]*Node, 0, len(node.children))
    childNames := make([]string, 0, len(node.children))
    for name, child := range node.children {
        children = append(children, child)
        childNames = append(childNames, name)
    }
    node.mu.RUnlock()
    
    if err := walkFn(path, info); err != nil {
        return err
    }
    
    if node.nodeType == DirNode {
        for i, child := range children {
            childPath := filepath.Join(path, childNames[i])
            if err := fs.walkRecursive(childPath, child, walkFn); err != nil {
                return err
            }
        }
    }
    
    return nil
}
```

## Uso del sistema de archivos

### Ejemplo básico

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    // Crear sistema de archivos
    fs := NewFileSystem()
    
    // Crear estructura de directorios
    if err := fs.MkdirAll("/home/user/documents", 0755); err != nil {
        log.Fatal(err)
    }
    
    // Crear archivos
    content := []byte("Hola, mundo!")
    if err := fs.WriteFile("/home/user/documents/saludo.txt", content); err != nil {
        log.Fatal(err)
    }
    
    // Leer archivo
    data, err := fs.ReadFile("/home/user/documents/saludo.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Contenido: %s\n", data)
    
    // Listar directorio
    files, err := fs.ListDir("/home/user/documents")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Archivos en /home/user/documents:")
    for _, file := range files {
        if file.IsDir {
            fmt.Printf("  [DIR]  %s\n", file.Name)
        } else {
            fmt.Printf("  [FILE] %s (%d bytes)\n", file.Name, file.Size)
        }
    }
    
    // Recorrer todo el sistema de archivos
    fmt.Println("\nÁrbol completo:")
    fs.Walk("/", func(path string, info FileInfo) error {
        indent := strings.Count(path, "/") * 2
        prefix := strings.Repeat(" ", indent)
        
        if info.IsDir {
            fmt.Printf("%s[%s/]\n", prefix, info.Name)
        } else {
            fmt.Printf("%s%s (%d bytes)\n", prefix, info.Name, info.Size)
        }
        return nil
    })
}
```

## Pruebas unitarias

```go
package minifs

import (
    "testing"
    "bytes"
)

func TestFileSystemOperations(t *testing.T) {
    fs := NewFileSystem()
    
    t.Run("CreateDir", func(t *testing.T) {
        err := fs.CreateDir("/test", 0755)
        if err != nil {
            t.Fatalf("Error creando directorio: %v", err)
        }
        
        if !fs.Exists("/test") {
            t.Error("El directorio no fue creado")
        }
    })
    
    t.Run("CreateFile", func(t *testing.T) {
        content := []byte("contenido de prueba")
        err := fs.WriteFile("/test/archivo.txt", content)
        if err != nil {
            t.Fatalf("Error creando archivo: %v", err)
        }
        
        data, err := fs.ReadFile("/test/archivo.txt")
        if err != nil {
            t.Fatalf("Error leyendo archivo: %v", err)
        }
        
        if !bytes.Equal(data, content) {
            t.Error("El contenido no coincide")
        }
    })
    
    t.Run("ListDir", func(t *testing.T) {
        files, err := fs.ListDir("/test")
        if err != nil {
            t.Fatalf("Error listando directorio: %v", err)
        }
        
        if len(files) != 1 {
            t.Errorf("Se esperaba 1 archivo, se encontraron %d", len(files))
        }
    })
    
    t.Run("Remove", func(t *testing.T) {
        err := fs.Remove("/test/archivo.txt")
        if err != nil {
            t.Fatalf("Error eliminando archivo: %v", err)
        }
        
        if fs.Exists("/test/archivo.txt") {
            t.Error("El archivo no fue eliminado")
        }
    })
}

func TestConcurrency(t *testing.T) {
    fs := NewFileSystem()
    fs.MkdirAll("/concurrent/test", 0755)
    
    // Crear múltiples archivos concurrentemente
    done := make(chan bool, 10)
    
    for i := 0; i < 10; i++ {
        go func(n int) {
            path := fmt.Sprintf("/concurrent/test/file%d.txt", n)
            content := []byte(fmt.Sprintf("Archivo %d", n))
            fs.WriteFile(path, content)
            done <- true
        }(i)
    }
    
    // Esperar a que terminen
    for i := 0; i < 10; i++ {
        <-done
    }
    
    // Verificar que todos los archivos fueron creados
    files, err := fs.ListDir("/concurrent/test")
    if err != nil {
        t.Fatalf("Error listando directorio: %v", err)
    }
    
    if len(files) != 10 {
        t.Errorf("Se esperaban 10 archivos, se encontraron %d", len(files))
    }
}
```

## Mejoras posibles

1. **Persistencia**: Serializar el sistema de archivos a disco
2. **Permisos avanzados**: Implementar usuarios y grupos
3. **Enlaces simbólicos**: Añadir soporte para symlinks
4. **Cuotas**: Limitar espacio por usuario/directorio
5. **Búsqueda**: Implementar find/grep
6. **Compresión**: Comprimir archivos grandes automáticamente
7. **Versionado**: Mantener historial de cambios
8. **Montaje**: Montar otros sistemas de archivos
9. **Cache**: Optimizar lecturas frecuentes
10. **Eventos**: Sistema de notificaciones de cambios

## Conclusión

Este mini sistema de archivos demuestra conceptos fundamentales:
- **Estructuras de datos**: Árboles y mapas
- **Concurrencia**: Mutex para operaciones seguras
- **Diseño de API**: Interfaz similar a os/filepath
- **Manejo de errores**: Validación exhaustiva
- **Testing**: Pruebas unitarias y de concurrencia

El código es extensible y puede servir como base para proyectos más complejos como:
- Sistemas de archivos virtuales
- Almacenamiento en memoria para testing
- Backends de almacenamiento distribuido
- Sistemas de caché de archivos