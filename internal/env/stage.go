package env

const (
	stageDev = "DEV"
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
