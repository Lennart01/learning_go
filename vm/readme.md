# Go Virtual Machine
This is a simple virtual machine written in Go that can execute a small set of instructions. The virtual machine uses a stack to store values and supports the following instructions:

- `PUSH` `<value>`: Pushes a value onto the stack
- `PLUS`: Pops the top two values from the stack, adds them together, and pushes the result onto the stack
- `MULTIPLY`: Pops the top two values from the stack, multiplies them together, and pushes the result onto the stack
The virtual machine is implemented in the `VM` struct in `main.go`. The Run method of the `VM` struct executes the instructions and returns the result.

## Usage
To use the virtual machine, create a new `VM` struct with the instructions you want to execute, and call the Run method. For example:
```go
vm := NewVM([]Code{
    NewPushCode(1),
    NewPushCode(2),
    NewPlusCode(),
    NewPushCode(3),
    NewMultiplyCode(),
})
result := vm.Run()
```
The result variable will contain the result of the calculation.

## Comparison to the original [C++ implementation](cpp_source)

### Classes and Methods vs. Structs and Functions
- **C++**: Code is organized around *classes* (`Code`, `VM`, `Optional`) and *methods* (`run`, `newPush`, `newPlus`, `newMult`). C++ classes allow encapsulation of data and methods, and support features like inheritance and polymorphism.
- **Go**: Doesn't have classes. It uses *structs* (`Optional`, `Code`, `VM`) and *functions*. You can define methods on structs in Go, which is how `Run`, `NewPushCode`, `NewPlusCode`, `NewMultiplyCode`, `IsJust`, `IsNothing`, `Value` are defined. However, Go does not support inheritance or polymorphism in the same way as C++.

### Enumerations
- **C++**: `OpCode_t` is an enumeration used to represent opcodes.
- **Go**: Doesn't have true enums, but it can mimic them using `const` and `iota`. `OpCode` in the Go code serves the same purpose as `OpCode_t` in C++.

### Optional/Maybe Types
- **C++**: Uses templates (`Optional<T>`) for Optional type, which allows for type-safe operations on any type `T`.
- **Go**: `Optional` struct uses an `interface{}` to hold the value, which can be any type. This is due to Go's lack of generics.

### Error Handling
- **Both**: Return a special 'nothing' or 'empty' value when an operation cannot be performed, as part of the Maybe/Optional paradigm.

### Type Conversion
- **Go**: Includes type assertions (e.g., `.(int)`) to convert values from the `interface{}` type to the needed specific type due to static typing and no automatic type conversion.
- **C++**: Also requires explicit type conversion in some cases, but it's not needed in the cpp source code.
