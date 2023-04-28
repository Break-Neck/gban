package main

import (
	"log"

	afs "github.com/spf13/afero"
)

func main() {
	app_env := AppEnv{
		AppFs:           afs.NewOsFs(),
		TargetLinksPath: "tmp/target_link_path",
		StorageRoot:     "tmp/storage_root",
		Logger:          log.Default(),
	}
	app_env.Build()
	log.Print("Done!")
}
