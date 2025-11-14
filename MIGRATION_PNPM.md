# Migration vers pnpm - Guide

## ✅ Migration terminée

Le monorepo a été migré avec succès de npm vers **pnpm workspaces**.

## 📦 Changements effectués

1. ✅ Suppression de `iot/package-lock.json`
2. ✅ Création de `pnpm-lock.yaml` à la racine
3. ✅ Configuration des workspaces dans `package.json`
4. ✅ Création de `pnpm-workspace.yaml`
5. ✅ Configuration de `.npmrc` pour React Native
6. ✅ Nettoyage des anciens packages npm

## 🚀 Utilisation

### Commandes depuis la racine

```bash
# Installer toutes les dépendances
pnpm install

# Commandes pour iot/
pnpm iot:start        # Démarrer Metro bundler
pnpm iot:ios          # Lancer sur iOS
pnpm iot:android      # Lancer sur Android
pnpm iot:test         # Lancer les tests

# Commandes pour native/
pnpm native:build     # Build le firmware
pnpm native:upload    # Upload vers l'ESP32
pnpm native:test      # Tests natifs
```

### Commandes depuis iot/

Vous pouvez toujours utiliser pnpm directement dans `iot/` :

```bash
cd iot
pnpm start
pnpm ios
pnpm android
pnpm test
```

## 📝 Notes importantes

- **Ne plus utiliser npm** dans `iot/` - utilisez `pnpm` à la place
- Le fichier `pnpm-lock.yaml` doit être commité dans Git
- Les `node_modules` sont maintenant gérés par pnpm (liens symboliques)
- Si vous avez des problèmes, supprimez `node_modules` et `pnpm-lock.yaml`, puis relancez `pnpm install`

## 🔧 Dépannage

### Réinstaller les dépendances

```bash
# Depuis la racine
rm -rf node_modules iot/node_modules pnpm-lock.yaml
pnpm install
```

### Vérifier l'installation

```bash
# Lister les packages installés
pnpm --filter carpe-app list

# Vérifier les workspaces
pnpm -r exec pwd
```

## 📚 Ressources

- [Documentation pnpm](https://pnpm.io/)
- [pnpm workspaces](https://pnpm.io/workspaces)
- [React Native avec pnpm](https://reactnative.dev/docs/environment-setup)

