# Extodan Language Specification

Extodan is a programming language that blends the best features of Elixir and Go. It is designed to be concise, expressive, and performant, making it suitable for a wide range of applications.

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
    - [Functions](#functions)
    - [Variables](#variables-1)
    - [Keywords](#keywords)
    - [Examples](#examples)

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

Variables in Extodan are declared using the 'var' keyword. The type of a variable is automatically inferred from the assigned value.

### Data Types

Extodan supports various data types, including:

- Integer: Represents whole numbers.
- Float: Represents floating-point numbers.
- String: Represents a sequence of characters.
- Boolean: Represents true or false.
- List: Represents a collection of elements of the same type.

### Control Flow

Extodan provides conditional statements (`if-else`) and loops (`while`) for controlling the flow of the program.

### Functions

Functions are defined using the `func` keyword, followed by the function name, parameters, and the `do` keyword to start the function body.

### Variables

Variables in Extodan are declared using the `var` keyword. The type of a variable is automatically inferred from the assigned value.

### Keywords

The Extodan language has several reserved keywords with special meanings, including `func`, `do`, `end`, and `var`.

### Examples

Here's a simple Extodan program that calculates the factorial of a number:

```extodan
func factorial(n) do
    if n <= 1 do
        return 1
    else
        return n * factorial(n - 1)
    end
end

var num = 5
var result = factorial(num)
print("Factorial of ", num, " is ", result)
