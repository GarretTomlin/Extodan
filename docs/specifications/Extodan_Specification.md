# Extodan Language Specification

Extodan is a programming language that blends the best features of Elixir and Go. It is designed to be concise, expressive, and performant, making it suitable for a wide range of applications

## Table of Contents

- [Extodan Language Specification](#extodan-language-specification)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Syntax](#syntax)
    - [Program Structure](#program-structure)
    - [Comments](#comments)
    - [Identifiers](#identifiers)
    - [Variables](#variables)
    - [Data Types](#data-types)
    - [Control Flow](#control-flow)
    - [Concurrency](#concurrency)
    - [Functions](#functions)
    - [Modules](#modules)
    - [Creating a Module](#creating-a-module)
    - [Example: Creating a MathOperations Module](#example-creating-a-mathoperations-module)
    - [Error Handling](#error-handling)
    - [Pipe Operator](#pipe-operator)
    - [Keywords](#keywords)
    - [Conclusion](#conclusion)

## Introduction

Extodan is a programming language designed as a mixture of Elixir and Go (Golang), combining the expressive power of Elixir and the simplicity of Go. This specification document outlines the syntax, grammar rules, keywords, data types, control flow constructs, and other language features that define the Extodan language.

## Syntax

### Program Structure

An Extodan program consists of one or more modules. Each module represents a separate unit of code. A module may contain function definitions, variable declarations, and other statements.

### Comments

Comments in Extodan are denoted by the '#' symbol. Anything after the '#' symbol on the same line is considered a comment and is ignored by the compiler.

### Identifiers

Identifiers in Extodan are sequences of letters, digits, and underscores. They must start with a letter or underscore.

### Variables

Variables in Extodan are declared using the (`var`) keyword. The type of a variable is automatically inferred from the assigned value.

```go
var age = 30
var name = "Garret Tomlin"
var isTrue = true
```

### Data Types

Extodan supports various data types, including:

- Integer: Represents whole numbers.
- Float: Represents floating-point numbers.
- String: Represents a sequence of characters.
- Boolean: Represents true or false.
- List: Represents a collection of elements of the same type.

### Control Flow

Extodan provides conditional statements (`if-else`) and loops (`while`) for controlling the flow of the program.

```elixir
func checkNumber(number) do
    if number > 0 do
        return "Positive"
    else if number < 0 do
        return "Negative"
    else
        return "Zero"
    end
endfunc
```


 The (`when`) keyword is used for pattern matching with guards

```elixir
func checkType(x) do
    case x do
        v when is_integer(v) -> return "Integer"
        v when is_string(v) -> return "String"
        v when is_list(v) -> return "List"
        _ -> return "Unknown Type"
    end
endfunc
```
### Concurrency
Extodan supports lightweight concurrency with goroutines, similar to Golang. Goroutines are created using the go keyword.

```elixir
func printNumbers() do
    for i in 1..10 do
        go func() do
            print(i)
        end
    end
endfunc
```

### Functions

Functions are defined using the `func` keyword, followed by the function name, parameters, and the `do` keyword to start the function body.
```elixir

  func add(a, b) do
        return a + b
    endfunc
```

### Modules

In Extodan, a module is a self-contained unit that encapsulates related functions, data, and logic. It allows you to organize your code into logical sections, making it easier to manage and maintain. Modules provide a way to group related functionality together, promoting modularity and code reusability.

### Creating a Module

To create a module in Extodan, you use the `module` keyword followed by the module name and the `do` keyword to start the module body. Inside the module body, you can define functions, variables, and other constructs that are relevant to the module's purpose.

Here's the syntax for creating a module:

```elixir
module ModuleName do
    # Functions, variables, and other code related to the module
endmodule
```


### Example: Creating a MathOperations Module
Let's create a MathOperations module that encapsulates functions for basic arithmetic operations:

```elixir
module MathOperations do

    # Function to add two numbers
    func add(a, b) do
        return a + b
    endfunc
    
    # Function to subtract two numbers
    func subtract(a, b) do
        return a - b
    endfunc
    
    # Function to multiply two numbers
    func multiply(a, b) do
        return a * b
    endfunc
    
    # Function to divide two numbers
    func divide(a, b) do
        return a / b
    endfunc
    
endmodule
```

Using a Module
Once you've defined a module, you can use its functions and other elements by referencing the module's name followed by the function or variable name. This helps maintain a clear separation of concerns and avoids naming conflicts.

Here's an example of how to use the (`MathOperations`) module:

```elixir
func main() do
    var result = MathOperations.add(5, 3)
    print("Addition result:", result)
    
    result = MathOperations.subtract(10, 4)
    print("Subtraction result:", result)
    
    result = MathOperations.multiply(6, 7)
    print("Multiplication result:", result)
    
    result = MathOperations.divide(15, 3)
    print("Division result:", result)
endfunc
```


### Error Handling
Extodan supports error handling using the try and catch keywords, similar to Golang's try and catch.

```elixir
func divide(a, b) do
    try do
        if b == 0 do
            raise "Cannot divide by zero"
        end
        return a / b
    catch error do
        return error
    end
endfunc
```

### Pipe Operator
Extodan incorporates the pipe operator |> from Elixir, which allows for a more readable and expressive code flow.

```elixir
func processNumber(number) do
    number
    |> add(5)
    |> subtract(3)
    |> checkNumber
end
```



### Keywords

The Extodan language has several reserved keywords with special meanings, including `func`, `do`, `end`, and `var`.


### Conclusion 
This specification aims to provide a comprehensive overview of the Extodan programming language, enabling developers to understand and use the language effectively for various applications.
