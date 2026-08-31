---
name: generador-de-skills
description: Genera automáticamente Skills completas (SKILL.md + referencias + scripts) a partir de documentación técnica, investigando esa documentación directamente con las herramientas propias de del agente (búsqueda y lectura web) — sin configuraciones externas de API keys, MCPs u otros. Usa esta skill siempre que el usuario valide la necesidad de crear la skill identificada por el agente spec-writer, quiera scaffolding rápido de una skill nueva, o mencione documentación de una herramienta/API que quiere que se sepa usar de forma repetible. También cúbrela cuando el usuario pida "validar esta skill" o "empaquetar esta skill en un .skill".
---

# Generador de Skills

## Qué hace

Apunta esta skill a cualquier documentación (una URL, varias páginas, o texto de documentación ya pegado en el chat) y obtén una Skill funcional: un `SKILL.md` con frontmatter correcto, una carpeta `referencias/` con los patrones extraídos, y opcionalmente scripts de apoyo. El resultado es un punto de partida sólido, no un producto terminado — siempre debe revisarse con criterio antes de usarse en producción.

## Requisitos previos

Ninguno.

Si el usuario ya tiene documentación pegada en el chat o adjunta como archivo, úsala directamente y salta la investigación web — si es un link web debes realizar la búsqueda en internet solo de lo que ya te dieron.

## Flujo de trabajo completo

1. **Aclarar el alcance.** Antes de investigar nada, pregunta (o infiere del contexto si ya está claro):
   - ¿Qué URL(s) o producto de documentación es la fuente?
   - ¿Qué casos de uso concretos debe cubrir la skill? (no "todo Stripe", sino "webhooks de Stripe" o "autenticación de Stripe")
   - ¿Hay flujos de trabajo específicos que el usuario ya sabe que necesita?

2. **Investigar la documentación con tus propias herramientas.** Usa `web_search` para ubicar las páginas relevantes (empieza por la página índice de la documentación, luego busca las secciones específicas: autenticación, endpoints, SDKs, guías) y `web_fetch` para leer el contenido completo de cada página que encuentres. Sigue el checklist de `referencias/plantilla-prompt-extraccion.md` — enumera exactamente qué información no debe faltar (identidad de la skill, workflows, endpoints, autenticación, ejemplos de código, gotchas). Escala la cantidad de búsquedas/fetches a la complejidad del servicio: una API pequeña con pocas operaciones puede resolverse en 3-5 fetches; una plataforma con muchos productos (como Stripe completo) puede necesitar 10-20.
   - **Respeta el copyright mientras investigas**: nunca copies párrafos completos de la documentación fuente al `SKILL.md` o a las referencias. Reescribe con tus propias palabras, usa como máximo fragmentos cortos entrecomillados cuando la redacción exacta importe (como el nombre exacto de un header), y prioriza siempre transcribir código (que no tiene el mismo problema de derechos de autor sobre prosa) antes que prosa explicativa larga.
   - Si necesitas organizar lo que vas encontrando antes de escribir el `SKILL.md`, puedes armar un archivo intermedio siguiendo el esquema de `referencias/esquema-extraccion.json` como checklist de campos — esto es opcional, no un paso obligatorio ni un formato que deba "devolver" ningún agente externo.

3. **Inicializar la estructura de carpetas.**

   ```bash
   python scripts/init_skill.py <nombre-de-la-skill> --path <ruta-destino>
   ```

   Esto crea `SKILL.md` (plantilla), `scripts/`, `referencias/` y `assets/` con placeholders.

4. **Rellenar el SKILL.md real** con lo que investigaste:
   - `name`: hyphen-case, minúsculas, sin espacios (`stripe-webhooks`, no `Stripe Webhooks`).
   - `description`: **debe decir cuándo activarse**, no solo qué hace la herramienta. Sé "insistente": enumera frases y escenarios de usuario explícitos. Menos de 1024 caracteres, sin `<` ni `>`.
   - El cuerpo del `SKILL.md` sigue uno de los patrones de estructura (por flujo, por tarea, por referencia, por capacidades) — ver `scripts/init_skill.py` para la guía completa de las 4 variantes.
   - Workflows en modo imperativo ("Ejecuta X", "Configura Y"), nunca en segunda persona ("deberías ejecutar X").
   - Ejemplos de código completos y ejecutables, sin placeholders tipo `TU_API_KEY` sin explicar de dónde sacarlo; usa variables de entorno para secretos. Ningún ejemplo debe incluir claves, tokens o credenciales reales aunque aparezcan de forma incidental en la fuente — reemplázalas siempre por `os.environ["API_KEY"]` o equivalente.

