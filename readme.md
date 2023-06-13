# Learning Go
This repo contains my personal learning experience with Go as part of my project work at the [University of Applied Sciences Hochschule Karlsruhe](https://www.h-ka.de). Supervised by [Prof. Dr. Martin Sulzmann (Hochschule Karlsruhe)](https://www.h-ka.de/die-hochschule-karlsruhe/organisation-personen/personen-a-z/person/martin-sulzmann)

## C++ Source Code
The C++ source code is by [Prof. Dr. Martin Sulzmann (Hochschule Karlsruhe)](https://www.h-ka.de/die-hochschule-karlsruhe/organisation-personen/personen-a-z/person/martin-sulzmann) and extracted from these [slides](https://sulzmann.github.io/SoftwareProjekt/schein-neu.html#(8)).

## Structure
````
.
├── ast (Abstract Syntax Tree)
│   ├── ast.go
│   ├── ast_test.go
│   ├── cpp_source (contains the c++ source wich was rewritten in go)
│   │   ├── ast.cpp
│   │   └── ast.h
│   └── readme.md
├── go.mod
├── parser (parse a String to an AST)
│   ├── cpp_source (contains the c++ source wich was rewritten in go)
│   │   ├── parser.cpp
│   │   ├── parser.h
│   │   ├── tokenizer.h
│   │   └── utility.h
│   ├── parser.go (recursive descent parser)
│   ├── parser_test.go
│   ├── pratt_parser
│   │   ├── pratt_parser.go (pratt parser)
│   │   └── pratt_parser_test.go
│   └── readme.md
├── readme.md
└── vm (virtual machine, takes an ast as input transforms it into opcodes and executes them)
    ├── cpp_source (contains the c++ source wich was rewritten in go)
    │   ├── utility.h
    │   ├── vm.cpp
    │   └── vm.h
    ├── readme.md
    ├── vm.go
    └── vm_test.go
````
