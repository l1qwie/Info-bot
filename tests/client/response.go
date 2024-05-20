package client

import (
	"InfoBot/fmtogram/formatter"
	"log"
)

func logs(fm *formatter.Formatter) {
	log.Printf(`fm.Message.Text = "%s"`, fm.Message.Text)
	log.Printf(`fm.Keyboard.Keyboard = "%v"`, fm.Keyboard.Keyboard)
	log.Printf(`fm.Message.ChatID = %d`, fm.Message.ChatID)
}

func greeteings(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString("Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас!", true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Клиент", "Админ"}, []string{"client", "admin"}, []string{"cmd", "cmd"}, true)
}

func showSchedule(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString("Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас! На этой неделе у нас будут кое-какие активности, и мы будем вас с нетерпением ждать! Ниже нажмите на кнопку, чтобы узнать подробнее", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1, 1, 1, 1}, []string{"Богослужение", "Молодежка", "Домашняя Группа", "Молитвы", "Мужское&Женское", "Фильм"},
		[]string{"worship", "youthmeeting", "homegroups", "prayers", "men&women", "film"}, []string{"cmd", "cmd", "cmd", "cmd", "cmd", "cmd"}, true)
}

func showWorship(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждую субботу в 13:00 у нас в главном здании церкви начинается богослужение.
	На протяжении примерно 20-25 минут в начале мы играем и поем песни, дальше некоторое количество
	времени уделяется молитвам, а потом главная проповедь. Обычно богослужение заканчивается ближе к 15:00,
	но никто никуда не расходится, так как можно пообщаться и выпить чаю. Очень ждем именно Вас!
	
	Начало: каждая суббота в 13:00
	Проповедь будет вести: Ахмед Абдулов
	
	Мы находимся тут: 
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/data=!4m6!3m5!1s0x145287d03816e835:0xbae794404ccd749!8m2!3d27.2508627!4d33.8319995!16s%2Fg%2F11c2ldpk6q?entry=ttu`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
}

func showYouth(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждое воскресенье у нас проходят молодежные встречи на дому (очень похоже на домашнюю группу, но только для молодых).
	На протяжении около 2 часов мы общаемся, кушаем вкусности, иногда во что-нибудь играем и обсуждаем Библию. Если Вам от 12 до 35 лет, мы очень ждем Вас!
	
	Начало: каждое воскресенье в 16:00
	Ведущий: Ислам Мусульманов
	
	Геолокация иногда меняется, поэтому лучше приходить до 16:00 в наше главное здание и спрашивать у ведущего, где сегодня будет молодежная встреча
	
	Наше главное здание тут:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/data=!4m6!3m5!1s0x145287d03816e835:0xbae794404ccd749!8m2!3d27.2508627!4d33.8319995!16s%2Fg%2F11c2ldpk6q?entry=ttu`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
}

func showHomeGroups(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждый четверг некоторые наши прихожане устраивают домашние группы у себя на дому.
	Если вы хотите обсудить Библию и пообщаться в приятной компании, то вам определенно к нам!
	В среднем такие домашние группы длятся около 2 часов и начинаются ближе к 6-7 часам вечера.
	Расписание может варьироваться в зависимости от того, кто ведущий
	
	Группа 1:
	Ведущий: Антон Ахмедов
	Начало: 18:00
	Адрес: Scandic Resort Hurghada
	Ссылка на адрес: https://www.google.com/maps/place/Scandic+Resort+Hurghada/@27.2429138,33.8417842,18.21z/data=!4m12!1m5!3m4!2zMzbCsDUzJzM2LjQiTiAzMMKwNDInMzQuNSJF!8m2!3d36.893445!4d30.709591!3m5!1s0x14528790e20fc35b:0xfed8fd8e22bb8a2b!8m2!3d27.2431856!4d33.8431078!16s%2Fg%2F11kjzg21ft?entry=ttu
	
	Группа 2:
	Ведущий: Авраам Израилев
	Начало: 19:00
	Адрес: Иерусалим
	Ссылка на адрес: https://www.google.com/maps/place/ZARIFFA+-+%D7%91%D7%99%D7%AA+%D7%A7%D7%A4%D7%94+%D7%99%D7%A8%D7%95%D7%A9%D7%9C%D7%9E%D7%99%E2%80%AD/@31.7556663,35.2016948,18.16z/data=!4m6!3m5!1s0x1502d7e8f28c54b9:0x3cb970274e20cd2d!8m2!3d31.755949!4d35.2027168!16s%2Fg%2F11bxf42czk?entry=ttu`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)

}

func showPrayers(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждый вторник у нас в главном здании проводится часовая молитва.
	А еще в среду у нас онлайн-молитва в Zoom. Если у вас есть нужда или же вы хотите просто помолиться, то мы ждем вас!
	
	Начало молитвы в здании (Вт): 9:00
	Начало молитвы в Zoom (Ср): 8:00
	
	Мы находимся тут:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/data=!4m6!3m5!1s0x145287d03816e835:0xbae794404ccd749!8m2!3d27.2508627!4d33.8319995!16s%2Fg%2F11c2ldpk6q?entry=ttu
	
	Ссылка на Zoom будет отсылаться в наш общий церковный чат. Присоединяйтесь!`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)

}