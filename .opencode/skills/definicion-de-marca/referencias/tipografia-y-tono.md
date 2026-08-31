# Tipografía y tono de voz

Guía para los pasos 3 y 4 del flujo de `definicion-de-marca`: elegir tipografía coherente con el tono del producto, y definir el tono de voz de forma concreta (no solo como lista de adjetivos).

## Tipografía por personalidad

| Personalidad del producto | Familia sugerida (categoría) | Notas |
|---|---|---|
| Corporativo / serio / B2B | Sans-serif geométrica o neutra (estilo Inter, Helvetica) | Pesos regular/medium/bold alcanzan; evita itálicas decorativas |
| Cálido / cercano / consumo | Sans-serif redondeada (estilo Nunito, Quicksand) | Pesos más livianos por defecto (regular/medium), reservar bold para títulos |
| Técnico / dev-tool / open-source | Monoespaciada para código o acentos, sans-serif neutra para el resto | Usar la monoespaciada con moderación (nombres de producto, código), no en párrafos largos |
| Calmado / mindfulness / bienestar | Sans-serif redondeada y liviana | Evitar pesos muy bold; dar espacio (line-height generoso) |
| Editorial / contenido largo | Serif para títulos + sans-serif para UI | Solo si el producto es principalmente de lectura (blog, docs) |

No hace falta elegir dos familias distintas para títulos y cuerpo — una sola familia con distintos pesos suele ser más simple de mantener y sigue viéndose consistente. Usa dos familias solo cuando el contraste entre ellas aporte algo real (monoespaciada para código en un dev-tool).

## Escala de tamaños

Una escala modular simple alcanza para la mayoría de proyectos. Punto de partida razonable (ajustar según necesidad real del feature, no completar los 8 pasos si el proyecto solo va a usar 3):

| Token | Tamaño | Uso típico |
|---|---|---|
| `text-xs` | 12px | Texto auxiliar, labels pequeños |
| `text-sm` | 14px | Texto secundario |
| `text-base` | 16px | Cuerpo de texto |
| `text-lg` | 18px | Texto destacado, subtítulos |
| `text-xl` | 24px | Títulos de sección |
| `text-2xl` | 32px | Títulos principales |
| `text-3xl` | 48px | Hero / landing |

No agregues pasos intermedios (`text-md`, `text-xl-2`, etc.) salvo que el feature que se está construyendo realmente los necesite.

## Definir el tono de voz

El tono de voz no es una lista de adjetivos sueltos — necesita quedar demostrado con ejemplos concretos para ser accionable.

### Paso 1 — Elegir 3-4 adjetivos

Ejemplos de pares que ayudan a acotar (elige un punto en cada eje, no los dos extremos):
- Formal ←→ Cercano
- Serio ←→ Lúdico
- Técnico/preciso ←→ Simple/accesible
- Discreto ←→ Enfático

### Paso 2 — Escribir 2-3 ejemplos de copy en ese tono

No alcanza con decir "cercano y simple" — hay que mostrarlo. Ejemplos de qué pedir/producir:
- Un mensaje de error (ej. de un formulario que falló).
- Un texto de botón de call-to-action.
- Una línea de bienvenida o de estado vacío ("no tienes tareas todavía").

### Paso 3 — Qué evitar explícitamente

Tan útil como decir qué es el tono es decir qué NO es. Ejemplos de formulaciones concretas (adaptar al producto real, no copiar
genéricamente):
- "Nunca usar jerga corporativa vacía tipo 'sinergia' o 'solución disruptiva'."
- "Nunca sonar condescendiente en mensajes de error — explicar qué pasó, no regañar al usuario."
- "Evitar exclamaciones excesivas si el tono es serio/profesional."

## Coherencia entre tipografía, tono y color

Antes de cerrar la definición, revisa que las tres piezas cuenten la misma historia: una paleta "Naturaleza y bienestar" con una tipografía monoespaciada y un tono de voz técnico-seco no es coherente — alguna de las tres piezas probablemente esté mal elegida. Si detectas esta incoherencia, señálasela al usuario en vez de entregar igual el documento.
