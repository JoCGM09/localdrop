# Pirámide de testing — ejemplos por nivel

Estos ejemplos son ilustrativos (sintaxis tipo Jest/pytest, adapta al stack real del proyecto). El punto no es la sintaxis sino qué se
prueba en cada nivel y por qué.

## Nivel 1 — Unitario

Prueba lógica pura, sin tocar BD, red, ni filesystem. Debe correr en milisegundos y no requerir mocks pesados.

```python
# calcular_total.py
def calcular_total(items, descuento_pct=0):
    subtotal = sum(item["precio"] * item["cantidad"] for item in items)
    return subtotal * (1 - descuento_pct / 100)

# test_calcular_total.py
def test_calcula_total_sin_descuento():
    items = [{"precio": 10, "cantidad": 2}]
    assert calcular_total(items) == 20

def test_aplica_descuento_porcentual():
    items = [{"precio": 100, "cantidad": 1}]
    assert calcular_total(items, descuento_pct=10) == 90

def test_lista_vacia_da_total_cero():
    assert calcular_total([]) == 0
```

Escribir "muchos" tests unitarios está bien porque son baratos. Pero "muchos" significa cubrir las reglas de negocio y casos límite reales (cero, negativo si aplica, descuento 100%), no cada combinación numérica posible sin razón.

## Nivel 2 — Integración

Prueba que varios módulos, o el código contra una BD/API real (o un stub fiel), funcionan juntos. Requiere setup/teardown.

```python
# test_reservas_integracion.py
def test_reservar_cita_guarda_en_bd(db_session, paciente_factory):
    paciente = paciente_factory()
    resultado = reservar_cita(
        db_session, paciente_id=paciente.id, fecha="2026-09-10T10:00"
    )
    cita_en_bd = db_session.query(Cita).filter_by(id=resultado.id).first()
    assert cita_en_bd is not None
    assert cita_en_bd.paciente_id == paciente.id

def test_reservar_cita_en_horario_ocupado_falla(db_session, cita_existente):
    with pytest.raises(HorarioOcupadoError):
        reservar_cita(
            db_session,
            paciente_id=cita_existente.paciente_id,
            fecha=cita_existente.fecha,
        )
```

Escribe los que cubran los flujos descritos en `requirements.md` ("no se puede doble-reservar un horario"), no todos los combos de inputs posibles cruzados con todos los estados de BD posibles.

## Nivel 3 — E2E

Solo los 2-3 flujos de negocio verdaderamente críticos. Son caros de mantener (se rompen con cambios de UI, son lentos) así que cada uno debe justificar su existencia por el impacto de negocio si falla en producción.

```python
# test_e2e_reserva_cita.py (con Playwright)
def test_paciente_puede_reservar_una_cita(page):
    page.goto("/login")
    page.fill("#email", "paciente@test.com")
    page.fill("#password", "****")
    page.click("text=Ingresar")

    page.click("text=Reservar cita")
    page.select_option("#especialidad", "Odontología")
    page.click("text=10:00 AM - 10 sep")
    page.click("text=Confirmar")

    assert page.locator("text=Cita confirmada").is_visible()
```

Antes de agregar un E2E nuevo, pregúntate: "¿qué otro nivel de test ya cubre esto de forma más barata?" Si un test de integración ya prueba que `reservar_cita` funciona contra la BD, el E2E solo debería justificarse por probar la cadena completa (UI → API → BD) en un flujo que el negocio marcó como crítico, no por duplicar la misma aserción.

## Cuándo subir de nivel

Sube de unitario a integración cuando la lógica que quieres probar depende genuinamente de una BD/API/otro módulo (no se puede aislar con datos en memoria sin perder el sentido del test). Sube de integración a E2E solo si el flujo completo (varias pantallas/ llamadas encadenadas) es en sí mismo parte de lo que el negocio considera crítico — no por "más cobertura es mejor".
