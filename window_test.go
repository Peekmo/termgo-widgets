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
	windowId = 1
	w := NewWindow()

	x, y := termbox.Size()
	if x != w.width || y != w.height {
		t.Errorf("Window size failed (%d, %d), exepected (%d, %d)", w.width, w.height, x, y)
	}

	if w.foreground != colors.Default || w.background != colors.Default {
		t.Error("Wrong foreground or background color")
	}

	if w.id != 1 {
		t.Errorf("Window id is not initialized (id %d)", w.id)
	}

	if w.margin == nil {
		t.Error("Margin is not init")
	}

	if w.padding == nil {
		t.Error("Padding is not init")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Creates 2 windows to check if the increment on id is working

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestNewWindowIdIncrement(t *testing.T) {
	windowId = 1
	w1, w2 := NewWindow(), NewWindow()

	if w2.id == w1.id || (w2.id != 2 && w1.id != 2) {
		t.Error("Window id is not incremented")
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
	if before == -1 {
		t.Error("Test is not valid anymore because we are expecting the window to be showed when addWindow is called")
	}

	err2 := w.Hide()
	if err2 != nil {
		t.Errorf("An error has been received during TestHide : %s", fmt.Sprint(err2))
	}

	if w.priority != -1 {
		t.Errorf("The window seems to be showed (priority %d)", w.priority)
	}

	err3 := w.Show()
	if err3 != nil {
		t.Errorf("An error has been received during TestSHow : %s", fmt.Sprint(err3))
	}

	if w.priority != before {
		t.Errorf("Show test failed (priority %d, expected %d)", w.priority, before)
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Gets an error on hiding a window without a program instance

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestHide_FailNoProg(t *testing.T) {
	w := NewWindow()
	err := w.Hide()

	if err == nil {
		t.Error("An error was expected, nil received")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Gets an error on hiding window which is not a part of the program

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestHide_FailNotAdded(t *testing.T) {
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
func TestShow_FailNoProg(t *testing.T) {
	w := NewWindow()
	err := w.Show()

	if err == nil {
		t.Error("An error was expected, nil received")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Gets an error on hiding window which is not a part of the program

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestShow_FailNotAdded(t *testing.T) {
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
