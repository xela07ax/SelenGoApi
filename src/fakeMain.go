package main

import (
	"fmt"
	"github.com/tebeka/selenium"
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
	if err := wd.Get("http://localhost:3036/"); err != nil {
		panic(err)
	}
	GetElement(selenium.ByID,"qf_region", wd).SendKeys(regionNum)
	GetElement(selenium.ByID,"qf_detailedDescription", wd).SendKeys(detailedDescription)
GetElement(selenium.ByID,"qf_fio", wd).SendKeys(fio)
GetElement(selenium.ByID,"qf_phone", wd).SendKeys(phone)
GetElement(selenium.ByID,"qf_dateTicket", wd).SendKeys(dateTicket)
GetElement(selenium.ByID,"qf_filePath", wd).SendKeys(filePath)
GetElement(selenium.ByID,"qf_email", wd).SendKeys(email)
GetElement(selenium.ByID,"qf_shortDescription", wd).SendKeys(shortDescription)
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