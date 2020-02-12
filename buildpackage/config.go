package buildpackage

import "github.com/buildpacks/pack/internal/dist"

type Config struct {
	Buildpack    dist.BuildpackURI `toml:"buildpack"`
	Dependencies []dist.ImageOrURI `toml:"dependencies"`
}

