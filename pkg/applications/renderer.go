package applications

import (
	"bytes"
	"fmt"
	"marcocd/pkg/infras/manifest_reader"
	"os"
	"text/template"

	nomad "github.com/hashicorp/nomad/api"
)

type TemplateRenderer interface {
	Render() ([]*nomad.Job, error)
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

func (renderer *ModuleTemplateRenderer) Render() ([]*nomad.Job, error) {
	manifest, err := renderer.manifestReader.Read(renderer.manifestPath)
	if err != nil {
		return nil, err
	}

	jobs := []*nomad.Job{}

	for _, deliverable := range manifest.Deliverables {
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
	fmt.Println(templateFilePath)

	f, err := os.ReadFile(templateFilePath)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(templateFilePath).Parse(string(f))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err = tmpl.Execute(&b, values); err != nil {
		return nil, err
	}

	fmt.Println(b.String())

	return nil, nil
}
