package main
import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "unicode"
)

func main() {
    word := "вылысыпыдыстычкы"
    guessed := make([]bool, len([]rune(word)))
    attempts := 6
    usedLetters := make(map[rune]bool)
    reader := bufio.NewReader(os.Stdin)
    
    for {
        // Показываем текущий прогресс
        display := ""
        for i, c := range []rune(word) {
            if guessed[i] {
                display += string(c)
            } else {
                display += "_"
            }
        }
        fmt.Println("\nСлово:", display)
        fmt.Println("Использованные буквы:", getUsedLettersString(usedLetters))
        
        // Проверка на победу
        if allGuessed(guessed) {
            fmt.Println("\nПоздравляем! Вы выиграли!")
            return
        }
        
        // Проверка на проигрыш
        if attempts <= 0 {
            fmt.Printf("\nИгра окончена! Загаданное слово было: %s\n", word)
            return
        }
        
        // Ввод буквы
        fmt.Print("Введите букву: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        
        // Проверка корректности ввода
        if len([]rune(input)) != 1 {
            fmt.Println("Пожалуйста, введите только одну букву!")
            continue
        }
        
        letter := []rune(input)[0]
        
        // Проверка на повторный ввод
        if usedLetters[letter] {
            fmt.Println("Эта буква уже была использована!")
            continue
        }
        
        usedLetters[letter] = true
        
        // Проверка угадывания
        found := false
        for i, c := range []rune(word) {
            if unicode.ToLower(c) == unicode.ToLower(letter) {
                guessed[i] = true
                found = true
            }
        }
        
        if !found {
            attempts--
            fmt.Printf("Такой буквы нет! Осталось попыток: %d\n", attempts)
        } else {
            fmt.Println("Верно! Буква есть в слове.")
        }
    }
}

func allGuessed(guessed []bool) bool {
    for _, g := range guessed {
        if !g {
            return false
        }
    }
    return true
}

func getUsedLettersString(usedLetters map[rune]bool) string {
    letters := make([]string, 0, len(usedLetters))
    for letter := range usedLetters {
        letters = append(letters, string(letter))
    }
    return strings.Join(letters, ", ")
}
