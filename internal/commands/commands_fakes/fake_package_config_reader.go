package commands_fakes

import (
	"github.com/buildpacks/pack/internal/buildpackage"
)

type FakePackageConfigReader struct {
	ReadCalledWithArg string
	ReadReturnConfig  buildpackage.Config
	ReadReturnError   error
}

func (r *FakePackageConfigReader) Read(path string) (buildpackage.Config, error) {
	r.ReadCalledWithArg = path

	return r.ReadReturnConfig, r.ReadReturnError
}

func NewFakePackageConfigReader(ops ...func(*FakePackageConfigReader)) *FakePackageConfigReader {
	fakePackageConfigReader := &FakePackageConfigReader{
		ReadReturnConfig: buildpackage.Config{},
		ReadReturnError:  nil,
	}

	for _, op := range ops {
		op(fakePackageConfigReader)
	}

	return fakePackageConfigReader
}
