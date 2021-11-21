# Learning-Golang

Repository used to learn Go.

## Types

Go has several built-in data types.

* Numbers 

For number Go has Integers and Floating point.

Integer data types in Go: `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`, `int64` and also `int` and `uint` whose size is dependent on the type of architecture you are using.
We also have `byte` which is the same as uint8 and `rune` which is the same as `int32`. 
Generally if you are working with integers you should use the `int` type. 

Floating point data types in Go: `float32`, `float64`.
Generally use `float64` when working with floats.

Go also has two special types for complex types: `complex64`, `complex128`.

* Booleans

A boolean is a 1 bit integer used to represent `true` and `false`.
We also have 3 logical operators to use with boolean types:
* && and
* || or
* ! not

## Variables

variables in Go are created using the `var` keyword, then specifying the variable name, the type and finally assigning a value to the variable.

```go
var name string = "Daniel"
```

When we are declaring and assigning a variable right-away we have a shorter syntax, which is a feature of Go called `type-inference`.

```go
name := "Daniel"
```

Here because we are also assigning the variable right-away the Go compiler infers the type automatically.
Generally use the shorter form whenever possible.