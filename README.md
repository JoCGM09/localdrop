# LocalDrop - Transferencia de Información Local Segura

LocalDrop es una aplicación web PWA diseñada para solucionar la fricción de compartir texto, imágenes y documentos pesados (hasta 50MB) entre una laptop y un celular. Funciona **estrictamente sobre tu red local** usando WebSockets y memoria RAM efímera, garantizando máxima seguridad sin que tus archivos viajen por internet ni se guarden en ningún disco duro.

## Características Principales
*   **Efímero por Diseño:** Los datos se borran inmediatamente tras ser descargados o al pasar 5 minutos.
*   **Zero-Storage:** Ningún archivo toca el disco del servidor, todo viaja directamente por la memoria RAM.
*   **Autenticación por PIN:** Generación y validación de códigos temporales de seguridad para cada transferencia (no necesitas cuentas).
*   **Soporte Multiformato:** Comparte textos largos, imágenes (fotos, capturas de pantalla) y documentos (PDF, Word, ZIP, etc) hasta 50MB.
*   **Sin instalación en móvil:** Funciona abriendo una URL local en el navegador de tu dispositivo móvil, escaneando el código QR (opcional).

## Requisitos de Entorno
*   [Node.js](https://nodejs.org/en) (v22.12.0 o superior)
*   [Go](https://go.dev/doc/install) (v1.21 o superior)
*   NPM (viene con Node)
*   Ngrok (Opcional, para túneles temporales en caso de bloqueos del router)

## Instalación y Arranque Rápido

LocalDrop está compuesto por un Backend en Go y un Frontend en Astro+Tailwind. Para hacer la ejecución sencilla, se incluye un script automático que arranca todo a la vez:

1.  Clona el repositorio:
    ```bash
    git clone https://github.com/tu-usuario/localdrop.git
    cd localdrop
    ```
2.  Instala las dependencias del frontend (solo la primera vez):
    ```bash
    cd frontend && npm install && cd ..
    ```
3.  **Inicia la Aplicación:**
    En la raíz del proyecto, ejecuta el Makefile desde tu terminal:
    ```bash
    # Para arrancar solo de forma local
    make dev

    # O si necesitas usar Ngrok para el celular (bloqueos de router)
    make dev-tunnel
    ```

    *(Nota: Si usas Git Bash o WSL en Windows y no tienes `make`, puedes ejecutar directamente `./start.sh`).*

Una vez iniciados los servicios:
*   La aplicación estará disponible en `http://localhost:4321` (y una URL pública de `loca.lt/ngrok` si activaste el túnel).

## Documentación Técnica Interna
Este proyecto fue desarrollado bajo una arquitectura SDD (Specification Driven Development). Puedes revisar las fases técnicas en los siguientes documentos:
*   [Misión del proyecto](specs/mission.md)
*   [Stack Tecnológico (Go + Astro)](specs/tech-stack.md)
*   [Tokens de diseño UI](specs/brand-definition.md)