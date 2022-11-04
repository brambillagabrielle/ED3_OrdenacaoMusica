package estruturasordenacao

import (
	arq "OrdenacaoMusica/arquivosmusicas"
)

func particionar(dados []arq.Dado, inicio, fim int) int {

	aux := dados[fim].Ordem
	pivo := inicio

	for j := inicio; j < fim; j++ {

		if dados[j].Ordem <= aux {
			dados[pivo], dados[j] = dados[j], dados[pivo]
			pivo++
		}

	}

	dados[pivo], dados[fim] = dados[fim], dados[pivo]
	return pivo

}

func QuickSort(dados []arq.Dado, inicio, fim int) {

	if inicio < fim {
		pivo := particionar(dados, inicio, fim)
		QuickSort(dados, inicio, pivo-1)
		QuickSort(dados, pivo+1, fim)
	}

}
