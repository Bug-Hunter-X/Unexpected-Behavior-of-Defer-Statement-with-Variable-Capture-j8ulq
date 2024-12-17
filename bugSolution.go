```go
func main() {
    var i int
    fmt.Println(i) // Output: 0

    defer func(j int) {
        fmt.Println(j) // Output: 1
    }(i)

    i = 1
    fmt.Println(i) // Output: 1
}
```