package govern

type (
	Result struct {
		Errors []Error
	}
)

func (r Result) IsOk() bool {
	return len(r.Errors) < 1
}
func (r Result) IsErr() bool {
	return len(r.Errors) > 0
}
