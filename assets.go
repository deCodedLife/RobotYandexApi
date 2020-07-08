package main

var (
	firstMessage          string   = "Здравствуйте, я проморобот. Вы можете управлять мной, или получать актуальную информацию."
	possibilities         string   = "Вы можете управлять мной используя комманды: Вперед, Назад, Налево, Направо или провести презентацию."
	presentationAnswer    string   = "ごめんなさい。 Сейчас данная функция на стадии разработки"
	presentationQuestions []string = []string{"презентация", "Презентация", "Проведи презентацию", "проведи презентацию", "начни презентацию", "Начни презегтацию", "провести презентацию", "Провести презентацию"}
	helpQuestions         []string = []string{"help", "Help", "Помощь", "помощь", "Инструкция", "инструкция", "Что делать?", "что делать?", "Что делать", "что делать"}
	abilityQuestions      []string = []string{"abilities", "Abilities", "possibilities", "Possibilities", "Возможности", "возможности", "Какие у тебя возможности", "какие у тебя возможности", "Какие у тебя возможности?", "какие у тебя возможности?", "Что ты умеешь", "что ты умеешь", "Что ты умеешь?", "что ты умеешь?"}
	controllCommands      []string = []string{"Вперед", "вперед", "Иди вперед", "иди вперед", "Назад", "назад", "Иди назад", "иди назад", "В лево", "в лево", "Влево", "влево", "На лево", "на лево", "В право", "в право", "На право", "на право"}
	stopCommands          []string = []string{"Выйти из программы", "выйти из программы"}
	timeQuestions         []string = []string{"Сколько время?", "сколько время?", "Сколько время", "сколько время", "сколько времени?", "Сколько времени?", "Узнай время", "узнай время", "Сколько сейчас?", "сколько сейчас?", "Сколько сейчас времени?", "сколько сейчас времени?", "Сколько сейчас времени", "сколько сейчас времени", "Время", "время", "Который час?", "который час?", "который час", "Который час", "なんじ", "何時", "なん時", "なんじなんぷん", "何時何分"}
)
