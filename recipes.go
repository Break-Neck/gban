package main

type BinarySource interface {
	Upgrade(appEnv *AppEnv)
}

type GithubBinSource struct {
	repo string
}

func (source *GithubBinSource) Upgrade(appEnv *AppEnv) {
	appEnv.Print("I'm in a github handler!")
}
