---
layout: center
---
# Sintaxis, Tipos de Datos y Sistema de Tipos en Go  
### Curso introductorio
<!-- 
Notas del presentador: Dar la bienvenida a la clase y presentar el tema. Mencionar que veremos la sintaxis básica del lenguaje Go, sus tipos de datos principales, cómo funciona su sistema de tipos (incluyendo los nuevos genéricos), la interoperabilidad con C y WebAssembly, y finalmente algunos recursos para seguir aprendiendo. Explicar que la sesión está pensada para principiantes e intermedios, durará ~2 horas, e incluirá ejemplos de código y ejercicios prácticos.
-->

---

## <span class="underline decoration-4 decoration-emerald-500">Temario</span> de la Presentación  

<v-clicks>

- <span class="border-2 border-sky-500 rounded px-2 py-0.5">**Sintaxis básica**</span>: Estructura de un programa Go, variables, constantes, control de flujo  
- <span class="border-2 border-purple-500 rounded px-2 py-0.5">**Tipos de datos**</span>: Primitivos y compuestos  
- <span class="border-2 border-rose-500 rounded px-2 py-0.5">**Sistema de tipos**</span>: Interfaces, polimorfismo y <span class="underline decoration-4 decoration-rose-500">genéricos</span>  
- <span class="border-2 border-amber-500 rounded px-2 py-0.5">**Interoperabilidad**</span>: CGo y WebAssembly  
- <span class="border-2 border-green-500 rounded px-2 py-0.5">**Recursos adicionales**</span>: Para continuar aprendiendo  
- <span class="bg-lime-200 px-2 rounded">**Ejercicios prácticos**</span> en cada sección

</v-clicks>
<!--
Notas del presentador: Dar un repaso de los puntos que cubriremos. Mencionar que empezaremos por la sintaxis básica, luego repasaremos los tipos de datos disponibles en Go, hablaremos sobre el sistema de tipos (incluyendo interfaces y genéricos), veremos cómo Go se puede integrar con otros entornos (C y WebAssembly), y finalmente compartiremos recursos. Destacar que después de cada sección habrá un ejercicio para practicar.
-->

---

## Sintaxis Básica

### <span v-click class="underline decoration-4 decoration-blue-500">¿Qué hace especial a Go?</span>

<v-clicks>

- **Sintaxis <span class="underline decoration-2 decoration-green-500">simple</span>**: Fácil de leer y escribir
- **Reglas <span class="underline decoration-2 decoration-purple-500">claras</span>**: Sin ambigüedades
- **Compilación <span class="underline decoration-2 decoration-orange-500">rápida</span>**: Feedback inmediato
- **Menos es <span class="underline decoration-2 decoration-cyan-500">más</span>**: No hay sintaxis innecesaria

</v-clicks>

<v-click>

💡 **Filosofía**: <span class="bg-yellow-200 px-2 py-1 rounded">Claridad sobre cleverness</span>

</v-click>

---

## <span class="border-2 border-indigo-500 rounded-lg px-3 py-1">Lo que Aprenderemos</span>

<span v-click class="bg-blue-200 px-2 py-1 rounded">**Elementos fundamentales**</span> de todo programa Go

<v-clicks>

- **Estructura del programa**: `package`, `import`, `func main`
- **Variables y constantes**: Declaración y tipos
- **Control de flujo**: `for`, `if/else`, `switch`
- **Particularidades**: Sin paréntesis, llaves obligatorias

</v-clicks>

<!-- 
Notas del presentador: Introducir la sintaxis básica de Go. Explicar que primero veremos cómo se estructura un programa Go: paquete, imports, función main, etc. Luego hablaremos de variables, constantes y estructuras de control (if, for, switch). Destacar que Go tiene una sintaxis concisa y algunas reglas particulares (por ejemplo, posicionamiento de llaves, ausencia de paréntesis en if/for, etc.).
-->

---

## <span class="border-2 border-cyan-500 rounded-lg px-3 py-1">Comandos Esenciales de Go</span>

<span v-click class="bg-green-200 px-2 py-1 rounded">**CLI de Go**</span> - Tu caja de herramientas

<v-clicks>

- **`go run`**: <span class="underline decoration-2 decoration-blue-500">Ejecuta</span> directamente archivos `.go`
- **`go build`**: <span class="underline decoration-2 decoration-green-500">Compila</span> el programa a binario
- **`go mod init`**: <span class="underline decoration-2 decoration-purple-500">Inicializa</span> un nuevo módulo
- **`go mod tidy`**: <span class="underline decoration-2 decoration-orange-500">Limpia</span> dependencias
- **`go fmt`**: <span class="underline decoration-2 decoration-red-500">Formatea</span> código automáticamente
- **`go test`**: <span class="underline decoration-2 decoration-cyan-500">Ejecuta</span> pruebas unitarias

</v-clicks>

<v-click>

💡 **Tip**: `go help <comando>` para más información

</v-click>

---

## <span class="border-2 border-emerald-500 rounded-lg px-3 py-1">Proyecto Práctico</span>

<span v-click class="bg-purple-200 px-2 py-1 rounded">**Algoritmo de Luhn**</span> - Validación de tarjetas de crédito

<v-clicks>

- 💳 **Qué hace**: Valida números de tarjetas de crédito
- 🔢 **Algoritmo**: Suma ponderada de dígitos
- 🎯 **Sintaxis Go**: loops, condiciones, conversiones
- 🛠️ **Herramientas**: Todos los comandos CLI

</v-clicks>

<v-click>

🎪 **¡Vamos a construirlo paso a paso!**

</v-click>

---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Inicializando el Proyecto</span>

<span v-click class="underline decoration-wavy decoration-cyan-400">**Primer paso**</span> - Crear nuestro workspace

```bash {1|2|3}
go mod init luhn-validator
cd luhn-validator
touch main.go
```

<v-clicks>

- `go mod init`: Crea archivo `go.mod` 📄
- `luhn-validator`: Nombre de nuestro módulo
- `main.go`: Archivo principal del programa

</v-clicks>

<v-click>

💡 **Resultado**: Proyecto Go listo para desarrollo

</v-click>

---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Estructura del Código</span>

<span v-click class="bg-amber-200 px-2 py-1 rounded">**main.go**</span> - Esqueleto de nuestro programa

```go {1-3|5-7|9-11}
package main

import "fmt"

func main() {
    // Aquí irá nuestra lógica
}

func luhnCheck(cardNumber string) bool {
    // Algoritmo de Luhn
}
```

<v-clicks>

- `package main`: <span class="underline decoration-2 decoration-green-500">Punto de entrada</span>
- `import "fmt"`: Para <span class="underline decoration-2 decoration-blue-500">input/output</span>
- `func main()`: <span class="underline decoration-2 decoration-purple-500">Función principal</span>
- `func luhnCheck()`: <span class="underline decoration-2 decoration-orange-500">Nuestra función</span>

</v-clicks>

---

## <span class="border-2 border-rose-500 rounded-lg px-3 py-1">Implementando Luhn</span>

<span v-click class="underline decoration-4 decoration-indigo-500">**El algoritmo completo**</span>

```go {1-2|4-6|8-10|12-15|17-19|21}
func luhnCheck(cardNumber string) bool {
    sum := 0
    
    // Recorrer dígitos de derecha a izquierda
    for i := len(cardNumber) - 1; i >= 0; i-- {
        digit := int(cardNumber[i] - '0')
        
        // Duplicar cada segundo dígito
        if (len(cardNumber)-i)%2 == 0 {
            digit *= 2
            
            // Si el resultado > 9, sumar sus dígitos
            if digit > 9 {
                digit = digit/10 + digit%10
            }
        }
        
        sum += digit
    }
    
    return sum%10 == 0
}
```

---

## <span class="border-2 border-cyan-500 rounded-lg px-3 py-1">Función Main Completa</span>

<span v-click class="bg-green-200 px-2 py-1 rounded">**Probando nuestro algoritmo**</span>

```go {1-3|5-8|10-15}
func main() {
    testCards := []string{
        "4532015112830366", // Visa válida
        "4532015112830367", // Visa inválida  
    }
    
    for _, card := range testCards {
        isValid := luhnCheck(card)
        
        if isValid {
            fmt.Printf("✅ %s es válida\n", card)
        } else {
            fmt.Printf("❌ %s es inválida\n", card)
        }
    }
}
```

<v-click>

🎯 **Sintaxis mostrada**: arrays, loops, condiciones, strings, conversiones

</v-click>

---

## Estructura de un programa Go  

### <span v-click class="underline decoration-4 decoration-sky-600">Elementos fundamentales</span>

---

## <span class="border-2 border-emerald-500 rounded-lg px-3 py-1">package main</span>

Define el <span v-click class="bg-amber-200 px-1 rounded">**paquete**</span> al que pertenece el archivo

<v-click>

- <span class="underline decoration-wavy decoration-blue-400">**Requerido**</span> para programas ejecutables
- Otros paquetes: `package utils`, `package models`, etc.

</v-click>

```go {1}
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, Go!")
}
```

---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">import "fmt"</span>

Importa <span v-click class="bg-lime-200 px-1 rounded">**paquetes**</span> necesarios

<v-click>

- `fmt`: <span class="underline decoration-2 decoration-blue-500">Formato de entrada/salida</span>
- Otros: `"os"`, `"net/http"`, `"strings"`

</v-click>

```go {3}
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, Go!")
}
```

---

## <span class="border-2 border-rose-500 rounded-lg px-3 py-1">func main()</span>

<span v-click class="underline decoration-4 decoration-rose-600">**Punto de entrada**</span> del programa

<v-click>

- Se ejecuta <span class="underline decoration-2 decoration-green-500">automáticamente</span> al correr el programa
- Solo una función `main` por paquete `main`

</v-click>

