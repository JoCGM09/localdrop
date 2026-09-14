#!/usr/bin/env bash

# Colores para la salida
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}========================================================${NC}"
echo -e "${GREEN}                LocalDrop - Arrancando${NC}"
echo -e "${BLUE}========================================================${NC}"
echo ""

# Funcion para matar los procesos si el script se interrumpe
cleanup() {
    echo -e "\n${YELLOW}Deteniendo servidores...${NC}"
    kill $(jobs -p) 2>/dev/null
    exit
}

# Capturar Ctrl+C (SIGINT)
trap cleanup SIGINT

echo -e "${GREEN}[1/2] Iniciando Backend (Go) en puerto 8080...${NC}"
(cd backend && go run main.go) &

echo -e "${GREEN}[2/2] Iniciando Frontend (Astro) en puerto 4321...${NC}"
(cd frontend && npm run dev) &

echo ""
read -p "¿Deseas iniciar Ngrok para acceso desde el celular? (y/N): " choice
case "$choice" in 
  y|Y ) 
    echo -e "${GREEN}[3/3] Iniciando túnel Ngrok...${NC}"
    (cd frontend && npm run dev:tunnel) &
    ;;
  * ) 
    echo -e "${YELLOW}Omitiendo Ngrok. Modo Localhost únicamente.${NC}"
    ;;
esac

echo ""
echo -e "${BLUE}========================================================${NC}"
echo -e "${GREEN}LocalDrop está corriendo en:${NC} http://localhost:4321"
echo -e "${YELLOW}Presiona Ctrl+C para detener todos los servidores.${NC}"
echo -e "${BLUE}========================================================${NC}"

# Esperar a que los procesos terminen (o hasta que el usuario presione Ctrl+C)
wait
