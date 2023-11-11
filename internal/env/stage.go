package env

import (
	"slices"

	"github.com/cockroachdb/errors"
)

const (
	stageDev  string = "DEV"
	stageTest string = "TEST"
)

const (
	errInvalidStage string = "invalid stage env: %v"
)

func stageCollection() []string {
	return []string{
		stageDev,
		stageTest,
	}
}

type Stage struct {
	stg string
}

func NewStage(stg string) (*Stage, error) {
	if !slices.Contains(stageCollection(), stg) {
		return nil, errors.Newf(errInvalidStage)
	}
	return &Stage{
		stg: stg,
	}, nil
}

func (s *Stage) isDev() bool {
	return s.stg == stageDev
}

func (s *Stage) isTest() bool {
	return s.stg == stageTest
}