```go {5-7}
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, Go!")
}
```  
<!-- 
Notas del presentador: Comentar que todo archivo Go inicia declarando un **paquete**. El paquete `main` indica un programa ejecutable:contentReference[oaicite:1]{index=1}. Luego vienen las declaraciones `import` para incluir paquetes estándar o de terceros; por ejemplo `fmt` proporciona funciones de formato de entrada/salida. Luego definimos funciones; en particular, `func main()` es donde comienza la ejecución del programa.  
Explicar también la sintaxis de llaves: en Go, la llave `{` debe ir en la misma línea que la declaración `func` o `if/for` correspondiente, no en la siguiente línea:contentReference[oaicite:2]{index=2}. El compilador inserta puntos y coma automáticamente al final de líneas, por lo que colocar `{` en la línea siguiente causaría errores.  
En el ejemplo a la derecha, tenemos un programa mínimo que imprime un mensaje. Destacar cómo usamos `fmt.Println` para imprimir texto en la consola.
-->

---

## Variables y Constantes en Go

### <span v-click class="underline decoration-4 decoration-indigo-500">Formas de declarar variables</span>

---

## <span class="border-2 border-green-500 rounded-lg px-3 py-1">Declaración Explícita</span>

<span v-click class="bg-yellow-200 px-2 rounded">**Especificamos el tipo**</span> claramente

```go {1}
var x int = 10
```

<v-click>

- Sintaxis: `var nombre tipo = valor`
- <span class="underline decoration-2 decoration-purple-500">Tipado estático</span> y explícito
- Útil cuando el tipo no es obvio

</v-click>

---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Inferencia de Tipo</span>

Go <span v-click class="underline decoration-wavy decoration-emerald-400">**deduce**</span> el tipo automáticamente

```go {1}
var y = 20
```

<v-click>

- El compilador <span class="underline decoration-2 decoration-cyan-500">infiere</span> que `y` es `int`
- Menos verboso, igual de seguro
- Go es inteligente con los tipos

</v-click>

---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Declaración Corta</span>

La forma <span v-click class="bg-lime-200 px-1 rounded">**más popular**</span> en Go

```go {1}
z := x + y
```

<v-click>

- Solo <span class="underline decoration-2 decoration-orange-500">dentro de funciones</span>
- Sintaxis: `nombre := valor`
- Declaración + asignación + inferencia

</v-click>

---

## <span class="border-2 border-amber-500 rounded-lg px-3 py-1">Constantes</span>

Valores <span v-click class="underline decoration-4 decoration-rose-600">**inmutables**</span>

```go {1}
const Pi = 3.14
```

<v-click>

- <span class="underline decoration-2 decoration-red-500">No cambian</span> durante la ejecución
- Se evalúan en <span class="underline decoration-2 decoration-indigo-500">tiempo de compilación</span>
- Pueden ser números, strings o booleanos

</v-click>  
<!-- 
Notas del presentador: Explicar que Go es un lenguaje **tipado estáticamente**, por lo que cada variable tiene un tipo fijo:contentReference[oaicite:3]{index=3}. Podemos declarar variables con `var nombre tipo = valor`. Si omitimos el tipo, Go **infier** el tipo a partir del valor asignado (por ejemplo, `y` será int porque 20 es int literal).  
Dentro de una función (por ejemplo en `main`), podemos usar la sintaxis corta `:=` para declarar e inicializar variables en una sola expresión, con inferencia de tipo automática. Notar que esta forma no se puede usar fuera de funciones (en nivel de paquete).  
Las **constantes** se declaran con `const` y un valor fijo que no cambia. Pueden ser numéricas, cadenas, booleanos, etc. Las constantes numéricas pueden ser "no tipadas" internamente hasta que se usan, lo cual les da flexibilidad (por ejemplo, Pi se podría usar como float32 o float64 según contexto).  
Resaltar que debido al tipado fuerte de Go, no se permiten conversiones implícitas: por ejemplo, no podemos asignar un `int32` a una variable `int` sin conversión explícita.  
-->

---

## Estructuras de Control

### <span v-click class="underline decoration-4 decoration-purple-600">Go simplifica el control de flujo</span>

---

## <span class="border-2 border-emerald-500 rounded-lg px-3 py-1">for: El único bucle</span>

Go tiene <span v-click class="bg-amber-200 px-1 rounded">**solo**</span> `for`, pero es muy flexible

<v-clicks>

- Bucle tradicional: `for i := 0; i < 10; i++`
- Estilo "while": `for condición { ... }`
- Bucle infinito: `for { ... }`
- Con `range`: `for i, v := range slice`

</v-clicks>

