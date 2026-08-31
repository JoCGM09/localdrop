# Framework SDD + SECDD + TDD + FDD para opencode

Este paquete es un punto de partida reusable para cualquier proyecto que armes con opencode: te da el flujo de "constitución → spec de feature → implementación → tests → seguridad → merge" siguiendo una adecuada disciplina de costo.

Los nombres de las carpetas y sintaxis aquí están verificados contra la documentación oficial (https://opencode.ai/docs) al momento de armar esto, pero si algo no carga, corre `opencode --help` o revisa esa URL antes de asumir que el framework está mal.

## 0. Instalación en tu proyecto

```bash
# Copia esto a la raíz de tu repo (o de un repo nuevo)
cp -r opencode-sdd-framework/. /ruta/a/tu/proyecto/
cd /ruta/a/tu/proyecto
git add . && git commit -m "chore: setup framework SDD/seguridad/testing para opencode"
```

Estructura:

```
.
├── opencode.json              # config: modelos, permisos, agentes
├── AGENTS.md                  # reglas de proyecto (se cargan siempre)
├── .opencode/
│   ├── agent/                 # subagentes especializados (.md)
│   │   ├── spec-writer.md
│   │   ├── test-writer.md
│   │   ├── security-reviewer.md
│   │   └── code-reviewer.md
│   ├── skills/                # conocimiento reusable (SKILL.md)
│   │   ├── seguridad-buenas-practicas/SKILL.md
│   │   |── test-strategy/SKILL.md
│   │   |── generador-de-skills/SKILL.md
|   |   |── definicion-de-marca/SKILL.md
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
    └── _template-feature/      # referencia — no se usa directo,
                                 # /feature crea una carpeta nueva
                                 # con este mismo esquema
```

## 1. Personaliza antes de usar

- Edita `AGENTS.md`: completa la secciones "Reglas del proyecto" y "Convenciones técnicas" con tu stack real (lenguaje, cómo correr tests, cómo lintear, etc.).
- Revisa `opencode.json`: cambia los strings de modelo (`anthropic/el Agente-sonnet-4-5`, `anthropic/el Agente-haiku-4-5`) por los que tengas disponibles/prefieras. La idea del split no cambia: **modelo capaz para decisiones, modelo barato para trabajo mecánico y repetitivo.**
- Si usas otro proveedor (OpenAI, Google, etc.), el formato de model string cambia (`openai/gpt-...`) — revisa `opencode models` en tu CLI.

## 2. El flujo completo, paso a paso

### Paso 1 — Constitución del proyecto (una sola vez)
```
Agent: spec-writer - command: /constitution
```
Esto lee tu `README.md` de stakeholders y, después de preguntarte lo necesario con la tool AskUserQuestions, genera:
- `specs/mission.md`
- `specs/tech-stack.md`
- `specs/roadmap.md`

Revísalos y ajústalos a mano si algo quedó mal — son la base de todo lo demás, vale la pena que sean correctos.

### Paso 2 — Nueva feature (por cada fase del roadmap)
```
Agent: spec-writter - command: /feature nombre-corto-de-la-feature
```
Lee roadmap.md, crea la rama, la carpeta `specs/YYYY-MM-DD-nombre-feature/` y, tras preguntarte, utiliza las plantillas de _template-feature y genera `requirements.md`, `plan.md`, `validation.md`.

### Paso 3 — Implementar
```
Agent: build - command: /implement
```
Ejecuta **un grupo de tareas a la vez** de `plan.md` (no todo de golpe) y se detiene a resumir — así puedes revisar antes de que siga, en vez de descubrir 40 archivos cambiados de una sola pasada.

Repite `/implement` hasta que `plan.md` esté completo.

### Paso 4 — Testing
```
Agent: test-writer - command: /test
```
El subagente `test-writer` (modelo barato) escribe y corre tests según `requirements.md`/`validation.md`, no "todo lo imaginable".

### Paso 5 — Seguridad
```
Agent: security-reviewer - command: /security-review
```
El subagente `security-reviewer` (solo lectura, no puede editar ni correr comandos) revisa el diff contra la checklist de `security-checklist`. Si hay hallazgos critical/high, vuelve a `/implement` para corregirlos antes de seguir.

### Paso 6 — Merge
Checklist en `validation.md` de la feature: tests OK, seguridad OK, plan.md completo → merge normal con tu flujo de git de siempre.

IMPORTANTE: Si necesitas limpiar la ventana de contexto (recomendado para casos mayores al 15%), cierra la sesión y al abrir una nueva ingresa el siguiente prompt:

'''
Lee la constitución en @docs/specs para recordar nuestras reglas de arquitectura. Luego, abre el archivo @specs/2026-08-31-perfiles-y-gestion/plan.md. Revisa cuáles tareas ya están marcadas como completadas, identifica el siguiente grupo de tareas pendiente y ejecútalo siguiendo exactamente las instrucciones de tu System Prompt. Usa tu herramienta AskUserQuestion si tienes dudas antes de escribir código.
'''

## 3. Qué NO automatizar con IA

- Decisiones de cumplimiento normativo de datos — la IA puede *sugerir*, pero la validación final es humana/legal.
- Reglas de negocio ambiguas (cancelaciones, reprogramaciones, prioridad de urgencias) — confírmalas con el stakeholder antes de que `/feature` las convierta en requirements.
- El reporte de `security-review` es un primer filtro, no un reemplazo de un pentest o auditoría real antes de producción.

## 5. Extendiendo el framework

- ¿Necesitas otro subagente? Copia el patrón de `.opencode/agent/test-writer.md`: frontmatter con `mode: subagent`, `tools` explícitos, `permission` acotado al mínimo necesario. 
- ¿Otra skill? Carpeta nueva en `.opencode/skills/<nombre>/SKILL.md` con frontmatter `name` + `description` corta y el contenido en Markdown plano.
- ¿Otro comando? Archivo nuevo en `.opencode/commands/<nombre>.md`, frontmatter con `description` y opcionalmente `agent:`/`model:`, cuerpo como prompt con `$ARGUMENTS` si necesita parámetros.
