package main

import (
	"fmt"
	"log"
	"os"
	"webs/scrapper"
	"webs/structs"

	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	product := "VINHO"
	var index int64 = 1
	var finalIndex int64 = 3
	url := "https://lista.mercadolivre.com.br/<<PRODUCT>>_Desde_<<INDEX>>_NoIndex_True"

	productChan := make(chan structs.Product)
	products := []structs.Product{}

	go scrapper.FormatUrls(product, url, index, finalIndex, productChan)

	for prod := range productChan {
		products = append(products, prod)
	}
	csv, err := gocsv.MarshalString(products)

	if err != nil {
		log.Fatal(err.Error())
	}
	file, _ := os.Create("products.csv")

	file.Write([]byte(csv))
	fmt.Print("Processo finalizado.")
}
