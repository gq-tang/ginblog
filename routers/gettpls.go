package routers

import (
	"html/template"
	"path/filepath"
	"strings"

	"ginblog/views"
	log "github.com/sirupsen/logrus"
)

func loadTemplate(funcmap template.FuncMap, ext string) (*template.Template, error) {
	t := template.New("")
	names := views.AssetNames()
	for _, name := range names {
		info, err := views.AssetInfo(name)
		if err != nil {
			log.WithError(err).Error("get asset info error")
			continue
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ext) {
			continue
		}
		data, err := views.Asset(name)
		if err != nil {
			log.WithError(err).Error("get asset data error")
			continue
		}

		t, err = t.New(filepath.Base(info.Name())).Funcs(funcmap).Parse(string(data))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
