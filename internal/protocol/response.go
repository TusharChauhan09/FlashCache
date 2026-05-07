// ! format responses/errors
package protocol

import "fmt"

func OK() string{
	return "OK"
}

func Pong() string{
	return "PONG"
}

func Value(value string) string {
	return value
}

func Integer(value int) string {
	return fmt.Sprintf("%d", value)
}

func Error(err error) string {
	return fmt.Sprintf("ERR %s", err.Error())
}