package command

type Executor struct {
	Commands []Command
}

func (e *Executor) ExecuteCommands() {
	for _, e := range e.Commands {
		e.Execute()
	}
}
