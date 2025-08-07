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
	mode    os.FileMode
	modTime time.Time
	size    int64
	mu      sync.RWMutex
}

// FileSystem representa nuestro sistema de archivos
type FileSystem struct {
	root *Node
	mu   sync.RWMutex
}

// FileInfo representa información de un archivo/directorio
type FileInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

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

	startNode := fs.root
	if path != "/" && path != "" {
		dir, name := filepath.Split(path)
		parent, err := fs.navigateTo(dir)
		if err != nil {
			return err
		}

		parent.mu.RLock()
		node, exists := parent.children[name]
		parent.mu.RUnlock()

		if !exists {
			return errors.New("no existe: " + path)
		}
		startNode = node
	}

	return fs.walkRecursive(path, startNode, walkFn)
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

// AppendFile añade contenido al final de un archivo existente
func (fs *FileSystem) AppendFile(path string, content []byte) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	dir, name := filepath.Split(path)

	parent, err := fs.navigateTo(dir)
	if err != nil {
		return err
	}

	parent.mu.RLock()
	node, exists := parent.children[name]
	parent.mu.RUnlock()

	if !exists {
		// Si no existe, lo creamos
		return fs.CreateFile(path, content, 0644)
	}

	if node.nodeType != FileNode {
		return errors.New("no es un archivo: " + path)
	}

	node.mu.Lock()
	defer node.mu.Unlock()

	node.content = append(node.content, content...)
	node.size = int64(len(node.content))
	node.modTime = time.Now()

	return nil
}

// Rename mueve o renombra un archivo o directorio
func (fs *FileSystem) Rename(oldPath, newPath string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	// Obtener el nodo origen
	oldDir, oldName := filepath.Split(oldPath)
	oldParent, err := fs.navigateTo(oldDir)
	if err != nil {
		return err
	}

	oldParent.mu.Lock()
	node, exists := oldParent.children[oldName]
	if !exists {
		oldParent.mu.Unlock()
		return errors.New("origen no existe: " + oldPath)
	}
	oldParent.mu.Unlock()

	// Obtener el directorio destino
	newDir, newName := filepath.Split(newPath)
	newParent, err := fs.navigateTo(newDir)
	if err != nil {
		return err
	}

	// Si es el mismo directorio y mismo nombre, no hacer nada
	if oldParent == newParent && oldName == newName {
		return nil
	}

	// Verificar que el destino no exista
	newParent.mu.Lock()
	if _, exists := newParent.children[newName]; exists {
		newParent.mu.Unlock()
		return errors.New("destino ya existe: " + newPath)
	}

	// Mover el nodo
	oldParent.mu.Lock()
	delete(oldParent.children, oldName)
	oldParent.modTime = time.Now()
	oldParent.mu.Unlock()

	node.name = newName
	node.parent = newParent
	newParent.children[newName] = node
	newParent.modTime = time.Now()
	newParent.mu.Unlock()

	return nil
}

// Size calcula el tamaño total de un directorio o archivo
func (fs *FileSystem) Size(path string) (int64, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	if path == "/" || path == "" {
		return fs.sizeRecursive(fs.root), nil
	}

	dir, name := filepath.Split(path)
	parent, err := fs.navigateTo(dir)
	if err != nil {
		return 0, err
	}

	parent.mu.RLock()
	node, exists := parent.children[name]
	parent.mu.RUnlock()

	if !exists {
		return 0, errors.New("no existe: " + path)
	}

	return fs.sizeRecursive(node), nil
}

func (fs *FileSystem) sizeRecursive(node *Node) int64 {
	node.mu.RLock()
	defer node.mu.RUnlock()

	if node.nodeType == FileNode {
		return node.size
	}

	var totalSize int64
	for _, child := range node.children {
		totalSize += fs.sizeRecursive(child)
	}

	return totalSize
}