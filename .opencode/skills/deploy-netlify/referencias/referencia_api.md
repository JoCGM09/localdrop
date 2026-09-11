# Referencia de Configuración Netlify

## Archivo netlify.toml

El archivo `netlify.toml` permite configurar el comportamiento del sitio sin usar la UI de Netlify.

```toml
[build]
  command = "npm run build"
  publish = "dist"

[dev]
  command = "npm run dev"
  port = 5173
  targetPort = 5173

[[redirects]]
  from = "/*"
  to = "/index.html"
  status = 200
```

## Comandos Útiles de la CLI

| Comando | Descripción |
| :--- | :--- |
| `netlify status` | Muestra información del usuario y del sitio vinculado. |
| `netlify open` | Abre el dashboard de Netlify en el navegador. |
| `netlify dev` | Levanta un servidor local con funciones de Netlify. |
| `netlify build` | Ejecuta el build localmente emulando el entorno de Netlify. |
| `netlify env:list` | Lista todas las variables de entorno configuradas. |

## Configuración con Supabase

Al usar Supabase, es crítico configurar las variables de entorno en Netlify para que el build de producción las incluya.

1. **VITE_SUPABASE_URL**: La URL de tu proyecto.
2. **VITE_SUPABASE_PUBLISHABLE_KEY**: La clave anon/public.

Comando de ejemplo:
```bash
netlify env:set VITE_SUPABASE_URL "https://xyz.supabase.co" --context all
```
