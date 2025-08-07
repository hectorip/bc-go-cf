Here’s a **Slidev** (Markdown-based) deck you can paste into a `slides.md` and run with Slidev. I kept the syntax simple and included runnable Go snippets.

````md
---
# Working with the Operating System in Go (1.21+)
subtitle: Portability first, escape hatches when needed
author: Your Name
theme: default
class: text-center
highlighter: shiki
lineNumbers: true
---

# Working with the Operating System in Go
Modern Go (1.21+) — macOS · Linux · Windows

<div class="text-sm opacity-70">Portability first, escape hatches when needed</div>

---

## What you'll learn
- Core concepts for OS work in Go
- Stdlib essentials (`os`, `io/fs`, `path/filepath`, `os/exec`, `os/signal`, `log/slog`)
- Useful OSS packages
- Cross-platform differences & patterns
- Notable Go projects that touch the OS deeply
- What’s hard / not feasible
- Exercises & mini-projects (theory → practice)

---

## Mental model
- Prefer **portable stdlib** for 90% of tasks
- Use **`golang.org/x/sys`** for OS-specific needs
- Think in **filesystems as interfaces** (`io/fs`)
- Handle **paths** via `path/filepath`
- Manage **processes** via `os/exec` + `context`
- Graceful **shutdown** with `os/signal` + `signal.NotifyContext`
- Structured **logging** with `log/slog` (Go 1.21)
- **Embed** resources with `embed` (Go 1.16)
- Keep **per-OS code** behind build tags (`//go:build`)

---

## Stdlib map (you’ll use these a lot)
- **Files & dirs**: `os`, `io`, `bufio`, `time`
- **Paths & walking**: `path/filepath`, `io/fs`
- **Processes**: `os/exec` (+ `context`)
- **Signals**: `os/signal` (+ `signal.NotifyContext`)
- **Logging**: `log/slog`
- **User dirs**: `os.UserConfigDir`, `UserCacheDir`, `UserHomeDir`
- **Embedding**: `embed` + `//go:embed`

---

## Paths: never concatenate strings
```go
package main

import (
 "fmt"
 "path/filepath"
 "runtime"
)

func main() {
 base := "/var"
 if runtime.GOOS == "windows" {
  base = `C:\ProgramData`
 }
 p := filepath.Join(base, "acme", "config.yaml")
 fmt.Println(p) // OS-appropriate separators
}
````

---

## Walking files with io/fs (portable)

```go
package main

import (
 "fmt"
 "io/fs"
 "os"
 "path/filepath"
)

func main() {
 root := "."
 err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
  if err != nil { return err }
  if d.IsDir() && (d.Name() == ".git" || d.Name() == "node_modules") {
   return fs.SkipDir
  }
  if !d.IsDir() {
   info, _ := d.Info()
   fmt.Printf("%s (%d bytes)\n", path, info.Size())
  }
  return nil
 })
 if err != nil { panic(err) }
}
```

---

## Embedding assets (Go 1.16+)

```go
package main

import (
 "embed"
 "fmt"
 "io/fs"
)

//go:embed static/*
var assets embed.FS

func main() {
 sub, _ := fs.Sub(assets, "static")
 b, _ := fs.ReadFile(sub, "hello.txt")
 fmt.Println(string(b))
}
```

---

## User config/cache dirs (don’t hard-code)

```go
package main

import (
 "fmt"
 "os"
 "path/filepath"
)

func main() {
 cfg, _ := os.UserConfigDir()
 cache, _ := os.UserCacheDir()
 fmt.Println("Config:", cfg)
 fmt.Println("Cache :", cache)

 confFile := filepath.Join(cfg, "acme", "config.yaml")
 fmt.Println("Conf file:", confFile)
}
```

---

## Running subprocesses with context

```go
package main

import (
 "context"
 "fmt"
 "os/exec"
 "time"
)

func main() {
 ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
 defer cancel()

 cmd := exec.CommandContext(ctx, "sh", "-c", "sleep 2 && echo done")
 out, err := cmd.CombinedOutput()
 fmt.Printf("out=%s err=%v\n", out, err)
}
```

---

## Graceful shutdown via signals

```go
package main

