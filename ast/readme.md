# Abstract Syntax Tree
## Design Goal
The goal is to replicate the C++ AST in Go.
## Current State
The C++ implementation uses an `Exp class` to represent all expressions.
The following expressions exist:
- `IntExp` for integer literals
- `PlusExp` for addition
- `MultExp` for multiplication
They all inherit from ´Exp´.
## Go Implementation
In Go, we can use interfaces to achieve the same result.