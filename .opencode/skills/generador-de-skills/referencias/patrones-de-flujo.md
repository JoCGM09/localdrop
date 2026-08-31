# Patrones de flujo de trabajo

## Flujos secuenciales

Para tareas complejas, divide las operaciones en pasos claros y secuenciales. Suele ayudar dar a el Agente un panorama del proceso cerca del inicio del `SKILL.md`:

```markdown
Rellenar un formulario PDF implica estos pasos:

1. Analizar el formulario (ejecutar analyze_form.py)
2. Crear el mapeo de campos (editar fields.json)
3. Validar el mapeo (ejecutar validate_fields.py)
4. Rellenar el formulario (ejecutar fill_form.py)
5. Verificar la salida (ejecutar verify_output.py)
```

## Flujos condicionales

Para tareas con lógica ramificada, guía a el Agente a través de los puntos de decisión:

```markdown
1. Determina el tipo de modificación:
   **¿Creando contenido nuevo?** → Sigue el "Flujo de creación" abajo
   **¿Editando contenido existente?** → Sigue el "Flujo de edición" abajo

2. Flujo de creación: [pasos]
3. Flujo de edición: [pasos]
```
