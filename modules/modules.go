package modules

var Commands = map[string]func(){}

func registerCommand(name string, command func()) {
	Commands[name] = command
}
