---
name: definicion-de-marca
description: Construye la definición de marca y las reglas visuales de un proyecto — paleta de color, tipografía, tono de voz, y design tokens — como parte de la etapa de /feature, ANTES de implementar cualquier UI. Usa esta skill cuando el usuario esté arrancando un proyecto/feature nuevo y pida "definir la marca", "reglas visuales", "design system", "paleta de colores del proyecto", "tono de voz", "identidad visual", o cuando esté a punto de construir una interfaz y todavía no exista una definición de marca en `specs/`. Esta skill NO genera imágenes finales (favicons, og-images, logos) — produce el documento de reglas que luego usan otras skills/features para construir esos assets o cualquier UI de forma consistente.
---

# Definición de marca

## Qué hace

Produce un documento de **definición de marca y reglas visuales** (`specs/brand-definition.md` + `specs/design-tokens.json`) que actúa como fuente de verdad para todo lo que se construya después en el proyecto: color, tipografía, tono de voz, y los tokens concretos que el código va a usar. Es un artefacto de la etapa `/feature` — se produce **antes** de escribir la primera línea de UI, no después.

## Checklist antes de empezar

1. **¿Ya existe `specs/brand-definition.md`?** Si existe, no lo regeneres desde cero — léelo, y si el pedido es extender o ajustar algo puntual (por ejemplo agregar un color semántico de error), edita esa sección específica en vez de rehacer el documento completo.
2. **¿Tengo el tono/personalidad del producto, o solo colores?** Una definición de marca sin tono de voz está incompleta — no te quedes solo con la paleta.

## Flujo de trabajo

### 1. Reunir contexto del producto

Antes de proponer nada, entiende:
- Nombre del proyecto/producto.
- Qué hace y para quién (una frase basta).
- Tono deseado: ¿serio/corporativo, cálido/cercano, técnico/minimalista, lúdico? Si el usuario no lo dice, pregúntalo — no lo asumas a partir del rubro (un producto de salud puede ser cálido, una app de juegos puede ser minimalista y seria; ver el patrón general en `ejemplos/ejemplos-de-uso.md`).
- Referencias que le gusten ("algo estilo Stripe", "onda Notion") —
  útiles como punto de partida de tono/paleta, nunca para copiar la
  identidad visual real de esa marca.
- Restricciones ya existentes: ¿hay un logo, un color ya elegido, una
  guía de marca corporativa que haya que respetar?

### 2. Definir la paleta de color

Usa `referencias/paletas-color.md` como banco de categorías e
inspiración (fría/profesional, cálida/enérgica, vibrante/audaz,
naturaleza/bienestar, degradados) — elige por **tono**, no por rubro
(ver el patrón general al final de `ejemplos/ejemplos-de-uso.md`).

Define como mínimo:
- Color primario (`#RRGGBB`)
- Color de acento
- Color de texto sobre fondo primario
- Colores semánticos si el proyecto los va a necesitar pronto: éxito,
  error, advertencia, información (no los inventes si el usuario no
  va a construir UI con estados todavía; agrégalos cuando haga falta)
- Fondo claro y fondo oscuro si el proyecto va a soportar ambos temas

**Verifica el contraste de cada combinación texto/fondo antes de
darla por definitiva.** Usa `scripts/check_contrast.py` en vez de
pedirle al usuario que lo verifique manualmente — corre el script vos
mismo y reporta el resultado (ratio y si pasa AA/AAA). Ver
`referencias/accesibilidad-color.md` para los umbrales y qué hacer si
una combinación falla.

### 3. Definir tipografía

Usa `referencias/tipografia-y-tono.md` para elegir:
- Familia tipográfica para títulos y para cuerpo (pueden ser la misma).
- Escala de tamaños (una escala modular simple alcanza para la mayoría
  de proyectos; no inventes una escala de 15 pasos si el proyecto no
  la necesita).
- Pesos a usar (regular/medium/bold es suficiente en la mayoría de
  casos).

La elección de tipografía debe ser coherente con el tono definido en
el paso 1 (una fuente redondeada y liviana para un producto calmado,
una monoespaciada para un producto técnico/dev-tool) — ver ejemplos en
`ejemplos/ejemplos-de-uso.md`.

### 4. Definir tono de voz

No es opcional. Usa `referencias/tipografia-y-tono.md` (sección de
tono) para llegar a:
- 3-4 adjetivos que describan cómo "habla" el producto.
- Un par de ejemplos de copy corto (ej. un mensaje de error, un botón
  de call-to-action) escritos en ese tono, para que quede concreto y
  no solo como lista de adjetivos.
- Qué evitar explícitamente (ej. "nunca usar jerga corporativa vacía",
  "nunca sonar condescendiente").

