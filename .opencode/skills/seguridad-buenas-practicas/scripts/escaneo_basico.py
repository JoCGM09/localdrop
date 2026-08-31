#!/usr/bin/env python3
"""
Escaneo estático básico de seguridad.

Herramienta defensiva de ayuda: recorre archivos de texto/código en una ruta
y reporta patrones de alto riesgo conocidos (posibles secretos hardcodeados,
usos peligrosos de eval/exec/os.system, concatenación de SQL, except vacíos,
shell=True, etc). No ejecuta ningún código del proyecto analizado, no hace
llamadas de red, y no modifica nada por sí solo: solo lee archivos y reporta.

Es un complemento a la revisión humana y al checklist de
referencias/checklist-seguridad.md, no un sustituto de ninguno de los dos.

Uso:
    python escaneo_basico.py <ruta-del-proyecto-o-skill>
"""

import re
import sys
from pathlib import Path

EXTENSIONES_TEXTO = {
    ".py", ".js", ".ts", ".jsx", ".tsx", ".sh", ".bash", ".rb", ".go",
    ".java", ".php", ".md", ".json", ".yaml", ".yml", ".env", ".cfg", ".ini",
}

# (nombre_regla, patrón, severidad, explicación)
REGLAS = [
    (
        "posible-secreto-hardcodeado",
        re.compile(
            r'(?i)(api[_-]?key|secret|token|password|passwd|access[_-]?key)\s*[:=]\s*'
            r'["\'][A-Za-z0-9_\-\/\+=]{12,}["\']'
        ),
        "CRÍTICO",
        "Parece un secreto (clave, token o contraseña) escrito literalmente en el archivo.",
    ),
    (
        "os-system",
        re.compile(r"\bos\.system\s*\("),
        "IMPORTANTE",
        "os.system ejecuta una cadena completa en el shell; es fácil de inyectar si contiene entrada externa.",
    ),
    (
        "subprocess-shell-true",
        re.compile(r"subprocess\.\w+\([^)]*shell\s*=\s*True"),
        "IMPORTANTE",
        "subprocess con shell=True es vulnerable a inyección si el comando incluye entrada externa.",
    ),
    (
        "eval-exec",
        re.compile(r"\b(eval|exec)\s*\("),
        "CRÍTICO",
        "eval()/exec() sobre datos no completamente controlados permite ejecución de código arbitrario.",
    ),
    (
        "concatenacion-sql-fstring",
        re.compile(r'(?i)(select|insert|update|delete)\s+.*["\']?\s*\+\s*\w+|f["\'].*\b(select|insert|update|delete)\b', re.IGNORECASE),
        "IMPORTANTE",
        "Posible consulta SQL construida por concatenación o f-string en vez de parámetros preparados.",
    ),
    (
        "except-vacio",
        re.compile(r"except\s*[:\(].*?:\s*\n\s*pass", re.DOTALL),
        "MENOR",
        "Bloque except que solo hace 'pass': revisa si oculta fallos de seguridad silenciosamente.",
    ),
    (
        "pickle-load",
        re.compile(r"\bpickle\.load\s*\("),
        "IMPORTANTE",
        "pickle.load sobre datos no confiables permite ejecución de código arbitrario al deserializar.",
    ),
    (
        "curl-pipe-shell",
        re.compile(r"curl[^\n]*\|\s*(sh|bash)\b"),
        "CRÍTICO",
        "Descargar y ejecutar un script remoto directamente (curl | sh) sin verificar su contenido o hash.",
    ),
]

IGNORAR_DIRECTORIOS = {".git", "node_modules", "__pycache__", ".venv", "venv", "dist", "build"}


def escanear_archivo(ruta):
    hallazgos = []
    try:
        contenido = ruta.read_text(encoding="utf-8", errors="ignore")
    except Exception:
        return hallazgos

    for nombre_regla, patron, severidad, explicacion in REGLAS:
        for match in patron.finditer(contenido):
            numero_linea = contenido[: match.start()].count("\n") + 1
            hallazgos.append({
                "archivo": str(ruta),
                "linea": numero_linea,
                "regla": nombre_regla,
                "severidad": severidad,
                "explicacion": explicacion,
                "fragmento": contenido.splitlines()[numero_linea - 1].strip()[:120],
            })
    return hallazgos


def escanear_ruta(ruta_base):
    ruta_base = Path(ruta_base)
    todos_hallazgos = []

    if ruta_base.is_file():
        archivos = [ruta_base]
    else:
        archivos = [
            p for p in ruta_base.rglob("*")
            if p.is_file()
            and p.suffix.lower() in EXTENSIONES_TEXTO
            and not any(parte in IGNORAR_DIRECTORIOS for parte in p.parts)
        ]

    for archivo in archivos:
        todos_hallazgos.extend(escanear_archivo(archivo))

    return todos_hallazgos


def imprimir_reporte(hallazgos):
    if not hallazgos:
        print("[ESCANEO SEGURIDAD] No se encontraron patrones de riesgo conocidos en el escaneo estático.")
        print("   Esto NO significa que el código sea seguro: sigue aplicando el checklist manual.")
        return

    orden_severidad = {"CRÍTICO": 0, "IMPORTANTE": 1, "MENOR": 2}
    hallazgos.sort(key=lambda h: (orden_severidad.get(h["severidad"], 99), h["archivo"], h["linea"]))

    print(f"[HALLAZGO SEGURIDAD] Se encontraron {len(hallazgos)} coincidencia(s). Revisa cada una manualmente:\n")
    for h in hallazgos:
        print(f"[{h['severidad']}] {h['regla']}")
        print(f"  Archivo: {h['archivo']}:{h['linea']}")
        print(f"  {h['explicacion']}")
        print(f"  > {h['fragmento']}")
        print()

    criticos = sum(1 for h in hallazgos if h["severidad"] == "CRÍTICO")
    if criticos:
        print(f"[HALLAZGO SEGURIDAD] {criticos} hallazgo(s) CRÍTICO(S). No se recomienda distribuir hasta corregirlos.")


def main():
    if len(sys.argv) != 2:
        print("Uso: python escaneo_basico.py <ruta-del-proyecto-o-skill>")
        sys.exit(1)

    ruta = sys.argv[1]
    if not Path(ruta).exists():
        print(f"[HALLAZGO SEGURIDAD] Error: la ruta no existe: {ruta}")
        sys.exit(1)

    hallazgos = escanear_ruta(ruta)
    imprimir_reporte(hallazgos)

    tiene_criticos = any(h["severidad"] == "CRÍTICO" for h in hallazgos)
    sys.exit(1 if tiene_criticos else 0)


if __name__ == "__main__":
    main()
