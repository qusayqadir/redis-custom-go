// Package server handles networking: it opens a TCP listener and runs a
// read/reply loop for each connected client.
package server

import (
	"log"
	"net"

	"github.com/qusayqadir/redis-custom-go/internal/store"
	// "github.com/qusayqadir/redis-custom-go/internal/resp"
	
)

// talks to the client 
func handleConn(conn net.Conn, db *store.Store) {

}


// if error nil then there is an issue 
func Run(port_num string) error {

	db := store.New()
	//reserver the port on my machine 
	listener, err := net.Listen("tcp", port_num) 
	
	if err != nil {
		log.Printf("unable to connect to port: %s", port_num)	
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue 
		}
	
		// go is keyword to run concurrently 
		// goroutine 
		go handleConn(conn, db) 
	}
	
} 



