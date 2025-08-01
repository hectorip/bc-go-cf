### Guía rápida: cómo **subrayar, resaltar y “encerrar”** elementos en Slidev

*(Markdown + componentes incorporados + Tailwind/UnoCSS que Slidev trae por defecto)*

---

## 1 · Subrayados

| Qué quieres lograr              | Ejemplo Markdown / HTML                                                                    | Vista previa\*                                                                |
| ------------------------------- | ------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------- |
| Subrayar simple                 | `<u>texto subrayado</u>`                                                                   | <u>texto subrayado</u>                                                        |
| Subrayado con color             | `<span class="underline decoration-4 decoration-emerald-500">Importante</span>`            | <span class="underline decoration-4 decoration-emerald-500">Importante</span> |
| Subrayado animado al hacer clic | `<span v-click class="underline decoration-wavy decoration-yellow-400">def palabra</span>` | *Aparece al avanzar*                                                          |

\*En la presentación verás el estilo; aquí es solo referencia.

**Claves**

* Puedes usar cualquier utilidad *underline / decoration-* de Tailwind/UnoCSS.
* El atributo `v-click` hace que el elemento aparezca en el siguiente clic.
* `decoration-wavy`, `decoration-dotted`, etc. te dan estilos diferentes.

---

## 2 · Resaltados (highlight)

| Opción            | Código                                         | Comentario                                                                                 |
| ----------------- | ---------------------------------------------- | ------------------------------------------------------------------------------------------ |
| Resaltado rápido  | `==texto==`                                    | El plugin `markdown-it-mark` ya viene activado en Slidev: `==texto==` → <mark>texto</mark> |
| Con `<mark>`      | `<mark class="px-1 bg-amber-200">texto</mark>` | Útil para aplicar utilidades Tailwind extra                                                |
| Resaltado gradual | `<span v-click>==¡Atención!==</span>`          | El texto aparece y ya viene marcado                                                        |

**Tip interactivo**

```md
Hazle clic 👉 <span v-click class="bg-lime-200 px-1 rounded">==detalle oculto==</span>
```

Cada clic avanza y muestra la parte resaltada.

---

## 3 · “Encerrar” (marco / círculo)

### Marco rectangular

```html
<span class="border-2 border-sky-500 rounded px-2 py-0.5">
  variable crítica
</span>
```

### “Pastilla” o círculo

```html
<span class="inline-block rounded-full border-4 border-rose-500 px-3 py-1">
  ⚠️
</span>
```

### Bocadillo con sombra + animación

```html
<div v-click
     class="inline-block border-2 border-purple-600 rounded-lg shadow-lg p-3
            transition duration-300 ease-out hover:scale-105">
  <strong>¡Nuevo!</strong> soporte a genéricos
</div>
```

* **`rounded` / `rounded-lg` / `rounded-full`** definen la forma.
* Combina `border-*` (ancho) y `border-color` para el contorno.
* Añade `shadow`, `hover:` y `transition` para dar dinamismo.

---

## 4 · Subrayar / resaltar **en código**

Slidev usa Shiki; puedes marcar líneas, palabras y hacer aparición gradual.

```go {1|3|5}
package main          // ⬅ línea resaltada

import "fmt"          // aparece al avanzar

func main() {         // ⬅ más tarde, línea 5 se ilumina
    fmt.Println("Hi")
}
```

* Entre llaves del bloque de código pon `{highlightLines:1,3,5}` o la forma breve mostrada arriba.
* Usa `v-click` en el bloque para que **todo** el snippet aparezca paso a paso:

  ````md
  ```go {v-click=1,4,6}   ← mostrará 1 → 4 → 6
  ````

---

## 5 · Combinar efectos

```md
¡Go 1.22 trae <span
  v-click="1"
  class="underline decoration-wavy decoration-indigo-500"
>==genéricos mejorados==</span> y
<span
  v-click="2"
  class="border-2 border-fuchsia-600 rounded px-1"
>canales optimizados</span>!
```

* **Paso 1:** subraya y resalta “genéricos mejorados”.
* **Paso 2:** “encierra” la segunda frase.
* Así mantienes el texto mínimo y guías la atención con cada avance.

---

## 6 · Buenas prácticas de diseño Slidev

1. **Menos texto, más énfasis**: usa subrayados/resaltados en frases cortas; el resto explícalo verbalmente o en *presenter notes*.
2. **Secuencia**: con `v-click` o `v-show` controla el flujo mental del público.
3. **Coherencia**: define un **color de énfasis** (por ejemplo `emerald-500`) y úsalo en todos los subrayados; otro color para marcos.
4. **Accesibilidad**: asegúrate de buen contraste (usa `bg-yellow-200` + texto negro, no amarillo puro).
5. **Evita abuso**: si todo está subrayado, nada lo está. Resalta solo lo esencial de cada diapositiva.

---

### “Snippet” listo para copiar

```md
Las tres palabras clave son
<u class="decoration-4 decoration-sky-600" v-click>paquete</u>,
<mark v-click>tipo</mark> y
<span class="border-2 border-rose-500 rounded px-1" v-click>interfaz</span>.
```

---

#### ¿Dudas?

Prueba estos fragmentos en tu `slides.md`, ajusta colores/clases a tu gusto y ¡listo!
