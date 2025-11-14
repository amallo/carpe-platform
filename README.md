# 🐟 CARPE Platform

Monorepo pour la plateforme CARPE - un système de messagerie décentralisé utilisant LoRa et BLE.

## 📁 Structure du Monorepo

```
carpe-platform/
├── iot/             # Module ESP32 (firmware embarqué)
└── native/          # Application mobile React Native
```

## 🎯 Projets

### `iot/` - CARPE Module

Firmware ESP32 implémentant l'authentification peer-to-peer sécurisée via BLE (Bluetooth Low Energy).

**Technologies :**
- C++17
- PlatformIO
- ESP32 (Arduino framework)
- NimBLE-Arduino

**Documentation :** Voir [`iot/README.md`](iot/README.md)

**Commandes principales :**
```bash
cd iot
pio run -e carpe-lora          # Build
pio run -e carpe-lora --target upload  # Upload
make test                      # Tests natifs
```

### `native/` - CARPE App

Application mobile React Native pour iOS et Android permettant de se connecter aux modules LoRa/BLE.

**Technologies :**
- React Native 0.78.1
- TypeScript
- Redux Toolkit
- React Navigation

**Documentation :** Voir [`native/README.md`](native/README.md)

**Commandes principales :**
```bash
# Depuis la racine du monorepo
pnpm native:start                 # Démarrer Metro bundler
pnpm native:ios                   # Lancer sur iOS
pnpm native:android               # Lancer sur Android

# Ou depuis native/
cd native
pnpm start                     # Démarrer Metro bundler
pnpm ios                       # Lancer sur iOS
pnpm android                   # Lancer sur Android
```

## 🚀 Getting Started

### Prérequis

**Pour `iot/` :**
- [PlatformIO](https://platformio.org/) installé
- ESP32 development board (testé avec TTGO LoRa32 v1)

**Pour `native/` :**
- Node.js >= 18
- pnpm >= 8 (gestionnaire de paquets recommandé pour le monorepo)
- React Native CLI
- Xcode (pour iOS)
- Android Studio (pour Android)

### Installation

```bash
# Cloner le monorepo
git clone <repository-url> carpe-platform
cd carpe-platform

# Installer pnpm si ce n'est pas déjà fait
npm install -g pnpm

# Installer toutes les dépendances du monorepo
pnpm install

# Installer les dépendances du module IoT
cd iot
# PlatformIO installera automatiquement les dépendances au premier build

# Pour iOS uniquement
cd ../native/ios
pod install
cd ../..
```

### Commandes du monorepo

Depuis la racine du monorepo, vous pouvez utiliser :

```bash
# Commandes pour native/ (app mobile)
pnpm native:install      # Installer les dépendances
pnpm native:start        # Démarrer Metro bundler
pnpm native:ios          # Lancer sur iOS
pnpm native:android      # Lancer sur Android
pnpm native:test         # Lancer les tests

# Commandes pour iot/ (firmware ESP32)
pnpm iot:build     # Build le firmware
pnpm iot:upload    # Upload vers l'ESP32
pnpm iot:test      # Lancer les tests natifs
```

## 🏗️ Architecture

### Communication Protocol

Les deux projets partagent le même protocole de communication binaire documenté dans [`iot/protocol.md`](iot/protocol.md).

### Séparation des responsabilités

- **`iot/`** : Gère l'authentification, le protocole BLE, et la communication LoRa
- **`native/`** : Interface utilisateur mobile, scan BLE, affichage des messages

## 📚 Documentation

- **Protocole de communication :** [`iot/protocol.md`](iot/protocol.md)
- **Module IoT :** [`iot/README.md`](iot/README.md)
- **Application mobile :** [`native/README.md`](native/README.md)
- **Whitepaper technique :** [`native/CARPEAPP_WHITEPAPER.md`](native/CARPEAPP_WHITEPAPER.md)

## 🔧 Développement

### Workflow recommandé

1. Développer et tester le firmware dans `iot/`
2. Tester l'intégration avec l'app dans `native/`
3. Itérer sur le protocole si nécessaire

### Tests

```bash
# Tests du module IoT (desktop, rapide)
cd iot
make test

# Tests de l'app mobile
pnpm native:test
# Ou depuis native/
cd native && pnpm test
```

## 📦 Gestion des dépendances

Ce monorepo utilise **pnpm workspaces** pour gérer les dépendances JavaScript/TypeScript. 

**Avantages de pnpm :**
- ⚡ Plus rapide que npm/yarn
- 💾 Plus efficace en espace disque (liens symboliques)
- 🎯 Meilleur pour les monorepos
- 🔒 Installation plus sécurisée (pas de dépendances fantômes)

**Migration :** Le projet a été migré de npm vers pnpm. Voir [`MIGRATION_PNPM.md`](MIGRATION_PNPM.md) pour les détails.

## 📝 Notes

- Les deux projets sont des repositories Git indépendants
- Le protocole de communication est partagé entre les deux projets
- Les changements de protocole doivent être synchronisés entre `iot/` et `native/`
- Le monorepo utilise pnpm workspaces pour `native/`, tandis que `iot/` utilise PlatformIO

## 🤝 Contribution

Chaque projet a ses propres guidelines de contribution. Voir les README respectifs.

---

**Motto :** "Exploration first. Outcome later."

