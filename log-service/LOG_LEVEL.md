# Log Levels and Use Cases

## 1. TRACE 🪶 (Most Detailed)
**Meaning:** Logs every step of execution, like step-by-step code tracing.

**Use Case:** Debugging complex issues, multi-service flows, or internal library problems.

- **Production:** ❌ Almost always off (too verbose, high I/O).
- **Development:** ✅ Enable for hard-to-debug bugs (race conditions, data flow issues).

**Example:**
```plaintext
TRACE 2025-08-09T12:10:05Z auth-service validating user input: {username: "john", password: "***"}
TRACE 2025-08-09T12:10:05Z auth-service calling DB query: SELECT * FROM users WHERE username = ?
```

## 2. DEBUG 🐛
**Meaning:** Developer-focused logs showing parameters, responses, and execution paths.

**Use Case:** 
- Check execution flow.
- Inspect function input/output.

- **Production:** ❌ Off, or enabled via feature flag for on-demand debugging.
- **Development:** ✅ Always on.

**Example:**
```plaintext
DEBUG 2025-08-09T12:10:05Z auth-service login payload received: {username: "john"}
DEBUG 2025-08-09T12:10:05Z auth-service database returned 1 user row
```

## 3. INFO ℹ️
**Meaning:** Logs important normal events for monitoring, not errors.

**Use Case:** 
- Start/end of important processes.
- Successful API calls.
- Service startup/shutdown.

- **Production:** ✅ Always on.
- **Development:** ✅ Always on.

**Example:**
```plaintext
INFO 2025-08-09T12:10:05Z auth-service user john logged in successfully
INFO 2025-08-09T12:00:00Z payment-service processed payment id=12345 amount=500.00
```

## 4. WARN ⚠️
**Meaning:** Logs abnormal events, but system continues to work.

**Use Case:** 
- Using default values due to incomplete config.
- API retry succeeded.
- Invalid input but system fallback works.

- **Production:** ✅ On for monitoring.
- **Development:** ✅ On.

**Example:**
```plaintext
WARN 2025-08-09T12:10:05Z auth-service missing optional header "X-Tracking-ID", using default
WARN 2025-08-09T12:10:05Z payment-service API timeout, retrying...
```

## 5. ERROR ❌
**Meaning:** Logs errors that cause process failure.

**Use Case:** 
- Database query failed.
- Third-party API error.
- Business logic failure.

- **Production:** ✅ On.
- **Development:** ✅ On.

**Example:**
```plaintext
ERROR 2025-08-09T12:10:05Z auth-service failed to query DB: connection refused
ERROR 2025-08-09T12:10:05Z payment-service payment id=12345 failed: insufficient funds
```

## 6. FATAL 💀
**Meaning:** Logs critical errors that stop the service.

**Use Case:** 
- Missing critical config (DB credentials).
- Port already in use.
- Service startup failure.

- **Production:** ✅ On (trigger immediate alert).
- **Development:** ✅ On.

**Example:**
```plaintext
FATAL 2025-08-09T12:10:05Z auth-service failed to start: missing DB connection string
```

## Summary Table: Usage in Each Environment (usage summary each environment)

| Level   | Production | Development                |
|---------|------------|---------------------------|
| TRACE   | ❌         | ✅ Only for hard debugging |
| DEBUG   | ❌ (except temporary) | ✅           |
| INFO    | ✅         | ✅                        |
| WARN    | ✅         | ✅                        |
| ERROR   | ✅         | ✅                        |
| FATAL   | ✅         | ✅                        |