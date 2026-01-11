# 🐟 CARPE Platform

Monorepo pour la plateforme CARPE - un système de messagerie décentralisé utilisant LoRa et BLE.

## 📁 Structure du Monorepo

```
carpe-platform/
├── packages/
│   ├── iot/         # Module ESP32 (firmware embarqué)
│   └── native/      # Application mobile React Native
├── package.json     # Configuration npm workspaces
└── package-lock.json
```

## 🎯 Projets

### `iot/` - CARPE Module

Firmware ESP32 implémentant l'authentification peer-to-peer sécurisée via BLE (Bluetooth Low Energy).

**Technologies :**
- C++17
- PlatformIO
- ESP32 (Arduino framework)
- NimBLE-Arduino

**Documentation :** Voir [`packages/iot/README.md`](packages/iot/README.md)

**Commandes principales :**
```bash
cd packages/iot
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

**Documentation :** Voir [`packages/native/README.md`](packages/native/README.md)

**Commandes principales :**
```bash
# Depuis la racine du monorepo
npm run native:start                 # Démarrer Metro bundler
npm run native:ios                   # Lancer sur iOS
npm run native:android               # Lancer sur Android

# Ou depuis packages/native/
cd packages/native
npm start                     # Démarrer Metro bundler
npm run ios                       # Lancer sur iOS
npm run android                   # Lancer sur Android
```

## 🚀 Getting Started

### Prérequis

**Pour `iot/` :**
- [PlatformIO](https://platformio.org/) installé
- ESP32 development board (testé avec TTGO LoRa32 v1)

**Pour `native/` :**
- Node.js >= 18
- npm >= 9 (gestionnaire de paquets pour le monorepo)
- React Native CLI
- Xcode (pour iOS)
- Android Studio (pour Android)

### Installation

```bash
# Cloner le monorepo
git clone <repository-url> carpe-platform
cd carpe-platform

# Installer toutes les dépendances du monorepo
npm install

# Installer les dépendances du module IoT
cd packages/iot
# PlatformIO installera automatiquement les dépendances au premier build

# Pour iOS uniquement
cd ../native/ios
pod install
cd ../../..
```

### Commandes du monorepo

Depuis la racine du monorepo, vous pouvez utiliser :

```bash
# Commandes pour native/ (app mobile)
npm run native:install      # Installer les dépendances
npm run native:start        # Démarrer Metro bundler
npm run native:ios          # Lancer sur iOS
npm run native:android      # Lancer sur Android
npm run native:test         # Lancer les tests

# Commandes pour iot/ (firmware ESP32)
npm run iot:build     # Build le firmware
npm run iot:upload    # Upload vers l'ESP32
npm run iot:test      # Lancer les tests natifs
```

## 🏗️ Architecture

### Communication Protocol

Les deux projets partagent le même protocole de communication binaire documenté dans [`packages/iot/protocol.md`](packages/iot/protocol.md).

### Séparation des responsabilités

- **`iot/`** : Gère l'authentification, le protocole BLE, et la communication LoRa
- **`native/`** : Interface utilisateur mobile, scan BLE, affichage des messages

## 📚 Documentation

- **Protocole de communication :** [`packages/iot/protocol.md`](packages/iot/protocol.md)
- **Module IoT :** [`packages/iot/README.md`](packages/iot/README.md)
- **Application mobile :** [`packages/native/README.md`](packages/native/README.md)
- **Whitepaper technique :** [`packages/native/CARPEAPP_WHITEPAPER.md`](packages/native/CARPEAPP_WHITEPAPER.md)

## 🔧 Développement

### Workflow recommandé

1. Développer et tester le firmware dans `iot/`
2. Tester l'intégration avec l'app dans `native/`
3. Itérer sur le protocole si nécessaire

### Tests

```bash
# Tests du module IoT (desktop, rapide)
cd packages/iot
make test

# Tests de l'app mobile
npm run native:test
# Ou depuis packages/native/
cd packages/native && npm test
```

## 📦 Gestion des dépendances

Ce monorepo utilise **npm workspaces** pour gérer les dépendances JavaScript/TypeScript. 

**Structure :**
- Les packages sont organisés dans le dossier `packages/`
- `packages/native/` : Application React Native (géré par npm)
- `packages/iot/` : Firmware Go/ESP32 (géré par PlatformIO)

## 📝 Notes

- Les deux projets sont organisés dans le dossier `packages/`
- Le protocole de communication est partagé entre les deux projets
- Les changements de protocole doivent être synchronisés entre `packages/iot/` et `packages/native/`
- Le monorepo utilise npm workspaces pour `packages/native/`, tandis que `packages/iot/` utilise PlatformIO/Go

## 🤝 Contribution

Chaque projet a ses propres guidelines de contribution. Voir les README respectifs.

---

**Motto :** "Exploration first. Outcome later."

