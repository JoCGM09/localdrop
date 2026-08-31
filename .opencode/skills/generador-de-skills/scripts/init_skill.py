#!/usr/bin/env python3
"""
Inicializador de skills - Crea una nueva skill a partir de una plantilla

Uso:
    init_skill.py <nombre-skill> --path <ruta>

Ejemplos:
    init_skill.py mi-skill-nueva --path .opencode/skills/public
    init_skill.py mi-ayudante-api --path skills/private
"""

import sys
from pathlib import Path


PLANTILLA_SKILL = """---
name: {skill_name}
description: [TODO: explicación completa e informativa de qué hace la skill y cuándo usarla. Incluye CUÁNDO usar esta skill - escenarios específicos, tipos de archivo o tareas que la activan.]
---

# {skill_title}

## Resumen

[TODO: 1-2 frases explicando qué habilita esta skill]

## Estructurando esta skill

[TODO: Elige la estructura que mejor se adapte al propósito de esta skill. Patrones comunes:

**1. Basada en flujo de trabajo** (ideal para procesos secuenciales)
- Funciona bien cuando hay procedimientos claros paso a paso
- Ejemplo: skill de DOCX con "Árbol de decisión del flujo" → "Lectura" → "Creación" → "Edición"
- Estructura: ## Resumen → ## Árbol de decisión → ## Paso 1 → ## Paso 2...

**2. Basada en tareas** (ideal para colecciones de herramientas)
- Funciona bien cuando la skill ofrece distintas operaciones/capacidades
- Ejemplo: skill de PDF con "Inicio rápido" → "Combinar PDFs" → "Dividir PDFs" → "Extraer texto"
- Estructura: ## Resumen → ## Inicio rápido → ## Categoría de tarea 1 → ## Categoría de tarea 2...

**3. Referencia/Lineamientos** (ideal para estándares o especificaciones)
- Funciona bien para guías de marca, estándares de código o requisitos
- Ejemplo: estilo de marca con "Lineamientos de marca" → "Colores" → "Tipografía" → "Características"
- Estructura: ## Resumen → ## Lineamientos → ## Especificaciones → ## Uso...

**4. Basada en capacidades** (ideal para sistemas integrados)
- Funciona bien cuando la skill ofrece múltiples características interrelacionadas
- Ejemplo: gestión de producto con "Capacidades principales" → lista numerada de capacidades
- Estructura: ## Resumen → ## Capacidades principales → ### 1. Función → ### 2. Función...

Los patrones se pueden combinar según se necesite. La mayoría de skills combinan patrones (empezar basada en tareas y agregar flujo para operaciones complejas).

Elimina toda esta sección "Estructurando esta skill" cuando termines - es solo guía.]

## [TODO: Reemplaza con la primera sección principal según la estructura elegida]

[TODO: Agrega contenido aquí. Ver ejemplos en skills existentes:
- Ejemplos de código para skills técnicas
- Árboles de decisión para flujos complejos
- Ejemplos concretos con solicitudes de usuario realistas
- Referencias a scripts/plantillas/referencias según se necesite]

## Recursos

Esta skill incluye directorios de recursos de ejemplo que muestran cómo organizar distintos tipos de recursos incluidos.

### scripts/
Código ejecutable (Python/Bash/etc.) que puede ejecutarse directamente para realizar operaciones específicas.

**Ejemplos de otras skills:**
- Skill de PDF: `fill_fillable_fields.py`, `extract_form_field_info.py` - utilidades para manipulación de PDF
- Skill de DOCX: `document.py`, `utilities.py` - módulos de Python para procesamiento de documentos

**Apropiado para:** scripts de Python, shell, o cualquier código ejecutable que realice automatización, procesamiento de datos u operaciones específicas.

**Nota:** los scripts pueden ejecutarse sin cargarse al contexto, pero también pueden ser leídos por el Agente para parchearlos o ajustar el entorno.

### referencias/
Documentación y material de referencia pensado para cargarse al contexto e informar el proceso y razonamiento de el Agente.

**Ejemplos de otras skills:**
- Gestión de producto: `communication.md`, `context_building.md` - guías detalladas de flujo de trabajo
- BigQuery: documentación de referencia de API y ejemplos de consultas
- Finanzas: documentación de esquemas, políticas de la empresa

**Apropiado para:** documentación extensa, referencias de API, esquemas de bases de datos, guías completas, o cualquier información detallada que el Agente deba consultar mientras trabaja.

### assets/
Archivos que no están pensados para cargarse al contexto, sino para usarse dentro de la salida que produce el Agente.

**Ejemplos de otras skills:**
- Estilo de marca: archivos de plantilla de PowerPoint (.pptx), archivos de logo
- Constructor de frontend: directorios boilerplate HTML/React
- Tipografía: archivos de fuente (.ttf, .woff2)

**Apropiado para:** plantillas, código boilerplate, plantillas de documentos, imágenes, íconos, fuentes, o cualquier archivo pensado para copiarse o usarse en la salida final.

---

**Cualquier directorio no necesario puede eliminarse.** No todas las skills requieren los tres tipos de recursos.
"""

