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

const (
	BorderSlimLinear *BorderStyle = &BorderStyle{'│', '┌', '─', '┐', '│', '┘', '─', '└'}
	BorderDashed     *BorderStyle = &BorderStyle{'|', '+', '-', '+', '|', '+', '-', '+'}
)
