/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  @author Axel Anceau - 2014
  Package termgow allows to creates console apps faster

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
package termgow

import (
	"github.com/peekmo/termbox-go"
)

/**
 * General informations about the program
 */
type Program struct {
	IsRunning bool
	windows   []*Window
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  NewProgram creates a new Program (only one program is allowed)

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func NewProgram() (*Program, error) {
	err := termbox.Init()
	if err != nil {
		return nil, err
	}

	return &Program{IsRunning: false}, nil
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Run the given program (until ctrl + c has been pressed)

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (p *Program) Run() {
	p.IsRunning = true

	for p.IsRunning {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				p.Close()
			}
		}
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Close the program. Must be called when the program is not needed
  anymore

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (p *Program) Close() {
	p.IsRunning = false

	if termbox.IsInit {
		termbox.Close()
	}
}
