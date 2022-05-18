package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"imageUrl"`
}

func main() {

	c := colly.NewCollector()

	items := []item{}

	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := item{
			Name:     h.ChildText("h2.product-title"),
			Price:    h.ChildText("div.sale-price"),
			ImageUrl: h.ChildAttr("img", "src"),
		}
		items = append(items, item)
	})

	c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
		nextPage := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://j2store.net/demo/index.php/shop")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("products.json", content, 0644)
}
