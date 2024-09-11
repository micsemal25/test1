package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	EMOJI_RACKET       = "\U0001F680" // 🚀
	EMOJI_COIN         = "\U0001F4B8" // 💸
	EMOGI_SHOPING_CART = "\U0001F6D2" // 🛒
	EMOJI_PACKAGE      = "\U0001F4E6" // 📦
	EMOJI_USER         = "\U0001F464"
	EMOJI_SUNGLASSES   = "\U0001F60E"
	EMOJI_SOS          = "\U0001F198"
	EMOJI_LIGHTNING    = "\U000026A1"
	EMOJI_STAR         = "\u2B50"
	EMOJI_BACK         = "\u2B05"
	EMOJI_CHECK        = "\u2705"
	EMOJI_HI           = "\U0001F44B"
	EMOJI_MONEYLOVE    = "\U0001F911"
	EMOJI_KUBOK        = "\U0001F3C6"
	EMOJI_VOSKL        = "\u2757"
	EMOJI_100          = "\U0001F4AF"
	EMOJI_STARLIGHT    = "\U0001F31F"
	EMOJI_TECH         = "\U0001F4AC"
)

var gBot *tgbotapi.BotAPI
var gToken string
var gChatId int64
var lastMessageID int
var adminChatId int64 = 1283075660

