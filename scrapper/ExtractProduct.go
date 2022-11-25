package scrapper

import (
	"webs/structs"

	"github.com/gocolly/colly"
)

func ExtractProductData(url string, productChan chan structs.Product) {
	defer close(productChan)
	product := structs.Product{}
	collyProductCollector := colly.NewCollector()

	collyProductCollector.OnHTML(".ui-pdp-title", func(e *colly.HTMLElement) {
		product.Name = e.DOM.Text()
	})

	collyProductCollector.OnHTML(".andes-money-amount__fraction", func(e *colly.HTMLElement) {
		product.Price = e.DOM.Text()
	})

	product.Link = url

	collyProductCollector.Visit(url)
	productChan <- product
}
