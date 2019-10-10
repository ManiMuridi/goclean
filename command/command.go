package command

type Command interface {
	Execute() *Result
}

type ValidatedCommand interface {
	Command
	Validate() error
}

func Execute(cmd Command) *Result {
	if _, ok := cmd.(ValidatedCommand); ok {
		if err := cmd.(ValidatedCommand).Validate(); err != nil {
			return &Result{
				Error: err,
				Data:  nil,
			}
		}
	}

	return cmd.Execute()
}
