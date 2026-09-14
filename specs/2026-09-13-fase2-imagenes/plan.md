# Plan de Implementación: Fase 2 - Transferencia de Imágenes

## Grupo 1: Backend (Soporte de Binarios) [HECHO]
1. Modificar la constante `MaxPayloadSize` en `main.go` a un valor seguro apenas superior a 50MB (ej. `52MB` para cubrir overhead JSON/Binario).
2. Adaptar la estructura `ClipboardItem` y el protocolo WS. Como los WebSockets soportan frames binarios nativos, diseñar el flujo: 
   - El cliente envía `{"action": "send_image", "ext": "png"}` -> Servidor responde `{"action": "ready_for_binary", "pin": "..."}`.
   - El cliente envía el binario (Blob).
   - O bien, unificar un modelo híbrido en memoria (`PayloadType: "text" | "image"`, `Data []byte`).
3. Modificar `Save` y `Retrieve` en Go para soportar el nuevo tipo de dato sin romper la Fase 1 (Texto).

## Grupo 2: Frontend Base UI (Imágenes)
1. En `index.astro`, agregar una solapa o toggle (Text / Image) en la vista "Enviar".
2. Crear la interfaz de Dropzone para la imagen:
   - Área punteada con el icono de imagen.
   - Texto "Arrastra aquí, pega (Ctrl+V) o haz click".
   - Input file oculto (`accept="image/*"`).

## Grupo 3: Frontend Lógica Emisor (Imágenes)
1. Implementar la captura de la imagen desde los 3 eventos: `drop`, `paste`, y `change` (input click).
2. Validar que el archivo sea una imagen y que no supere los 50MB.
3. Generar una vista previa local (URL.createObjectURL) al seleccionar la imagen.
4. Adaptar la lógica WebSocket del botón "Generar PIN" para enviar el Blob binario (o la estructura acordada en el Grupo 1) al backend.

## Grupo 4: Frontend Lógica Receptor (Imágenes)
1. Modificar la escucha en la vista "Recibir" para interpretar el payload si es texto (como string) o si es una imagen (Blob).
2. Si es imagen, crear un object URL e inyectarlo en una etiqueta `<img>` para que el receptor pueda verla.
3. Añadir el botón "Descargar Imagen" (`<a download="localdrop-image..." href="...">`) para guardar el archivo.
4. Asegurar de limpiar las URLs de objeto (`URL.revokeObjectURL`) cuando la vista se resetee para no causar *memory leaks*.

## Grupo 5: Pruebas Manuales (Imágenes)
1. Exponer la aplicación frontend usando ngrok (`ngrok http 4321`) para poder acceder desde el celular temporalmente.
2. Enviar una imagen pequeña mediante Drag & Drop (PC -> Celular) y validar que se ve y descarga bien.
3. Copiar y pegar una captura de pantalla en el emisor y validar el flujo.
4. Enviar una imagen pesada (ej. 20-40MB) y asegurar que ni el cliente ni el backend Go crashean y se respeta el tiempo.
5. Intentar enviar una imagen de 60MB para probar que el frontend lo rechace de inmediato.
