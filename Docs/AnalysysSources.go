package src

import (
	"github.com/tebeka/selenium/chrome"
	"github.com/tebeka/selenium/log"
	"time"
)

// Краткий анализ исходного кода:
// Переведено с https://blog.51cto.com/9406836/2499426

// Первая категория Разное

//Удалить разговор
func DeleteSession(urlPrefix, id string) error

// Включаем и отключаем отладочную отладку
func SetDebug(debug bool)

// Установить прокси
func (c Capabilities) AddProxy(p Proxy)
// Установить уровень журнала
func (c Capabilities) SetLogLevel(typ log.Type, level log.Level)

// Второй тип фонового сервиса seleium

// Необязательно для экземпляра службы
type ServiceOption func(*Service) error

// Добавить информацию о пути Chrome, вернуть тип serviceOption
type ServiceOption func(*Service) error

// Структура сервиса, включая скрытые типы
type Service struct {
// содержит отфильтрованные или неэкспортированные поля
}

// Запускаем сервер браузера Chrome и возвращаем указатель типа сервиса
func NewChromeDriverService(path string, port int, opts ...ServiceOption) (*Service, error)

// Закрыть сервис, не забудьте отложить закрытие
func (s *Service) Stop() error

// Третий тип операции Chrome

// Установка совместимости браузера, типа карты, например совместимости браузера Chrome.
// caps: = selenium.Capabilities {"browserName": "chrome"}
type Capabilities map[string]interface{}

// Добавляем совместимость с Chrome, вызывая функцию
func (c Capabilities) AddChrome(f chrome.Capabilities)

// Запускаем экземпляр веб-драйвера
func NewRemote(capabilities Capabilities, urlPrefix string) (WebDriver, error)

