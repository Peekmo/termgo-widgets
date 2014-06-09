package styles

// BorderStyle contains all elements which constitute a border
type BorderStyle struct {
	Left              rune
	LeftTopCorner     rune
	Top               rune
	RightTopCorner    rune
	Right             rune
	RightBottomCorner rune
	Bottom            rune
	LeftBottomCorner  rune
}

// Border types (for a container for example)
var (
	BorderSlimLinear  *BorderStyle = &BorderStyle{'│', '┌', '─', '┐', '│', '┘', '─', '└'}
	BorderDashed      *BorderStyle = &BorderStyle{'|', '+', '-', '+', '|', '+', '-', '+'}
	BorderTransparent *BorderStyle = &BorderStyle{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
)

// Alignement type (for a text or a widget)
type Alignment uint8

// Alignement possibilities (for text & widgets)
const (
	AlignCenter Alignment = iota
	AlignLeft
	AlignRight
	AlignTop
	AlignBottom
)
