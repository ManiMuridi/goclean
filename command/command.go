package command

import "github.com/ManiMuridi/goclean/syserr"

type Command interface {
	Execute() *Result
}

type ValidatedCommand interface {
	Command
	Validate() error
}

type ReversibleCommand interface {
	Revert() error
}

type NotFoundCommand struct{}

func (n *NotFoundCommand) Execute() *Result {
	return &Result{
		Error: syserr.ValidationError{"NotFound": "Command not found"},
		Data:  nil,
	}
}

func Execute(cmd Command) *Result {
	// TODO: check for nil command since this is an interface nil can be passed in
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

func MultiExec(cmds ...Command) []*Result {
	results := make([]*Result, 0)

	for i := range cmds {
		results = append(results, Execute(cmds[i]))
	}

	return results
}

func ExecRevert(rCmd ReversibleCommand) error {
	return rCmd.Revert()
}
