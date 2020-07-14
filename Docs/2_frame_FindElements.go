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
	// Получить тестовую страницу
	if err := wd.Get("http://cdn1.python3.vip/files/selenium/sample2.html"); err != nil {
		panic(err)
	}

	// 4. Переключиться на соответствующий frame
	//wd.SwitchFrame(webelement, который можно получить по id или frame), мы используем два метода для достижения отдельно.

	//4.1 Поиск по id frame, в это время id=frame1
	/*
	    err = wd.SwitchFrame("frame1")
	    if err != nil {
	        panic(err)
	    }

	    // На этом этапе найдите в HTML-код iframe, а затем используйте bycssselector
	   // Поскольку animal содержит несколько объектов, мы используем findelements
	    wes, err := wd.FindElements(selenium.ByCSSSelector, ".animal")
	    if err != nil {
	        panic(err)
	    }

	   // Цикл для получения информации о каждом
	    for _,we := range  wes {
	        text, err := we.Text()

	        if err != nil {
	            panic(err)
	        }
	        fmt.Println(text)
	    }
	*/

	//4.2 webelement, полученный frame, реализуется путем переключения webelement-a.
	// найти объект webelement из ifname
	element, err := wd.FindElement(selenium.ByCSSSelector, "#frame1")
	/// Различные способы получения элементов
	//element, err := wd.FindElement(selenium.ByCSSSelector, `iframe[name="innerFrame"]`)
	if err != nil {
		panic(err)
	}
	// Переключаемся на iframe
	err = wd.SwitchFrame(element)
	if err != nil {
		panic(err)
	}
	// На этом этапе найдите в HTML-код iframe, а затем используйте bycssselector
	// Поскольку animal содержит несколько
	wes, err := wd.FindElements(selenium.ByCSSSelector, ".animal")
	if err != nil {
		panic(err)
	}

	// Цикл для получения информации о каждом
	for _, we := range wes {
		text, err := we.Text()
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
	}

	// 5. Переключиться обратно на фрейм верхнего уровня, потому что внешним элементом значения нельзя манипулировать в фрейме во время переключения, поэтому мы должны его отключить
	// frame = nil - переключиться обратно на верхний фрейм
	err = wd.SwitchFrame(nil)
	if err != nil {
		panic(err)
	}
	// Выбор элементов на основе имени класса
	we, err := wd.FindElement(selenium.ByCSSSelector, ".baiyueheiyu")
	if err != nil {
		panic(err)
	}
	// Просмотр содержимого элемента верхнего
	fmt.Println(we.Text())

	// Выход через 20 секунд сна
	time.Sleep(20 * time.Second)
}