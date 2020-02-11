package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// LoadTemplatesGlob recursively loads the templates within the dir and sub-dirs.
func LoadTemplatesGlob() *template.Template {
	tpl := template.New("")
	err := filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = tpl.ParseFiles(path)
			if err != nil {
				log.Fatal(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return tpl
}
