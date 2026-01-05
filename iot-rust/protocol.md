## CARPE Messaging Protocol v0 – Authentication (BLE)

This document defines the minimal binary exchange between the App and the Device for authentication challenge initiation.

### Scope (v0)
- App ↔ Device over BLE characteristic write/notify
- Authentication challenge initiation
- User messages (public and private)

### Conventions
- Byte order: big-endian (network order) for multi-byte integers.
- All numeric values are unsigned unless stated otherwise.
- All fields are mandatory unless stated otherwise.

### Message Frame
Each message is a concatenation of HEADER then PAYLOAD.

```
[HEADER][PAYLOAD]
```

#### HEADER (3 bytes)
- TYPE (1 byte): Message type code
- NONCE (2 bytes): Anti-replay value (big-endian, see Security)

#### PAYLOAD (variable)
Depends on TYPE. See per-message tables below.

### Message Types
| TYPE | Name                              | Direction      | Description |
|------|-----------------------------------|----------------|-------------|
| 0x04 | INITIATE_AUTH_CHALLENGE           | Device → App   | Initiate authentication challenge with challengeId |
| 0x05 | AUTH_CHALLENGE_NEGOTIATION_SUCCESS| Device → App   | Authentication challenge negotiation succeeded with sessionId |
| 0x06 | AUTH_CHALLENGE_NEGOTIATION_FAILURE| Device → App   | Authentication challenge negotiation failed with reason and remaining attempts |
| 0x07 | USER_MESSAGE                      | App → Device   | User message (public or private) with senderId, recipientId, sessionId, and content |

---

### INITIATE_AUTH_CHALLENGE (TYPE = 0x04)

Header
| Field | Size | Notes |
|-------|------|-------|
| TYPE  | 1    | 0x04 |
| NONCE | 2    | Random per message (Device generated) |

Payload
| Field        | Size | Type | Description |
|--------------|------|------|-------------|
| CHALLENGE_ID | 16   | bytes| Challenge identifier (fixed size, null-padded if shorter) |

Example (challengeId="challenge-123", NONCE=0xAABB):
```
04 AA BB 63 68 61 6C 6C 65 6E 67 65 2D 31 32 33 00 00 00
```
Where:
- `04` is TYPE
- `AA BB` is NONCE (big-endian)
- `63 68 61 6C 6C 65 6E 67 65 2D 31 32 33` is "challenge-123" (13 bytes)
- `00 00 00` is padding to reach 16 bytes

---

### AUTH_CHALLENGE_NEGOTIATION_SUCCESS (TYPE = 0x05)

Header
| Field | Size | Notes |
|-------|------|-------|
| TYPE  | 1    | 0x05 |
| NONCE | 2    | Random per message (Device generated) |

Payload
| Field        | Size | Type | Description |
|--------------|------|------|-------------|
| SESSION_ID   | 16   | bytes| Session identifier (fixed size, null-padded if shorter) |
| CHALLENGE_ID | 16   | bytes| Challenge identifier (fixed size, null-padded if shorter) |

Example (sessionId="session-1", challengeId="challenge-123", NONCE=0xCCDD):
```
05 CC DD 73 65 73 73 69 6F 6E 2D 31 00 00 00 00 00 63 68 61 6C 6C 65 6E 67 65 2D 31 32 33 00 00 00
```
Where:
- `05` is TYPE
- `CC DD` is NONCE (big-endian)
- `73 65 73 73 69 6F 6E 2D 31` is "session-1" (9 bytes) + 7 bytes padding
- `63 68 61 6C 6C 65 6E 67 65 2D 31 32 33` is "challenge-123" (13 bytes) + 3 bytes padding

---

### AUTH_CHALLENGE_NEGOTIATION_FAILURE (TYPE = 0x06)

Header
| Field | Size | Notes |
|-------|------|-------|
| TYPE  | 1    | 0x06 |
| NONCE | 2    | Random per message (Device generated) |

