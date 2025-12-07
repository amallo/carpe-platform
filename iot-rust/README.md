# IoT Rust - ESP32 Project

Projet Rust pour ESP32 utilisant esp-rs.

## Prérequis

1. Installer esp-rs :
```bash
cargo install espup
espup install
```

2. Configurer l'environnement :
```bash
source $HOME/.espressif/rust/export-esp.sh
```

3. Ajouter le target Rust pour ESP32 :
```bash
rustup target add riscv32imc-unknown-none-elf  # Pour ESP32-C3
# ou pour ESP32 classique (Xtensa), esp-rs l'ajoutera automatiquement
```

## Compilation

```bash
cargo build --release
```

## Flash sur ESP32

```bash
espflash flash --monitor target/xtensa-esp32-espidf/release/iot-rust
```

## Note

Ce projet utilise cargo directement avec esp-rs. PlatformIO n'est pas nécessaire pour le développement Rust sur ESP32.
