package arquivosmusicas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Dado struct {
	Arquivo int    `json:"arq"`
	Ordem   int    `json:"ordem"`
	Notas   string `json:"notas"`
}

func verifErro(erro error) {

	if erro != nil {
		fmt.Println(erro)
	}

}

func LerJson(arquivo string) []Dado {

	dados := make([]Dado, 0)

	arqJson, erro := os.Open(arquivo)
	verifErro(erro)

	defer arqJson.Close()

	bytesJson, erro := ioutil.ReadAll(arqJson)
	verifErro(erro)

	erro = json.Unmarshal(bytesJson, &dados)
	verifErro(erro)

	return dados

}

func EscreverPorArquivo(dados []Dado) {

	for i := range dados {

		nomeArquivo := "arquivosmusicas/arquivosordenados/arquivo" + fmt.Sprintf("%d", dados[i].Arquivo) + ".txt"

		arquivo, erro := os.OpenFile(nomeArquivo, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		verifErro(erro)

		_, erro = fmt.Fprintln(arquivo, dados[i].Notas)
		verifErro(erro)

	}

}
