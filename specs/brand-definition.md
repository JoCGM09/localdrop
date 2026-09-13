# Definición de Marca: LocalDrop

## Tono de Voz y Personalidad
*   **Amigable:** Cercano, sin jerga técnica innecesaria. Habla directo al usuario como si fuera una herramienta de su día a día.
*   **Moderno:** Fresco, rápido, sin fricciones. Inspirado en la fluidez de herramientas como AirDrop.
*   **Directo:** Minimiza los pasos. Los mensajes deben ser concisos y orientados a la acción.

**Ejemplos de Copy:**
*   *Call to Action:* "Enviar texto" / "Recibir" (En vez de "Transmitir carga útil").
*   *Mensaje de éxito:* "¡Texto copiado al portapapeles!"
*   *Mensaje de espera:* "Listo para recibir. Ingresa el código PIN."

## Paleta de Color
Se ha elegido una paleta moderna con un azul vibrante pero accesible como color principal, transmitiendo confianza y tecnología limpia. Los fondos son muy claros para maximizar la legibilidad, y el texto es un gris muy oscuro para reducir la fatiga visual en contraste extremo puro (#000000).

*   **Color Primario (Acento):** `#1e40af` (Blue 800 de Tailwind)
*   **Fondo Principal:** `#ffffff` (Blanco)
*   **Fondo Secundario (Tarjetas/Inputs):** `#f8fafc` (Slate 50)
*   **Texto Principal:** `#0f172a` (Slate 900)
*   **Texto Secundario:** `#334155` (Slate 700)

**Colores Semánticos (Verificados por accesibilidad):**
*   **Éxito:** `#166534` (Green 800)
*   **Error:** `#991b1b` (Red 800)

### Accesibilidad (Contraste Verificado)
Se corrió `check_contrast.py` para asegurar que las combinaciones pasen las normativas WCAG.
*   Texto Principal (`#0f172a`) sobre Fondo Blanco (`#ffffff`): **17.85:1 (Pasa AAA)**
*   Texto Secundario (`#334155`) sobre Fondo Secundario (`#f8fafc`): **9.90:1 (Pasa AAA)**
*   Color Primario (`#1e40af`) sobre Blanco (Para botones/links): **8.72:1 (Pasa AAA)**
*   Texto Éxito (`#166534`) sobre Blanco: **7.13:1 (Pasa AAA)**
*   Texto Error (`#991b1b`) sobre Blanco: **8.31:1 (Pasa AAA)**

## Tipografía
*   **Familia (General):** Inter (o fuentes sans-serif nativas del sistema como `-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto`). Es neutra, moderna y muy legible en interfaces de usuario.
*   **Familia (Monoespaciada):** Para mostrar el código PIN (ej. `JetBrains Mono` o `ui-monospace`). Ayuda a distinguir claramente los números.
*   **Escala básica:**
    *   Títulos grandes (Código PIN): 3rem (48px) - Bold
    *   Títulos (Headers): 1.5rem (24px) - Semibold
    *   Cuerpo de texto: 1rem (16px) - Regular
    *   Texto pequeño (Labels): 0.875rem (14px) - Medium

## Iconografía (Opcional/Futuro)
*   Estilo: Línea (Stroke), no rellenos (Fill), con un grosor de 2px.
*   Esquinas redondeadas para mantener el tono amigable.
*   (Ej: Lucide Icons se ajusta perfectamente a este estilo).
