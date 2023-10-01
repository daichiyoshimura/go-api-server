package env

import "os"

const (
	stageDev  = "DEV"
	stageTest = "TEST"
)

type Stage struct {
	stg string
}

func NewStage() *Stage {
	return &Stage{
		stg: os.Getenv("STAGE"),
	}
}

func (s *Stage) isDev() bool {
	return s.stg == stageDev
}

func (s *Stage) isTest() bool {
	return s.stg == stageTest
}
