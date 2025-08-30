package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hectorip/minifs"
)

func main() {
	fmt.Println("=== Mini Sistema de Archivos - Ejemplo de Uso ===\n")

	// Crear sistema de archivos
	fs := minifs.NewFileSystem()

	// 1. Crear estructura de directorios
	fmt.Println("1. Creando estructura de directorios...")
	dirs := []string{
		"/home/user/documents",
		"/home/user/downloads",
		"/home/user/projects/golang",
		"/home/user/projects/python",
		"/var/log",
		"/var/cache",
		"/etc/config",
		"/tmp",
	}

	for _, dir := range dirs {
		if err := fs.MkdirAll(dir, 0755); err != nil {
			log.Printf("Error creando %s: %v", dir, err)
		} else {
			fmt.Printf("   ✓ Creado: %s\n", dir)
		}
	}

	// 2. Crear archivos con contenido
	fmt.Println("\n2. Creando archivos...")
	files := map[string]string{
		"/home/user/documents/readme.txt":      "Este es un documento de ejemplo",
		"/home/user/documents/notas.md":        "# Notas\n\n- Item 1\n- Item 2\n- Item 3",
		"/home/user/projects/golang/main.go":   "package main\n\nfunc main() {\n    println(\"Hello, World!\")\n}",
		"/home/user/projects/python/app.py":    "#!/usr/bin/env python3\n\nprint('Hello, World!')",
		"/etc/config/app.conf":                 "server.port=8080\nserver.host=localhost",
		"/var/log/system.log":                  "2024-01-01 10:00:00 Sistema iniciado\n",
		"/home/user/.bashrc":                   "export PATH=$PATH:/usr/local/bin",
		"/home/user/.gitconfig":                "[user]\n    name = Usuario\n    email = usuario@example.com",
	}

	for path, content := range files {
		if err := fs.WriteFile(path, []byte(content)); err != nil {
			log.Printf("Error creando %s: %v", path, err)
		} else {
			fmt.Printf("   ✓ Archivo creado: %s (%d bytes)\n", path, len(content))
		}
	}

	// 3. Leer y mostrar contenido de un archivo
	fmt.Println("\n3. Leyendo archivo /home/user/projects/golang/main.go:")
	if content, err := fs.ReadFile("/home/user/projects/golang/main.go"); err != nil {
		log.Printf("Error leyendo archivo: %v", err)
	} else {
		fmt.Println("   Contenido:")
		fmt.Println("   " + strings.ReplaceAll(string(content), "\n", "\n   "))
	}

	// 4. Añadir contenido a un archivo (log)
	fmt.Println("\n4. Añadiendo entradas al log...")
	logEntries := []string{
		"2024-01-01 10:05:00 Usuario conectado",
		"2024-01-01 10:10:00 Proceso completado",
		"2024-01-01 10:15:00 Sistema actualizado",
	}

	for _, entry := range logEntries {
		if err := fs.AppendFile("/var/log/system.log", []byte(entry+"\n")); err != nil {
			log.Printf("Error añadiendo al log: %v", err)
		} else {
			fmt.Printf("   ✓ Entrada añadida: %s\n", entry)
		}
	}

	// 5. Listar contenido de directorios
	fmt.Println("\n5. Listando contenido de directorios:")
	dirsToList := []string{"/home/user", "/home/user/projects", "/var"}

	for _, dir := range dirsToList {
		fmt.Printf("\n   Contenido de %s:\n", dir)
		if files, err := fs.ListDir(dir); err != nil {
			log.Printf("   Error listando %s: %v", dir, err)
		} else {
			for _, file := range files {
				if file.IsDir {
					fmt.Printf("   📁 %-20s (dir)\n", file.Name+"/")
				} else {
					fmt.Printf("   📄 %-20s (%d bytes)\n", file.Name, file.Size)
				}
			}
		}
	}

	// 6. Obtener información de archivos
	fmt.Println("\n6. Información de archivos:")
	pathsToStat := []string{
		"/home/user/documents/readme.txt",
		"/home/user/projects",
		"/var/log/system.log",
	}

	for _, path := range pathsToStat {
		if info, err := fs.Stat(path); err != nil {
			log.Printf("Error obteniendo info de %s: %v", path, err)
		} else {
			fmt.Printf("\n   %s:\n", path)
			fmt.Printf("   - Tipo: %s\n", map[bool]string{true: "Directorio", false: "Archivo"}[info.IsDir])
			fmt.Printf("   - Tamaño: %d bytes\n", info.Size)
			fmt.Printf("   - Permisos: %o\n", info.Mode)
			fmt.Printf("   - Modificado: %s\n", info.ModTime.Format(time.RFC3339))
		}
	}

	// 7. Calcular tamaño de directorios
	fmt.Println("\n7. Tamaño de directorios:")
	dirsToSize := []string{"/home/user/documents", "/home/user/projects", "/var", "/"}

	for _, dir := range dirsToSize {
		if size, err := fs.Size(dir); err != nil {
			log.Printf("Error calculando tamaño de %s: %v", dir, err)
		} else {
			fmt.Printf("   %s: %d bytes\n", dir, size)
		}
	}

	// 8. Renombrar archivos
	fmt.Println("\n8. Renombrando archivos...")
	renameOps := []struct {
		old, new string
	}{
		{"/home/user/documents/readme.txt", "/home/user/documents/README.md"},
		{"/home/user/projects/python", "/home/user/projects/python3"},
	}

	for _, op := range renameOps {
		if err := fs.Rename(op.old, op.new); err != nil {
			log.Printf("Error renombrando %s a %s: %v", op.old, op.new, err)
		} else {
			fmt.Printf("   ✓ Renombrado: %s → %s\n", op.old, op.new)
		}
	}

	// 9. Eliminar archivos y directorios
	fmt.Println("\n9. Eliminando archivos...")
	toRemove := []string{
		"/tmp",
		"/home/user/.bashrc",
		"/var/cache",
	}

	for _, path := range toRemove {
		if err := fs.Remove(path); err != nil {
			// Intentar con RemoveAll si es un directorio no vacío
			if err := fs.RemoveAll(path); err != nil {
				log.Printf("Error eliminando %s: %v", path, err)
			} else {
				fmt.Printf("   ✓ Eliminado (recursivo): %s\n", path)
			}
		} else {
			fmt.Printf("   ✓ Eliminado: %s\n", path)
		}
	}

	// 10. Recorrer todo el árbol de archivos
	fmt.Println("\n10. Árbol completo del sistema de archivos:")
	fmt.Println()
	
	var printTree func(string, int)
	printTree = func(path string, level int) {
		fs.Walk(path, func(p string, info minifs.FileInfo) error {
			// Calcular nivel de indentación
			depth := strings.Count(p, "/") - strings.Count(path, "/")
			if depth != level {
				return nil
			}
			
			indent := strings.Repeat("  ", depth)
			
			// Obtener solo el nombre del archivo/directorio
			name := info.Name
			if name == "/" {
				name = "/"
			}
			
			// Imprimir con formato de árbol
			if info.IsDir {
				fmt.Printf("%s📁 %s/\n", indent, name)
				// Recursivamente imprimir subdirectorios
				if p != "/" {
					printTree(p, level+1)
				} else {
					printTree(p, 1)
				}
			} else {
				fmt.Printf("%s📄 %s (%d bytes)\n", indent, name, info.Size)
			}
			
			return nil
		})
	}
	
	printTree("/", 0)

	// 11. Estadísticas finales
	fmt.Println("\n11. Estadísticas finales:")
	
	var totalFiles, totalDirs int
	var totalSize int64
	
	fs.Walk("/", func(path string, info minifs.FileInfo) error {
		if info.IsDir {
			totalDirs++
		} else {
			totalFiles++
			totalSize += info.Size
		}
		return nil
	})
	
	fmt.Printf("   - Total de directorios: %d\n", totalDirs)
	fmt.Printf("   - Total de archivos: %d\n", totalFiles)
	fmt.Printf("   - Tamaño total: %d bytes\n", totalSize)
	fmt.Printf("   - Tamaño promedio por archivo: %.2f bytes\n", float64(totalSize)/float64(totalFiles))

	// 12. Prueba de concurrencia
	fmt.Println("\n12. Prueba de operaciones concurrentes...")
	done := make(chan bool, 20)
	
	// Crear archivos concurrentemente
	for i := 0; i < 10; i++ {
		go func(n int) {
			path := fmt.Sprintf("/home/user/downloads/file%d.txt", n)
			content := fmt.Sprintf("Archivo concurrente #%d", n)
			fs.WriteFile(path, []byte(content))
			done <- true
		}(i)
	}
	
	// Leer archivos concurrentemente
	for i := 0; i < 10; i++ {
		go func(n int) {
			path := fmt.Sprintf("/home/user/downloads/file%d.txt", n)
			fs.ReadFile(path)
			done <- true
		}(i)
	}
	
	// Esperar a que terminen todas las operaciones
	for i := 0; i < 20; i++ {
		<-done
	}
	
	// Verificar archivos creados
	if files, err := fs.ListDir("/home/user/downloads"); err == nil {
		fmt.Printf("   ✓ Archivos creados concurrentemente: %d\n", len(files))
	}

	fmt.Println("\n=== Ejemplo completado exitosamente ===")
}