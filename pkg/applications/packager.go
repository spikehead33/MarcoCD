package applications

import (
	"marcocd/pkg/infras/manifest_reader"
	"marcocd/pkg/infras/tar_executor"
)

type Packager interface {
	Package() error
}

type packager struct {
	in             string
	ou             string
	version        string
	manifestReader manifest_reader.ModuleManifestReader
	tarExecutor    tar_executor.TarExecutor
}

func NewPackager(
	in, ou, version string,
	manifestReader manifest_reader.ModuleManifestReader,
	tarExecutor tar_executor.TarExecutor) Packager {
	return &packager{
		in:             in,
		ou:             ou,
		version:        version,
		manifestReader: manifestReader,
		tarExecutor:    tarExecutor,
	}
}

func (pack *packager) Package() error {
	manifest, err := pack.manifestReader.Read(pack.in)
	if err != nil {
		return err
	}

	deliverables := manifest.Deliverables
	resources := []string{pack.in} // add the marcocd.yaml, the manifest into the tar target

	for _, deliverables := range deliverables {
		resources = append(resources, deliverables.Resources...)
	}

	output := pack.ou
	if pack.ou == "" {
		output = manifest.Name + "_" + pack.version + ".tar"
	}

	return pack.tarExecutor.Tar(output, resources)
}
