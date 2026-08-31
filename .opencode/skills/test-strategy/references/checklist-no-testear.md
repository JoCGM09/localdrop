# Checklist: qué NO testear (o no testear todavía)

## Patrones de sobre-testeo a evitar

- **Getters/setters triviales.** Si no hay lógica (validación, transformación, efecto secundario), no aporta un test dedicado.
- **Combinatoria exhaustiva sin justificación.** Si una función tiene 3 parámetros booleanos, no necesitas los 8 casos si 2-3 casos representativos ya cubren las reglas de negocio distintas. Prueba los casos límite (vacío, cero, máximo, el "camino feliz" y el camino de error más probable), no el producto cartesiano completo.
- **Reprobar lo que el framework ya garantiza.** No testees que `useState` de React actualiza el estado, o que el ORM guarda una fila — eso ya está cubierto por los tests del framework/librería. Testea tu lógica encima de eso.
- **Duplicar el mismo caso en dos niveles sin razón.** Si un test de integración ya prueba "reservar en horario ocupado falla" contra la BD real, no necesitas repetir exactamente el mismo caso como test unitario con un mock de BD, salvo que quieras aislar la regla pura de validación de horario (en cuyo caso el unitario prueba la regla, el de integración prueba que se aplica correctamente contra BD real — son cosas distintas, no una duplicación).
- **Tests que dependen de detalles de implementación frágiles** (por ejemplo contar cuántas veces se llamó una función interna) cuando lo que importa es el resultado observable. Esto genera tests que se rompen con refactors que no cambian el comportamiento — costo alto, valor bajo.

## Casos que requieren confirmación humana antes de testear (no asumir)

- **Reglas de negocio ambiguas.**: "¿qué pasa si el usuario cancela una cita a menos de 1 hora del horario?" Si no está en `requirements.md`/`validation.md` ni lo dijo el usuario explícitamente, no inventes el comportamiento esperado para escribirle un test — pregúntalo o márcalo como pendiente.
- **Compliance/legal.**: Tiempos de retención de datos, requisitos de auditoría, consentimiento. Estos no se asumen desde lo que "parece razonable" — deben validarse con el stakeholder (legal/compliance), y el test se escribe después de esa confirmación, citando la fuente de la regla.
- **Comportamiento de terceros no documentado.** Si una integración externa (pasarela de pago, API de un tercero) tiene un comportamiento no confirmado en su documentación oficial, no escribas un test que asuma ese comportamiento como contrato — es fácil que quede desactualizado o directamente mal.

## Cómo reportarlos

Cuando encuentres alguno de estos casos, no lo saltees en silencio: agrégalo a la sección "Pendiente de confirmar con el usuario" del reporte final, usando `assets/templates/pendiente-confirmacion.md` como formato. Esto evita que el gap quede invisible y que alguien asuma más adelante que "ya está testeado".
