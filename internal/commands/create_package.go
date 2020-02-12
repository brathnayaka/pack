package commands

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/buildpacks/pack"
	pubbldpkg "github.com/buildpacks/pack/buildpackage"
	"github.com/buildpacks/pack/internal/style"
	"github.com/buildpacks/pack/logging"
)

type CreatePackageFlags struct {
	PackageTomlPath string
	Publish         bool
	NoPull          bool
}

type PackageConfigReader interface {
	Read(path string) (pubbldpkg.Config, error)
}

type PackageCreator interface {
	CreatePackage(ctx context.Context, options pack.CreatePackageOptions) error
}

func CreatePackage(logger logging.Logger, client PackageCreator, packageConfigReader PackageConfigReader) *cobra.Command {
	var flags CreatePackageFlags
	ctx := createCancellableContext()
	cmd := &cobra.Command{
		Use:   "create-package <image-name> --package-config <package-config-path>",
		Args:  cobra.ExactArgs(1),
		Short: "Create package",
		RunE: logError(logger, func(cmd *cobra.Command, args []string) error {
			config, err := packageConfigReader.Read(flags.PackageTomlPath)
			if err != nil {
				return errors.Wrap(err, "reading config")
			}

			imageName := args[0]
			if err := client.CreatePackage(ctx, pack.CreatePackageOptions{
				Name:    imageName,
				Config:  config,
				Publish: flags.Publish,
				NoPull:  flags.NoPull,
			}); err != nil {
				return err
			}
			action := "created"
			if flags.Publish {
				action = "published"
			}
			logger.Infof("Successfully %s package %s", action, style.Symbol(imageName))
			return nil
		}),
	}
	cmd.Flags().StringVarP(&flags.PackageTomlPath, "package-config", "p", "", "Path to package TOML config (required)")
	cmd.MarkFlagRequired("package-config")
	cmd.Flags().BoolVar(&flags.Publish, "publish", false, "Publish to registry")
	cmd.Flags().BoolVar(&flags.NoPull, "no-pull", false, "Skip pulling packages before use")
	AddHelpFlag(cmd, "create-package")

	cmd.Hidden = true

	return cmd
}
