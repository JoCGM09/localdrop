# Disciplina de contexto — patrones concretos

Objetivo: resolver la tarea leyendo la menor cantidad de texto posible sin perder precisión. Cada patrón de abajo tiene el "en vez de esto" (caro) y el "haz esto" (barato).

## Ubicar una función o clase

En vez de: `view` el archivo completo para encontrar `def calcular_total`.

Haz esto:
```bash
grep -rn "def calcular_total" --include="*.py" .
```
Luego usa `view` con `view_range` solo sobre las ~30-50 líneas
alrededor del resultado, no el archivo entero.

## Confirmar que algo existe (endpoint, variable de entorno, import)

En vez de: abrir cada archivo candidato uno por uno.

Haz esto:
```bash
grep -rln "STRIPE_SECRET_KEY" .
```
`-l` te da solo los nombres de archivo que matchean; ábrelos solo si necesitas el contexto exacto de la línea.

## Ver todos los usos de un símbolo antes de refactorizar

```bash
grep -rn "\bUserRepository\b" --include="*.ts" src/
```
Esto te da todos los call sites de una. Si son pocos (<10), lee cada línea con contexto (`grep -n -C 3`). Si son muchos, agrupa por carpeta antes de decidir si vale la pena abrir cada archivo.

## Explorar la estructura de un repo nuevo

En vez de: `find .` o `ls -R` desde la raíz.

Haz esto:
```bash
view <ruta>          # listado de 2 niveles, ignora node_modules
```
Si necesitas más profundidad, hazlo carpeta por carpeta según lo que vayas necesitando, no de una vez.

## Cuándo SÍ conviene leer el archivo completo

- El archivo es corto (<150 líneas aprox.) y vas a modificarlo de todas formas — el overhead de dos llamadas (`grep` + `view`) no
  compensa.
- Necesitas entender el flujo completo de un módulo pequeño y cohesivo antes de tocarlo (ej. un solo componente de UI).
- Ya usaste `grep` y el resultado son >5 matches dispersos en el mismo archivo — en ese punto es más barato leer el archivo completo una vez que hacer 5 `view_range` distintos.

## Regla de "no releer"

Si un archivo ya apareció en el output de un `view`, `grep` o `str_replace` anterior en esta misma sesión y no ha sido modificado desde entonces, no lo vuelvas a leer completo — referencia el contenido que ya tienes. Excepción: siempre re-`view` un archivo inmediatamente antes de un `str_replace` sobre él si hiciste ediciones en el medio, porque el offset de líneas puede haber cambiado.

## Presupuesto mental por tarea

Para tareas mecánicas y acotadas, un buen presupuesto de referencia:
- 1-3 `grep` para ubicar lo relevante.
- 1-2 `view` con rango acotado.
- El resto del "trabajo" en `bash`/`str_replace`, no en texto razonado sobre archivos completos pegados en el contexto.

Si vas más allá de esto para una tarea que se describe como "chica", es una señal de que el scope creció o que estás leyendo de más — vale la pena pausar y replantear el approach en vez de seguir acumulando contexto.
