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

## Similarities to the original [C++ implementation](cpp_source)
- Both implementations use a recursive descent parsing technique to generate an abstract syntax tree (AST) from the input expression.
- Both implementations use a similar approach to evaluate the expression by recursively traversing the AST and applying the corresponding operations to the operands.
- Both implementations support a limited set of operators and operands, including addition, multiplication, integer literals, and parentheses.

## Diferences to the original [C++ implementation](cpp_source)
- The Go implementation uses interfaces to define the AST nodes, while the C++ implementation uses a class hierarchy.
- The Go implementation uses Go's built-in error handling mechanism, while the C++ implementation uses exceptions.
  
## Limitations

- It only supports a limited set of operators and operands wich includes:
  - Addition (+)
  - Multiplication (*)
  - Integer literals (0, 1, 2, ...)
  - Parentheses for grouping expressions
- It does not support variables or functions
- It does not handle syntax errors gracefully

## Bonus (Pratt Parser)

The Pratt Parser is a top-down operator precedence parser that is used to parse expressions. It was invented by Vaughan Pratt in 1973 and is widely used in programming languages and compilers.

### How It Works

The Pratt Parser works by recursively parsing expressions based on operator precedence. It uses a set of parsing rules that are defined for each operator, which specify how to parse expressions containing that operator.

The parser starts by parsing the leftmost token in the expression, which is usually an operand. It then looks ahead to the next token to determine if it is an operator. If it is, the parser checks the parsing rules for that operator to determine how to parse the expression.

The parsing rules for an operator specify the precedence level of the operator, as well as the associativity (left or right) and the behavior of the operator when it is applied to its operands. The parser uses these rules to recursively parse the expression, building up an abstract syntax tree (AST) as it goes.

### Implementation
The Pratt Parser is implemented in the `pratt_parser.go` file. The file defines several types, including `Token`, `Expression`, `Number`, and `BinaryOp`. `Token` represents a token in the input string, `Expression` represents an expression in the input string, `Number` represents a numeric value in the input string, and `BinaryOp` represents a binary operation in the input string.

The file also defines a `Parser` struct that takes a list of tokens and uses a `parseExpression` method to parse the input string and return the resulting expression. The `parseExpression` method is the main entry point for parsing expressions, and takes a `precedence level` as an argument. It first parses the leftmost expression using the `parseAtom` method, and then iteratively parses infix expressions using the `parseExpression ` method until it reaches an operator with a lower precedence level than the current precedence.

The `parseAtom` method parses an atomic expression, which can be a number or a subexpression enclosed in parentheses. The parseExpression method parses an expression with the given precedence level, using the parsing rules for each operator to determine how to parse the expression.



