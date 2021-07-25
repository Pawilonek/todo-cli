package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

    "github.com/Pawilonek/nozbe-cli/tasks"
    "github.com/Pawilonek/nozbe-cli/storage"
)

type TaskBox struct {
	*tview.Table
	List *tasks.List
}

var app *tview.Application
var pages *tview.Pages
var inputField *tview.InputField

func textObject(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetText(text).
		SetWordWrap(true)
}

func (task *TaskBox) Draw(screen tcell.Screen) {

	for i, t := range task.List.List() {

		var checkIcon string
		var textColor tcell.Color
		if t.Done {
			// Checked
			checkIcon = " [#A3BE8C]\u2611 "
			textColor = tcell.NewHexColor(0x4C566A)
		} else {
			// Unchecked
			checkIcon = " [#BF616A]\u2610 "
			textColor = tcell.NewHexColor(0xD8DEE9)
		}

		tcCheck := tview.NewTableCell(checkIcon)
		task.Table.SetCell(i, 0, tcCheck)

		tcName := tview.NewTableCell(t.Name).SetTextColor(textColor).SetExpansion(1)
		task.Table.SetCell(i, 1, tcName)
	}

	task.Table.SetSelectable(true, false)
	task.Table.SetSelectedStyle(tcell.ColorDefault, tcell.NewHexColor(0x3B4252), tcell.AttrNone)
	task.Table.
		SetFixed(1, 2).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				app.Stop()
			}
			if key == tcell.KeyEnter {
				task.Table.SetSelectable(true, true)
			}
		}).
		SetSelectedFunc(func(row int, column int) {
			task.List.ToggleDone(row)
            // Moze selected task to a next one
            // TODO: Make sure we do not overflow cells
            task.Table.Select(row + 1, column)
		}).
		SetInputCapture(func(capture *tcell.EventKey) *tcell.EventKey {
            // Open adding new task when pressing "c"
            if capture.Rune() == 'c' {
                pages.ShowPage("modal")
                app.SetFocus(inputField)

				return nil
		    }

			return capture
		})

	task.Table.Draw(screen)
}


func Main() {
	app = tview.NewApplication()

    pages = tview.NewPages()

    storage := storage.NewDisk("storage.json")
    list, err := storage.LoadTasks()
    if err != nil {
        panic(err)
    }

    defer func() {
        storage.SaveTasks(list)
    }()

	tbox := TaskBox{
        Table: tview.NewTable(),
		List:  &list,
	}

	tbox.
		SetBackgroundColor(tcell.NewHexColor(0x2E3440))

	modal := func(p tview.Primitive, width, height int) tview.Primitive {

        frame := tview.NewFrame(p).
            SetBorders(0, 1, 0, 0, 1, 1).
            AddText(" === New task === ", true, 1, tcell.NewHexColor(0xD8DEE9))

        frame.SetBorder(false).
            SetBackgroundColor(tcell.NewHexColor(0x3B4252))

        return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
                AddItem(frame, height + 2, 1, false).
				AddItem(nil, 0, 1, false), width + 2, 1, false).
			AddItem(nil, 0, 1, false)
	}

    inputField = tview.NewInputField().
	  SetFieldTextColor(tcell.NewHexColor(0xD8DEE9)).
      SetFieldBackgroundColor(tcell.NewHexColor(0x4C566A))

    donefunc := func(key tcell.Key) {
        text := inputField.GetText()

        if key == tcell.KeyEnter && text != "" {
            list.Add(text)
        }

        inputField.SetText("")
        pages.HidePage("modal")
        //pages.SwitchToPage("background")
        app.SetFocus(&tbox)
    }

    inputField.SetDoneFunc(func(key tcell.Key) {
		    //app.Stop()
            donefunc(key)
	})

    flex := tview.NewFlex().
		SetFullScreen(true).
		SetDirection(tview.FlexRow).
		AddItem(&tbox, 0, 1, false)

    newModal := modal(inputField, 40, 1)

	pages.
		AddPage("background", flex, true, true).
		AddPage("modal", newModal, true, true)

    pages.HidePage("modal")

	if err := app.SetRoot(pages, true).SetFocus(&tbox).Run(); err != nil {
		panic(err)
	}
}
