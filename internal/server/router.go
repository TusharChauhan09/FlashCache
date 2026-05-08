// ! route commands to store methods
package server

import (
	"fmt"

	"github.com/TusharChauhan09/flashcache/internal/protocol"
	"github.com/TusharChauhan09/flashcache/internal/store"
)

func (s *Server) Route(cmd *protocol.Command) string {
	
	switch cmd.Name{

	case "PING":
		return protocol.Pong()

	case "SET":
		s.store.Set(cmd.Args[0],cmd.Args[1])
		return protocol.OK()
	
	case "GET": 
		value, err := s.store.Get(cmd.Args[0])
		if err != nil {
			return protocol.Error(err)
		}
		return protocol.Value(value)

	case "DEL":
		deleted := s.store.Delete(cmd.Args[0])
		if !deleted {
			return protocol.Error(store.ErrKeyNotFound)
		}
		return protocol.OK()

	case "EXIST":
		exists := s.store.Exist(cmd.Args[0])
		if !exists {
			return protocol.Integer(1)
		}
		return protocol.Integer(0)
	
	case "PUSH":
		s.store.Push(cmd.Args[0], cmd.Args[1])
		return protocol.OK()

	case "POP":
		value, err := s.store.Pop(cmd.Args[0])
		if err != nil {
			return protocol.Error(err)
		}
		return protocol.Value(value)

	case "LEN":
		length := s.store.Len(cmd.Args[0])
		return protocol.Integer(length)

	case "STATS":
		return protocol.Value("FlashCache running")

	default:
		return protocol.Error(fmt.Errorf("unknown command"))
	}
}