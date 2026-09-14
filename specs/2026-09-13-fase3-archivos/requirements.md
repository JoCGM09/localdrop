# Requirements: Fase 3 - Transferencia de archivos pequeños

## Problema a Resolver
Actualmente la aplicación soporta la transferencia de texto plano e imágenes en memoria. Sin embargo, para cumplir con el alcance completo del roadmap, el usuario necesita poder transferir otro tipo de archivos de uso diario (como `.pdf`, `.docx`, `.zip`, `.csv`) entre dispositivos, de manera igualmente rápida y segura.

## Reglas de Negocio y Restricciones
*   **Límite Estricto:** Los archivos no pueden superar los **50MB**. El frontend debe prevenir la subida de archivos más grandes inmediatamente, y el backend debe rechazar payloads que excedan el `MaxPayloadSize` por seguridad.
*   **Tipos Permitidos:** Se permitirá **cualquier extensión de archivo** (sin lista blanca restrictiva), asumiendo que el usuario transfiere archivos de confianza en su entorno local.
*   **Efimeridad (Zero-Storage):** Los archivos deben persistirse únicamente en la memoria RAM del servidor backend en Go (`Store`). No deben tocar el disco (HDD/SSD) del servidor bajo ninguna circunstancia.
*   **Consumo Único:** Al igual que el texto y las imágenes, una vez el archivo es recibido y descargado (o expira el PIN tras 5 minutos), este se elimina inmediatamente de la memoria RAM.

## Experiencia de Usuario (UI/UX)
*   **Emisor:**
    *   La misma Dropzone implementada en la Fase 2 (Imágenes) debe admitir archivos generales. 
    *   Si se sube un archivo (que no sea imagen), se debe mostrar un ícono genérico de documento, su nombre y peso (ej. `reporte.pdf (2.4 MB)`).
*   **Receptor:**
    *   Al ingresar el PIN y recibir un archivo genérico, la interfaz **NO debe descargar automáticamente** el archivo.
    *   Debe mostrarse una vista con el nombre del archivo, su extensión, su peso, y un botón explícito de "Descargar Documento".

## Fuera de Alcance
*   Envío de múltiples archivos al mismo tiempo bajo un mismo PIN.
*   Transferencia de archivos masivos (>50MB).
*   Streaming de datos; todo el archivo se carga en memoria.