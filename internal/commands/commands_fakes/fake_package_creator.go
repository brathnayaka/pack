package commands_fakes

import (
	"context"

	"github.com/buildpacks/pack"
)

type FakePackageCreator struct {
	CreateCalledWithOptions pack.CreatePackageOptions
}

func (c *FakePackageCreator) CreatePackage(ctx context.Context, opts pack.CreatePackageOptions) error {
	c.CreateCalledWithOptions = opts

	return nil
}
