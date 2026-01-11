.PHONY: build flash monitor clean help test test-verbose test-coverage

# Port série par défaut (modifier selon votre configuration)
SERIAL_PORT ?= /dev/ttyUSB0

# Cible par défaut
help:
	@echo "Commandes disponibles:"
	@echo "  make build         - Compiler le firmware pour ESP32"
	@echo "  make flash         - Flasher le firmware sur l'ESP32"
	@echo "  make monitor       - Ouvrir le moniteur série"
	@echo "  make test          - Exécuter les tests unitaires"
	@echo "  make test-verbose  - Exécuter les tests avec détails"
	@echo "  make test-coverage - Générer un rapport de couverture"
	@echo "  make clean         - Nettoyer les fichiers de build"
	@echo ""
	@echo "Variables:"
	@echo "  SERIAL_PORT        - Port série (défaut: /dev/ttyUSB0)"

# Compiler le firmware
# Note: Les fichiers *_test.go sont automatiquement exclus du build grâce au build tag !tinygo
build:
	@echo "Compilation du firmware ESP32..."
	tinygo build -size short -o firmware.bin -target=esp32-coreboard-v2 -opt=z ./main.go
	@echo "Compilation terminée: firmware.bin"

# Flasher sur l'ESP32 (méthode 1: avec tinygo flash - plus simple)
flash:
	@echo "Flash du firmware sur ESP32..."
	tinygo flash -size short -target=esp32-coreboard-v2 -port=$(SERIAL_PORT) -opt=z ./main.go
	@echo "Flash terminé"

# Flasher avec esptool.py (méthode 2: si tinygo flash ne fonctionne pas)
flash-esptool:
	@echo "Flash du firmware sur ESP32 avec esptool..."
	@if [ ! -f firmware.bin ]; then \
		echo "Erreur: firmware.bin introuvable. Exécutez 'make build' d'abord."; \
		exit 1; \
	fi
	esptool.py --chip esp32 --port $(SERIAL_PORT) write_flash -z 0x1000 firmware.bin
	@echo "Flash terminé"

# Monitorer le port série
monitor:
	@echo "Ouverture du moniteur série sur $(SERIAL_PORT)..."
	@echo "Appuyez sur Ctrl+] pour quitter"
	tinygo monitor -port $(SERIAL_PORT)

# Nettoyer les fichiers de build
clean:
	rm -f firmware.bin firmware.elf
	@echo "Nettoyage terminé"

# Exécuter les tests unitaires (host-only, utilise go test)
# Les tests sont exécutés localement sur la machine de développement avec go test
# Le build tag !tinygo garantit que les fichiers de test sont exclus du build firmware
# Note: On utilise go test plutôt que tinygo test car tinygo test respecte aussi
#       le build tag !tinygo et ne voit donc pas les fichiers de test
test:
	@echo "Exécution des tests (host-only)..."
	go test ./core/... -v

# Exécuter les tests avec plus de détails
test-verbose:
	@echo "Exécution des tests (mode verbeux)..."
	go test ./core/... -v -count=1

# Générer un rapport de couverture
test-coverage:
	@echo "Génération du rapport de couverture..."
	go test ./core/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Rapport généré: coverage.html"

# Compiler et flasher en une commande
all: build flash

