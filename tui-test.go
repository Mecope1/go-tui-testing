package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewList()
	p.Rows = []string{"Hello World!"}

	termWidth, termHeight := ui.TerminalDimensions()
	p.Rows = append(p.Rows, "width: "+strconv.Itoa(termWidth)+" height: "+strconv.Itoa(termHeight))
	p.SetRect(0, 0, 100, 30)

	// Replace this with the joined channel's number.
	num := 1

	p.Title = "Channel " + strconv.Itoa(num)
	//p.WrapText = true

	// These are included to move the cursor to the end of the list.
	// Otherwise the terminal will be 2 place behind the newest element, unless the user moves it themselves.
	p.ScrollDown()
	p.ScrollDown()

	q := widgets.NewParagraph()
	q.WrapText = true
	q.SetRect(0, 0, 100, 4)
	//p.SetRect(0, 0, termWidth, termHeight-50)

	grid := ui.NewGrid()

	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.6/2,
			ui.NewCol(1, p),
		),
		ui.NewRow(0.4/2,
			ui.NewCol(1, q),
		),
	)

	ui.Render(grid)

	for e := range ui.PollEvents() {

		//if e.Type == ui.ResizeEvent {
		//	ui.Render(p)
		//}

		if e.Type == ui.ResizeEvent {
			termWidth, termHeight := ui.TerminalDimensions()
			//p.SetRect(0, 0, termWidth-50, termHeight-50)
			//p.Rows = append(p.Rows, "\nwidth: " +  strconv.Itoa(termWidth) + " height: " + strconv.Itoa(termHeight))

			// These two are here for the same reason that they appeared above. This just pushes the cursor to the front.
			//p.ScrollDown()
			//p.ScrollDown()
			// This line redefines the size of the application, which is necessary when the terminal itself is resized.
			grid.SetRect(0, 0, termWidth, termHeight)
		} else if e.Type == ui.KeyboardEvent || e.Type == ui.MouseEvent {
			switch e.ID {
			case "<C-c>":
				fallthrough
			case "<Escape>":
				ui.Close()
				os.Exit(0)
			case "<PageUp>":
				p.ScrollPageUp()
			case "<PageDown>":
				p.ScrollPageDown()
			case "<Up>":
				fallthrough
			case "<MouseWheelUp>":
				p.ScrollUp()
			case "<Down>":
				fallthrough
			case "<MouseWheelDown>":
				p.ScrollDown()
			case "<Enter>":
				p.Rows = append(p.Rows, q.Text)
				q.Text = ""
				p.ScrollDown()
			case "<Backspace>":
				if len(q.Text) > 0 {
					chars := strings.Split(q.Text, "")

					q.Text = strings.Join(chars[0:len(chars)-1], "")
				}
			case "<Space>":
				q.Text += " "
			case "<Insert>":
			case "<Delete>":
			case "<Home>":
			case "<End>":
			case "<F1>":
			case "<F2>":
			case "<F3>":
			case "<F4>":
			case "<F5>":
			case "<F6>":
			case "<F7>":
			case "<F8>":
			case "<F9>":
			case "<F10>":
			case "<F11>":
			case "<F12>":
			case "<Left>":
			case "<Right>":
			case "<MouseRelease>":
			case "<MouseLeft>":
			case "<MouseMiddle>":
			case "<MouseRight>":
			default:
				q.Text += e.ID

			}

		}

		ui.Clear()
		ui.Render(grid)
	}
}
