# Accesibilidad de color

Detalle de umbrales y verificación para el paso de accesibilidad de `definicion-de-marca`. Esto se aplica **antes** de dar por definitiva cualquier combinación de color en `specs/brand-definition.md`.

## Umbrales WCAG

| Tipo de texto | Relación de contraste mínima (AA) | Relación mínima (AAA) |
|---|---|---|
| Texto normal (< 18px o < 14px bold) | 4.5:1 | 7:1 |
| Texto grande (≥ 18px o ≥ 14px bold) | 3:1 | 4.5:1 |
| Elementos de UI no textuales (bordes de inputs, íconos funcionales) | 3:1 | — |

Para la mayoría de proyectos, apuntar a AA es suficiente. Solo exige AAA si el usuario lo pide explícitamente o el producto tiene requisitos de accesibilidad elevados (como sector público, salud, educación).

## Cómo verificar

Usa `scripts/check_contrast.py` en vez de estimarlo a ojo o pedirle al usuario que lo revise en una herramienta externa:

```bash
python3 scripts/check_contrast.py "#FFFFFF" "#0066CC"
```

El script imprime la relación de contraste exacta y si pasa AA/AAA para texto normal y para texto grande. Corre esto para **cada** combinación texto/fondo que vaya a quedar en `design-tokens.json` (texto sobre primario, texto sobre acento, texto sobre fondo si hay modo oscuro, etc.), no solo la combinación principal.

## Qué hacer si una combinación falla

No la dejes pasar "porque se ve bien". Opciones, en este orden de preferencia:
1. Oscurecer o aclarar uno de los dos colores ligeramente (ajustes de luminosidad) hasta que pase, y volver a correr el script.
2. Si el color de marca no puede tocarse (porque viene de una guía corporativa externa), usar ese color solo para elementos grandes o decorativos, y un color de texto distinto (típicamente negro o blanco puro) para el texto real.
3. Si ninguna variante razonable pasa, decírselo explícitamente al usuario en el documento final en vez de entregar una combinación que no cumple, con la relación de contraste real anotada.

## Daltonismo y semántica de color

- Nunca uses rojo/verde como única forma de distinguir un estado (error/éxito). Acompaña siempre con ícono, texto, o un patrón distinto además del color.
- Si el proyecto va a definir colores semánticos (éxito/error/advertencia/info), verifica también su contraste contra el fondo donde se vayan a usar (un badge de error sobre fondo blanco, por ejemplo), no solo contra el color primario.

## Qué documentar en el resultado final

En `specs/brand-definition.md`, cada combinación texto/fondo relevante debe listar el resultado real de la verificación, no solo un "verificado" genérico. Ejemplo de formato:

```
- Texto `#FFFFFF` sobre primario `#0066CC` → contraste 5.1:1 (pasa AA para texto normal y AAA para texto grande).
```