// Через интерфейс WebDriver вы можете увидеть метод реализации конкретной страницы, которая является интерфейсом.
type WebDriver interface {
	// Возвращаем информацию о версии серверной среды
	// Статус возвращает различную информацию о серверной среде.
	Status() (*Status, error)

	// Создать новую сессию
	// NewSession начинает новый сеанс и возвращает идентификатор сеанса.
	NewSession() (string, error)

	// Создать новый сеанс (не рекомендуется)
	// SessionId возвращает идентификатор текущей сессии
	// Устаревший: этот идентификатор неверен в стиле Go. Использовать SessionID
	// вместо
	SessionId() string

	// Получить новый ssionid
	// SessionID возвращает идентификатор текущей сессии.
	SessionID() string

	// Переключить сессию
	// SwitchSession переключается на данный идентификатор сессии.
	SwitchSession(sessionID string) error

	// Возвращаем совместимость
	// Возможности возвращают возможности текущего сеанса.
	Capabilities() (Capabilities, error)

	// Установить время выполнения асинхронного скрипта
	// SetAsyncScriptTimeout устанавливает количество времени, которое асинхронные сценарии
	// разрешено запускать до того, как они прерваны. Время ожидания будет округлено
	// с точностью до миллисекунды.
	SetAsyncScriptTimeout(timeout time.Duration) error

	// Установите время ожидания поискового элемента, цель: если результат страницы возвращается медленно, вам нужно дождаться полного возврата содержимого страницы, а затем выполнить операцию элемента страницы.
	// SetImplicitWaitTimeout устанавливает количество времени, которое драйвер должен ждать, когда
	// поиск элементов. Время ожидания округляется до ближайшей миллисекунды.
	SetImplicitWaitTimeout(timeout time.Duration) error

	// Устанавливаем время ожидания страницы
	// SetPageLoadTimeout устанавливает количество времени, которое драйвер должен ждать, когда
	// загрузка страницы. Время ожидания округляется до ближайшей миллисекунды.
	SetPageLoadTimeout(timeout time.Duration) error

	// Установить сессию для выхода
	// Выход завершает текущий сеанс. Экземпляр браузера будет закрыт.
	Quit() error

	// Получить текущий дескриптор окна, последовательность чисел, открыть окно и дескриптор
	// CurrentWindowHandle возвращает идентификатор текущего дескриптора окна.
	CurrentWindowHandle() (string, error)

	// Получить все дескрипторы открытого окна, получить все дескрипторы окна
	// WindowHandles возвращает идентификаторы текущих открытых окон.
	WindowHandles() ([]string, error)

	// Возвращаем URL текущей страницы подключения
	// CurrentURL возвращает текущий URL браузера.
	CurrentURL() (string, error)

	// Получить заголовок текущей страницы
	// Заголовок возвращает заголовок текущей страницы.
	Title() (string, error)

	// Возвращаем все содержимое текущей страницы
	// PageSource возвращает источник текущей страницы.
	PageSource() (string, error)

	// Закрыть текущее окно
	// Закрыть закрывает текущее окно.
	Close() error

	// Переключение фрейма, полный фреймворк внедряется во фрейм, если контент в операции должен ввести iframe. switchframe (ноль), возврат на верхний уровень
	// SwitchFrame переключается на данный кадр. Параметр frame может быть
	// идентификатор кадра в виде строки, его экземпляр WebElement возвращается
	// GetElement или nil для переключения на текущий контекст просмотра верхнего уровня.
	SwitchFrame(frame interface{}) error

	// Переключить окна на указанное окно
	// SwitchWindow переключает контекст на указанное окно.
	SwitchWindow(name string) error

	// закрываем окно
	// CloseWindow закрывает указанное окно.
	CloseWindow(name string) error

	// Установить развернутое окно
	// MaximizeWindow максимизирует окно. Если имя пустое, текущий
	// окно будет развернуто.
	MaximizeWindow(name string) error

	// Установить размер окна
	// ResizeWindow изменяет размеры окна. Если имя пустое,
	// текущее окно будет развернуто.
	ResizeWindow(name string, width, height int) error

	// Перейти к соответствующему интерфейсу через URL. Основной вариант - открыть адрес URL.
	// Get перемещает браузер к предоставленному URL.
	Get(url string) error

	// поворот вперед
	// Вперед движется вперед в истории.
	Forward() error

	// поворот назад
	// Назад движется назад в истории.
	Back() error

	// Обновить
	// Refresh обновляет страницу.
	Refresh() error

	// Найти и найти HTML-элемент.
	// FindElement находит ровно один элемент в DOM текущей страницы.
	FindElement(by, value string) (WebElement, error)

	// Найти и найти несколько HTML-элементов
	// FindElement находит потенциально много элементов в DOM текущей страницы.
	FindElements(by, value string) ([]WebElement, error)

	// Получить текущий элемент фокуса
	// ActiveElement возвращает текущий активный элемент на странице.
	ActiveElement() (WebElement, error)

	// DecodeElement декодирует ответ одного элемента.
	DecodeElement([]byte) (WebElement, error)

	// Декодируем ответы нескольких элементов
	// DecodeElements декодирует многоэлементный ответ.
	DecodeElements([]byte) ([]WebElement, error)

	// Получить все куки
	// GetCookies возвращает все куки в банке браузера.
	GetCookies() ([]Cookie, error)

	// Получить указанный файл cookie
	// GetCookie возвращает указанный файл cookie в банке, если он есть.
	// реализовано только для Firefox.
	GetCookie(name string) (Cookie, error)

	// Добавить куки в банку
	// AddCookie добавляет куки в банку браузера.
	AddCookie(cookie *Cookie) error

	// Удалить все куки
	// DeleteAllCookies удаляет все куки в банке браузера.
	DeleteAllCookies() error

	// Удалить указанный файл cookie
	// DeleteCookie удаляет куки в банку браузера.
	DeleteCookie(name string) error

	// Нажмите кнопку мыши
	// Щелчок щелкает кнопку мыши. Кнопка должна быть одной из RightButton,
	// MiddleButton или LeftButton.
	Click(button int) error

	// Двойной щелчок мышью
	// DoubleClick дважды щелкает левой кнопкой мыши.
	DoubleClick() error

	// Нажмите мышку
	// ButtonDown вызывает нажатие левой кнопки мыши.
	ButtonDown() error

	// Поднимаем мышь
	// ButtonUp вызывает отпускание левой кнопки мыши.
	ButtonUp() error

	// Отправляем изменения в активный элемент (отбрасывается)
	// SendModifier отправляет ключ модификатора активному элементу.
	// может быть одним из ShiftKey, ControlKey, AltKey, MetaKey.
	//
	// Устаревший: используйте взамен KeyDown или KeyUp.
	SendModifier(modifier string, isDown bool) error

	// Отправить последовательность клавиш активному элементу
	// KeyDown отправляет последовательность нажатий клавиш активному элементу.
	// похож на SendKeys, но без неявного завершения.
	// не освобождается в конце каждого вызова.
	KeyDown(keys string) error

	// Освобождаем отправленный элемент
	// KeyUp указывает, что предыдущее нажатие клавиши, отправленное KeyDown, должно быть
	// релиз
	KeyUp(keys string) error

	// сделать снимок
	// Снимок экрана делает снимок окна браузера.
	Screenshot() ([]byte, error)

	// Журнал соскоб
	// Журнал извлекает журналы. Типы журналов должны быть предварительно настроены в
	// возможности.
	//
	// ПРИМЕЧАНИЕ: вернет ошибку (не реализовано) в драйверах IE11 или Edge.
	Log(typ log.Type) ([]log.Message, error)

	//Все чисто
	// DismissAlert отклоняет текущее предупреждение.
	DismissAlert() error

	// Принять предупреждение
	// AcceptAlert принимает текущее предупреждение.
	AcceptAlert() error

	// Возврат к текущему содержанию тревоги
	// AlertText возвращает текущий текст предупреждения.
	AlertText() (string, error)

	// Отправить содержание оповещения
	// SetAlertText устанавливает текущий текст предупреждения.
	SetAlertText(text string) error

	// Выполнить скрипт
	// ExecuteScript выполняет скрипт.
	ExecuteScript(script string, args []interface{}) (interface{}, error)

	// Выполнить скрипт асинхронно
	// ExecuteScriptAsync асинхронно выполняет скрипт.
	ExecuteScriptAsync(script string, args []interface{}) (interface{}, error)

	// Выполнить исходный скрипт
	// ExecuteScriptRaw выполняет сценарий, но не выполняет JSON-декодирование.
	ExecuteScriptRaw(script string, args []interface{}) ([]byte, error)

	// Выполнить исходный скрипт асинхронно
	// ExecuteScriptAsyncRaw асинхронно выполняет скрипт, но не
	// выполнить декодирование JSON.
	ExecuteScriptAsyncRaw(script string, args []interface{}) ([]byte, error)

	// условие ожидания верно
	// WaitWithTimeoutAndInterval ожидает, когда условие оценивается как true.
	WaitWithTimeoutAndInterval(condition Condition, timeout, interval time.Duration) error

	//время ожидания
	// WaitWithTimeout работает как WaitWithTimeoutAndInterval, но с интервалом опроса по умолчанию.
	WaitWithTimeout(condition Condition, timeout time.Duration) error

	//Подождите
	// Ожидание работает как WaitWithTimeoutAndInterval, но использует время ожидания по умолчанию и интервал опроса.
	Wait(condition Condition) error
}