5. **Distribuir contenido extenso en `referencias/`.** Si el `SKILL.md` superaría ~500 líneas, mueve detalle de API, ejemplos avanzados y esquemas a archivos de referencia con nombres descriptivos, y enlázalos desde el `SKILL.md` explicando cuándo leer cada uno. Usa los patrones de `referencias/patrones-de-salida.md` y `referencias/patrones-de-flujo.md` para decidir el formato.

6. **Validar antes de entregar.**

   ```bash
   python scripts/quick_validate.py <ruta-de-la-skill>
   ```

   Corrige cualquier error de frontmatter (claves no permitidas, nombre fuera de hyphen-case, descripción con `<`/`>`, longitudes excedidas) antes de continuar.

7. **Pasar la skill por la revisión de seguridad.** Antes de empaquetar o entregar cualquier skill generada automáticamente, aplica el checklist de la skill `seguridad-buenas-practicas` de este mismo paquete (si está disponible) — una skill generada a partir de documentación externa puede arrastrar ejemplos de código con credenciales de muestra, `eval()` innecesario, o instrucciones que piden ejecutar comandos sin verificación. No omitas este paso solo porque el contenido "vino de la documentación oficial".

8. **Empaquetar para distribución (opcional).**

   ```bash
   python scripts/package_skill.py <ruta-de-la-skill-folder> [directorio-salida]
   ```

   Esto valida de nuevo automáticamente y produce un `.skill` (zip) listo para compartir o instalar.

## Cuando la documentación no está disponible en la web

Si el usuario pega texto de documentación directamente en el chat, o adjunta un PDF/archivo con la documentación, sáltate el paso 2 de investigación web: usa ese contenido tal cual como fuente, aplicando el mismo checklist de `referencias/plantilla-prompt-extraccion.md` para no dejar campos sin cubrir.

Si una página relevante no es accesible (requiere login, está detrás de un paywall, o `web_fetch` la rechaza), dilo explícitamente al usuario en vez de inventar contenido para rellenar el hueco. Es preferible una skill con un `referencias/pendiente.md` señalando qué falta, que una skill con detalles de API inventados.

## Calidad de la extracción

- **Nombre**: hyphen-case, minúsculas, sin espacios.
- **Descripción**: bajo 1024 caracteres, sin `<` ni `>`, y con disparadores concretos de activación.
- **Workflows**: pasos en modo imperativo, extraídos de lo que la documentación realmente describe, no inventados.
- **Ejemplos de código**: completos y ejecutables, sin placeholders sin contexto; usa variables de entorno para secretos.
- **Gotchas**: límites de tasa, comportamientos no obvios, casos borde — solo si están documentados, nunca inventados.
- **Atribución**: si el `SKILL.md` o una referencia cita una frase corta textual de la documentación fuente, que sea breve (menos de 15 palabras) y quede claro de dónde salió; el resto siempre parafraseado.

## Referencias

- `referencias/plantilla-prompt-extraccion.md` — checklist completo de qué información no debe faltar al investigar la documentación (identidad de la skill, workflows, endpoints, autenticación, ejemplos, gotchas), pensado para que el Agente lo siga por sí mismo mientras usa `web_search`/`web_fetch`.
- `referencias/esquema-extraccion.json` — esquema JSON opcional para organizar la información recolectada antes de escribir el `SKILL.md`; también sirve como checklist de qué campos no deben faltar.
- `referencias/patrones-de-salida.md` — cómo estructurar la salida de una skill (plantillas estrictas vs. flexibles, patrón de ejemplos).
- `referencias/patrones-de-flujo.md` — cómo describir flujos de trabajo secuenciales y condicionales dentro de un `SKILL.md`.

## Scripts

- `scripts/init_skill.py` — crea el andamiaje de una skill nueva (`SKILL.md` + `scripts/` + `referencias/` + `assets/`).
- `scripts/quick_validate.py` — valida el frontmatter YAML de un `SKILL.md` contra las reglas del formato.
- `scripts/package_skill.py` — valida y empaqueta una carpeta de skill en un archivo `.skill` (zip) distribuible.
