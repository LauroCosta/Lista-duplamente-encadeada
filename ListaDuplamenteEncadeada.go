package listaDuplamenteEncadeada

import "fmt"

type listaEncadeada struct {
	inicio  *no
	fim     *no
	tamanho int
}
type no struct {
	valor    int
	proximo  *no
	anterior *no
}

func novoItem(valor int) *no {
	return &no{valor, nil, nil}
}

func (l *listaEncadeada) addInicio(novoItem *no) {

	if novoItem == nil {
		return
	}

	if l.tamanho == 0 {
		l.inicio = novoItem
		l.fim = novoItem
		l.tamanho++
	} else {
		novoItem.proximo = l.inicio
		l.inicio.anterior = novoItem
		l.inicio = novoItem
		l.tamanho++
	}
}
func (l *listaEncadeada) removerInicio() {

	if l.tamanho == 0 {
		return
	} else if l.tamanho == 1 {
		l.inicio = nil
		l.fim = nil
		l.tamanho--
	} else {
		l.inicio = l.inicio.proximo
		l.inicio.anterior = nil
		l.tamanho--
	}
}

func (l *listaEncadeada) addFim(novoItem *no) {

	if l.tamanho == 0 {
		l.fim = novoItem
		l.inicio = novoItem
		l.tamanho++
	} else {
		l.fim.proximo = novoItem
		novoItem.anterior = l.fim
		l.fim = novoItem
		l.tamanho++
	}

}

func (l *listaEncadeada) removerFim() {
	if l.tamanho == 0 {
		return
	} else if l.tamanho == 1 {
		l.inicio = nil
		l.fim = nil
		l.tamanho--
	} else {
		l.fim = l.fim.anterior
		l.fim.proximo = nil
		l.tamanho--
	}
}

func (l *listaEncadeada) addPosicao(novoItem *no, posicao int) {

	if posicao >= 0 && posicao <= l.tamanho {

		if posicao == 0 {
			l.addInicio(novoItem)
		} else if posicao == l.tamanho {
			l.addFim(novoItem)
		} else {
			i := 0
			aux := l.inicio
			for i < posicao-1 {
				aux = aux.proximo
				i++
			}
			novoItem.proximo = aux.proximo
			novoItem.anterior = aux.anterior
			aux.proximo = novoItem
			novoItem.proximo.anterior = novoItem
			l.tamanho++

		}

	} else {
		fmt.Println("posição inválida")
	}

}

func (l *listaEncadeada) removerPosicao(posicao int) {

	if posicao >= 0 && posicao < l.tamanho {

		if posicao == 0 {
			l.removerInicio()
		} else if posicao == l.tamanho-1 {
			l.removerFim()
		} else {
			i := 0
			aux := l.inicio
			for i < posicao {
				aux = aux.proximo
				i++
			}
			aux.anterior.proximo = aux.proximo
			aux.proximo.anterior = aux.anterior
			l.tamanho--

		}

	} else {
		fmt.Println("posição inválida")
	}
}

func (l *listaEncadeada) listar() {

	if l.tamanho == 0 {
		fmt.Println("lista vazia")
	} else {
		proximo := l.inicio
		for proximo != nil {
			fmt.Print(proximo.valor, " ")
			proximo = proximo.proximo
		}
		fmt.Println()
	}
}
