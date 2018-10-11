package client

const (
	PublishCatalogConfigType         = "publishCatalogConfig"
	PublishCatalogConfigFieldCatalog = "catalog"
	PublishCatalogConfigFieldPath    = "path"
	PublishCatalogConfigFieldVersion = "version"
)

type PublishCatalogConfig struct {
	Catalog string `json:"catalog,omitempty" yaml:"catalog,omitempty"`
	Path    string `json:"path,omitempty" yaml:"path,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
