package cli

type event struct {
	cmdName string
	args    []string
	start   func(cmdName string, args []string)
	finish  func(cmdName string, args []string, recv interface{})
}

func newEvent() *event {
	return &event{
		start:  func(cmdName string, args []string) {},
		finish: func(cmdName string, args []string, recv interface{}) {},
	}
}

func (e *event) setCmdNameArgs(cmdName string, args []string) {
	e.cmdName = cmdName
	e.args = args
}

func (e *event) setStart(start func(cmdName string, args []string)) {
	if nil == start {
		return
	}
	e.start = start
}

func (e *event) setFinish(finish func(cmdName string, args []string, recv interface{})) {
	if nil == finish {
		return
	}
	e.finish = finish
}

func (e *event) executeStart() {
	e.start(e.cmdName, e.args)
}

func (e *event) executeFinish() {
	e.finish(e.cmdName, e.args, recover())
}
