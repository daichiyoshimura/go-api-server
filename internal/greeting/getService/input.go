package getService

type Input struct {
	id uint
}

func NewInput(id uint) *Input {
	return &Input{
		id: id,
	}
}

func (i *Input) ID() uint {
	return i.id
}
