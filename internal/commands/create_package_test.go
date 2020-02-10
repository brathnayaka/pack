package commands_test

import (
	"github.com/buildpacks/pack/internal/commands"
	h "github.com/buildpacks/pack/testhelpers"
	"github.com/heroku/color"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreatePackageCommand(t *testing.T) {
	color.Disable(true)
	defer color.Disable(false)
	spec.Run(t, "Commands", testCreatePackageCommand, spec.Parallel(), spec.Report(report.Terminal{}))
}

func testCreatePackageCommand(t *testing.T, when spec.G, it spec.S) {
	when("CreatePackage#Execute", func() {
		it("reads package config from the configured path", func() {
			createPackageCommand := commands.CreatePackage(nil, fakeClient, fakePackageConfigReader)

			os.Create(ioutil.TempFile("", ""))
			createPackageCommand.SetArgs([]string{"-package-config", "/path/to/some/file"})

			err := createPackageCommand.Execute()
			h.AssertNil(t, err)

			h.AssertEq(t, fakePackageConfigReader.readCalledWithArg, "path/to/some/file")
		})

		it("logs an error and exits when package toml is invalid", func() {
			fakePackageConfigReader(returnsForRead(errors.New("it went wrong")))


		})
	})
}