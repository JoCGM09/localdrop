# Patrones de salida

Usa estos patrones cuando las skills necesiten producir resultados consistentes y de alta calidad.

## Patrón de plantilla

Proporciona plantillas para el formato de salida. Ajusta el nivel de rigidez a lo que necesites.

**Para requisitos estrictos (como respuestas de API o formatos de datos):**

```markdown
## Estructura del informe

SIEMPRE usa esta estructura de plantilla exacta:

# [Título del análisis]

## Resumen ejecutivo
[Panorama de un párrafo con los hallazgos clave]

## Hallazgos clave
- Hallazgo 1 con datos de respaldo
- Hallazgo 2 con datos de respaldo
- Hallazgo 3 con datos de respaldo

## Recomendaciones
1. Recomendación específica y accionable
2. Recomendación específica y accionable
```

**Para orientación flexible (cuando la adaptación es útil):**

```markdown
## Estructura del informe

Aquí hay un formato por defecto razonable, pero usa tu criterio:

# [Título del análisis]

## Resumen ejecutivo
[Panorama general]

## Hallazgos clave
[Adapta las secciones según lo que descubras]

## Recomendaciones
[Ajusta al contexto específico]

Ajusta las secciones según lo requiera el tipo de análisis específico.
```

## Patrón de ejemplos

Para skills donde la calidad de la salida depende de ver ejemplos, proporciona pares de entrada/salida:

```markdown
## Formato de mensaje de commit

Genera mensajes de commit siguiendo estos ejemplos:

**Ejemplo 1:**
Entrada: Se agregó autenticación de usuario con tokens JWT
Salida:
```
feat(auth): implementar autenticación basada en JWT

Agrega endpoint de login y middleware de validación de tokens
```

**Ejemplo 2:**
Entrada: Se corrigió un error donde las fechas se mostraban incorrectamente en los reportes
Salida:
```
fix(reports): corregir formato de fecha en la conversión de zona horaria

Usar timestamps UTC de forma consistente en la generación de reportes
```

Sigue este estilo: tipo(alcance): descripción breve, luego explicación detallada.
```

Los ejemplos ayudan a el Agente a entender el estilo y nivel de detalle deseado con más claridad que las descripciones por sí solas.
