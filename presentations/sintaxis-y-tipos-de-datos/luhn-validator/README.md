# 🎪 Validador de Tarjetas de Crédito - Algoritmo de Luhn

Este proyecto implementa el **algoritmo de Luhn** en Go para validar números de tarjetas de crédito.

## 🎯 ¿Qué es el Algoritmo de Luhn?

El algoritmo de Luhn es una fórmula de suma de verificación simple utilizada para validar números de tarjetas de crédito, códigos IMEI, números de identificación nacional de muchos países, y otros números de identificación.

## 🚀 Uso

### Ejecutar directamente:
```bash
go run main.go
```

### Compilar y ejecutar:
```bash
go build
./luhn-validator
```

## 📋 Tipos de Tarjetas Soportadas

El programa reconoce y valida:

- **Visa**: Números de 16 dígitos que comienzan con 4
- **Mastercard**: Números de 16 dígitos que comienzan con 51-55
- **American Express**: Números de 15 dígitos que comienzan con 34 o 37
- **Discover**: Números de 16 dígitos que comienzan con 6011

## 🔍 Casos de Prueba Incluidos

El programa incluye 8 casos de prueba:
- ✅ 4 números válidos (uno de cada tipo de tarjeta)
- ❌ 4 números inválidos (para demostrar la detección de errores)

## 🎓 Conceptos de Go Demostrados

Este proyecto muestra múltiples aspectos de la sintaxis de Go:

- **Variables y tipos**: `int`, `string`, `bool`
- **Bucles**: `for` con diferentes patrones
- **Condicionales**: `if/else`, `switch`
- **Slices**: `[]string`
- **Funciones**: Definición, parámetros y valores de retorno
- **Conversiones de tipos**: `int(cardNumber[i] - '0')`
- **String indexing**: Acceso a caracteres individuales
- **Operadores**: Aritméticos, módulo, comparación
- **Printf**: Formateo de salida

## 💡 Datos Curiosos

- El algoritmo detecta aproximadamente el **70%** de errores de transcripción
- Fue desarrollado por **Hans Peter Luhn** en IBM en 1954
- Es utilizado por prácticamente todas las compañías de tarjetas de crédito

## 🔧 Comandos Go Utilizados

```bash
go mod init luhn-validator  # Inicializar módulo
go run main.go             # Ejecutar directamente
go build                   # Compilar binario
go fmt                     # Formatear código
```

---

**Parte del curso**: *Sintaxis, Tipos de Datos y Sistema de Tipos en Go*