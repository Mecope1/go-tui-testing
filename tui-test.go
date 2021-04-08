package main


import (
	"log"
	"strconv"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewList()
	p.Rows = []string{"Hello World!"}

	termWidth, termHeight := ui.TerminalDimensions()
	p.Rows = append(p.Rows, "width: " +  strconv.Itoa(termWidth) + " height: " + strconv.Itoa(termHeight))
	p.SetRect(0, 0, 100, 30)
	p.ScrollDown()
	p.ScrollDown()

	q := widgets.NewParagraph()
	q.Text = "Typing goes here!"
	q.SetRect(0, 0, 100, 4)
	//p.SetRect(0, 0, termWidth, termHeight-50)

	grid := ui.NewGrid()

	grid.SetRect(0,0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.6/2,
			ui.NewCol(1,p),
		),
		ui.NewRow(0.4/2,
			ui.NewCol(1,q),
		),
	)


	ui.Render(grid)

	for e := range ui.PollEvents() {

		//if e.Type == ui.ResizeEvent {
		//	ui.Render(p)
		//}

		if e.ID == "<C-c>" || e.ID =="<Escape>" {
			break
		} else if e.Type == ui.ResizeEvent {
			termWidth, termHeight := ui.TerminalDimensions()
			//p.SetRect(0, 0, termWidth-50, termHeight-50)
			p.Rows = append(p.Rows, "\nwidth: " +  strconv.Itoa(termWidth) + " height: " + strconv.Itoa(termHeight))
			p.ScrollDown()
			p.ScrollDown()
			grid.SetRect(0,0, termWidth, termHeight)

		} else if e.ID == "<PageUp>"{
			p.ScrollPageUp()
		}else if e.ID == "<PageDown>" {
			p.ScrollPageDown()
		}else if e.ID == "<Up>" || e.ID == "<MouseWheelUp>"{
			p.ScrollUp()
		} else if e.ID == "<Down>" || e.ID == "<MouseWheelDown>" {
			p.ScrollDown()
		} else if e.ID == "<Enter>" {
			p.Rows = append(p.Rows, q.Text)
			q.Text = ""
			p.ScrollDown()
		} else if e.ID == "<Backspace>" {

			if len(q.Text) > 0 {
				chars := strings.Split(q.Text, "")

				q.Text = strings.Join(chars[0:len(chars)-1], "")
			}
		} else if e.ID == "<Space>" {
			q.Text += " "
		} else if e.Type == ui.KeyboardEvent {
				q.Text += e.ID
		}

		ui.Clear()
		ui.Render(grid)
	}

	//q.Text =
}