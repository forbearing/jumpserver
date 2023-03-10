package jumperserver

type Option interface {
	apply(*Jumpserver) error
}

type optionFunc func(*Jumpserver) error

func (f optionFunc) apply(j *Jumpserver) error {
	return f(j)
}
