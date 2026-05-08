// ! start TCP listener
package server

import (
	"fmt"
	"net"
)

func (s *Server) Start() error {
	address := s.config.Host + ":" + s.config.Port

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	defer listener.Close()

	fmt.Printf("FlashCache server running on %s\n", address)

	for{
		conn,err := listener.Accept()
		if err!=nil{
			fmt.Println("accept error: ",err)
			continue // Keep server alive , Never crash server because one accept failed
		}
		
		// Each client gets separate worker
		go s.Handle(conn)
	}

}