```go {1}
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">if/else: Sin paréntesis</span>

<span v-click class="underline decoration-wavy decoration-sky-400">**Sintaxis limpia**</span>, llaves obligatorias

```go {2-6}
for i := 1; i <= 5; i++ {
    if i % 2 == 0 {
        fmt.Println(i, "es par")
    } else {
        fmt.Println(i, "es impar")
    }
}
```

<v-clicks>

- <span class="underline decoration-2 decoration-yellow-500">Sin paréntesis</span> en la condición
- Llaves `{}` siempre requeridas
- Permite declaración corta: `if err := foo(); err != nil`

</v-clicks>

---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">switch: Sin break</span>

Switch <span v-click class="bg-lime-200 px-1 rounded">**inteligente**</span> y seguro

```go
switch dia {
case "lunes":
    fmt.Println("Inicio de semana")
case "viernes":
    fmt.Println("¡Fin de semana cerca!")
default:
    fmt.Println("Día normal")
}
```

<v-clicks>

- <span class="underline decoration-2 decoration-pink-500">No necesita</span> `break` (rompe automáticamente)
- `fallthrough` para continuar al siguiente case
- Puede evaluar expresiones, no solo valores

</v-clicks>  
<!-- 
Notas del presentador: **For:** En Go, `for` reemplaza a `while` y `do-while`. Tiene tres formas principales: (1) con inicialización, condición y post (como en C, ej. `for i:=0; i<10; i++`), (2) solo con condición (`for cond { ... }` actúa como while), y (3) bucle infinito `for { ... }`. También existe `for ... range` para iterar sobre arrays, slices, mapas, canales o strings.  
En el ejemplo, usamos la forma tradicional con un índice `i`.  

**If/else:** Observar que la sintaxis de `if` en Go no lleva paréntesis alrededor de la condición. Por ejemplo, escribimos `if i % 2 == 0 { ... }` sin `()`. Las llaves `{}` son obligatorias incluso para una sola instrucción, a diferencia de lenguajes como Python o Ruby donde se usan indentaciones, o C donde se pueden omitir en caso de una sola línea (en Go *no* se pueden omitir). También es posible combinar una declaración corta con if, por ejemplo: `if err := Foo(); err != nil { ... }`, lo cual declara `err` y luego evalúa la condición.

**Switch:** Go tiene un switch muy potente. Por defecto no hace *fall-through* automático, es decir, cada `case` rompe al final automáticamente (no hay que escribir `break`). Si se desea ejecutar el siguiente caso deliberadamente, se usa la palabra `fallthrough`. Un `switch` en Go puede usarse sin expresión para hacer múltiples if/else más claros. También existe el *type switch* (`switch x.(type)`) para ramificar según el tipo de una interfaz, pero eso es más avanzado.

En el código de ejemplo, combinamos un `for` con un `if/else` dentro para imprimir si cada número es par o impar.
\-->

---

### **Ejercicio:** Sintaxis Básica

Implementa un programa en Go que recorra los números del **1 al 10** e indique para cada uno si es “par” o “impar”.

* Usa un bucle **for** para iterar del 1 al 10.
* Dentro del bucle, emplea una condición **if** para verificar si el número actual es divisible por 2.
* Imprime por pantalla mensajes como por ejemplo: “2 es par”, “3 es impar”, etc.

<!-- 
Notas del presentador: Este ejercicio refuerza el uso de `for` e `if`. Indicar a los estudiantes que pueden basarse en el ejemplo visto (aunque en el ejemplo iteramos 1 a 5, aquí es 1 a 10).  
Pautas para la solución: Inicializar un `for i := 1; i <= 10; i++`. Dentro, usar `if i % 2 == 0` para detectar números pares. Si es par, imprimir `<i> es par`, de lo contrario imprimir `<i> es impar`.  
Si ya conocen el operador `%` de módulo, estará claro; si no, aclarar que `i % 2 == 0` verifica si el residuo de dividir i por 2 es 0 (número par).  
Solución esperada (pseudocódigo):  
```go
for i := 1; i <= 10; i++ {
    if i % 2 == 0 {
        fmt.Println(i, "es par")
    } else {
        fmt.Println(i, "es impar")
    }
}
```  
-->

---

## Tipos de Datos en Go

*(Valores básicos y compuestos que maneja el lenguaje)*

<!-- 
Notas del presentador: Introducir la sección de tipos de datos. Mencionar que Go tiene una variedad de tipos básicos (numéricos, booleanos, cadenas) y tipos compuestos o estructurados (arrays, slices, mapas, structs, punteros, etc.).  
Explicar brevemente que en Go todas las variables tienen un tipo definido en compilación. Vamos a repasar los tipos disponibles y sus características principales. También comentar el concepto de **valor cero** (*zero value*): en Go, las variables no inicializadas toman un valor por defecto según su tipo (0 para números, false para bool, "" cadena vacía, nil para punteros, slices, mapas, etc.). Esto evita valores indefinidos:contentReference[oaicite:4]{index=4}.  
-->

---

### Tipos básicos

### <span v-click class="underline decoration-4 decoration-emerald-500">Los fundamentos de Go</span>

---

## <span class="border-2 border-green-500 rounded-lg px-3 py-1">Enteros</span>

Números <span v-click class="bg-amber-200 px-1 rounded">**enteros**</span> de diferentes tamaños

<v-clicks>

- `int`: Tamaño <span class="underline decoration-2 decoration-emerald-500">natural</span> de la arquitectura (32 o 64 bits)
- Específicos: `int8`, `int16`, `int32`, `int64`
- Sin signo: `uint8`, `uint16`, `uint32`, `uint64`
- Alias especiales: `byte` (uint8), `rune` (int32 para Unicode)

</v-clicks>

```go
var edad int = 25
var contador uint32 = 100
var letra rune = '€'
```

---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Flotantes</span>

Números con <span v-click class="underline decoration-wavy decoration-cyan-400">**punto decimal**</span>

<v-clicks>

- `float32`: <span class="underline decoration-2 decoration-teal-500">32 bits</span> de precisión
- `float64`: <span class="underline decoration-2 decoration-violet-500">64 bits</span> de precisión (recomendado)
- `complex64`, `complex128`: Para números <span class="underline decoration-2 decoration-amber-500">complejos</span>

</v-clicks>

```go
var precio float64 = 29.99
var pi float32 = 3.14159
var complejo complex128 = 1 + 2i
```

---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Booleanos</span>

Solo <span v-click class="bg-lime-200 px-1 rounded">**dos valores**</span> posibles

```go
var activo bool = true
var terminado bool = false
```

<v-clicks>

- <span class="underline decoration-2 decoration-lime-500">Solo</span> `true` o `false`
- No se convierten automáticamente a números
- Resultado de comparaciones y operaciones lógicas

</v-clicks>

---

## <span class="border-2 border-rose-500 rounded-lg px-3 py-1">Strings (cadenas)</span>

Texto en <span v-click class="underline decoration-4 decoration-pink-600">**UTF-8**</span>

```go
var saludo string = "¡Hola, mundo! 🌍"
var multilinea string = `Esto es una
cadena de múltiples
líneas`
```

<v-clicks>

- <span class="underline decoration-2 decoration-slate-500">Inmutables</span> (no se pueden modificar)
- Comillas dobles `"`: permiten escapes (`\n`, `\t`)
- Comillas invertidas `` ` ``: texto crudo (raw strings)

</v-clicks>

<!-- 
Notas del presentador: Repasar cada categoría:  
**Enteros:** Explicar que `int` suele ser 64 bits en sistemas de 64 bits (así lo es en la mayoría de casos modernos) y 32 bits en sistemas de 32 bits. Mencionar que además están las versiones con tamaño explícito y las sin signo. `byte` es sinónimo de uint8 (útil para datos binarios), `rune` es sinónimo de int32 y representa un carácter Unicode (un código de punto).  

**Flotantes:** Indicar que para la mayoría de los cálculos `float64` es recomendado por su precisión. Los complejos quizás no se usan tanto, pero es bueno saber que existen (puede mencionarse que `complex64` tiene parte real e imaginaria float32 cada una, etc., pero no profundizar mucho si no es necesario).  

**Booleanos:** Nada muy especial salvo que solo admiten true/false y no se pueden convertir a números (no hay equivalentes 0/1 implícitos).  

**Strings:** Explicar que Go usa UTF-8 en sus strings, por lo que maneja bien caracteres Unicode. Destacar la diferencia entre strings con comillas dobles (donde `\n` representa salto de línea, etc.) y raw strings con comillas invertidas, donde el texto se toma tal cual, permitiendo incluir nuevas líneas sin escapes. Los strings en Go son inmutables, lo que significa que operaciones que "modifican" cadenas realmente devuelven nuevas cadenas; no se puede, por ejemplo, hacer `s[0] = 'H'` si s es string (habría que convertir a slice de bytes, modificar, y reconstruir).
-->

---

### Tipos compuestos

### <span v-click class="underline decoration-4 decoration-sky-600">Estructuras de datos más complejas</span>

---

## <span class="border-2 border-green-500 rounded-lg px-3 py-1">Array</span>

Colección de <span v-click class="bg-amber-200 px-1 rounded">**longitud fija**</span>

```go
var numeros [5]int = [5]int{1, 2, 3, 4, 5}
var nombres [3]string = [3]string{"Ana", "Luis", "María"}
```

<v-clicks>

- Tamaño <span class="underline decoration-2 decoration-gray-500">fijo</span> definido en compilación
- El tamaño es parte del tipo: `[5]int` ≠ `[6]int`
- Acceso por índice: `numeros[0]`

</v-clicks>
---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Slice</span>

Arrays <span v-click class="underline decoration-wavy decoration-cyan-400">**dinámicos**</span> y flexibles

```go
var edades []int = []int{25, 30, 35}
edades = append(edades, 40)  // Agregar elemento
```

<v-clicks>

- <span class="underline decoration-2 decoration-sky-500">Longitud variable</span> durante la ejecución
- Vista a un array subyacente
- Tienen longitud (`len`) y capacidad (`cap`)
- <span class="underline decoration-2 decoration-emerald-600">Más usados</span> que los arrays

</v-clicks>
---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Map</span>

Diccionarios <span v-click class="bg-lime-200 px-1 rounded">**clave-valor**</span>

```go
var edades map[string]int = map[string]int{
    "Ana":   25,
    "Luis":  30,
    "María": 35,
}
```

<v-clicks>

- Sintaxis: `map[TipoClave]TipoValor`
- Las claves deben ser <span class="underline decoration-2 decoration-purple-600">comparables</span>
- Acceso: `edades["Ana"]`

</v-clicks>
---

## <span class="border-2 border-amber-500 rounded-lg px-3 py-1">Struct</span>

<span v-click class="underline decoration-4 decoration-orange-600">**Agrupación**</span> de campos relacionados

```go
type Persona struct {
    Nombre string
    Edad   int
    Email  string
}
```

<v-clicks>

- <span class="underline decoration-2 decoration-orange-600">Campos nombrados</span> con tipos específicos
- Go no tiene clases, usa structs
- Se pueden asociar métodos

</v-clicks>

* **Struct:** Estructura de campos nombrados, cada uno con su propio tipo. Permite crear tipos compuestos definidos por el usuario. Equivalente a “registro” o “objeto” simple (Go no tiene clases, pero los structs pueden tener métodos asociados).


<!-- 
Notas del presentador: Explicar cada uno:  
**Array:** Ejemplificar que una vez definido un array de cierto tamaño, ese tamaño no cambia. Acceso por índice (0 basado). Pueden comentar que por valor, asignar o pasar un array a una función copia todo el contenido (lo que puede ser costoso para arrays grandes).  

**Slice:** Comentar que un slice es una estructura interna con un puntero a un array, una longitud y una capacidad. Cuando se hace append, si excede capacidad puede realocar un array más grande. Mencionar que muchas funciones del paquete estándar usan slices (p.ej. bytes.Buffer).  

**Map:** Notar que es similar al concepto de diccionario o mapa hash en otros lenguajes. Las operaciones básicas: asignación `m[key] = valor`, lectura `m[key]`, eliminación con `delete(m, key)`. Si se accede a una clave inexistente, devuelve el cero valor del tipo de valor. Podemos usar sintaxis especial para checar existencia (`val, ok := m[key]`).  

**Struct:** Ejemplificar que un struct es como un "objeto" sin métodos en sí mismo (aunque se le pueden asociar métodos luego). Sirve para agrupar datos. Notar que a diferencia de lenguajes con clases, no hay herencia, pero se puede usar composición (un struct puede incluir otro struct como campo, incluso anónimo, para reutilizar código).  

**Puntero:** Destacar que en Go se usan punteros para referenciar estructuras o datos grandes en funciones (evitando copia) o para estructuras compartidas. Sin embargo, Go maneja la memoria con recolector de basura, así que no hay `free` manual. No hay pointer arithmetic: los punteros se utilizan principalmente para pasar referencias y para indicar ausencia (nil).  
-->

---

### Ejemplo: Array vs Slice

```go {1-2|4-6}
var arr [3]string = [3]string{"go", "es", "genial"}
fmt.Println(arr[0])   // "go"

s := []string{"bien", "venido"}
s = append(s, "a Go")
fmt.Println(s, len(s))  // [bien venido a Go] 3
```

<!-- 
Notas del presentador: Aquí comparamos un **array** y un **slice** en código.  
En la primera línea definimos un array `arr` de longitud 3 con valores iniciales `"go", "es", "genial"`. Al imprimir `arr[0]` obtenemos `"go"` (elemento inicial). Este array siempre tendrá tamaño 3.  

Luego definimos `s` como un slice de string, inicialmente con dos elementos `"bien", "venido"`. Los slices se pueden definir con sintaxis literal similar a arrays pero sin especificar tamaño (`[]string{...}`). Aplicamos `append(s, "a Go")` para agregar otro elemento al slice. `append` retorna el nuevo slice (puede o no ser el mismo array debajo según capacidad). Guardamos el resultado en `s`. Ahora `s` contiene tres elementos `["bien", "venido", "a Go"]`.  
Imprimimos `s` y también `len(s)` que es la longitud. Vemos la salida `[bien venido a Go] 3`.  

Notar que si el array subyacente se reubica al crecer, Go se encarga de eso internamente. Los programadores no suelen usar arrays directamente salvo en casos especiales; normalmente se usan slices para listas de tamaño variable.  
-->

---

### Ejemplo: Map

```go {1-2|3-5}
m := map[string]int{"Alice": 23, "Bob": 35}
fmt.Println(m["Alice"])  // 23

m["Charlie"] = 29
for name, age := range m {
    fmt.Println(name, age)
}
```

<!-- 
Notas del presentador: Ejemplo de uso de un **map**.  
En la primera línea creamos un mapa `m` que asocia `string` a `int`, inicializándolo con dos pares: `"Alice":23`, `"Bob":35`.  
Luego imprimimos `m["Alice"]`, lo que debería mostrar `23` (la edad de Alice).  

Después añadimos un nuevo par con `m["Charlie"] = 29`. No se necesita ninguna función especial, simplemente asignamos a la clave "Charlie". Si la clave no existía, se crea; si existía, se actualiza su valor.  

Finalmente recorremos el map con un `for range`. La sintaxis `for name, age := range m` itera sobre cada par clave-valor del mapa, asignando la clave a `name` y valor a `age`. Imprimimos ambos. El orden de iteración en los mapas **no está garantizado**, es pseudo-aleatorio.  

Podemos comentar que si necesitáramos un orden específico, tendríamos que extraer las claves y ordenarlas por separado.  
-->

---

### Ejemplo: Struct

```go {1-4|5-8}
type Persona struct {
    Nombre string
    Edad   int
}
p := Persona{"Ana", 30}
fmt.Println(p.Nombre)  // "Ana"
p.Edad = 31
fmt.Println(p)         // {Ana 31}
```

<!-- 
Notas del presentador: Ejemplo de definición y uso de un **struct**.  
Definimos un nuevo tipo `Persona` con la palabra clave `type ... struct`. Este struct tiene dos campos: `Nombre` (string) y `Edad` (int). Notar la sintaxis: llaves conteniendo campos con nombre y tipo.  

Luego, para crear una instancia, usamos sintaxis literal: `Persona{"Ana", 30}`. Esto asigna "Ana" a `Nombre` y 30 a `Edad` (en orden de campos definidos). También podríamos usar sintaxis con nombres: `Persona{Nombre: "Ana", Edad: 30}`.  

Imprimimos `p.Nombre`, obteniendo "Ana". Luego modificamos `p.Edad` asignándole 31 (los structs son mutables, incluso si `p` es una variable no puntero; se puede modificar sus campos directamente).  

Finalmente imprimimos `p` entero, lo que muestra `{Ana 31}`. Por defecto, `fmt.Println` de un struct muestra los campos en orden entre llaves.  

Se puede comentar que en Go la visibilidad de campos viene dada por la capitalización: aquí `Nombre` y `Edad` empiezan con mayúscula, así que estos campos serían **exportados** (accesibles desde otros paquetes). Si los nombráramos en minúscula, serían campos no exportados (privados al paquete). Esto es relevante si después uno crea paquetes y exporta tipos.  
-->

---

### Ejemplo: Puntero

```go {1-2|3-5}
x := 42
var px *int = &x

