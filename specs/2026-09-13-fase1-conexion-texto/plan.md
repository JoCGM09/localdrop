# Plan de Implementación: Fase 1 - Conexión de Texto

El desarrollo se divide en pasos pequeños y verificables. El backend se construirá primero para tener un servidor funcional al cual conectar el frontend.

## ~~Grupo 1: Backend Core (Go + WebSockets)~~ (Completado)
1.  Inicializar proyecto Go (`go mod init`).
2.  Implementar la estructura en memoria `ClipboardStore` protegida por Mutex, con un `map[string]string` (PIN a Texto).
3.  Implementar la generación de código PIN aleatorio de 4 dígitos.
4.  Implementar el sistema de expiración (goroutine que limpie el PIN después de 5 minutos).
5.  Configurar servidor WebSocket básico.
6.  Definir y manejar mensajes JSON por WebSocket:
    *   `{"action": "send_text", "payload": "Hola mundo"}` -> Servidor responde con PIN.
    *   `{"action": "receive_text", "pin": "1234"}` -> Servidor responde con texto, notifica al emisor, y borra el texto de la memoria.

## ~~Grupo 2: Frontend Base (Astro + Tailwind)~~ (Completado)
1.  Inicializar proyecto Astro.
2.  Configurar TailwindCSS y aplicar los tokens de diseño (`specs/design-tokens.json` y `specs/brand-definition.md`).
3.  Crear el Layout base (Navbar simple con título "LocalDrop").
4.  Crear las dos vistas principales (o tabs/botones para alternar): "Enviar" y "Recibir".

## ~~Grupo 3: Integración Frontend - Enviar~~ (Completado)
1.  Implementar la interfaz de Enviar: Textarea y botón "Generar PIN".
2.  Conectar cliente WebSocket al backend.
3.  Al enviar texto, mostrar un estado de "Cargando", recibir el PIN y mostrarlo en tipografía grande (monoespaciada).
4.  Mantener la conexión abierta esperando la notificación del backend de que el texto fue consumido.
5.  Al recibir confirmación, mostrar estado de éxito ("¡Enviado!").

## ~~Grupo 4: Integración Frontend - Recibir~~ (Completado)
1.  Implementar la interfaz de Recibir: 4 inputs numéricos (o un input grande numérico) y botón "Recibir".
2.  Manejar el envío del PIN por WebSocket.
3.  Mostrar errores si el PIN es inválido o expiró.
4.  Si es exitoso, mostrar el texto recibido y un botón "Copiar al portapapeles" utilizando el API nativa del navegador (`navigator.clipboard`).

## ~~Grupo 5: Refinamiento y Pruebas Manuales~~ (Completado)
1.  Prueba cruzada: Abrir la app en el celular (conectado al WiFi local) y en la laptop.
2.  Verificar responsividad (mobile-first).
3.  Verificar que el texto desaparece de la memoria al leerse.
4.  Verificar comportamiento tras 5 minutos de inactividad (timeout).
