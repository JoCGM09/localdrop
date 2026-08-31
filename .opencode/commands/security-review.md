---
description: Revisión de seguridad de solo lectura sobre el diff actual antes de mergear
agent: security-reviewer
---

Usa la skill `seguridad-buenas-practicas` contra `git diff main...HEAD` (o el
rango que indique el usuario). Reporta hallazgos con severidad. No
edites nada — si encuentras algo, repórtalo para que el agente `build`
lo arregle en un paso aparte.

$ARGUMENTS
