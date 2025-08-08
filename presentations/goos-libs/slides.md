---
title: Go + SO — Paquetes estándar para interactuar con el sistema
info: "Cobertura: Go 1.22/1.23 — slides con poco texto"
---

# Go + Sistema Operativo
Mapa de paquetes estándar útiles para el SO

---
layout: center
---

## Categorías
- Archivos y rutas
- Procesos y señales
- Usuario/entorno/tiempo
- Red
- Logs
- Runtime
- Binarios y formatos
- Empaquetado/compresión
- Extras (no estándar)

---

## Archivos y rutas — `os`
**Qué:** FS real + procesos/env  
**Para:** abrir/escribir, permisos, env

```go
f,_:=os.Create("out.txt"); defer f.Close()
os.Setenv("MODE","prod")
```

---

## Archivos y rutas — `path/filepath`
**Qué:** rutas según SO  
**Para:** `Join`, `WalkDir`, `Glob`

```go
p:=filepath.Join("/var","log","app.log")
```

---

## Archivos y rutas — `io/fs`
**Qué:** interfaz FS (solo lectura)  
**Para:** código portable (OS, embed, test)

```go
b,_:=fs.ReadFile(fsys,"cfg.json")
```

---

## Archivos y rutas — `embed`
**Qué:** archivos dentro del binario  
**Para:** assets/templates sin disco

```go
//go:embed static/*
var staticFS embed.FS
```

---

## Archivos y rutas — `path`
**Qué:** rutas con `/` (no OS)  
**Para:** imports/URL-like

```go
path.Clean("a/b/../c")
```

---

## Archivos y rutas — `testing/fstest`
**Qué:** FS de prueba en memoria  
**Para:** tests de código `fs.FS`

```go
fstest.MapFS{"f.txt":{Data:[]byte("ok")}}
```

---

## Procesos — `os/exec`
**Qué:** ejecutar comandos  
**Para:** pipelines, capturar salida

```go
out,_:=exec.Command("sh","-c","echo hi").Output()
```

---

## Señales — `os/signal`
**Qué:** señales del SO  
**Para:** cierre limpio (Ctrl-C)

```go
ctx,_:=signal.NotifyContext(context.Background(),os.Interrupt)
<-ctx.Done()
```

---

## Bajo nivel — `syscall` (congelado)
**Qué:** syscalls crudos  
**Para:** casos nicho (usar `x/sys`)

```go
cmd.SysProcAttr=&syscall.SysProcAttr{Setpgid:true}
```

---

## Usuario — `os/user`
**Qué:** info de usuario/grupos  
**Para:** home, UID/GID

```go
u,_:=user.Current(); _=u.HomeDir
```

---

## Tiempo — `time`
**Qué:** tiempos/timeout  
**Para:** `After`,`Ticker`, deadlines

```go
<-time.After(500*time.Millisecond)
```

---

## Red — `net`
**Qué:** TCP/UDP/Unix sockets  
**Para:** servidores/IFs/DNS

```go
ln,_:=net.Listen("tcp",":8080"); _=ln
```

---

## Red — `net/netip`
**Qué:** IPs inmutables/comparables  
**Para:** `Addr`,`Prefix` claros

```go
ip, _ := netip.ParseAddr("192.0.2.1")
```

---

## Red — `net/http` (capa alta)
**Qué:** HTTP(S) cliente/servidor  
**Para:** sobre `net`/TLS del SO

```go
http.Get("https://example.com")
```

---

## Logs — `log`
**Qué:** logger básico  
**Para:** stdout/archivo

```go
log.Println("arrancó")
```

---

## Logs — `log/slog`
**Qué:** logging estructurado  
**Para:** niveles + attrs

```go
slog.New(slog.NewJSONHandler(os.Stdout,nil)).Info("ok")
```

---

## Logs — `log/syslog` (Unix)
**Qué:** syslog del sistema  
**Para:** integrar con daemon

```go
w,_:=syslog.Dial("", "", syslog.LOG_INFO, "app")
```

---

## Runtime — `runtime`
**Qué:** GOOS/CPU/hilos  
**Para:** introspección/ajustes

```go
runtime.GOOS; runtime.NumCPU()
```

---

## Runtime — `runtime/debug`
**Qué:** build info + GC  
**Para:** diagnósticos puntuales

```go
info,_:=debug.ReadBuildInfo(); _=info.Main.Path
```

---

## Runtime — `runtime/metrics`
**Qué:** métricas estables del runtime  
**Para:** observabilidad

```go
// consultar keys y samples
```

---

## Binarios — `plugin`
**Qué:** cargar `.so` (Linux/macOS)  
**Para:** extensiones en runtime

```go
p,_:=plugin.Open("mod.so"); _,_=p.Lookup("Sym")
```

---

## Binarios — `debug/elf`
**Qué:** formato ELF (Linux)  
**Para:** inspección binarios

```go
f,_:=elf.Open("bin"); _=f.Sections
```

---

## Binarios — `debug/macho`
**Qué:** formato Mach-O (macOS)  
**Para:** tooling

```go
m,_:=macho.Open("bin"); _=m.FileHeader
```

---

## Binarios — `debug/pe`
**Qué:** formato PE (Windows)  
**Para:** tooling

```go
p,_:=pe.Open("app.exe"); _=p.Sections
```

---

## Binarios — `debug/dwarf`
**Qué:** datos DWARF  
**Para:** debuggers/analizadores

```go
// leer secciones DWARF
```

---

## Empaquetado — `archive/tar`
**Qué:** tar stream  
**Para:** backups/artefactos

```go
tw:=tar.NewWriter(w); _=tw.Close()
```

---

## Empaquetado — `archive/zip`
**Qué:** ZIP archivos  
**Para:** empaquetar/procesar

```go
zw:=zip.NewWriter(w); _=zw.Create("f.txt")
```

---

## Compresión — `compress/gzip`
**Qué:** gzip stream  
**Para:** comprimir logs/archivos

```go
gz:=gzip.NewWriter(w); _=gz.Close()
```

---

## Certs — `crypto/x509`
**Qué:** X.509 + trust del SO  
**Para:** `SystemCertPool`

```go
pool,_:=x509.SystemCertPool()
```

---

## Apoyo — `context` y `io/bufio`
**Qué:** cancelación + I/O eficiente  
**Para:** tiempo de OS y buffers

```go
ctx,cancel:=context.WithTimeout(context.Background(),1*time.Second); defer cancel()
```

---

## ¿Qué NO está en stdlib?
- Terminal TTY: `golang.org/x/term`
- Syscalls modernos: `golang.org/x/sys`
- File watching: `github.com/fsnotify/fsnotify`

---

## Cheatsheet final
- **FS real:** `os` + `filepath`
- **FS portable:** `io/fs` + `embed`
- **Procesos:** `os/exec`, **Señales:** `os/signal`
- **Red:** `net` (+ `netip`, `http`)
- **Logs:** `log` / `log/slog` / `log/syslog*`
- **Runtime:** `runtime`, `debug`, `metrics`
- **Binarios:** `plugin`, `debug/*`
- **Empaquetado:** `archive/*`, `compress/*`
