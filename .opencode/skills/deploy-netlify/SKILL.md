---
name: deploy-netlify
description: Skill para automatizar el despliegue de aplicaciones web en Netlify utilizando Netlify CLI y configuración manual/automatizada. Actívala cuando el usuario pida "desplegar en Netlify", "configurar Netlify", "subir mi web a Netlify" o necesite automatizar el CI/CD con Netlify.
---

# Deploy Netlify

## Resumen

Esta skill habilita al Agente para configurar y ejecutar despliegues en Netlify. Cubre desde la instalación de la CLI, autenticación, inicialización de proyectos, hasta el despliegue manual o mediante CI/CD (GitHub/GitLab).

## Árbol de decisión del flujo

1.  **¿El proyecto ya está en Netlify?**
    *   Sí: Usa `netlify link` para conectar el directorio local.
    *   No: Usa `netlify init` para crear un nuevo sitio y configurar CI/CD.
2.  **¿Se requiere despliegue manual o automático?**
    *   Manual: `netlify deploy` (draft) o `netlify deploy --prod`.
    *   Automático: Configura el repositorio de Git durante `netlify init`.
3.  **¿Hay variables de entorno (ej. Supabase)?**
    *   Usa `netlify env:set KEY VALUE` para sincronizarlas.

## Flujos de Trabajo

### 1. Preparación e Instalación
Antes de cualquier operación, asegura que la CLI esté disponible:
```bash
npm install -g netlify-cli
```

### 2. Autenticación
Si no hay una sesión activa, solicita al usuario ejecutar:
```bash
netlify login
```
*Nota: En entornos de CI, usa la variable de entorno `NETLIFY_AUTH_TOKEN`.*

### 3. Vinculación o Creación
- **Para proyectos existentes:**
  ```bash
  netlify link
  ```
- **Para nuevos proyectos (con Git):**
  ```bash
  netlify init
  ```

### 4. Despliegue Manual
- **Despliegue de prueba (Draft):**
  ```bash
  netlify deploy --dir=dist
  ```
- **Despliegue a producción:**
  ```bash
  netlify deploy --dir=dist --prod
  ```
*Reemplaza `dist` por la carpeta de salida real (ej. `build`, `public`).*

### 5. Gestión de Variables de Entorno
Para configurar servicios externos como Supabase:
```bash
netlify env:set VITE_SUPABASE_URL "https://your-project.supabase.co"
netlify env:set VITE_SUPABASE_PUBLISHABLE_KEY "your-key"
```

## Recursos

### referencias/
- `referencia_api.md`: Detalles técnicos de comandos comunes y configuración de `netlify.toml`.

### scripts/
- `check_deploy_status.py`: Script para verificar el estado del último deploy (placeholder).