func isStartMessage(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

func printSystemMessageWithDelay(delayInSec uint8, message string) {
	msg := tgbotapi.NewMessage(gChatId, message)
	msg.ParseMode = "HTML"
	gBot.Send(msg)
	time.Sleep(time.Second * time.Duration(delayInSec))
}

func printIntro(update *tgbotapi.Update) {
	printSystemMessageWithDelay(2, "Вас приветствует <b>MickyBot</b>!"+EMOJI_HI)
	printSystemMessageWithDelay(1, "В нашем магазине вы можете заказать разработку персонального бота для мессенджеров: \n"+EMOJI_CHECK+"<b>Telegram</b> \n"+EMOJI_CHECK+"<b>Discord</b> \n"+EMOJI_CHECK+"<b>VK</b>")
	printSystemMessageWithDelay(4, EMOJI_STARLIGHT+"Делаем ботов строго под ваши задачи и в срок. \n"+EMOJI_STARLIGHT+"Разрабатываем ботов полностью с нуля на языке программирования GO.\n"+EMOJI_STARLIGHT+"Наши боты имеют высокую производительность и способны выделить вас среди конкурентов. ")
	printSystemMessageWithDelay(4, EMOJI_VOSKL+"<b>ВНИМАНИЕ</b>"+EMOJI_VOSKL+"\nВ честь открытия магазина дарим каждому клиенту <b>скидку 50%</b>"+EMOJI_MONEYLOVE+" на любой товар.")
}

func init() {
	_ = os.Setenv("telegram-bot-1", "7350279682:AAHyNZtabou2nQy9SHuKIqwb-P488BonX50")
	gToken = os.Getenv("telegram-bot-1")
	var err error
	if gBot, err = tgbotapi.NewBotAPI(gToken); err != nil {
		log.Panic(err)
	}
	gBot.Debug = true
}

func main() {
	log.Printf("Authorized on account %s", gBot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	for update := range gBot.GetUpdatesChan(updateConfig) {
		if isStartMessage(&update) { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			gChatId = update.Message.Chat.ID
			printIntro(&update)
			// главное меню
			button1 := tgbotapi.NewInlineKeyboardButtonData("Заказать бота"+EMOJI_COIN, "callback_data_1")
			button2 := tgbotapi.NewInlineKeyboardButtonData("Профиль"+EMOJI_USER, "callback_data_2")
			button3 := tgbotapi.NewInlineKeyboardButtonURL("FAQ"+EMOJI_SOS, "https://t.me/+bke3X2XVmlthYjky")
			button4 := tgbotapi.NewInlineKeyboardButtonURL("Отзывы"+EMOJI_KUBOK, "https://t.me/+LVXFHgP7yv83MTk6")
			if lastMessageID != 0 {
				deleteMsg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, lastMessageID)
				_, err := gBot.Request(deleteMsg)
				if err != nil {
					log.Printf("Failed to delete message: %v", err)
				}
			}
			// Создаем клавиатуру с кнопками
			keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button1, button2), tgbotapi.NewInlineKeyboardRow(button4, button3))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите опцию:")
			msg.ReplyMarkup = keyboard
			sentMessage, err := gBot.Send(msg)
			if err != nil {
				log.Printf("Failed to send message: %v", err)
			}

			// Сохраняем ID последнего отправленного сообщения
			lastMessageID = sentMessage.MessageID

		} else if update.CallbackQuery != nil {
			// Обрабатываем нажатие на кнопку
			callback := update.CallbackQuery
			switch callback.Data {
			// меню выбора товаров
			case "callback_data_1":
				button1 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_LIGHTNING+"Бот для Telegram", "callback_data_1_1")
				button2 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_LIGHTNING+"Бот для Discord", "callback_data_1_2")
				button3 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_LIGHTNING+"Бот для VK", "callback_data_1_3")
				button4 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_BACK+"Назад", "callback_data_1_5")
				keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button1), tgbotapi.NewInlineKeyboardRow(button2), tgbotapi.NewInlineKeyboardRow(button3), tgbotapi.NewInlineKeyboardRow(button4))
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Что желаете приобрести?")
				msg.ReplyMarkup = keyboard
				if lastMessageID != 0 {
					deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, lastMessageID)
					_, err := gBot.Request(deleteMsg)
					if err != nil {
						log.Printf("Failed to delete message: %v", err)
					}
				}
				// Отправляем обновленное сообщение
				sentMessage, err := gBot.Send(msg)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}

				// Сохраняем ID последнего отправленного сообщения
				lastMessageID = sentMessage.MessageID
			// раздел с телеграм подписками
			case "callback_data_1_1":
				button1 := tgbotapi.NewInlineKeyboardButtonData("Подать заявку"+EMOJI_CHECK, "callback_data_zayavka")
				button2 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_BACK+"Назад", "callback_data_back_buy")
				keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button1), tgbotapi.NewInlineKeyboardRow(button2))
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Подайте заявку:")
				msg.ReplyMarkup = keyboard
				if lastMessageID != 0 {
					deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, lastMessageID)
					_, err := gBot.Request(deleteMsg)
					if err != nil {
						log.Printf("Failed to delete message: %v", err)
					}
				}
				// Отправляем обновленное сообщение
				sentMessage, err := gBot.Send(msg)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}

				// Сохраняем ID последнего отправленного сообщения
				lastMessageID = sentMessage.MessageID
				// Отправляем обновленное сообщение
			// оформление заявки
			case "callback_data_1_2":
				button1 := tgbotapi.NewInlineKeyboardButtonData("Сделать заявку"+EMOJI_CHECK, "callback_data_zayavka")
				button2 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_BACK+"Назад", "callback_data_back_buy")
				keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button1), tgbotapi.NewInlineKeyboardRow(button2))
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Подайте заявку:")
				msg.ReplyMarkup = keyboard
				if lastMessageID != 0 {
					deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, lastMessageID)
					_, err := gBot.Request(deleteMsg)
					if err != nil {
						log.Printf("Failed to delete message: %v", err)
					}
				}
				// Отправляем обновленное сообщение
				sentMessage, err := gBot.Send(msg)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}

				// Сохраняем ID последнего отправленного сообщения
				lastMessageID = sentMessage.MessageID
			// Отправляем обновленное сообщение
			case "callback_data_1_3":
				button1 := tgbotapi.NewInlineKeyboardButtonData("Подать заявку"+EMOJI_CHECK, "callback_data_zayavka")
				button2 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_BACK+"Назад", "callback_data_back_buy")
				keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button1), tgbotapi.NewInlineKeyboardRow(button2))
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Сделайте заявку:")
				msg.ReplyMarkup = keyboard
				if lastMessageID != 0 {
					deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, lastMessageID)
					_, err := gBot.Request(deleteMsg)
					if err != nil {
						log.Printf("Failed to delete message: %v", err)
					}
				}
				// Отправляем обновленное сообщение
				sentMessage, err := gBot.Send(msg)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}

				// Сохраняем ID последнего отправленного сообщения
				lastMessageID = sentMessage.MessageID
				// Отправляем обновленное сообщение
			case "callback_data_zayavka":
				user := update.CallbackQuery.From

				//fmt.Println(user)
				// Формируем ссылку на профиль пользователя
				var profileLink string
				if user.UserName != "" {
					profileLink = fmt.Sprintf("https://t.me/%s", user.UserName)
				} else {
					profileLink = fmt.Sprintf("tg://user?id=%d", user.ID)
				}
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Отлично, ваша заявка одобрена!"+EMOJI_CHECK)
				msg2 := tgbotapi.NewMessage(callback.Message.Chat.ID, "В скором времени с вами свяжется менеджер Михаил"+EMOJI_TECH)

				adminMsg := fmt.Sprintf("📝 <b>Новая заявка!</b>\n👤 Имя: %s\n🔗 Профиль: %s", user.FirstName, profileLink)
				adminMessage := tgbotapi.NewMessage(adminChatId, adminMsg)

				// Отправляем сообщение администратору
				_, err := gBot.Send(adminMessage)
				if err != nil {
					log.Printf("Ошибка отправки сообщения администратору: %v", err)
				}
				if lastMessageID != 0 {
					deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, lastMessageID)
					_, err := gBot.Request(deleteMsg)
					if err != nil {
						log.Printf("Failed to delete message: %v", err)
					}
				}
				sentMessage, err := gBot.Send(msg)
				sentMessage, err = gBot.Send(msg2)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}
				lastMessageID = sentMessage.MessageID
			// меню выбора товаров
			case "callback_data_back_buy":
				button1 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_LIGHTNING+"Бот для Telegram", "callback_data_1_1")
				button2 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_LIGHTNING+"Бот для Discord", "callback_data_1_2")
				button3 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_LIGHTNING+"Бот для VK", "callback_data_1_3")
				button4 := tgbotapi.NewInlineKeyboardButtonData(EMOJI_BACK+"Назад", "callback_data_1_5")
				keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button1), tgbotapi.NewInlineKeyboardRow(button2), tgbotapi.NewInlineKeyboardRow(button3), tgbotapi.NewInlineKeyboardRow(button4))
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Что желаете приобрести?")
				msg.ReplyMarkup = keyboard
				if lastMessageID != 0 {
					deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, lastMessageID)
					_, err := gBot.Request(deleteMsg)
					if err != nil {
						log.Printf("Failed to delete message: %v", err)
					}
				}
				// Отправляем обновленное сообщение
				sentMessage, err := gBot.Send(msg)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}

				// Сохраняем ID последнего отправленного сообщения
				lastMessageID = sentMessage.MessageID
			// главное меню
			case "callback_data_1_5":
				button1 := tgbotapi.NewInlineKeyboardButtonData("Заказать бота"+EMOJI_COIN, "callback_data_1")
				button2 := tgbotapi.NewInlineKeyboardButtonData("Профиль"+EMOJI_USER, "callback_data_2")
				button3 := tgbotapi.NewInlineKeyboardButtonURL("FAQ"+EMOJI_SOS, "https://t.me/+bke3X2XVmlthYjky")
				button4 := tgbotapi.NewInlineKeyboardButtonURL("Отзывы"+EMOJI_KUBOK, "https://t.me/+LVXFHgP7yv83MTk6")

				// Создаем клавиатуру с кнопками
				keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button1, button2), tgbotapi.NewInlineKeyboardRow(button4, button3))
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Выберите опцию:")
				msg.ReplyMarkup = keyboard

				// Отправляем обновленное сообщение
				if lastMessageID != 0 {
					deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, lastMessageID)
					_, err := gBot.Request(deleteMsg)
					if err != nil {
						log.Printf("Failed to delete message: %v", err)
					}
				}
				sentMessage, err := gBot.Send(msg)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}
				lastMessageID = sentMessage.MessageID
			case "callback_data_2":
				user := update.CallbackQuery.From
				profileInfo := fmt.Sprintf(
					"👤 Ваш профиль:\n\nИмя: %s\nID: %d\nUsername: @%s",
					user.FirstName,
					user.ID,
					user.UserName,
				)
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, profileInfo)
				gBot.Send(msg)
			}

		}

	}
}
