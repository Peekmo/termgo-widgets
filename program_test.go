package termgow

import (
	"fmt"
	"testing"
)

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Test a successfull NewProgram

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestNewProgram(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Errorf("New program failed : %s", fmt.Sprint(err))
	}

	defer p.Close()

	if p.IsRunning == true {
		t.Error("New program failed : IsRunning is true")
	}

	if p.windows == nil {
		t.Error("New program failed : windows's map not init")
	}

	if p.showed == nil {
		t.Error("New program failed : showed windows map not init")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Tests to adds a window to the program
  Checks that it's showed by default too

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestAddWindow(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Error("Unable to start termbox")
	}

	defer p.Close()

	winNumberBefore := len(p.windows)
	showNumberBefore := len(p.showed)

	w := NewWindow()
	err2 := p.AddWindow(w)

	if err2 != nil {
		t.Errorf("An error occured %s", fmt.Sprint(err2))
	}

	if winNumberBefore+1 != len(p.windows) {
		t.Error("Window not successfully added to windows's map")
	}

	if showNumberBefore+1 != len(p.showed) {
		t.Error("Window not successfully added to showed windows map")
	}

	if w.parent == nil {
		t.Error("The window must have a parent when added to a program")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Add the same window to the program
  Expects an error

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestAddSameWindow(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Error("Unable to start termbox")
	}

	defer p.Close()

	w := NewWindow()
	err2 := p.AddWindow(w)
	if err2 != nil {
		t.Errorf("An error occured %s", fmt.Sprint(err2))
	}

	err3 := p.AddWindow(w)
	if err3 == nil {
		t.Error("An error was expected")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Add the different windows to the program

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestAddDifferentWindows(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Error("Unable to start termbox")
	}

	defer p.Close()

	winNumberBefore := len(p.windows)
	showNumberBefore := len(p.showed)

	w := NewWindow()
	err2 := p.AddWindow(w)

	if err2 != nil {
		t.Errorf("An error occured (1st win) %s", fmt.Sprint(err2))
	}

	w2 := NewWindow()
	err3 := p.AddWindow(w2)
	if err3 != nil {
		t.Errorf("An error occured (2nd win) %s", fmt.Sprint(err2))
	}

	if winNumberBefore+2 != len(p.windows) {
		t.Error("Window not successfully added to windows's map")
	}

	if showNumberBefore+2 != len(p.showed) {
		t.Error("Window not successfully added to showed windows map")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Test to check if the window is removed from the program

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestRemoveWindow(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Error("Unable to start termbox")
	}

	defer p.Close()

	w := NewWindow()
	err2 := p.AddWindow(w)

	if err2 != nil {
		t.Errorf("An error occured (1st win) %s", fmt.Sprint(err2))
	}

	winNumberBefore := len(p.windows)
	showNumberBefore := len(p.showed)

	err3 := p.RemoveWindow(w)
	if err3 != nil {
		t.Errorf("An error occured (1st win) %s", fmt.Sprint(err2))
	}

	if winNumberBefore-1 != len(p.windows) {
		t.Error("Window not successfully removed from windows's map")
	}

	if showNumberBefore-1 != len(p.showed) {
		t.Error("Window not successfully removed from showed windows map")
	}

	if w.parent != nil {
		t.Error("Window must not have a parent when it's removed from the program")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Test to check if an error is received when we are trying
  to remove a window which is not in the program

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestRemoveUnexistingWindow(t *testing.T) {
	p, err := NewProgram()
	if err != nil {
		t.Error("Unable to start termbox")
	}

	defer p.Close()

	w := NewWindow()

	err2 := p.RemoveWindow(w)
	if err2 == nil {
		t.Error("An error was expected")
	}
}
