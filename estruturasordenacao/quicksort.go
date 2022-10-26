package estruturasordenacao

import (
	arq "OrdenacaoMusica/arquivosmusicas"
)

func particionar(dados []arq.Dado, inicio, fim int) ([]arq.Dado, int) {

	aux := dados[fim].Ordem
	pivo := inicio

	for j := inicio; j < fim; j++ {

		if dados[j].Ordem <= aux {
			dados[pivo], dados[j] = dados[j], dados[pivo]
			pivo++
		}

	}

	dados[pivo], dados[fim] = dados[fim], dados[pivo]
	return dados, pivo

}

func QuickSort(dados []arq.Dado, inicio, fim int) []arq.Dado {

	if inicio < fim {
		dados, pivo := particionar(dados, inicio, fim)
		QuickSort(dados, inicio, pivo-1)
		QuickSort(dados, pivo+1, fim)
	}

	return dados

}
