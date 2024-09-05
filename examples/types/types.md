# Types in Go

Go is a statically-typed language, which means that the type of every variable is known at compile time. This section will provide an overview of the most important types in Go and how they are used.

## Basic Types
Go has a variety of basic types, including:

- **Boolean**: `bool`
  - Values: `true` or `false`
- **Numeric Types**:
  - **Integers**: Signed integers (`int`, `int8`, `int16`, `int32`, `int64`) and unsigned integers (`uint`, `uint8`, `uint16`, `uint32`, `uint64`).
  - **Floating-point**: `float32`, `float64`
  - **Complex numbers**: `complex64`, `complex128`
  - **Byte**: `byte` (alias for `uint8`)
  - **Rune**: `rune` (alias for `int32`, represents Unicode code points)
- **String**: Sequence of characters enclosed in double quotes (`""`). Strings are immutable.

## Composite Types
Go provides several composite types that allow grouping other types together:

### Arrays
- **Definition**: Fixed-length sequences of elements of the same type.
- **Syntax**: 
  ```go
  var arr [5]int
  arr := [3]int{1, 2, 3}
  ```
- Arrays have a fixed size and cannot be resized after creation.

### Slices
- **Definition**: Dynamic, resizable view over an array. More flexible than arrays.
- **Syntax**: 
  ```go
  var s []int
  s := []int{1, 2, 3}
  s = make([]int, 3)
  ```
- Slices are a reference type, meaning they point to an underlying array. They have both **length** and **capacity**.

### Maps
- **Definition**: Unordered collection of key-value pairs.
- **Syntax**: 
  ```go
  var m map[string]int
  m := make(map[string]int)
  ```
- Maps must be initialized using `make` or a map literal before use.

### Structs
- **Definition**: Collections of fields grouped together to form custom types.
- **Syntax**: 
  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```
- Structs are value types. They can hold any type of data, including other structs and arrays.

## Type Declarations
You can define your own types in Go to improve readability and enforce stricter typing:

### Type Alias
- **Definition**: Creates an alias for an existing type.
- **Syntax**:
  ```go
  type MyInt = int
  ```

### New Type
- **Definition**: Creates a new distinct type based on an existing type.
- **Syntax**:
  ```go
  type Age int
  ```

## Interfaces
- **Definition**: Defines a set of method signatures that a type must implement.
- **Syntax**:
  ```go
  type Stringer interface {
      String() string
  }
  ```
- Any type that implements the methods of an interface satisfies that interface.

## Type Parameters (Generics)
- **Definition**: Allows defining functions, methods, and types that work with any data type.
- **Syntax**:
  ```go
  func PrintSlice[T any](s []T) {
      for _, v := range s {
          fmt.Println(v)
      }
  }
  ```
- Generics enable type-safe code reuse for multiple types.

## Pointers
- **Definition**: Variables that store the memory address of another variable.
- **Syntax**: 
  ```go
  var p *int
  ```
- Use the `&` operator to get the address of a variable and `*` to dereference it.

## Type Assertions & Type Switches
- **Type Assertion**: Used to extract the concrete type from an interface.
  ```go
  var i interface{} = 42
  val := i.(int)
  ```
- **Type Switch**: Switch over the dynamic type stored in an interface:
  ```go
  switch v := i.(type) {
  case int:
      // handle int
  case string:
      // handle string
  }
  ```

## Type Inference
- **Definition**: Go can infer the type of a variable based on its initial value.
- **Syntax**:
  ```go
  x := 42 // inferred as int
  ```

## Type Safety & Static Typing
- Go ensures **type safety** at compile time, meaning invalid operations between mismatched types are caught early.
- Every expression has a specific type, and Go strictly enforces this at compile time.
