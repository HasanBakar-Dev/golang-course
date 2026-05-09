## 1. Language Foundations (Core Syntax & Semantics)

* Go philosophy (simplicity, orthogonality, explicitness)
* Program structure (`package`, `main`, `init`)
* Identifiers, keywords, naming conventions
* Comments and documentation comments
* Unicode and UTF-8 handling
* Lexical scoping
* Blocks and lifetime
* Zero values
* Declaration vs initialization
* Short variable declaration (`:=`)
* Constants (`iota`, typed vs untyped)
* Type inference
* Type identity vs type equivalence
* Compile-time vs runtime behavior

---

## 2. Built-in Types & Values

* Boolean type
* Numeric types

  * Signed/unsigned integers
  * Floating-point
  * Complex numbers
  * Overflow behavior
* Strings (immutability, runes vs bytes)
* Runes (`int32`)
* Bytes (`uint8`)
* Nil value semantics
* `any` (alias for `interface{}`)

---

## 3. Composite Data Types

* Arrays (value semantics)
* Slices

  * Slice header (ptr, len, cap)
  * Slice expressions
  * Re-slicing
  * Copy behavior
  * Append growth strategy
* Maps

  * Hashing
  * Zero value behavior
  * Iteration order randomness
  * Reference semantics
* Structs

  * Field alignment & padding
  * Embedded fields
  * Anonymous structs
  * Tags (JSON, DB, validation)
* Pointers

  * Addressability
  * Escape analysis
* Type aliases vs defined types

---

## 4. Control Flow

* `if` (with initialization)
* `switch`

  * Expression switch
  * Type switch
* `for`

  * Traditional
  * While-style
  * Infinite loop
* `range`

  * Over slices, maps, strings, channels
* `break`, `continue`, `goto`
* `defer`

  * Stack behavior
  * Evaluation timing
* `panic` and `recover`
* Error propagation patterns

---

## 5. Functions & Methods

* Function declarations
* Multiple return values
* Named return values
* Variadic functions
* First-class functions
* Closures
* Method sets
* Value vs pointer receivers
* Method expressions
* Method values
* Function inlining
* Recursion
* Tail-call (non-guaranteed)

---

## 6. Type System & Interfaces (Go’s Intellectual Core)

* Structural typing
* Interfaces
* Empty interface
* Interface satisfaction (implicit)
* Interface embedding
* Nil interface pitfalls
* Type assertions
* Type switches
* Duck typing in Go
* Compile-time interface checks
* Polymorphism without inheritance
* Composition over inheritance
* Liskov substitution (practical Go version)

---

## 7. Error Handling Paradigm

* `error` interface
* Sentinel errors
* Wrapped errors (`fmt.Errorf`, `%w`)
* Error unwrapping
* `errors.Is`, `errors.As`
* Custom error types
* When to panic vs return error
* Error context propagation
* Go 1.20+ error patterns

---

## 8. Packages & Modules

* Package design principles
* Exported vs unexported identifiers
* Package initialization order
* Circular dependency avoidance
* `go.mod`, `go.sum`
* Semantic versioning
* Module replacement
* Vendoring
* Workspace mode
* Internal packages
* Dependency hygiene

---

## 9. Concurrency Model (CSP, Not Threads)

* Goroutines
* Stack growth
* Scheduler (G-M-P model)
* Channels

  * Buffered vs unbuffered
  * Directional channels
  * Closing channels
* Select statement
* Fan-in / fan-out
* Worker pools
* Pipelines
* Backpressure
* Race conditions
* Deadlocks
* Livelocks
* Starvation
* Memory visibility
* Happens-before guarantees
* `sync` package

  * Mutex
  * RWMutex
  * Cond
  * WaitGroup
  * Once
  * Atomic operations
* Context propagation (`context.Context`)
* Cancellation patterns
* Timeouts and deadlines

---

## 10. Memory Management & Runtime

* Garbage collection (tri-color, concurrent)
* Allocation patterns
* Stack vs heap
* Escape analysis
* Pointer scanning
* GC tuning
* Memory profiling
* Runtime scheduler
* Preemption
* Syscalls handling
* Signal handling

---

## 11. Standard Library Mastery

* `fmt`
* `os`, `io`, `io/fs`
* `bufio`
* `net`, `net/http`
* `encoding/*` (json, xml, gob, csv)
* `time`
* `math`, `math/big`
* `regexp`
* `sync`, `sync/atomic`
* `context`
* `database/sql`
* `testing`
* `reflect`
* `unsafe`
* `log`, `log/slog`
* `embed`

---

## 12. Reflection & Metaprogramming

* Reflection basics
* `reflect.Type` vs `reflect.Value`
* Kind vs type
* Settable values
* Tag inspection
* Performance costs
* Alternatives to reflection
* `unsafe` package
* Memory layout manipulation
* When reflection is a design smell

---

## 13. Toolchain & Build System

* `go build`
* `go run`
* `go test`
* `go vet`
* `go fmt`
* `go mod`
* `go work`
* Cross-compilation
* Build tags
* Linker behavior
* Static vs dynamic binaries

---

## 14. Testing & Quality

* Unit testing
* Table-driven tests
* Subtests
* Test fixtures
* Benchmarking
* Fuzz testing
* Race detector
* Code coverage
* Property-based testing philosophy

---

## 15. Performance Engineering

* Profiling (CPU, memory, block)
* pprof
* Allocation minimization
* Cache friendliness
* Escape reduction
* Zero-copy patterns
* Avoiding premature optimization
* Measuring before tuning

---

## 16. Design Idioms (Unwritten but Essential)

* Accept interfaces, return structs
* Small interfaces
* Explicit dependencies
* Avoid global state
* Composition-first design
* Context-first APIs
* Error-first returns
* No exceptions philosophy
* Predictability over cleverness
* Readability as performance

---

## 17. Advanced Language Topics

* Generics (type parameters)

  * Constraints
  * Type inference
  * Comparable
  * Performance implications
* Generic data structures
* Generic APIs design
* Monomorphization behavior
* Backward compatibility issues

---

## 18. Distributed & Systems-Level Go

* Network programming
* HTTP servers & clients
* Middleware patterns
* Graceful shutdown
* Load shedding
* Observability
* Metrics
* Tracing
* Logging strategies
* Microservices vs monoliths (Go tradeoffs)

---

## 19. Security Concepts in Go

* Secure randomness
* TLS configuration
* Input validation
* Serialization risks
* Race-condition security bugs
* Memory safety boundaries
* Supply-chain risks

---

## 20. Ecosystem & Real-World Practice

* Go project layout conventions
* CLI tools
* Long-running services
* Containers & Go
* Cross-platform binaries
* Interop with C (cgo)
* When *not* to use Go

---

### Final intellectual checkpoint

Mastery comes from understanding **what Go refuses to do for you**.


```golang
package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	balance int = 1000
)

func deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	balance += amount
	fmt.Println("Deposited", amount, " Current Balance:", balance)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go deposit(500, &wg)
	fmt.Println("Previous Balance:", balance)
	wg.Add(1)
	go deposit(4000, &wg)
	wg.Wait()
	fmt.Println("Final Balance:", balance)

}
```