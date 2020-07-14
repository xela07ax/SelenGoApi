package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"strings"
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

	// 3. Открыть многостраничный экземпляр Chrome
	// В настоящее время я думаю о двух способах открытия,
	// Во-первых, на странице есть URL-соединение, которое открывается щелчком click().
	// Второй способ - открыть его скриптом. wd.ExecuteScript
	if err := wd.Get("http://cdn1.python3.vip/files/selenium/sample3.html"); err != nil {
		panic(err)
	}

	// Первый способ - найти адрес URL на странице и перейти на страницу.
	we, err := wd.FindElement(selenium.ByTagName, "a")
	if err != nil {
		panic(err)
	}
	we.Click()

	// Второй способ - открыть новое окно, запустив обычный js-скрипт, потому что нам пока не нужно получать результат, полученный операцией, поэтому мы не получаем возвращаемое значение.
	wd.ExecuteScript(`window.open("https://www.qq.com", "_blank");`, nil)
	wd.ExecuteScript(`window.open("https://www.runoob.com/jsref/obj-window.html", "_blank");`, nil)


	// Эта строка предназначена для отправки предупреждающего сообщения. Цель написания этой строки - узнать, какое из них является текущим главным окном.
	wd.ExecuteScript(`window.alert(location.href);`, nil)

	// Просмотр значения дескриптора текущего окна
	handle, err := wd.CurrentWindowHandle()
	if err != nil {
		panic(err)
	}
	fmt.Println(handle)
	fmt.Println("--------------------------")

	// Просмотр значения дескриптора всех страниц
	handles, err := wd.WindowHandles()
	if err != nil {
		panic(err)
	}
	for _, handle := range handles {
		fmt.Println(handle)
	}
	fmt.Println("--------------------------")

	// 4. Перейти на указанную веб-страницу
	// Хотя мы открыли несколько страниц, наше текущее значение дескриптора по-прежнему является первой страницей, мы должны найти способ получить его.
	// Не забудьте сохранить значение дескриптора текущей главной страницы
	// mainhandle: = handle

	// Соответствующая веб-страница определяется условием суждения
	// Получить все значения дескриптора

	handles, err = wd.WindowHandles()
	if err != nil {
		panic(err)
	}

	// Обходим все значения дескриптора, находим целевую страницу по URL, а когда они равны, вырываемся и останавливаемся на соответствующей странице.
	for _, handle := range handles {
		wd.SwitchWindow(handle)
		url, _ := wd.CurrentURL()
		if strings.Contains(url, "qq.com") {
			break
		}
	}
	// Просмотр дескриптора этой страницы
	handle, err = wd.CurrentWindowHandle()
	if err != nil {
		panic(err)
	}
	fmt.Println(handle)
	// Эта строка предназначена для отправки предупреждающего сообщения. Цель написания этой строки - узнать, какое из них является текущим главным окном.
	wd.ExecuteScript(`window.alert(location.href);`, nil)

	// Вернемся к первой странице
	//wd.SwitchWindow(mainhandle)

	// Выход через 20 секунд сна
	time.Sleep(20 * time.Second)
}