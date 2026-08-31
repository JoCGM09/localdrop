# Plantilla de test unitario — Arrange / Act / Assert
#
# Copia esta estructura por cada regla de negocio o validación pura
# que quieras cubrir. Un test unitario NO debe tocar BD, red, ni
# filesystem real.
#
# Si el test que estás por escribir cubre un requisito de
# requirements.md, agrega el comentario "covers: REQ-XX" para que
# scripts/check_requirements_coverage.py lo detecte.

def test_nombre_descriptivo_del_caso():  # covers: REQ-XX
    # Arrange: prepara los datos/inputs mínimos necesarios
    entrada = ...

    # Act: ejecuta la unidad bajo prueba
    resultado = funcion_a_probar(entrada)

    # Assert: verifica el resultado esperado según requirements.md
    assert resultado == esperado


def test_caso_limite():  # covers: REQ-XX
    # Casos límite típicos a considerar: vacío, cero, valor máximo,
    # el primer/último elemento, tipo inesperado si aplica.
    ...