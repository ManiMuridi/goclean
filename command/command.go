package command

type Command interface {
	Execute() *Result
}

func Execute(cmd Command) *Result {
	return cmd.Execute()
}
