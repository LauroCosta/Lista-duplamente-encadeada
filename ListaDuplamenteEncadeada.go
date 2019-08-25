package listaDuplamenteEncadeada

import "fmt"

type ListaEncadeada struct {
	Inicio  *no
	Fim     *no
	Tamanho int
}
type no struct {
	Valor    int
	Proximo  *no
	Anterior *no
}

func NovoItem(valor int) *no {
	return &no{valor, nil, nil}
}

func (l *ListaEncadeada) AddInicio(novoItem *no) {

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

func (l *ListaEncadeada) AddFim(novoItem *no) {

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

func (l *ListaEncadeada) AddPosicao(novoItem *no, posicao int) {

	if posicao >= 0 && posicao <= l.Tamanho {

		if posicao == 0 {
			l.AddInicio(novoItem)
		} else if posicao == l.Tamanho {
			l.AddFim(novoItem)
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
