# Ejemplo de Comunicación UDP con Go

Aplicación de ejemplo que demuestra la comunicación entre dos servidores usando UDP con la biblioteca estándar de Go.

## Arquitectura

El sistema consta de dos servidores que se comunican mediante UDP:

- **Server A** (Puerto 5001): Envía heartbeats periódicos y puede enviar mensajes personalizados
- **Server B** (Puerto 5002): Recibe heartbeats, responde con ACKs, y envía PINGs periódicos

## Características

- ✅ Comunicación bidireccional UDP
- ✅ Heartbeats automáticos cada 3 segundos
- ✅ Sistema de ACK/respuesta para confirmar recepción
- ✅ Timeouts configurados para evitar bloqueos
- ✅ Estadísticas en tiempo real
- ✅ Modo interactivo para enviar mensajes manualmente
- ✅ Manejo graceful de señales (Ctrl+C)
- ✅ Sin dependencias externas (solo stdlib)

## Instalación

```bash
# Clonar o crear el directorio
cd udp-servers-example

# Compilar
go build -o udp-servers

# O ejecutar directamente
go run .
```

## Uso

### Modo básico (ambos servidores)

```bash
# Ejecutar ambos servidores con configuración por defecto
go run .

# O con el binario compilado
./udp-servers
```

### Modo interactivo

```bash
# Habilitar modo interactivo para enviar mensajes manualmente
go run . -interactive

# Comandos disponibles en modo interactivo:
# > send Hello World   # Envía mensaje a Server B
# > help               # Muestra ayuda
# > exit               # Sale del programa
```

### Ejecutar servidores por separado

```bash
# Terminal 1: Solo Server A
go run . -mode a -port-a 5001 -port-b 5002

# Terminal 2: Solo Server B
go run . -mode b -port-a 5001 -port-b 5002
```

### Configuración personalizada

```bash
# Cambiar puertos
go run . -port-a 6001 -port-b 6002

# Ver todas las opciones
go run . -help
```

## Flujo de comunicación

```
Server A (5001)                    Server B (5002)
     |                                   |
     |-------- HEARTBEAT:1 ------------->|
     |<-------- ACK:HEARTBEAT:1 ---------|
     |                                   |
     |-------- HEARTBEAT:2 ------------->|
     |<-------- ACK:HEARTBEAT:2 ---------|
     |                                   |
     |<------------ PING ----------------|
     |------------ PONG ---------------->|
     |<-------- RECEIVED:PONG:OK --------|
     |                                   |
```

## Formato de mensajes

- **Heartbeat**: `HEARTBEAT:<counter>:ServerA:Status=OK`
- **ACK**: `ACK:HEARTBEAT:<counter>`
- **Mensaje**: `MSG:<content>:ServerA:<timestamp>`
- **Respuesta**: `RECEIVED:<content>:OK`
- **Ping/Pong**: Mensajes simples para verificar conectividad

## Características técnicas

- **Goroutines**: Cada servidor usa 3-4 goroutines para manejar concurrentemente:
  - Recepción de mensajes
  - Procesamiento de mensajes
  - Tareas periódicas (heartbeats, estadísticas)
  - Entrada interactiva (opcional)

- **Channels**: Comunicación segura entre goroutines usando channels tipados

- **Timeouts**: Configurados en lecturas UDP (1 segundo) para evitar bloqueos

- **Buffers**: 1024 bytes para mensajes UDP (suficiente para este ejemplo)

## Salida esperada

```
15:30:45.123456 === Sistema de Comunicación UDP ===
15:30:45.123567 Configuración: Server A puerto 5001, Server B puerto 5002
15:30:45.123678 Iniciando ambos servidores...
15:30:45.123789 [Server A] Iniciado en [::]:5001
15:30:45.123890 [Server B] Iniciado en [::]:5002
15:30:46.124001 === Sistema iniciado correctamente ===

15:30:48.124567 [Server A] Heartbeat #1 enviado a Server B
15:30:48.124678 [Server B] Heartbeat #1 de ServerA - Status: Status=OK
15:30:48.124789 [Server A] Respuesta de 127.0.0.1:5002: ACK:HEARTBEAT:1

15:30:51.124567 [Server A] Heartbeat #2 enviado a Server B
15:30:51.124678 [Server B] Heartbeat #2 de ServerA - Status: Status=OK

15:30:55.124567 [Server B] === ESTADÍSTICAS ===
15:30:55.124678 [Server B] Heartbeats recibidos: 4
15:30:55.124789 [Server B] Mensajes recibidos: 4
15:30:55.124890 [Server B] Mensajes procesados: 4
15:30:55.124901 [Server B] Último heartbeat hace: 1 segundos
15:30:55.125012 [Server B] ===================

15:31:00.124567 [Server B] PING enviado a Server A
15:31:00.124678 [Server A] Respuesta de 127.0.0.1:5002: PING
15:31:00.124789 [Server A] Recibido PING, enviando PONG...
15:31:00.124890 [Server A] Mensaje enviado: PONG
15:31:00.124901 [Server B] Mensaje de ServerA: PONG
15:31:00.125012 [Server B] PONG recibido, conexión bidireccional confirmada
```

## Detener el programa

Presiona `Ctrl+C` para detener gracefully ambos servidores.

## Extensiones posibles

- Añadir encriptación de mensajes
- Implementar protocolo de descubrimiento de servicios
- Añadir persistencia de estadísticas
- Implementar rate limiting
- Añadir compresión para mensajes grandes
- Crear protocolo de fragmentación para mensajes > MTU