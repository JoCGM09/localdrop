const fs = require('fs');
const path = require('path');

const rootDir = path.join(__dirname, '..', '..');
const agentsFile = path.join(rootDir, 'AGENTS.md');
const skillsDir = path.join(rootDir, '.opencode', 'skills');
const agentDir = path.join(rootDir, '.opencode', 'agent');
const commandDir = path.join(rootDir, '.opencode', 'commands');

console.log('Sincronizando reglas y comandos de SDD Framework...');

try {
  const agentsMdContent = fs.existsSync(agentsFile) ? fs.readFileSync(agentsFile, 'utf8') : '';
  
  // 1. Mapeo de Subagentes
  let rolesInfo = '## Roles y Subagentes\n\nEste proyecto define roles específicos para distintas tareas de IA:\n';
  if (fs.existsSync(agentDir)) {
      const agents = fs.readdirSync(agentDir).filter(f => f.endsWith('.md'));
      agents.forEach(agent => {
          const content = fs.readFileSync(path.join(agentDir, agent), 'utf8');
          const match = content.match(/description:\s*(.+)/);
          const desc = match ? match[1] : 'Rol del sistema';
          rolesInfo += `- **${agent.replace('.md', '')}**: ${desc} (Ver detalles en \`.opencode/agent/${agent}\`)\n`;
      });
  }

  // 2. Mapeo de Skills
  let skillsInfo = '## Skills Disponibles\n\nPuedes consultar o invocar estas habilidades:\n';
  if (fs.existsSync(skillsDir)) {
      const skills = fs.readdirSync(skillsDir);
      skills.forEach(skill => {
          const skillPath = path.join(skillsDir, skill, 'SKILL.md');
          if (fs.existsSync(skillPath)) {
            skillsInfo += `- **${skill}**: (Instrucciones en \`.opencode/skills/${skill}/SKILL.md\`)\n`;
          }
      });
  }

  // 3. Mapeo EXPLICITO de Comandos (La clave de la interoperabilidad)
  let commandsInfo = '## Comandos del Proyecto\n\nEste repositorio soporta los siguientes comandos que simulan el comportamiento nativo. Si el usuario te pide ejecutar alguno, aplica la lógica descrita:\n\n';
  if (fs.existsSync(commandDir)) {
      const commands = fs.readdirSync(commandDir).filter(f => f.endsWith('.md'));
      commands.forEach(cmd => {
          const content = fs.readFileSync(path.join(commandDir, cmd), 'utf8');
          const matchDesc = content.match(/description:\s*(.+)/);
          const matchAgent = content.match(/agent:\s*(.+)/);
          
          const desc = matchDesc ? matchDesc[1] : 'Ejecuta esta acción';
          const agent = matchAgent ? matchAgent[1] : 'cualquiera';
          
          commandsInfo += `### Comando: \`/${cmd.replace('.md', '')}\`\n`;
          commandsInfo += `- **Qué hace:** ${desc}\n`;
          commandsInfo += `- **Rol que debes asumir:** \`${agent}\` (Lee \`.opencode/agent/${agent}.md\`)\n`;
          commandsInfo += `- **Lógica de ejecución:** Sigue paso a paso las instrucciones detalladas en \`.opencode/commands/${cmd}\`.\n\n`;
      });
  }

  const baseContext = `${agentsMdContent}\n\n---\n\n${rolesInfo}\n\n---\n\n${skillsInfo}\n\n---\n\n${commandsInfo}`;

  // Archivo para Claude Code
  const claudeMdContent = `# Instrucciones de Proyecto (Claude Code)\n\n${baseContext}\n\n## Instrucciones para el Agente Claude\n- Si el usuario menciona un comando que empiece con \`/\` (ej. \`/constitution\`), búscalo en la lista de comandos de arriba y ejecuta la lógica exacta de su archivo asociado. Asume automáticamente el rol requerido.`;
  fs.writeFileSync(path.join(rootDir, 'CLAUDE.md'), claudeMdContent);
  console.log('Generado CLAUDE.md (con comandos mapeados)');

  // Archivo para Cursor / Windsurf
  const cursorRulesContent = `# Reglas de Sistema (Cursor / Windsurf)\n\n${baseContext}\n\n## Comportamiento del Agente\n- Si el usuario te pide ejecutar \`/feature\` o \`/implement\`, obedece las instrucciones listadas en la sección de Comandos usando la definición en \`.opencode/commands/\`.`;
  fs.writeFileSync(path.join(rootDir, '.cursorrules'), cursorRulesContent);
  fs.writeFileSync(path.join(rootDir, '.windsurfrules'), cursorRulesContent);
  console.log('Generado .cursorrules y .windsurfrules (con comandos mapeados)');

  // Archivo para Codex / Aider
  const conventionsContent = `# Convenciones Globales del Proyecto\n\n${baseContext}`;
  fs.writeFileSync(path.join(rootDir, 'CONVENTIONS.md'), conventionsContent);
  console.log('Generado CONVENTIONS.md');
  
  console.log('Sincronizacion completada exitosamente.');

} catch (error) {
  console.error('Error durante la sincronizacion:', error);
}