package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    score := 0
    
    for i := 0; i < 5; i++ {
        a, b := rand.Intn(10), rand.Intn(10)
        fmt.Printf("%d + %d = ", a, b)
        
        var answer int
        fmt.Scanln(&answer)
        
        if answer == a + b {
            score++
            fmt.Println("✅ Верно!")
        } else {
            fmt.Println("❌ Неверно!")
        }
    }
    fmt.Printf("Твой счет: %d/5\n", score)
}
