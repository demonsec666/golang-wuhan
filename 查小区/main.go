package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/johnmccabe/go-bitbar"
)

func main() {
	app := bitbar.New()

	submenu := app.NewSubMenu()

	resp, err := http.Get("http://m.bj.bendibao.com/news/xiaoquchaxun/?p=detail")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	// fmt.Println(string(body))
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find(".data").Each(func(i int, contentSelection *goquery.Selection) {

		title := contentSelection.Find(".data-item").Text()
		titlea := strings.Replace(title, "\n                    ", "", -1)
		titlee := strings.Replace(titlea, "确诊人数                  ", "*", -1)
		titleb := strings.Replace(titlee, "治愈人数                  ", "*", -1)
		titlef := strings.Replace(titleb, "死亡人数                  ", "", -1)
		feiyan := strings.Split(titlef, "*")
		quezhen := feiyan[0]
		zhiyu := feiyan[1]
		siwang := feiyan[2]
		all_city := "确诊: " + quezhen + " 治愈: " + zhiyu + " 死亡: " + siwang
		app.StatusLine(all_city)
	})
	dom.Find(".list .list-border").Each(func(i int, contentSelection *goquery.Selection) {
		str := contentSelection.Find(".list-title").Text()
		title := contentSelection.Find(".list-detail a").Text()
		title = strings.Replace(title, "\n                                                    \n                                                \n                                                        ", "*", -1)
		title = strings.Replace(title, "\n                                                    \n                                                ", "", -1)
		title = strings.Replace(title, "\n                                                        ", "", -1)
		a := strings.Split(title, "*")
		b := strconv.Itoa(len(a))
		c := str + " " + b + "处"
		submenu.Line(c).Color("black")

		for i := 0; i < len(a); i++ {
			subsubmenu := submenu.NewSubMenu()
			subsubmenu.Line(a[i]).Color("black")
		}

	})
	app.Render()
}
