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
	@bash -c "source $$HOME/export-esp.sh && cargo build --release --target xtensa-esp32-none-elf"

dev: ## Compile le projet en mode debug (plus rapide)
	@echo "$(GREEN)🔨 Compilation du projet (debug)...$(NC)"
	@bash -c "source $$HOME/export-esp.sh && cargo build --target xtensa-esp32-none-elf"

flash: build ## Flash le firmware sur l'ESP32
	@echo "$(GREEN)📤 Flash du firmware sur ESP32...$(NC)"
	@espflash flash --monitor $(TARGET_DIR)/$(BUILD_MODE)/$(BINARY)

monitor: ## Ouvre le moniteur série
	@echo "$(GREEN)📺 Ouverture du moniteur série...$(NC)"
	@espflash monitor

clean: ## Nettoie les fichiers de build
	@echo "$(GREEN)🧹 Nettoyage...$(NC)"
	@cargo clean

test: ## Lance les tests unitaires sur Mac (host)
	@echo "$(GREEN)🧪 Lancement des tests unitaires...$(NC)"
	@echo "$(YELLOW)Note: Les tests dans #[cfg(test)] sont automatiquement exclus du build release.$(NC)"
	@echo "$(YELLOW)Pour tester du code dépendant d'ESP32, utilisez des abstractions/traits.$(NC)"
	@echo ""
	@echo "$(GREEN)Exécution des tests dans src/lib.rs (sans dépendances ESP32)...$(NC)"
	@rustup override unset 2>/dev/null || true
	@RUSTFLAGS="" CARGO_BUILD_RUSTFLAGS="" cargo +stable test --lib --no-default-features 2>&1 | tail -30
	@rustup override set esp 2>/dev/null || true

check: ## Vérifie le code sans compiler
	@echo "$(GREEN)✓ Vérification du code...$(NC)"
	@bash -c "source $$HOME/export-esp.sh && cargo check --target xtensa-esp32-none-elf"

clippy: ## Lance clippy (linter Rust)
	@echo "$(GREEN)🔍 Analyse avec clippy...$(NC)"
	@bash -c "source $$HOME/export-esp.sh && cargo clippy --target xtensa-esp32-none-elf -- -D warnings"

fmt: ## Formate le code avec rustfmt
	@echo "$(GREEN)✨ Formatage du code...$(NC)"
	@cargo fmt

all: clean build ## Nettoie et compile le projet
