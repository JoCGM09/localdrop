# Paletas de color de marca

Banco de categorías e inspiración para el paso 2 de la skill `definicion-de-marca` (definir la paleta del proyecto). Estas paletas son puntos de partida, no una lista para copiar literalmente: elige la categoría según el **tono** que se definió con el usuario, y ajusta los valores exactos si hace falta.

## Empresas tech (estilo de referencia, no reproducir logos ni identidad exacta)

Útiles cuando el usuario menciona una marca conocida como referencia de tono ("algo estilo Stripe"). Usa solo como punto de partida y aclara siempre en el documento final que es "inspirado en", nunca una réplica de la identidad visual real de esa marca.

### Estilo "Slack"
- Primario: `#611F69` (Púrpura profundo)
- Acento: `#E01E5A` (Coral)
- Texto: `#FFFFFF`

### Estilo "Discord"
- Primario: `#5865F2` (Azul "Blurple")
- Acento: `#57F287` (Verde)
- Texto: `#FFFFFF`

### Estilo "Vercel"
- Primario: `#000000` (Negro)
- Acento: `#FFFFFF` (Blanco)
- Texto: `#FFFFFF`

### Estilo "Stripe"
- Primario: `#635BFF` (Azul iris)
- Acento: `#D0D4FF` (Lavanda)
- Texto: `#FFFFFF`

## Startups modernas

### Cálida y enérgica
- Primario: `#FF6B35` (Coral)
- Acento: `#FFA500` (Naranja)
- Texto: `#FFFFFF` o `#2D3436`

### Fría y profesional
- Primario: `#0066CC` (Azul cobalto)
- Acento: `#00B4D8` (Azul cielo)
- Texto: `#FFFFFF`

### Vibrante y audaz
- Primario: `#7928CA` (Púrpura)
- Acento: `#FF0080` (Rosa)
- Texto: `#FFFFFF`

## Naturaleza y bienestar

### Tema bosque
- Primario: `#2D6A4F` (Verde bosque)
- Acento: `#52B788` (Verde salvia)
- Texto: `#FFFFFF`

### Tema océano
- Primario: `#0077B6` (Azul profundo)
- Acento: `#00B4D8` (Aguamarina)
- Texto: `#FFFFFF`

## Degradados

Útiles para el campo `gradient` de `design-tokens.json` cuando el proyecto va a usar fondos con degradado (en una futura og-image o en el hero de una landing).

### Atardecer
```
desde: #FF6B6B (Rojo)
hasta: #FFA500 (Naranja)
texto: #FFFFFF
```

### Océano
```
desde: #0077B6 (Azul)
hasta: #00B4D8 (Cian)
texto: #FFFFFF
```

### Bosque
```
desde: #2D6A4F (Verde)
hasta: #52B788 (Verde claro)
texto: #FFFFFF
```

### Crepúsculo
```
desde: #5865F2 (Azul)
hasta: #7928CA (Púrpura)
texto: #FFFFFF
```

## Cómo elegir entre categorías

Elige por el tono acordado con el usuario en el paso 1 del flujo, no por el rubro del producto:
- Un producto B2B serio → "Fría y profesional".
- Un producto de consumo enérgico → "Cálida y enérgica" o "Vibrante y audaz".
- Salud/mindfulness/sustentabilidad → "Naturaleza y bienestar".
- Un producto técnico para developers → cualquiera de las de "Empresas tech" suele encajar mejor que una paleta de consumo.

No asumas la categoría solo por la industria: un producto serio en una industria "divertida" (o viceversa) debe reflejar el tono que pidió el usuario, no el estereotipo del rubro.

## Consejos de accesibilidad

Ver `referencias/accesibilidad-color.md` para el detalle completo de umbrales y verificación. Resumen rápido:
- Asegura una relación de contraste ≥ 4.5:1 para texto sobre fondo.
- Evita combinaciones rojo/verde como única forma de distinguir un estado.
- Verifica siempre con `scripts/check_contrast.py`, no a ojo.
