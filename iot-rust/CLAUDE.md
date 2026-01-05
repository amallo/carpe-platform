Project Posture (AI Rules)

This project follows a strict engineering posture. Any AI assistance must comply with the rules below.

General Mindset

Prefer explicit, simple, and deterministic designs over clever abstractions.

Optimize first for clarity, testability, and robustness, not for flexibility or performance.

Assume the target environment is embedded / constrained / no_std by default.

Treat this codebase as long-lived firmware, not a prototype.

Architecture Principles

Architecture is Event / Command / State, inspired by PLC scan cycles.

No implicit magic:

No hidden observers

No global event bus

No reflection-like behavior

Control flow must be traceable end-to-end.

State & Logic

State is explicit and modeled with enums and structs.

State transitions are deterministic and testable.

Reducers compute state only (no side effects).

Side effects are triggered explicitly via Commands / Effects.

Dependencies

All external interactions (BLE, LoRa, storage, time, queues, etc.) are dependencies, injected via traits.

No hard-coded singletons.

The runtime orchestrates execution but contains no domain logic.

Testing Philosophy

Core logic must be testable without async runtime, without hardware.

Tests favor step-by-step execution and inspection over mocks-heavy setups.

Determinism > realism in tests.

What to Avoid

Over-engineered patterns (CQRS frameworks, generic buses, actor systems).

Dynamic dispatch when static dispatch is sufficient.

Hidden background tasks or uncontrolled concurrency.

AI Behavior Rule

When in doubt:

Ask whether this abstraction is truly necessary for an embedded, offline-first, deterministic device.

If not clearly necessary, propose the simpler alternative.