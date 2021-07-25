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
				// TODO: Close the application
				//app.Stop()
			}
			if key == tcell.KeyEnter {
				// TODO Open view of a sinlge task
				task.Table.SetSelectable(true, true)
			}
		}).
		SetSelectedFunc(func(row int, column int) {
			task.List.ToggleDone(row)
			// TODO: open task details
		}).
		SetInputCapture(func(capture *tcell.EventKey) *tcell.EventKey {
			if capture.Rune() == ' ' {
				// TODO: mark task as done

				return nil
			}

			return capture
		})

	task.Table.Draw(screen)
}

func Main() {
	app := tview.NewApplication()

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

	flex := tview.NewFlex().
		SetFullScreen(true).
		SetDirection(tview.FlexRow).
		AddItem(&tbox, 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(&tbox).Run(); err != nil {
		panic(err)
	}
}
