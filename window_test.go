package termgow

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/peekmo/termgo-widgets/colors"
	"testing"
)

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Creates a new window and checks default params

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestNewWindow(t *testing.T) {
	w := NewWindow()

	x, y := termbox.Size()
	if x != w.Width || y != w.Height {
		t.Errorf("Window size failed (%d, %d), exepected (%d, %d)", w.Width, w.Height, x, y)
	}

	if w.foreground != colors.Default || w.background != colors.Default {
		t.Error("Wrong foreground or background color")
	}

	if w.id != 1 {
		t.Error("Window id is not initialized")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Creates 2 windows to check if the increment on id is working

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestNewWindowIdIncrement(t *testing.T) {
	w1, w2 := NewWindow(), NewWindow()

	if w2.id == w1.id || (w2.id != 2 && w1.id != 2) {
		t.Error("Window id is not incremented")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks the method Window.SetSize

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestWindowSetSize(t *testing.T) {
	w := NewWindow()
	w.SetSize(5, 7)

	if w.Width != 5 || w.Height != 7 {
		t.Error("SetSize error (%d, %d), expected (5, 7)", w.Width, w.Height)
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks Window.SetForeground

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestSetForeground(t *testing.T) {
	w := NewWindow()
	w.SetForeground(colors.Red)

	if w.foreground != colors.Red {
		t.Error("SetForeground is bugged !")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks Window.SetBackground

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestSetBackground(t *testing.T) {
	w := NewWindow()
	w.SetBackground(colors.Red)

	if w.background != colors.Red {
		t.Error("SetForeground is bugged !")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks Window.draw() by filling the buffer and checks every cell

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestDraw(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Error("Unable to start termbox")
	}

	defer p.Close()

	w := NewWindow()
	w.SetForeground(colors.Cyan)
	w.SetBackground(colors.Blue)

	w.draw()

	for _, cell := range termbox.CellBuffer() {
		if cell.Fg != colors.Cyan {
			t.Error("Drawing error : FG color")
		}

		if cell.Bg != colors.Blue {
			t.Error("Drawing error : BG color")
		}
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks if the window is hide and show with its methods
  It supposes that when a window is add to a program, it is showed by
  default

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestHideAndShow(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Error("Unable to start termbox")
	}

	defer p.Close()

	w := NewWindow()
	p.AddWindow(w)

	before := len(p.showed)
	err2 := w.Hide()

	if err2 != nil {
		t.Errorf("An error has been received during TestHide : %s", fmt.Sprint(err2))
	}

	if after := len(p.showed); before == after {
		t.Errorf("The window seems to be showed (%d windows showed before, and %d after)", before, after)
	}

	err3 := w.Show()
	if err3 != nil {
		t.Errorf("An error has been received during TestSHow : %s", fmt.Sprint(err3))
	}

	if after := len(p.showed); before != after {
		t.Errorf("The window seems to be hided (%d windows showed before, and %d after)", before, after)
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Gets an error on hiding a window without a program instance

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestHideFailNoProg(t *testing.T) {
	w := NewWindow()
	err := w.Hide()

	if err == nil {
		t.Error("An error was expected, nil received")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Gets an error on hiding window which is not a part of the program

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestHideFailNotAdded(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Error("Unable to start termbox")
	}

	defer p.Close()

	w := NewWindow()

	err2 := w.Hide()
	if err2 == nil {
		t.Error("An error was expected, nil received")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Gets an error on hiding a window without a program instance

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestShowFailNoProg(t *testing.T) {
	w := NewWindow()
	err := w.Show()

	if err == nil {
		t.Error("An error was expected, nil received")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Gets an error on hiding window which is not a part of the program

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestShowFailNotAdded(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Error("Unable to start termbox")
	}

	defer p.Close()

	w := NewWindow()

	err2 := w.Show()
	if err2 == nil {
		t.Error("An error was expected, nil received")
	}
}
