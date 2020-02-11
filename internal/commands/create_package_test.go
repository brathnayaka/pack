package commands_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/heroku/color"
	"github.com/pkg/errors"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/spf13/cobra"

	"github.com/buildpacks/pack/internal/buildpackage"
	"github.com/buildpacks/pack/internal/commands"
	"github.com/buildpacks/pack/internal/commands/commands_fakes"
	"github.com/buildpacks/pack/internal/dist"
	"github.com/buildpacks/pack/internal/logging"
	h "github.com/buildpacks/pack/testhelpers"
)

func TestCreatePackageCommand(t *testing.T) {
	color.Disable(true)
	defer color.Disable(false)
	spec.Run(t, "Commands", testCreatePackageCommand, spec.Parallel(), spec.Report(report.Terminal{}))
}

func testCreatePackageCommand(t *testing.T, when spec.G, it spec.S) {
	when("CreatePackage#Execute", func() {
		it("reads package config from the configured path", func() {
			fakePackageConfigReader := commands_fakes.NewFakePackageConfigReader()
			expectedConfigPath := "/path/to/some/file"

			createPackageCommand := createPackageCommand(
				withConfigReader(fakePackageConfigReader),
				withConfigPath(expectedConfigPath),
			)

			err := createPackageCommand.Execute()
			h.AssertNil(t, err)

			h.AssertEq(t, fakePackageConfigReader.ReadCalledWithArg, expectedConfigPath)
		})

		it("logs an error and exits when package toml is invalid", func() {
			outBuf := &bytes.Buffer{}
			expectedErr := errors.New("it went wrong")

			createPackageCommand := createPackageCommand(
				withLogger(logging.NewLogWithWriters(outBuf, outBuf)),
				withConfigReader(
					commands_fakes.NewFakePackageConfigReader(whereReadReturns(buildpackage.Config{}, expectedErr)),
				),
			)

			err := createPackageCommand.Execute()
			h.AssertNotNil(t, err)

			h.AssertContains(t, outBuf.String(), fmt.Sprintf("ERROR: reading config: %s", expectedErr))
		})

		it("creates package with correct image name", func() {
			fakePackageCreator := &commands_fakes.FakePackageCreator{}

			createPackageCommand := createPackageCommand(
				withImageName("my-specific-image"),
				withPackageCreator(fakePackageCreator),
			)

			err := createPackageCommand.Execute()
			h.AssertNil(t, err)

			receivedOptions := fakePackageCreator.CreateCalledWithOptions

			h.AssertEq(t, receivedOptions.Name, "my-specific-image")
		})

		it("creates package with config returned by the reader", func() {
			fakePackageCreator := &commands_fakes.FakePackageCreator{}

			myConfig := buildpackage.Config{
				Buildpack: dist.BuildpackURI{URI: "test"},
			}

			createPackageCommand := createPackageCommand(
				withPackageCreator(fakePackageCreator),
				withConfigReader(commands_fakes.NewFakePackageConfigReader(whereReadReturns(myConfig, nil))),
			)

			err := createPackageCommand.Execute()
			h.AssertNil(t, err)

			receivedOptions := fakePackageCreator.CreateCalledWithOptions

			h.AssertEq(t, receivedOptions.Config, myConfig)
		})
	})
}

type packageCommandConfig struct {
	logger         *logging.LogWithWriters
	configReader   *commands_fakes.FakePackageConfigReader
	packageCreator *commands_fakes.FakePackageCreator

	imageName  string
	configPath string
}

type packageCommandOption func(config *packageCommandConfig)

func createPackageCommand(ops ...packageCommandOption) *cobra.Command {
	config := &packageCommandConfig{
		logger:         logging.NewLogWithWriters(&bytes.Buffer{}, &bytes.Buffer{}),
		configReader:   commands_fakes.NewFakePackageConfigReader(),
		packageCreator: &commands_fakes.FakePackageCreator{},

		imageName:  "some-image-name",
		configPath: "/path/to/some/file",
	}

	for _, op := range ops {
		op(config)
	}

	cmd := commands.CreatePackage(config.logger, config.packageCreator, config.configReader)
	cmd.SetArgs([]string{config.imageName, "--package-config", config.configPath})

	return cmd
}

func withLogger(logger *logging.LogWithWriters) packageCommandOption {
	return func(config *packageCommandConfig) {
		config.logger = logger
	}
}

func withConfigReader(reader *commands_fakes.FakePackageConfigReader) packageCommandOption {
	return func(config *packageCommandConfig) {
		config.configReader = reader
	}
}

func withPackageCreator(creator *commands_fakes.FakePackageCreator) packageCommandOption {
	return func(config *packageCommandConfig) {
		config.packageCreator = creator
	}
}

func withImageName(name string) packageCommandOption {
	return func(config *packageCommandConfig) {
		config.imageName = name
	}
}

func withConfigPath(path string) packageCommandOption {
	return func(config *packageCommandConfig) {
		config.configPath = path
	}
}

func whereReadReturns(config buildpackage.Config, err error) func(*commands_fakes.FakePackageConfigReader) {
	return func(r *commands_fakes.FakePackageConfigReader) {
		r.ReadReturnConfig = config
		r.ReadReturnError = err
	}
}
