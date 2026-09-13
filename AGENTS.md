# AGENTS.md — Reglas del proyecto

Este archivo se carga SIEMPRE (ver `instructions` en opencode.json).

<!-- Completa esto con todos los detalles de tu proyecto con conocimiento general compartido -->

Necesito construir una aplicación que solucione el siguiente problema: Necesito copiar y pegar información rápidamente entre mi celular y mi computadora, la información puede ser mucho texto, el clipboard, una imagen o hasta un archivo pequeño como .pdf, .docx u otro, pero necesito poder hacerlo de manera controlada y segura dentro de un entorno compartido entre mi celular y mi laptop, podría ser la red local de mi casa. Necesito una transferencia de información rápida y segura a través de un código de seguridad que se generará cuando elija la opción de enviar y luego deba colocarlo del otro lado con la opción de recibir.

Todo lo que pongas aquí no hace falta repetirlo en cada prompt: es la forma más barata de "entrenar" al agente para tu proyecto. Mantenlo corto (&lt; 1 página) — cada línea de más se paga en tokens en cada turno, de cada sesión, para siempre.

## Flujo de trabajo obligatorio (SDD)

1. No se escribe código sin un `plan.md` aprobado en `specs/<fecha>-<feature>/`.
2. Toda feature nueva empieza en una rama nueva desde `master`.
3. Antes de mergear: `test-writer` corrió y los tests pasan, `security-reviewer` no dejó hallazgos "high/critical" sin resolver.
4. Si el agente no está seguro de un requisito, pregunta — no asume.

## Convenciones técnicas

<!-- Completa esto una vez con tu stack real; ver specs/tech-stack.md -->

- Lenguaje / framework: Typescript, Astro, backend con go y chadcdn.
- Estilo de commits: Conventional Commits (`feat:`, `fix:`, `chore:`...)
- Gestor de paquetes: a tu elección
- Cómo correr tests localmente: a tu elección
- Cómo correr el linter: a tu elección

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

