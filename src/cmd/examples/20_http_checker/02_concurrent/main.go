package main

import (
	"fmt"
	"net/http"
	"sync"
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

	var wg sync.WaitGroup
	var mu sync.Mutex
	var erros []string

	inicio := time.Now()

	fmt.Printf("%sIniciando verificações...%s\n", Yellow, Reset)

	for _, url := range urls {
		wg.Add(1)
		go verifyServiceStatus(url, &wg, &mu, &erros)
	}

	wg.Wait()

	fmt.Printf("%sVerificações finalizado. Tempo: %.2f%s\n",
		Yellow, time.Since(inicio).Seconds(), Reset)

	if len(erros) > 0 {
		fmt.Printf("\n%sErros encontrados:%s\n", Red, Reset)
		for _, e := range erros {
			fmt.Printf("%s- %s%s\n", Red, e, Reset)
		}
	}
}

func verifyServiceStatus(url string, wg *sync.WaitGroup, mu *sync.Mutex, erros *[]string) {
	defer wg.Done()

	start := time.Now()

	res, err := http.Get(url)
	periodo := time.Since(start).Seconds()

	if err != nil {
		mu.Lock()
		*erros = append(*erros, fmt.Sprintf("%s[%s]%s — %s[%.2fs]%s",
			Blue, url, Reset,
			Yellow, periodo, Reset))
		mu.Unlock()
		return
	}
	defer res.Body.Close()

	fmt.Printf("%sSucesso%s ao consultar o serviço — %s[%s]%s — %s[%.2fs]%s\n",
		Green, Reset,
		Blue, url, Reset,
		Yellow, periodo, Reset)
}
