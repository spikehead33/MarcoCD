package applications

import (
	"marcocd/pkg/domains"
	"marcocd/pkg/infras/tar_executor"
)

type Packager interface {
	Package() error
}

type packager struct {
	manifestPath   string
	ou             string
	version        string
	moduleManifest *domains.ModuleManifest
	tarExecutor    tar_executor.TarExecutor
}

func NewPackager(
	in, ou, version string,
	moduleManifest *domains.ModuleManifest,
	tarExecutor tar_executor.TarExecutor) Packager {
	return &packager{
		ou:             ou,
		version:        version,
		moduleManifest: moduleManifest,
		tarExecutor:    tarExecutor,
	}
}

func (pack *packager) Package() error {
	deliverables := pack.moduleManifest.Deliverables
	resources := []string{pack.manifestPath} // add the marcocd.yaml, the manifest into the tar target

	for _, deliverables := range deliverables {
		resources = append(resources, deliverables.Resources...)
	}

	output := pack.ou
	if pack.ou == "" {
		output = pack.moduleManifest.Name + "_" + pack.version + ".tar"
	}

	return pack.tarExecutor.Tar(output, resources)
}
