package styles

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

var (
	BorderSlimLinear  *BorderStyle = &BorderStyle{'│', '┌', '─', '┐', '│', '┘', '─', '└'}
	BorderDashed      *BorderStyle = &BorderStyle{'|', '+', '-', '+', '|', '+', '-', '+'}
	BorderTransparent *BorderStyle = &BorderStyle{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
)
