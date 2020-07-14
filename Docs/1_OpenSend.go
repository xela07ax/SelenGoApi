package main

import (
"fmt"
"github.com/tebeka/selenium"
"time"
)

const (
	// Установить константы Установить адрес и локальный порт вызова chromedriver.exe
	seleniumPath = `H:\webdriver\chromedriver.exe`
	port         = 9515
)

func main() {
	// 1. Запустить службу selenium
	// Установить параметры для сервиса selium
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	// Задержка выключения сервиса
	defer service.Stop()

	// 2. Вызов браузера
	// Установить совместимость браузера, мы устанавливаем имя браузера в Chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	// Вызываем браузер urlPrefix: тестовая ссылка:
	// DefaultURLPrefix := "http://127.0.0.1:4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		panic(err)
	}
	// Задержка из-за хрома
	defer wd.Quit()

	// 3. Работа с элементами страницы
	// Получить страницу Baidu
	if err := wd.Get("https://www.baidu.com/"); err != nil {
		panic(err)
	}
	// Находим идентификатор поля ввода Baidu
	we, err := wd.FindElement(selenium.ByID, "kw")
	if err != nil {
		panic(err)
	}
	// Отправить "" в поле ввода
	err = we.SendKeys("Первый в мире")
	if err != nil {
		panic(err)
	}

	// Найти идентификатор кнопки отправки Baidu
	we, err = wd.FindElement(selenium.ByID, "su")
	if err != nil {
		panic(err)
	}
	// Нажмите, чтобы отправить
	err = we.Click()
	if err != nil {
		panic(err)
	}

	// Выход через 20 секунд сна
	time.Sleep(20 * time.Second)
}