# Problemas Conocidos y Bloqueos (Known Issues)

## Problema de Conectividad LAN (Fase 1)
**Fecha:** 13 de Septiembre de 2026
**Estado:** Bloqueado / Pendiente de resolución en el entorno local.

### Descripción
El entorno de desarrollo (Astro y Go) se ejecuta correctamente en `localhost` dentro de la laptop. Sin embargo, al intentar acceder a la aplicación desde un dispositivo móvil conectado a la misma red Wi-Fi (usando la IP local `http://192.168.1.5:4321`), el navegador del celular se queda cargando (timeout) y no logra resolver la página.

### Intentos de Solución Realizados
1. Se configuró Astro para exponerse a la red local añadiendo la bandera `--host` en `package.json` (`npm run dev -- --host`).
2. Se sugirió abrir el puerto `4321` (y `8080`) en el Firewall de Windows Defender mediante PowerShell.
3. Se sugirió cambiar el perfil de la red Wi-Fi en Windows de "Público" a "Privado".

### Posibles Causas Restantes
* **AP Isolation (Aislamiento de AP):** El router de la red doméstica podría tener activada una función de aislamiento que impide que los dispositivos inalámbricos se comuniquen entre sí, incluso si están en la misma subred.
* **Firewall de Antivirus de Terceros:** Algún software antivirus (McAfee, Norton, ESET, etc.) podría estar bloqueando las conexiones entrantes sobreescribiendo las reglas de Windows Defender.
* **Redes Virtuales:** Interfaces de red de Docker, WSL, o VPNs podrían estar interfiriendo con el enrutamiento de la IP.

### Workarounds Propuestos para Sesiones Futuras
Si el problema de red persiste, se pueden evaluar estas alternativas para continuar las pruebas:
1. **Crear un Mobile Hotspot desde la Laptop:** Usar la función de Windows "Zona con cobertura inalámbrica móvil" para que el celular se conecte directamente a la laptop (bypass del router).
2. **Prueba Cruzada Local:** Abrir dos navegadores distintos en la misma laptop (ej. Chrome y Firefox) simulando ser el emisor y el receptor. Esto permite validar la lógica y la UI sin depender del celular por el momento.
