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
	for x := 0; x < c.Width; x++ {
		// Margin
		if x >= c.margin.Left && x < c.Width-c.margin.Right {
			for y := 0; y < c.Height; y++ {
				if y >= c.margin.Top && y < c.Height-c.margin.Bottom {
					termbox.SetCell(x, y, ' ', c.foreground, c.background)
				}
			}
		}
	}

	termbox.Flush()
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
