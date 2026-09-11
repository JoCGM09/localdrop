# Instrucciones de Proyecto (Claude Code)

# AGENTS.md — Reglas del proyecto

Este archivo se carga SIEMPRE (ver `instructions` en opencode.json).
<!-- Completa esto con todos los detalles de tu proyecto con conocimiento general compartido -->
Todo lo que pongas aquí no hace falta repetirlo en cada prompt: es la forma más barata de "entrenar" al agente para tu proyecto. Mantenlo corto (< 1 página) — cada línea de más se paga en tokens en cada turno, de cada sesión, para siempre.

## Flujo de trabajo obligatorio (SDD)

1. No se escribe código sin un `plan.md` aprobado en `specs/<fecha>-<feature>/`.
2. Toda feature nueva empieza en una rama nueva desde `master`.
3. Antes de mergear: `test-writer` corrió y los tests pasan, `security-reviewer` no dejó hallazgos "high/critical" sin resolver.
4. Si el agente no está seguro de un requisito, pregunta — no asume.

## Convenciones técnicas
<!-- Completa esto una vez con tu stack real; ver specs/tech-stack.md -->
- Lenguaje / framework:
- Estilo de commits: Conventional Commits (`feat:`, `fix:`, `chore:`...)
- Gestor de paquetes:
- Cómo correr tests localmente:
- Cómo correr el linter:

## Seguridad — no negociable
- Nunca hardcodear secrets, tokens o API keys. Usar variables de entorno.
- Toda entrada de usuario se valida y sanitiza antes de tocar la BD.
- Nunca loguear PII (DNI, contraseñas, etc) en texto plano.
- Cualquier endpoint que toque datos sensibles requiere autenticación y autorización explícita — nunca "por defecto abierto".

## Disciplina de costo/tokens
- No leas archivos completos si con `grep`/`glob` alcanza para ubicar lo que necesitas.
- No repitas contexto que ya está en este archivo o en `specs/`.
- Para tareas mecánicas (tests, docs, refactors chicos) usa el subagente correspondiente con modelo económico — no el agente principal.
- Si una tarea puede resolverse leyendo 1 archivo, no listes todo el repo primero.


---

## Roles y Subagentes

Este proyecto define roles específicos para distintas tareas de IA:
- **code-reviewer**: Revisión de calidad de código de solo lectura (no seguridad, eso es security-reviewer). Modelo barato. (Ver detalles en `.opencode/agent/code-reviewer.md`)
- **security-reviewer**: Revisión de seguridad de solo lectura antes de mergear. No edita nada — solo produce un reporte. (Ver detalles en `.opencode/agent/security-reviewer.md`)
- **spec-writer**: Redacta la constitución del proyecto y specs de feature (mission, tech-stack, roadmap, requirements, plan, validation). NUNCA escribe código de la app. (Ver detalles en `.opencode/agent/spec-writer.md`)
- **test-writer**: Escribe y corre tests (unitarios, integración, e2e básicos) a partir de requirements.md y validation.md de la feature actual. Modelo barato: es trabajo mecánico y acotado. (Ver detalles en `.opencode/agent/test-writer.md`)


---

## Skills Disponibles

Puedes consultar o invocar estas habilidades:
- **cost-guard**: (Instrucciones en `.opencode/skills/cost-guard/SKILL.md`)
- **definicion-de-marca**: (Instrucciones en `.opencode/skills/definicion-de-marca/SKILL.md`)
- **deploy-netlify**: (Instrucciones en `.opencode/skills/deploy-netlify/SKILL.md`)
- **generador-de-skills**: (Instrucciones en `.opencode/skills/generador-de-skills/SKILL.md`)
- **seguridad-buenas-practicas**: (Instrucciones en `.opencode/skills/seguridad-buenas-practicas/SKILL.md`)
- **test-strategy**: (Instrucciones en `.opencode/skills/test-strategy/SKILL.md`)


---

## Comandos del Proyecto

Este repositorio soporta los siguientes comandos que simulan el comportamiento nativo. Si el usuario te pide ejecutar alguno, aplica la lógica descrita:

### Comando: `/constitution`
- **Qué hace:** Crea/actualiza la constitución del proyecto (mission, tech-stack, roadmap) en specs/
- **Rol que debes asumir:** `spec-writer` (Lee `.opencode/agent/spec-writer.md`)
- **Lógica de ejecución:** Sigue paso a paso las instrucciones detalladas en `.opencode/commands/constitution.md`.

### Comando: `/feature`
- **Qué hace:** Arranca una feature nueva - siguiente fase del roadmap, branch, y specs (requirements/plan/validation)
- **Rol que debes asumir:** `spec-writer` (Lee `.opencode/agent/spec-writer.md`)
- **Lógica de ejecución:** Sigue paso a paso las instrucciones detalladas en `.opencode/commands/feature.md`.

### Comando: `/implement`
- **Qué hace:** Implementa el siguiente grupo de tareas pendiente del plan.md de la feature actual
- **Rol que debes asumir:** `build` (Lee `.opencode/agent/build.md`)
- **Lógica de ejecución:** Sigue paso a paso las instrucciones detalladas en `.opencode/commands/implement.md`.

### Comando: `/security-review`
- **Qué hace:** Revisión de seguridad de solo lectura sobre el diff actual antes de mergear
- **Rol que debes asumir:** `security-reviewer` (Lee `.opencode/agent/security-reviewer.md`)
- **Lógica de ejecución:** Sigue paso a paso las instrucciones detalladas en `.opencode/commands/security-review.md`.

### Comando: `/test`
- **Qué hace:** Genera y corre tests para la feature actual, según requirements.md/validation.md
- **Rol que debes asumir:** `test-writer` (Lee `.opencode/agent/test-writer.md`)
- **Lógica de ejecución:** Sigue paso a paso las instrucciones detalladas en `.opencode/commands/test.md`.



## Instrucciones para el Agente Claude
- Si el usuario menciona un comando que empiece con `/` (ej. `/constitution`), búscalo en la lista de comandos de arriba y ejecuta la lógica exacta de su archivo asociado. Asume automáticamente el rol requerido.