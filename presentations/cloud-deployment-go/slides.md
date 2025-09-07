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
class: text-center
---

# Framework de Decisión
## Preguntas clave para elegir tu plataforma

<v-click>

Responde estas 9 preguntas en orden para encontrar la mejor opción

</v-click>

---

# Pregunta 1: ¿Cuál es tu presupuesto?

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### 💸 Presupuesto mínimo
- **Tráfico bajo o variable**
- **Comenzando el proyecto**

→ **Serverless** (Cloud Run, Lambda)
→ Pay-per-use, free tiers generosos

</div>

<div v-click>

### 💰 Presupuesto establecido
- **Tráfico constante**
- **Proyecto en producción**

→ **PaaS** (Heroku, Render) o **VMs reservadas**
→ Costos predecibles mensuales

</div>

</div>

<v-click>

<div class="mt-8 p-4 bg-blue-500 bg-opacity-10 rounded">
💡 **Tip**: Con tráfico < 100 req/min, serverless casi siempre es más barato
</div>

</v-click>

---

# Pregunta 2: ¿Qué tan variable es tu tráfico?

<v-clicks>

### 📊 Patrones de tráfico

**Muy variable** (0 → 10,000 requests/segundo)
- Picos impredecibles
- Eventos virales posibles
→ **Serverless** escala automáticamente

**Predecible** (100-1000 requests/segundo constante)
- Carga estable 24/7
- Crecimiento gradual
→ **VMs** o **Containers** con auto-scaling

**Intermitente** (pocas horas al día)
- Jobs programados
- APIs internas
→ **Serverless** o **PaaS** con scale-to-zero

</v-clicks>

---

# Pregunta 3: ¿Cuánto control necesitas?

<div class="text-2xl mt-10">

<v-clicks>

**Control total del OS y red**
→ Bare Metal o IaaS (VMs)

**Control del runtime y configuración**
→ Containers (Docker, Kubernetes)

**Solo control del código**
→ PaaS (Heroku, App Engine)

**Mínimo control, máxima simplicidad**
→ Serverless (Cloud Run, Lambda)

</v-clicks>

</div>

---

# Pregunta 4: ¿Qué experiencia tiene tu equipo?

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### 👨‍💻 Sin equipo DevOps

- Solo desarrolladores
- Tiempo limitado para infra

**Recomendación:**
- PaaS (Heroku, Render)
- Serverless (Cloud Run)
- Evitar: Kubernetes, VMs

</div>

<div v-click>

### 🚀 Con equipo DevOps

- Experiencia en infraestructura
- Pueden mantener sistemas

**Opciones abiertas:**
- Kubernetes para microservicios
- IaaS para optimización
- Hybrid: PaaS + custom

</div>

</div>

---

# Pregunta 5: ¿Dónde están tus usuarios?

<v-clicks>

### 🌍 Distribución geográfica

**Local/Regional** (un país o región)
- Latencia < 50ms aceptable
→ **Un datacenter regional** en cualquier plataforma

**Nacional** (todo un país grande)
- Necesitas baja latencia nacional
→ **Multi-zona** en una región (AWS/GCP/Azure)

**Global** (usuarios mundiales)
- Latencia < 20ms requerida
→ **Edge deployment** (Fly.io, Cloudflare)
→ **Multi-región** con CDN

</v-clicks>

---

# Pregunta 6: ¿Qué tipo de aplicación es?

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### API REST / GraphQL
- Stateless
- Request-response

**Ideal**: Cloud Run, Lambda, PaaS

</div>

<div v-click>

### WebSockets / Real-time
- Conexiones persistentes
- Estado en memoria

**Ideal**: VMs, Containers, Fly.io

</div>

</div>

<div class="grid grid-cols-2 gap-8 mt-4">

<div v-click>

### Background Jobs
- Procesamiento asíncrono
- Tareas programadas

**Ideal**: Cloud Functions, Lambda

</div>

<div v-click>

### Monolito tradicional
- Una sola aplicación
- Base de datos integrada

**Ideal**: PaaS, VM única

</div>

</div>

---

# Pregunta 7: ¿Necesitas cumplir regulaciones?

<v-clicks>

### 🔒 Requerimientos de compliance

**HIPAA / PCI-DSS / SOC2**
- Aislamiento estricto
- Auditoría completa
→ **VMs dedicadas** o **Bare Metal**
→ Proveedores certificados (AWS, Azure, GCP)

**GDPR / Data residency**
- Datos en región específica
- Control de localización
→ **Deployment regional** específico
→ Evitar edge/CDN global automático

**Sin regulaciones especiales**
→ Cualquier opción es válida

</v-clicks>

---

# Pregunta 8: ¿Cuál es tu velocidad de desarrollo?

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### 🏃 Desarrollo rápido
**MVP, Startup, Hackathon**

