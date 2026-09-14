@echo off
echo ========================================================
echo                 LocalDrop - Arrancando
echo ========================================================
echo.
echo Iniciando Backend (Go) en el puerto 8080...
start cmd /k "cd backend && go run main.go"

echo Iniciando Frontend (Astro) en el puerto 4321...
start cmd /k "cd frontend && npm run dev"

echo.
echo Presiona [1] para abrir Ngrok (Túnel público para el celular)
echo Presiona [Cualquier otra tecla] para continuar solo en Localhost.
set /p choice="Opcion: "

if "%choice%"=="1" (
    echo Iniciando túnel Ngrok...
    start cmd /k "cd frontend && npm run dev:tunnel"
)

echo.
echo ========================================================
echo Los servidores se han abierto en nuevas ventanas.
echo LocalDrop esta corriendo en: http://localhost:4321
echo ========================================================
pause