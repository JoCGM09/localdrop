# Criterios de Validación: Fase 3 - Archivos

La Fase 3 se considerará lista para mergear cuando se cumplan todos los siguientes puntos:

## Funcionalidad
1. [ ] **Archivos de cualquier tipo:** Un usuario puede arrastrar o seleccionar un archivo que no sea imagen (ej. un `.pdf` o `.docx`) y enviarlo correctamente.
2. [ ] **Metadatos en la Recepción:** El dispositivo receptor muestra claramente el nombre original del archivo enviado y permite descargarlo conservando esa extensión.
3. [ ] **Descarga Exitosa:** El archivo descargado desde el cliente Web se puede abrir localmente en el SO sin advertencias de corrupción.
4. [ ] **Límite Estricto:** Si se intenta cargar un archivo >50MB en el Frontend, este bloquea la carga inmediatamente y muestra un mensaje de error sin enviar datos al WebSocket.
5. [ ] **Tunneling Funcional:** El script con Ngrok permite probar el flujo completo entre PC y Celular sin depender del enrutamiento de la red local.

## Calidad de Código
1. [ ] **Sin persistencia en disco:** No se usa el paquete `os/exec` ni llamadas a `os.WriteFile` en el backend para guardar los datos temporalmente. Todo sigue viviendo en la memoria RAM `Store.items`.
2. [ ] **Limpieza de Blobs:** En el Frontend, tras descargar un archivo o resetear la vista, se llama a `URL.revokeObjectURL()` para no consumir memoria innecesaria del navegador.