// Последующее выполнение связанных элементов, тип интерфейса, который является методом реализации
type WebElement interface {
	// щелкнуть выбранный элемент
	// Кликаем по элементу.
	Click() error

	// Отправляем данные на выбранный элемент
	// SendKeys вводит в элемент.
	SendKeys(keys string) error

	// Отправить кнопку
	// Отправить отправляет кнопку.
	Submit() error

	// Пустая кнопка
	// Clear очищает элемент.
	Clear() error

	// Переместить элемент в соответствующие координаты
	// MoveTo перемещает мышь к относительным координатам от центра элемента, если
	// элемент не виден, он будет прокручиваться в поле зрения.
	MoveTo(xOffset, yOffset int) error

	// Найти дочерние элементы
	// FindElement находит дочерний элемент.
	FindElement(by, value string) (WebElement, error)

	// Найти несколько дочерних элементов
	// FindElement находит несколько дочерних элементов.
	FindElements(by, value string) ([]WebElement, error)

	// Возвращаем имя метки
	// TagName возвращает имя элемента.
	TagName() (string, error)

	// Возвращаем содержимое элемента
	// Text возвращает текст элемента.
	Text() (string, error)

	// Элемент выбран и возвращает true
	// IsSelected возвращает true, если элемент выбран.
	IsSelected() (bool, error)

	// Возвращаем true, если элемент включен
	// IsEnabled возвращает true, если элемент включен.
	IsEnabled() (bool, error)

	// IsDisplayed возвращает true, если элемент отображается.
	IsDisplayed() (bool, error)
	// Получить имя элемента
	// GetAttribute возвращает именованный атрибут элемента.
	GetAttribute(name string) (string, error)

	// Положение элемента диапазона
	// Location возвращает местоположение элемента.
	Location() (*Point, error)

	// Возвращаем позицию элемента после прокрутки
	// LocationInView возвращает местоположение элемента после его прокрутки
	// в поле зрения
	LocationInView() (*Point, error)

	// Возвращаем размер элемента
	// Размер возвращает размер элемента.
	Size() (*Size, error)

	// Возвращаем приоритет css
	// CSSProperty возвращает значение указанного свойства CSS
	// элемент.
	CSSProperty(name string) (string, error)

	// Возвращаем моментальный снимок прокрутки атрибута
	// Снимок экрана делает снимок экрана атрибута scroll'ing, если это необходимо.
	Screenshot(scroll bool) ([]byte, error)
}