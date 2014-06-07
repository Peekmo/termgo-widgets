package termgow

import (
	"github.com/nsf/termbox-go"
	"github.com/peekmo/termgo-widgets/colors"
	"testing"
	// "github.com/peekmo/termgo-widgets/styles"
)

/*******************************************************************************
 *
 * ELEMENT
 *
 ******************************************************************************/

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks the method SetSize

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestElement_SetSize(t *testing.T) {
	w := &Element{}
	w.SetSize(5, 7)

	if w.width != 5 || w.height != 7 {
		t.Error("SetSize error (%d, %d), expected (5, 7)", w.width, w.height)
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks SetForeground

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestElement_SetForeground(t *testing.T) {
	w := &Element{}
	w.SetForeground(colors.Red)

	if w.foreground != colors.Red {
		t.Error("SetForeground is bugged !")
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks SetBackground

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestElement_SetBackground(t *testing.T) {
	w := &Element{}
	w.SetBackground(colors.Red)

	if w.background != colors.Red {
		t.Error("SetForeground is bugged !")
	}
}

/*******************************************************************************
 *
 * CONTAINER
 *
 ******************************************************************************/

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks the method SetMargin

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestContainer_SetMargin(t *testing.T) {
	w := &Container{}
	w.SetMargin(&Spacing{Left: 1, Top: 2, Right: 3, Bottom: 4})

	if w.margin.Left != 1 || w.margin.Top != 2 || w.margin.Right != 3 || w.margin.Bottom != 4 {
		t.Errorf("SetSize error (%d, %d, %d, %d), expected (1,2,3,4)", w.margin.Left,
			w.margin.Top, w.margin.Right, w.margin.Bottom)
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks the method SetPadding

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestContainer_SetPadding(t *testing.T) {
	w := &Container{}
	w.SetPadding(&Spacing{Left: 1, Top: 2, Right: 3, Bottom: 4})

	if w.padding.Left != 1 || w.padding.Top != 2 || w.padding.Right != 3 || w.padding.Bottom != 4 {
		t.Errorf("SetSize error (%d, %d, %d, %d), expected (1,2,3,4)", w.padding.Left,
			w.padding.Top, w.padding.Right, w.padding.Bottom)
	}
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Checks draw() by filling the buffer and checks every cell

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func TestContainer_Draw(t *testing.T) {
	w := &Container{}
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
