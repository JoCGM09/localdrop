#!/usr/bin/env python3
"""
check_contrast.py — calcula la relación de contraste WCAG entre dos colores hex y reporta si pasa AA/AAA, sin depender de una herramienta externa.

Uso:
    python3 check_contrast.py "#FFFFFF" "#0066CC"
    python3 check_contrast.py "#FFFFFF" "#0066CC" --large

Con --large, evalúa contra los umbrales de "texto grande" (>=18px o >=14px bold) en vez de texto normal.
"""

import argparse
import sys


def _hex_to_rgb(hex_color: str) -> tuple[int, int, int]:
    hex_color = hex_color.strip().lstrip("#")
    if len(hex_color) == 3:
        hex_color = "".join(c * 2 for c in hex_color)
    if len(hex_color) != 6:
        raise ValueError(f"Color hex inválido: {hex_color!r} (esperado #RRGGBB)")
    try:
        r = int(hex_color[0:2], 16)
        g = int(hex_color[2:4], 16)
        b = int(hex_color[4:6], 16)
    except ValueError as e:
        raise ValueError(f"Color hex inválido: {hex_color!r}") from e
    return r, g, b


def _relative_luminance(rgb: tuple[int, int, int]) -> float:
    def channel(c: int) -> float:
        c_srgb = c / 255.0
        if c_srgb <= 0.03928:
            return c_srgb / 12.92
        return ((c_srgb + 0.055) / 1.055) ** 2.4

    r, g, b = rgb
    r_lin, g_lin, b_lin = channel(r), channel(g), channel(b)
    return 0.2126 * r_lin + 0.7152 * g_lin + 0.0722 * b_lin


def contrast_ratio(hex_a: str, hex_b: str) -> float:
    lum_a = _relative_luminance(_hex_to_rgb(hex_a))
    lum_b = _relative_luminance(_hex_to_rgb(hex_b))
    lighter, darker = max(lum_a, lum_b), min(lum_a, lum_b)
    return (lighter + 0.05) / (darker + 0.05)


def evaluate(ratio: float, large_text: bool) -> dict:
    if large_text:
        aa_threshold, aaa_threshold = 3.0, 4.5
    else:
        aa_threshold, aaa_threshold = 4.5, 7.0
    return {
        "aa": ratio >= aa_threshold,
        "aaa": ratio >= aaa_threshold,
        "aa_threshold": aa_threshold,
        "aaa_threshold": aaa_threshold,
    }


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("color_a", help="Primer color, formato #RRGGBB")
    parser.add_argument("color_b", help="Segundo color, formato #RRGGBB")
    parser.add_argument(
        "--large",
        action="store_true",
        help="Evaluar contra umbrales de texto grande (>=18px o >=14px bold)",
    )
    args = parser.parse_args()

    try:
        ratio = contrast_ratio(args.color_a, args.color_b)
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

    result = evaluate(ratio, args.large)
    tipo_texto = "texto grande" if args.large else "texto normal"

    print(f"[CHECK CONTRAST] Colores: {args.color_a} vs {args.color_b}`]")
    print(f"[CHECK CONTRAST] Relación de contraste: {ratio:.2f}:1")
    print(f"Evaluando como: {tipo_texto}")
    print(
        f"[CHECK CONTRAST] AA  (mínimo {result['aa_threshold']}:1): "
        f"{'[CHECK CONTRAST] pasa' if result['aa'] else '❌ no pasa'}"
    )
    print(
        f"[CHECK CONTRAST] AAA (mínimo {result['aaa_threshold']}:1): "
        f"{'[CHECK CONTRAST] pasa' if result['aaa'] else '❌ no pasa'}"
    )

    if not result["aa"]:
        sys.exit(1)


if __name__ == "__main__":
    main()
