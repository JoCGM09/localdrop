.PHONY: dev dev-tunnel backend frontend tunnel

# Objetivo por defecto si solo se escribe "make"
help:
	@echo "Comandos disponibles:"
	@echo "  make dev         - Inicia el backend (Go) y frontend (Astro) localmente"
	@echo "  make dev-tunnel  - Inicia backend, frontend y abre un tunel con Ngrok"
	@echo "  make backend     - Inicia solo el backend"
	@echo "  make frontend    - Inicia solo el frontend"

backend:
	@echo "Iniciando Backend..."
	cd backend && go run main.go

frontend:
	@echo "Iniciando Frontend..."
	cd frontend && npm run dev

tunnel:
	@echo "Iniciando Tunnel Ngrok..."
	cd frontend && npm run dev:tunnel

# Ejecuta backend y frontend en paralelo usando -j (jobs)
dev:
	@echo "Iniciando LocalDrop en http://localhost:4321..."
	@$(MAKE) -j 2 backend frontend

# Ejecuta backend, frontend y ngrok en paralelo
dev-tunnel:
	@echo "Iniciando LocalDrop con tunel publico..."
	@$(MAKE) -j 3 backend frontend tunnel
