#!/usr/bin/env python3
"""
check_requirements_coverage.py — compara requisitos declarados en un requirements.md/validation.md contra los tests existentes, para detectar qué requisitos no tienen test asociado y evitar tanto gaps como re-testeo de lo ya cubierto.

Convención esperada:
- En requirements.md/validation.md, cada requisito es una línea de lista con un ID entre corchetes al inicio, EJ: [REQ-01] El paciente no puede reservar un horario ya ocupado.
- En los archivos de test, referencia el ID en un comentario o en el nombre/docstring del test, EJ: def test_no_permite_doble_reserva():  # covers: REQ-01

Uso:
    python3 check_requirements_coverage.py \
        --requirements path/to/requirements.md \
        --tests-dir path/to/tests

Si tu proyecto no sigue esta convención de IDs, usa este script como
punto de partida y ajusta las expresiones regulares, o simplemente
haz la revisión manual guiándote por references/checklist-no-testear.md.
"""

import argparse
import re
import sys
from pathlib import Path

REQ_PATTERN = re.compile(r"\[(REQ-[\w\-]+)\]\s*(.+)")
COVERS_PATTERN = re.compile(r"covers:\s*(REQ-[\w\-]+)", re.IGNORECASE)


def parse_requirements(path: Path) -> dict[str, str]:
    reqs = {}
    for line in path.read_text(encoding="utf-8").splitlines():
        m = REQ_PATTERN.search(line)
        if m:
            reqs[m.group(1)] = m.group(2).strip()
    return reqs


def find_covered_ids(tests_dir: Path) -> set[str]:
    covered = set()
    for f in tests_dir.rglob("*"):
        if f.is_file() and f.suffix in {".py", ".ts", ".js", ".tsx", ".jsx"}:
            text = f.read_text(encoding="utf-8", errors="ignore")
            covered.update(COVERS_PATTERN.findall(text))
    return covered


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--requirements", required=True, type=Path)
    parser.add_argument("--tests-dir", required=True, type=Path)
    args = parser.parse_args()

    if not args.requirements.exists():
        print(f"No existe: {args.requirements}")
        sys.exit(1)
    if not args.tests_dir.exists():
        print(f"No existe: {args.tests_dir}")
        sys.exit(1)

    reqs = parse_requirements(args.requirements)
    covered = find_covered_ids(args.tests_dir)

    if not reqs:
        print(
            "No se encontraron requisitos con formato [REQ-XX] en "
            f"{args.requirements}. Revisa la convención esperada en "
            "el docstring de este script."
        )
        sys.exit(0)

    missing = {rid: desc for rid, desc in reqs.items() if rid not in covered}

    print(f"Requisitos totales: {len(reqs)}")
    print(f"Con test asociado:  {len(reqs) - len(missing)}")
    print(f"Sin test asociado:  {len(missing)}")
    print()

    if missing:
        print("Requisitos SIN test (agregar o marcar como pendiente):")
        for rid, desc in missing.items():
            print(f"  - {rid}: {desc}")
    else:
        print("Todos los requisitos tienen al menos un test asociado.")


if __name__ == "__main__":
    main()
