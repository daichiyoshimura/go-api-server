package env

const (
	stageDev  = "DEV"
	stageTest = "TEST"
)

type Stage struct {
	stg string
}

func NewStage(stg string) *Stage {
	return &Stage{
		stg: stg,
	}
}

func (s *Stage) isDev() bool {
	return s.stg == stageDev
}

func (s *Stage) isTest() bool {
	return s.stg == stageTest
}
