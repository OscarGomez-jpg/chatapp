// Package server implements the TCP chat server logic.
package server

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Server struct {
	clients map[net.Conn]bool
	mu      sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients: make(map[net.Conn]bool),
	}
}

func (s *Server) Broadcast(sender net.Conn, msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for conn := range s.clients {
		if conn != sender {
			fmt.Fprintln(conn, msg)
		}
	}
}

func (s *Server) HandleConn(conn net.Conn) {
	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
		conn.Close()
	}()
	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		s.Broadcast(conn, msg)
	}
}

func (s *Server) Listen(address string) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go s.HandleConn(conn)
	}
}
