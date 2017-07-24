package views

import (
	"html/template"
	"path/filepath"
	"net/http"
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

// Don't fully get this Render func.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

type View struct {
	Template *template.Template
	Layout string
}
