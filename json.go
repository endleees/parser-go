package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

type Item struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

func main() {
	// Отправляем GET-запрос на веб-страницу для парсинга
	array := [5]string{
		"https://arpa-hpl.ru/catalog/arpa-hpl/",
		"https://www.arpa-hpl.ru/catalog/arpa-hpl/?PAGEN_1=2",
		"https://www.arpa-hpl.ru/catalog/arpa-hpl/?PAGEN_1=3",
		"https://www.arpa-hpl.ru/catalog/arpa-hpl/?PAGEN_1=4",
		"https://www.arpa-hpl.ru/catalog/arpa-hpl/?PAGEN_1=5",
	}
	for _, s := range array {
		response, err := http.Get(s)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		// Загружаем HTML-страницу для парсинга
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Создаем слайс для хранения извлеченных данных
		items := []Item{}

		doc.Find(".item-wrap").Each(func(index int, element *goquery.Selection) {
			img, _ := element.Find("img").Attr("src")

			title := element.Find(".name").Text()
			// Создаем новый элемент и добавляем его в слайс
			item := Item{
				Title: title,
				Image: "https://arpa-hpl.ru" + img,
			}
			items = append(items, item)
		})
		file, err := os.OpenFile("output.json", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Декодируем JSON-файл в структуру данных
		existingItems := []Item{}
		err = json.NewDecoder(file).Decode(&existingItems)
		if err != nil {
			log.Fatal(err)
		}

		// Добавляем новые данные к существующим
		items = append(existingItems, items...)

		// Устанавливаем указатель файла в начало
		_, err = file.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}

		// Записываем обновленные данные в файл
		err = file.Truncate(0)
		if err != nil {
			log.Fatal(err)
		}

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		err = encoder.Encode(items)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Данные успешно добавлены в output.json")
	}

}
