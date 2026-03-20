package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func main() {
	urls := []string{
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

	inicio := time.Now()

	fmt.Printf("%sIniciando verificações...%s\n\n", Blue, Reset)

	for _, url := range urls {
		verifyServiceStatus(url)
	}

	fmt.Printf("\n%sFINALIZADO EM: %.2fs%s\n",
		Green, time.Since(inicio).Seconds(), Reset)
}

func verifyServiceStatus(url string) {
	start := time.Now()
	fmt.Printf("%s[START]%s - Iniciando verificação: %s\n", Yellow, Reset, url)

	res, err := http.Get(url)
	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Printf("%s[ERRO] %s - Está fora do ar ou é inválido! (%.2fs)%s\n",
			Red, url, elapsed, Reset)
		return
	}
	defer res.Body.Close()

	fmt.Printf("%s[OK] %s - Status: %d - Tempo: %.2fs%s\n",
		Green, url, res.StatusCode, elapsed, Reset)
}
