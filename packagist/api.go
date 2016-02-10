package packagist

type PackageWrapper struct {
	Package Package `json:"package"`
}

type Package struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Repository  string                    `json:"repository"`
	Versions    map[string]PackageVersion `json:"versions"`
}

type PackageVersion struct {
	Require map[string]string `json:"require"`
}
