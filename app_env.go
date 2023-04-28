package main

import (
	"log"
	"os"

	afs "github.com/spf13/afero"
)

type AppEnv struct {
	AppFs           afs.Fs
	TargetLinksPath string
	StorageRoot     string
	*log.Logger
}

func (app_env *AppEnv) Build() error {
	pathes := []string{app_env.TargetLinksPath, app_env.StorageRoot}
	for _, path := range pathes {
		err := app_env.AppFs.MkdirAll(path, os.FileMode(0777))
		if err != nil {
			return err
		}
	}
	return nil
}
