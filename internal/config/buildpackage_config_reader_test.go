package config

import (
	"io/ioutil"
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
	spec.Run(t, "Buildpackage Config", testBuildpackageConfigReader, spec.Parallel(), spec.Report(report.Terminal{}))
}

func testBuildpackageConfigReader(t *testing.T, when spec.G, it spec.S) {
	when("#Read", func() {
		it("returns correct config when provided toml file is valid", func() {
			configSource, err := filepath.Abs(filepath.Join("testdata", "package.toml"))
			h.AssertNil(t, err)

			configLocation := ioutil.TempFile()
		})
	})
}
