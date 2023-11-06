package applications

import (
	"bytes"
	"fmt"
	"marcocd/pkg/infras/manifest_reader"
	"os"
	"text/template"
)

type TemplateRenderer interface {
	Render() ([]string, error)
}

type ModuleTemplateRenderer struct {
	manifestPath   string
	manifestReader manifest_reader.ModuleManifestReader
}

func NewModuleTemplateRenderer(path string, manifestReader manifest_reader.ModuleManifestReader) TemplateRenderer {
	return &ModuleTemplateRenderer{
		manifestPath:   path,
		manifestReader: manifestReader,
	}
}

func (renderer *ModuleTemplateRenderer) Render() ([]string, error) {
	manifest, err := renderer.manifestReader.Read(renderer.manifestPath)
	if err != nil {
		return nil, err
	}

	jobSpecs := []string{}

	for _, deliverable := range manifest.Deliverables {
		for _, templateFilePath := range deliverable.Resources {
			jobSpec, err := render(templateFilePath, manifest.Values)
			if err != nil {
				return nil, err
			}

			jobSpecs = append(jobSpecs, jobSpec)
		}
	}

	return jobSpecs, nil
}

func render(templateFilePath string, values map[string]interface{}) (string, error) {
	fmt.Println(templateFilePath)

	f, err := os.ReadFile(templateFilePath)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New(templateFilePath).Parse(string(f))
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	if err = tmpl.Execute(&b, values); err != nil {
		return "", err
	}

	return b.String(), nil
}
