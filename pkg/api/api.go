package api

func init() {
	initDocker()
}

func CloseAll() {
	CloseDockerClient()
}