### 5. Reglas visuales adicionales (si aplican)

Solo si el proyecto ya lo necesita — no las inventes por completar el
documento:
- Espaciado/grid básico, si ya se sabe que se va a construir UI pronto.
- Reglas de iconografía (estilo de línea, grosor, tamaño base).
- Reglas de logo si el usuario ya tiene uno (zona de exclusión, tamaño
  mínimo, qué no hacer con él) — esta skill no diseña el logo en sí.

### 6. Producir el documento final

Escribe `specs/brand-definition.md` usando
`assets/templates/definicion-de-marca-template.md` como estructura
base, y `specs/design-tokens.json` usando
`assets/templates/design-tokens.json.template` con los valores
concretos ya decididos, listos para que el código los consuma
directamente (variables CSS, tema de un framework de UI, etc.).

Si el proyecto ya tiene `AGENTS.md` (convención de la skill
`cost-guard`), agrega ahí una línea que apunte a
`specs/brand-definition.md`, para que sesiones futuras lo encuentren
sin tener que preguntar de nuevo.

### 7. Marcar lo que quedó sin confirmar

Si en cualquier paso no tuviste una respuesta clara del usuario (tono
ambiguo, sin preferencia de color, sin ejemplos de referencia), no lo
resuelvas con un supuesto silencioso: proponé una opción concreta pero
márcala en el documento final bajo una sección "Pendiente de
confirmar", igual que se marcarían reglas de negocio no confirmadas en
otras specs del proyecto. Esto evita que una decisión de marca
inventada se trate después como definitiva.

## Accesibilidad (obligatorio revisar antes de entregar)

- Relación de contraste ≥ 4.5:1 para texto normal sobre fondo, ≥ 3:1
  para texto grande — verificado con `scripts/check_contrast.py`, no
  a ojo.
- Evita que el rojo/verde sea la única forma de distinguir un estado
  semántico (éxito/error); acompaña con ícono o texto.
- Deja constancia en el documento final de qué combinaciones se
  verificaron y con qué resultado — no solo que "se verificó".
- Detalle completo de umbrales y qué hacer ante un fallo en
  `referencias/accesibilidad-color.md`.

## Qué NO hacer

- No generes las imágenes finales (favicon, og-image, logo) — esta
  skill entrega las reglas, no los archivos de imagen.
- No copies la identidad visual real de una marca de referencia
  mencionada por el usuario; úsala solo como punto de partida de tono
  y aclara siempre que el resultado es "inspirado en", no una réplica.
- No inventes reglas de compliance de marca corporativa (ej. "así lo
  exige nuestra guía de marca") si el usuario no las mencionó — si el
  proyecto ya tiene una guía de marca externa, pídela o pregunta por
  sus restricciones en vez de asumirlas.
- No agregues tokens que el proyecto todavía no necesita (escalas de
  espaciado completas, temas oscuro/claro, colores semánticos) solo
  por completar el documento — agrégalos cuando el feature que se está
  construyendo los requiera.

## Referencias

- `referencias/paletas-color.md` — paletas de referencia rápida por
  categoría (tech, startups, naturaleza/bienestar, degradados).
- `referencias/tipografia-y-tono.md` — cómo elegir tipografía coherente
  con el tono, y cómo definir tono de voz con adjetivos + ejemplos de
  copy.
- `referencias/accesibilidad-color.md` — umbrales de contraste, cómo
  interpretar el resultado de `scripts/check_contrast.py`, y qué hacer
  cuando una combinación no pasa.

## Scripts

- `scripts/check_contrast.py` — calcula la relación de contraste WCAG
  entre dos colores hex y reporta si pasa AA/AAA, sin depender de una
  herramienta externa.

## Ejemplos completos

Ver `ejemplos/ejemplos-de-uso.md` para tres casos resueltos de
principio a fin: una startup tech B2B (paleta fría/profesional), un
producto de bienestar (paleta de naturaleza) y un proyecto open-source
(paleta vibrante) — cada uno mostrando cómo queda el
`brand-definition.md` final, no las imágenes.

## Siguientes pasos tras producir la definición

1. Guarda `specs/brand-definition.md` y `specs/design-tokens.json` en
   el repo, versionados como cualquier otra spec.
2. Referencia `specs/brand-definition.md` desde `AGENTS.md` si el
   proyecto lo usa (convención de la skill `cost-guard`).
3. Cuando se implemente la primera UI del feature, esa implementación
   debe leer `specs/design-tokens.json` en vez de elegir colores o
   tipografía de nuevo.
4. Si más adelante se necesitan los assets finales (favicon, og-image,
   logo), esos tokens son el input — no se vuelven a decidir ahí.
