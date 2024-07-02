package main

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
	"jediproj.io/tview-tests/applications"
	"jediproj.io/tview-tests/internal"
	"jediproj.io/tview-tests/internal/providers/cd/argocd"
	"jediproj.io/tview-tests/widgets"
	"jediproj.io/tview-tests/window"
)

const (
	defaultPage = "applications"
	title       = "Release Manager"
)

var (
	defaultConfigPath = path.Join(os.Getenv("HOME"), ".portsentinel.yaml")
)

func main() {
	configuration := &internal.Configuration{
		Argocd: &internal.ArgoCDConfiguration{},
	}
	configData, err := os.ReadFile(defaultConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	yaml.Unmarshal(configData, configuration)
	var list = widgets.NewButtonList()
	var w = window.Window{}
	var a = applications.NewManager(defaultPage).SetWindow(&w)
	w.InitWindow()
	provider := argocd.ArgoCDProvider{
		Configuration: *configuration.Argocd,
	}
	list.SetBorder(true).
		SetTitle(title)
	for _, applicaiton := range provider.GetApplications() {
		app := applications.NewApplication(string(applicaiton)).SetWindowManager(a)
		list.AddItem(app)
	}
	w.Pages.AddPage(defaultPage, list, true, true)
	w.App.Run()
}
