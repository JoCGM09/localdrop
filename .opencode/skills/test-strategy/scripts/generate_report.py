#!/usr/bin/env python3
"""
generate_report.py — arma el reporte final de una sesión de testing en el formato de assets/templates/reporte-final.md, a partir de
resultados que le pasas como argumentos o de un JSON simple.

Uso directo con argumentos:
    python3 generate_report.py \
        --unitarios-agregados 12 --integracion-agregados 3 --e2e-agregados 1 \
        --pasan 15 --fallan 1 \
        --fallo "test_reserva_horario_limite: bug real, no valida fin de día" \
        --gap "REQ-07: sin test, requiere confirmar regla de cancelación"

Uso con JSON (más cómodo si el runner de tests ya te da un resumen):
    python3 generate_report.py --json resumen.json

Formato esperado del JSON:
{
  "unitarios_agregados": 12,
  "integracion_agregados": 3,
  "e2e_agregados": 1,
  "pasan": 15,
  "fallan": 1,
  "fallos": ["test_x: bug real, ..."],
  "gaps": ["REQ-07: ..."],
  "pendientes_confirmacion": ["¿qué pasa si se cancela a <1h?"]
}
"""

import argparse
import json
import sys
from datetime import date
from pathlib import Path


def build_report(data: dict) -> str:
    total_agregados = (
        data.get("unitarios_agregados", 0)
        + data.get("integracion_agregados", 0)
        + data.get("e2e_agregados", 0)
    )

    lines = []
    lines.append(f"# Reporte de testing — {date.today().isoformat()}")
    lines.append("")
    lines.append("## Tests agregados")
    lines.append(f"- Unitarios: {data.get('unitarios_agregados', 0)}")
    lines.append(f"- Integración: {data.get('integracion_agregados', 0)}")
    lines.append(f"- E2E: {data.get('e2e_agregados', 0)}")
    lines.append(f"- **Total: {total_agregados}**")
    lines.append("")
    lines.append("## Resultados")
    lines.append(f"- Pasan: {data.get('pasan', 0)}")
    lines.append(f"- Fallan: {data.get('fallan', 0)}")
    lines.append("")

    fallos = data.get("fallos", [])
    if fallos:
        lines.append("## Detalle de fallos (bug real vs. test mal escrito)")
        for f in fallos:
            lines.append(f"- {f}")
        lines.append("")

    gaps = data.get("gaps", [])
    if gaps:
        lines.append("## Requisitos sin test asociado")
        for g in gaps:
            lines.append(f"- {g}")
        lines.append("")

    pendientes = data.get("pendientes_confirmacion", [])
    if pendientes:
        lines.append("## Pendiente de confirmar con el usuario")
        lines.append(
            "> No testeado sin validación — ver "
            "references/checklist-no-testear.md"
        )
        for p in pendientes:
            lines.append(f"- {p}")
        lines.append("")

    return "\n".join(lines)


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--json", type=Path, help="Ruta a un JSON con el resumen")
    parser.add_argument("--unitarios-agregados", type=int, default=0)
    parser.add_argument("--integracion-agregados", type=int, default=0)
    parser.add_argument("--e2e-agregados", type=int, default=0)
    parser.add_argument("--pasan", type=int, default=0)
    parser.add_argument("--fallan", type=int, default=0)
    parser.add_argument("--fallo", action="append", default=[], dest="fallos")
    parser.add_argument("--gap", action="append", default=[], dest="gaps")
    parser.add_argument(
        "--pendiente", action="append", default=[], dest="pendientes_confirmacion"
    )
    parser.add_argument("--out", type=Path, default=None)
    args = parser.parse_args()

    if args.json:
        if not args.json.exists():
            print(f"No existe: {args.json}", file=sys.stderr)
            sys.exit(1)
        data = json.loads(args.json.read_text(encoding="utf-8"))
    else:
        data = {
            "unitarios_agregados": args.unitarios_agregados,
            "integracion_agregados": args.integracion_agregados,
            "e2e_agregados": args.e2e_agregados,
            "pasan": args.pasan,
            "fallan": args.fallan,
            "fallos": args.fallos,
            "gaps": args.gaps,
            "pendientes_confirmacion": args.pendientes_confirmacion,
        }

    report = build_report(data)

    if args.out:
        args.out.write_text(report, encoding="utf-8")
        print(f"Reporte escrito en {args.out}")
    else:
        print(report)


if __name__ == "__main__":
    main()