fmt.Println(*px)  // 42
*px = 21
fmt.Println(x)    // 21
```

<!-- 
Notas del presentador: Ejemplo de uso de **punteros** en Go.  
Primero tenemos una variable entera `x` con valor 42. Luego declaramos `px` como `*int` (puntero a int) y le asignamos la dirección de `x` usando `&x`. Ahora `px` “apunta” a `x`.  

Cuando hacemos `fmt.Println(*px)`, estamos des-referenciando el puntero (obteniendo el valor al que apunta). Esto imprime `42`, el valor actual de `x`.  

Luego hacemos `*px = 21`. Esto significa: “asignar 21 al valor apuntado por px”. En otras palabras, estamos modificando `x` indirectamente a través del puntero.  

Después imprimimos `x`, que ahora vale `21`. Esto demuestra que cambiar `*px` afectó a `x` porque `px` apuntaba a `x`.  

Apuntar que en Go, a diferencia de C/C++, no podemos hacer aritmética de punteros. No existe `px++` para moverse a otro entero en memoria, por ejemplo. Los punteros se usan de forma segura solo para referencia. También, mencionar que Go tiene un recolector de basura: no necesitamos liberar la memoria manualmente; cuando `x` ya no se use, el GC la limpiará.  
-->

---

### **Ejercicio:** Tipos de Datos

Crea un slice de **5 números enteros** a tu elección y calcula la **suma** de todos sus elementos.

* Declara un slice (por ejemplo `[]int`) con 5 valores iniciales.
* Recorre el slice con un bucle (for o range) acumulando la suma en una variable.
* Imprime el resultado de la suma.

*Opcional:* Calcula también el **promedio** (media) de esos números.

<!-- 
Notas del presentador: Este ejercicio refuerza el manejo de slices, bucles y operaciones aritméticas sencillas.  
Pautas para la solución:  
1. Crear un slice de 5 enteros, p. ej.: `nums := []int{4, 8, 15, 16, 23}` (los números pueden ser cualquiera).  
2. Inicializar una variable suma en 0.  
3. Usar un `for` (posiblemente `for _, valor := range nums`) para iterar sobre cada elemento y sumarlo: `suma += valor`.  
4. Tras el bucle, imprimir la suma.  

Si se calcula el promedio: dividir la suma por la cantidad (5). Notar que para tener decimal podría convertirse a float, pero si no se menciona, pueden dejarlo como int (división entera).  

Solución esperada (ejemplo):  
```go
nums := []int{4, 8, 15, 16, 23}
suma := 0
for _, v := range nums {
    suma += v
}
fmt.Println("Suma:", suma)
// Opcional promedio
promedio := float64(suma) / float64(len(nums))
fmt.Println("Promedio:", promedio)
```  
Comprobar con algunos valores manualmente.  
-->

---

## Sistema de Tipos de Go

### <span v-click class="underline decoration-4 decoration-purple-600">La filosofía de tipos en Go</span>

---

## <span class="border-2 border-emerald-500 rounded-lg px-3 py-1">Estático y Fuerte</span>

Go verifica tipos en <span v-click class="bg-amber-200 px-1 rounded">**compilación**</span>, no en ejecución

<v-clicks>

- <span class="underline decoration-2 decoration-indigo-600">Estático</span>: Los tipos se determinan antes de ejecutar
- <span class="underline decoration-2 decoration-red-600">Fuerte</span>: No hay conversiones implícitas entre tipos
- Ejemplo: `int32` no se asigna automáticamente a `int64`
- Esto aporta <span class="underline decoration-2 decoration-green-600">seguridad</span> y <span class="underline decoration-2 decoration-blue-600">rendimiento</span>

</v-clicks>

```go
var x int32 = 100
var y int64 = 200
// y = x  // ❌ Error de compilación
y = int64(x)  // ✅ Conversión explícita
```

---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Tipado Estructural</span>

<span v-click class="underline decoration-wavy decoration-cyan-400">**"Duck typing"**</span> con interfaces

<v-clicks>

- Si <span class="underline decoration-2 decoration-yellow-600">camina como pato</span> y <span class="underline decoration-2 decoration-cyan-600">hace cuac</span>, es un pato 🦆
- No necesitas declarar `implements`
- Los tipos satisfacen interfaces <span class="underline decoration-2 decoration-pink-600">implícitamente</span>
- Flexibilidad sin sacrificar seguridad

</v-clicks>

```go
type Volador interface {
    Volar() string
}

type Pato struct{}
func (p Pato) Volar() string { return "Volando como pato" }
// Pato implementa Volador automáticamente ✨
```

---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Composición sobre Herencia</span>

Go no tiene <span v-click class="bg-red-200 px-1 rounded">**clases**</span>, usa <span class="underline decoration-2 decoration-teal-600">composición</span>

<v-clicks>

- **Sin herencia**: No hay `extends` o `inheritance`
- **Composición**: Incluir structs dentro de otros structs
- **Interfaces**: Para definir comportamientos polimórficos
- <span class="underline decoration-2 decoration-violet-600">Más flexible</span> que jerarquías de clases

</v-clicks>

```go
type Motor struct {
    Potencia int
}

