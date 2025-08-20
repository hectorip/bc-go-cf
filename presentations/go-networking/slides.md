---
title: Redes con Go (TCP/UDP)
description: Guía práctica y concisa para crear apps en red con Go
author: Tú + ChatGPT
theme: default
favicon: https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo\_Blue.png

---

# Redes con Go (TCP/UDP)

* Meta: entender **intuitivamente** redes y protocolos con Go
* Combina **conceptos** + **código**
* Enfoque **multiplataforma** (Linux/Mac/Windows)
* De estándar → a librerías externas

---

# Agenda (mini)

* Fundamentos (IP, puertos, sockets)
* TCP vs UDP + handshake
* Stdlib: `net`, `net/http`, `net/url`
* Ejemplos: chat, files, microservicio, IoT, juegos
* Librerías: WebSockets, MQ, gRPC
* Observabilidad y tips

---

# Fundamentos rápidos

* IP = “dirección” del host
* **Puerto** (0–65535) = servicio en el host
* **Socket** = IP\:PUERTO (extremo de comunicación)
* **Cliente/Servidor**: cliente inicia, servidor escucha

---

# TCP vs UDP

* **TCP**: conexión, fiable, ordenado (stream de bytes)
* **UDP**: sin conexión, sin garantías, **rápido** (datagramas)
* Elige TCP para **integridad**; UDP para **latencia**

---

# Handshake TCP (3 pasos)

1. Cliente → **SYN**
2. Servidor → **SYN-ACK**
3. Cliente → **ACK**
   → Canal listo para enviar **stream** de bytes

---

# Paquetes Go clave

* `net`: TCP/UDP, DNS, `Dial`/`Listen`/`Accept`
* `net/http`: servidor/cliente HTTP sencillo
* `net/url`: parseo/armado seguro de URLs

---

# Patrón servidor TCP (flujo)

* `Listen("tcp", ":puerto")`
* `Accept()` en bucle
* 1 goroutine por cliente
* `Read`/`Write` sobre `net.Conn`
* Cerrar al terminar

---

# Código: TCP Echo Server

```go
package main

import (
  "bufio"
  "net"
)

func main() {
  ln, _ := net.Listen("tcp", ":9000")
  defer ln.Close()
  for {
    c, _ := ln.Accept()
    go func(conn net.Conn) {
      defer conn.Close()
      in := bufio.NewScanner(conn)
      for in.Scan() {
        conn.Write(append(in.Bytes(), '\n'))
      }
    }(c)
  }
}
```

---

# Cliente TCP (flujo)

* `Dial("tcp", "host:puerto")`
* Leer stdin / escribir al socket
* Leer respuestas en goroutine
* Cerrar con `Close()`

---

# Código: TCP Cliente mínimo

```go
package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
)

func main() {
  conn, _ := net.Dial("tcp", "127.0.0.1:9000")
  defer conn.Close()

  go func() {
    s := bufio.NewScanner(conn)
    for s.Scan() { fmt.Println("Srv:", s.Text()) }
  }()

  in := bufio.NewScanner(os.Stdin)
  for in.Scan() {
    conn.Write(append([]byte(in.Text()), '\n'))
  }
}
```

---

# UDP (ideas clave)

* **No hay** `Accept()`
* `ListenUDP`/`ReadFromUDP`/`WriteToUDP`
* Tamaño de datagrama limitado
* Usa **timeouts** (`SetReadDeadline`)

---

# Código: UDP Echo Server

```go
package main

import (
  "fmt"
  "net"
)

func main() {
  addr, _ := net.ResolveUDPAddr("udp", ":9999")
  conn, _ := net.ListenUDP("udp", addr)
  defer conn.Close()
  buf := make([]byte, 1024)
  for {
    n, cli, _ := conn.ReadFromUDP(buf)
    fmt.Printf("RX %q de %v\n", string(buf[:n]), cli)
    conn.WriteToUDP(buf[:n], cli)
  }
}
```

---

# Código: UDP Cliente

