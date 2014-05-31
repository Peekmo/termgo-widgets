/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  @author Axel Anceau - 2014
  Package termgow allows to creates console apps faster

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
package termgow

import (
	"github.com/peekmo/termbox-go"
)

/**
 * Window is the base of a view with termgow
 */
type Window struct {
	Width  int
	Height int
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  NewWindow returns a new Window with terminal's size

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func NewWindow() *Window {
	var win Window
	win.Width, win.Height = termbox.Size()

	return &win
}
