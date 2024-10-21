package ErrorHandling

import "fmt"

func CausePainc () {
	fmt.Println("About to panic!")
    panic("Something went wrong")
    fmt.Println("This will not print")
}

func PanicRecovery () {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    fmt.Println("About to panic!")
    panic("Something went wrong")
    fmt.Println("This will not print")
}