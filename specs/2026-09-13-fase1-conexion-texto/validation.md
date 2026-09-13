# Validación: Fase 1 - Conexión de Texto

La Fase 1 se considerará lista para mergear cuando se cumplan los siguientes criterios comprobables de manera local (Laptop + Celular):

## Criterios de Aceptación (Funcionales)
1.  [ ] **Generación de PIN:** Al introducir texto en el cliente A y presionar "Enviar", el servidor debe retornar un PIN numérico de exactamente 4 dígitos.
2.  [ ] **Recepción Exitosa:** El cliente B puede ingresar ese código de 4 dígitos, solicitar el texto y visualizarlo correctamente.
3.  [ ] **Consumo Único (Seguridad):** Una vez que el cliente B recibe el texto, un tercer intento (ej. Cliente C) ingresando el mismo PIN debe recibir un error de "PIN no encontrado o expirado".
4.  [ ] **Expiración de Tiempo:** Si el cliente A envía un texto y genera un PIN, ese PIN debe volverse inválido automáticamente después de 5 minutos si nadie lo consume.
5.  [ ] **Feedback en Tiempo Real:** El cliente A (emisor) debe ver en su pantalla que la transferencia fue exitosa inmediatamente después de que el cliente B consume el texto, sin necesidad de recargar la página.
6.  [ ] **Portapapeles:** El cliente B debe poder copiar el texto recibido al portapapeles de su dispositivo presionando un solo botón.

## Criterios de Calidad (No Funcionales)
1.  [ ] **Consistencia de Marca:** La interfaz refleja fielmente la tipografía, colores y tono de voz definidos en `specs/brand-definition.md`.
2.  [ ] **Red Local:** El sistema funciona accediendo a la IP local de la laptop desde el navegador del teléfono (ej. `http://192.168.1.100:4321` o el puerto que corresponda a Astro/Go).
3.  [ ] **Linter/Compilación:** El código backend en Go compila sin advertencias, y el build de Astro no presenta errores de tipado.