```go
package main

import (
  "fmt"
  "net"
  "time"
)

func main() {
  srv, _ := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
  c, _ := net.DialUDP("udp", nil, srv)
  defer c.Close()

  c.Write([]byte("ping"))
  c.SetReadDeadline(time.Now().Add(2 * time.Second))
  buf := make([]byte, 1024)
  n, from, err := c.ReadFromUDP(buf)
  if err != nil { fmt.Println("timeout"); return }
  fmt.Println("RX:", string(buf[:n]), "de", from)
}
```

---

# HTTP con `net/http`

* Handlers por **ruta**
* Respuestas **texto/JSON**
* Cliente HTTP simple (`http.Get`)
* Concurrencia automática

---

# Código: HTTP Server mini

```go
package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hola, mundo!")
  })
  http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" { name = "Invitado" }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Hola, " + name})
  })
  http.ListenAndServe(":8080", nil)
}
```

---

# Código: Cliente HTTP

```go
resp, err := http.Get("https://httpbin.org/ip")
if err != nil { log.Fatal(err) }
defer resp.Body.Close()
b, _ := io.ReadAll(resp.Body)
fmt.Println(resp.Status)
fmt.Println(string(b))
```

---

# URLs con `net/url`

* `url.Parse()` → partes seguras
* `Hostname()/Port()`
* `Query()` (map de listas)
* Construye y `String()`

---

# Código: parseo de URL

```go
u, _ := url.Parse("https://user:pass@x.com:8080/p?q=1#sec")
fmt.Println(u.Scheme, u.Hostname(), u.Port(), u.Path)
fmt.Println(u.Query().Get("q"))
u.Path = "/nuevo"
fmt.Println(u.String())
```

---

# Ejemplos prácticos (set)

* Chat TCP multi-cliente
* Transferencia de archivos
* Microservicio REST
* IoT (UDP / MQTT)
* Esqueleto multijugador (UDP)

---

# Chat TCP (idea)

* Mantén `[]Conn` activos
* 1 goroutine por cliente
* **Broadcast** al resto
* Sin bloquear escritura/lectura

---

# Código: broadcast (núcleo)

```go
var (
  clients = make(map[net.Conn]bool)
  mu      sync.Mutex
)

func broadcast(sender net.Conn, msg string) {
  mu.Lock(); defer mu.Unlock()
  for c := range clients {
    if c != sender { fmt.Fprintln(c, msg) }
  }
}
```

---

# Transferencia de archivos (idea)

* Servidor recibe y **guarda**
* Cliente **lee** archivo y envía
* `io.Copy` simplifica mucho
* Un archivo por conexión

---

# Código: servidor de archivos

```go
ln, _ := net.Listen("tcp", ":4000")
for {
  c, _ := ln.Accept()
  go func(conn net.Conn) {
    defer conn.Close()
    f, _ := os.Create("upload.bin")
    defer f.Close()
    io.Copy(f, conn) // hasta EOF
  }(c)
}
```

---

# Código: cliente de archivos

```go
f, _ := os.Open("foto.jpg")
defer f.Close()
c, _ := net.Dial("tcp", "127.0.0.1:4000")
defer c.Close()
io.Copy(c, f)
```

---

# Microservicio (idea)

* Endpoints JSON
* Métodos: GET/POST
* Mutex para estado en memoria
* Validación simple

---

# Código: /tasks (mini)

```go
type Task struct{ ID int; Text string }
var (
  tasks []Task
  id int
  mu sync.Mutex
)

http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost {
    var t Task; json.NewDecoder(r.Body).Decode(&t)
    if t.Text == "" { http.Error(w, "bad", 400); return }
    mu.Lock(); id++; t.ID = id; tasks = append(tasks, t); mu.Unlock()
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(201); json.NewEncoder(w).Encode(t); return
  }
  w.Header().Set("Content-Type","application/json")
  mu.Lock(); json.NewEncoder(w).Encode(tasks); mu.Unlock()
})
```

---

# IoT (rápido)

