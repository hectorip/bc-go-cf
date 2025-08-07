package minifs

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
	"time"
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

	t.Run("CreateNestedDirs", func(t *testing.T) {
		err := fs.MkdirAll("/test/nested/deep/dir", 0755)
		if err != nil {
			t.Fatalf("Error creando directorios anidados: %v", err)
		}

		if !fs.Exists("/test/nested/deep/dir") {
			t.Error("Los directorios anidados no fueron creados")
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

	t.Run("OverwriteFile", func(t *testing.T) {
		newContent := []byte("nuevo contenido")
		err := fs.WriteFile("/test/archivo.txt", newContent)
		if err != nil {
			t.Fatalf("Error sobrescribiendo archivo: %v", err)
		}

		data, err := fs.ReadFile("/test/archivo.txt")
		if err != nil {
			t.Fatalf("Error leyendo archivo: %v", err)
		}

		if !bytes.Equal(data, newContent) {
			t.Error("El contenido no fue actualizado")
		}
	})

	t.Run("ListDir", func(t *testing.T) {
		// Crear algunos archivos más
		fs.WriteFile("/test/archivo2.txt", []byte("contenido 2"))
		fs.CreateDir("/test/subdir", 0755)

		files, err := fs.ListDir("/test")
		if err != nil {
			t.Fatalf("Error listando directorio: %v", err)
		}

		if len(files) < 3 {
			t.Errorf("Se esperaban al menos 3 elementos, se encontraron %d", len(files))
		}

		// Verificar que los nombres estén presentes
		names := make(map[string]bool)
		for _, f := range files {
			names[f.Name] = true
		}

		if !names["archivo.txt"] || !names["archivo2.txt"] || !names["subdir"] {
			t.Error("No se encontraron todos los archivos esperados")
		}
	})

	t.Run("Stat", func(t *testing.T) {
		info, err := fs.Stat("/test/archivo.txt")
		if err != nil {
			t.Fatalf("Error obteniendo información del archivo: %v", err)
		}

		if info.IsDir {
			t.Error("Se esperaba un archivo, no un directorio")
		}

		if info.Name != "archivo.txt" {
			t.Errorf("Nombre incorrecto: %s", info.Name)
		}

		if info.Size == 0 {
			t.Error("El tamaño del archivo no debería ser 0")
		}
	})

	t.Run("Remove", func(t *testing.T) {
		err := fs.Remove("/test/archivo2.txt")
		if err != nil {
			t.Fatalf("Error eliminando archivo: %v", err)
		}

		if fs.Exists("/test/archivo2.txt") {
			t.Error("El archivo no fue eliminado")
		}
	})

	t.Run("RemoveNonEmptyDir", func(t *testing.T) {
		err := fs.Remove("/test")
		if err == nil {
			t.Error("Se esperaba un error al eliminar directorio no vacío")
		}
	})

	t.Run("RemoveAll", func(t *testing.T) {
		err := fs.RemoveAll("/test/nested")
		if err != nil {
			t.Fatalf("Error eliminando directorio recursivamente: %v", err)
		}

		if fs.Exists("/test/nested") {
			t.Error("El directorio no fue eliminado")
		}
	})

	t.Run("AppendFile", func(t *testing.T) {
		original := []byte("inicio")
		fs.WriteFile("/test/append.txt", original)

		addition := []byte(" fin")
		err := fs.AppendFile("/test/append.txt", addition)
		if err != nil {
			t.Fatalf("Error añadiendo al archivo: %v", err)
		}

		data, _ := fs.ReadFile("/test/append.txt")
		expected := append(original, addition...)
		if !bytes.Equal(data, expected) {
			t.Errorf("Contenido incorrecto: got %s, want %s", data, expected)
		}
	})

	t.Run("Rename", func(t *testing.T) {
		fs.WriteFile("/test/old.txt", []byte("contenido"))

		err := fs.Rename("/test/old.txt", "/test/new.txt")
		if err != nil {
			t.Fatalf("Error renombrando archivo: %v", err)
		}

		if fs.Exists("/test/old.txt") {
			t.Error("El archivo viejo todavía existe")
		}

		if !fs.Exists("/test/new.txt") {
			t.Error("El archivo nuevo no existe")
		}

		data, _ := fs.ReadFile("/test/new.txt")
		if string(data) != "contenido" {
			t.Error("El contenido no se preservó")
		}
	})

	t.Run("Size", func(t *testing.T) {
		// Crear estructura de prueba
		fs.MkdirAll("/size/test", 0755)
		fs.WriteFile("/size/test/file1.txt", bytes.Repeat([]byte("a"), 100))
		fs.WriteFile("/size/test/file2.txt", bytes.Repeat([]byte("b"), 200))
		fs.WriteFile("/size/file3.txt", bytes.Repeat([]byte("c"), 50))

		size, err := fs.Size("/size")
		if err != nil {
			t.Fatalf("Error calculando tamaño: %v", err)
		}

		if size != 350 {
			t.Errorf("Tamaño incorrecto: got %d, want 350", size)
		}
	})
}

func TestWalk(t *testing.T) {
	fs := NewFileSystem()

	// Crear estructura de prueba
	fs.MkdirAll("/walk/dir1/subdir", 0755)
	fs.MkdirAll("/walk/dir2", 0755)
	fs.WriteFile("/walk/file1.txt", []byte("contenido1"))
	fs.WriteFile("/walk/dir1/file2.txt", []byte("contenido2"))
	fs.WriteFile("/walk/dir1/subdir/file3.txt", []byte("contenido3"))

	var paths []string
	err := fs.Walk("/walk", func(path string, info FileInfo) error {
		paths = append(paths, path)
		return nil
	})

	if err != nil {
		t.Fatalf("Error recorriendo árbol: %v", err)
	}

	// Ordenar para comparación consistente
	sort.Strings(paths)

	expected := []string{
		"/walk",
		"/walk/dir1",
		"/walk/dir1/file2.txt",
		"/walk/dir1/subdir",
		"/walk/dir1/subdir/file3.txt",
		"/walk/dir2",
		"/walk/file1.txt",
	}
	sort.Strings(expected)

	if len(paths) != len(expected) {
		t.Errorf("Número de rutas incorrecto: got %d, want %d", len(paths), len(expected))
	}
}

func TestConcurrency(t *testing.T) {
	fs := NewFileSystem()
	fs.MkdirAll("/concurrent/test", 0755)

	// Crear múltiples archivos concurrentemente
	done := make(chan bool, 100)

	// Escritores
	for i := 0; i < 50; i++ {
		go func(n int) {
			path := fmt.Sprintf("/concurrent/test/file%d.txt", n)
			content := []byte(fmt.Sprintf("Archivo %d", n))
			fs.WriteFile(path, content)
			done <- true
		}(i)
	}

	// Lectores
	for i := 0; i < 50; i++ {
		go func(n int) {
			// Leer archivos que pueden o no existir aún
			path := fmt.Sprintf("/concurrent/test/file%d.txt", n%10)
			fs.ReadFile(path)
			done <- true
		}(i)
	}

	// Esperar a que terminen
	for i := 0; i < 100; i++ {
		<-done
	}

	// Verificar que los archivos fueron creados
	files, err := fs.ListDir("/concurrent/test")
	if err != nil {
		t.Fatalf("Error listando directorio: %v", err)
	}

	if len(files) != 50 {
		t.Errorf("Se esperaban 50 archivos, se encontraron %d", len(files))
	}
}

func TestEdgeCases(t *testing.T) {
	fs := NewFileSystem()

	t.Run("RootOperations", func(t *testing.T) {
		// No se debe poder eliminar la raíz
		err := fs.Remove("/")
		if err == nil {
			t.Error("Se permitió eliminar la raíz")
		}

		err = fs.RemoveAll("/")
		if err == nil {
			t.Error("Se permitió eliminar la raíz con RemoveAll")
		}

		// Se debe poder listar la raíz
		files, err := fs.ListDir("/")
		if err != nil {
			t.Error("No se pudo listar la raíz")
		}

		if files == nil {
			t.Error("La lista de archivos no debería ser nil")
		}
	})

	t.Run("InvalidPaths", func(t *testing.T) {
		// Intentar crear archivo en directorio inexistente
		err := fs.WriteFile("/nonexistent/file.txt", []byte("test"))
		if err == nil {
			t.Error("Se permitió crear archivo en directorio inexistente")
		}

		// Intentar leer archivo inexistente
		_, err = fs.ReadFile("/nonexistent.txt")
		if err == nil {
			t.Error("Se permitió leer archivo inexistente")
		}
	})

	t.Run("FileAsDirectory", func(t *testing.T) {
		// Crear un archivo
		fs.WriteFile("/file.txt", []byte("contenido"))

		// Intentar usarlo como directorio
		err := fs.CreateDir("/file.txt/subdir", 0755)
		if err == nil {
			t.Error("Se permitió crear directorio dentro de un archivo")
		}

		err = fs.WriteFile("/file.txt/another.txt", []byte("test"))
		if err == nil {
			t.Error("Se permitió crear archivo dentro de un archivo")
		}
	})

	t.Run("EmptyNames", func(t *testing.T) {
		err := fs.CreateDir("/", 0755)
		if err == nil {
			t.Error("Se permitió crear directorio con nombre vacío")
		}

		err = fs.WriteFile("/", []byte("test"))
		if err == nil {
			t.Error("Se permitió crear archivo con nombre vacío")
		}
	})
}

func TestMetadata(t *testing.T) {
	fs := NewFileSystem()

	t.Run("ModificationTime", func(t *testing.T) {
		// Crear archivo y verificar tiempo de modificación
		beforeCreate := time.Now()
		fs.WriteFile("/test.txt", []byte("inicial"))
		afterCreate := time.Now()

		info, _ := fs.Stat("/test.txt")
		if info.ModTime.Before(beforeCreate) || info.ModTime.After(afterCreate) {
			t.Error("Tiempo de modificación fuera del rango esperado")
		}

		// Esperar un momento y modificar
		time.Sleep(10 * time.Millisecond)
		beforeModify := time.Now()
		fs.WriteFile("/test.txt", []byte("modificado"))
		afterModify := time.Now()

		info, _ = fs.Stat("/test.txt")
		if info.ModTime.Before(beforeModify) || info.ModTime.After(afterModify) {
			t.Error("Tiempo de modificación no se actualizó correctamente")
		}
	})

	t.Run("FileMode", func(t *testing.T) {
		// Crear archivo con permisos específicos
		fs.CreateFile("/executable.sh", []byte("#!/bin/bash"), 0755)
		info, _ := fs.Stat("/executable.sh")
		if info.Mode != 0755 {
			t.Errorf("Permisos incorrectos: got %v, want %v", info.Mode, 0755)
		}

		// Crear directorio con permisos
		fs.CreateDir("/private", 0700)
		info, _ = fs.Stat("/private")
		if info.Mode != 0700 {
			t.Errorf("Permisos de directorio incorrectos: got %v, want %v", info.Mode, 0700)
		}
	})

	t.Run("FileSize", func(t *testing.T) {
		content := []byte("Este es el contenido del archivo de prueba")
		fs.WriteFile("/sized.txt", content)

		info, _ := fs.Stat("/sized.txt")
		if info.Size != int64(len(content)) {
			t.Errorf("Tamaño incorrecto: got %d, want %d", info.Size, len(content))
		}

		// Verificar que directorios reportan tamaño 0 en Stat
		fs.CreateDir("/emptydir", 0755)
		info, _ = fs.Stat("/emptydir")
		if info.Size != 0 {
			t.Errorf("Tamaño de directorio debería ser 0, got %d", info.Size)
		}
	})
}

func BenchmarkWriteFile(b *testing.B) {
	fs := NewFileSystem()
	content := []byte("benchmark content")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		path := fmt.Sprintf("/bench/file%d.txt", i)
		fs.WriteFile(path, content)
	}
}

func BenchmarkReadFile(b *testing.B) {
	fs := NewFileSystem()
	content := []byte("benchmark content")
	fs.WriteFile("/bench.txt", content)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.ReadFile("/bench.txt")
	}
}

func BenchmarkListDir(b *testing.B) {
	fs := NewFileSystem()
	
	// Crear muchos archivos
	for i := 0; i < 100; i++ {
		path := fmt.Sprintf("/file%d.txt", i)
		fs.WriteFile(path, []byte("content"))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.ListDir("/")
	}
}

func BenchmarkConcurrentWrites(b *testing.B) {
	fs := NewFileSystem()
	fs.MkdirAll("/concurrent", 0755)

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			path := fmt.Sprintf("/concurrent/file%d.txt", i)
			fs.WriteFile(path, []byte("concurrent content"))
			i++
		}
	})
}

func BenchmarkWalk(b *testing.B) {
	fs := NewFileSystem()
	
	// Crear estructura de árbol
	for i := 0; i < 10; i++ {
		dir := fmt.Sprintf("/dir%d", i)
		fs.CreateDir(dir, 0755)
		for j := 0; j < 10; j++ {
			file := fmt.Sprintf("%s/file%d.txt", dir, j)
			fs.WriteFile(file, []byte("content"))
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.Walk("/", func(path string, info FileInfo) error {
			return nil
		})
	}
}