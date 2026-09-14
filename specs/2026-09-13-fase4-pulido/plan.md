# Plan de Implementación: Fase 4 - Pulido Final

## Grupo 1: Integración de shadcn y UI/UX Base [HECHO]
1. Inicializar la configuración base de `shadcn/ui` en el proyecto Astro (con soporte React o Vanilla si se prefiriese usar WebComponents simples para evitar más librerías) o crear los componentes base utilizando Tailwind puro basándose en la especificación de `specs/brand-definition.md`.
2. Reemplazar los componentes nativos (botones, inputs, textarea) de la vista de "Enviar" y "Recibir" para que cumplan 100% con los Design Tokens de la marca.

## Grupo 2: Sistema de Notificaciones (Toasts)
1. Implementar o instalar un sistema de Toasts (ej. Sonner, o uno custom con Tailwind + JS) global en el `Layout.astro`.
2. Eliminar todos los mensajes de error `<p id="send-error">` estáticos del HTML.
3. Actualizar la lógica en JavaScript para que:
   - Los errores de WebSocket disparen un Toast Rojo de error.
   - La acción de Copiar texto o terminar transferencia dispare un Toast Verde de éxito.
   - La subida de un archivo excedido dispare una advertencia.

## Grupo 3: Refinamiento del Flujo WebSockets
1. Ajustar los estados visuales (`state-waiting`, `state-success`) añadiendo micro-interacciones (ej. transiciones suaves de opacidad) al cambiar entre ellos en lugar de un cambio abrupto `display: none/block`.
2. Validar que la interfaz se reinicie correctamente sin fugas de memoria al presionar "Volver" o cancelar el envío en cualquier punto.

## Grupo 4: Automatización y Empaquetado
1. Crear un script en la raíz del proyecto (ej. `start.bat` o `start.sh`) que realice las siguientes acciones en paralelo:
   - Levantar el servidor Go en el directorio `backend` (`go run main.go`).
   - Levantar el servidor Astro en el directorio `frontend` (`npm run dev`).
   - (Opcional) Proporcionar un comando integrado de túnel si se solicita entorno externo.
2. Actualizar el `README.md` principal del repositorio para documentar el uso de este script único como la manera principal de ejecutar la aplicación.