SCRIPT_EJEMPLO = '''#!/usr/bin/env python3
"""
Script de ejemplo para {skill_name}

Este es un script placeholder que puede ejecutarse directamente.
Reemplázalo con la implementación real o elimínalo si no se necesita.

Ejemplos de scripts reales de otras skills:
- pdf/scripts/fill_fillable_fields.py - rellena campos de formularios PDF
- pdf/scripts/convert_pdf_to_images.py - convierte páginas PDF a imágenes
"""

def main():
    print("Este es un script de ejemplo para {skill_name}")
    # TODO: agrega aquí la lógica real del script
    # Esto podría ser procesamiento de datos, conversión de archivos, llamadas a APIs, etc.

if __name__ == "__main__":
    main()
'''

REFERENCIA_EJEMPLO = """# Documentación de referencia para {skill_title}

Este es un placeholder para documentación de referencia detallada.
Reemplázalo con contenido de referencia real o elimínalo si no se necesita.

Ejemplos de documentos de referencia reales de otras skills:
- product-management/references/communication.md - guía completa para actualizaciones de estado
- product-management/references/context_building.md - profundización en cómo reunir contexto
- bigquery/references/ - referencias de API y ejemplos de consultas

## Cuándo son útiles los documentos de referencia

Los documentos de referencia son ideales para:
- Documentación completa de API
- Guías detalladas de flujo de trabajo
- Procesos complejos de varios pasos
- Información demasiado extensa para el SKILL.md principal
- Contenido que solo se necesita para casos de uso específicos

## Sugerencias de estructura

### Ejemplo de referencia de API
- Resumen
- Autenticación
- Endpoints con ejemplos
- Códigos de error
- Límites de tasa

### Ejemplo de guía de flujo de trabajo
- Prerrequisitos
- Instrucciones paso a paso
- Patrones comunes
- Resolución de problemas
- Buenas prácticas
"""

ASSET_EJEMPLO = """# Archivo de recurso de ejemplo

Este placeholder representa dónde se almacenarían los archivos de recursos (assets).
Reemplázalo con archivos de recursos reales (plantillas, imágenes, fuentes, etc.) o elimínalo si no se necesita.

Los archivos de assets NO están pensados para cargarse al contexto, sino para usarse
dentro de la salida que produce el Agente.

Ejemplos de archivos de assets de otras skills:
- Lineamientos de marca: logo.png, slides_template.pptx
- Constructor de frontend: directorio hello-world/ con boilerplate HTML/React
- Tipografía: fuente-personalizada.ttf, familia-de-fuente.woff2
- Datos: sample_data.csv, test_dataset.json

## Tipos de assets comunes

- Plantillas: .pptx, .docx, directorios boilerplate
- Imágenes: .png, .jpg, .svg, .gif
- Fuentes: .ttf, .otf, .woff, .woff2
- Código boilerplate: directorios de proyecto, archivos de arranque
- Íconos: .ico, .svg
- Archivos de datos: .csv, .json, .xml, .yaml

Nota: esto es un placeholder de texto. Los assets reales pueden ser de cualquier tipo de archivo.
"""


def nombre_a_titulo(skill_name):
    """Convierte un nombre en hyphen-case a Título para mostrar."""
    return ' '.join(palabra.capitalize() for palabra in skill_name.split('-'))


