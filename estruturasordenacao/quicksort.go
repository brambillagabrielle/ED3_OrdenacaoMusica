package estruturasordenacao

import (
	arq "OrdenacaoMusica/arquivosmusicas"
)

func particionar(dados []arq.Dado, inicio, fim int) ([]arq.Dado, int) {

	pivo := dados[fim].Ordem
	i := inicio

	for j := inicio; j < fim; j++ {

		if dados[j].Ordem <= pivo {
			dados[i], dados[j] = dados[j], dados[i]
			i++
		}

	}

	dados[i], dados[fim] = dados[fim], dados[i]
	return dados, i

}

func QuickSort(dados []arq.Dado, inicio, fim int) []arq.Dado {

	if inicio < fim {
		dados, pivo := particionar(dados, inicio, fim)
		QuickSort(dados, inicio, pivo-1)
		QuickSort(dados, pivo+1, fim)
	}

	return dados

}
