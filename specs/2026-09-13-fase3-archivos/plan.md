# Plan de Implementación: Fase 3 - Archivos y Ngrok

## Grupo 1: Configuración de Entorno (Ngrok)
1. Instalar `ngrok` de manera local o requerir su uso mediante `npx` en el proyecto Frontend.
2. Añadir un script al `package.json` de Astro (ej. `"dev:tunnel": "ngrok http 4321"`) o un comando similar que facilite la exposición pública del puerto local de Astro para resolver las limitaciones de red LAN reportadas.

## Grupo 2: Backend (Metadatos Generales)
1. Adaptar el `ClipboardItem` en `backend/main.go` para manejar un tipo de payload genérico (ej. `PayloadType: "file"`).
2. Añadir el campo `FileName` (string) al `ClipboardItem` y al `WSMessage`, necesario para que el receptor sepa cómo se llama el archivo original y su extensión.
3. Asegurar que las validaciones de límite (ya ajustadas a 52MB en la fase previa) aplican también a esta nueva modalidad.

## Grupo 3: Frontend Lógica Emisor (Archivos)
1. Modificar la vista "Enviar" (la Dropzone de Fase 2) para aceptar archivos en general (`accept="*"` en el input file oculto, y manejar los eventos de *drag/drop* generalizados).
2. Si el archivo capturado no es una imagen (sin *preview* visual), mostrar una caja informativa que indique el nombre del archivo (ej. `documento.pdf`) y su tamaño en MB.
3. Modificar la lógica del WebSocket emisor para notificar al servidor `{"action": "send_file", "filename": "doc.pdf"}` y posteriormente enviar el Blob/ArrayBuffer.

## Grupo 4: Frontend Lógica Receptor (Archivos)
1. Modificar el cliente WebSocket para reaccionar a un nuevo tipo de acción entrante (ej. `file_received`).
2. Diseñar e inyectar el componente UI correspondiente: En lugar de un `<textarea>` o una `<img>`, debe mostrar un contenedor con el ícono de un documento, el nombre (`filename` enviado desde el servidor) y un botón prominente de "Descargar".
3. Al hacer clic en descargar, convertir el Payload binario recibido en un Blob temporal (`URL.createObjectURL(blob)`), forzar la descarga usando un `<a download="filename">` invisible y limpiar la memoria (`URL.revokeObjectURL`).

## Grupo 5: Pruebas Funcionales (Archivos + Tunnel)
1. Iniciar los servidores: Backend Go (`go run main.go`), y el tunnel (`npm run dev:tunnel` o `ngrok`).
2. Acceder a la URL de Ngrok desde el dispositivo móvil.
3. Enviar un archivo de texto o PDF desde la PC (emisor).
4. Recibirlo en el móvil e intentar la descarga.
5. Comprobar las validaciones de límite (subir archivo de >50MB).