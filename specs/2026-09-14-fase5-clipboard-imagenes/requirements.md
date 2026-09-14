# Requirements: Fase 5 - Portapapeles de Imágenes

## Problema a Resolver
El usuario frecuentemente copia imágenes directamente al portapapeles de su sistema operativo (ej. mediante recortes de pantalla o `PrtScn`) y necesita una forma rápida de enviarlas. Actualmente, tendría que guardarlas primero en disco como archivo para poder enviarlas por el tab de "Archivo". Se requiere una opción directa y separada para pegar (`Ctrl+V`) y visualizar imágenes.

## Decisiones y Reglas de Negocio
*   **Aislamiento de la Función:** Se agregará un tercer Tab ("Imagen") dedicado exclusivamente al manejo del portapapeles de imágenes, separado de "Texto" y "Archivo".
*   **Entrada de Datos (Emisor):** La vista de "Imagen" constará de un área contenedora donde el usuario simplemente enfocará y presionará `Ctrl+V` (o pegará desde el menú contextual) para capturar la imagen cruda desde su portapapeles.
*   **Visualización (Ambos Lados):**
    *   *Emisor:* Al pegar la imagen, esta debe renderizarse inmediatamente como un *preview* en la caja.
    *   *Receptor:* Al consumir el PIN, la imagen transferida debe mostrarse gráficamente renderizada (`<img>`) en pantalla, no como un archivo genérico.
*   **Acciones del Receptor:** El receptor tendrá acceso a dos botones obligatorios:
    1.  **Copiar al Portapapeles:** Para emular el comportamiento rápido inverso.
    2.  **Descargar:** Para guardar la imagen en disco si lo desea.
*   **Restricciones de Tamaño:** Se mantiene la restricción estricta global de 50MB por payload.

## Fuera de Alcance
*   Edición de imágenes en el navegador (recortar, dibujar, redimensionar).
*   Soporte para pegar múltiples imágenes simultáneamente (solo se capturará y enviará la primera imagen detectada en el portapapeles).