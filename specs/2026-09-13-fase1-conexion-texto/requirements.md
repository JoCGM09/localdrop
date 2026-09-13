# Requerimientos: Fase 1 - Conexión de Texto

## Alcance
La Fase 1 implementará la capacidad core de la aplicación **LocalDrop**: permitir que un dispositivo (ej. computadora) envíe un bloque de texto y otro dispositivo (ej. celular) lo reciba dentro de la misma red local (LAN) utilizando un código PIN temporal.

## Funcionalidad Principal
1.  **Emisor (Enviar):**
    *   Interfaz con un área de texto grande.
    *   Botón para enviar el texto.
    *   Al enviar, la aplicación muestra un código PIN de 4 dígitos generado por el servidor.
    *   La UI se queda esperando a que el receptor consuma el texto.
2.  **Receptor (Recibir):**
    *   Interfaz con un input numérico para el código PIN de 4 dígitos.
    *   Botón para solicitar el texto.
    *   Al introducir el código correcto, recibe el texto, lo muestra en pantalla y proporciona un botón para copiarlo al portapapeles del dispositivo local.
3.  **Backend (Go):**
    *   Manejo en memoria del par `(PIN -> Texto)`.
    *   Comunicación bidireccional en tiempo real vía **WebSockets**.
    *   El texto se elimina inmediatamente después de ser leído.
    *   El texto expira automáticamente a los **5 minutos** si no es leído.

## Decisiones Tomadas
*   **Código PIN:** 4 dígitos numéricos (para máxima velocidad y facilidad de tecleo en móvil).
*   **Protocolo:** WebSockets. Esto permite que la pantalla del emisor se actualice automáticamente diciendo "Transferencia completada" tan pronto como el receptor consuma el PIN.
*   **Persistencia:** Estrictamente efímera en la memoria del proceso Go. No hay bases de datos.
*   **UI/Frontend:** Astro puro (con HTML/TailwindCSS/JS vanilla) siguiendo las directrices de diseño de `specs/brand-definition.md`. No se usará React.
*   **Timeout:** 5 minutos.

## Fuera de Alcance (Non-Goals)
*   **Imágenes y Archivos:** Estrictamente bloqueados para esta fase. Solo texto.
*   **Cifrado E2E (End-to-End):** Dado que funciona en red local y usa un PIN temporal consumible de un solo uso, el cifrado HTTP plano sobre LAN se considera aceptable para esta primera iteración. (Se podrá evaluar HTTPS/TLS local más adelante, pero no es bloqueante hoy).
*   **Historial:** No se guarda ningún tipo de registro de transferencias pasadas.
*   **Uso sobre Internet (WAN):** Expresamente prohibido/no soportado.

## Referencias
*   `specs/mission.md`
*   `specs/tech-stack.md`
*   `specs/brand-definition.md`
