#!/usr/bin/env bash
# context_budget_check.sh
#
# Uso: ./context_budget_check.sh <ruta-archivo> [patron-grep-opcional]
#
# Antes de hacer `view` sobre un archivo grande, corre este script.
# Te dice cuántas líneas/tokens aproximados tiene, y si superan el
# umbral, te sugiere usar grep en vez de leer todo el archivo.
#
# Si pasas un segundo argumento (patrón), además corre grep -n con
# ese patrón y te muestra dónde están los matches, para que decidas
# si con eso te alcanza.

set -euo pipefail

LINE_THRESHOLD=300

if [ "$#" -lt 1 ]; then
  echo "Uso: $0 <ruta-archivo> [patron-grep-opcional]"
  exit 1
fi

FILE="$1"
PATTERN="${2:-}"

if [ ! -f "$FILE" ]; then
  echo "No existe el archivo: $FILE"
  exit 1
fi

LINES=$(wc -l < "$FILE" | tr -d ' ')
CHARS=$(wc -c < "$FILE" | tr -d ' ')
# Estimación gruesa: ~4 caracteres por token
APPROX_TOKENS=$((CHARS / 4))

echo "Archivo: $FILE"
echo "Líneas: $LINES"
echo "Tokens aproximados: ~$APPROX_TOKENS"
echo

if [ "$LINES" -gt "$LINE_THRESHOLD" ]; then
  echo "Este archivo supera el umbral de $LINE_THRESHOLD líneas."
  echo "Antes de leerlo completo con 'view', intenta ubicar la"
  echo "sección relevante con grep -n."
  if [ -n "$PATTERN" ]; then
    echo "Resultados para el patrón '$PATTERN':"
    grep -n "$PATTERN" "$FILE" || echo "  (sin matches)"
  else
    echo "   Pasa un patrón como segundo argumento para buscarlo ya."
  fi
else
  echo "Archivo dentro del umbral ($LINE_THRESHOLD líneas). Leerlo"
  echo "completo con 'view' es razonable si lo vas a modificar."
fi
