package getService

type Output struct {
	msg string
}

func NewOutput(msg string) *Output {
	return &Output{
		msg:msg,
	}
}

func (o *Output) Message() string {
	return o.msg
}
