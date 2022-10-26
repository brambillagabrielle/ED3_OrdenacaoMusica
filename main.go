package main

import (
	arq "OrdenacaoMusica/arquivosmusicas"
	ord "OrdenacaoMusica/estruturasordenacao"
	"fmt"
	"time"
)

func main() {

	inicio := time.Now()

	dados := arq.LerJson(`arquivosmusicas/arquivosjson/songs7JSONvector.txt`)
	fmt.Println("Leu o JSON")

	// dados = ord.QuickSort(dados, 0, len(dados)-1)
	dados = ord.CountingSort(dados)
	fmt.Println("Ordenou as notas")
	tempoOrd := time.Since(inicio)

	arq.EscreverPorArquivo(dados)
	fmt.Println("Escreveu no arquivo correto")
	tempoTotal := time.Since(inicio)

	fmt.Println("Tempo de execução para ordenação: " + tempoOrd.String())
	fmt.Println("Tempo de execução total: " + tempoTotal.String())

}
