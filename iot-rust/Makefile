.PHONY: build flash monitor clean test help

# Configuration
BINARY = iot-rust
BUILD_MODE = release
TARGET_DIR = target/xtensa-esp32-none-elf

# Couleurs pour les messages
GREEN = \033[0;32m
YELLOW = \033[1;33m
NC = \033[0m # No Color

help: ## Affiche cette aide
	@echo "$(GREEN)Commandes disponibles:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'

build: ## Compile le projet en mode release
	@echo "$(GREEN)🔨 Compilation du projet...$(NC)"
	@bash -c "source $$HOME/export-esp.sh && cargo build --release"

dev: ## Compile le projet en mode debug (plus rapide)
	@echo "$(GREEN)🔨 Compilation du projet (debug)...$(NC)"
	@bash -c "source $$HOME/export-esp.sh && cargo build"

flash: build ## Flash le firmware sur l'ESP32
	@echo "$(GREEN)📤 Flash du firmware sur ESP32...$(NC)"
	@espflash flash --monitor $(TARGET_DIR)/$(BUILD_MODE)/$(BINARY)

monitor: ## Ouvre le moniteur série
	@echo "$(GREEN)📺 Ouverture du moniteur série...$(NC)"
	@espflash monitor

clean: ## Nettoie les fichiers de build
	@echo "$(GREEN)🧹 Nettoyage...$(NC)"
	@cargo clean

test: ## Lance les tests (si disponibles)
	@echo "$(GREEN)🧪 Lancement des tests...$(NC)"
	@bash -c "source $$HOME/export-esp.sh && cargo test"

check: ## Vérifie le code sans compiler
	@echo "$(GREEN)✓ Vérification du code...$(NC)"
	@bash -c "source $$HOME/export-esp.sh && cargo check"

clippy: ## Lance clippy (linter Rust)
	@echo "$(GREEN)🔍 Analyse avec clippy...$(NC)"
	@bash -c "source $$HOME/export-esp.sh && cargo clippy -- -D warnings"

fmt: ## Formate le code avec rustfmt
	@echo "$(GREEN)✨ Formatage du code...$(NC)"
	@cargo fmt

all: clean build ## Nettoie et compile le projet