def init_skill(skill_name, path):
    """
    Inicializa un nuevo directorio de skill con un SKILL.md plantilla.

    Args:
        skill_name: Nombre de la skill
        path: Ruta donde debe crearse el directorio de la skill

    Returns:
        Ruta al directorio de la skill creado, o None si hay error
    """
    skill_dir = Path(path).resolve() / skill_name

    if skill_dir.exists():
        print(f"[SKILL CREATION ERROR] Error: el directorio de la skill ya existe: {skill_dir}")
        return None

    try:
        skill_dir.mkdir(parents=True, exist_ok=False)
        print(f"[SKILL CREATION] Directorio de la skill creado: {skill_dir}")
    except Exception as e:
        print(f"[SKILL CREATION ERROR] Error al crear el directorio: {e}")
        return None

    skill_title = nombre_a_titulo(skill_name)
    contenido_skill = PLANTILLA_SKILL.format(
        skill_name=skill_name,
        skill_title=skill_title
    )

    ruta_skill_md = skill_dir / 'SKILL.md'
    try:
        ruta_skill_md.write_text(contenido_skill)
        print(f"[SKILL CREATION] SKILL.md creado: {ruta_skill_md}")
    except Exception as e:
        print(f"[SKILL CREATION ERROR] Error al crear SKILL.md: {e}")
        return None

    try:
        scripts_dir = skill_dir / 'scripts'
        scripts_dir.mkdir(exist_ok=True)
        script_ejemplo = scripts_dir / 'ejemplo.py'
        script_ejemplo.write_text(SCRIPT_EJEMPLO.format(skill_name=skill_name))
        script_ejemplo.chmod(0o755)
        print(f"[SKILL CREATION] scripts/ejemplo.py creado: {script_ejemplo}")

        referencias_dir = skill_dir / 'referencias'
        referencias_dir.mkdir(exist_ok=True)
        referencia_ejemplo = referencias_dir / 'referencia_api.md'
        referencia_ejemplo.write_text(REFERENCIA_EJEMPLO.format(skill_title=skill_title))
        print(f"[SKILL CREATION] referencias/referencia_api.md creado: {referencia_ejemplo}")

        assets_dir = skill_dir / 'assets'
        assets_dir.mkdir(exist_ok=True)
        asset_ejemplo = assets_dir / 'ejemplo_asset.txt'
        asset_ejemplo.write_text(ASSET_EJEMPLO)
        print(f"[SKILL CREATION] assets/ejemplo_asset.txt creado: {asset_ejemplo}")
    except Exception as e:
        print(f"[SKILL CREATION ERROR] Error al crear los directorios de recursos: {e}")
        return None

    print(f"\n[SKILL CREATION] Skill '{skill_name}' inicializada correctamente en {skill_dir}")
    print("\nPróximos pasos:")
    print("1. Edita SKILL.md para completar los TODO y actualizar la descripción")
    print("2. Personaliza o elimina los archivos de ejemplo en scripts/, referencias/ y assets/")
    print("3. Ejecuta el validador cuando esté lista para verificar la estructura de la skill")

    return skill_dir


def main():
    if len(sys.argv) < 4 or sys.argv[2] != '--path':
        print("Uso: init_skill.py <nombre-skill> --path <ruta>")
        print("\nRequisitos del nombre de la skill:")
        print("  - Identificador en hyphen-case (ej. 'analizador-de-datos')")
        print("  - Solo letras minúsculas, dígitos y guiones")
        print("  - Máximo 40 caracteres")
        print("  - Debe coincidir exactamente con el nombre del directorio")
        print("\nEjemplos:")
        print("  init_skill.py mi-skill-nueva --path skills/public")
        print("  init_skill.py mi-ayudante-api --path skills/private")
        print("  init_skill.py skill-personalizada --path /ruta/personalizada")
        sys.exit(1)

    skill_name = sys.argv[1]
    path = sys.argv[3]

    print(f"Inicializando skill: {skill_name}")
    print(f"Ubicación: {path}")
    print()

    resultado = init_skill(skill_name, path)

    if resultado:
        sys.exit(0)
    else:
        sys.exit(1)


if __name__ == "__main__":
    main()
