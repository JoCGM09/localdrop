# Roadmap

<!-- Generado/actualizado por /constitution. Fases MUY pequeñas: cada fase debe ser demostrable de principio a fin, no una capa técnica aislada ("no: construir toda la capa de base de datos" / "sí: el usuario puede ver la lista de franjas horarias disponibles"). -->

- [x] **Fase 1: Conexión básica y transferencia de texto**
  - Interfaz web simple (Astro) y backend base (Go).
  - Generación y validación de código PIN de seguridad.
  - Envío y recepción bidireccional fiable exclusivamente de texto plano/clipboard.
- [x] **Fase 2: Transferencia de imágenes**
  - Soporte de envíos binarios en memoria para mime types de imágenes.
  - Funcionalidad de subir, enviar y previsualizar imágenes en ambos clientes.
- [x] **Fase 3: Transferencia de archivos pequeños**
  - Soporte genérico para documentos (ej. .pdf, .docx).
  - Aplicación de límite estricto de tamaño de archivo (ej. 50MB) desde cliente y servidor.
- [ ] **Fase 4: Pulido final y manejo de fallos**
  - Refinamiento de la interfaz (UI/UX) utilizando componentes de shadcn.
  - Manejo integral de errores (código inválido, expiración de PIN/timeout, archivo pesado).
  - Empaquetado final y documentación para ejecución fácil.
