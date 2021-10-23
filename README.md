# namedreturn

**DO NOT use `named return values` in Go**

`namedreturn` finds a named return value, you must remove it.

```go
package a

func _()                {}              // OK
func _() int            { return 1 }    // OK
func _() (int, int)     { return 1, 1 } // OK
func _() (a int)        { return 1 }    // want "a is named return value"
func _() (a int, b int) { return 1, 1 } // want "a is named return value" "b is named return value"
func _() (a, b int)     { return 1, 1 } // want "a is named return value" "b is named return value"
```

## Installation

```
go install github.com/thara/namedreturn/cmd/namedreturn@latest
```

## Usage

```
go vet -vettool=$(which namedreturn) ./...
```

## Author

Tomochika Hara (a.k.a [thara](https://thara.dev))
