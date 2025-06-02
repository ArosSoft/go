package main
import (
    "log"
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "fmt"
    "github.com/Knetic/govaluate"
    "strings"
    "unicode"
)

// Структура для хранения состояния игры
type GameState struct {
    Word string
    Guessed []bool
    Attempts int
    UsedLetters map[rune]bool
    IsPlaying bool
}

// Глобальная карта для хранения состояния игры для каждого пользователя
var gameStates = make(map[int64]*GameState)

func main() {
	// бот Arossoft_bot
    bot, err := tgbotapi.NewBotAPI("8056397360:AAF3hrv4nw45NtT-TQ4q4RE-w-Dp39mZ7f8")
    if err != nil {
        log.Panic(err)
    }

    u := tgbotapi.NewUpdate(0)
    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Command() {
		case "start":
			msg.Text = "Привет! Я бот на Go!\nКоманды:\n/game - игра в виселицу\n/calc - калькулятор"
		case "game":
			// Инициализация новой игры
			gameState := &GameState{
				Word: "программирование",
				Guessed: make([]bool, len([]rune("программирование"))),
				Attempts: 6,
				UsedLetters: make(map[rune]bool),
				IsPlaying: true,
			}
			gameStates[update.Message.Chat.ID] = gameState
			msg.Text = getGameStatus(gameState)
		case "calc":
			// Разбираем выражение типа "/calc 2+2"
			// Используем CommandArguments() для получения текста после команды /calc
			expressionText := update.Message.CommandArguments()
			
			// Если аргументов нет, просим пользователя ввести выражение
			if expressionText == "" {
				expressionText = "Пожалуйста, введите выражение после /calc"
				msg.Text = expressionText
			} else {
				msg.Text = calculate(expressionText)
			}
		default:
			// Обработка обычных сообщений
			if gameState, exists := gameStates[update.Message.Chat.ID]; exists && gameState.IsPlaying {
				msg.Text = processGameInput(update.Message.Text, gameState)
			}
		}
		
        bot.Send(msg)
		
    }
}

func getGameStatus(gameState *GameState) string {
    display := ""
    for i, c := range []rune(gameState.Word) {
        if gameState.Guessed[i] {
            display += string(c)
        } else {
            display += "_"
        }
    }
    
    usedLetters := getUsedLettersString(gameState.UsedLetters)
    return fmt.Sprintf("Слово: %s\nИспользованные буквы: %s\nОсталось попыток: %d", 
        display, usedLetters, gameState.Attempts)
}

func processGameInput(input string, gameState *GameState) string {
    input = strings.TrimSpace(input)
    
    // Проверка корректности ввода
    if len([]rune(input)) != 1 {
        return "Пожалуйста, введите только одну букву!"
    }
    
    letter := []rune(input)[0]
    
    // Проверка на повторный ввод
    if gameState.UsedLetters[letter] {
        return "Эта буква уже была использована!"
    }
    
    gameState.UsedLetters[letter] = true
    
    // Проверка угадывания
    found := false
    for i, c := range []rune(gameState.Word) {
        if unicode.ToLower(c) == unicode.ToLower(letter) {
            gameState.Guessed[i] = true
            found = true
        }
    }
    
    if !found {
        gameState.Attempts--
        if gameState.Attempts <= 0 {
            gameState.IsPlaying = false
            return fmt.Sprintf("Игра окончена! Загаданное слово было: %s", gameState.Word)
        }
        return fmt.Sprintf("Такой буквы нет! Осталось попыток: %d", gameState.Attempts)
    }
    
    // Проверка на победу
    if allGuessed(gameState.Guessed) {
        gameState.IsPlaying = false
        return fmt.Sprintf("Поздравляем! Вы выиграли!\nЗагаданное слово: %s", gameState.Word)
    }
    
    return getGameStatus(gameState)
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

func calculate(expression string) string {
	// Создаем объект выражения из строки
	expr, err := govaluate.NewEvaluableExpression(expression)
	// Вычисляем выражение.
	// Второй аргумент предназначен для параметров (переменных) в выражении,
	// которых у нас нет в простом калькуляторе, поэтому передаем nil.
	result, err := expr.Evaluate(nil)
	if err != nil {
		// Обрабатываем ошибки вычисления (например, деление на ноль)
		return fmt.Sprintf("Ошибка вычисления выражения: %v", err)
	}

	// Преобразуем результат в строку
	return fmt.Sprintf("%v", result)
}
