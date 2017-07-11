package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

// Func that globs all the go templates
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func NewView(layout string, files ...string) *View {
	// Reads all files out of the dir
	files = append(files, layoutFiles()...)
	// Expands all the files out of the files slice
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
