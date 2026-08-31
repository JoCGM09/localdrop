# Árbol de decisión: qué agente/modelo usar

## Paso 1 — Clasifica la tarea

```
¿La tarea solo lee/resume/planea (sin escribir código final)?
├── Sí → agente `plan` o subagente de revisión, modelo barato
└── No, escribe/modifica código o archivos
    ├── ¿Es mecánica y de scope acotado?
    │     (tests, docs, refactor chico, renombrado, formato,
    │      traducción de strings, generar boilerplate repetitivo)
    │   ├── Sí → subagente especializado, modelo barato
    │   └── No, requiere juicio no trivial
    │         (arquitectura, seguridad, specs, decisiones de
    │          producto, cambios que tocan compliance/legal)
    │       └── → agente `build`, `spec-writer` o
    │             `security-reviewer`, modelo capaz
```

## Ejemplos resueltos

| Tarea | Clasificación | Agente |
|---|---|---|
| "Lee el módulo de pagos y dime cómo está estructurado" | lectura/planeación | `plan` |
| "Agrega tests unitarios para `calcular_total`" | mecánica acotada | subagente `test-writer`, modelo barato |
| "Renombra `getUser` a `fetchUser` en todo el repo" | mecánica acotada | subagente barato + `bash`/`grep` |
| "Diseña el esquema de la tabla `appointments`" | arquitectura | `build`/`spec-writer`, modelo capaz |
| "Revisa si este endpoint cumple con retención de datos" | compliance/seguridad | `security-reviewer`, modelo capaz |
| "Escribe el CHANGELOG de este release" | mecánica acotada | subagente barato |
| "Decide si migramos de REST a GraphQL" | arquitectura | modelo capaz + probablemente confirmar con el usuario, no decidir solo |

## Casos límite

**Un subagente barato se traba o produce algo incorrecto en una tarea "mecánica".**
No escales automáticamente el modelo para el resto del proyecto.
Reintenta una vez dándole más contexto puntual (por ejemplo, el `requirements.md` relevante). Si vuelve a fallar, esa tarea específica sube a modelo capaz, pero el resto de tareas similares se mantienen en el barato — no fue el tipo de tarea el que falló, fue esa instancia.

**No estoy seguro si algo es "arquitectura" o "mecánico".**
Pregúntate: si el subagente barato se equivoca, ¿el costo de corregirlo después es alto (tocar compliance, seguridad, contratos con el usuario) o bajo (un test mal escrito que se detecta al correrlo)? Costo alto → modelo capaz. Costo bajo → modelo barato, aunque te equivoques a veces es más barato que subir todo por default.

**La tarea mezcla ambas cosas** (por ejemplo, "implementa el endpoint Y escribe tests").
Divide: la implementación del endpoint puede requerir juicio (modelo capaz si toca lógica de negocio nueva), los tests son mecánicos una vez que el endpoint existe (modelo barato, subagente `test-writer`, ver la skill `test-strategy`).

## Regla general

Por default, asume que la tarea es más barata de lo que parece. Delega a subagente barato primero; solo sube de modelo cuando haya evidencia concreta de que se necesita más capacidad (ambigüedad real, falla repetida, o la categoría "arquitectura seguridad/specs" de la tabla de arriba).
