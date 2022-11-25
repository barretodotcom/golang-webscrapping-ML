package scrapper

import (
	"strconv"
	"strings"
	"webs/structs"
)

func FormatUrls(product string, url string, index int8, finalIndex int8, structChan chan structs.Product) {

	defer close(structChan)

	var linksChan []chan []structs.Product

	product = strings.Replace(product, " ", "-", -1)
	url = strings.Replace(url, "<<PRODUCT>>", product, 1)

	for i := 0; i <= int(finalIndex); i++ {
		linksChan = append(linksChan, make(chan []structs.Product))
		pageUrl := strings.Replace(url, "<<INDEX>>", strconv.FormatInt(int64(i), 8), 1)
		go ScrapUrls(pageUrl, linksChan[i])
	}
	for _, v := range linksChan {
		productValue := <-v

		for i, _ := range productValue {
			structChan <- productValue[i]
		}
	}
}
