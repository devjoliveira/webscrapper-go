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

	var wg sync.WaitGroup //espera que todas as goroutines terminem
	chNoticias := make(chan Noticia) // canal para enviar as notícias extraídas
	fmt.Printf("A raspar %d páginas...\n", len(urls))

	for _, url := range urls { // 1. Iteramos sobre as URLs que queremos raspar.
		wg.Add(1) // 2. Incrementamos o contador do WaitGroup para cada goroutine.
		go webScrapper(url, &wg, chNoticias) //3. Lançamos a goroutine com a palavra-chave 'go'.
	}

	// 3. Lançamos uma goroutine especial para esperar e fechar o canal.
	// Isto é crucial para que o nosso loop de recolha de resultados saiba quando parar.
	go func() {
		wg.Wait()      // Espera que todas as goroutines em wg.Add() chamem wg.Done().
		close(chNoticias) // Fecha o canal, sinalizando que não haverá mais resultados.
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