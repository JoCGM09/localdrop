#!/usr/bin/env python3
"""
pick_model.py — sugiere qué agente/modelo usar para una tarea.

Uso:
    python3 pick_model.py "agrega tests unitarios para calcular_total"
    python3 pick_model.py "diseña el esquema de la tabla appointments"

No es un clasificador perfecto: es un heurístico rápido basado en
palabras clave, pensado para dar una primera sugerencia antes de leer
references/model-routing.md si el caso es ambiguo.
"""

import sys

ARCHITECTURE_KEYWORDS = [
    "arquitectura", "diseña", "diseño de", "esquema de", "seguridad",
    "compliance", "legal", "retención de datos", "spec", "specs",
    "decide si", "decisión de", "migrar de", "elegir entre",
    "security", "architecture", "design the schema",
]

MECHANICAL_KEYWORDS = [
    "test", "tests", "renombra", "rename", "formatea", "format",
    "docs", "documentación", "changelog", "refactor chico",
    "boilerplate", "traduce", "traducción", "mueve", "move",
    "busca y reemplaza", "find and replace",
]

PLANNING_KEYWORDS = [
    "lee", "resume", "resumen", "explica cómo está",
    "explora", "mapea", "revisa (sin modificar)", "analiza",
    "read", "summarize", "explain how",
]


def classify(task: str) -> tuple[str, str]:
    t = task.lower()

    if any(k in t for k in ARCHITECTURE_KEYWORDS):
        return (
            "Arquitectura / seguridad / specs",
            "Agente `build`, `spec-writer` o `security-reviewer` — modelo capaz.",
        )
    if any(k in t for k in MECHANICAL_KEYWORDS):
        return (
            "Mecánica y acotada",
            "Subagente especializado (ej. `test-writer`) — modelo barato.",
        )
    if any(k in t for k in PLANNING_KEYWORDS):
        return (
            "Lectura / planeación",
            "Agente `plan` o subagente de revisión — modelo barato, sin permiso de escribir.",
        )
    return (
        "No clasificado con confianza",
        "Por default trátala como mecánica y delega a subagente barato; "
        "escala solo si aparece ambigüedad real. Ver references/model-routing.md.",
    )


def main() -> None:
    if len(sys.argv) < 2:
        print(__doc__)
        sys.exit(1)

    task = " ".join(sys.argv[1:])
    category, suggestion = classify(task)

    print(f"Tarea: {task}")
    print(f"Categoría sugerida: {category}")
    print(f"Recomendación: {suggestion}")


if __name__ == "__main__":
    main()
