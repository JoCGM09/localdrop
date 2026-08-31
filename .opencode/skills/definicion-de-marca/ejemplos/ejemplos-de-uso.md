# Ejemplos de uso — definicion-de-marca

Tres casos resueltos de principio a fin, mostrando la solicitud del usuario, las decisiones de paleta/tipografía/tono, y cómo queda el `specs/brand-definition.md` resultante.
---

## Ejemplo 1 — Startup tech B2B

**Solicitud del usuario:**
> "Estoy por empezar a construir Fluxa, una herramienta que ayuda a equipos a automatizar flujos de aprobación."

**Decisiones:**
- Tono acordado: serio pero no frío — "profesional, directo, sin jerga corporativa".
- Categoría de paleta: "Fría y profesional" (`referencias/paletas-color.md`), coherente con el tono.
- Primario: `#0066CC` (azul cobalto). Acento: `#00B4D8` (azul cielo).
- Verificación de contraste con `scripts/check_contrast.py`: texto `#FFFFFF` sobre `#0066CC` → 5.1:1, pasa AA para texto normal.
- Tipografía: sans-serif neutra (estilo Inter), sin familia secundaria coherente con "directo, sin adornos".
- Tono de voz: adjetivos "directo, confiable, sin rodeos". Ejemplo de copy de error: "No pudimos guardar el flujo. Revisa el paso 2 e intenta de nuevo." Qué evitar: "nunca usar frases tipo 'sinergia' o 'solución disruptiva'".

**Resultado (`specs/brand-definition.md`):**
- Paleta, tipografía y tono documentados con la estructura de `assets/templates/definicion-de-marca-template.md`.
- `specs/design-tokens.json` con los valores concretos listos para que la primera pantalla del producto los use directamente.
- Sin pendientes: el usuario confirmó tono y referencias en la conversación.

---

## Ejemplo 2 — Producto de bienestar / mindfulness

**Solicitud del usuario:**
> "Voy a armar Respira, una app de meditación. Quiero dejar la identidad de marca definida antes de tocar código — algo tranquilo, tipo naturaleza."

**Decisiones:**
- Tono acordado: "calmado, cercano, nada corporativo".
- Categoría de paleta: "Naturaleza y bienestar" → tema bosque.
- Primario: `#2D6A4F` (verde bosque). Acento: `#52B788` (verde salvia).
- Verificación de contraste: texto `#FFFFFF` sobre `#2D6A4F` → 5.2:1, pasa AA y AAA para texto normal.
- Tipografía: sans-serif redondeada y liviana, pesos regular/medium únicamente (se evitó bold pesado, coherente con "calmado").
- Tono de voz: adjetivos "calmado, cercano, sin urgencia". Ejemplo de estado vacío: "Todavía no hiciste tu primera sesión. Cuando quieras, acá vas a poder empezar." Qué evitar: "nunca usar lenguaje de urgencia o gamificación agresiva ('¡no pierdas tu racha!')".

**Resultado (`specs/brand-definition.md`):**
- Documento con la coherencia explícita entre paleta/tipografía/tono
  señalada (las tres piezas apuntan a "calma", como pide `referencias/tipografia-y-tono.md`).
- Un ítem quedó en "Pendiente de confirmar": el usuario mencionó tener una idea de logo minimalista pero no lo definió — se marcó como pendiente en vez de inventar reglas de logo.

---

## Ejemplo 3 — Proyecto open-source

**Solicitud del usuario:**
> "Voy a publicar una librería de JavaScript. Quiero definir la identidad antes de armar la página de docs — algo con onda dev-tool, no corporativo."

**Decisiones:**
- Tono acordado: "desenfadado pero técnico, directo".
- Categoría de paleta: "Vibrante y audaz", coherente con el tono desenfadado común en dev-tools open-source.
- Primario: `#7928CA` (púrpura). Acento: `#FF0080` (rosa), como el degradado "Crepúsculo" invertido de `referencias paletas-color.md`.
- Verificación de contraste: texto `#FFFFFF` sobre `#7928CA` → 4.9:1, pasa AA para texto normal.
- Tipografía: sans-serif neutra para el cuerpo + monoespaciada como familia secundaria, reservada para nombres de comandos y snippets de código (no para párrafos largos, según la guía de `referencias/tipografia-y-tono.md`).
- Tono de voz: adjetivos "técnico, directo, con humor moderado". Ejemplo de mensaje de instalación exitosa: "Listo. `npm run dev` y ya estás andando." Qué evitar: "nunca sonar como una landing de ventas — este es un proyecto para developers".

**Resultado (`specs/brand-definition.md`):**
- Se dejó anotado que la paleta está pensada para verse bien tanto en la página de docs como embebida en el README de GitHub, para que quien implemente esa UI más adelante lo tenga en cuenta.

---

## Patrón general que se repite en los tres ejemplos

1. Elegir la categoría de paleta y la tipografía según el **tono** acordado con el usuario, no solo por la industria — un producto serio en una industria "divertida" (o viceversa) debe reflejar el tono, no el rubro.
2. Verificar contraste con `scripts/check_contrast.py` antes de dar por definitiva cualquier combinación, y dejar el resultado exacto documentado, no solo "verificado".
3. Definir el tono de voz con adjetivos **y** ejemplos de copy — nunca solo la lista de adjetivos.
4. Revisar que paleta, tipografía y tono cuenten la misma historia antes de cerrar el documento.
5. Marcar explícitamente cualquier decisión que el usuario no haya confirmado, en vez de completarla con un supuesto.
6. El entregable es el documento de reglas y sus tokens — no una imagen. La generación de assets finales es un paso posterior y separado que consume este resultado.
