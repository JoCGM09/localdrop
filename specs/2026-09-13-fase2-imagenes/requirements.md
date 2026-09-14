# Requerimientos: Fase 2 - Transferencia de Imágenes

## Alcance
La Fase 2 extenderá las capacidades de LocalDrop para permitir a los usuarios enviar y recibir una única imagen por sesión en su calidad original. La comunicación seguirá usando el código PIN generado y la misma persistencia en memoria local del servidor.

## Funcionalidad Principal
1. **Emisor (Enviar):**
   * UI que soporta 3 modos de entrada: Botón para seleccionar archivo, Área Drag & Drop, y Pegar (Ctrl+V) desde el portapapeles.
   * Previsualización miniatura de la imagen antes de generar el PIN.
   * Validación estricta del tamaño de la imagen en el cliente (max 50MB).
2. **Receptor (Recibir):**
   * Al recibir correctamente (a través del código PIN de 4 dígitos), se muestra la imagen en pantalla.
   * Botón para descargar/guardar la imagen original en el dispositivo local.
3. **Backend (Go):**
   * Extender el esquema actual de `Store` o el mensaje WS para soportar la recepción de binarios (`[]byte`), o integrar una estructura que diferencie texto de binario.
   * Modificar el límite de lectura del payload de WS para soportar hasta `50MB + overhead`.

## Decisiones Tomadas
* **Límite de tamaño**: Estrictamente `50MB` por imagen.
* **Cantidad**: Únicamente **1 imagen por cada código PIN generado**.
* **Calidad**: **Original**. El cliente no aplicará compresión alguna sobre la imagen.
* **Formato de Transmisión**: **Binario Puro (ArrayBuffer/Blob)** vía WebSockets en lugar de Base64, para maximizar la eficiencia y reducir el consumo de memoria en RAM de Go.

## Fuera de Alcance (Non-Goals)
* Transferencia de múltiples imágenes a la vez en un solo PIN (galerías).
* Envío de formatos que no sean imágenes estándares nativas de navegador (ej. PDF, DOCX se verán en la Fase 3).
* Herramientas de edición en el cliente (recorte, filtros).
