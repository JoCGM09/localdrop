# Criterios de Validación: Fase 5 - Portapapeles de Imágenes

La Fase 5 se considerará lista para mergear cuando se cumplan los siguientes criterios:

## Pruebas Funcionales (Manuales)
1. [ ] **Pegado Directo:** Al ubicarse en la pestaña "Imagen" de la vista de envío y hacer `Ctrl+V` (tras haber copiado un recorte de pantalla), la imagen aparece renderizada en la caja de previsualización.
2. [ ] **Recepción Gráfica:** Al consumir el PIN en otro dispositivo, la imagen se muestra previsualizada en grande en la pantalla de éxito.
3. [ ] **Acciones del Receptor:**
    - [ ] Hacer clic en "Copiar" guarda la imagen nativamente en el portapapeles del dispositivo receptor.
    - [ ] Hacer clic en "Descargar" inicia la descarga del archivo `.png` o `.jpg` original.
4. [ ] **Filtro Estricto:** Si se intenta pegar texto estando en la pestaña "Imagen", el sistema lo ignora y no se rompe.
5. [ ] **Prevención de Memoria (Leaks):** Al quitar una imagen previsualizada o al darle al botón "Volver" tras una recepción, se ejecutan explícitamente llamadas a `URL.revokeObjectURL()` para no ahogar la RAM del navegador.