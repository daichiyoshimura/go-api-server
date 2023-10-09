package getService

type Input struct {
	id int
}

func NewInput(id int) *Input {
	return &Input{
		id: id,
	}
}

func (i *Input) ID() int {
	return i.id
}
