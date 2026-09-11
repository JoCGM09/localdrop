# Framework SDD + SECDD + TDD + FDD Universal para Agentes de IA

Este paquete es un punto de partida reusable para cualquier proyecto. Te da el flujo de "constitucion -> spec de feature -> implementacion -> tests -> seguridad -> merge" siguiendo una adecuada disciplina de costo. Originalmente disenado para **OpenCode**, ahora es totalmente interoperable con **Claude Code, Cursor, Windsurf, Codex y Aider**.

Los nombres de las carpetas y sintaxis aqui estan verificados contra la documentacion oficial (https://opencode.ai/docs) al momento de armar esto, pero si algo no carga, corre `opencode --help` o revisa esa URL antes de asumir que el framework esta mal.

## 0. Instalacion e Inicio Rapido

```bash
# Copia esto a la raiz de tu repo (o de un repo nuevo)
git clone https://github.com/tu-usuario/sdd-es-opencode.git mi-proyecto
cd mi-proyecto
# Generar archivos de contexto para otros proveedores (Claude, Cursor, etc.)
node .opencode/scripts/sync.js
git add . && git commit -m "chore: setup framework SDD/seguridad/testing universal"
```

Estructura:

```
.
├── opencode.json              # config: modelos, permisos, agentes
├── AGENTS.md                  # reglas de proyecto (se cargan siempre - FUENTE UNICA DE VERDAD)
├── .opencode/
│   ├── scripts/               # Scripts de sincronizacion y gestion (sync.js, etc.)
│   ├── agent/                 # subagentes especializados (.md)
│   │   ├── spec-writer.md
│   │   ├── test-writer.md
│   │   ├── security-reviewer.md
│   │   └── code-reviewer.md
│   ├── skills/                # conocimiento reusable (SKILL.md)
│   │   ├── seguridad-buenas-practicas/SKILL.md
│   │   |── test-strategy/SKILL.md
│   │   |── generador-de-skills/SKILL.md
│   |   |── definicion-de-marca/SKILL.md
│   │   └── cost-guard/SKILL.md
│   └── commands/               # comandos /slash reusables
│       ├── constitution.md
│       ├── feature.md
│       ├── implement.md
│       ├── test.md
│       └── security-review.md
└── specs/
    ├── mission.md              # plantilla, la llena /constitution
    ├── tech-stack.md           # plantilla, la llena /constitution
    ├── roadmap.md              # plantilla, la llena /constitution
    └── _template-feature/      # referencia - no se usa directo,
                                 # el agente crea una carpeta nueva
                                 # con este mismo esquema
```

## 1. Personaliza antes de usar

- Edita `AGENTS.md`: completa la secciones "Reglas del proyecto" y "Convenciones tecnicas" con tu stack real (lenguaje, como correr tests, como lintear, etc.).
- Revisa `opencode.json`: cambia los strings de modelo (`anthropic/el Agente-sonnet-4-5`, `anthropic/el Agente-haiku-4-5`) por los que tengas disponibles/prefieras. La idea del split no cambia: **modelo capaz para decisiones, modelo barato para trabajo mecanico y repetitivo.**
- Si usas otro proveedor (OpenAI, Google, etc.), el formato de model string cambia (`openai/gpt-...`) - revisa `opencode models` en tu CLI.
- **Ejecuta `node .opencode/scripts/sync.js`** despues de cualquier cambio en `AGENTS.md`, subagentes, comandos o skills para que todos los agentes (Claude, Cursor, etc.) se enteren.

## 2. Interactuar con tu Agente de IA

Gracias a la sincronizacion, todas las IAs leen las mismas reglas, subagentes y skills pero a traves de sus propios archivos (que estan ocultos en tu editor para mantenerlo limpio).

### Si usas OpenCode:
- Ejecuta `/constitution` para inicializar especificaciones.
- Usa `/feature <nombre>` para crear ramas.
- Usa `/implement` para ejecutar tareas de `plan.md`.

### Si usas Claude Code, Cursor, Windsurf, Aider, etc.:
El archivo `CLAUDE.md`, `.cursorrules` o `CONVENTIONS.md` ha sido generado instruyendo al agente para que reconozca los mismos comandos.
Solo diles:
- "Ejecuta /constitution"
- "Ejecuta /feature mi-nueva-feature"
- "Ejecuta /implement"
- "Actua como security-reviewer y revisa el diff"

El agente leera las definiciones en `.opencode/commands/` y `.opencode/agent/` y seguira los mismos pasos.

## 3. El flujo completo, paso a paso (Logica de Roles)

### Paso 1 - Constitucion del proyecto (una sola vez)
```
Comando: /constitution
```
Esto lee tu `README.md` de stakeholders y, despues de preguntarte lo necesario con la tool AskUserQuestions, genera:
- `specs/mission.md`
- `specs/tech-stack.md`
- `specs/roadmap.md`

Revisalos y ajustalos a mano si algo quedo mal - son la base de todo lo demas, vale la pena que sean correctos.

### Paso 2 - Nueva feature (por cada fase del roadmap)
```
Comando: /feature nombre-corto-de-la-feature
```
Lee roadmap.md, crea la rama, la carpeta `specs/YYYY-MM-DD-nombre-feature/` y, tras preguntarte, utiliza las plantillas de _template-feature y genera `requirements.md`, `plan.md`, `validation.md`.

### Paso 3 - Implementar
```
Comando: /implement
```
Ejecuta **un grupo de tareas a la vez** de `plan.md` (no todo de golpe) y se detiene a resumir - asi puedes revisar antes de que siga, en vez de descubrir 40 archivos cambiados de una sola pasada.

Repite esto hasta que `plan.md` este completo.

### Paso 4 - Testing
```
Comando: /test
```
El subagente `test-writer` (modelo barato) escribe y corre tests segun `requirements.md`/`validation.md`, no "todo lo imaginable".

### Paso 5 - Seguridad
```
Comando: /security-review
```
El subagente `security-reviewer` (solo lectura, no puede editar ni correr comandos) revisa el diff contra la checklist de `security-checklist`. Si hay hallazgos critical/high, vuelve a la fase de implementacion para corregirlos antes de seguir.

### Paso 6 - Merge
Checklist en `validation.md` de la feature: tests OK, seguridad OK, plan.md completo -> ejecuta merge normal con tu flujo de git de siempre.

IMPORTANTE: Si necesitas limpiar la ventana de contexto (recomendado para casos mayores al 15%), cierra la sesion y al abrir una nueva ingresa el siguiente prompt:

'''
Lee la constitucion en @docs/specs para recordar nuestras reglas de arquitectura. Luego, abre el ultimo archivo @specs/ y revisa el /plan.md. Revisa cuales tareas ya estan marcadas como completadas, identifica el siguiente grupo de tareas pendiente y ejecutalo siguiendo exactamente las instrucciones de tu System Prompt. Usa tu herramienta AskUserQuestion si tienes dudas antes de escribir codigo.
'''

## 4. Que NO automatizar con IA

- Decisiones de cumplimiento normativo de datos - la IA puede *sugerir*, pero la validacion final es humana/legal.
- Reglas de negocio ambiguas (cancelaciones, reprogramaciones, prioridad de urgencias) - confirmalas con el stakeholder antes de que `/feature` las convierta en requirements.
- El reporte de `security-review` es un primer filtro, no un reemplazo de un pentest o auditoria real antes de produccion.

## 5. Extendiendo el framework

- Necesitas otro subagente? Copia el patron de `.opencode/agent/test-writer.md`: frontmatter con `mode: subagent`, `tools` explicitos, `permission` acotado al minimo necesario. 
- Otra skill? Carpeta nueva en `.opencode/skills/<nombre>/SKILL.md` con frontmatter `name` + `description` corta y el contenido en Markdown plano.
- Otro comando? Archivo nuevo en `.opencode/commands/<nombre>.md`, frontmatter con `description` y opcionalmente `agent:`/`model:`, cuerpo como prompt con `$ARGUMENTS` si necesita parametros.
- **Recuerda:** Siempre que agregues un agente, comando o skill, ejecuta `node .opencode/scripts/sync.js` para actualizar las reglas de los demas proveedores.

## 6. Ocultamiento Visual
Archivos como `CLAUDE.md`, `.cursorrules` o `CONVENTIONS.md` se generan en la raiz para que los agentes los lean, pero el archivo `.vscode/settings.json` los oculta visualmente en tu editor para no generar ruido. Siguen ahi y funcionan perfectamente.