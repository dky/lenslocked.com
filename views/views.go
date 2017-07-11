package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func NewView(layout string, files ...string) *View {
	/*
	files = append(files,
		"views/layouts/footer.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/bootstrap.gohtml")
	*/
	files = append(files, layoutFiles()...)
	//template.ParseFiles("a", "b", "c") ... below "unravels the slice"
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View { 
		Template: t,
		Layout: layout,
	}
}

type View struct {
	Template *template.Template
	Layout string
}
