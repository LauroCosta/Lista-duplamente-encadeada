package lde

import "fmt"

type ListaEncadeada struct {
	Inicio  *No
	Fim     *No
	Tamanho int
}
type No struct {
	Valor    int
	Chave    int
	Proximo  *No
	Anterior *No
}

func NovoItem(chave int, valor int) *No {
	return &No{valor, chave, nil, nil}
}

func (l *ListaEncadeada) AddInicio(valor int, chave int) {

	novoItem := NovoItem(valor, chave)

	if novoItem == nil {
		return
	}

	if l.Tamanho == 0 {
		l.Inicio = novoItem
		l.Fim = novoItem
		l.Tamanho++
	} else {
		novoItem.Proximo = l.Inicio
		l.Inicio.Anterior = novoItem
		l.Inicio = novoItem
		l.Tamanho++
	}
}
func (l *ListaEncadeada) RemoverInicio() {

	if l.Tamanho == 0 {
		return
	} else if l.Tamanho == 1 {
		l.Inicio = nil
		l.Fim = nil
		l.Tamanho--
	} else {
		l.Inicio = l.Inicio.Proximo
		l.Inicio.Anterior = nil
		l.Tamanho--
	}
}

func (l *ListaEncadeada) AddFim(valor int, chave int) {

	novoItem := NovoItem(valor, chave)

	if l.Tamanho == 0 {
		l.Fim = novoItem
		l.Inicio = novoItem
		l.Tamanho++
	} else {
		l.Fim.Proximo = novoItem
		novoItem.Anterior = l.Fim
		l.Fim = novoItem
		l.Tamanho++
	}

}

func (l *ListaEncadeada) RemoverFim() {
	if l.Tamanho == 0 {
		return
	} else if l.Tamanho == 1 {
		l.Inicio = nil
		l.Fim = nil
		l.Tamanho--
	} else {
		l.Fim = l.Fim.Anterior
		l.Fim.Proximo = nil
		l.Tamanho--
	}
}

func (l *ListaEncadeada) AddPosicao(valor int, chave int, posicao int) {

	novoItem := NovoItem(valor, chave)

	if posicao >= 0 && posicao <= l.Tamanho {

		if posicao == 0 {
			l.AddInicio(valor, chave)
		} else if posicao == l.Tamanho {
			l.AddFim(valor, chave)
		} else {
			i := 0
			aux := l.Inicio
			for i < posicao-1 {
				aux = aux.Proximo
				i++
			}
			novoItem.Proximo = aux.Proximo
			novoItem.Anterior = aux.Anterior
			aux.Proximo = novoItem
			novoItem.Proximo.Anterior = novoItem
			l.Tamanho++

		}

	} else {
		fmt.Println("posição inválida")
	}

}

func (l *ListaEncadeada) RemoverPosicao(posicao int) {

	if posicao >= 0 && posicao < l.Tamanho {

		if posicao == 0 {
			l.RemoverInicio()
		} else if posicao == l.Tamanho-1 {
			l.RemoverFim()
		} else {
			i := 0
			aux := l.Inicio
			for i < posicao {
				aux = aux.Proximo
				i++
			}
			aux.Anterior.Proximo = aux.Proximo
			aux.Proximo.Anterior = aux.Anterior
			l.Tamanho--

		}

	} else {
		fmt.Println("posição inválida")
	}
}

func (l *ListaEncadeada) Buscar(chave int) int {

	if l.Tamanho != 0 {
		if l.Inicio.Chave == chave {
			return l.Inicio.Valor
		} else if l.Fim.Chave == chave {
			return l.Fim.Valor
		} else {
			i := 0
			aux := l.Inicio.Proximo
			for i < l.Tamanho-1 {
				fmt.Println(aux.Chave, "  ", chave)
				if aux.Chave == chave {
					fmt.Print("entrou aqui")
					return aux.Valor
				} else {
					aux = aux.Proximo
				}
				i++
			}
			return -1
		}
	}
	return -1
}

func (l *ListaEncadeada) Remover(chave int) {

	if l.Tamanho != 0 {
		if l.Inicio.Chave == chave {
			l.RemoverInicio()
		} else if l.Fim.Chave == chave {
			l.RemoverFim()
		} else {
			i := 1
			aux := l.Inicio.Proximo
			for i < l.Tamanho-1 {
				if aux.Chave == chave {
					l.RemoverPosicao(i)
				} else {
					aux = aux.Proximo
				}
				i++
			}
		}
	}
}

func (l *ListaEncadeada) Listar() {

	if l.Tamanho == 0 {
		fmt.Println("lista vazia")
	} else {
		proximo := l.Inicio
		for proximo != nil {
			fmt.Print(proximo.Valor, " ")
			proximo = proximo.Proximo
		}
		fmt.Println()
	}
}
