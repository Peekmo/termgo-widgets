package colors

import (
	"github.com/nsf/termbox-go"
	"testing"
)

func TestColorMapping(t *testing.T) {
	if Default != termbox.ColorDefault {
		t.Error("Color mapping error : Default")
	}

	if Black != termbox.ColorBlack {
		t.Error("Color mapping error : Black")
	}

	if Blue != termbox.ColorBlue {
		t.Error("Color mapping error : Blue")
	}

	if Yellow != termbox.ColorYellow {
		t.Error("Color mapping error : Yellow")
	}

	if Magenta != termbox.ColorMagenta {
		t.Error("Color mapping error : Magenta")
	}

	if Green != termbox.ColorGreen {
		t.Error("Color mapping error : Green")
	}

	if Red != termbox.ColorRed {
		t.Error("Color mapping error : Red")
	}

	if Cyan != termbox.ColorCyan {
		t.Error("Color mapping error : Cyan")
	}

	if White != termbox.ColorWhite {
		t.Error("Color mapping error : White")
	}
}
