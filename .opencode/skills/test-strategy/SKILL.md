---
name: test-strategy
description: Cómo decidir qué testear, con qué nivel (unitario/integración/e2e) y con qué formato de reporte, para no gastar tokens ni tiempo probando de más. Úsala siempre que se escriban o revisen tests — desde el subagente test-writer, el comando /test, o cualquier tarea que mencione "agregar tests", "cobertura", "verificar que funciona", aunque no se use la palabra "test" explícitamente.
---

# Estrategia de testing costo-eficiente

Antes de escribir un solo test, recorre este checklist. El detalle extenso (ejemplos de código por nivel, checklist de qué NO cubrir, formato de reporte) vive en `references/` y `assets/templates/` — esta skill te dice cuándo ir a cada uno.

## 1. Pirámide simplificada (en este orden de prioridad)

1. **Unitario** — lógica pura, reglas de negocio, validaciones. Rápido, barato, escribir muchos está bien.
2. **Integración** — lo que toca BD, APIs externas, o varios módulos juntos. Escribe los necesarios para cubrir los flujos de `requirements.md`, no todos los combos posibles.
3. **E2E** — solo los 2-3 flujos críticos de negocio ("un paciente puede reservar una cita"). Son caros de mantener y correr: úsalos con moderación.

Ejemplos concretos de test en cada nivel (con código) están en `references/pyramide-detallada.md`. Plantillas listas para copiar están en `assets/templates/`.

## 2. Regla para no sobre-testear

Antes de escribir un test, pregúntate: "¿este test verifica algo que está en `requirements.md`/`validation.md`, o algo que el usuario realmente pidió?" Si la respuesta es no, no lo escribas — es gasto de tokens sin valor.

Casos típicos de sobre-testeo a evitar (detalle en `references/checklist-no-testear.md`):
- Getters/setters triviales sin lógica.
- Combinatoria exhaustiva de inputs cuando un par de casos límite ya cubre la regla de negocio.
- Tests que solo verifican que un framework/librería externa hace lo que su propia documentación ya garantiza.
- Duplicar el mismo caso en unitario e integración sin que aporte algo distinto.

## 3. Qué NO cubrir con tests generados por IA sin revisión humana

- Reglas de negocio ambiguas o no confirmadas por el usuario.
- Casos que dependen de compliance/legal (retención de datos) — estos deben validarse con el stakeholder, no asumirse.

Si detectas alguno de estos casos mientras escribes tests, no lo inventes ni lo asumas: márcalo explícitamente en el reporte final (sección "Pendiente de confirmar con el usuario") en vez de escribir un test sobre un supuesto no confirmado. Usa `assets/templates pendiente-confirmacion.md` como formato para listarlos.

## 4. Antes de escribir: revisa qué ya existe

- Corre `scripts/check_requirements_coverage.py` contra `requirements.md`/`validation.md` del feature para ver qué requisitos ya tienen test asociado y cuáles no — evita re-testear lo cubierto y te da la lista real de gaps.
- Si el feature no tiene `requirements.md`, no inventes requisitos: pregunta al usuario o usa solo lo que pidió explícitamente en el chat spec.

## 5. Formato de salida esperado

Al terminar, reporta usando `assets/templates/reporte-final.md` como base:
- Cuántos tests se agregaron, por nivel (unitario/integración/e2e).
- Cuántos pasan, cuántos fallan.
- Por cada fallo: si es un bug real o un test mal escrito.
- Requisitos de `requirements.md` sin test asociado (si los hay).
- Casos marcados como "pendiente de confirmar" (sección 3).

`scripts/generate_report.py` arma este reporte automáticamente a partir del output del runner de tests (pytest/jest/etc.) más el resultado de `check_requirements_coverage.py`.

## Referencias y recursos

- `references/pyramide-detallada.md` — ejemplos de código real por nivel (unitario/integración/e2e) y cuándo cruzar de un nivel a otro.
- `references/checklist-no-testear.md` — lista extendida de patrones de sobre-testeo y de casos que requieren confirmación humana antes de testear.
- `scripts/check_requirements_coverage.py` — compara `requirements.md`/`validation.md` contra los tests existentes y
  la lista qué requisitos no tienen test.
- `scripts/generate_report.py` — genera el reporte final en el formato de `assets/templates/reporte-final.md` a partir del output
  del test runner.
- `assets/templates/unit-test-template.*` — esqueleto de test unitario con estructura Arrange-Act-Assert.
- `assets/templates/integration-test-template.*` — esqueleto de test de integración con setup/teardown de dependencias externas.
- `assets/templates/e2e-test-template.*` — esqueleto de test E2E para un flujo crítico de negocio.
- `assets/templates/reporte-final.md` — formato del reporte de cierre.
- `assets/templates/pendiente-confirmacion.md` — formato para listar casos ambiguos/compliance que no se testearon sin validación.
