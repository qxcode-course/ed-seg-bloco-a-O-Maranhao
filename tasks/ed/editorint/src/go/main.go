package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	text   *List[*List[rune]] // a lista de linhas
	itLine *Node[*List[rune]] // iterador para a linha corrente
	itChar *Node[rune]        // iterador para o caracter do cursor
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.itChar = e.itLine.Value.Insert(e.itChar, r) // insere antes do elemento apontado pelo cursor
	e.itChar = e.itChar.Next()                    // move o cursor para próxima posição
}

func (e *Editor) KeyLeft() {
	if e.itChar != e.itLine.Value.Front() { // Se o cursor não está no início da linha
		e.itChar = e.itChar.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.itLine != e.text.Front() { // Se não está na primeira linha
		e.itLine = e.itLine.Prev()      // Atualiza iterador de linha para linha anterior
		e.itChar = e.itLine.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	e.text.Insert(e.itLine.Next(), NewList[rune]()) // cria uma nova linha e insere abaixo da linha corrente
	e.itLine = e.itLine.Next()                      // vai pra próxima linha
	e.itChar = e.itLine.Value.Front()               // move o cursor para o início da linha
}

func (e *Editor) KeyRight() { //esse aqui é só o contrário do keyleft né
	if e.itChar != e.itLine.Value.End() {
		e.itChar = e.itChar.Next()
		return
	}
	//agora no fim da linha
	if e.itLine.Next() != e.text.End() { //não está na última linha
		e.itLine = e.itLine.Next()
		e.itChar = e.itLine.Value.Front()
	}
}

func (e *Editor) KeyUp() {
	if e.itLine != e.text.Front() { //se não é a primeira linha
		e.itLine = e.itLine.Prev()        //aqui sobe uma linha
		e.itChar = e.itLine.Value.Front() //e esse volta pro início
	}
}

func (e *Editor) KeyDown() {
	if e.itLine.Next() != e.text.End() { //Se não está na última linha
		e.itLine = e.itLine.Next()        //desce uma linha
		e.itChar = e.itLine.Value.Front() //e volta pro início dnv q nem no keyUp
	}

}

func (e *Editor) KeyBackspace() {
	//tem dois casos eu acho, que é quando tá no começo da linha de baixo e vai pra cima, e quando é só pra apagar um caractere da linha mesmo
	if e.itChar != e.itLine.Value.Front() {
		e.itChar = e.itLine.Value.Erase(e.itChar.Prev())
		return
	}
	if e.itLine != e.text.Front() { //só faz algo se não tá na primeira linah, se o cursor está no início, junta com a anterior
		prevLine := e.itLine.Prev()     //guarda linha anterior
		e.itChar = prevLine.Value.End() // cursor vai pro fim dessa linha

		//Esse for medonho é pra mover todos os chars da linha atual pra anterior; nem é tão medonho vendo assim, mas é que é bem horizontal
		for it := e.itLine.Value.Front(); it != e.itLine.Value.End(); it = it.Next() {
			prevLine.Value.Insert(prevLine.Value.End(), it.Value)
		}

		e.text.Erase(e.itLine) //remove a linha atual
		e.itLine = prevLine    //atualiza
	}
}

func (e *Editor) KeyDelete() { //último graças
	//primeiro caso não tá no fim da linha ent apaga o caractere atual
	if e.itChar != e.itLine.Value.End() {
		e.itChar = e.itLine.Value.Erase(e.itChar)
		return
	}
	//segundo caso: está no fim da linha então junta com a seguinte
	if e.itLine.Next() != e.text.End() {
		nextLine := e.itLine.Next()
		for it := nextLine.Value.Front(); it != nextLine.Value.End(); it = it.Next() {
			e.itLine.Value.Insert(e.itLine.Value.End(), it.Value)
		}
		e.text.Erase(nextLine)
	}
}

func main() { // Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey: //tcell
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.text = NewList[*List[rune]]()
	e.text.PushBack(NewList[rune]())
	e.itLine = e.text.Front()
	e.itChar = e.itLine.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.text.Front(); line != e.text.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '↲'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.itChar {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
