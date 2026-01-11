# Guide de Tests avec Go

Ce projet utilise le package `testing` standard de Go. Les tests sont exécutés localement sur la machine de développement avec `go test` et sont exclus du build firmware ESP32 grâce au build tag `!tinygo`.

**Note**: On utilise `go test` plutôt que `tinygo test` car `tinygo test` respecte aussi le build tag `!tinygo` et ne verrait donc pas les fichiers de test. Le build tag `!tinygo` garantit que les tests ne sont jamais inclus dans le build firmware avec `tinygo build`.

## Exécution des tests

```bash
# Tous les tests (utilise go test)
make test

# Tests avec détails
make test-verbose

# Tests avec couverture de code
make test-coverage

# Tests d'un package spécifique
go test ./core/device/...

# Tests avec sortie détaillée
go test ./core/... -v
```

## Structure des tests

### Build tags

Tous les fichiers de test doivent commencer par le build tag `!tinygo` pour être exclus du build firmware:

```go
//go:build !tinygo

package device

import "testing"
```

### Convention de nommage

- Fichiers de test: `*_test.go`
- Fonctions de test: `TestXxx(t *testing.T)` où `Xxx` commence par une majuscule
- Tests de benchmark: `BenchmarkXxx(b *testing.B)`

## Exemples de tests

### Tests unitaires simples

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
    
    if got := device.GetName(); got != "Test Device" {
        t.Errorf("GetName() = %q, want %q", got, "Test Device")
    }
}
```

### Gestion d'erreurs

```go
func TestBoot(t *testing.T) {
    device := NewDevice("test", "Test")
    err := device.Boot()
    
    if err != nil {
        t.Fatalf("Boot() returned error: %v", err)
    }
    
    if got := device.GetStatus(); got != StatusOnline {
        t.Errorf("GetStatus() = %q, want %q", got, StatusOnline)
    }
}
```

### Tests avec sous-tests

```go
func TestStatusTransitions(t *testing.T) {
    device := NewDevice("test", "Test")
    
    t.Run("initial state", func(t *testing.T) {
        if got := device.GetStatus(); got != StatusOffline {
            t.Errorf("GetStatus() = %q, want %q", got, StatusOffline)
        }
    })
    
    t.Run("after boot", func(t *testing.T) {
        err := device.Boot()
        if err != nil {
            t.Fatalf("Boot() returned error: %v", err)
        }
        if got := device.GetStatus(); got != StatusOnline {
            t.Errorf("GetStatus() = %q, want %q", got, StatusOnline)
        }
    })
}
```

### Mocks manuels

Pour créer des mocks sans dépendances externes:

```go
//go:build !tinygo

package config

import (
    "errors"
    "testing"
)

// MockConfigStorage est un mock manuel pour ConfigStorage
type MockConfigStorage struct {
    storage map[string]string
    errors  map[string]error
}

func NewMockConfigStorage() *MockConfigStorage {
    return &MockConfigStorage{
        storage: make(map[string]string),
        errors:  make(map[string]error),
    }
}

func (m *MockConfigStorage) SetError(key string, err error) {
    m.errors[key] = err
}

func (m *MockConfigStorage) Get(key string) (string, error) {
    if err, ok := m.errors[key]; ok {
        return "", err
    }
    value, ok := m.storage[key]
    if !ok {
        return "", errors.New("key not found")
    }
    return value, nil
}

func (m *MockConfigStorage) Set(key, value string) error {
    if err, ok := m.errors[key]; ok {
        return err
    }
    m.storage[key] = value
    return nil
}

func TestConfigManagerWithMock(t *testing.T) {
    mockStorage := NewMockConfigStorage()
    mockStorage.Set("device_id", "test-123")
    configManager := NewConfigManager(mockStorage)
    
    value, err := configManager.GetConfig("device_id")
    if err != nil {
        t.Fatalf("GetConfig() returned error: %v", err)
    }
    
    if value != "test-123" {
        t.Errorf("GetConfig() = %q, want %q", value, "test-123")
    }
}
```

## Méthodes de testing.T

### Signaler des erreurs

- `t.Error(args...)` - Signale une erreur mais continue le test
- `t.Errorf(format, args...)` - Signale une erreur avec formatage
- `t.Fatal(args...)` - Signale une erreur et arrête le test
- `t.Fatalf(format, args...)` - Signale une erreur avec formatage et arrête le test

### Exemples d'assertions

```go
// Égalité
if got != want {
    t.Errorf("got %q, want %q", got, want)
}

// Nil
if obj == nil {
    t.Fatal("obj is nil")
}

// Erreurs
if err != nil {
    t.Fatalf("unexpected error: %v", err)
}

// Booléens
if !condition {
    t.Error("condition is false")
}

// Contenus (pour strings)
if !strings.Contains(str, substr) {
    t.Errorf("%q does not contain %q", str, substr)
}

// Longueurs (pour slices)
if len(slice) != expected {
    t.Errorf("len(slice) = %d, want %d", len(slice), expected)
}
```

## Bonnes pratiques

1. **Un test, une responsabilité**: Chaque test doit vérifier une seule chose
2. **Noms descriptifs**: `TestDeviceBoot_SetsStatusToOnline` plutôt que `TestBoot`
3. **Tests indépendants**: Les tests ne doivent pas dépendre les uns des autres
4. **Mocks pour les dépendances**: Utilisez des mocks manuels pour isoler les tests unitaires
5. **Build tags**: Toujours ajouter `//go:build !tinygo` en première ligne des fichiers de test
6. **Messages d'erreur clairs**: Utilisez `t.Errorf` avec des messages descriptifs incluant les valeurs attendues et obtenues

## Séparation host/embarqué

### Tests (host-only)

- Exécutés uniquement sur la machine de développement
- Utilisent `go test` pour l'exécution (pas `tinygo test` car il respecte aussi le build tag `!tinygo`)
- Peuvent utiliser le package `testing` standard complet
- Exclus du build firmware grâce au build tag `!tinygo`

### Build firmware

- Les fichiers `*_test.go` sont automatiquement exclus lors du build
- Le build tag `!tinygo` garantit l'exclusion explicite
- Aucun code de test n'est inclus dans le firmware final

## Exemples dans le projet

- `core/device/device_test.go` - Tests unitaires du device
- `core/config/config_test.go` - Tests avec mocks manuels

## Ressources

- [Documentation Go testing](https://pkg.go.dev/testing)
- [Go Testing Best Practices](https://golang.org/doc/effective_go#testing)
- [TinyGo Documentation](https://tinygo.org/docs/)
