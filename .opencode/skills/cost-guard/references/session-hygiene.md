# Higiene de sesiones

## Por qué importa
Cada mensaje previo en una sesión larga es contexto que se vuelve a pagar en cada turno siguiente. Una sesión que arrastra 5 features anteriores paga por ese historial en cada nueva pregunta, aunque el feature actual no tenga nada que ver con los anteriores.

## Cuándo cerrar/abrir sesión nueva

- Al terminar un feature y empezar otro que no depende del anterior.
- Cuando la sesión actual ya resolvió su objetivo original y lo que sigue es un tema distinto (pasaste de "implementar reservas" a "arreglar el pipeline de CI").
- Cuando notas que estás re-explicando contexto que ya deberías tener persistido en `AGENTS.md`/`specs/` en vez de en el chat.

## Qué SÍ persistir en AGENTS.md / specs/

- Convenciones del proyecto (estilo de commits, estructura de carpetas, naming).
- Decisiones de arquitectura ya tomadas y su razón (para no redebatirlas cada sesión).
- Requisitos de negocio confirmados por el usuario (lo que en `test-strategy` se referencia como `requirements.md`/
  `validation.md`).
- Deuda técnica conocida y su prioridad.
- Credenciales/config de entorno (rutas, no secretos) necesarias para correr el proyecto localmente.

## Qué NO vale la pena persistir

- Detalles de implementación que cambian seguido (se desactualizan rápido y generan confusión si no se mantienen).
- El razonamiento paso a paso de cómo se llegó a una decisión, solo la decisión final y su justificación en 1-2 líneas.
- Contenido que ya vive en el código mismo y es fácil de grepear (no dupliques la fuente de verdad).

## Formato sugerido

Usa `assets/AGENTS.md.template` como punto de partida y `assets/specs/feature-template.md` por cada feature nuevo. Mantén
`AGENTS.md` corto (es contexto que se carga siempre); mueve el detalle extenso a `specs/<feature>.md` y referencia desde `AGENTS.md` solo el puntero.

## Señal de que una sesión se volvió cara

- Empiezas a pegar resúmenes de "lo que hicimos antes" para que el modelo recuerde.
- El usuario tiene que repetir una decisión que ya había tomado hace varios turnos.
- La sesión mezcla temas sin relación (feature A, luego bug B, luego feature C) sin que compartan contexto real.

Cualquiera de estas es señal de cerrar sesión, volcar lo importante a `AGENTS.md`/`specs/`, y arrancar limpio.
