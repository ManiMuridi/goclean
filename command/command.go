package command

type Command interface {
	Execute() *Result
}
