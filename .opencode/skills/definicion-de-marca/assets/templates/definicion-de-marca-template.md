# Definición de marca — <nombre del proyecto>

> Este documento es la fuente de verdad de identidad visual y tono del proyecto. Cualquier feature que construya UI, copy, o assets de marca debe referenciarlo en vez de decidir estos valores de nuevo. Los valores concretos (hex, tamaños) también están disponibles en `specs/design-tokens.json` para consumo directo desde código.

## Producto

- **Nombre:** <nombre>
- **Qué hace / para quién:** <una frase>
- **Tono buscado:** <"cercano y directo, sin sonar corporativo">

## Paleta de color

| Token | Valor | Uso |
|---|---|---|
| `color-primary` | `#RRGGBB` | <uso> |
| `color-accent` | `#RRGGBB` | <uso> |
| `color-text-on-primary` | `#RRGGBB` | <uso> |
| `color-background` | `#RRGGBB` | <uso> |
<!-- Agregar solo los colores semánticos que el proyecto realmente necesita todavía -->

### Verificación de contraste

- Texto `<color>` sobre `<color>` → contraste `<X.X>`:1 (`<pasa/no pasa>` AA <detalle>)
- <repetir por cada combinación relevante>

## Tipografía

- **Familia principal:** <nombre>
- **Familia secundaria (si aplica):** <nombre> — uso: <ej. "código, nombres de producto">
- **Escala de tamaños:**

| Token | Tamaño | Uso |
|---|---|---|
| `text-base` | 16px | Cuerpo |
| `text-xl` | 24px | Títulos de sección |
<!-- Agregar solo los pasos que el proyecto va a usar -->

## Tono de voz

**Adjetivos:** <adjetivo 1>, <adjetivo 2>, <adjetivo 3>

**Ejemplos de copy en este tono:**
- Mensaje de error: "<ejemplo>"
- Call-to-action: "<ejemplo>"
- Estado vacío: "<ejemplo>"

**Qué evitar:**
- <regla 1>
- <regla 2>

## Reglas visuales adicionales

<solo si aplican todavía al proyecto — espaciado, iconografía, reglas de logo>

## Pendiente de confirmar

> No asumido — requiere respuesta del usuario antes de tratarse como definitivo.
- <ítem pendiente, si lo hay>

## Referencias de marca usadas como inspiración

- <marca mencionada por el usuario, si aplica> — usada solo como punto de partida de tono/paleta, no como identidad a replicar.
