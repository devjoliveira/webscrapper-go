package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"github.com/PuerkitoBio/goquery"
)

type Noticia struct {
	Titulo string
	Link   string
}

func webScrapper(url string, wg *sync.WaitGroup, ch chan<- Noticia) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erro ao realizar requisição HTTP: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Erro: o site retornou um status code inesperado: %d %s", resp.StatusCode, resp.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler o arquivo HTML: %v", err)
		return
	}

	doc.Find(".titleline > a").Each(func(i int, s *goquery.Selection) {
		titulo := s.Text()
		link, exists := s.Attr("href")

		if exists {
			if !strings.HasPrefix(link, "http") {
				link = "https://news.ycombinator.com/" + link
			}

			ch <- Noticia{
				Titulo: titulo,
				Link:   link,
			}
		}
	})
}

func main() {
	urls := []string{
		"https://news.ycombinator.com/news",
		"https://news.ycombinator.com/news?p=2",
		"https://news.ycombinator.com/news?p=3",
	}

	var wg sync.WaitGroup
	chNoticias := make(chan Noticia)
	fmt.Printf("A raspar %d páginas...\n", len(urls))

	for _, url := range urls {
		wg.Add(1)
		go webScrapper(url, &wg, chNoticias)
	}

	go func() {
		wg.Wait()
		close(chNoticias)
	}()

	var todasAsNoticias []Noticia
	for noticia := range chNoticias {
		todasAsNoticias = append(todasAsNoticias, noticia)
	}

	fmt.Printf("\n--- Total de %d Notícias Extraídas ---\n", len(todasAsNoticias))
	for _, noticia := range todasAsNoticias {
		fmt.Printf("Título: %s\nLink: %s\n\n", noticia.Titulo, noticia.Link)
	}
}