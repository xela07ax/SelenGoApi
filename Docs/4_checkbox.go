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

	// 3 radioс одиночным выбором, checkbox с множественным выбором, select выбора ящика (функция должна быть улучшена, https://github.com/tebeka/selenium/issues/141)
	if err := wd.Get("http://cdn1.python3.vip/files/selenium/test2.html"); err != nil {
		panic(err)
	}
	//3.1 Операция radio
	we, err := wd.FindElement(selenium.ByCSSSelector, `#s_radio > input[type=radio]:nth-child(3)`)
	if err != nil {
		panic(err)
	}
	we.Click()

	//3.2操作多选checkbox
	//Удалить checkbox по умолчанию
	we, err = wd.FindElement(selenium.ByCSSSelector, `#s_checkbox > input[type=checkbox]:nth-child(5)`)
	if err != nil {
		panic(err)
	}
	we.Click()
	// Выберите опцию
	we, err = wd.FindElement(selenium.ByCSSSelector, `#s_checkbox > input[type=checkbox]:nth-child(1)`)
	if err != nil {
		panic(err)
	}
	we.Click()
	we, err = wd.FindElement(selenium.ByCSSSelector, `#s_checkbox > input[type=checkbox]:nth-child(3)`)
	if err != nil {
		panic(err)
	}
	we.Click()

	//3.3 select多选
	// Удалить опцию по умолчанию

	// Выберите элемент по умолчанию
	we, err = wd.FindElement(selenium.ByCSSSelector, `#ss_multi > option:nth-child(3)`)
	if err != nil {
		panic(err)
	}
	we.Click()

	we, err = wd.FindElement(selenium.ByCSSSelector, `#ss_multi > option:nth-child(2)`)
	if err != nil {
		panic(err)
	}
	we.Click()

	// Выход через 20 секунд сна
	time.Sleep(20 * time.Second)
}