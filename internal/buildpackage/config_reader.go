package buildpackage

import (
	"github.com/BurntSushi/toml"
	pubbldpkg "github.com/buildpacks/pack/buildpackage"
)

type ConfigReader struct {}

func (r *ConfigReader) Read(path string) (pubbldpkg.Config, error) {
	config := pubbldpkg.Config{}

	toml.DecodeFile(path, &config)

	return config, nil
}

//func ReadPackageConfig(path string) (buildpackage.Config, error) {
//	config := buildpackage.Config{}
//
//	configDir, err := filepath.Abs(filepath.Dir(path))
//	if err != nil {
//		return config, err
//	}
//
//	_, err = toml.DecodeFile(path, &config)
//	if err != nil {
//		return config, errors.Wrapf(err, "reading config %s", path)
//	}
//
//	absPath, err := paths.ToAbsolute(config.Buildpack.URI, configDir)
//	if err != nil {
//		return config, errors.Wrapf(err, "getting absolute path for %s", style.Symbol(config.Buildpack.URI))
//	}
//	config.Buildpack.URI = absPath
//
//	for i := range config.Dependencies {
//		uri := config.Dependencies[i].URI
//		if uri != "" {
//			absPath, err := paths.ToAbsolute(uri, configDir)
//			if err != nil {
//				return config, errors.Wrapf(err, "getting absolute path for %s", style.Symbol(uri))
//			}
//
//			config.Dependencies[i].URI = absPath
//		}
//	}
//
//	return config, nil
//}