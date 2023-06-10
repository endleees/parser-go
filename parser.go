package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Отправляем GET-запрос на веб-страницу
	resp, err := http.Get("https://arpa-hpl.ru/catalog/arpa-hpl/")
	if err != nil {
		log.Fatal("Ошибка при выполнении GET-запроса:", err)
	}
	defer resp.Body.Close()

	// Загружаем HTML документ с использованием goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Ошибка при загрузке HTML:", err)
	}

	// Извлекаем все теги <a> и выводим их содержимое и атрибуты
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// Извлекаем текст тега <a>
		text := strings.TrimSpace(s.Text())

		// Извлекаем значение атрибута "href"
		href, exists := s.Attr("href")
		if exists {
			fmt.Printf("Текст: %s\n", text)
			fmt.Printf("Ссылка: %s\n", href)
			fmt.Println("-----------------")
		}
	})

	// Извлекаем все теги <img> и выводим их атрибуты
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		// Извлекаем значение атрибута "src"
		src, exists := s.Attr("src")
		if exists {
			fmt.Printf("Изображение: https://arpa-hpl.ru%s\n", src)
			fmt.Println("-----------------")
		}
	})
}
