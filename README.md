# go-learn
Jotting down some code + notes as I learn go for the first time!

### May 13th
- Creating a Go project: https://go-zero.dev/en/docs/tasks/create/project/from/vscode
- When declaring and initializing variables inside a function, `f := “apples”` is a short assignnment statement, can be used to replace `var f string = “apples”`
- The format for delaring variables is `var [name] [type]`

    - Eg. `var connected bool`
    - Eg. `var hi, hello, goodbye bool` (list declaration)

- Type can be omitted if initalizer value is present: `var i, j = 1, 2`
- Basic data types are: 

    - `bool`
    - `string`
    - `int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr`
    - `byte` (alias for uint8)
    - `rune` (alias for int32)
    - `float32 float64`
    - `complex64 complex128` (1double complex, 16 bytes wide total, 8 bytes for real (lower), 8 for imag)

- `int uint uintptr` are 32 bits wide on 32-bit systems and 64 bits on 64-bit systems
- No need to use sized or unsigned integer type unless you have specific reason 
- `%T` string formatter gives you the data type of a variable
- Not initialized variables default to `0`, `false`, or `""`
- Use `float32(x) uint(x)` (`T(v)`) operator to cast 
- Type of untyped numeric constants depend on the prevision of the constant

    - `42` -> int
    - `23.322` -> float64
    - `34.32 + 0.63i` -> complex128

- `const Hi = 5` for constants, cannot be declared by `:=`
- `for i := 0; i < 10; i++ { }` for loop syntax
- A while loop is just `for sum < 100 { }`
- While true loop is just `for { }`





