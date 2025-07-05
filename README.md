# Web Scraper Concorrente em Go

![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## 📖 Sobre o Projeto

Este projeto é um web scraper desenvolvido em Go como um exercício prático de programação. 

O objetivo é extrair de forma concorrente os títulos e links das notícias das primeiras páginas do **Hacker News** (`news.ycombinator.com`).

O projeto serve como um estudo de caso para aplicar conceitos fundamentais e avançados de Go, demonstrando o poder da linguagem para tarefas de rede e processamento de dados em paralelo.

## ✨ Funcionalidades Principais

* **Raspagem de Múltiplas Páginas**: Extrai dados das 3 primeiras páginas do Hacker News.
* **Processamento Concorrente**: Utiliza **goroutines** para fazer os pedidos HTTP em paralelo, otimizando drasticamente o tempo de execução em comparação com uma abordagem sequencial.
* **Sincronização Segura**: Emprega **`sync.WaitGroup`** para garantir que o programa espere pela conclusão de todas as tarefas e **channels** para a comunicação segura dos resultados entre as goroutines, evitando *race conditions*.
* **Parsing de HTML**: Usa a popular biblioteca `goquery` para analisar o documento HTML e extrair dados de forma eficiente usando seletores CSS.
* **Normalização de Dados**: Implementa uma lógica de limpeza para converter URLs relativas em absolutas, garantindo a usabilidade e consistência dos links extraídos.

## 🛠️ Tecnologias Utilizadas

* [**Go**](https://go.dev/): A linguagem de programação principal.
* [**goquery**](https://github.com/PuerkitoBio/goquery): Biblioteca utilizada para o parsing de HTML e manipulação do DOM.

## 🚀 Como Executar

Para executar este projeto localmente, siga os passos abaixo.

### Pré-requisitos

É necessário ter o **Go (versão 1.18 ou superior)** instalado e configurado no seu sistema.

### Passos para a Execução

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/SEU-USUARIO/SEU-REPOSITORIO.git](https://github.com/SEU-USUARIO/SEU-REPOSITORIO.git)
    ```
    *(Substitua `SEU-USUARIO/SEU-REPOSITORIO` pelo seu nome de utilizador e nome do repositório no GitHub)*

2.  **Navegue até a pasta do projeto:**
    ```bash
    cd SEU-REPOSITORIO
    ```

3.  **Instale as dependências:**
    Este comando irá verificar o ficheiro `go.mod` e descarregar as dependências necessárias (neste caso, a `goquery`).
    ```bash
    go mod tidy
    ```

4.  **Execute o programa:**
    ```bash
    go run main.go
    ```

Após a execução, o programa irá imprimir no terminal a lista de todos os títulos e links extraídos das páginas configuradas.

## 🏗️ Estrutura do Código

O código está organizado da seguinte forma para maximizar a clareza e a eficiência:

* **`main()`**: A função principal que atua como a "orquestradora". Ela configura as URLs alvo, inicializa o `WaitGroup` e o `channel`, e lança as goroutines. No final, é responsável por recolher e exibir os resultados.
* **`webScrapper(url string, wg *sync.WaitGroup, ch chan<- Noticia)`**: Cada instância desta função corre numa goroutine separada e é responsável por raspar uma única página, extrair os dados e enviar os resultados de volta para a `main` através do channel.
* **`Noticia struct`**: Uma estrutura de dados simples para armazenar de forma organizada as informações extraídas (Título e Link).

## 📄 Licença

Este projeto está licenciado sob a Licença MIT. Veja o ficheiro `LICENSE` para mais detalhes.
