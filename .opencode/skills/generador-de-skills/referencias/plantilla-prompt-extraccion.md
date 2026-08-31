# Checklist de investigación para generar una Skill

Este documento es una guía para investigar documentación técnica y extraer de ella todo lo necesario para armar una Skill. No es
un prompt para un agente externo: es el checklist que el agente mismo sigue mientras usa `web_search` y `web_fetch` (o mientras lee documentación pegada/adjuntada por el usuario).

## Contexto

Las Skills son paquetes modulares que extienden las capacidades de los agentes con conocimiento especializado, flujos de
trabajo e integraciones de herramientas. Piénsalas como "guías de incorporación" (onboarding) para un dominio específico.

## Cómo investigar

1. Empieza por la página índice/overview de la documentación con `web_fetch` si el usuario dio una URL concreta, o con `web_search` si solo dio el nombre del servicio.
2. Desde ahí, identifica las secciones que necesitas cubrir (ver checklist abajo) y sigue los enlaces relevantes con `web_fetch` uno por uno: autenticación, referencia de API, guías de inicio rápido, SDKs por lenguaje.
3. Si la documentación es muy extensa, prioriza las secciones que
   correspondan a los casos de uso que el usuario pidió cubrir, no toda la documentación del producto.
4. Anota (mentalmente o en un borrador) qué campos del checklist siguen sin cubrir después de cada fetch, para no terminar la investigación con huecos que luego rellenes inventando.

## Checklist de extracción

### 1. Identidad de la skill (CRÍTICO)

- **nombre**: genera un nombre en hyphen-case (minúsculas, guiones). Ejemplos: `pdf-editor`, `stripe-webhooks`.
- **descripción**: DEBE incluir CUÁNDO usar la skill. Este es el mecanismo principal de activación. Incluye:
  - Qué hace la herramienta.
  - Escenarios de usuario específicos que deberían activarla.
  - Ejemplo: "Extracción y scraping web. Úsala cuando: (1) se solicite raspar sitios web, (2) extraer datos estructurados de URLs, (3) rastrear múltiples páginas, (4) convertir contenido web a markdown."
- **overview**: 1-2 frases que resuman la capacidad principal.
- **triggers**: lista de 5 a 10 frases de usuario que deberían activar esta skill.

### 2. Workflows (OBLIGATORIO)

Extrae TODAS las operaciones principales que soporta la herramienta. Para cada workflow:
- **nombre**: nombre de acción claro ("Scraping de página única", "Rastreo por lotes").
- **descripción**: qué logra.
- **pasos**: instrucciones paso a paso en forma IMPERATIVA ("Ejecuta X", "Configura Y", NO "Deberías ejecutar X").
- **ejemplo de código**: incluye código funcional cuando esté disponible en la documentación fuente.

### 3. Detalles de la API (si aplica)

Para cada endpoint extrae:
- Método (GET/POST/etc.)
- Ruta (path)
- Todos los parámetros con tipo, si son obligatorios, y descripción.
- Estructura de la respuesta.

### 4. Autenticación

- Tipo de método (API Key, Bearer Token, OAuth).
- Nombre del header.
- Nombre sugerido de variable de entorno.
- Instrucciones de configuración.

### 5. Ejemplos de código

Extrae ejemplos en TODOS los lenguajes disponibles en la documentación fuente: Python, TypeScript/JavaScript, cURL, otros SDKs. Para cada ejemplo:
- Código completo y ejecutable (sin placeholders tipo "TU_API_KEY" sin contexto).
- Reemplaza cualquier secreto de ejemplo por una variable de entorno.
- Muestra uso simple y uso avanzado si la documentación los distingue.

### 6. Buenas prácticas y gotchas

- Patrones comunes que funcionan bien.
- Anti-patrones a evitar.
- Límites de tasa (rate limits).
- Enfoques de manejo de errores.
- Casos borde.

### 7. Planificación de divulgación progresiva

Si el contenido superaría 500 líneas en el `SKILL.md`, identifica qué debería ir en archivos de referencia separados:
- Detalles de referencia de API.
- Ejemplos avanzados.
- Documentación de esquemas.
- Guías de resolución de problemas.

## Criterios de calidad

1. **Sé completo pero conciso** — incluye suficiente detalle para que el agente use la herramienta sin tener que releer la documentación original en cada uso.
2. **Usa forma imperativa** — "Ejecuta", "Configura", "Establece" (no "deberías ejecutar").
3. **Prioriza ejemplos prácticos** — código real sobre teoría.
4. **Incluye manejo de errores** — qué puede salir mal y cómo manejarlo, si la documentación lo cubre.
5. **Anota límites de tasa y costos** — si están documentados.
6. **No inventes nada.** Si la documentación no cubre algo del checklist, no lo completes con una suposición razonable — déjalo señalado como pendiente (ver más abajo) o pregúntale al usuario.

## Copyright al extraer

Mientras lees la documentación con `web_fetch`, no traslades párrafos completos al `SKILL.md` ni a las referencias. Reescribe cada sección con tus propias palabras; los únicos fragmentos que pueden citarse casi textuales son detalles técnicos exactos donde la redacción literal importa (el nombre exacto de un header, un mensaje de error específico), y siempre en fragmentos cortos. El código de ejemplo no tiene este problema del mismo modo que la prosa explicativa, así que prioriza transcribir bloques de código reales antes que resumir la prosa que los rodea.

## Uso opcional del esquema JSON

Si te resulta útil organizar lo investigado antes de redactar el `SKILL.md` — especialmente en documentaciones grandes con muchos
endpoints — puedes armar un archivo `extraccion.json` intermedio siguiendo la estructura de `referencias/esquema-extraccion.json`. Esto es un borrador de trabajo tuyo, no un formato de intercambio con ningún sistema externo; puedes saltarte este paso y escribir el `SKILL.md` directamente si el alcance es chico.

## Salida

El resultado final de este checklist es el propio `SKILL.md` (y sus referencias), no un JSON que se "devuelve". Antes de darlo por terminado, verifica:
- `descripción` tiene menos de 1024 caracteres.
- `nombre` tiene menos de 64 caracteres y está en hyphen-case.
- Todos los ejemplos de código son sintácticamente válidos.
- Los pasos son accionables y específicos, no genéricos.

## Nota de seguridad

Ningún ejemplo extraído debe incluir claves, tokens o credenciales reales, aunque aparezcan de forma incidental en la documentación fuente (por ejemplo, en una captura de pantalla o un log de ejemplo). Reemplázalos siempre por variables de entorno o placeholders explícitos como `os.environ["API_KEY"]`. Antes de aceptar la extracción, pásala por el checklist de la skill `seguridad-buenas-practicas` si está disponible en el mismo paquete.
