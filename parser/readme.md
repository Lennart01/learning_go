# Expression Parser and Evaluator

This is a simple expression parser and evaluator written in Go. It supports parsing and evaluation of simple arithmetic expressions containing the following operators:

- Addition (+)
- Multiplication (*)

The parser can handle expressions containing the following operands:

- Integer literals (0, 1, 2, ...)
- Parentheses for grouping expressions

## How It Works

The parser works by taking an arithmetic expression as input and generating an abstract syntax tree (AST) from it. The AST is a tree-like representation of the structure of the expression, with each node in the tree representing an operation or operand.

The parser uses a recursive descent parsing technique to generate the AST. It does this by breaking the expression down into its component parts, starting with the highest level of precedence (parentheses), and working its way down to the lowest level of precedence (addition and multiplication).

The parser first looks for expressions in parentheses, and recursively parses the contents of the parentheses to generate a subtree of the AST. If there are no parentheses, the parser looks for multiplication expressions, and recursively parses the left and right operands to generate a subtree of the AST. If there are no multiplication expressions, the parser looks for addition expressions, and again recursively parses the left and right operands to generate a subtree of the AST.

Once the AST has been generated, the evaluator can then traverse the tree and evaluate the expression by recursively evaluating the nodes of the tree from the bottom up.

The evaluator works by recursively evaluating each node of the tree. For integer literals, it simply returns the integer value. For addition and multiplication nodes, it recursively evaluates the left and right operands, and applies the corresponding operation to the results.

## Diferences to the original [C++ implementation](cpp_source)

### Parsing strategy

The Go implementation uses a recursive descent parsing strategy, where each level of precedence is handled by a separate parsing function.

The C++ implementation, on the other hand, uses a bottom-up parsing strategy known as operator precedence parsing, where operators are parsed in order of their precedence. .

## Limitations

- It only supports a limited set of operators and operands
- It does not support variables or functions
- It does not handle syntax errors gracefully