type Coche struct {
    Motor  // Composición
    Marca string
}
```

---

## <span class="border-2 border-rose-500 rounded-lg px-3 py-1">Genéricos (Go 1.18+)</span>

<span v-click class="underline decoration-4 decoration-pink-600">**Código reutilizable**</span> para múltiples tipos

<v-clicks>

- <span class="underline decoration-2 decoration-rose-600">Una función</span> para múltiples tipos
- Sintaxis: `func Nombre[T any](param T)`
- **Constraints**: Restringen los tipos permitidos
- Mantiene <span class="underline decoration-2 decoration-green-600">seguridad de tipos</span> en compilación

</v-clicks>

```go
func Max[T int | float64](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

<!-- 
Notas del presentador: Describir la filosofía del sistema de tipos de Go.  
**Estático y fuerte:** Go es estáticamente tipado, o sea, todos los tipos se determinan antes de ejecutar el programa (no hay, por ejemplo, variables que cambien de tipo en runtime). “Fuerte” significa que no hay muchas conversiones automáticas; por ejemplo, un int32 no se asigna a un int64 sin conversión, y tampoco se puede usar un número como booleano, etc. Esto previene errores de tipo.  

**Tipado estructural (duck typing):** Introducir el concepto de **interfaces** en Go. Go emplea un sistema de tipos estructural: un tipo implementa una interfaz si tiene los métodos requeridos, sin declarar nada explícito. Esto es diferente a otros lenguajes donde debe declarar “implements” o heredar. Muchas veces se dice que Go usa “duck typing” en las interfaces: *"si camina como pato y hace cuac, es un pato"*.  
Mencionar que esta característica hace el sistema de tipos de Go muy flexible y permite polimorfismo desacoplado.  

**Composición vs herencia:** Señalar que al no haber clases, no hay herencia clásica. Go favorece la composición: por ejemplo, un struct puede incluir otro struct (anónimamente) y “heredar” sus campos y métodos de forma composicional. Los patrones de diseño en Go usan interfaces para definir comportamientos polimórficos. Esto resulta en programas modulares y flexibles:contentReference[oaicite:5]{index=5}.  

**Genéricos:** Explicar brevemente que fue una de las características más solicitadas y finalmente añadidas. Permite, por ejemplo, definir una función `Sumar` que sume elementos de cualquier tipo numérico, en vez de tener que escribir versiones separadas para int, float, etc. Los genéricos de Go soportan **constraints** (restricciones) que limitan qué tipos se pueden usar en un parámetro de tipo. Por ejemplo, se puede restringir a “tipos numéricos” o a “tipos comparables” etc. Veremos un ejemplo concreto a continuación.  
-->

---

## Interfaces en Go

### <span v-click class="underline decoration-4 decoration-indigo-500">Polimorfismo elegante en Go</span>

---

## <span class="border-2 border-emerald-500 rounded-lg px-3 py-1">¿Qué es una interfaz?</span>

Un conjunto de <span v-click class="bg-amber-200 px-1 rounded">**métodos abstractos**</span> que define comportamiento

<v-clicks>

- <span class="underline decoration-2 decoration-emerald-700">Solo firmas</span> de métodos, sin implementación
- Define **qué** puede hacer un tipo, no **cómo**
- Contratos que los tipos deben cumplir
- Ejemplos: `io.Reader`, `fmt.Stringer`

</v-clicks>

```go
type Escritor interface {
    Escribir(texto string) error
    Cerrar() error
}
```

* **¿Qué es una interfaz?** Un conjunto de métodos abstractos (firmas) que define un comportamiento. No contiene implementación, solo los métodos que un tipo debe tener para “cumplir” esa interfaz.
---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Implementación Implícita</span>

<span v-click class="underline decoration-wavy decoration-sky-400">**No hay `implements`**</span> en Go

<v-clicks>

- Si un tipo <span class="underline decoration-2 decoration-blue-700">tiene los métodos</span>, implementa la interfaz
- **Automático** y **transparente**
- Elimina acoplamiento entre tipos e interfaces
- <span class="underline decoration-2 decoration-purple-700">Flexibilidad máxima</span>

</v-clicks>

```go
type Archivo struct { nombre string }

func (a Archivo) Escribir(texto string) error { /* ... */ }
func (a Archivo) Cerrar() error { /* ... */ }

// Archivo implementa Escritor automáticamente! ✨
```

* **Implementación implícita:** En Go no se declara que un tipo implementa una interfaz; simplemente, si el tipo tiene todos los métodos que la interfaz requiere, entonces satisface esa interfaz. Esto elimina la necesidad de palabras clave como `implements`. Cualquier tipo puede implementar múltiples interfaces simplemente definiendo los métodos necesarios.
---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Polimorfismo Dinámico</span>

Trata tipos diferentes de forma <span v-click class="bg-lime-200 px-1 rounded">**uniforme**</span>

<v-clicks>

- <span class="underline decoration-2 decoration-indigo-700">Una interfaz</span>, múltiples implementaciones
- **Dynamic dispatch** en tiempo de ejecución
- Funciones genéricas sin sacrificar tipos
- Ejemplo: `fmt.Stringer` para personalizar impresión

</v-clicks>

```go
func ProcesarEscritor(w Escritor, datos string) {
    w.Escribir(datos)  // Funciona con cualquier implementación
    w.Cerrar()
}
```

* **Polimorfismo dinámico:** Las interfaces permiten tratar diferentes tipos de forma uniforme. Por ejemplo, si varios tipos implementan la interfaz `fmt.Stringer` (método `String()`), cualquiera de ellos puede usarse donde se espera un `fmt.Stringer`. Las interfaces son tipos de primera clase: una variable de tipo interfaz puede contener valores de distintos tipos que implementen esa interfaz.
---

## <span class="border-2 border-amber-500 rounded-lg px-3 py-1">Interfaz Vacía: any</span>

<span v-click class="underline decoration-4 decoration-orange-600">**Todos los tipos**</span> la implementan

```go
var cualquierCosa any = "Hola"
cualquierCosa = 42
cualquierCosa = []int{1, 2, 3}
```

<v-clicks>

- `any` es alias de `interface{}` (Go 1.18+)
- <span class="underline decoration-2 decoration-orange-700">Almacena cualquier valor</span>
- **Trade-off**: Pierdes seguridad de tipos
- Necesitas **type assertions** para usarla

</v-clicks>

```go
if str, ok := cualquierCosa.(string); ok {
    fmt.Println("Es string:", str)
}
```

* **Interfaz vacía (`interface{}` o `any`):** Es una interfaz sin métodos, por tanto **todos los tipos** la implementan automáticamente. Se usa para valores genéricos (similar a `Object` en otros lenguajes). Sin embargo, al usar una interfaz vacía, perdemos información de tipo estática y a menudo necesitamos hacer *type assertions* o reflección para usarlos. Nota: Go 1.18 introdujo el alias `any` equivalente a `interface{}`.

<!-- 
Notas del presentador: Definir interfaces y su uso:  
Recalcar que una **interfaz** en Go es un tipo “abstracto” que define un conjunto de métodos. Dar algún ejemplo simple: podría mencionar `io.Reader` (tiene un método Read) o `fmt.Stringer` (tiene String()).  

Explicar la **implementación implícita**: no necesitamos declarar “MiTipo implements X”. Simplemente si MiTipo tiene los métodos necesarios, Go lo considera implementado. Esto hace que el acoplamiento sea bajo: nuestros tipos no dependen de paquetes de interfaces directamente. Por ejemplo, en la librería estándar muchas funciones aceptan interfaces (como `io.Reader`); cualquier tipo nuestro con método `Read` adecuado puede usarse allí sin modificación.  

**Polimorfismo:** Indicar que las interfaces permiten escribir funciones más genéricas. Podemos tener una función que reciba una interfaz y podrá aceptar diferentes tipos concretos. Internamente, una variable de interfaz puede contener un valor de alguno de esos tipos (concepto de *dynamic dispatch* en tiempo de ejecución). Se puede mencionar cómo imprimir con `fmt.Println` usa la interfaz `fmt.Stringer` o la vacía para manejar argumentos de cualquier tipo.  

**Interfaz vacía (`interface{}`):** Importante subrayar que es implementada por todos, así que sirve para almacenar "cualquier cosa". Antes de generics, se usaba mucho en colecciones genéricas (ej. `[]interface{}` para lista heterogénea). El alias `any` hace el código más legible. También advertir que usar demasiado la interfaz vacía va en contra del tipado estático, porque luego hay que comprobar manualmente de qué tipo es el contenido (con type assertions o un type switch).  

Se puede comentar brevemente de *type assertion*: p.ej. `if val, ok := i.(string); ok { ... }` para convertir interfaz vacía a string si es ese tipo. Pero profundizar en eso podría no ser necesario a este nivel, salvo que pregunten.  
-->

---

### Ejemplo: Implementación de una interfaz

```go {1|3-4|6,8}
type Forma interface {
    Area() float64
}

type Circulo struct { Radio float64 }
func (c Circulo) Area() float64 { return 3.14 * c.Radio * c.Radio }

func imprimirArea(f Forma) { fmt.Println(f.Area()) }

imprimirArea(Circulo{10})  // 314
```

<!-- 
Notas del presentador: Veamos un ejemplo práctico con interfaces.  
Definimos una interfaz `Forma` que requiere un método `Area() float64`. Cualquier figura geométrica que tenga ese método cumple la interfaz.  

Luego definimos un tipo concreto `Circulo` con un campo `Radio`. Implementamos el método `Area()` para `Circulo` (como receptor valor en este caso). La fórmula de área que usamos es `π * Radio^2`; para simplificar usamos 3.14 como π.  

Definimos una función `imprimirArea(f Forma)` que recibe algo que cumple la interfaz `Forma`. Dentro, llama `f.Area()` y lo imprime. Esta función no sabe si f es un Circulo, un rectángulo u otra forma; le da igual mientras tenga Area().  

Finalmente, llamamos `imprimirArea(Circulo{10})`. Aquí un `Circulo` de radio 10 se pasa donde se espera un `Forma`. Esto es válido porque `Circulo` implementa `Area()`. La salida será 314 (aproximadamente el área de un círculo de radio 10).  

Este ejemplo ilustra el polimorfismo: la función `imprimirArea` puede trabajar con cualquier "Forma". Si tuviéramos otra struct, digamos `Cuadrado` con su propio Area(), también podríamos hacer `imprimirArea(Cuadrado{...})` sin cambiar la función.  

Destacar que no hubo que declarar explícitamente "Circulo implements Forma". Go lo dedujo. Si `Circulo` no tuviera Area(), la asignación a `Forma` daría error en compilación.  

También se puede mencionar que si pasáramos un puntero (`&Circulo{10}`), también funcionaría si el método Area está definido con receptor valor (Go automáticamente trata `&Circulo` que implementa mediante su valor).  

Este mecanismo es muy poderoso para diseñar APIs flexibles.  
-->

---

## Genéricos (Generics)

### <span v-click class="underline decoration-4 decoration-rose-600">La revolución de Go 1.18</span>

---

## <span class="border-2 border-emerald-500 rounded-lg px-3 py-1">Funciones y Tipos Genéricos</span>

<span v-click class="bg-amber-200 px-1 rounded">**Parametriza por tipo**</span> para máxima reutilización

<v-clicks>

- Sintaxis: `func Nombre[T any](param T)`
- <span class="underline decoration-2 decoration-slate-700">Parámetros de tipo</span> entre corchetes `[...]`
- El compilador <span class="underline decoration-2 decoration-cyan-500">infiere</span> o especificas el tipo
- También structs genéricos: `type Lista[T any] struct`

</v-clicks>

```go
func Intercambiar[T any](a, b T) (T, T) {
    return b, a
}

type Pila[T any] struct {
    elementos []T
}
```

* **Funciones y tipos genéricos:** Permiten parametrizar por tipo. Por ejemplo, `func Bar[T any](x T) { ... }` define una función genérica con un parámetro de tipo `T`. Al usarla, se puede especificar qué tipo toma `T` (o el compilador lo infiere). Igualmente se pueden definir estructuras genéricas: `type MiStruct[T any] struct { campo T }`.
---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Reutilización de Código</span>

<span v-click class="underline decoration-wavy decoration-cyan-400">**Una función**</span>, múltiples tipos

<v-clicks>

- <span class="underline decoration-2 decoration-gray-700">Antes</span>: Funciones duplicadas para cada tipo
- <span class="underline decoration-2 decoration-red-700">Problema</span>: `SumarInts`, `SumarFloats`, `SumarStrings`...
- <span class="underline decoration-2 decoration-green-700">Solución</span>: Genéricos mantienen seguridad de tipos
- Sin usar `interface{}` que perdía información

</v-clicks>

```go
// ❌ Antes: Duplicación
func SumarInts(a, b int) int { return a + b }
func SumarFloats(a, b float64) float64 { return a + b }

// ✅ Ahora: Una función genérica
func Sumar[T int | float64](a, b T) T { return a + b }
```

* **Reutilización de código:** Los genéricos evitan tener que escribir la misma función para distintos tipos. Antes de Go 1.18, a veces se usaban interfaces vacías o generadores de código para lograr algo similar, pero con pérdida de seguridad de tipo. Con generics, el compilador chequea que el tipo concreto cumple las restricciones y aplica el código apropiado para cada instanciación.
---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Constraints (Restricciones)</span>

<span v-click class="bg-lime-200 px-1 rounded">**Limita**</span> qué tipos se pueden usar

<v-clicks>

- `any`: <span class="underline decoration-2 decoration-blue-800">Cualquier tipo</span> (sin restricciones)
- `comparable`: Tipos que permiten `==` y `!=`
- Unión de tipos: `int | float64 | string`
- Interfaces personalizadas como constraints

</v-clicks>

```go
func Buscar[T comparable](slice []T, valor T) int {
    for i, v := range slice {
        if v == valor {  // Solo funciona con tipos comparables
            return i
        }
    }
    return -1
}
```

* **Constraints (restricciones):** Dentro de `[...]` se pueden imponer restricciones a los tipos permitidos. `any` indica que se acepta cualquier tipo. Existen constraints predefinidas como `comparable` (tipos que soportan `==`/`!=`) o se pueden usar interfaces como constraints (incluso con listas de tipos, ej. `interface{ ~int | ~int64 }`). Esto permite limitar, por ejemplo, que `T` sea “algún número” para poder usar operadores aritméticos.
---

## <span class="border-2 border-amber-500 rounded-lg px-3 py-1">Consideraciones</span>

Genéricos <span v-click class="underline decoration-4 decoration-orange-600">**simples**</span> por diseño

<v-clicks>

- <span class="underline decoration-2 decoration-purple-800">Limitados intencionalmente</span> para mantener simplicidad
- **No hay**: Sobrecarga de operadores custom
- **No hay**: Metaprogramación compleja (como C++)
- **Cuándo usar**: Evitar duplicación real de código

</v-clicks>

```go
// ✅ Buen uso: Estructuras de datos genéricas
type Cola[T any] struct { items []T }

// ❌ Overkill: Para casos simples
func Multiplicar[T int](a, b T) T { return a * b }  // Solo int? Usar int directamente
```

<v-click>

<span class="underline decoration-2 decoration-teal-800">Las interfaces siguen siendo útiles</span> para polimorfismo tradicional

</v-click>

* **Consideraciones:** El uso de generics en Go es intencionalmente limitado para mantener la simplicidad (no hay sobrecarga de operadores, ni metaprogramación compleja como en C++ templates). En muchos casos, las interfaces siguen siendo útiles para polimorfismo. Se recomienda usar generics cuando realmente evita duplicación significativa de código o provee abstracciones claras (por ejemplo, estructuras de datos genéricas como listas, colas, etc.).

<!-- 
Notas del presentador: Introducir el concepto de genéricos con entusiasmo ya que fue un cambio grande en Go 1.18:contentReference[oaicite:6]{index=6}:contentReference[oaicite:7]{index=7}.  
Explicar la sintaxis básica: las funciones genéricas llevan parámetros de tipo entre corchetes `[...]`. En la declaración podemos poner constraints en esos parámetros. Por defecto usamos `any` si no requerimos nada específico (any es alias de interface{}). Por ejemplo `func PrintSlice[T any](s []T) { ... }` sería una función genérica que acepta slice de cualquier tipo T.  

**Reutilización de código:** Mencionar cómo antes se tenían que escribir varias funciones similares (como en el post que duplicaban `sumarEnteros` y `sumarFlotantes`:contentReference[oaicite:8]{index=8}, etc.), o usar interface{} que perdía chequeo de tipos. Con generics, una sola función cubre múltiples tipos sin sacrificar seguridad.  

**Constraints:** Dar algún ejemplo sencillo, como `comparable` para decir "T debe permitir comparaciones con ==" (por ejemplo tipos básicos, string). Esto es útil para funciones de búsqueda en slice, etc. Otro ejemplo: se puede restringir a `~int | ~float64` etc., como en el ejemplo de suma que mostraremos. Comentar que el paquete `constraints` del estándar (desde Go 1.18) define `constraints.Ordered` que abarca tipos ordenables (números y strings).  

**Limitaciones intencionales:** Resaltar que Go implementó generics de manera sencilla: no hay especialización excesiva en tiempo de compilación del tipo de C++ (que genera código duplicado e iteraciones pesadas del compilador), aquí se basa en diccionarios de métodos en tiempo de ejecución (pero puede monomorfizar en algunos casos simples, detalle quizás no necesario mencionar). Lo importante es que no hay sobrecarga de operadores custom, ni herencia de generics; es más parecido a Java/C# en funcionalidad básica aunque con sintaxis propia.  

Podemos mencionar que la comunidad está adaptando librerías para usar generics donde conviene (por ejemplo, en estructuras de datos genéricas).  
-->

---

### Ejemplo: Función Genérica

```go {1|3-4}
func Sumar[T int | float64](x, y T) T {
    return x + y
}

fmt.Println(Sumar(3, 4))     // 7
fmt.Println(Sumar(3.5, 4.2)) // 7.7
```

<!-- 
Notas del presentador: Aquí definimos una función genérica `Sumar` que funciona tanto para enteros como para flotantes.  
En `func Sumar[T int | float64]` estamos declarando un parámetro de tipo `T` con la restricción de que T puede ser un int **o** un float64. Dentro de la función, `x` y `y` son del tipo genérico T, por lo que podemos sumarlos con `x + y` **solo** porque hemos restringido T a tipos numéricos que soportan el operador `+`.  

Llamamos `Sumar(3, 4)`. Aquí no indicamos explícitamente el tipo para T; el compilador infiere que T es `int` porque le pasamos enteros literales. Devuelve 7 (entero). Luego llamamos `Sumar(3.5, 4.2)`. Ahora T se infiere como `float64` (los literales con decimal se tratan como float64 por defecto). Retorna 7.7.  

Notemos que si intentáramos `Sumar("hola", "mundo")` no compilaría, porque string no está en la lista de tipos permitidos `int | float64`. Así, los genéricos con constraints garantizan que solo se usen tipos válidos y que las operaciones dentro de la función son seguras para esos tipos.  

Este ejemplo de constraint con unión de tipos es una forma sencilla de admitir un par de tipos. Para admitir "cualquier número", podríamos listar todos (int, int32, float64, etc.) o usar interfaces predefinidas (por ejemplo, una interfaz `Number` hipotética). La propuesta oficial incluyó interfaces tipo `constraints.Ordered` para casos comunes.  

En la práctica, podríamos sobrecargar esta función con más tipos si quisiéramos. Pero ojo: generics no significa que `+` funcione para cualquier T arbitrario; por eso hay que restringirlo a tipos que definan `+`.  
-->

---

### **Ejercicio:** Genéricos e Interfaces

Implementa una función **genérica** llamada `PrintSlice` que imprima todos los elementos de un slice de cualquier tipo. Por ejemplo:

```go
PrintSlice([]int{1,2,3})    // debería imprimir 1, 2, 3 en líneas separadas
PrintSlice([]string{"a","b","c"})  // debería imprimir a, b, c
```

Pistas:

* Define la función con un parámetro de tipo `T` sin restricciones especiales (usa `any`).
* Recorre el slice pasado e imprime cada elemento. (Puedes usar un `for range`).

*Reflexión:* ¿Podrías haber hecho lo mismo usando una interfaz en lugar de un genérico? ¿Qué ventajas ofrece el enfoque genérico en este caso?

<!-- 
Notas del presentador: Este ejercicio pide escribir una función genérica que imprima los elementos de un slice sin importar su tipo.  

Solución esperada:  
```go
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```  
Explicación: La función `PrintSlice` tiene un parámetro de tipo `T` que no tiene restricciones (any). Toma un slice de T (`s []T`). Luego simplemente itera sobre cada elemento `v` en `s` y lo imprime. `fmt.Println` puede imprimir cualquier cosa porque todos los tipos implementan la interfaz vacía (efectivamente usa reflection).  

Comentar que, alternativamente, podríamos haber hecho una función que tome `[]interface{}` antes de generics, pero habría detalles (por ejemplo, no podríamos pasar directamente `[]int` a `[]interface{}` sin copiar elementos). El genérico nos permite aceptar directamente `[]T` del tipo específico sin convertir.  

En la reflexión, la idea es que con interface vacía se podría lograr, pero con generics evitamos conversiones y mantenemos el chequeo de tipo (por ejemplo, si por error alguien le pasara algo que no es slice, ni siquiera compilaría).  

También mencionar que la versión genérica se especializa en tiempo de compilación para cada tipo usado, obteniendo potencialmente mejor rendimiento que hacer todo via interface{}.  
-->

---

## Interoperabilidad: CGo y WebAssembly

*(Integración de Go con otros lenguajes/entornos)*

<!-- 
Notas del presentador: Presentar la idea de que Go, además de su propio ecosistema, puede interactuar con código en otros lenguajes:  
1. **CGo:** integración con código C/C++ (principalmente C). Permite llamar bibliotecas o funciones escritas en C desde Go, o exponer funciones Go para que sean llamadas desde C. Útil para reutilizar librerías existentes en C o funcionalidades de sistema no expuestas directamente en Go.  
2. **WebAssembly:** la capacidad de compilar programas Go a WebAssembly (WASM) para ejecutarlos en un navegador o en entornos WASM fuera del navegador (por ejemplo WASI). Esto permite llevar lógica escrita en Go al front-end web, por ejemplo.  

Adelantar que veremos cada uno brevemente, sus usos y consideraciones.
-->

---

## CGo: Llamando código C desde Go

### <span v-click class="underline decoration-4 decoration-orange-500">Integrando Go con C</span>

---

## <span class="border-2 border-emerald-500 rounded-lg px-3 py-1">¿Qué es CGo?</span>

<span v-click class="bg-amber-200 px-1 rounded">**Puente**</span> entre Go y código C

<v-clicks>

- Permite a paquetes Go invocar código <span class="underline decoration-2 decoration-blue-500">**C**</span>
- Usa pseudo-paquete especial `"C"`
- Código C en comentarios especiales del archivo `.go`
- CGo genera los <span class="underline decoration-2 decoration-green-500">enlaces necesarios</span>

</v-clicks>

```go
/*
#include <stdio.h>
void hello() { printf("Hello from C!\n"); }
*/
import "C"

func main() {
    C.hello()
}
```

---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Uso Básico</span>

<span v-click class="underline decoration-wavy decoration-cyan-400">**Sintaxis simple**</span> pero poderosa

<v-clicks>

- Código C en comentario `/* ... */` antes de `import "C"`
- Llamar funciones: `C.nombreFuncion()`
- Tipos C: `C.int`, `C.char`, `C.double`
- <span class="underline decoration-2 decoration-purple-500">Conversión</span> entre tipos C y Go necesaria

</v-clicks>

```go
/*
#include <stdlib.h>
*/
import "C"

var cInt C.int = 42
var goInt int = int(cInt)  // Conversión explícita
```

---

## <span class="border-2 border-amber-500 rounded-lg px-3 py-1">Costos y Restricciones</span>

CGo tiene <span v-click class="bg-red-200 px-1 rounded">**trade-offs**</span> importantes

<v-clicks>

- <span class="underline decoration-2 decoration-red-500">**Overhead**</span>: ~100ns por llamada
- **Seguridad**: El código C puede causar segfaults
- **Memoria**: Gestión manual en C (`malloc`/`free`)
- <span class="underline decoration-2 decoration-orange-500">**Portabilidad**</span>: Requiere compilador C

</v-clicks>

<v-click>

💡 **Recomendación**: Usar solo cuando <span class="underline decoration-2 decoration-green-600">realmente necesario</span>

</v-click>

---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Ejemplos de Uso</span>

<span v-click class="bg-lime-200 px-1 rounded">**Casos comunes**</span> para CGo

<v-clicks>

- **Librerías estándar C**: `math.h`, `string.h`
- **Librerías de terceros**: SQLite, OpenCV
- **APIs del sistema**: Windows API, Unix syscalls
- **Al revés**: Exportar Go a C (`//export`)

</v-clicks>

```go
//export GoCallback
func GoCallback(x C.int) C.int {
    return x * 2
}

// Función C puede llamar GoCallback
```

<!-- 
Notas del presentador: Explicar en términos sencillos cómo usar CGo.  
Describir la necesidad: por ejemplo, "Quiero usar una librería escrita en C porque no hay en Go". CGo nos permite hacerlo.  

El bullet 1 define conceptualmente qué es cgo. Añadir que `import "C"` no es un paquete real sino una instrucción para el compilador cgo:contentReference[oaicite:14]{index=14}.  

Bullet 2 explica la sintaxis. Aclarar que el bloque de comentario con código C suele venir inmediatamente después de los imports de Go (o incluso antes, como en ejemplo). Dentro de ese bloque se pueden poner includes (`#include <...>`), definiciones de funciones C, etc. CGo compilará ese fragmento con un compilador C y enlazará.  
Luego, en Go, cualquier símbolo C incluido se puede referenciar como `C.symbol`. Los tipos también: por ejemplo `C.int` es el tipo entero de C (que en Go se representa de forma segura). Antes de usarlo en cálculos Go hay que convertirlo a un tipo Go (int32, int, etc.):contentReference[oaicite:15]{index=15}.  

Bullet 3 habla de costes. Recalcar que hay un costo de cambio de contexto entre Go y C. No es enorme (p.ej ~100 ns por llamada según algunas mediciones), pero comparado con llamar a una función Go trivial es bastante más. Si se llama en bucles muy intensivos puede afectar. Por tanto, conviene minimizar llamadas frecuentes dentro de loops críticos. Mencionar también que CGo obliga a que ciertas cosas no estén completamente en el runtime de Go: por ejemplo, las llamadas cgo pueden bloquear threads del scheduler, etc., y que hay implicaciones en la concurrencia (no entrar en detalle profundo, pero citar que "pone en riesgo la promesa de concurrencia de Go" como en Dave Cheney, pero suena muy técnico).  

Bullet 3 también: Indicar que al meter C, hay que compilarlo, lo que rompe la facilidad de compilar para múltiples plataformas sin tener ese compilador. Y que hay que cuidar la memoria: si C aloja memoria, Go no la liberará automáticamente (a menos que devolvamos un pointer y lo convirtamos a slice usando runtime finalizers... demasiado detalle; simplemente decir que memoria en C -> free en C manual).  

Bullet 4: Dar ejemplos concretos donde cgo se usa: interfazar con librerías de sistema (por ej, llamar a una función de Windows API, etc.), usar una librería criptográfica optimizada en C, etc.  

Añadir si se quiere que cgo se usa en la biblioteca estándar en algunos casos (por ej, `net` en ciertas resoluciones de DNS puede usar cgo).  

Concluir con: "No estás escribiendo solo Go, así que úsalo con precaución".  
-->

---

### Ejemplo: Llamando a una función de C (`math.h`)

```go {4|7}
/*
#include <math.h>
*/
import "C"

func main() {
    result := C.sqrt(16)
    fmt.Println(result)
}
```

<!-- 
Notas del presentador: Este código de ejemplo muestra el uso de CGo para llamar la función `sqrt` de la librería matemática de C (`math.h`).  

En el comentario arriba del import, incluimos `<math.h>`, que declara sqrt en C. Luego `import "C"` le indica al compilador que procese el bloque de C.  

En `main()`, llamamos `C.sqrt(16)`. Esto invoca la función C `sqrt` pasando 16 (que en C será considerado un entero, pero la función sqrt en C espera un `double`). CGo maneja automáticamente la conversión literal a `double` (16.0). `result` tendrá el tipo `C.double` en Go. Al hacer `fmt.Println(result)`, se imprimirá `4` (o `4e+00` dependiendo formato, pero usualmente `4`), porque sqrt(16) = 4.  

**Conversión de tipos:** Si quisiéramos usar `result` como un float64 de Go, podríamos convertir: `float64(result)`. En este caso `fmt.Println` puede manejarlo directamente aunque sea `C.double` porque implementa fmt.Stringer internamente, pero es mejor ser explícitos normalmente.  

**Nota de seguridad:** Llamar `C.sqrt` está bien porque sqrt no modifica globales ni mantiene estado; es pura. Si llamáramos funciones C más complejas, recordar manejar errores de C adecuadamente (por ejemplo, cgo no lanza excepciones, habría que checar códigos de retorno, etc.).  

Esto es un caso básico. Con CGo podemos también acceder a macros o constantes definidas en C (cgo las convierte si posibles, sino hay mecanismos). Por ejemplo, `C.LONG_MAX` si incluimos `<limits.h>`.  

Recordar que para compilar este código hace falta tener un compilador C disponible en el sistema, ya que cgo lo invocará para compilar la parte C.  
-->

---

## Go en WebAssembly (WASM)

### <span v-click class="underline decoration-4 decoration-cyan-500">Go llega al navegador</span>

---

## <span class="border-2 border-emerald-500 rounded-lg px-3 py-1">Go y WASM</span>

<span v-click class="bg-blue-200 px-1 rounded">**WebAssembly**</span> es un formato binario portable

<v-clicks>

- <span class="underline decoration-2 decoration-blue-500">Corre en navegadores</span> y otros entornos
- Go soporta WASM <span class="underline decoration-2 decoration-green-500">desde v1.11</span> (2018)
- Ejecuta lógica Go en el <span class="underline decoration-2 decoration-purple-500">cliente</span>
- También en runtimes WASM del servidor

</v-clicks>

```bash
# ¡Go en el navegador! 🚀
GOOS=js GOARCH=wasm go build -o app.wasm main.go
```

---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Compilación a .wasm</span>

<span v-click class="underline decoration-wavy decoration-cyan-400">**Variables de entorno**</span> especiales

<v-clicks>

- `GOOS=js`: <span class="underline decoration-2 decoration-orange-500">Navegador</span> como "OS"
- `GOARCH=wasm`: Arquitectura WebAssembly
- Resultado: Archivo <span class="underline decoration-2 decoration-rose-500">`.wasm`</span>
- Solo paquetes `main`, no librerías

</v-clicks>

```bash
# Comando completo
GOOS=js GOARCH=wasm go build -o mi-app.wasm main.go

# Optimizado para producción
GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o app.wasm main.go
```

---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Integración con JavaScript</span>

<span v-click class="bg-lime-200 px-1 rounded">**Puente**</span> entre Go y JS

<v-clicks>

- Go provee `wasm_exec.js` para <span class="underline decoration-2 decoration-green-500">soporte</span>
- Cargar con `WebAssembly.instantiateStreaming`
- Paquete `syscall/js` para <span class="underline decoration-2 decoration-cyan-500">DOM</span>
- Funciones bidireccionales Go ↔ JS

</v-clicks>

```go
import "syscall/js"

func main() {
    document := js.Global().Get("document")
    body := document.Get("body")
    body.Call("appendChild", h1Element)
}
```

---

## <span class="border-2 border-amber-500 rounded-lg px-3 py-1">Consideraciones</span>

WASM de Go tiene <span v-click class="bg-yellow-200 px-1 rounded">**trade-offs**</span>

<v-clicks>

- **Tamaño**: ~2MB mínimo (incluye runtime Go)
- **Rendimiento**: Similar a JS, más lento que Go nativo
- **WASI**: Go 1.21+ soporta ejecución fuera del navegador
- <span class="underline decoration-2 decoration-green-600">**Ventaja**</span>: Mismo código en backend y frontend

</v-clicks>

<v-click>

💡 **Ideal para**: Compartir lógica, algoritmos complejos, herramientas

</v-click>

<!-- 
Notas del presentador:  
Bullet 1: Explicar qué es WebAssembly brevemente (formato binario para web, cercano a ensamblador de stack). Mencionar que varios lenguajes pueden compilar a WASM (C, Rust, etc.) y Go es uno de ellos desde 2018 (v1.11):contentReference[oaicite:19]{index=19}.  

Bullet 2: Mostrar la variable de entorno GOOS y GOARCH. `GOOS=js` indica que el sistema operativo destino es un entorno JavaScript (el navegador actúa como un "OS"), y `GOARCH=wasm` la arquitectura. El comando dado compila el paquete main de un programa. Notar que solo se pueden compilar paquetes `main` a WASM, no librerías sueltas (por restricción actual):contentReference[oaicite:20]{index=20}.  

Bullet 3: Comentar que el `.wasm` por sí solo no corre en el navegador sin más; necesita el JS glue code. El `wasm_exec.js` se distribuye con Go (en `$GOROOT/misc/wasm/wasm_exec.js` hasta 1.23, luego movido a `lib/wasm/wasm_exec.js`). Este script implementa el objeto `Go` que prepara el entorno, maneja cosas como time, etc. En HTML se instancia el módulo WASM y luego se llama `go.run(instance)` para empezar la ejecución (como se ve en la documentación):contentReference[oaicite:21]{index=21}:contentReference[oaicite:22]{index=22}.  

También mencionar el paquete `syscall/js` disponible en Go: permite interaccionar con el DOM y demás. Por ejemplo, se puede llamar funciones JS o definir funciones para que JS las llame.  

Bullet 4: Hablar de limitaciones: tamaño - un simple hello world WASM en Go es ~2MB (aunque puede reducirse un poco con flags `-s -w` y compresión). Para comparar, Rust genera WASM más pequeño porque no incluye gran runtime. Sin embargo, la conveniencia de escribir en Go puede justificarlo en ciertos casos.  

Rendimiento: el código Go en WASM es generalmente más lento que el mismo en Go nativo, y similar o un poco más lento que JS puro debido al overhead del runtime. Pero puede ser suficientemente rápido para muchas aplicaciones.  

Mencionar WASI (WebAssembly fuera del browser): Con WASI, Go puede crear utilidades que corren en entornos sandbox. Desde v1.21 hay un target experimental `GOOS=wasip1 GOARCH=wasm`.  

Concluir que WebAssembly es un área en evolución y que es genial poder reutilizar conocimientos de Go para aplicaciones web sin tener que escribir en JS para ciertas lógicas.  
-->

---

### **Ejercicio:** CGo y WASM en la práctica

Piensa en escenarios donde utilizarías estas tecnologías:

* **CGo:** Menciona una librería o funcionalidad escrita en C/C++ que te gustaría aprovechar desde un programa Go. ¿Cómo te beneficiaría usar CGo en ese caso?
* **WebAssembly:** Describe una situación en la que compilar código Go a WebAssembly sería útil (por ejemplo, llevar parte de la lógica de tu aplicación al navegador). ¿Qué ventajas tiene usar Go en el navegador en lugar de JavaScript puro en ese contexto?

*Discute tus respuestas:* Comparte tus ideas con el grupo.

<!-- 
Notas del presentador: Este ejercicio es más conceptual para verificar comprensión y estimular ideas.  
Para **CGo**: Esperamos que mencionen casos como uso de librerías existentes (p.ej. usar una librería de compresión en C, o una API del sistema operativo, o funciones de rendimiento crítico escritas en C). Si nadie menciona, dar ejemplos: motores de base de datos en C (SQLite), librerías gráficas (OpenCV en C++), algoritmos optimizados en C (por ejemplo, una implementación de encriptación altamente optimizada). Usar CGo permitiría no reescribir todo eso en Go y simplemente integrarlo. Beneficios: aprovechar código maduro, posiblemente rendimiento (aunque la llamada cgo tiene overhead, el procesamiento intensivo dentro de C puede ser rápido).  

Para **WASM**: Esperamos respuestas como: uso de Go en frontend para reutilizar lógica (por ejemplo, validaciones que ya existen en backend, o motores de juego/escritura en Go ejecutados en web). Otra: crear aplicaciones web sin escribir JS, beneficiándose de la concurrencia de Go (aunque en WASM actualmente es single-thread por restricciones de navegador a menos que se use threads webasm experimental). Ventaja: mismo lenguaje en backend y frontend, seguridad del tipado de Go, posibilidad de usar paquetes Go en el navegador (por ejemplo, un parser, o lógica de negocios).  

Después de unos minutos, pedir a algunos estudiantes que compartan. Guiar la discusión hacia las consideraciones correctas: por ejemplo, "usar Go en WASM podría ser bueno para una aplicación que haga procesamiento intensivo (como decodificar video) aprovechando Go, pero hay que sopesar el tamaño del binario". O "CGo es útil si ya tienes un módulo en C comprobado, pero recuerda los riesgos de estabilidad y performance".  

Este ejercicio también hace que piensen en las limitaciones: Por ejemplo, si alguien dice "usaría CGo para todo así no aprendo a hacer X en Go", aclarar que abusar de CGo no es ideal, uno pierde muchos beneficios de Go al hacerlo.  

No hay una respuesta única correcta; se busca que comprendan cuándo **sí** vale la pena.  
-->

---

## Recursos para seguir aprendiendo Go

### <span v-click class="underline decoration-4 decoration-green-500">Tu viaje en Go apenas comienza</span>

---

## <span class="border-2 border-emerald-500 rounded-lg px-3 py-1">Tutoriales Interactivos</span>

Aprende <span v-click class="bg-amber-200 px-1 rounded">**practicando**</span> directo en el navegador

<v-clicks>

- **A Tour of Go**: <span class="underline decoration-2 decoration-blue-500">Tutorial oficial</span> interactivo
- Sintaxis básica, interfaces y <span class="underline decoration-2 decoration-purple-500">concurrencia</span>
- <span class="underline decoration-2 decoration-green-500">Disponible en español</span>
- Perfecto para <span class="underline decoration-2 decoration-orange-500">consolidar</span> lo aprendido hoy

</v-clicks>

🔗 **go.dev/tour**

---

## <span class="border-2 border-blue-500 rounded-lg px-3 py-1">Guías y Ejemplos</span>

Recursos <span v-click class="underline decoration-wavy decoration-cyan-400">**esenciales**</span> para código limpio

<v-clicks>

- **Effective Go**: <span class="underline decoration-2 decoration-green-600">Buenas prácticas</span> y estilo idiomático
- **Go by Example**: Ejemplos <span class="underline decoration-2 decoration-blue-600">concisos</span> y completos
- Cada concepto = <span class="underline decoration-2 decoration-purple-600">programa funcional</span>
- <span class="underline decoration-2 decoration-orange-600">Referencia rápida</span> perfecta

</v-clicks>

🔗 **gobyexample.com**

---

## <span class="border-2 border-purple-500 rounded-lg px-3 py-1">Documentación Oficial</span>

La <span v-click class="bg-lime-200 px-1 rounded">**fuente de verdad**</span> sobre Go

<v-clicks>

- **go.dev/doc**: <span class="underline decoration-2 decoration-indigo-500">Documentación completa</span>
- **Language Specification**: Detalles técnicos del lenguaje
- **Getting Started**: Para <span class="underline decoration-2 decoration-green-600">comenzar</span> desde cero
- **FAQ**: Respuestas a <span class="underline decoration-2 decoration-red-500">preguntas frecuentes</span>

</v-clicks>

📚 **Siempre actualizada y precisa**

---

## <span class="border-2 border-amber-500 rounded-lg px-3 py-1">Comunidad</span>

<span v-click class="underline decoration-4 decoration-orange-600">**Conecta**</span> con otros Gophers

<v-clicks>

- **Stack Overflow**: Tag `go` para <span class="underline decoration-2 decoration-blue-700">dudas técnicas</span>
- **Slack Gophers**: Canal oficial con expertos
- **Gophers Latam**: <span class="underline decoration-2 decoration-green-700">Comunidad en español</span>
- **Blog oficial**: go.dev/blog para novedades

</v-clicks>

<v-click>

💡 **¡No programes solo!** La comunidad Go es muy acogedora 🤗

</v-click>

---

## <span class="border-2 border-rose-500 rounded-lg px-3 py-1">Libros y Cursos</span>

Para <span v-click class="bg-pink-200 px-1 rounded">**profundizar**</span> aún más

<v-clicks>

- **"The Go Programming Language"** (Kernighan & Donovan)
- **"Head First Go"**: Más <span class="underline decoration-2 decoration-cyan-700">introductorio</span>
- **Platzi, Udemy**: Cursos en línea en español
- **Go en Español**: Canal de Nacho Pacheco

</v-clicks>

<v-click>

🚀 **Mejor forma de aprender**: ¡Escribe código! Haz un proyecto pequeño esta semana

</v-click>

<!-- 
Notas del presentador: Recomendar recursos adicionales:  
1. **Tour of Go:** Si no lo han hecho, animarlos a completarlo, ya que cubre muchos temas que vimos con ejercicios guiados:contentReference[oaicite:26]{index=26}.  
2. **Effective Go:** Explicar que es casi lectura obligada para entender cómo escribir buen código Go, convenciones, trucos, etc:contentReference[oaicite:27]{index=27}.  
3. **Go by Example:** Útil para rápidamente ver "cómo se hace X en Go" con código de ejemplo. Mencionar que es libre en web.  
4. **Official docs & spec:** Si alguien quiere detalles profundos, la spec es un documento formal pero legible. La página de documentación de go.dev tiene tutoriales (Getting Started, etc.), la wiki (Go Wiki) también tiene muchos tópicos. Para hispanohablantes, el sitio *¡Go en Español!* tiene traducciones de muchos documentos oficiales:contentReference[oaicite:28]{index=28}.  
5. **Community:** Animar a aprovechar la comunidad. Mencionar que preguntar en StackOverflow suele dar buenas respuestas técnicas. En Slack, existe el workspace Gophers con canales por tema e idioma (puede haber #spanish por ejemplo). También existen grupos en Discord, Telegram, etc., para Go en español. Si hay un Go meetup en su ciudad o virtual, recomendarlos.  
6. **Libros/Cursos:** Recomendar K&D "The Go Programming Language" como un libro de referencia completa (está en inglés, no sé si hay traducción al español; posiblemente no, pero se puede mencionar). También *Head First Go* para algo más introductorio pero en inglés también. Si conocen plataformas en español (Platzi tiene un curso de Go, etc.), mencionarlas.  

Terminar motivando: la mejor forma de aprender es *escribir código*: proponer que intenten un proyectito pequeño en Go para practicar.  
-->

---

## ¿Preguntas?

### ¡Gracias por su atención! 🙌

<!-- 
Notas del presentador: Abrir espacio para preguntas finales de los estudiantes. Animar a que compartan cualquier duda o que comenten su experiencia realizando los ejercicios.  
Reforzar que la práctica continua es clave: sugerir que implementen un pequeño programa en Go esta semana (puede ser algo sencillo como un conversor de unidades, un web scraper básico, etc.) para afianzar lo aprendido.  
Agradecer la participación y el tiempo. 
-->
