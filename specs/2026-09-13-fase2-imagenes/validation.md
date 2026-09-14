# Validación: Fase 2 - Transferencia de Imágenes

La Fase 2 estará lista para ser mergeada cuando se cumplan los siguientes criterios comprobables:

## Criterios de Aceptación (Funcionales)
1. [ ] **Drag & Drop:** Al arrastrar un archivo `.png` o `.jpg` a la zona de envío, la interfaz cambia y muestra una miniatura previa sin necesidad de hacer click.
2. [ ] **Pegar:** Al presionar `Ctrl+V` (o equivalente) teniendo una imagen en el portapapeles del SO, la app la captura e inicia el proceso para enviar.
3. [ ] **Límite Estricto:** Seleccionar intencionalmente una imagen > 50MB debe mostrar un error de validación inmediato en el cliente ("Archivo muy grande"), sin iniciar conexión ni consumir memoria del backend.
4. [ ] **Transmisión de Binarios:** El archivo es recibido en su totalidad y calidad por el receptor, mostrando la imagen real y sin corrupción provocada por conversiones base64 intermedias de mala calidad.
5. [ ] **Descarga Exitosa:** El botón de "Descargar" en el receptor guarda el archivo en el dispositivo con la extensión correspondiente.
6. [ ] **Retrocompatibilidad:** La transferencia de texto (Fase 1) debe seguir funcionando impecablemente.

## Criterios de Calidad (No Funcionales)
1. [ ] **Memory Leaks Frontend:** Confirmar mediante inspección del código que cada vez que se usa `URL.createObjectURL` en el receptor o en la preview del emisor, en algún punto del ciclo de vida se usa `URL.revokeObjectURL` para liberar la memoria del navegador.
2. [ ] **Control en Go:** El backend no crashea (OOM - Out of memory) si se levantan y cancelan múltiples envíos de imágenes en simultáneo (respetando los límites de la Fase 1).
