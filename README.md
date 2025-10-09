# SHA-256 Learning Implementation

This repository contains a simple implementation of SHA-256, created for **learning purposes only**.  
⚠️ It is **not intended for production use**.

---

## Running the Example Solutions

To view the answers for the tasks from the assignment sheet, run:

```bash
go run main.go
```

**Alternatively**, you can check the `hash_test.go` file, which also contains tests related to the exercises.

---

## Running the SHA-256 Function

To compute a SHA-256 hash for a custom message, run:

```bash
go run cmd/main.go hash <your-message>
```

## Arguments

| Argument         | Description                                        |
| ---------------- | -------------------------------------------------- |
| `<your-message>` | **Message to be hashed.** Example: `"hello world"` |

This will print logs to help you **follow and understand the process**.

---

## Running with Weak Hashing (for experimentation)

You can also try a weaker version of the SHA-256 function with additional parameters:

```bash
go run cmd/main.go hash <your-message> <iterations> false
```

## Arguments

| Argument         | Description                                                                                |
| ---------------- | ------------------------------------------------------------------------------------------ |
| `<your-message>` | **Message to be hashed.** Example: `"hello world"`                                         |
| `<iterations>`   | **Number of loops (iterations)** – defines how many times the hashing process is repeated. |
| `false`          | **Uses a weaker version of the hash function** – useful for experimentation and learning.  |

This allows you to experiment with reduced hashing strength for learning purposes.

---

## Repository Structure

```
.
├── cmd/              # Contains the CLI entrypoint (main command logic)
│   └── main.go
├── main.go           # Basic runner for exercises
├── hash_test.go      # Tests for the SHA-256 implementation and exercises
├── core/...          # The Logic of the code
└── README.md         # Project documentation
```
