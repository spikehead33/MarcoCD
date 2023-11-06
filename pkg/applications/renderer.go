package applications

import (
	"bytes"
	"marcocd/pkg/domains"
	"os"
	"text/template"
)

type TemplateRenderer interface {
	Render() ([]string, error)
}

type ModuleTemplateRenderer struct {
	moduleManifest *domains.ModuleManifest
}

func NewModuleTemplateRenderer(moduleManifest *domains.ModuleManifest) TemplateRenderer {
	return &ModuleTemplateRenderer{
		moduleManifest: moduleManifest,
	}
}

func (renderer *ModuleTemplateRenderer) Render() ([]string, error) {
	jobSpecs := []string{}
	for _, deliverable := range renderer.moduleManifest.Deliverables {
		for _, templateFilePath := range deliverable.Resources {
			jobSpec, err := render(
				templateFilePath, renderer.moduleManifest.Values)
			if err != nil {
				return nil, err
			}

			jobSpecs = append(jobSpecs, jobSpec)
		}
	}

	return jobSpecs, nil
}

func render(templateFilePath string, values map[string]interface{}) (string, error) {
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
