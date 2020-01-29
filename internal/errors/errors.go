package errors

type Error string

func (err Error) Error() string {
	return string(err)
}

const (
	Inconsistent Error = "inconsistent"
)
