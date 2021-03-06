package command

type GrpcCommandStruct struct {
	Name   string
	Method func(args []string) (string)
}

type CommandMap map[string]func(args []string) (string)

var Commands CommandMap

func init(){
	Commands =make(CommandMap)
}
func addCommand(cmd GrpcCommandStruct ){
	Commands[cmd.Name]=cmd.Method
}
func init() {
	Commands = make(CommandMap)
	addCommand(GrpcCommandStruct{Name: "getScreenSize", Method: getScreenSize})
	addCommand(GrpcCommandStruct{Name: "renameFile", Method: renameFile})
}


