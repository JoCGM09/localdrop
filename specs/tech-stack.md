# Stack Tecnológico

<!-- Generado/actualizado por /constitution -->

## Stack Seleccionado
| Capa         | Elección | Justificación |
|--------------|--------|-----------|
| Frontend     | Astro + TypeScript + shadcn | Ligero, tipado seguro y con componentes UI limpios y rápidos de implementar. |
| Backend      | Go | Alto rendimiento, concurrencia nativa y genera un único binario ejecutable fácilmente. |
| Base de datos| Memoria efímera | Los datos se mantienen en la memoria del proceso de Go y se eliminan tras ser recibidos. |
| Autenticación| Código PIN temporal | Generado al enviar, consumido al recibir. Ideal y seguro para sesiones efímeras locales. |
| Hosting/CI   | Local (Laptop host) | El backend corre en la laptop y el celular se conecta como cliente mediante la IP local. |

## Alternativas Descartadas
- **Aplicación móvil nativa (Android/iOS):** Se descartó para evitar la complejidad de desarrollo y distribución que requiere el software nativo, priorizando una solución web local ágil.
- **Bases de datos persistentes (PostgreSQL, SQLite):** Se descartaron ya que la naturaleza de un portapapeles es temporal; persistir la información de manera duradera es un sobreesfuerzo innecesario.

## Estándares Técnicos, Buenas Prácticas y Seguridad
- **Limpieza de memoria:** El servidor debe liberar explícitamente el payload de la memoria una vez que el código PIN sea consumido o expire por tiempo (timeout).
- **Validación de tamaño:** El backend debe rechazar cargas que excedan el límite de tamaño permitido (50MB - 100MB) para prevenir el agotamiento de la memoria local.
- **Entorno LAN:** El servidor no se expondrá públicamente a Internet, funcionando exclusivamente de manera local.
