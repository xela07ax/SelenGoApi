package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/xela07ax/toolsXela/tp"
	"time"
)

const (
	//设置常量 分别设置chromedriver.exe的地址和本地调用端口
	seleniumPath = `resources/chromedriver.exe`
	port         = 9515
	regionNum = "101"
	fio = "Иванов Иван Иванович"
	phone = "84950000000"
	email = "ivanov@ivan.ivanovich"
	shortDescription = "Тест скороти заведения заявки"
	detailedDescription = "Подробный текст: Тест скороти заведения заявки"
	dateTicket = "10.07.2020 22:14:04"
	filePath = "C:\\windows-version.txt"
)

func main() {
	//1.开启selenium服务
	//设置selium服务的选项,设置为空。根据需要设置。
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	//延迟关闭服务
	defer service.Stop()

	//2.调用浏览器
	//设置浏览器兼容性，我们设置浏览器名称为chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//调用浏览器urlPrefix: 测试参考：DefaultURLPrefix = "http://127.0.0.1:4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		panic(err)
	}
	//延迟退出chrome
	defer wd.Quit()
	if err := wd.Get("http://localhost:3456"); err != nil {
		panic(err)
	}

	FindSetID("username","user123", wd)
	FindSetID("password","pa$$word", wd)
	we, err := wd.FindElement(selenium.ByID, "loginbutton")
	if err != nil {
		panic(err)
	}
	we.Click()
	fmt.Println("Входим")
	btn := GetElement(selenium.ByLinkText,"Новая заявка", wd)
	btn.Click()
	// Эта  кнопка бывает не срабатывает, сделаем повтор если не сработала
	we, err = GetElementWithErr(selenium.ByID,"m79c53660-img", wd)
	if err != nil {
		btn := GetElement(selenium.ByLinkText,"Новая заявка", wd)
		btn.Click()
		we, err = GetElementWithErr(selenium.ByID,"m79c53660-img", wd)
	}
	//w_b is a WebDriver
	GetElement(selenium.ByID,"m79c53660-img", wd).Click()
	GetElement(selenium.ByID,"menu0_REG_OPTION_a_tnode", wd).Click()
	GetElement(selenium.ByID,"m3644c556-img", wd).Click()
	GetElement(selenium.ByID,"menu0_blank__OPTION_a", wd).Click()


	// Выбираем регион
	GetElement(selenium.ByID,"mdcb35a0b-img", wd).Click()
	we = GetElement(selenium.ByID,"lookup_page1_tfrow_\\[C\\:0\\]_txt-tb", wd)
	fmt.Println("Полле ввода региона")
	err = we.SendKeys("101\n")
	if err != nil {
		panic(err)
	}
	//GetElement(selenium.ByID,"lookup_page1_tdrow_\\[C\\:0\\]-c\\[R\\:0\\]", wd).Click()
	var c selenium.Condition= func(wb selenium.WebDriver)(bool,error){
		//when i delete the line, it can wait
		var err error
		we,err = wb.FindElement(selenium.ByID, "lookup_page1_tbod-tbd")
		if err != nil {return true,nil}
		bt, _ := we.Screenshot(true)
		f,_ := tp.OpenWriteFile("file.png")
		f.Write(bt)
		wes,err := wb.FindElements(selenium.ByCSSSelector, "#lookup_page1_tbod-tbd tbody tr")
		if (len(wes) < 3) || (len(wes) > 4){
			fmt.Println("Ошибка при чтении региона со страницы")
			return false,err
		} else if len(wes) == 3{
			fmt.Println(`Ошибка, регион ${regionNum} ненайден`)
			return false,err
		} else {
			fmt.Println(`Регион ${regionNum} найден`)
			we, err = wes[3].FindElement(selenium.ByCSSSelector,"span")
			if err != nil {return false,err}
			we.Click()
		}
		fmt.Println("Полле региона click")
		return false,err
	}
	_ = wd.WaitWithTimeoutAndInterval(c,40*time.Second,1*time.Second)
	time.Sleep(1*time.Second)
	FindSetID("m448773e5-tb",fio, wd)
	FindSetID("m928d2451-tb",phone, wd)
	FindSetID("m97c232d4-tb",email, wd)
	FindSetID("m8672e47c-tb",shortDescription, wd)
	FindSetID("m1f7bb5c6-ta",detailedDescription, wd)

	GetElement(selenium.ByCSSSelector,"button#m74a6b003_bg_button_ssaddnewattachmentfile-pb", wd).Click()
	err = wd.SwitchFrame("upload_iframe")
	if err != nil {panic(err)}
	we, err = wd.FindElement(selenium.ByID, "file")
	if err != nil {
		panic(err)
	}
	//Send '' to input box
	err = we.SendKeys("C:\\windows-version.txt")
	if err != nil {
		panic(err)
	}

	//FindSetID("file",filePath, wd)
	// Переключиться обратно на фрейм верхнего уровня
	err = wd.SwitchFrame(nil)
	if err != nil {panic(err)}
	GetElement(selenium.ByCSSSelector,"button#me7a99c75-pb", wd).Click()
	FindSetID("m1090c1af-tb",dateTicket, wd)
	FindSetID("m29404739-tb","88.", wd)
	time.Sleep(1*time.Hour)

}
func GetElement(bySelector string, selector string, wd selenium.WebDriver) selenium.WebElement  {
	we, err := GetElementWithErr(bySelector,selector,wd)
	if err != nil {
		panic(err)
	}
	return we
}

func GetElementWithErr(bySelector string, selector string, wd selenium.WebDriver) (selenium.WebElement,error)  {
	var c selenium.Condition= func(wb selenium.WebDriver)(bool,error){
		//when i delete the line, it can wait
		fmt.Printf("Ищем:%s\n",selector)
		var err error
		_,err = wb.FindElement(bySelector, selector)
		if err != nil {
			return false,nil
		}
		return true,err
	}
	var err error
	err = wd.WaitWithTimeoutAndInterval(c,20*time.Second,1*time.Second)
	if err != nil {
		return nil,err
	}
	var we selenium.WebElement
	we,err = wd.FindElement(bySelector, selector)
	return we,err
}

func FindSetID(id string,val string, wd selenium.WebDriver) {
	//Find Baidu input box id
	we, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		panic(err)
	}
	//Send '' to input box
	err = we.SendKeys(val)
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
}