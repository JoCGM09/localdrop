---
name: seguridad-buenas-practicas
description: "Aplica, revisa y corrige buenas prácticas de seguridad en código, scripts, configuraciones y skills de agentes antes de entregarlos o distribuirlos. Úsala siempre que se genere código nuevo, se empaquete una skill para distribución, se revise un pull request, o el usuario pida 'revisar la seguridad de esto', 'auditar este código', '¿esto es seguro?', 'corrige los problemas de seguridad', o cuando cualquier contenido incluya credenciales, tokens, llamadas de red, ejecución de comandos del sistema, o manejo de datos de usuario. Es un complemento obligatorio de flujos como generador-de-skills, ya que nada generado automáticamente debe distribuirse sin pasar por este checklist."
---

# Seguridad: buenas prácticas

## Alcance y postura

Esta skill es **defensiva**: revisa, detecta y corrige. No sirve para diseñar ataques, evadir controles de acceso, ni analizar objetivos ajenos sin autorización. Si una solicitud pide explotar una vulnerabilidad en un sistema de terceros, redirige a un proceso de divulgación responsable en vez de producir el exploit.

El objetivo es que, al terminar de usar esta skill, el artefacto revisado (código, script, skill, configuración) sea seguro por defecto para quien lo ejecute o instale, sin sorpresas ocultas.

## Cuándo se activa

- Antes de empaquetar o distribuir cualquier skill (propia o generada automáticamente).
- Al escribir o revisar código que maneje secretos, entradas de usuario, red o sistema de archivos.
- Cuando el usuario pide explícitamente una auditoría o revisión de seguridad.
- Como paso final de cualquier flujo de generación de código, incluso si nadie lo pide explícitamente: es responsabilidad del agente aplicarlo de forma proactiva.

## Flujo de trabajo

1. **Inventariar la superficie.** Antes de revisar línea por línea, identifica: ¿hay secretos o credenciales en el código? ¿hay llamadas de red? ¿hay ejecución de comandos del sistema (`subprocess`, `eval`, `exec`, `os.system`)? ¿hay entrada de usuario que llega a una consulta, un comando o una plantilla sin sanear? ¿hay dependencias externas nuevas?

2. **Aplicar el checklist** de `referencias/checklist-seguridad.md`, categoría por categoría. No saltar categorías solo porque "el código es simple" — los checklists cortos son exactamente donde se cuelan los descuidos.

3. **Consultar el catálogo de antipatrones** en `referencias/antipatrones-comunes.md` cuando encuentres algo que "se ve raro" pero no sabes nombrar. La mayoría de los problemas reales caen en un puñado de categorías conocidas.

4. **Ejecutar el escaneo estático básico** cuando el artefacto sea código o una carpeta de skill:

   ```bash
   python scripts/escaneo_basico.py <ruta-del-proyecto-o-skill>
   ```

   Esto detecta patrones de alto riesgo comunes (ver detalle en el script) de forma determinística, como red de seguridad adicional a la revisión manual — no la reemplaza.

5. **Corregir, no solo señalar.** Cuando encuentres un problema, propone y aplica la corrección concreta (por ejemplo: mover una clave hardcodeada a una variable de entorno, parametrizar una consulta en vez de concatenar strings, añadir validación de entrada). Si la corrección implica una decisión de producto (por ejemplo, qué política de expiración de tokens usar), señala la decisión al usuario en vez de asumirla.

6. **Reportar con severidad.** Al final de la revisión, resume los hallazgos con tres niveles:
   - **Crítico**: credenciales expuestas, inyección de comandos/SQL, ejecución de código arbitrario desde entrada no confiable, permisos excesivos. Bloquea la entrega hasta corregirse.
   - **Importante**: falta de validación de entrada, manejo de errores que filtra información sensible, dependencias sin fijar versión, ausencia de límites de tasa en operaciones costosas.
   - **Menor / mejora**: falta de comentarios de seguridad, logging insuficiente para auditoría, nombres de variables que ocultan el propósito sensible del dato.

## Principios que rigen toda corrección

- **Nunca reintroducir un secreto en texto plano.** Si encuentras uno, reemplázalo por una referencia a variable de entorno o gestor de secretos, y confirma que no quede también en el historial de ejemplos, logs o documentación.
- **Preferir el rechazo explícito sobre el manejo silencioso de errores.** Un `except: pass` alrededor de una operación sensible es un antipatrón, no una simplificación.
- **Validar en el borde, no confiar en el llamador.** Cualquier función que reciba datos de fuera del proceso (red, archivo, argumento de línea de comandos, entrada de otro agente) debe validar antes de usar ese dato en una operación sensible.
- **El principio de menor privilegio también aplica a skills.** Un script que solo necesita leer un archivo no debería tener permisos de escritura ni acceso de red "por si acaso".
- **La documentación es parte de la superficie de seguridad.** Un `README` con una clave de API real de ejemplo, aunque sea de una cuenta de prueba, es un hallazgo, no un detalle cosmético.

## Casos específicos para skills de el Agente

Cuando el artefacto a revisar es una el Agente Skill (como las generadas por `generador-de-skills`):

- Revisa que ningún script incluido ejecute comandos arbitrarios construidos a partir de la entrada del usuario sin sanear.
- Revisa que los ejemplos de código en `SKILL.md` o en `referencias/` no contengan tokens, claves o cookies reales copiadas de una sesión de recon o de documentación.
- Revisa que las instrucciones de la skill no le pidan a el Agente que omita advertencias de seguridad, ejecute código sin mostrarlo al usuario, o envíe datos a un endpoint no documentado.
- Si la skill incluye un `scripts/` con dependencias de terceros, confirma que estén ancladas a una versión y que provengan de un origen verificable.

## Referencias

- `referencias/checklist-seguridad.md` — checklist detallado por categoría (secretos, inyección, autenticación, dependencias, manejo de errores, superficie de red).
- `referencias/antipatrones-comunes.md` — catálogo de antipatrones frecuentes con su corrección recomendada.

## Scripts

- `scripts/escaneo_basico.py` — escaneo estático local (sin red) que busca patrones de alto riesgo: posibles secretos hardcodeados, usos de `eval`/`exec`/`os.system`, concatenación de SQL, `except` vacíos, y banderas de shell peligrosas. Es una ayuda determinística, no un sustituto de la revisión humana ni de herramientas SAST completas.
