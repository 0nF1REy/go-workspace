package checker

import (
	"time"

	"github.com/0nF1REy/go-workspace/services/bubble_tea/01_health_checker/internal/model"
)

var services = []string{
	"https://www.google.com",
	"https://www.github.com",
	"https://www.youtube.com",
	"https://www.cursor.com",
	"https://www.microsoft.com",
	"https://www.chatgpt.com",
	"https://www.linkedin.com",
	"https://jashinchan.com",
	"https://kill-la-kill.jp",
	"https://www.toei-anim.co.jp",
	"https://www.ghibli.jp",
	"https://www.tv-tokyo.co.jp",
	"https://www.crunchyroll.com",
	"https://www.aniplex.co.jp",
	"https://www.nhk.or.jp",
	"https://www.sony.co.jp",
	"https://www.nintendo.co.jp",
	"https://www.bandai.co.jp",
	"https://www.shonenjump.com",
	"https://www.pixiv.net",
	"https://www.nicovideo.jp",
	"https://amnibus.com",
	"https://www.animate-onlineshop.jp",
	"https://p-bandai.jp",
	"https://www.geestore.com",
	"https://hicul.jp",
	"https://meqri.com",
	"https://www.amiami.com",
	"https://www.djsdushds.com",
}

func Start() chan model.Result {
	out := make(chan model.Result)

	go func() {
		ticker := time.NewTicker(5 * time.Second)

		for {
			runChecks(services, out)
			<-ticker.C
		}
	}()
	return out
}
