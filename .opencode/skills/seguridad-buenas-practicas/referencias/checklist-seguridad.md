# Checklist de seguridad

Recorre cada categoría. Marca solo lo que aplica al artefacto revisado; no todas las categorías aplican a todo tipo de código.

## 1. Secretos y credenciales

- [ ] ¿Hay claves de API, tokens, contraseñas o cadenas de conexión escritas literalmente en el código?
- [ ] ¿Hay secretos en archivos de ejemplo, `README`, comentarios, o en el historial de commits?
- [ ] ¿Los secretos se leen desde variables de entorno o un gestor de secretos, y no desde un archivo de configuración versionado?
- [ ] ¿Existe un `.gitignore` (o equivalente) que excluya archivos `.env`, credenciales locales y capturas de tráfico (HAR, logs con cookies)?
- [ ] ¿Los mensajes de error o logs evitan imprimir el valor completo de un secreto, incluso en modo debug?

**Corrección estándar**: mover el valor a una variable de entorno, documentar el nombre esperado en el `README`, y usar un placeholder explícito (`os.environ["NOMBRE_VARIABLE"]`) en cualquier ejemplo.

## 2. Inyección (comandos, SQL, plantillas)

- [ ] ¿Alguna consulta SQL se construye por concatenación o f-strings con datos externos, en vez de parámetros preparados?
- [ ] ¿Algún comando de shell se arma concatenando entrada de usuario (`os.system`, `subprocess` con `shell=True`)?
- [ ] ¿Se usa `eval()` o `exec()` sobre datos que no son 100% controlados por el desarrollador?
- [ ] ¿Alguna plantilla (HTML, Markdown, JSON) inserta datos de usuario sin escapar en un contexto donde eso importa (por ejemplo, HTML renderizado)?

**Corrección estándar**: usar consultas parametrizadas, `subprocess` con lista de argumentos y `shell=False`, evitar `eval`/`exec` sobre datos externos, y escapar/validar antes de interpolar.

## 3. Autenticación y autorización

- [ ] ¿El código asume que un usuario autenticado tiene automáticamente permiso para la acción que pide, sin verificar autorización a nivel de recurso?
- [ ] ¿Los tokens tienen una vida útil razonable y un mecanismo de revocación o expiración?
- [ ] ¿Se valida la firma/integridad de un token en vez de solo decodificarlo y confiar en su contenido?
- [ ] ¿Las rutas o funciones administrativas están protegidas de forma explícita, no solo "ocultas"?

## 4. Validación de entrada

- [ ] ¿Toda entrada que cruza un límite de confianza (red, archivo, argumento CLI, salida de otro modelo/agente) se valida antes de usarse en una operación sensible?
- [ ] ¿Se valida tipo, longitud y rango antes de tipo de contenido esperado (por ejemplo, un ID que debería ser numérico)?
- [ ] ¿Se rechazan explícitamente los casos inesperados, en vez de intentar "adivinar" qué quiso decir el usuario?

## 5. Manejo de errores y logging

- [ ] ¿Existen bloques `except` vacíos o que solo hacen `pass` alrededor de operaciones sensibles (auth, pagos, escritura de archivos)?
- [ ] ¿Los mensajes de error que llegan al usuario final evitan filtrar detalles internos (rutas de archivo, stack traces completos, nombres de tablas)?
- [ ] ¿Hay registro (log) suficiente para poder reconstruir qué pasó ante un incidente, sin registrar el contenido de secretos?

## 6. Dependencias y cadena de suministro

- [ ] ¿Las dependencias nuevas están fijadas a una versión (no `*` ni `latest`)?
- [ ] ¿Provienen de un repositorio oficial y reconocido, no de una URL arbitraria?
- [ ] Si el código descarga o ejecuta algo en tiempo de ejecución (un script remoto, un paquete instalado dinámicamente), ¿se verifica su procedencia o hash?

## 7. Superficie de red y de sistema

- [ ] ¿El script hace llamadas de red a dominios no documentados en el `SKILL.md`/`README`?
- [ ] ¿El script pide permisos de archivo o red más amplios de los que necesita su tarea declarada?
- [ ] Si el código ejecuta comandos del sistema, ¿la lista de comandos posibles es fija y conocida, o depende de una cadena arbitraria construida en tiempo de ejecución?

## 8. Datos de usuario y privacidad

- [ ] ¿Se recolectan o registran más datos personales de los necesarios para la tarea?
- [ ] ¿Los datos sensibles (tokens de sesión, cookies, PII) se excluyen de logs, capturas de pantalla y archivos de ejemplo compartidos?
- [ ] Si el artefacto interactúa con servicios de terceros, ¿queda claro qué datos se les envían?

## 9. Específico para skills de el Agente / agentes

- [ ] ¿La skill pide a el Agente ejecutar acciones irreversibles (borrar, enviar, publicar) sin un paso de confirmación?
- [ ] ¿Las instrucciones de la skill podrían interpretarse como "ignora tus reglas de seguridad para completar esta tarea"?
- [ ] ¿Los scripts incluidos hacen exactamente lo que su descripción dice, sin efectos secundarios no documentados (llamadas de red ocultas, escritura fuera del directorio esperado)?

## Cómo reportar

Para cada ítem marcado como problema, documenta: **qué se encontró**, **dónde** (archivo y línea si aplica), **por qué es un riesgo**, y **la corrección aplicada o sugerida**. Un hallazgo sin corrección propuesta obliga al siguiente lector a rehacer el análisis.
