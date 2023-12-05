package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pessoa struct {
	Nome      string
	Idade     int
	Pontuacao int
}

type PorNome []Pessoa
type PorIdade []Pessoa

func (a PorNome) Len() int           { return len(a) }
func (a PorNome) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorNome) Less(i, j int) bool { return a[i].Nome < a[j].Nome }

func (a PorIdade) Len() int           { return len(a) }
func (a PorIdade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorIdade) Less(i, j int) bool { return a[i].Idade < a[j].Idade }

func main() {
	arquivoOrigem := os.Args[1]
	arquivoDestino := os.Args[2]

	// Leitura do arquivo de entrada
	arquivo, err := os.Open(arquivoOrigem)
	if err != nil {
		log.Fatal(err)
	}
	defer arquivo.Close()

	leitorCSV := csv.NewReader(arquivo)

	_, err = leitorCSV.Read()
	if err != nil {
		log.Fatal(err)
	}

	linhas, err := leitorCSV.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Processamento dos dados
	pessoas := make([]Pessoa, 0, len(linhas))
	for _, linha := range linhas {
		idade, _ := strconv.Atoi(linha[1])
		pontuacao, _ := strconv.Atoi(linha[2])
		p := Pessoa{Nome: linha[0], Idade: idade, Pontuacao: pontuacao}
		pessoas = append(pessoas, p)
	}

	extensao := strings.LastIndex(arquivoDestino, ".")

	if extensao < 0 {
		arquivoDestino = arquivoDestino + ".csv"
		extensao = strings.LastIndex(arquivoDestino, ".")
	}

	// Ordenação por nome
	sort.Sort(PorNome(pessoas))
	salvarOrdenado(arquivoDestino[:extensao] + "_ordenado_por_nome" + arquivoDestino[extensao:], pessoas)

	// Ordenação por idade
	sort.Sort(PorIdade(pessoas))
	salvarOrdenado(arquivoDestino[:extensao] + "_ordenado_por_idade" + arquivoDestino[extensao:], pessoas)
}

func salvarOrdenado(nomeArquivo string, pessoas []Pessoa) {
	arquivoSaida, err := os.Create(nomeArquivo)
	if err != nil {
		log.Fatal(err)
	}
	defer arquivoSaida.Close()

	escritorCSV := csv.NewWriter(arquivoSaida)
	defer escritorCSV.Flush()

	for _, pessoa := range pessoas {
		escritorCSV.Write([]string{pessoa.Nome, strconv.Itoa(pessoa.Idade), strconv.Itoa(pessoa.Pontuacao)})
	}
}