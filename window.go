/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  @author Axel Anceau - 2014
  Package termgow allows to creates console apps faster

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
package termgow

import (
	"errors"
	"github.com/nsf/termbox-go"
	"github.com/peekmo/termgo-widgets/colors"
)

/**
 * Window is the base of a view with termgow
 */
type Window struct {
	Container

	id       int
	priority int

	parent *Program
}

var (
	windowId int = 1 // Current window id value
)

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  NewWindow returns a new Window with terminal's size

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func NewWindow() *Window {
	var win Window
	win.Width, win.Height = termbox.Size()

	win.margin = &Spacing{0, 0, 0, 0}
	win.padding = &Spacing{0, 0, 0, 0}

	win.id = windowId
	win.priority = -1

	win.foreground = colors.Default
	win.background = colors.Default

	windowId++
	return &win
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Hide the window from the program

  If the program has not been created, an error will be returned

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) Hide() error {
	if w.parent == nil {
		return errors.New("The window does not have any parent")
	}

	err := program.hide(w)
	if err != nil {
		return err
	}

	w.priority = -1
	return nil
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Show the window in the program.
  If the window is already showed, it will be displayed in foreground

  If the program has not been created or the window is not in its list,
  an error will be returned

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) Show() error {
	if w.parent == nil {
		return errors.New("The window does not have any parent")
	}

	return program.show(w)
}
