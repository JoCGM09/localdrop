---
description: Arranca una feature nueva - siguiente fase del roadmap, branch, y specs (requirements/plan/validation)
agent: spec-writer
---

1. Lee `specs/roadmap.md` y ubica la siguiente fase no implementada.
2. RUN `git checkout -b feature/$ARGUMENTS` (si el usuario no dio nombre, propone uno corto en kebab-case basado en la fase).
3. Crea el directorio `specs/YYYY-MM-DD-<nombre-feature>/` (usa la fecha de hoy).
4. Dentro, utiliza las plantillas de `specs/_template-feature` crea:
   - `requirements.md` — alcance, decisiones tomadas, contexto y lo que está explícitamente fuera de alcance.
   - `plan.md` — grupos de tareas numerados, cada uno chico y verificable por separado.
   - `validation.md` — cómo se sabe que la implementación está lista para mergear (criterios concretos).

Refiérete a `specs/mission.md` y `specs/tech-stack.md` para mantener consistencia.

Si los requerimientos o el plan específico para esta feature lo requiere, utiliza la skill `generador-de-skills` para añadir skills que se utilizarán posteriormente en la etapa de implementación. Lee atentamente el Readme.md correspondiente a la skill primero.

Inicialmente, utilizar la skill `definicion-de-marca` para plantear las bases del diseño, visuales, UX/UI e interfaces de la aplicación.

IMPORTANTE: antes de escribir CUALQUIER archivo, usa la herramienta de preguntas agrupando las preguntas de los tres documentos en una sola tanda. No escribas a disco sin haber confirmado con el usuario.
