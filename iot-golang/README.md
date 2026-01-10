# CARPE Module - ESP32 Go

Projet ESP32 développé en Go avec TinyGo pour le module CARPE.

## 🎯 Vue d'ensemble

Ce projet implémente le firmware ESP32 en Go pour le système d'authentification peer-to-peer sécurisé via BLE (Bluetooth Low Energy) du module CARPE.

## 📋 Prérequis

### Outils de développement (macOS uniquement)

Sur macOS, vous devez installer les outils de développement en ligne de commande de Xcode.

**Sur macOS Sequoia (15.0+)**, Xcode seul n'est pas suffisant. Vous devez installer séparément les Command Line Tools:

1. **Méthode automatique** (ouvre une fenêtre de dialogue):
   ```bash
   xcode-select --install
   ```

2. **Méthode manuelle** (si la méthode automatique ne fonctionne pas):
   - Téléchargez depuis [developer.apple.com/downloads](https://developer.apple.com/downloads/)
   - Recherchez "Command Line Tools for Xcode"
   - Installez le package `.dmg`

**Note**: Sur macOS Sequoia, même si Xcode est installé, les Command Line Tools doivent être installés séparément.

### Installation de TinyGo

1. **macOS** (avec Homebrew):
   ```bash
   brew tap tinygo-org/tools
   brew install tinygo
   ```

2. **Linux**:
   ```bash
   wget https://github.com/tinygo-org/tinygo/releases/download/v0.31.0/tinygo0.31.0.linux-amd64.tar.gz
   tar -xzf tinygo0.31.0.linux-amd64.tar.gz
   export PATH=$PATH:$(pwd)/tinygo/bin
   ```

3. **Windows**: Télécharger depuis [tinygo.org](https://tinygo.org/getting-started/install/)

### Installation d'esptool.py

```bash
pip install esptool
```

### Vérification de l'installation

```bash
tinygo version
esptool.py version
```

## 🚀 Utilisation

### Compilation

```bash
make build
```

Cela génère `firmware.bin` prêt à être flashé sur l'ESP32.

### Flash sur l'ESP32

```bash
# Utiliser le port série par défaut (/dev/ttyUSB0)
make flash

# Ou spécifier un port différent
make flash SERIAL_PORT=/dev/tty.usbserial-*
```

### Monitorer le port série

```bash
make monitor SERIAL_PORT=/dev/ttyUSB0
```

### Compilation et flash en une commande

```bash
make all SERIAL_PORT=/dev/ttyUSB0
```

## 🏗️ Structure du projet

```
iot-golang/
├── main.go          # Point d'entrée principal
├── go.mod           # Dépendances Go
├── Makefile         # Commandes de build
├── README.md        # Documentation
└── .gitignore       # Fichiers ignorés
```

## 🔧 Configuration

### Cible TinyGo

Le projet utilise la cible `esp32-coreboard-v2` par défaut. Pour utiliser une autre carte ESP32, modifiez la cible dans le `Makefile`:

- `esp32-coreboard-v2` - ESP32 Core Board V2
- `esp32` - ESP32 générique
- `esp32c3` - ESP32-C3
- `esp32s2` - ESP32-S2

### Port série

Définissez la variable `SERIAL_PORT` pour spécifier le port série:

```bash
export SERIAL_PORT=/dev/ttyUSB0
make flash
```

## 📦 Dépendances

Les dépendances sont gérées via `go.mod`. Pour ajouter une nouvelle dépendance:

```bash
go get tinygo.org/x/drivers@latest
```

## 🧪 Tests

Le projet utilise le package `testing` standard de Go. Les tests sont exécutés localement sur la machine de développement avec `go test` et sont exclus du build firmware ESP32 grâce au build tag `!tinygo`.

**Note**: On utilise `go test` plutôt que `tinygo test` car `tinygo test` respecte aussi le build tag `!tinygo` et ne verrait donc pas les fichiers de test. Le build tag `!tinygo` garantit que les tests ne sont jamais inclus dans le build firmware avec `tinygo build`.

### Exécuter les tests

```bash
# Tous les tests (utilise go test)
make test

# Tests avec détails
make test-verbose

# Tests avec couverture de code
make test-coverage
```

Ou directement avec Go:

```bash
go test ./core/...
go test ./core/... -v
go test ./core/... -cover
```

### Structure des tests

Les tests suivent la convention Go standard:
- Fichiers de test: `*_test.go` avec build tag `//go:build !tinygo`
- Tests unitaires: fonctions `TestXxx(t *testing.T)`
- Les tests sont host-only et ne sont jamais inclus dans le build firmware

### Exemples de tests

Le projet inclut des exemples de tests dans:
- `core/device/device_test.go` - Tests unitaires du device
- `core/config/config_test.go` - Tests avec mocks manuels

### Utilisation du package testing standard

**Tests simples:**
```go
//go:build !tinygo

package device

import "testing"

func TestNewDevice(t *testing.T) {
    device := NewDevice("test-001", "Test Device")
    
    if device == nil {
        t.Fatal("NewDevice returned nil")
    }
    
    if got := device.GetID(); got != "test-001" {
        t.Errorf("GetID() = %q, want %q", got, "test-001")
    }
}
```

**Gestion d'erreurs:**
```go
func TestBoot(t *testing.T) {
    device := NewDevice("test", "Test")
    err := device.Boot()
    
    if err != nil {
        t.Fatalf("Boot() returned error: %v", err)
    }
}
```

**Mocks manuels:**
```go
type MockStorage struct {
    storage map[string]string
}

func (m *MockStorage) Get(key string) (string, error) {
    value, ok := m.storage[key]
    if !ok {
        return "", errors.New("key not found")
    }
    return value, nil
}
```

### Séparation host/embarqué

- **Tests**: Exécutés uniquement sur la machine de développement avec `tinygo test`
- **Build firmware**: Les fichiers `*_test.go` sont exclus grâce au build tag `!tinygo`
- **Avantage**: Les tests peuvent utiliser le package `testing` standard sans dépendances externes

## 🧪 Développement

### Formatage du code

```bash
go fmt ./...
```

### Vérification statique

```bash
go vet ./...
```

### Build avec optimisations

Le Makefile utilise déjà `-opt=z` pour optimiser la taille. Pour d'autres options:

- `-opt=0` - Pas d'optimisation (debug)
- `-opt=s` - Optimisation pour la taille
- `-opt=z` - Optimisation agressive pour la taille (défaut)

## 🔍 Debugging

### Affichage de la taille du binaire

Le Makefile utilise `-size short` pour afficher la taille des sections:

```
   code  data     bss |   flash     ram
  1234   567     890 |   1801    1457
```

### Moniteur série

Utilisez `make monitor` pour voir les sorties `println()` et les logs.

## 📚 Ressources

- [Documentation TinyGo](https://tinygo.org/docs/)
- [TinyGo Drivers](https://github.com/tinygo-org/drivers)
- [ESP32 avec TinyGo](https://tinygo.org/microcontrollers/esp32/)
- [Protocol CARPE](../iot/protocol.md)

## 🤝 Contribution

Ce projet fait partie de la plateforme CARPE. Voir le README principal pour les guidelines de contribution.

## 📄 License

[À définir]