import (
 "context"
 "fmt"
 "net/http"
 "os/signal"
 "time"
)

func main() {
 srv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "hello")
 })}

 ctx, stop := signal.NotifyContext(context.Background(), /* os.Interrupt etc. */)
 defer stop()

 go func() {
  _ = srv.ListenAndServe()
 }()

 <-ctx.Done()
 shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
 defer cancel()
 _ = srv.Shutdown(shutdownCtx)
}
```

---

## Structured logging with slog (Go 1.21)

```go
package main

import (
 "log/slog"
 "os"
)

func main() {
 logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
 logger.Info("starting", "component", "worker", "version", "1.2.3")
 logger.Warn("disk space low", "freeMB", 512)
}
```

---

## Watching files (fsnotify)

```go
package main

import (
 "fmt"
 "log"
 "time"

 "github.com/fsnotify/fsnotify"
)

func main() {
 w, err := fsnotify.NewWatcher()
 if err != nil { log.Fatal(err) }
 defer w.Close()

 if err := w.Add("."); err != nil { log.Fatal(err) }

 debounce := time.NewTimer(0)
 if !debounce.Stop() {}
 for {
  select {
  case ev := <-w.Events:
   if !debounce.Stop() {}
   debounce.Reset(200 * time.Millisecond)
   go func(e fsnotify.Event) {
    <-debounce.C
    fmt.Println("event:", e)
   }(ev)
  case err := <-w.Errors:
   fmt.Println("error:", err)
  }
 }
}
```

---

## Low-level OS access (`golang.org/x/sys`)

* `x/sys/unix`: ioctls, xattrs, capabilities, advanced perms
* `x/sys/windows`: tokens, ACLs, registry, services
* Prefer `x/sys` over deprecated `syscall`
* Use **build tags** to isolate platform code

---

## Build tags & per-OS files

```go
// platform_unix.go
//go:build unix

package platform

func DefaultSocket() string { return "/var/run/acme.sock" }
```

```go
// platform_windows.go
//go:build windows

package platform