Payload
| Field              | Size | Type | Description |
|--------------------|------|------|-------------|
| CHALLENGE_ID       | 16   | bytes| Challenge identifier (fixed size, null-padded if shorter) |
| REASON             | 32   | bytes| Failure reason (fixed size, null-padded if shorter) |
| REMAINING_ATTEMPTS | 1    | byte | Number of remaining attempts (-1 if challenge ended, 0-255 otherwise) |

Example (challengeId="challenge-123", reason="Invalid PIN", remainingAttempts=2, NONCE=0xEEFF):
```
06 EE FF 63 68 61 6C 6C 65 6E 67 65 2D 31 32 33 00 00 00 49 6E 76 61 6C 69 64 20 50 49 4E 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 02
```
Where:
- `06` is TYPE
- `EE FF` is NONCE (big-endian)
- `63 68 61 6C 6C 65 6E 67 65 2D 31 32 33` is "challenge-123" (13 bytes) + 3 bytes padding
- `49 6E 76 61 6C 69 64 20 50 49 4E` is "Invalid PIN" (11 bytes) + 21 bytes padding
- `02` is remainingAttempts (2)

---

### USER_MESSAGE (TYPE = 0x07)

Header
| Field        | Size | Notes |
|--------------|------|-------|
| TYPE         | 1    | 0x07 |
| NONCE        | 2    | Random per message (App generated) |
| SENDER_ID    | 16   | bytes| Sender identifier (fixed size, null-padded if shorter) |
| RECIPIENT_ID | 16   | bytes| Recipient identifier (fixed size, null-padded if shorter). `0x00...00` for public messages |

Payload
| Field      | Size | Type | Description |
|------------|------|------|-------------|
| SESSION_ID | 16   | bytes| Session identifier (fixed size, null-padded if shorter) |
| CONTENT    | 128  | bytes| Message content (fixed size, null-padded if shorter) |




### Validation Rules (Receiver)
1. Verify header size:
   - For INITIATE_AUTH_CHALLENGE, AUTH_CHALLENGE_NEGOTIATION_SUCCESS, AUTH_CHALLENGE_NEGOTIATION_FAILURE: 3 bytes (TYPE + NONCE)
   - For USER_MESSAGE: 35 bytes (TYPE + NONCE + SENDER_ID + RECIPIENT_ID)
2. Verify payload size matches expected size for TYPE:
   - INITIATE_AUTH_CHALLENGE: 16 bytes
   - AUTH_CHALLENGE_NEGOTIATION_SUCCESS: 32 bytes (16 + 16)
   - AUTH_CHALLENGE_NEGOTIATION_FAILURE: 49 bytes (16 + 32 + 1)
   - USER_MESSAGE: 144 bytes (16 + 128)
3. Check anti-replay: NONCE not seen recently for this peer/session.
4. For INITIATE_AUTH_CHALLENGE: Extract and validate CHALLENGE_ID (16 bytes).
5. For AUTH_CHALLENGE_NEGOTIATION_SUCCESS: Extract and validate SESSION_ID (16 bytes) and CHALLENGE_ID (16 bytes).
6. For AUTH_CHALLENGE_NEGOTIATION_FAILURE: Extract and validate CHALLENGE_ID (16 bytes), REASON (32 bytes), and REMAINING_ATTEMPTS (1 byte).
7. For USER_MESSAGE: Extract and validate SENDER_ID (16 bytes in header), RECIPIENT_ID (16 bytes in header), SESSION_ID (16 bytes), and CONTENT (128 bytes).

### Error Handling (Recommendations)
- On malformed frame or header/type mismatch: ignore silently or respond with ERROR depending on UX strategy.
- On payload size mismatch: reject the message.

### Future Extensions (Out of scope v0)
- Additional message types for configuration and data
- Optional MIC/HMAC and encryption
- Larger NONCE and synchronized counters


