// @ts-check
import { defineConfig } from 'astro/config';
import tailwindcss from '@tailwindcss/vite';

// https://astro.build/config
export default defineConfig({
  vite: {
    plugins: [tailwindcss()],
    server: {
      allowedHosts: true, // Permite cualquier host (como ngrok) en modo desarrollo
      proxy: {
        '/ws': {
          target: 'ws://localhost:8080',
          ws: true,
          rewrite: (path) => path
        }
      }
    }
  }
});