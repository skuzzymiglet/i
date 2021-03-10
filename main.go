package main

import (
	"log"
	"os/exec"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	in := ""

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	w, h := ui.TerminalDimensions()
	out := widgets.NewParagraph()
	out.SetRect(0, 0, w, h-10)
	inp := widgets.NewParagraph()
	inp.SetRect(0, h-10, w, h)
	ui.Render(inp, out)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			if e.ID == "<C-c>" {
				panic("fuck this god-awful language")
			}
			in += e.ID
			cmd := exec.Command("sh", "-c", in)
			var buf strings.Builder
			cmd.Stdout = &buf
			cmd.Stderr = &buf
			err := cmd.Run()
			if err != nil {
				out.Text = err.Error() + "\n" + buf.String()
				out.SetRect(0, 0, w, h)
				ui.Render(out)
			}
			out.Text = buf.String()
			out.SetRect(0, 0, w, h)
			ui.Render(out)
		}
	}
}
