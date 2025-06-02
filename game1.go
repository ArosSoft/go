package main
import (
    "fmt"
    "strings"
)

func main() {
    word := "mouse"
    guessed := make([]bool, len(word))
    attempts := 6
    usedLetters := make(map[string]bool)
    
    for {
        // Показываем текущий прогресс
        display := ""
        for i, c := range word {
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
        var letter string
        fmt.Print("Введите букву: ")
        fmt.Scanln(&letter)
        letter = strings.ToLower(letter)
        
        // Проверка корректности ввода
        if len(letter) != 1 {
            fmt.Println("Пожалуйста, введите только одну букву!")
            continue
        }
        
        // Проверка на повторный ввод
        if usedLetters[letter] {
            fmt.Println("Эта буква уже была использована!")
            continue
        }
        
        usedLetters[letter] = true
        
        // Проверка угадывания
        found := false
        for i, c := range word {
            if string(c) == letter {
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

func getUsedLettersString(usedLetters map[string]bool) string {
    letters := make([]string, 0, len(usedLetters))
    for letter := range usedLetters {
        letters = append(letters, letter)
    }
    return strings.Join(letters, ", ")
}
