# 🚀 LocalDrop - Transferencia Efímera y Segura en Red Local

**LocalDrop** es una aplicación PWA diseñada para eliminar la fricción al transferir información entre tu laptop y tu celular. Permite compartir **texto plano, imágenes copiadas directamente al portapapeles (`Ctrl+V`) y archivos de cualquier tipo de hasta 50MB**, todo a través de la red local (LAN) mediante códigos PIN efímeros de 4 dígitos.

---

## ⚡ Aspectos Clave y Arquitectura

*   **🔒 Zero-Storage & Memoria Efímera:** Ningún archivo o texto toca el disco duro ni se almacena en bases de datos. Todo vive temporalmente en la memoria RAM del servidor Go y se borra inmediatamente tras ser recibido o al expirar tras 5 minutos.
*   **🔑 Sin Cuentas ni Registros:** La autenticación se realiza mediante un código PIN numérico de 4 dígitos generado al instante para cada transferencia.
*   **📋 Clipboard Nativo de Imágenes:** Puedes hacer un recorte de pantalla (`Win + Shift + S` o `PrtScn`), presionar `Ctrl+V` en la pestaña de Imagen, enviar el PIN, y en tu celular presionar "Copiar" para tener la imagen directamente en el portapapeles de tu teléfono.
*   **🛡️ Seguridad Integrada:** Protección contra ataques de fuerza bruta asociados a la IP del cliente, validación estricta de tamaño en cliente y servidor, y sanitización de orígenes CORS/CSWSH.

---

## 📋 Requisitos Previos

Antes de comenzar, asegúrate de tener instalado:

