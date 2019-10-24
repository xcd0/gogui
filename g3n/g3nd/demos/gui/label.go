package gui

import (
	"strings"
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.label"] = &GuiLabel{}
}

type GuiLabel struct{}

// Start is called once at the start of the demo.
func (t *GuiLabel) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	l1 := gui.NewLabel("label1")
	l1.SetPosition(10, 10)
	a.DemoPanel().Add(l1)

	l2 := gui.NewLabel("label2")
	l2.SetPosition(60, 10)
	l2.SetBorders(1, 1, 1, 1)
	l2.SetBordersColor(math32.NewColor("black"))
	l2.SetPaddings(2, 2, 2, 2)
	a.DemoPanel().Add(l2)

	l3 := gui.NewLabel("label3")
	l3.SetPosition(120, 10)
	l3.SetBgColor(math32.NewColor("green"))
	l3.SetBorders(1, 1, 1, 1)
	l3.SetPaddings(4, 6, 4, 6)
	a.DemoPanel().Add(l3)

	l4 := gui.NewLabel("label4")
	l4.SetPosition(200, 10)
	l4.SetBgColor(math32.NewColor("blue"))
	l4.SetColor(math32.NewColor("white"))
	l4.SetBorders(1, 1, 1, 1)
	l4.SetPaddings(4, 20, 4, 20)
	l4.SetFontSize(20)
	a.DemoPanel().Add(l4)

	l5 := gui.NewLabel("label5")
	l5.SetPosition(320, 10)
	l5.SetFontSize(28)
	l5.SetColor(math32.NewColor("red"))
	l5.SetBorders(1, 1, 1, 1)
	l5.SetBordersColor(math32.NewColor("white"))
	l5.SetPaddings(4, 20, 4, 20)
	l5.SetSize(100, 100)
	a.DemoPanel().Add(l5)

	l6 := gui.NewLabel("label6")
	l6.SetPosition(450, 10)
	l6.SetColor(math32.NewColor("red"))
	l6.SetBorders(1, 1, 1, 1)
	l6.SetBordersColor(math32.NewColor("white"))
	l6.SetPaddings(4, 20, 4, 20)
	l6.SetSize(100, 100)
	l6.SetFontSize(28)
	a.DemoPanel().Add(l6)

	lines := []string{
		"We are merely picking up pebbles on the beach",
		"while the great ocean of truth",
		"lays completely undiscovered before us.",
	}
	l7 := gui.NewLabel(strings.Join(lines, "\n"))
	l7.SetPosition(10, 120)
	l7.SetBordersColor(math32.NewColor("red"))
	l7.SetBgColor(math32.NewColor("green"))
	l7.SetColor(math32.NewColor("blue"))
	l7.SetBorders(10, 4, 10, 4)
	l7.SetPaddings(4, 20, 4, 20)
	l7.SetFontSize(22)
	a.DemoPanel().Add(l7)
}

// Update is called every frame.
func (t *GuiLabel) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiLabel) Cleanup(a *app.App) {}
