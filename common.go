package termgow

import (
	"github.com/nsf/termbox-go"
	"github.com/peekmo/termgo-widgets/styles"
)

// Spacing type contains spacing values for each side of a rect
type Spacing struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

/*******************************************************************************
 *
 * ELEMENT
 *
 ******************************************************************************/

// Most basic type. Everything is an element
type Element struct {
	width  int
	height int

	foreground termbox.Attribute
	background termbox.Attribute
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetSize changes windows's width & height value

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (e *Element) SetSize(width, height int) {
	e.width, e.height = width, height
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
func (e *Element) SetBackground(color termbox.Attribute) {
	e.background = color
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetForeground sets foreground's color (characters)

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (e *Element) SetForeground(color termbox.Attribute) {
	e.foreground = color
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  GetBackground gets background's color

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (e *Element) GetBackground() termbox.Attribute {
	return e.background
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  GetForeground gets foreground (chars) color

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (e *Element) GetForeground() termbox.Attribute {
	return e.foreground
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Size returns width & height of the current element

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (e *Element) Size() (int, int) {
	return e.width, e.height
}

/*******************************************************************************
 *
 * CONTAINER
 *
 ******************************************************************************/

// A container is a visual basic element
type Container struct {
	Element

	margin  *Spacing
	padding *Spacing

	hasBorder   bool
	borderStyle *styles.BorderStyle

	parent Basic
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Draw the container

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (c *Container) draw() {
	var printed bool = false

	for x := 0; x < c.width; x++ {
		// Margin
		if x >= c.margin.Left && x < c.width-c.margin.Right {
			for y := 0; y < c.height; y++ {
				printed = false
				if y >= c.margin.Top && y < c.height-c.margin.Bottom {
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
			termbox.SetCell(x, y, c.borderStyle.LeftTopCorner, c.foreground, c.parent.GetBackground())
		case (c.height - c.margin.Bottom - 1):
			termbox.SetCell(x, y, c.borderStyle.LeftBottomCorner, c.foreground, c.parent.GetBackground())
		default:
			termbox.SetCell(x, y, c.borderStyle.Left, c.foreground, c.parent.GetBackground())
		}

		return true
	case (c.width - c.margin.Right - 1):
		switch y {
		case c.margin.Top:
			termbox.SetCell(x, y, c.borderStyle.RightTopCorner, c.foreground, c.parent.GetBackground())
		case (c.height - c.margin.Bottom - 1):
			termbox.SetCell(x, y, c.borderStyle.RightBottomCorner, c.foreground, c.parent.GetBackground())
		default:
			termbox.SetCell(x, y, c.borderStyle.Right, c.foreground, c.parent.GetBackground())
		}

		return true
	default:
		switch y {
		case c.margin.Top:
			termbox.SetCell(x, y, c.borderStyle.Top, c.foreground, c.parent.GetBackground())
			return true
		case (c.height - c.margin.Bottom - 1):
			termbox.SetCell(x, y, c.borderStyle.Bottom, c.foreground, c.parent.GetBackground())
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
		c.margin.Right+border+c.padding.Right >= c.width-x ||
		c.margin.Bottom+border+c.padding.Bottom >= c.height-y {

		termbox.SetCell(x, y, ' ', c.foreground, termbox.ColorDefault)
		return true
	}
	return false
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

/*******************************************************************************
 *
 * INTERFACES
 *
 ******************************************************************************/

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Basic interface contains all methods that every basic element
  must implements

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
type Basic interface {
	Size() (int, int)
	GetForeground() termbox.Attribute
	GetBackground() termbox.Attribute
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Layout interface contains methods that a layout must implement
  to be able to manage a set of widgets

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
type Layout interface {
	AddWidget(widget *Widget) error
	RemoveWidget(widget *Widget) error
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Widget interface manages widget's different states

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
type Widget interface {
	Hide() error
	Show() error
}
