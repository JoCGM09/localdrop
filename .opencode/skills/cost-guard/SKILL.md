---
name: cost-guard
description: Disciplina de contexto, tokens y elección de modelo/agente a aplicar en CUALQUIER tarea de este proyecto (lectura de código, tests, refactors, specs, sesiones largas). Aplica siempre de forma proactiva, no solo cuando se invoca explícitamente — úsala incluso si el usuario no menciona "costo" o "tokens", cualquier vez que vayas a leer archivos, elegir un subagente/modelo, o decidir si abrir una sesión nueva.
---

# Disciplina de costo/tokens

Esta skill es un checklist operativo, no una lectura pasiva. Antes de ejecutar cualquier tarea del proyecto, recorre las secciones en orden y aplica la regla correspondiente. Los detalles largos (patrones de grep, árbol de decisión de modelo, plantillas) viven en `references/` y `assets/` para no inflar este archivo — ve ahí solo cuando la
sección te lo indique.

## 1. Contexto: lee lo mínimo necesario

- Prefiere `grep -n`/`glob` a leer archivos completos cuando solo necesitas ubicar algo, confirmar que existe, o extraer una función.
- No releas archivos que ya están en el contexto de esta sesión.
- No listes el repo completo (`ls -R`, `find .`) si sabes qué carpeta buscar. Usa rutas específicas.
- Si una tarea es puramente mecánica (renombrar, mover, formato, buscar-y-reemplazar), resuélvela con `bash`/`edit` directo en vez de razonar en texto sobre ella primero.
- **Señal de alerta**: si estás a punto de pegar un archivo de +300 líneas solo para que el modelo "tenga contexto", detente y usa `grep -n` para la sección relevante. Ver `references/context-discipline.md` para patrones concretos (buscar una función, un endpoint, un uso de una variable, etc.) y `scripts/context_budget_check.sh` para automatizar esta revisión antes de leer un archivo grande.

## 2. Modelo/agente correcto para la tarea

No todas las tareas requieren el modelo más capaz. Antes de delegar o ejecutar, clasifica la tarea:

| Tipo de tarea | Agente/modelo | Ejemplos |
|---|---|---|
| Lectura/planeación | `plan` o subagente de revisión (modelo barato, sin permiso de escribir) | leer specs, mapear el repo, resumir un PR |
| Mecánica y acotada | skill y/o agente especializado, modelo barato | tests, docs, refactors chicos, renombrados |
| Arquitectura/seguridad/specs | Agente `build`, `spec-writer`, `security-reviewer` (modelo capaz) | decisiones de diseño, revisión de compliance, specs nuevas |

Si tienes duda de a qué categoría pertenece una tarea, trátala como mecánica y delega a un subagente y/o skill con el modelo barato primero; solo escala a un modelo capaz si el subagente reporta ambigüedad. El árbol de decisión completo, con más ejemplos y casos límite, está en `references/model-routing.md`. `scripts/pick_model.py` da una sugerencia rápida por palabra clave si no estás seguro.

## 3. Sobre las sesiones:

- Cierra/empieza una sesión nueva por feature en vez de arrastrar el historial completo de features anteriores — el historial viejo se vuelve contexto pagado que ya no aporta.
- Usa `AGENTS.md` y `specs/` para todo lo que deba persistir entre sesiones, no el historial de chat. Plantillas listas para copiar: `assets/AGENTS.md.template` y `assets/specs/feature-template.md`.
- Antes de cerrar una sesión, vuelca a `specs/` o `AGENTS.md` cualquier decisión, convención o dato que la siguiente sesión necesitaría releer del chat. Ver `references/session-hygiene.md` para qué vale la pena persistir y qué no.

## 4. Checklist rápido antes de cada tarea

1. ¿Puedo resolver esto con `grep`/`glob` en vez de leer todo el archivo?
2. ¿Ya tengo esta info en contexto de esta sesión? Si sí, no la releas.
3. ¿Esta tarea es mecánica? → subagente barato. ¿Es de arquitectura/seguridad/specs? → modelo capaz.
4. ¿Esto debería vivir en `AGENTS.md`/`specs/` en vez de solo en el chat?
5. ¿Este es un feature nuevo? → considera sesión nueva en vez de seguir en la actual.

## Referencias y recursos

- `references/context-discipline.md` — patrones de grep/glob por caso de uso, cuándo SÍ vale la pena leer el archivo completo.
- `references/model-routing.md` — árbol de decisión completo con casos límite (¿qué hago si un subagente barato se traba?).
- `references/session-hygiene.md` — qué persistir en `AGENTS.md`/ `specs/`, qué formato usar, cuándo cerrar una sesión.
- `scripts/context_budget_check.sh` — antes de leer un archivo grande, imprime tamaño/líneas y sugiere si conviene grep en vez de `view` completo.
- `scripts/pick_model.py` — clasifica una descripción de tarea en texto libre y sugiere agente/modelo según la tabla de la sección 2.
- `assets/AGENTS.md.template` — plantilla para arrancar el archivo de contexto persistente de un proyecto nuevo.
- `assets/specs/feature-template.md` — plantilla de spec por feature, pensada para reemplazar el "pegar todo el historial" al abrir una sesión nueva.
