package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	godotenv.Load()

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "To add torrent write command:\n\n /add url_or_magnet_link_here")
	})

	b.Handle("/help", func(m *tb.Message) {
		b.Send(m.Sender, "To add torrent write command: /add url_or_magnet_link_here")
	})

	b.Handle("/add", func(m *tb.Message) {
		link_arr := strings.Split(m.Text, " ")

		if len(link_arr) < 2 {
			fmt.Println("No parameters")
			b.Send(m.Sender, "No parameters")
			return
		}

		req := CreateRequest(link_arr[1])
		result := SendRequest(req)
		b.Send(m.Sender, result)
	})

	b.Start()

}

func CreateRequest(link string) *http.Request {
	qbit_url := "http://127.0.0.1:8080/api/v2/torrents/add"

	data := url.Values{}
	data.Set("urls", link)
	data.Set("autoTMM", "false")
	data.Set("savepath", os.Getenv("TORRENT_SAVE_PATH"))
	data.Set("cookie", "")
	data.Set("rename", "")
	data.Set("category", "")
	data.Set("paused", "false")
	data.Set("root_folder", "true")
	data.Set("dlLimit", "NaN")
	data.Set("upLimit", "NaN")

	req, err := http.NewRequest("POST", qbit_url, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	return req
}

func SendRequest(req *http.Request) string {

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Qbittorrent web interface don't respond.")
		log.Print(err)
		return "Qbittorrent web interface don't respond. Check is it working..."
	}

	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