func DefaultSocket() string { return `\\.\pipe\acme` }
```

---

## Cross-platform differences (quick view)

| Area     | Linux/macOS            | Windows                         | Go approach               |
| -------- | ---------------------- | ------------------------------- | ------------------------- |
| Paths    | `/` (case-sens varies) | `\` (case-insensitive)          | `path/filepath`           |
| Perms    | POSIX modes            | ACLs, limited modes             | `os.Chmod` + `x/sys`      |
| Signals  | POSIX set              | Limited (Ctrl+C)                | `os/signal`               |
| IPC      | UDS sockets            | Named pipes / AF\_UNIX (limits) | feature-detect & fallback |
| Services | systemd/launchd        | SCM                             | per-OS code or helpers    |

---

## Notable Go OS-heavy projects

* Kubernetes, containerd, runc
* Tailscale (TUN), Syncthing, rclone
* gVisor (user-space kernel boundary)
* FUSE filesystems (`go-fuse`)

---

## Things that are hard / not feasible

* True `fork` semantics on Windows (assume `CreateProcess` model)
* Fine-grained ACLs/attrs: require `x/sys` and privileges
* Full Unix sockets behavior on Windows (fallback to named pipes/TCP)
* Service integration: per-OS code paths (systemd/launchd/SCM)

---

## Common practices checklist

* ✅ Use `filepath.*`, never string concat
* ✅ Use `UserConfigDir`/`UserCacheDir`
* ✅ Wrap blocking ops with `context`
* ✅ Graceful shutdown via `signal.NotifyContext`
* ✅ Structured logs (`slog`)
* ✅ Minimal, clean build-tagged per-OS files
* ✅ Abstract FS via `io/fs`

---

## Use cases

* CLI tools (backup/sync), file watchers, hot-reloaders
* Daemons/agents with clean shutdown and logs
* Dev tooling: task runners, build pipelines
* Local services: choose TCP vs UDS vs named pipes
* Virtual filesystems (FUSE)

---

## Exercises (theory → practice)

### 1) Cross-platform file walk

* List all files with size & mode, skip `.git` and `node_modules`
* Output JSON; log errors via `slog`
* **Stretch**: `--roots` uses `os.PathListSeparator`

### 2) Env & config loader

* Priority: flag > env > `${UserConfigDir}/acme/config.yaml`
* Print where config was read from

### 3) Graceful HTTP shutdown

* Server on `:8080`, stop on Ctrl+C with a timeout
* **Stretch**: second Ctrl+C forces exit

---

### 4) Subprocess wrapper

* Run `tar`/`zip`, stream stdout/stderr, propagate exit codes
* Add context timeout

### 5) File watcher

* Watch a directory; report create/modify/remove
* Debounce bursts (200ms)

### 6) Terminal raw mode

* Small REPL with raw mode (use `x/term`)
* Restore terminal on exit

---

### 7) Embedded assets

* Serve static files from `embed.FS`
* Dev mode switch uses OS FS with `fs.Sub`

### 8) Build tags demo

* `DefaultSocketPath()` with per-OS implementations
* Unit test that asserts the path format

### 9) Registry vs dotfiles

* Windows: store a setting in Registry (`x/sys/windows/registry`)
* Unix: store same setting under `${UserConfigDir()}`

### 10) Local IPC echo

* Unix/macOS: Unix domain socket
* Windows: named pipe (or AF\_UNIX if available)
* Fallback to TCP loopback

---

## Mini-projects (capstones)

### A) Cross-platform log tailer

* Tail files with rotation support (`fsnotify`)
* Normalize CRLF/LF, backpressure handling
* Structured JSON logs with `slog`

### B) Service runner

* Start child process, restart with backoff
* Install/run as systemd unit (Linux) or Windows service
* Graceful shutdown on signals

---

### C) Pluggable file sync

* Sync `io/fs` sources (DirFS / embed.FS) to disk
* Filters: glob patterns, `.gitignore`-style
* Dry-run mode with diff summary

### D) Pipeline runner

* YAML steps → external commands with env/cwd/timeout
* Per-OS default shell selection
* Logs and exit code aggregation

### E) TTY file manager

* Minimal ncurses-like browser using raw mode
* Resize handling, keybindings, status bar

---

## Closing

* Start with stdlib abstractions
* Reach for `x/sys` only when needed
* Design for portability; isolate OS-specific code
* Log richly, shut down gracefully, test with `io/fs`

---

## Appendix: tiny code snippets

### Windows Registry read/write

```go
//go:build windows
package main

import (
 "fmt"
 "golang.org/x/sys/windows/registry"
)

func main() {
 k, _, _ := registry.CreateKey(registry.CURRENT_USER, `Software\Acme`, registry.ALL_ACCESS)
 defer k.Close()
 _ = k.SetStringValue("Theme", "dark")
 s, _, _ := k.GetStringValue("Theme")
 fmt.Println("Theme:", s)
}
```

### Unix domain socket vs fallback

```go
//go:build unix
package main

import (
 "fmt"
 "net"
)

func main() {
 l, err := net.Listen("unix", "/tmp/acme.sock")
 if err != nil { panic(err) }
 fmt.Println("listening on UDS:", l.Addr())
}
```

```go
//go:build windows
package main

import (
 "fmt"
 "net"
)

func main() {
 // If AF_UNIX is not suitable, fallback to TCP loopback
 l, err := net.Listen("tcp", "127.0.0.1:0")
 if err != nil { panic(err) }
 fmt.Println("listening on TCP:", l.Addr())
}
```

---

## Q\&A

* What needs to be OS-specific in your current projects?
* Where can you replace shell scripts with Go for portability?

```

If you want, I can split this into separate files (slides + code starters) and add a basic `package.json`/`slides.md` scaffold for quick `npx slidev` runs.
::contentReference[oaicite:0]{index=0}
```
