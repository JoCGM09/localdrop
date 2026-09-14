# Plan de Implementación: Fase 5 - Portapapeles de Imágenes

## Grupo 1: Interfaz Base de "Imagen" (Emisor) [HECHO]
1. En `index.astro`, modificar el grupo de Tabs interno (`send-type-*`) para incluir el tercer botón "Imagen".
2. Crear un nuevo contenedor `image-container` que reemplace dinámicamente al de Texto y Archivo cuando se seleccione.
3. El `image-container` debe incluir:
   - Un área/caja vacía que indique "Presiona Ctrl+V para pegar una imagen".
   - Un contenedor oculto de *preview* (`<img>`) donde se renderizará la imagen pegada.
   - Un botón (ej. "X") para limpiar la imagen seleccionada.

## Grupo 2: Lógica de Captura del Portapapeles
1. Modificar el *event listener* global de `paste` en JavaScript.
2. Si el usuario está en el modo `image`, interceptar `e.clipboardData.items`.
3. Filtrar los items para encontrar el primero de tipo `image/*`.
4. Extraer el archivo crudo (`getAsFile()`), crear un `Blob URL` y asignarlo al `src` de la etiqueta `<img>` de previsualización (Recordar liberar el Blob previo con `revokeObjectURL` para evitar leaks).

## Grupo 3: Envío por WebSocket (Adaptación Híbrida)
1. Extender la lógica del botón "Generar PIN" para el modo `image`.
2. Reutilizar el flujo implementado en la Fase 3 (Archivos). Enviar `{"action": "send_file", "filename": "clipboard-image.png"}` al servidor Go.
3. Tras recibir el `ready_for_binary`, enviar el archivo de imagen cruda por el socket. *(Nota: El backend en Go no requiere cambios, ya que su lógica de metadatos generalizados maneja cualquier `Blob` binario perfectamente).*

## Grupo 4: Lógica de Recepción y Visualización
1. En el frontend (vista de recibir), modificar el evento de recepción de WebSocket.
2. Crear un nuevo estado visual (`receive-state-image-success`) que renderice la imagen usando un `Blob URL`.
3. Detectar, usando el `filename` recibido (o por inspección MIME del Blob), si el payload recibido es una imagen (termina en `.png`, `.jpg`, etc.) para renderizar el estado de Imagen en lugar del estado de Archivo genérico.
4. Agregar los botones "Copiar imagen" (utilizando `navigator.clipboard.write([new ClipboardItem(...)])`) y "Descargar".
5. Liberar las URLs de los blobs gráficos al resetear la vista.