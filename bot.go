package main
import (
    "log"
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "fmt"
    "github.com/Knetic/govaluate"
)

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
			msg.Text = "Привет! Я бот на Go!\nКоманды:\n/game - игра\n/calc - калькулятор"
		case "game":
			msg.Text = "Угадай число от 1 до 10!"
			// Логика игры
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
		}
		
        bot.Send(msg)
		
    }
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
