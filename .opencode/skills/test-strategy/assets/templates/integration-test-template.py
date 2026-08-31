# Plantilla de test de integración
#
# Usa esta estructura cuando el comportamiento a probar depende
# genuinamente de una BD, API externa, o la interacción entre varios
# módulos. Incluye setup/teardown explícito de esas dependencias.
#
# Antes de escribir este test, confirma que corresponde a un flujo
# descrito en requirements.md/validation.md — no generes combos
# extra "por las dudas".

import pytest


@pytest.fixture
def setup_dependencias():
    # Setup: crea la conexión/datos necesarios (BD de test, fixtures,
    # stub de API externa, etc.)
    recursos = ...
    yield recursos
    # Teardown: limpia lo creado para no afectar otros tests
    ...


def test_flujo_de_requirements(setup_dependencias):  # covers: REQ-XX
    # Arrange: usa los recursos del fixture
    ...

    # Act: ejecuta el flujo que cruza módulos/BD/API
    resultado = ejecutar_flujo(...)

    # Assert: verifica tanto el resultado como el efecto secundario
    # persistido (ej. la fila en BD, la llamada a la API externa)
    assert resultado == esperado
    # assert estado_en_bd == esperado
