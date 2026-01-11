package config

// ConfigStorage définit l'interface pour le stockage de configuration
type ConfigStorage interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
}

// ConfigManager gère la configuration
type ConfigManager struct {
	storage ConfigStorage
}

// NewConfigManager crée un nouveau gestionnaire de configuration
func NewConfigManager(storage ConfigStorage) *ConfigManager {
	return &ConfigManager{storage: storage}
}

// GetConfig récupère une valeur de configuration
func (cm *ConfigManager) GetConfig(key string) (string, error) {
	return cm.storage.Get(key)
}

// SetConfig définit une valeur de configuration
func (cm *ConfigManager) SetConfig(key, value string) error {
	return cm.storage.Set(key, value)
}

