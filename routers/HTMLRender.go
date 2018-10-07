package routers

import (
	"github.com/gin-contrib/multitemplate"
	log "github.com/sirupsen/logrus"
	"html/template"
	"path/filepath"
)

func loadMultiTemplates(layoutDir, includeDir, ext string, funcmap template.FuncMap) multitemplate.Renderer {
	render := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(layoutDir + "*." + ext)
	if err != nil {
		log.WithField("layoutdir", layoutDir).Fatal(err)
	}
	includes, err := filepath.Glob(layoutDir + includeDir + "*." + ext)
	if err != nil {
		log.WithField("includedir", includeDir).Fatal(err)
	}

	for _, layout := range layouts {
		files := []string{layout}
		files = append(files, includes...)
		log.WithField("templateName", filepath.Base(layout)).Debug("add layout template")
		render.AddFromFilesFuncs(filepath.Base(layout), funcmap, files...)
	}
	return render
}
