package loops

import "fmt"

func Switch() string {
	i := 2
    fmt.Println("Switch for i = ", i, " goes to: ")
    var result string
    switch i {
    case 1:
        fmt.Println("one")
        result = "one"
    case 2:
        fmt.Println("two")
        result = "two"
    case 3:
        fmt.Println("three")
        if result == "" {
            result = "three"
        }
    case 4:
        fmt.Println("four")
        if result == "" {
            result = "four"
        }
    case 5, 6:
        fmt.Println("five or six")
        if result == "" {
            result = "five or six"
        }
    default:
        fmt.Println("default")
        if result == "" {
            result = "default"
        }
    }
    return result
}