1.  **[Node.js](https://nodejs.org/es)** (v22.12.0 o superior) - Incluye `npm`.
2.  **[Go](https://go.dev/doc/install)** (v1.21 o superior) - Lenguaje en el que corre el backend.
3.  **Ngrok** (Opcional) - Si tu router doméstico tiene activado el *Aislamiento de AP* (bloqueo entre dispositivos de la misma Wi-Fi), Ngrok permite conectar tu celular por un túnel seguro.

---

## 📦 Instalación Inicial

```bash
# 1. Clona el repositorio
git clone https://github.com/tu-usuario/mobile-clipboard.git
cd mobile-clipboard

# 2. Instala las dependencias del Frontend
cd frontend
npm install
cd ..
```

---

## 🚦 Guía de Arranque (¿Qué Script Utilizar?)

LocalDrop cuenta con scripts automatizados adaptados a cada sistema operativo y entorno de consola. Elige la opción que mejor se adapte a tu terminal:

### 🟢 Opción A: Git Bash, Linux o macOS (Recomendada)
Si usas la terminal **Git Bash** en Windows, macOS o Linux, ejecuta:

```bash
./start.sh
```
* **Ventajas:** Maneja los servidores Go y Astro en la misma ventana de terminal, te pregunta interactivamente si deseas activar Ngrok (`y/N`) y limpia todos los procesos automáticamente al presionar `Ctrl + C`.

---

### 🟢 Opción B: Windows CMD / Doble Clic
Si usas la consola estándar de Windows (`cmd.exe`) o prefieres usar el ratón:

```cmd
start.bat
```
* **O bien:** Haz **doble clic** sobre el archivo `start.bat` desde el Explorador de Archivos de Windows.
* **Ventajas:** Abre ventanas individuales para el Backend, Frontend y Ngrok (opción `1`), fácil de usar en entornos Windows sin Git Bash.

---

### 🟢 Opción C: GNU Make
Si dispones de la herramienta `make` en tu sistema:

```bash
# Para arrancar solo en red local
make dev

# Para arrancar backend, frontend y túnel Ngrok en paralelo
make dev-tunnel
```

---

### 🟢 Opción D: Arranque Manual (Dos Terminales)
Si deseas ver los logs independientes de cada servicio para depuración:

*   **Terminal 1 (Backend Go):**
    ```bash
    cd backend
    go run main.go
    ```
*   **Terminal 2 (Frontend Astro):**
    ```bash
    cd frontend
    npm run dev
    ```

---

## 📱 Guía de Uso Paso a Paso

Una vez que la aplicación esté corriendo, entra a `http://localhost:4321` desde tu navegador.

### 📄 1. Enviar Texto o Enlaces
1. En la pestaña **Enviar**, selecciona **Texto**.
2. Escribe o pega tu texto/enlace en el área correspondiente.
3. Haz clic en **Generar PIN**. Se mostrará un código de 4 dígitos.
4. En tu celular, entra a la aplicación, ve a la pestaña **Recibir**, ingresa el PIN y presiona **Recibir**.
5. Haz clic en **Copiar texto** para llevarlo al portapapeles de tu teléfono.

### 🖼️ 2. Enviar Imágenes desde el Portapapeles (`Ctrl+V`)
1. Copia cualquier imagen o captura de pantalla (`Win+Shift+S`, `PrtScn` o Clic Derecho > Copiar imagen).
2. En LocalDrop (pestaña **Enviar**), selecciona la solapa **Imagen**.
3. Haz clic en la caja de recuadro y presiona **`Ctrl + V`**. Verás la previsualización gráfica de la foto.
4. Haz clic en **Generar PIN**.
5. En el celular, ingresa el PIN en la vista **Recibir**.
6. Aparecerá la imagen renderizada en grande con dos opciones:
   * **Copiar:** Transforma la foto a `image/png` e inyecta la foto directamente en el portapapeles nativo de tu celular (para pegar directo en WhatsApp, Telegram, Notas, etc.).
   * **Descargar:** Guarda el archivo de imagen en tu dispositivo.

### 📁 3. Enviar Archivos o Documentos (Hasta 50MB)
1. En la pestaña **Enviar**, selecciona **Archivo**.
2. Arrastra y suelta (Drag & Drop) cualquier archivo (`.pdf`, `.docx`, `.zip`, etc.) o haz clic para buscarlo en tu equipo.
3. Presiona **Generar PIN**.
4. Ingresa el PIN en el receptor y haz clic en **Descargar Documento**.

---

## 📲 Conectar el Celular a la Laptop

### Método 1: Por Red Wi-Fi Directa (LAN)
Si tu router permite la comunicación entre dispositivos, puedes ingresar desde la barra de direcciones de tu celular a la IP local de tu laptop (Astro te mostrará la IP en consola al arrancar, ej. `http://192.168.1.15:4321`).

> **💡 Truco de Código QR:** En Chrome o Edge en tu PC, haz clic derecho sobre cualquier parte de la página de LocalDrop y selecciona **"Crear código QR para esta página"**. Escanea el código con la cámara de tu teléfono para entrar al instante.

### Método 2: Mediante Túnel Ngrok (Recomendado si la Wi-Fi bloquea conexiones)
Si la página no carga en tu celular usando la IP local (por cortafuegos o aislamiento de AP del router):
1. Al ejecutar `./start.sh` o `start.bat`, activa la opción de **Ngrok**.
2. Se generará un enlace público temporal (ej. `https://xxxx.ngrok-free.app`).
3. Abre esa URL en tu celular. Gracias al proxy integrado en Astro, tanto la interfaz como el WebSocket de backend funcionarán a través del túnel seguro.

---

## 🧪 Pruebas Automatizadas

El proyecto cuenta con una suite completa de pruebas unitarias e integración en Go que validan la generación de PINs, el flujo de sockets, la transmisión de binarios y los límites de seguridad:

```bash
cd backend
go test -v
```

---

## 📂 Estructura del Proyecto

```
mobile-clipboard/
├── backend/                  # Servidor WebSocket en Go
│   ├── main.go               # Lógica de memoria, WS handler y CORS
│   └── main_test.go          # Pruebas unitarias e integración
├── frontend/                 # Cliente Web PWA en Astro + Tailwind
│   ├── astro.config.mjs      # Configuración de Vite y Proxy WebSocket
│   ├── package.json          # Scripts de frontend y ngrok
│   └── src/
│       ├── layouts/Layout.astro   # Layout con sistema global de Toast
│       └── pages/index.astro      # Interfaz completa de Enviar/Recibir
├── specs/                    # Documentación SDD (Misión, Stack y Fases)
├── start.sh                  # Script de arranque para Bash (Git Bash/Linux/Mac)
├── start.bat                 # Script de arranque para Windows CMD
└── Makefile                  # Automatización con GNU Make
```

---

## 📄 Licencia

Este proyecto es de uso personal y de código abierto bajo la licencia MIT.
