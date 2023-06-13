# Abstract Syntax Tree
This is a simple example of an abstract syntax tree (AST) in C++ and Go.
## C++ Implementation
The C++ implementation uses an `Exp class` to represent all expressions.
The following expressions exist:
- `IntExp` for integer literals
- `PlusExp` for addition
- `MultExp` for multiplication
They all inherit from ´Exp´.
## Go Implementation
In Go, we can use interfaces to achieve the same result.
The interface is called `Exp` and is defined like this:
```go
type Exp interface {
	Eval() int
	Pretty() string
}
```
All other expressions implement this interface implicitly because they implement the two methods `Eval` and `Pretty`.
This is the main difference between the C++ and the Go implementation as Go handles inheritance differently.