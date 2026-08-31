#!/usr/bin/env python3
"""
Script de validación rápida para skills - versión mínima
"""

import sys
import re
import yaml
from pathlib import Path

def validate_skill(skill_path):
    """Validación básica de una skill"""
    skill_path = Path(skill_path)

    skill_md = skill_path / 'SKILL.md'
    if not skill_md.exists():
        return False, "No se encontró SKILL.md"

    content = skill_md.read_text()
    if not content.startswith('---'):
        return False, "No se encontró frontmatter YAML"

    match = re.match(r'^---\n(.*?)\n---', content, re.DOTALL)
    if not match:
        return False, "Formato de frontmatter inválido"

    frontmatter_text = match.group(1)

    try:
        frontmatter = yaml.safe_load(frontmatter_text)
        if not isinstance(frontmatter, dict):
            return False, "El frontmatter debe ser un diccionario YAML"
    except yaml.YAMLError as e:
        return False, f"YAML inválido en el frontmatter: {e}"

    PROPIEDADES_PERMITIDAS = {'name', 'description', 'license', 'allowed-tools', 'metadata'}

    claves_inesperadas = set(frontmatter.keys()) - PROPIEDADES_PERMITIDAS
    if claves_inesperadas:
        return False, (
            f"Clave(s) inesperada(s) en el frontmatter de SKILL.md: {', '.join(sorted(claves_inesperadas))}. "
            f"Las propiedades permitidas son: {', '.join(sorted(PROPIEDADES_PERMITIDAS))}"
        )

    if 'name' not in frontmatter:
        return False, "Falta 'name' en el frontmatter"
    if 'description' not in frontmatter:
        return False, "Falta 'description' en el frontmatter"

    name = frontmatter.get('name', '')
    if not isinstance(name, str):
        return False, f"'name' debe ser una cadena de texto, se recibió {type(name).__name__}"
    name = name.strip()
    if name:
        if not re.match(r'^[a-z0-9-]+$', name):
            return False, f"El nombre '{name}' debe estar en hyphen-case (minúsculas, dígitos y guiones únicamente)"
        if name.startswith('-') or name.endswith('-') or '--' in name:
            return False, f"El nombre '{name}' no puede empezar/terminar con guión ni contener guiones consecutivos"
        if len(name) > 64:
            return False, f"El nombre es demasiado largo ({len(name)} caracteres). El máximo es 64 caracteres."

    description = frontmatter.get('description', '')
    if not isinstance(description, str):
        return False, f"'description' debe ser una cadena de texto, se recibió {type(description).__name__}"
    description = description.strip()
    if description:
        if '<' in description or '>' in description:
            return False, "La descripción no puede contener corchetes angulares (< o >)"
        if len(description) > 1024:
            return False, f"La descripción es demasiado larga ({len(description)} caracteres). El máximo es 1024 caracteres."

    return True, "¡La skill es válida!"

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Uso: python quick_validate.py <directorio_de_la_skill>")
        sys.exit(1)

    valido, mensaje = validate_skill(sys.argv[1])
    print(mensaje)
    sys.exit(0 if valido else 1)
