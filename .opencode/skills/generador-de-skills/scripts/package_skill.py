#!/usr/bin/env python3
"""
Empaquetador de skills - Crea un archivo .skill distribuible de una carpeta de skill

Uso:
    python scripts/package_skill.py <ruta/a/carpeta-skill> [directorio-salida]

Ejemplo:
    python scripts/package_skill.py skills/public/mi-skill
    python scripts/package_skill.py skills/public/mi-skill ./dist
"""

import sys
import zipfile
from pathlib import Path
from quick_validate import validate_skill


def package_skill(skill_path, output_dir=None):
    """
    Empaqueta una carpeta de skill en un archivo .skill.

    Args:
        skill_path: Ruta a la carpeta de la skill
        output_dir: Directorio de salida opcional para el archivo .skill (por defecto el directorio actual)

    Returns:
        Ruta al archivo .skill creado, o None si hay error
    """
    skill_path = Path(skill_path).resolve()

    if not skill_path.exists():
        print(f"[SKILL PACKAGING ERROR] Error: no se encontró la carpeta de la skill: {skill_path}")
        return None

    if not skill_path.is_dir():
        print(f"[SKILL PACKAGING ERROR] Error: la ruta no es un directorio: {skill_path}")
        return None

    skill_md = skill_path / "SKILL.md"
    if not skill_md.exists():
        print(f"[SKILL PACKAGING ERROR] Error: no se encontró SKILL.md en {skill_path}")
        return None

    print("🔍 Validando skill...")
    valido, mensaje = validate_skill(skill_path)
    if not valido:
        print(f"[SKILL PACKAGING ERROR] Validación fallida: {mensaje}")
        print("   Corrige los errores de validación antes de empaquetar.")
        return None
    print(f"[SKILL PACKAGING] {mensaje}\n")

    skill_name = skill_path.name
    if output_dir:
        output_path = Path(output_dir).resolve()
        output_path.mkdir(parents=True, exist_ok=True)
    else:
        output_path = Path.cwd()

    skill_filename = output_path / f"{skill_name}.skill"

    try:
        with zipfile.ZipFile(skill_filename, 'w', zipfile.ZIP_DEFLATED) as zipf:
            for file_path in skill_path.rglob('*'):
                if file_path.is_file():
                    arcname = file_path.relative_to(skill_path.parent)
                    zipf.write(file_path, arcname)
                    print(f"  Agregado: {arcname}")

        print(f"\n[SKILL PACKAGING] Skill empaquetada correctamente en: {skill_filename}")
        return skill_filename

    except Exception as e:
        print(f"[SKILL PACKAGING ERROR] Error al crear el archivo .skill: {e}")
        return None


def main():
    if len(sys.argv) < 2:
        print("Uso: python scripts/package_skill.py <ruta/a/carpeta-skill> [directorio-salida]")
        print("\nEjemplo:")
        print("  python scripts/package_skill.py skills/public/mi-skill")
        print("  python scripts/package_skill.py skills/public/mi-skill ./dist")
        sys.exit(1)

    skill_path = sys.argv[1]
    output_dir = sys.argv[2] if len(sys.argv) > 2 else None

    print(f"📦 Empaquetando skill: {skill_path}")
    if output_dir:
        print(f"   Directorio de salida: {output_dir}")
    print()

    resultado = package_skill(skill_path, output_dir)

    if resultado:
        sys.exit(0)
    else:
        sys.exit(1)


if __name__ == "__main__":
    main()