* UDP para **telemetría rápida**
* MQTT/**NATS** para fiabilidad
* Parseo sencillo de mensajes
* Persistir o alertar

---

# Código: UDP IoT listener

```go
conn, _ := net.ListenPacket("udp", ":2000")
defer conn.Close()
buf := make([]byte, 256)
for {
  n, addr, _ := conn.ReadFrom(buf)
  msg := string(buf[:n])
  fmt.Println("De", addr, "=>", msg)
}
```

---

# Juegos tiempo real (UDP)

* Cliente envía **inputs** frecuentes
* Servidor simula y **broadcast** estado
* Manejar pérdida: repetir/compensar
* 20–60 Hz (sleep entre ticks)

---

# Código: bucle (esqueleto)

```go
for {
  drainInputs(conn, &state)  // lee datagramas
  update(&state)             // física/juego
  pkt := encode(state)       // delta/estado
  for _, p := range players { conn.WriteTo(pkt, p.addr) }
  time.Sleep(50 * time.Millisecond)
}
```

---

# WebSockets (en Go)

* Full-duplex sobre HTTP
* Ideal: chat, dashboards, colab
* Librerías: **nhooyr/websocket**, gobwas/ws
* Alternativa: gRPC streaming

---

# Código: WS eco (nhooyr)

```go
import "nhooyr.io/websocket"

http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
  c, _ := websocket.Accept(w, r, nil)
  defer c.Close(websocket.StatusNormalClosure, "")
  for {
    _, data, err := c.Read(r.Context())
    if err != nil { return }
    c.Write(r.Context(), websocket.MessageText, data)
  }
})
```

---

# Pub/Sub y colas

* **NATS**: simple, veloz (pub/sub)
* **RabbitMQ**: AMQP, enrutamiento, durable
* **MQTT**: IoT ligero, tópicos
* Elige según **fiabilidad** y **volumen**

---

# Código: NATS (mini)

```go
nc, _ := nats.Connect(nats.DefaultURL)
defer nc.Drain()

nc.Subscribe("sensors.temp", func(m *nats.Msg) {
  fmt.Println("temp:", string(m.Data))
})
nc.Publish("sensors.temp", []byte("22.5"))
```

---

# RPC moderno (gRPC)

* Esquemas **protobuf**
* HTTP/2, binario, *streaming*
* Tipado fuerte cliente/servidor
* Ideal microservicios internos

---

# Código: servidor gRPC (esbozo)

```go
type Greeter struct{ pb.UnimplementedGreeterServer }
func (g *Greeter) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
  return &pb.HelloReply{Message: "Hola, " + r.Name}, nil
}

lis, _ := net.Listen("tcp", ":50051")
s := grpc.NewServer()
pb.RegisterGreeterServer(s, &Greeter{})
s.Serve(lis)
```

---

# Observabilidad (imprescindible)

* **pprof**: CPU/mem/goroutines
* **Prometheus**: métricas `/metrics`
* **OpenTelemetry**: trazas/metricas/logs
* Prueba de carga: `hey`, `wrk`

---

# Código: pprof + métricas

```go
import (
  _ "net/http/pprof"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
  http.Handle("/metrics", promhttp.Handler())
  go http.ListenAndServe(":6060", nil) // pprof + metrics
  // ... tu servidor real aquí ...
}
```

---

# Buenas prácticas

* 1 goroutine por conexión (TCP)
* **Timeouts/Deadlines** siempre
* Reintentos (UDP/MQTT)
* Maneja cerrados/errores con cuidado
* Testea bajo carga

---

# Errores comunes

* Bloquear el `Accept()`/lecturas
* No cerrar `Conn()`/`Body`
* Compartir mapas sin mutex
* Leer **líneas** y olvidar `\n`
* No validar entrada del usuario

---

# Siguientes pasos

* Extiende el **chat** (nicks, rooms)
* Añade **TLS** a HTTP/TCP
* Prueba **NATS** en microservicio
* Instrumenta con **Prometheus**
* Explora **gRPC streaming**

---

# Recursos (rápidos)

* `go doc net`, `net/http`, `net/url`
* nhooyr/websocket, nats.go, streadway/amqp
* grpc-go, protobuf
* Prometheus client\_golang, OpenTelemetry

---

# Fin 🧪

* Copia/ejecuta los ejemplos
* Itera y mide
* ¡Construye tu primera app en red con Go!