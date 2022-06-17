package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	window := a.NewWindow(WindowTitle)
	window.Resize(fyne.NewSize(WindowWidth, WindowHeight))

	commitTypes := widget.NewList(
		func() int {
			return len(list1Items)
		},
		func() fyne.CanvasObject {
			return &widget.Label{}
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(list1Items[id])
		},
	)
	commitTypePage := container.NewVBox(
		widget.NewLabel("gameplay-discovery"),
		commitTypes,
	)

	list2 := widget.NewTree(
		func(uid string) (children []string) {
			switch uid {
			case "Branches":
				children = append(children, Branches...)
			case "Remotes":
				children = append(children, Remotes...)
			case "Tags":
				children = append(children, Tags...)
			case "Stashes":
				children = append(children, Stashes...)
			case "Submodules":
				children = append(children, Submodules...)
			default:
				children = append(children, "Branches", "Remotes", "Tags", "Stashes", "Submodules")
			}
			return
		},
		func(uid string) bool {
			switch uid {
			case "", "Branches", "Remotes", "Tags", "Stashes", "Submodules":
				return true
			default:
				return false
			}
		},
		func(branch bool) fyne.CanvasObject {
			return &widget.Label{}
		},
		func(uid string, branch bool, node fyne.CanvasObject) {
			node.(*widget.Label).SetText(uid)
		},
	)

	leftContainer := container.NewVSplit(commitTypePage, list2)
	rightContainer := container.NewVSplit(commitTypePage, list2)
	wholeContainer := container.NewHSplit(leftContainer, rightContainer)

	window.SetContent(wholeContainer)
	window.ShowAndRun()
}

const (
	WindowTitle  = "Datahub Desktop"
	WindowWidth  = 1000.0
	WindowHeight = 1000.0
)

var (
	list1Items = []string{
		"Local Changes",
		"All Commits",
	}
)

var treeMap = make(map[string]string)

var (
	Branches   = []string{"master", "develop"}
	Remotes    = []string{"origin/master", "origin/develop"}
	Tags       = []string{"0.1.0", "0.2.0"}
	Stashes    = []string{"保存代码1"}
	Submodules []string
)
