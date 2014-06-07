/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  @author Axel Anceau - 2014
  Package termgow allows to creates console apps faster

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
package termgow

import (
	"errors"
	"github.com/nsf/termbox-go"
)

var program *Program

/*
  General informations about the program
*/
type Program struct {
	Element
	IsRunning bool            // If the program is currently running or not
	windows   map[int]*Window // All windows added to the program (showed or not) (key : window id)
	showed    map[int]*Window // Only showed windows (key : priority)
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  NewProgram creates a new Program (only one program is allowed)

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func NewProgram() (*Program, error) {
	if program != nil && program.IsRunning {
		return nil, errors.New("A program is already running")
	}

	err := termbox.Init()
	if err != nil {
		return nil, err
	}

	program = &Program{IsRunning: false, windows: make(map[int]*Window), showed: make(map[int]*Window)}
	program.width, program.height = termbox.Size()
	program.foreground, program.background = termbox.ColorDefault, termbox.ColorDefault

	return program, nil
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Run the given program (until ctrl + c has been pressed)

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (p *Program) Run() {
	p.IsRunning = true

	for p.IsRunning {
		for y := 0; y < p.height; y++ {
			for x := 0; x < p.width; x++ {
				termbox.SetCell(x, y, ' ', p.foreground, p.background)
			}
		}

		for _, win := range p.showed {
			win.draw()
		}

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

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  AddWindow adds a new window to the set of window to print
  An error is returned if the window already exists

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (p *Program) AddWindow(win *Window) error {
	if p.windows[win.id] != nil {
		return errors.New("This window is already added")
	}

	win.parent, win.base = p, p
	p.windows[win.id] = win

	err := win.Show()

	return err
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  RemoveWindow removes a window from the program
  An error is returned if the window does not exists

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (p *Program) RemoveWindow(win *Window) error {
	if p.windows[win.id] == nil {
		return errors.New("This window is not in the program")
	}

	if win.priority != -1 {
		win.Hide()
	}

	delete(p.windows, win.id)
	win.parent, win.base = nil, nil

	return nil
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  hide is a function to removes the given window from the showed ones

  Returns an error if the window is not a part of the program

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (p *Program) hide(win *Window) error {
	if p.windows[win.id] == nil {
		return errors.New("This window does not exists")
	}

	// BUG(p.hide) deleting win.priority from p.showed hides all other windows when .show() is called back
	// delete(p.showed, win.priority)

	for i := win.priority + 1; i <= len(p.showed); i++ {
		p.showed[i].priority -= 1
		p.showed[i-1] = p.showed[i]
	}

	delete(p.showed, len(p.showed))
	return nil
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  show adds the given window to the showed ones
  If already showed, its priority will be increase

  If the window is not in the window's list, an error will be returned

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (p *Program) show(win *Window) error {
	if p.windows[win.id] == nil {
		return errors.New("This window does not exists")
	}

	if win.priority != -1 {
		win.Hide()
	}

	win.priority = len(p.showed) + 1
	p.showed[win.priority] = win

	return nil
}
