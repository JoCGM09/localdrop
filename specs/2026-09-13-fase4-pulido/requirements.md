# Requirements: Fase 4 - Pulido Final y Empaquetado

## Problema a Resolver
La aplicación es 100% funcional y permite enviar texto, imágenes y archivos de forma segura por la red local a través de códigos PIN. Sin embargo, carece de feedback elegante para el usuario en caso de éxito o error, la interfaz visual no ha alcanzado el pulido final según los estándares iniciales, y actualmente arrancar el proyecto requiere iniciar múltiples terminales de forma manual.

## Decisiones y Reglas de Negocio
*   **Alertas y Feedback:** Se integrará un sistema de **Notificaciones Toast** flotantes (que desaparecen solas) para comunicar al usuario eventos como "Texto copiado al portapapeles", "Transferencia completada", "PIN incorrecto" o errores de tamaño, en lugar de utilizar bloques estáticos de error.
*   **Seguridad / Tiempos:** Se mantiene la expiración efímera del PIN en **5 minutos**. Las validaciones de límite de 50MB se mantendrán estrictas en el lado del cliente Web para dar respuesta rápida.
*   **Experiencia de Arranque (DX):** Se creará un script único (ej. archivo `.bat` o de shell script) que se encargue de compilar e iniciar el Backend en Go y el Frontend en Astro con un solo comando o clic.

## Fuera de Alcance
*   Implementación de cuentas de usuario o perfiles.
*   Solución definitiva al bloqueo LAN del router doméstico del usuario. Se asume el uso de red local configurada por el usuario, o en su defecto, que seguirá utilizando el túnel temporal de Ngrok como workaround.
*   Rediseño completo de la arquitectura; el objetivo es refinar lo existente.