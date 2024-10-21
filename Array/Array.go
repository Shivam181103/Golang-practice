package Array

import "fmt"

func Array() {
    numbers := []int{1, 2, 3, 4}

    fmt.Println("Array elements:")
    for i := 0; i < len(numbers); i++ {
        fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
    }
}
