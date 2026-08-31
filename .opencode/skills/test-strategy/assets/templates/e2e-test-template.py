# Plantilla de test E2E
#
# ÚSALA SOLO para flujos de negocio verdaderamente críticos (2-3 por
# proyecto, no más). Antes de agregar uno nuevo, confirma con el
# checklist de la skill: ¿este flujo completo es en sí mismo crítico
# para el negocio, o ya está cubierto por un test de integración más
# barato?
#
# Ejemplo con Playwright (adapta al framework E2E del proyecto).

def test_flujo_critico_de_negocio(page):  # covers: REQ-XX (flujo crítico)
    # 1. Login / estado inicial
    page.goto("/login")
    page.fill("#email", "usuario@test.com")
    page.fill("#password", "****")
    page.click("text=Ingresar")

    # 2. Pasos del flujo crítico, uno por uno
    page.click("text=<acción principal>")
    ...

    # 3. Confirmación del resultado observable por el usuario
    assert page.locator("text=<confirmación esperada>").is_visible()

    # Nota: evita aserciones sobre detalles internos (llamadas a
    # funciones, estructura de la respuesta HTTP) — en E2E lo que
    # importa es lo que el usuario final observa.
