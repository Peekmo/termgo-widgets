package termgow

import (
	"github.com/nsf/termbox-go"
	"github.com/peekmo/termgo-widgets/styles"
)

type Spacing struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

type Container struct {
	Width  int
	Height int

	margin  *Spacing
	padding *Spacing

	hasBorder   bool
	borderStyle *styles.BorderStyle

	foreground termbox.Attribute
	background termbox.Attribute
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Draw the container

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) draw() {
	var printed bool = false

	for x := 0; x < c.Width; x++ {
		// Margin
		if x >= c.margin.Left && x < c.Width-c.margin.Right {
			for y := 0; y < c.Height; y++ {
				printed = false
				if y >= c.margin.Top && y < c.Height-c.margin.Bottom {
					if c.hasBorder {
						printed = c.printBorder(x, y)
					}

					if !printed && (c.padding.Top != 0 || c.padding.Bottom != 0 || c.padding.Left != 0 || c.padding.Right != 0) {
						printed = c.printPadding(x, y)
					}

					if !printed {
						termbox.SetCell(x, y, ' ', c.foreground, c.background)
					}
				}
			}
		}
	}

	termbox.Flush()
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  printBorder is an internal function to print container's border
  if needed.
  It returns true/false if it printed something or not

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) printBorder(x, y int) bool {
	switch x {
	case c.margin.Left:
		switch y {
		case c.margin.Top:
			termbox.SetCell(x, y, c.borderStyle.LeftTopCorner, c.foreground, termbox.ColorDefault)
		case (c.Height - c.margin.Bottom - 1):
			termbox.SetCell(x, y, c.borderStyle.LeftBottomCorner, c.foreground, termbox.ColorDefault)
		default:
			termbox.SetCell(x, y, c.borderStyle.Left, c.foreground, termbox.ColorDefault)
		}

		return true
	case (c.Width - c.margin.Right - 1):
		switch y {
		case c.margin.Top:
			termbox.SetCell(x, y, c.borderStyle.RightTopCorner, c.foreground, termbox.ColorDefault)
		case (c.Height - c.margin.Bottom - 1):
			termbox.SetCell(x, y, c.borderStyle.RightBottomCorner, c.foreground, termbox.ColorDefault)
		default:
			termbox.SetCell(x, y, c.borderStyle.Right, c.foreground, termbox.ColorDefault)
		}

		return true
	default:
		switch y {
		case c.margin.Top:
			termbox.SetCell(x, y, c.borderStyle.Top, c.foreground, termbox.ColorDefault)
			return true
		case (c.Height - c.margin.Bottom - 1):
			termbox.SetCell(x, y, c.borderStyle.Bottom, c.foreground, termbox.ColorDefault)
			return true
		}
	}

	return false
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  printPadding prints container's padding (if any)

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) printPadding(x, y int) bool {
	var border int = 0
	if c.hasBorder {
		border = 1
	}

	if c.margin.Left+border+c.padding.Left > x ||
		c.margin.Top+border+c.padding.Top > y ||
		c.margin.Right+border+c.padding.Right >= c.Width-x ||
		c.margin.Bottom+border+c.padding.Bottom >= c.Height-y {

		termbox.SetCell(x, y, ' ', c.foreground, termbox.ColorDefault)
		return true
	}
	return false
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetSize changes windows's width & height value

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) SetSize(width, height int) {
	c.Width, c.Height = width, height
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetMargin sets the margin (in terms of chars) for every bounds
  of the window

  Note : If an element has a specified x/y pos, the margin will not
  be considered

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) SetMargin(s *Spacing) {
	c.margin = s
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetPadding sets the padding (in terms of chars) for every bounds
  of the window

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) SetPadding(s *Spacing) {
	c.padding = s
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetBackground changes background's color

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) SetBackground(color termbox.Attribute) {
	c.background = color
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetForeground sets foreground's color (characters)

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) SetForeground(color termbox.Attribute) {
	c.foreground = color
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetBorder sets the border style of the container

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) SetBorder(style *styles.BorderStyle) {
	c.hasBorder = true
	c.borderStyle = style
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  RemoveBorder removes the border from the container (if any)

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) RemoveBorder() {
	c.hasBorder = false
	c.borderStyle = nil
}

type Element interface {
	draw()
	SetSize(width, height int)
	SetSpacing(s *Spacing)
}

type Layout interface {
	AddWidget(widget *Widget) error
	RemoveWidget(widget *Widget) error
}

type Widget interface {
	SetBackground(color termbox.Attribute)
	SetForeground(color termbox.Attribute)
	Hide() error
	Show() error
}
