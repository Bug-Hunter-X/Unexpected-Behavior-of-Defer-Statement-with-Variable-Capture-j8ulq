```go
func main() {
    var i int
    fmt.Println(i) // Output: 0

    defer func() {
        fmt.Println(i) // Output: 0
    }()

    i = 1
    fmt.Println(i) // Output: 1
}
```