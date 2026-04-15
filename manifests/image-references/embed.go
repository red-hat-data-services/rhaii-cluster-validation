package imagereferences

import (
	_ "embed"
)

// ImageReferencesYAML contains the embedded default image references.
//
//go:embed image-references.yaml
var ImageReferencesYAML string
