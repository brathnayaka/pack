package buildpackage_test

import (
	"github.com/buildpacks/pack/internal/buildpackage"
	"path/filepath"
	"testing"

	h "github.com/buildpacks/pack/testhelpers"
	"github.com/heroku/color"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestBuildpackageConfigReader(t *testing.T) {
	color.Disable(true)
	defer color.Disable(false)
	spec.Run(t, "Buildpackage Config Reader", testBuildpackageConfigReader, spec.Parallel(), spec.Report(report.Terminal{}))
}

func testBuildpackageConfigReader(t *testing.T, when spec.G, it spec.S) {
	when("#Read", func() {
		it("returns correct config when provided toml file is valid", func() {
			configSource := filepath.Join("testdata", "package.toml")

			packageConfigReader := buildpackage.ConfigReader{}

			config, err := packageConfigReader.Read(configSource)
			h.AssertNil(t, err)

			h.AssertEq(t, config.Buildpack.URI, "https://example.com/bp/a.tgz")
			h.AssertEq(t, len(config.Dependencies), 2)
			h.AssertEq(t, config.Dependencies[0].URI, "https://example.com/bp/b.tgz")
			h.AssertEq(t, config.Dependencies[1].ImageRef.ImageName, "registry.example.com/bp/c")
		})
	})
}

// todo: Handle relative paths for URI

// todo: add public wrapper to allow reading config for API consumers



