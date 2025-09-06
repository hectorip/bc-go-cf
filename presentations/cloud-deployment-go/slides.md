---
theme: seriph
background: https://images.unsplash.com/photo-1451187580459-43490279c0fa?w=1920
title: Desplegando Aplicaciones Go en la Nube
info: |
  ## Guía completa de deployment para proyectos Go
  
  Una presentación sobre las opciones de despliegue para aplicaciones Go,
  desde servidores bare metal hasta funciones serverless.
  Aprende a elegir la plataforma correcta para tu proyecto.
class: text-center
drawings:
  persist: false
transition: slide-left
mdc: true
fonts:
  sans: 'Helvetica Neue'
  mono: 'Fira Code'
---

# Desplegando Aplicaciones Go
## De Bare Metal a Serverless

<div class="pt-12">
  <span @click="$slidev.nav.next" class="px-2 py-1 rounded cursor-pointer" hover:bg="white op-10">
    Presiona <kbd>espacio</kbd> para continuar →
  </span>
</div>

<div class="abs-br m-6 text-xl">
  <a href="https://github.com/slidevjs/slidev" target="_blank" class="slidev-icon-btn">
    <carbon:logo-github />
  </a>
</div>

<style>
h1 {
  background-color: #00ADD8;
  background-image: linear-gradient(45deg, #00ADD8 25%, #5AC8E2 50%);
  background-size: 100%;
  -webkit-background-clip: text;
  -moz-background-clip: text;
  -webkit-text-fill-color: transparent;
  -moz-text-fill-color: transparent;
}
</style>

---
transition: fade-out
layout: two-cols
---

# Agenda

<v-clicks>

- Teoría de Deployments
- Modelos de Despliegue
- Plataformas Disponibles
- Framework de Decisión

</v-clicks>

::right::

<div class="pl-10 pt-10">
  <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.svg" class="w-60 opacity-80" />
</div>

<style>
h1 {
  background-color: #00ADD8;
  background-image: linear-gradient(45deg, #00ADD8 10%, #5AC8E2 20%);
  background-size: 100%;
  -webkit-background-clip: text;
  -moz-background-clip: text;
  -webkit-text-fill-color: transparent;
  -moz-text-fill-color: transparent;
}
</style>

---
layout: center
class: text-center
---

# ¿Por qué es importante elegir bien?

<div class="text-4xl mt-10">

<v-clicks>

**Costos** pueden variar 10x

**Performance** impacta usuarios

**Mantenimiento** consume tiempo

**Escalabilidad** define el futuro

</v-clicks>

</div>

---
transition: slide-up
---

# Evolución del Deployment

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

## Antes (2000s)
- Servidores físicos propios
- Instalación manual
- Configuración artesanal
- Escalamiento = comprar hardware
- Semanas para provisionar

</div>

<div v-click>

## Ahora (2020s)
- Infraestructura como código
- CI/CD automatizado
- Auto-escalamiento
- Pago por uso
- Minutos para desplegar

</div>

</div>

---

# Teoría de Deployment: Arquetipos
## Los 4 niveles de distribución geográfica

<v-clicks>

Cada arquetipo define dónde y cómo se distribuye tu aplicación

La elección determina disponibilidad, latencia y costo

Veamos cada uno en detalle...

</v-clicks>

---

# Arquetipo Zonal
## Una sola zona de disponibilidad

<v-clicks>

- Todos los recursos en un solo datacenter
- Sin redundancia geográfica  
- Latencia mínima entre componentes

</v-clicks>

---

# Arquetipo Zonal: Casos de Uso

<v-clicks>

- Ambientes de desarrollo
- Testing y staging
- Aplicaciones internas
- Prototipos y MVPs

</v-clicks>

---

# Arquetipo Zonal: Trade-offs

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Ventajas
- **Costo mínimo**
- **Simplicidad**

</div>

<div v-click>

### Desventajas
- **Riesgo alto**
- **Sin HA**

</div>

</div>

---

# Arquetipo Regional
## Múltiples zonas en una región

<v-clicks>

- Distribuido en 2-3 zonas de disponibilidad
- Misma región geográfica
- Latencia < 2ms entre zonas

</v-clicks>

---

# Regional: Casos de Uso

<v-clicks>

- Producción estándar
- Aplicaciones B2B
- APIs regionales
- La mayoría de aplicaciones web

</v-clicks>

---

# Regional: Trade-offs

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Ventajas
- **Alta disponibilidad**
- **Balance costo-beneficio**
- **Baja latencia intra-región**

</div>

<div v-click>

### Desventajas  
- **Latencia para usuarios lejanos**
- 100ms+ cross-continental

</div>

</div>

---

# Arquetipo Multi-Regional
## Varias regiones geográficas

<v-clicks>

- Presencia en 2-4 regiones
- Replicación activa de datos
- Routing inteligente por geolocalización

</v-clicks>

---

# Multi-Regional: Casos de Uso

<v-clicks>

- E-commerce global
- SaaS con clientes internacionales
- Aplicaciones críticas de negocio
- Compliance con data residency

</v-clicks>

---

# Multi-Regional: Trade-offs

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Ventajas
- **Baja latencia global**
- **Disaster recovery**
- **Data sovereignty**

</div>

<div v-click>

### Desventajas
- **Costo 3-4x mayor**
- **Complejidad alta**

</div>

</div>

---

# Arquetipo Global
## Presencia mundial en el edge

<v-clicks>

- 50+ puntos de presencia (PoPs)
- Edge computing cerca del usuario
- CDN + compute en el edge
- Latencia < 20ms globalmente

</v-clicks>

---

# Global: Casos de Uso

<v-clicks>

- Streaming de video (Netflix, YouTube)
- Gaming online
- Redes sociales globales
- APIs de ultra-baja latencia

</v-clicks>

---

# Global: Trade-offs

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Ventajas
- **Latencia < 20ms mundial**
- **Escala masiva**
- **Resilencia extrema**

</div>

<div v-click>

### Desventajas
- **Costo 10x+**
- **Complejidad extrema**

</div>

</div>

---
layout: image-right
image: https://images.unsplash.com/photo-1558494949-ef010cbdcc31?w=1920
---

# Modelos de Responsabilidad
## La base para elegir

<v-clicks>

Cada modelo de despliegue implica diferentes niveles de responsabilidad entre tú y el proveedor

La elección correcta depende de cuánto control necesitas vs cuánto quieres gestionar

Veamos en detalle cada modelo...

</v-clicks>

---

# On-Premises
## Tú controlas todo

<v-clicks>

- Servidores físicos en tu datacenter
- Tú compras y mantienes el hardware
- Control total del stack completo

</v-clicks>

---

# On-Premises: Responsabilidades

<v-clicks>

- **Hardware**: Servidores, switches, routers
- **Datacenter**: Cooling, energía, seguridad
- **Software**: OS, parches, drivers
- **Operaciones**: Backups, DR, monitoring

</v-clicks>

---

# On-Premises: ¿Cuándo usarlo?

<v-clicks>

- Bancos con regulaciones estrictas
- Gobierno y defensa
- Datos ultra-sensibles
- Control total requerido

</v-clicks>

---

# Infrastructure as a Service (IaaS)
## El cloud maneja el hardware

<v-clicks>

- Rentas máquinas virtuales en la nube
- El proveedor maneja el hardware físico
- Tú controlas desde el OS hacia arriba

</v-clicks>

---

# IaaS: División de Responsabilidades

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Proveedor maneja
- Hardware físico
- Virtualización
- Red física
- Datacenter

</div>

<div v-click>

### Tú manejas
- Sistema operativo
- Runtime de Go
- Tu aplicación
- Seguridad del OS

</div>

</div>

---

# IaaS: Proveedores Principales

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

### Líderes del mercado

<v-clicks>

**AWS EC2**
- Más de 400 tipos de instancias
- Presencia global (30+ regiones)
- Integración con 200+ servicios AWS

**Google Compute Engine**
- Facturación por segundo
- Live migration (sin downtime)
- Precios competitivos

**Azure Virtual Machines**
- Mejor integración con Microsoft
- Hybrid cloud con Azure Arc
- Windows Server optimizado

</v-clicks>

</div>

<div>

### Alternativas económicas

<v-clicks>

**DigitalOcean Droplets**
- Desde $4/mes
- Interface simple
- Popular con startups

**Linode**
- $5/mes plan básico
- Buen soporte
- 11 datacenters globales

**Vultr**
- $2.50/mes el más barato
- Deploy en 60 segundos
- 25+ ubicaciones

</v-clicks>

</div>

</div>

---

# Platform as a Service (PaaS)
## El cloud maneja la plataforma

<v-clicks>

- Solo subes tu código
- La plataforma maneja todo lo demás
- Tú te enfocas en tu aplicación

</v-clicks>

---

# PaaS: ¿Qué incluye?

<v-clicks>

- Sistema operativo gestionado
- Runtime de Go actualizado
- Balanceo de carga automático
- SSL/TLS certificates
- Auto-scaling incluido

</v-clicks>

---

# PaaS: Tu única responsabilidad

<v-clicks>

- Tu código de aplicación
- Configuración de la app
- Los datos de tu aplicación

<span v-mark.highlight.yellow>Todo lo demás es automático</span>

</v-clicks>

---

# PaaS: Proveedores y Características

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

### Tradicionales

<v-clicks>

**Heroku**
- Pioneer del git push deployment
- Buildpacks para Go incluido
- Add-ons marketplace (DB, cache, etc.)
- $7-25/mes por dyno

**Google App Engine**
- Standard: Auto-scale to zero
- Flexible: Contenedores personalizados
- Integración nativa con GCP
- Free tier generoso

**AWS Elastic Beanstalk**
- Maneja EC2, RDS, ELB automáticamente
- Soporte para Go
- Integración con servicios AWS
- Pagas por recursos subyacentes

</v-clicks>

</div>

<div>

### Nueva generación

<v-clicks>

**Render**
- Heroku moderno con precios justos
- Free tier para proyectos pequeños
- Preview environments automáticos
- Desde $7/mes

**Railway**
- Deploy en segundos
- Detección automática de Go
- Colaboración en tiempo real
- $5 de crédito gratis/mes

**Fly.io**
- Deploy global en edge locations
- Firecracker microVMs
- Persistent volumes disponibles
- Cobra por uso real

</v-clicks>

</div>

</div>

---

# Container as a Service (CaaS)
## El cloud orquesta tus contenedores

<v-clicks>

- Creas contenedores Docker
- El proveedor los orquesta
- Punto medio entre PaaS y IaaS

</v-clicks>

---

# CaaS: Responsabilidades

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Proveedor
- Infraestructura
- Orquestación
- Networking
- Registry

</div>

<div v-click>

### Tú
- Dockerfile
- Configuración
- Recursos
- Secretos

</div>

</div>

---

# CaaS: Kubernetes Administrado

<div class="grid grid-cols-3 gap-6 mt-8">

<div>

### AWS EKS

<v-clicks>

- Kubernetes certificado
- $0.10/hora por cluster
- Integración con IAM
- Fargate para serverless pods

</v-clicks>

</div>

<div>

### Google GKE

<v-clicks>

- Creador de Kubernetes
- Autopilot mode disponible
- $0.10/hora (gratis el primero)
- Mejor experiencia

</v-clicks>

</div>

<div>

### Azure AKS

<v-clicks>

- Control plane gratis
- Azure AD integrado
- Windows containers
- Azure Arc enabled

</v-clicks>

</div>

</div>

<v-click>

### Alternativas más simples a K8s

<div class="grid grid-cols-2 gap-8 mt-4">

<div>

**AWS ECS + Fargate**
- Más simple que Kubernetes
- Sin gestión de servidores
- Pago por tarea ejecutada

</div>

<div>

**Google Cloud Run**
- Contenedores serverless
- Scale to zero automático
- Ideal para APIs Go

</div>

</div>

</v-click>

---

# Functions as a Service (FaaS)
## Serverless puro

<v-clicks>

- Escribes funciones individuales
- Se ejecutan por eventos
- Pagas por milisegundos
- Cero administración

</v-clicks>

---

# FaaS: Lo único que haces

<v-clicks>

1. Escribir tu código de función
2. Configurar los triggers

<span v-mark.circle.orange>Todo lo demás es automático</span>

</v-clicks>

---
transition: slide-left
---

# Go en AWS Lambda

<v-clicks>

- Soporte nativo para Go
- 1M requests gratis/mes
- Cold start ~100ms
- Hasta 15 min de ejecución

</v-clicks>

---

# FaaS: Más Opciones

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

### Con soporte nativo de Go

<v-clicks>

**Google Cloud Functions**
- Go 1.11+
- Triggers: HTTP, Pub/Sub, Storage
- 2M invocaciones gratis/mes
- Integración con Firebase

**IBM Cloud Functions**
- Basado en Apache OpenWhisk
- Go support incluido
- Pago por GB-segundo
- Buena integración con Watson

</v-clicks>

</div>

<div>

### Requieren workarounds

<v-clicks>

**Azure Functions**
- Custom handlers para Go
- Más complejo de configurar
- Mejor para .NET/Node.js

**Cloudflare Workers**
- Requiere compilar Go a WASM
- Ejecución en edge (0ms latency)
- No es Go nativo
- Limitaciones de WASM

**Vercel/Netlify Functions**
- Principalmente para JavaScript
- Posible con Go pero no ideal
- Mejor para JAMstack

</v-clicks>

</div>

</div>

---

# Resumen: ¿Quién maneja qué?

<div class="text-sm mt-4">

| Responsabilidad | On-Prem | IaaS | PaaS | CaaS | FaaS |
|----------------|---------|------|------|------|------|
| **Tu código Go** | Tú | Tú | Tú | Tú | Tú |
| **Datos** | Tú | Tú | Tú | Tú | Tú |
| **Runtime Go** | Tú | Tú | Cloud | Tú* | Cloud |
| **Container** | Tú | Tú | Cloud | Tú | Cloud |
| **OS** | Tú | Tú | Cloud | Cloud | Cloud |
| **Virtualización** | Tú | Cloud | Cloud | Cloud | Cloud |
| **Servidores** | Tú | Cloud | Cloud | Cloud | Cloud |
| **Almacenamiento** | Tú | Cloud | Cloud | Cloud | Cloud |
| **Red física** | Tú | Cloud | Cloud | Cloud | Cloud |
| **Datacenter** | Tú | Cloud | Cloud | Cloud | Cloud |

</div>

<v-click>

<div class="mt-4 p-3 bg-yellow-500 bg-opacity-10 rounded">
*En CaaS, el runtime Go está en tu contenedor, pero la gestión del contenedor es del cloud
</div>

</v-click>

---

# Bare Metal
## Control Total, Máxima Responsabilidad

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Ventajas
- Control completo del hardware
- Máximo rendimiento
- Aislamiento físico
- Sin overhead de virtualización

</div>

<div v-click>

### Desventajas
- Costoso de mantener
- Escalamiento lento
- Requiere expertise
- Sin elasticidad

</div>

</div>

---

# Bare Metal: Cuándo Usar

### Solo si realmente necesitas:

<v-clicks>

- Hardware especializado (GPU, FPGA, ASIC)
- Latencia ultra-baja (< 1ms)
- Compliance estricto que requiere aislamiento físico
- Control total sobre el stack de hardware
- Aplicaciones de alto rendimiento computacional

</v-clicks>

<v-click>

<div class="mt-8 p-4 bg-yellow-500 bg-opacity-10 rounded">

### Solo si realmente necesitas:

- Hardware especializado
- Latencia < 1ms
- Compliance estricto
- Aislamiento físico total

</div>

</v-click>

---

# Virtual Machines (IaaS)
## Flexibilidad sin Hardware

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Ventajas
- Control total del OS
- Flexibilidad en configuración
- Escalamiento horizontal
- Sin hardware físico
- Pago por hora

</div>

<div v-click>

### Desventajas
- Mantenimiento del OS
- Gestión de parches
- Configuración manual
- Costo 24/7

</div>

</div>

---

# IaaS: Opciones Principales

<div class="grid grid-cols-3 gap-4 mt-8">

<v-click>

### AWS EC2
- **AWS EC2** - Líder del mercado

</v-click>

<v-click>

### Google Compute
- **Google Compute Engine** - Buena integración con GCP

</v-click>

<v-click>

### Azure VMs
- **Azure VMs** - Ideal para empresas Microsoft
- **DigitalOcean Droplets** - Simple y económico
- **Linode/Vultr** - Alternativas cost-effective

</v-click>

</div>

<v-click>

<div class="mt-8 grid grid-cols-2 gap-4">

<div class="p-4 bg-green-500 bg-opacity-10 rounded">
**Ideal para**: Migraciones, control del OS, software legacy
</div>

<div class="p-4 bg-red-500 bg-opacity-10 rounded">
**Evitar si**: Quieres simplicidad, no tienes equipo DevOps
</div>

</div>

</v-click>

---
layout: center
---

# Containers & Kubernetes
## El estándar de la industria

<div class="text-xl mt-8">

<v-clicks>

Los contenedores encapsulan tu aplicación Go y sus dependencias

Kubernetes orquesta y gestiona estos contenedores a escala

Es el estándar de facto para aplicaciones cloud-native

</v-clicks>

</div>

---

# Dockerfile para Go

### Ventajas de containerizar aplicaciones Go

<v-clicks>

- **Portabilidad total**: Funciona igual en cualquier ambiente
- **Imagen mínima**: Go compila a binario estático (~10MB)
- **Rápido arranque**: Sin JVM ni interpretes (< 1s)
- **Fácil distribución**: Un solo artefacto inmutable
- **Aislamiento**: Cada servicio en su contenedor

</v-clicks>

<v-click>

### Mejores prácticas

- Usar multi-stage builds
- Imágenes desde scratch o alpine
- Compilar con CGO_ENABLED=0
- No ejecutar como root

</v-click>

<v-click>

<div class="mt-4 grid grid-cols-2 gap-4">

<div>
**Imagen final**: ~10MB
</div>

<div>
**Startup time**: <1s
</div>

</div>

</v-click>

---

# Kubernetes
## Orquestación de Contenedores

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Ventajas
- Orquestación automática
- Auto-healing y scaling
- Service discovery
- Rolling updates
- Rico ecosistema

</div>

<div v-click>

### Desventajas
- Alta complejidad
- Curva empinada
- Overhead para MVPs
- Requiere DevOps
- Costo mínimo alto

</div>

</div>

---

# Kubernetes: ¿Realmente lo necesitas?

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

## Sí, si tienes:

<v-clicks>

- 10+ microservicios
- Equipo DevOps dedicado
- Necesidad de orquestación compleja
- Presupuesto para la curva de aprendizaje

</v-clicks>

</div>

<div>

## No, si:

<v-clicks>

- Es tu primer proyecto
- Tienes 1-3 servicios
- Equipo pequeño (<5 devs)
- Presupuesto limitado

</v-clicks>

</div>

</div>

<v-click>

<div class="mt-8 p-4 bg-orange-500 bg-opacity-10 rounded text-center">
<span v-mark.highlight.yellow>"No uses Kubernetes porque Google lo usa. Úsalo cuando tengas problemas similares a Google"</span>
</div>

</v-click>

---

# Platform as a Service
## Deploy en minutos, no días

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Ventajas
- Deploy con git push
- Sin servidores
- SSL automático
- CI/CD integrado
- Add-ons fáciles

</div>

<div v-click>

### Desventajas
- Menos control
- Costo premium
- Límites de plataforma
- Vendor lock-in
- Menos optimización

</div>

</div>

---
layout: two-cols
---

# PaaS: Opciones Populares

<v-clicks>

### Heroku
- Deploy con git push
- Buildpacks automáticos para Go
- Add-ons marketplace
- $7/mes plan básico

### Render
- Alternativa moderna a Heroku
- Free tier disponible
- Auto-deploy desde GitHub
- SSL automático

</v-clicks>

::right::

<v-click>

### Railway
- Detección automática de Go
- Deploy instantáneo
- Variables de entorno fáciles
- Pricing por uso

</v-click>

<v-click>

### Fly.io
- Deploy global en múltiples regiones
- Contenedores en Firecracker VMs
- Escala automáticamente
- $5 crédito mensual gratis

</v-click>

<v-click>

<div class="mt-8 p-3 bg-green-500 bg-opacity-10 rounded">
**Perfecto para**: MVPs, startups, prototipos
</div>

</v-click>

---

# Serverless: El futuro es sin servidor

<div class="mt-8">

### Características principales

<v-clicks>

- **Pago por uso real**: $0 cuando no hay tráfico
- **Escalamiento automático**: De 0 a miles de instancias
- **Sin gestión de servidores**: El proveedor maneja todo
- **Alta disponibilidad**: Multi-AZ por defecto
- **Ideal para Go**: Cold starts rápidos (~100ms)

</v-clicks>

<v-click>

### Opciones principales

- **AWS Lambda**: Líder del mercado, gran ecosistema
- **Google Cloud Functions**: Integración con GCP
- **Azure Functions**: Requiere custom handlers para Go
- **Cloudflare Workers**: Requiere compilar a WASM

</v-click>

</div>

<v-click>

<div class="grid grid-cols-3 gap-4 mt-8">

<div class="text-center">
<div class="text-3xl">Costo</div>
**$0** cuando no hay tráfico
</div>

<div class="text-center">
<div class="text-3xl">Escalabilidad</div>
Escala a **miles** de requests/seg
</div>

<div class="text-center">
<div class="text-3xl">Velocidad</div>
Cold start Go: **~100ms**
</div>

</div>

</v-click>

---

# Google Cloud Run
## Lo mejor de ambos mundos

<v-clicks>

- **Containers** pero sin Kubernetes
- **Serverless** pricing (scale to zero)
- **Global** con un comando
- **Simple** como PaaS

</v-clicks>

<v-click>

### Características clave
- Deploy en 2 comandos simples
- Autoscaling de 0 a miles de instancias
- HTTPS automático incluido
- Integración con CI/CD de Google

</v-click>

<v-click>

<div class="mt-8 grid grid-cols-2 gap-4">

<div class="p-4 bg-blue-500 bg-opacity-10 rounded">
**Free tier generoso**
- 2M requests/mes
- 360,000 GB-segundos
- 180,000 vCPU-segundos
</div>

<div class="p-4 bg-green-500 bg-opacity-10 rounded">
**Ideal para**
- APIs REST
- Webhooks
- Microservicios
- Sites dinámicos
</div>

</div>

</v-click>

---
layout: center
---

# I/O Bound vs CPU Bound
## La clave para elegir dónde deployar

<v-clicks>

### ¿Por qué importa?

El tipo de carga de tu aplicación Go determina:
- Qué recursos necesitas optimizar
- Cuánto pagarás en la nube
- Qué plataforma te dará mejor performance

### La gran ventaja de Go

Go fue diseñado para I/O concurrente con goroutines
- Millones de goroutines con poca memoria
- Modelo CSP (Communicating Sequential Processes)
- Perfect para microservicios y APIs

</v-clicks>

---

# I/O Bound: Esperando al mundo

<v-clicks>

- La mayor parte del tiempo **esperando**
- Base de datos queries
- Llamadas a APIs externas
- Lectura/escritura de archivos
- WebSockets y streaming

</v-clicks>

---

# I/O Bound: Ejemplos en Go

<v-clicks>

- API REST que consulta PostgreSQL
- Proxy/Gateway que llama otros servicios
- Chat server con WebSockets
- Servicio de uploads a S3
- Web scraper

</v-clicks>

---

# I/O Bound: Lo que necesitas

<v-clicks>

- **Muchas conexiones concurrentes**
- **Baja latencia de red**
- **No mucho CPU**

</v-clicks>

---

# CPU Bound: Calculando sin parar

<v-clicks>

- Uso intensivo del procesador
- Cálculos matemáticos complejos
- Procesamiento de datos
- Transformaciones pesadas
- Algoritmos computacionales

</v-clicks>

---

# CPU Bound: Ejemplos en Go

<v-clicks>

- Procesamiento de imágenes/video
- Compresión/encriptación
- Machine Learning inference
- Análisis de datos grandes
- Compiladores/parsers

</v-clicks>

---

# CPU Bound: Lo que necesitas

<v-clicks>

- **CPUs potentes o múltiples cores**
- **Memoria para datasets**
- **Posiblemente GPUs**

</v-clicks>

---

# Go y la Concurrencia I/O

<v-clicks>

### Por qué Go brilla en I/O bound

**Goroutines baratas**
- Solo 2KB de stack inicial
- Puedes tener 100,000+ goroutines
- El runtime las multiplexea eficientemente

**Non-blocking I/O nativo**
- net/http server ya es concurrente
- database/sql con connection pooling
- io.Reader/Writer interfaces async-friendly

### El antipatrón
```go
// MAL: Bloquea un OS thread completo
for {
    conn, _ := listener.Accept()
    go handleConnection(conn) // BIEN: Goroutine barata
}
```

</v-clicks>

---

# Decisión I/O Bound: Mejores Opciones

<v-clicks>

**1. Serverless (Lambda, Cloud Run)**
- Pagas solo cuando procesas
- Escala automática con requests

**2. PaaS (Heroku, Render)**
- Múltiples conexiones fácil
- Connection pooling incluido

**3. Kubernetes con HPA**
- Autoscaling por requests/segundo
- Múltiples pods pequeños

</v-clicks>

---

# I/O Bound: Configuración Óptima

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### Serverless
- Memory: 256-512MB
- Timeout: 30s-5min
- Concurrency: 1000/inst

</div>

<div v-click>

### Containers/VMs
- vCPUs: 0.5-2
- RAM: 512MB-2GB
- Muchas instancias pequeñas

</div>

</div>

---

# Decisión CPU Bound: Mejores Opciones

<v-clicks>

**1. VMs con CPU optimizado**
- EC2 C5/C6 instances
- GCP C2 machines
- Azure F-series

**2. Kubernetes con pods grandes**
- Guaranteed QoS
- CPU requests = limits

**3. Bare Metal**
- Para ML training
- Procesamiento 24/7

</v-clicks>

---

# CPU Bound: Configuración Óptima

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### VMs potentes
- vCPUs: 8-64
- RAM: 16GB+
- Storage: SSD NVMe

</div>

<div v-click>

### Go runtime
- GOMAXPROCS = NumCPU
- Worker pools sized to cores
- Parallel algorithms

</div>

</div>

---

# CPU Bound: Qué Evitar

<v-clicks>

- Serverless (límites de tiempo)
- Instancias compartidas
- Burstable performance

<span v-mark.underline.red>Usa recursos dedicados</span>

</v-clicks>

---

# Casos Híbridos: I/O + CPU

<v-clicks>

### Ejemplos comunes
- API que procesa imágenes on-demand
- ETL pipeline que transforma datos
- Servicio de reportes con cálculos pesados

### Estrategia: Separar componentes

**Frontend (I/O bound)** → Serverless/PaaS
- Recibe requests
- Valida input
- Encola trabajos

**Backend (CPU bound)** → VMs/Kubernetes Jobs
- Procesa en background
- Workers con CPU dedicado
- Autoscaling por queue depth

### Arquitectura recomendada
Cloud Run (API) → Pub/Sub → GKE Jobs (Processing)

</v-clicks>

---

# Benchmarking: Conoce tu Aplicación

<v-clicks>

### Herramientas para Go

**pprof** - Profiling nativo
```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```
Te muestra exactamente dónde gastas CPU

**hey** - Load testing
```bash
hey -n 10000 -c 100 http://api.example.com
```
Simula carga real para ver comportamiento

**metrics** - Runtime stats
```go
runtime.NumGoroutine() // Cuántas goroutines
runtime.MemStats       // Uso de memoria
```

### La regla de oro
> "Mide primero, optimiza después, elige plataforma al final"

</v-clicks>

---

# Matriz de Decisión I/O vs CPU

<div class="text-sm mt-4">

| Característica | I/O Bound | CPU Bound | Híbrido |
|---------------|-----------|-----------|----------|
| **Serverless** | ⭐⭐⭐⭐⭐ Ideal | ⭐ No recomendado | ⭐⭐⭐ Para frontend |
| **PaaS** | ⭐⭐⭐⭐ Muy bueno | ⭐⭐ Costoso | ⭐⭐⭐ Posible |
| **Kubernetes** | ⭐⭐⭐⭐ Flexible | ⭐⭐⭐⭐ Bueno | ⭐⭐⭐⭐⭐ Ideal |
| **VMs** | ⭐⭐⭐ Funciona | ⭐⭐⭐⭐⭐ Ideal | ⭐⭐⭐⭐ Bueno |
| **Bare Metal** | ⭐⭐ Overkill | ⭐⭐⭐⭐⭐ Máximo | ⭐⭐⭐ Complejo |

</div>

<v-click>

### Ejemplos de configuración

| Tipo | Plataforma | Config | Costo/mes |
|------|------------|--------|-----------|
| **API REST** (I/O) | Cloud Run | 256MB, 0.5 CPU | $0-50 |
| **Image Processing** (CPU) | EC2 c5.2xlarge | 8 vCPU, 16GB | $250 |
| **Data Pipeline** (Híbrido) | GKE | 2 + 8 nodes | $500 |

</v-click>

---

# Anti-patrones a Evitar

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

### Para I/O Bound

<v-clicks>

❌ **No hagas esto**
- VMs gigantes para APIs simples
- Reserved instances para tráfico variable
- Bare metal para microservicios

✅ **Mejor haz esto**
- Serverless o PaaS
- Autoscaling agresivo
- Múltiples instancias pequeñas

</v-clicks>

</div>

<div>

### Para CPU Bound

<v-clicks>

❌ **No hagas esto**
- Lambda para procesamiento largo
- Instancias burstable
- Compartir CPU con I/O

✅ **Mejor haz esto**
- VMs dedicadas o bare metal
- CPU optimized instances
- Worker pools del tamaño de cores

</v-clicks>

</div>

</div>

---

# Framework de Decisión

### Preguntas clave para decidir

<v-clicks>

**1. ¿Cómo es tu tráfico?**
- Constante 24/7 → IaaS o PaaS
- Variable/picos → Serverless

**2. ¿Cuánto control necesitas?**
- Total → IaaS o Bare Metal
- Mínimo → PaaS o Serverless

**3. ¿Cuántos servicios tienes?**
- 1-3 servicios → PaaS
- 10+ servicios → Kubernetes

**4. ¿Qué tipo de aplicación?**
- API REST → Cloud Run/App Runner
- Eventos → Lambda/Functions
- Web tradicional → PaaS

</v-clicks>

---

# Factores Clave de Decisión

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

## Costo

<v-clicks>

- **Serverless**: Paga por uso real
- **PaaS**: Costo fijo mensual
- **IaaS**: Por hora (24/7 = caro)
- **Bare Metal**: CAPEX alto

</v-clicks>

</div>

<div>

## Escalabilidad

<v-clicks>

- **Serverless**: Automática e infinita
- **PaaS**: Semi-automática
- **Kubernetes**: Configurable
- **VMs**: Manual o scripted

</v-clicks>

</div>

</div>

<v-click>

<div class="mt-8 p-4 bg-blue-500 bg-opacity-10 rounded">

### Regla de oro

<span v-mark.circle.orange>"Empieza con la mayor abstracción que cumpla tus requisitos"</span>

</div>

</v-click>

---

# Análisis de Costos Reales

<div class="mt-8">

| Plataforma | 1K req/día | 100K req/día | 1M req/día |
|------------|------------|--------------|------------|
| **Lambda** | $0 | $2 | $20 |
| **Cloud Run** | $0 | $0 | $15 |
| **Heroku** | $7 | $7 | $25-250 |
| **EC2 t3.micro** | $8 | $8 | $8-80* |
| **EKS (3 nodos)** | $216 | $216 | $216 |

<div class="text-sm mt-2">*Requiere escalar horizontalmente</div>

</div>

<v-click>

<div class="mt-8 grid grid-cols-2 gap-4">

<div class="p-4 bg-green-500 bg-opacity-10 rounded">
**Tráfico variable**: Serverless gana
</div>

<div class="p-4 bg-orange-500 bg-opacity-10 rounded">
**Tráfico constante alto**: VMs pueden ser más baratas
</div>

</div>

</v-click>

---
layout: two-cols
---

# Caso 1: Startup SaaS

## Aplicación B2B

**Requisitos:**
- 100-1000 usuarios
- Necesita escalar rápido
- Presupuesto limitado
- Time to market crítico

<v-click>

### Solución: PaaS

- Deploy automático desde GitHub
- Escalamiento elástico incluido  
- Base de datos como add-on
- SSL y dominio personalizado incluidos

**Costo**: $19-50/mes

</v-click>

::right::

# Caso 2: E-commerce

## Tienda online

**Requisitos:**
- Picos en Black Friday
- Global (baja latencia)
- Costo-eficiente
- Alta disponibilidad

<v-click>

### Solución: Cloud Run + CDN

- Google Cloud Run para la API
- Cloudflare CDN para contenido estático
- Deploy multi-región automático
- Scale to zero cuando no hay tráfico

**Costo**: $0-500/mes (por uso)

</v-click>

---

# Caso 3: Procesamiento de Datos

<div class="grid grid-cols-2 gap-8">

<div>

## ETL Pipeline

**Requisitos:**
- Procesa 1TB/día
- Ejecución programada
- No necesita estar siempre activo

<v-click>

### Solución: Serverless

- AWS Lambda o Cloud Functions
- Triggered por eventos o schedule
- Procesamiento paralelo automático
- Sin servidores idle
- Integración nativa con data warehouses

</v-click>

</div>

<div>

<v-click>

## Comparación de costos

| Opción | Costo Mensual |
|--------|---------------|
| **Serverless** | $50 |
| **VM 24/7** | $300 |
| **Kubernetes** | $500 |

<div class="mt-4 p-3 bg-green-500 bg-opacity-10 rounded">
**Ahorro**: 83% vs VM always-on
</div>

</v-click>

</div>

</div>

---

# Plataformas Disponibles (2024)

<div class="grid grid-cols-4 gap-3 text-sm mt-6">

<div>

### IaaS/VMs
- AWS EC2
- Google Compute
- Azure VMs
- DigitalOcean
- Linode
- Vultr

</div>

<div>

### PaaS
- Heroku
- Render
- Railway
- Fly.io
- Google App Engine
- Azure App Service

</div>

<div>

### Serverless
- AWS Lambda
- Google Cloud Run
- Azure Functions
- Vercel
- Netlify Functions
- Cloudflare Workers*

</div>

<div>

### Kubernetes
- AWS EKS
- Google GKE
- Azure AKS
- DigitalOcean K8s
- IBM Cloud K8s
- Red Hat OpenShift

</div>

</div>

<v-click>

<div class="mt-8 p-4 bg-yellow-500 bg-opacity-10 rounded">
*Cloudflare Workers requiere compilar Go a WASM
</div>

</v-click>

---

# Ejemplo Práctico: API REST

<div class="grid grid-cols-2 gap-4">

<div>

## Código Go

### Una simple API Go puede desplegarse en:

<v-clicks>

- **Heroku**: Git push y listo
- **Cloud Run**: Contenedor serverless
- **Fly.io**: Deploy global instantáneo
- **Railway**: Detección automática
- **AWS Lambda**: Función serverless
- **VMs**: Máximo control

</v-clicks>

</div>

<div>

## Deploy Options

<v-clicks>

### La elección depende de:

- Presupuesto disponible
- Experiencia del equipo  
- Requisitos de escalamiento
- Necesidades de control

</v-clicks>

</div>

</div>

---

# Migración entre Plataformas

### Evolución típica de un proyecto

<v-clicks>

**Fase 1: MVP**
- Empieza con PaaS o Serverless
- Enfoque en desarrollo rápido
- Costos mínimos

**Fase 2: Crecimiento**
- Considera containers si necesitas más control
- Evalúa Kubernetes si tienes múltiples servicios
- Optimiza costos con instancias reservadas

**Fase 3: Madurez**
- IaaS para optimización máxima
- Arquitectura híbrida según necesidades
- Balance entre costo y complejidad

</v-clicks>

<v-click>

<div class="mt-8 grid grid-cols-3 gap-4">

<div class="p-3 bg-blue-500 bg-opacity-10 rounded">
**Fase 1**: MVP en PaaS
</div>

<div class="p-3 bg-orange-500 bg-opacity-10 rounded">
**Fase 2**: Escalar con containers
</div>

<div class="p-3 bg-red-500 bg-opacity-10 rounded">
**Fase 3**: Optimizar con IaaS
</div>

</div>

</v-click>

---

# Monitoreo y Observabilidad

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

## Métricas Clave

<v-clicks>

- **Latencia**: P50, P95, P99
- **Throughput**: Requests/segundo
- **Error rate**: % de fallos
- **Costo**: $/request

</v-clicks>

</div>

<div>

## Herramientas por Plataforma

<v-clicks>

- **Serverless**: CloudWatch, Stackdriver
- **PaaS**: Incluido (Heroku Metrics, etc.)
- **Kubernetes**: Prometheus + Grafana
- **VMs**: Datadog, New Relic

</v-clicks>

</div>

</div>

<v-click>

### Integración de monitoreo

- Prometheus + Grafana para Kubernetes
- CloudWatch para AWS
- Stackdriver para GCP
- Application Insights para Azure
- Datadog/New Relic para cualquier plataforma

</v-click>

---

# Manejo de Secretos en Go
## El problema más crítico en cloud

<v-clicks>

### El desastre más común

**Credenciales en el código = Game Over**
- Miles de repos con API keys en GitHub
- Bots escanean constantemente buscando secrets
- En minutos: crypto-miners usando tu cuenta

### Los números del horror
- 100,000+ repos con secrets expuestos (GitGuardian 2023)
- Tiempo promedio de detección por bots: **20 segundos**
- Costo promedio por incidente: $1.2M

### La regla #1
> "Si está en git, ya no es un secreto"

</v-clicks>

---

# Anatomía de un Desastre de Secretos

<v-clicks>

### Caso real: Uber 2016
- Credenciales AWS en GitHub privado
- Hackers accedieron a S3
- 57 millones de registros expuestos
- Multa: $148 millones

### Caso real: CircleCI 2023
- Secretos de clientes comprometidos
- Todos tuvieron que rotar TODAS las credenciales
- Impacto: miles de empresas

### El patrón común
1. Developer comete secreto "temporalmente"
2. Lo olvida o piensa que el repo es privado
3. Repo se hace público o es comprometido
4. Desastre instantáneo

</v-clicks>

---

# Manejo de Secretos: Por Plataforma

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

### Variables de Entorno

<v-clicks>

**IaaS (VMs)**
- Archivos .env (NUNCA en git)
- SystemD environment files
- Cloud-init user data (encriptado)

**PaaS**
- Heroku: `heroku config:set`
- Render: Dashboard UI
- Railway: Variables en UI
- Fly.io: `fly secrets set`

</v-clicks>

</div>

<div>

### Servicios Nativos

<v-clicks>

**AWS**
- Secrets Manager
- Systems Manager Parameter Store
- KMS para encriptación

**Google Cloud**
- Secret Manager
- Cloud KMS

**Azure**
- Key Vault
- Managed Identity

</v-clicks>

</div>

</div>

---

# Secretos en Kubernetes

<v-clicks>

### Opciones disponibles

**Kubernetes Secrets nativos**
- Base64 encoded (NO es encriptación)
- Almacenados en etcd
- Necesitan encriptación at rest

**Sealed Secrets**
- Encriptados que pueden ir en Git
- Controller los desencripta en cluster

**External Secrets Operator**
- Sincroniza desde AWS/GCP/Azure
- Rotación automática
- Single source of truth

**HashiCorp Vault**
- Gestión empresarial
- Rotación dinámica
- Auditoría completa

</v-clicks>

---

# Mejores Prácticas con Go

<v-clicks>

### 1. Nunca hardcodees secretos

❌ **MAL**
```go
apiKey := "sk-1234567890abcdef"
```

✅ **BIEN**
```go
apiKey := os.Getenv("API_KEY")
```

### 2. Usa bibliotecas especializadas

**Viper** - Configuración flexible
- Lee de ENV, archivos, Consul, etcd
- Valores por defecto
- Hot reload de config

**godotenv** - Para desarrollo local
- Lee archivos .env
- SOLO para desarrollo
- .env en .gitignore SIEMPRE

</v-clicks>

---

# Patrón de Configuración Segura en Go

<v-clicks>

### Estructura recomendada

1. **Config struct centralizada**
   - Todos los secretos en un lugar
   - Validación al inicio
   - Panic si falta algo crítico

2. **Inicialización temprana**
   - Cargar en main()
   - Validar antes de iniciar servidor
   - Fail fast si hay problemas

3. **Jerarquía de fuentes**
   - Defaults → Config file → ENV → Secret Manager
   - ENV siempre gana
   - Secretos nunca en archivos

4. **Validación estricta**
   - Verificar formato
   - Verificar permisos
   - Verificar conectividad

</v-clicks>

---

# Rotación de Secretos

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

### Cuándo rotar

<v-clicks>

- **Inmediatamente** si hay exposición
- **Regularmente** (30-90 días)
- **Al cambiar personal** con acceso
- **Después de auditorías**

</v-clicks>

</div>

<div>

### Estrategia de rotación

<v-clicks>

1. Crear nuevo secreto
2. Actualizar aplicación para aceptar ambos
3. Desplegar con dual support
4. Cambiar al nuevo secreto
5. Remover el viejo
6. Verificar que todo funciona

</v-clicks>

</div>

</div>

<v-click>

### Automatización
- AWS Secrets Manager: Rotación automática con Lambda
- GCP Secret Manager: Versioning automático
- Vault: Dynamic secrets con TTL

</v-click>

---

# Herramientas de Detección

<v-clicks>

### Pre-commit hooks
**git-secrets** (AWS)
- Bloquea commits con patterns peligrosos
- Configurable con regex
- Prevención en origen

**detect-secrets** (Yelp)
- Escanea antes de commit
- Baseline para ignorar falsos positivos
- Integración con CI/CD

### Scanning continuo
**GitGuardian**
- Monitoreo en tiempo real
- Alertas instantáneas
- Revocación automática

**TruffleHog**
- Escanea historia completa de Git
- Detecta secrets eliminados
- Verifica si siguen activos

</v-clicks>

---

# Ejemplo: Config Segura para Go

<v-clicks>

### Estructura de proyecto
```
myapp/
├── .env.example     # Template con valores fake
├── .gitignore       # INCLUYE .env
├── config/
│   └── config.go    # Carga configuración
├── main.go
└── go.mod
```

### .env.example (SÍ va en Git)
```
DATABASE_URL=postgres://user:pass@localhost/db
API_KEY=your-api-key-here
JWT_SECRET=your-secret-here
```

### .env (NUNCA en Git)
```
DATABASE_URL=postgres://prod:real@rds.aws/proddb
API_KEY=sk-real-key-123456
JWT_SECRET=actual-secret-key
```

</v-clicks>

---

# Checklist: Antes de Hacer Commit

<v-clicks>

- [ ] ¿Ejecuté git-secrets?
- [ ] ¿El .env está en .gitignore?
- [ ] ¿No hay IPs/URLs de producción?

</v-clicks>

---

# Checklist: Antes de Deployment

<v-clicks>

- [ ] ¿Secretos en secret manager?
- [ ] ¿Variables de entorno configuradas?
- [ ] ¿Permisos IAM mínimos?

</v-clicks>

---

# Checklist: En Producción

<v-clicks>

- [ ] ¿Rotación automática?
- [ ] ¿Alertas por acceso?
- [ ] ¿Logs de auditoría?

</v-clicks>

---

# Si Hay una Brecha

<v-clicks>

1. Rotar TODOS los secretos
2. Auditar logs de 90 días
3. Notificar a seguridad
4. Documentar el incidente

<span v-mark.circle.red>ACTUAR RÁPIDO</span>

</v-clicks>

---

# Seguridad por Capa

<div class="mt-8">

| Capa | Bare Metal | IaaS | PaaS | Serverless |
|------|------------|------|------|------------|
| **Red** | Tú | Tú | Provider | Provider |
| **OS** | Tú | Tú | Provider | Provider |
| **Runtime** | Tú | Tú | Provider | Provider |
| **App** | Tú | Tú | Tú | Tú |
| **Datos** | Tú | Tú | Tú | Tú |

</div>

<v-click>

<div class="mt-8 grid grid-cols-2 gap-4">

<div class="p-4 bg-yellow-500 bg-opacity-10 rounded">
**Siempre tu responsabilidad**:
- Seguridad del código
- Gestión de secretos
- Autenticación/Autorización
</div>

<div class="p-4 bg-green-500 bg-opacity-10 rounded">
**Mejores prácticas**:
- Usar variables de entorno
- Rotar secretos regularmente
- Principio de menor privilegio
</div>

</div>

</v-click>

---

# Performance: Go en Diferentes Plataformas

<div class="mt-8">

### Cold Start (tiempo de arranque en frío)

<v-clicks>

- **AWS Lambda Go**: 100-200ms
- **Cloud Run**: 200-500ms 
- **Heroku**: 1000-3000ms
- **VM/Kubernetes**: No aplica (siempre activo)

</v-clicks>

<v-click>

### Throughput máximo (requests/segundo)

- **Lambda**: ~1000 por instancia
- **Cloud Run**: ~1000 por contenedor
- **PaaS**: 100-500 por dyno
- **VM optimizada**: 5000-10000+

</v-click>

</div>

<v-click>

<div class="mt-4 p-4 bg-blue-500 bg-opacity-10 rounded">
**Go advantages**: Compilado, binario único, baja memoria, rápido startup
</div>

</v-click>

---

# ¿Cuándo Usar Serverless?

<v-clicks>

- [ ] Tráfico variable o impredecible
- [ ] Quieres pagar solo por uso
- [ ] No tienes equipo DevOps
- [ ] Necesitas escalar automáticamente

</v-clicks>

---

# ¿Cuándo Usar PaaS?

<v-clicks>

- [ ] Necesitas deploy rápido
- [ ] Tienes presupuesto fijo
- [ ] Quieres abstracciones simples
- [ ] Tu app es estándar (web/API)

</v-clicks>

---

# ¿Cuándo Usar Kubernetes?

<v-clicks>

- [ ] Tienes muchos microservicios
- [ ] Necesitas orquestación compleja
- [ ] Tienes equipo DevOps
- [ ] Requieres alta personalización

</v-clicks>

---

# ¿Cuándo Usar VMs?

<v-clicks>

- [ ] Necesitas control total
- [ ] Tienes requisitos especiales
- [ ] Migras sistema legacy
- [ ] Requieres software específico

</v-clicks>

---
layout: center
class: text-center
---

# Historias de Terror en la Nube
## Cuando el costo se sale de control

<div class="text-xl mt-8">
<v-clicks>

Antes de continuar, veamos algunos casos reales...

...para que no te pase a ti

</v-clicks>
</div>

---

# El Desastre de los $72,000 de Milkie Way

<v-clicks>

### La historia
**Marzo 2020**: Startup Milkie Way casi quiebra por un error de configuración

### ¿Qué pasó?
- Desarrollaban un scraper web con Firebase + Cloud Run
- Crearon recursión infinita accidental (páginas que se enlazaban entre sí)
- 1000 instancias consultando Firebase cada milisegundo

### Los números del horror
- **116 mil millones** de lecturas a Firestore en menos de 1 hora
- **1 mil millones** de requests por minuto en el pico
- Factura final: **$72,000 USD**
- Presupuesto configurado: $7 USD

</v-clicks>

---

# El Desastre de Milkie Way: Lecciones

<v-clicks>

### Problemas descubiertos

**No hay límites reales de gasto**
- Los "budgets" solo envían alertas, no detienen el gasto
- Firebase se auto-actualizó del plan gratuito sin avisar

**Facturación con retraso**
- GCP tarda 24+ horas en mostrar costos reales
- Dashboard de Firebase mostraba 42,000 lecturas cuando eran 116 mil millones
- Diferencia: **86,585,365%** de error

### El final
- El fundador ya estudiaba capítulos de bancarrota
- Google perdonó la deuda como "gesto único"
- Conclusión del fundador: *"Fail fast con cloud es una mala idea"*

</v-clicks>

---

# Más Historias de Terror Cloud

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

### El Bug Recursivo de AWS Lambda

<v-clicks>

- Función que se triggerea a sí misma
- Upload a S3 → Lambda → modifica archivo → nuevo upload
- **Costo overnight**: $4,000-15,000
- Sin exceder concurrencia de 1

</v-clicks>

</div>

<div>

### Ataques de Crypto-Mining

<v-clicks>

- Credenciales AWS expuestas en GitHub
- Hackers crean instancias GPU para minar
- **Casos documentados**: $2,000 - $2.3M
- Tiempo de detección: días o semanas

</v-clicks>

</div>

</div>

<v-click>

### Pinterest: Sorpresa Navideña
- Tráfico navideño excedió estimaciones
- Ya habían pagado $170M por adelantado
- **Sobrecosto adicional**: $20 millones

</v-click>

---

# Por Qué Suceden Estos Desastres

<v-clicks>

### La velocidad del desastre
- Los costos pueden acumularse en **minutos**
- Una función recursiva es "la inundación flash de los desastres cloud"
- Para cuando te enteras, ya es demasiado tarde

### La falta de límites duros
- AWS, GCP, Azure: ninguno tiene límites de gasto reales
- Solo "alertas" que llegan cuando ya gastaste
- Diseñado para "no limitar tu crecimiento"

### El miedo a experimentar
> "Normalmente aprendo rompiendo cosas. Con AWS no me siento cómodo haciendo eso"
> — Desarrollador anónimo

### Resultado
Startups han muerto porque necesitaban salir de AWS pero no podían hacerlo a tiempo

</v-clicks>

---

# Caso de Éxito: 37signals

<v-clicks>

- **Empresa**: 37signals (Basecamp, HEY)
- **Factura cloud 2022**: $3.2 millones/año
- **Decisión**: Salir completamente de AWS

</v-clicks>

---

# La Migración (2023)

<v-clicks>

- 7 aplicaciones de AWS a hardware propio
- Sin contratar personal adicional
- Inversión en hardware Dell: **$700,000**

</v-clicks>

---

# Resultados Año 1

<v-clicks>

- **Ahorro 2024**: $2 millones
- **Nueva factura**: $1.3 millones (solo S3)
- Hardware pagado en el primer año

</v-clicks>

---

# Proyección a 5 Años

<v-clicks>

- **Ahorro total**: Más de $10 millones
- **Reducción de costos**: 60-66%

<span v-mark.highlight.green>ROI excepcional</span>

</v-clicks>

---

# Próximo Paso: Salir de S3

<v-clicks>

- Migrar a Pure Storage on-premise
- Capacidad: 18 petabytes
- Costo: igual a 1 año de S3
- **Ahorro anual después**: $1.3M

</v-clicks>

---

# La Filosofía de DHH

<v-click>

> "Rentar computadoras es (mayormente) un mal negocio para empresas medianas con crecimiento estable"
> — DHH, CTO de 37signals

</v-click>

---

# ¿Para Quién NO Aplica?

<v-clicks>

- Startups en etapa inicial
- Empresas con carga muy irregular
- Proyectos experimentales

</v-clicks>

---

# ¿Para Quién SÍ Aplica?

<v-clicks>

- Empresas medianas establecidas
- Crecimiento predecible
- Carga estable 24/7

<span v-mark.underline.yellow>37signals lo demostró</span>

</v-clicks>

---

# No Hagas Esto

<v-clicks>

- Usar K8s para 1 servicio
- Optimizar costos prematuramente
- Ignorar vendor lock-in
- No planear la migración
- Subestimar la complejidad

</v-clicks>

---

# Mejor Haz Esto

<v-clicks>

- Empieza simple, evoluciona
- Mide antes de optimizar
- Usa abstracciones estándar
- Ten estrategia de salida
- Considera el TCO completo

</v-clicks>

---

# La Regla de Oro

<div class="text-center mt-16">

<span v-mark.underline.red>"La complejidad prematura es la raíz de todos los males"</span>

</div>

---

# Mejores Prácticas para Deployment Seguro

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

### Prevención de desastres de costos

<v-clicks>

- Configurar alertas de presupuesto en múltiples niveles
- Implementar circuit breakers en funciones
- Usar rate limiting agresivo
- Auditar permisos IAM regularmente
- Nunca commitear credenciales

</v-clicks>

</div>

<div>

### Monitoreo proactivo

<v-clicks>

- Dashboards de costo en tiempo real
- Alertas por anomalías de uso
- Revisión diaria de gastos
- Tags para tracking de costos
- Autoscaling con límites máximos

</v-clicks>

</div>

</div>

<v-click>

### La regla de oro

> "Si no entiendes completamente cómo se cobra un servicio, no lo uses en producción"

</v-click>

---

# Optimización: Ganancias Inmediatas

<v-clicks>

- Apagar recursos de desarrollo nocturnos
- Usar instancias spot
- Comprimir datos en S3
- Eliminar snapshots antiguos

</v-clicks>

---

# Optimización: Mediano Plazo

<v-clicks>

- Reserved instances para carga base
- Consolidar cuentas para descuentos
- Migrar a regiones baratas
- Auto-scaling eficiente

</v-clicks>

---

# Optimización: Largo Plazo

<v-clicks>

- Evaluar repatriación (como 37signals)
- Negociar contratos enterprise
- Arquitectura serverless
- Considerar hybrid cloud

</v-clicks>

---

# Platform Engineering: La Tendencia 2024-2025

<v-clicks>

### ¿Qué es Platform Engineering?

Crear plataformas internas de autoservicio que aceleren el desarrollo

### Beneficios documentados
- **60%** mejor utilización de recursos
- **45%** reducción en overhead operacional
- **55%** mayor eficiencia en deployments
- **35%** reducción en costos operativos

### Componentes clave
- Portal de autoservicio para developers
- Templates de infraestructura pre-aprobados
- CI/CD pipelines estandarizados
- Observabilidad integrada

### El resultado
Los desarrolladores se enfocan en código, no en infraestructura

</v-clicks>

---

# FinOps: Gestión Financiera del Cloud

<v-clicks>

### Los 3 pilares de FinOps

**Informar**: Visibilidad total de costos
- Quién gasta qué y por qué
- Dashboards por equipo/proyecto

**Optimizar**: Reducir desperdicios
- Rightsizing de recursos
- Eliminación de recursos idle
- Uso de descuentos y spots

**Operar**: Governance continuo
- Políticas de gasto
- Aprobaciones para recursos costosos
- Revisiones mensuales de optimización

### Resultado típico
35% de reducción en gasto cloud en el primer año

</v-clicks>

---

# Recursos y Herramientas

<div class="grid grid-cols-3 gap-6 mt-8">

<div>

### Documentación
- [Go Cloud Development Kit](https://gocloud.dev)
- [12 Factor App](https://12factor.net)
- [Cloud Native Go](https://www.oreilly.com/library/view/cloud-native-go/9781492076322/)

</div>

<div>

### Herramientas
- [ko](https://ko.build/) - Build containers
- [Skaffold](https://skaffold.dev/) - K8s workflow
- [Terraform](https://terraform.io) - IaC

</div>

<div>

### Calculadoras
- [AWS Calculator](https://calculator.aws)
- [GCP Calculator](https://cloud.google.com/products/calculator)
- [Azure Calculator](https://azure.microsoft.com/pricing/calculator/)

</div>

</div>

<v-click>

<div class="mt-8 p-4 bg-green-500 bg-opacity-10 rounded">
**Pro tip**: Usa las calculadoras ANTES de elegir plataforma
</div>

</v-click>

---
layout: center
class: text-center
transition: fade
---

# Conclusiones

<div class="text-2xl mt-8 space-y-4">

<v-clicks>

<div>
No existe una solución única
</div>

<div>
Empieza simple, escala cuando lo necesites
</div>

<div>
Considera el costo total (tiempo + dinero)
</div>

<div>
Go funciona bien en todas las plataformas
</div>

</v-clicks>

</div>

<v-click>

<div class="mt-12 p-6 bg-gradient-to-r from-blue-500 to-green-500 bg-opacity-10 rounded-lg">
<span v-mark.circle.orange class="text-xl">"La mejor arquitectura es la que resuelve tu problema actual, no el que podrías tener"</span>
</div>

</v-click>

---
layout: center
class: text-center
transition: slide-up
---

# ¡Gracias!

## ¿Preguntas?

<div class="mt-12">
  <a href="https://github.com" target="_blank" class="mx-2">
    <carbon:logo-github class="text-3xl"/>
  </a>
  <a href="https://twitter.com" target="_blank" class="mx-2">
    <carbon:logo-twitter class="text-3xl"/>
  </a>
  <a href="https://linkedin.com" target="_blank" class="mx-2">
    <carbon:logo-linkedin class="text-3xl"/>
  </a>
</div>

<div class="mt-8 text-sm opacity-75">
Presentación creada con Slidev
</div>

<style>
h1 {
  background-color: #00ADD8;
  background-image: linear-gradient(45deg, #00ADD8 25%, #5AC8E2 50%);
  background-size: 100%;
  -webkit-background-clip: text;
  -moz-background-clip: text;
  -webkit-text-fill-color: transparent;
  -moz-text-fill-color: transparent;
}
</style>