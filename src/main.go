package main

import (
	"fmt"

	// "github.com/Pawilonek/nozbe-cli/tasks"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// RadioButtons implements a simple primitive for radio button selections.
type RadioButtons struct {
	*tview.Box
	options       []string
	currentOption int
}

// NewRadioButtons returns a new radio button primitive.
func NewRadioButtons(options []string) *RadioButtons {
	return &RadioButtons{
		Box:     tview.NewBox(),
		options: options,
	}
}

// Draw draws this primitive onto the screen.
func (r *RadioButtons) Draw(screen tcell.Screen) {
	r.Box.Draw(screen)
	x, y, width, height := r.GetInnerRect()

	for index, option := range r.options {
		if index >= height {
			break
		}

		radioButton := "[#BF616A:-:-]\u2610" // Unchecked.
		if index == r.currentOption {
			radioButton = "[#A3BE8C:#3B4252:b]\u2611" // Checked.
		}
		line := fmt.Sprintf(`%s[#ECEFF4] %s`, radioButton, option)
		tview.Print(screen, line, x, y+index, width, tview.AlignLeft, tcell.ColorYellow)
	}
}

// InputHandler returns the handler for this primitive.
func (r *RadioButtons) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return r.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch event.Key() {
		case tcell.KeyUp:
			r.currentOption--
			if r.currentOption < 0 {
				r.currentOption = 0
			}
		case tcell.KeyDown:
			r.currentOption++
			if r.currentOption >= len(r.options) {
				r.currentOption = len(r.options) - 1
			}
		case tcell.KeyCtrlSpace:
			fmt.Println(r.currentOption)
		}
	})
}

func main() {
	radioButtons := NewRadioButtons([]string{"Lions", "Elephants", "Giraffes", "afd", "asdf", "dddd", "asdfasdf"})
	radioButtons.
		// SetBackgroundColor(tcell.ColorDefault)
		SetBackgroundColor(tcell.NewHexColor(0x2E3440))

		//		SetRect(0, 0, 30, 5)

	if err := tview.NewApplication().SetRoot(radioButtons, false).Run(); err != nil {
		panic(err)
	}
}

// func main() {

// 	list := tasks.NewList()

// 	list.Add("test 1")
// 	list.Add("A small second task")
// 	list.Add("A third one!")
// 	list.Add("Why i'm counting this?")
// 	list.Add("Hello!")

// 	list.ToggleDone(1)
// 	list.ToggleDone(0)
// 	list.ToggleDone(1)

// 	list.ToggleDone(3)

// 	taskList := list.List()
// 	var doneCharacter string
// 	for i := 0; i < len(taskList); i++ {
// 		if taskList[i].Done {
// 			doneCharacter = "\033[1;32m☑\033[0m"
// 		} else {
// 			doneCharacter = "\033[1;31m☐\033[0m"
// 		}

// 		fmt.Printf("%s %s \n", doneCharacter, taskList[i].Name)
// 	}

// 	app := tview.NewApplication()

// 	// checkbox := tview.NewCheckbox().SetLabel("Hit Enter to check box: ")

// 	a := tview.NewList().
// 		AddItem("List item 1", "Some explanatory text", '', nil).
// 		AddItem("List item 2", "Some explanatory text", 'b', nil).
// 		AddItem("List item 3", "Some explanatory text", 'c', nil).
// 		AddItem("List item 4", "Some explanatory text", 'd', nil).
// 		AddItem("Quit", "Press to exit", 'q', func() {
// 			app.Stop()
// 		})

// 	if err := app.SetRoot(a, true).Run(); err != nil {
// 		panic(err)
// 	}
// }
