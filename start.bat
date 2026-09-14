@echo off
title LocalDrop - Lanzador
echo ========================================================
echo                 LocalDrop - Arrancando
echo ========================================================
echo.

set "ROOT_DIR=%~dp0"

echo [1/2] Iniciando Backend (Go) en puerto 8080...
start "LocalDrop Backend" cmd /k "cd /d "%ROOT_DIR%backend" && go run main.go"

echo [2/2] Iniciando Frontend (Astro) en puerto 4321...
start "LocalDrop Frontend" cmd /k "cd /d "%ROOT_DIR%frontend" && npm run dev"

echo.
echo Presiona [1] para abrir Ngrok (Tunel publico para acceso movil).
echo Presiona [Enter] para continuar solo en Localhost.
set /p choice="Opcion: "

if "%choice%"=="1" (
    echo Iniciando tunel Ngrok...
    start "LocalDrop Ngrok" cmd /k "cd /d "%ROOT_DIR%frontend" && npm run dev:tunnel"
)

echo.
echo ========================================================
echo Los servidores estan corriendo en http://localhost:4321
echo Cierra las ventanas individuales de CMD para detenerlos.
echo ========================================================
pause