Necesitas:
- Deploy en minutos
- Iteración rápida
- Zero config

→ **PaaS** (Railway, Render)
→ **Serverless** (Vercel, Netlify)

</div>

<div v-click>

### 🐢 Desarrollo estable
**Enterprise, Sistema crítico**

Puedes permitirte:
- Setup inicial largo
- Optimización profunda
- Configuración detallada

→ **Kubernetes**
→ **IaaS** con IaC

</div>

</div>

---

# Pregunta 9: ¿Qué esperas a largo plazo?

<div class="grid grid-cols-2 gap-8 mt-8">

<div v-click>

### 📈 Alto crecimiento esperado
**10x-100x en 12 meses**

- Tráfico exponencial
- De 100 a 1M usuarios
- Scaling frecuente

**Costo óptimo:**
- Comenzar con **Serverless**
- Migrar a **VMs reservadas** al crecer
- **Kubernetes** para > 100k usuarios

</div>

<div v-click>

### 📊 Crecimiento estable
**2x-3x anual**

- Predicción confiable
- Base de usuarios establecida
- Tráfico 24/7 constante

**Costo óptimo:**
- **VMs con reservas** (1-3 años)
- **PaaS con planes anuales**
- Evitar serverless para alto volumen

</div>

</div>

<v-click>

<div class="mt-8 p-4 bg-red-500 bg-opacity-10 rounded">
⚠️ **Punto de inflexión**: ~50k requests/día = serverless puede ser más caro que VM dedicada
</div>

</v-click>

---


# Matriz de Decisión Rápida

<div class="text-sm">

| Criterio | Serverless | PaaS | Containers | VMs | Bare Metal |
|----------|------------|------|------------|-----|------------|
| **Costo inicial** | 💚 Mínimo | 🟡 Bajo | 🟡 Medio | 🟠 Alto | 🔴 Muy alto |
| **Costo alto volumen** | 🔴 Muy alto | 🟠 Alto | 🟡 Medio | 💚 Bajo | 💚 Mínimo |
| **Escalabilidad** | 💚 Automática | 🟡 Semi-auto | 🟡 Configurable | 🟠 Manual | 🔴 Física |
| **Time to market** | 💚 Minutos | 💚 Horas | 🟡 Días | 🟠 Semanas | 🔴 Meses |
| **Control** | 🔴 Mínimo | 🟠 Limitado | 🟡 Bueno | 💚 Total | 💚 Absoluto |
| **Complejidad ops** | 💚 Ninguna | 💚 Baja | 🟠 Alta | 🔴 Muy alta | 🔴 Extrema |
| **Vendor lock-in** | 🔴 Alto | 🟠 Medio | 🟡 Bajo | 💚 Ninguno | 💚 Ninguno |

</div>

<v-click>

<div class="mt-4 p-3 bg-green-500 bg-opacity-10 rounded">
💚 = Excelente | 🟡 = Bueno | 🟠 = Regular | 🔴 = Pobre
</div>

</v-click>

---

# Recomendaciones por Escenario

<div class="grid grid-cols-2 gap-6 text-sm">

<div v-click>

### 🚀 Startup/MVP
**Objetivo**: Velocidad y bajo costo

1. Comenzar con **Serverless** (Cloud Run)
2. Si no encaja, usar **PaaS** (Render/Railway)
3. Migrar a Containers cuando crezcas

</div>

<div v-click>

### 🏢 Empresa establecida
**Objetivo**: Confiabilidad y control

1. **Kubernetes** para microservicios
2. **VMs** para apps legacy
3. **Serverless** para nuevos features

</div>

<div v-click>

### 🎓 Proyecto personal
**Objetivo**: Aprender y experimentar

1. **Free tiers** (Cloud Run, Railway)
2. **DigitalOcean** $5/mes
3. **Oracle Cloud** VMs gratis

</div>

<div v-click>

### 🌍 SaaS Global
**Objetivo**: Performance mundial

1. **Fly.io** para edge deployment
2. **Multi-region** en cloud mayor
3. **CDN** + API en múltiples zonas

</div>

</div>

---

# Checklist Final de Decisión

<v-clicks>

✅ **Antes de elegir, verifica:**

1. ¿El presupuesto permite esta opción a largo plazo?
2. ¿Tu equipo puede mantener esta solución?
3. ¿La plataforma soporta tus requisitos técnicos?
4. ¿Hay vendor lock-in aceptable?
5. ¿La latencia será adecuada para tus usuarios?
6. ¿Puedes migrar si necesitas cambiar?
7. ¿Los costos escalan linealmente con el crecimiento?

</v-clicks>

<v-click>

<div class="mt-8 p-4 bg-yellow-500 bg-opacity-10 rounded">
⚠️ **Regla de oro**: Empieza simple, migra cuando sea necesario
</div>

</v-click>
