package applications

import (
	"marcocd/pkg/infras/manifest_reader"
	"os"
	"text/template"

	nomad "github.com/hashicorp/nomad/api"
)

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

func (renderer *ModuleTemplateRenderer) Render() ([]*nomad.Job, error) {
	manifest, err := renderer.manifestReader.Read(renderer.manifestPath)
	if err != nil {
		return nil, err
	}

	jobs := []*nomad.Job{}

	deliverables := manifest.Deliverables
	for _, deliverable := range deliverables {
		for _, templateFilePath := range deliverable.Resources {
			job, err := render(templateFilePath, manifest.Values)
			if err != nil {
				return nil, err
			}

			jobs = append(jobs, job)
		}
	}

	return jobs, nil
}

func render(templateFilePath string, values map[string]interface{}) (*nomad.Job, error) {
	tmpl, err := template.New(templateFilePath).ParseFiles(templateFilePath)
	if err != nil {
		return nil, err
	}

	f, err := os.CreateTemp("", "tmp")
	if err != nil {
		return nil, err
	}

	if err = tmpl.Execute(f, values); err != nil {
		return nil, err
	}

	return nil, nil
}
