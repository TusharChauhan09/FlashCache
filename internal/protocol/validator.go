// ! validate command args
package protocol

import "fmt"

func Validate(cmd *Command)error {
	switch cmd.Name {
	case "PING","STATS":
		if len(cmd.Args) != 0 {
			return fmt.Errorf("%s takes no arguments", cmd.Name)
		}
	case "GET","DEL","EXISTS","POP","LEN":
		if len(cmd.Args)!=1 {
			return fmt.Errorf("%s requires exactly 1 argument",cmd.Name)
		}
	case "SET","PUSH":
		if len(cmd.Args)!=2{
			return fmt.Errorf("%s requires exactly 2 arguments",cmd.Name)
		}
	default:
		return fmt.Errorf("unkown command: %s",cmd.Name)
	}
	return nil
}