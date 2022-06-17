package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	window := a.NewWindow(WindowTitle)
	window.Resize(fyne.NewSize(WindowWidth, WindowHeight))

	SetFakeData()

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
			var err error
			switch uid {
			case "Branches":
				children, err = branchList.Get()
			case "Remotes":
				children, err = remoteList.Get()
			case "Tags":
				children, err = tagList.Get()
			case "Stashes":
				children, err = stashList.Get()
			case "Submodules":
				children, err = submoduleList.Get()
			default:
				children = append(children, "Branches", "Remotes", "Tags", "Stashes", "Submodules")
			}
			checkErr(err)
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

var (
	branchList    = binding.NewStringList()
	remoteList    = binding.NewStringList()
	tagList       = binding.NewStringList()
	stashList     = binding.NewStringList()
	submoduleList = binding.NewStringList()
)

func SetFakeData() {
	err := branchList.Set([]string{"master", "develop"})
	checkErr(err)

	err = remoteList.Set([]string{"origin/master", "origin/develop"})
	checkErr(err)

	err = tagList.Set([]string{"0.1.0", "0.2.0"})
	checkErr(err)

	err = stashList.Set([]string{"保存代码1"})
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
