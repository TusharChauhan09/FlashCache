// ! handle each client connection
package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/TusharChauhan09/flashcache/internal/protocol"
)

func (s *Server) Handle(conn net.Conn){
	defer conn.Close()
	fmt.Printf("client connected: %s\n",conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){ // scans until sender keeps on sending 
		input := scanner.Text()  // gets the line : Raw Sting 

		cmd,err := protocol.Parse(input)
		if err != nil {
			s.write(conn,protocol.Error(err))
			continue
		}

		err = protocol.Validate(cmd)
		if err != nil {
			s.write(conn,protocol.Error(err))
			continue
		}

		response := s.Route(cmd)

		s.write(conn, response)
	}

	fmt.Printf("client disconnected: %s\n", conn.RemoteAddr())
}

func (s *Server) write(conn net.Conn, msg string){
	_,_ = conn.Write([]byte(msg+"\n"))
} 
