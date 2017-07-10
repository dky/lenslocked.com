package views

import "html/template"

func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	//template.ParseFiles("a", "b", "c") ... below "unravels the slice"
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View { 
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
