package pgrpc

import (
	"net"

	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	sess *Session

	addr   string
	config *yamux.Config
	hook   Hook
}

func NewServer(addr string, hook Hook, conf *yamux.Config, opt ...grpc.ServerOption) *Server {
	return &Server{
		Server: grpc.NewServer(opt...),
		addr:   addr,
		config: conf,
		hook:   hook,
	}
}

func (s *Server) Serve() error {
	if err := s.dial(); err != nil {
		return err
	}

	RegisterStreamPingServer(s.Server, &ping{s.addr, s.sess, s.hook})

	return errors.Wrap(s.Server.Serve(s.sess), "serve grpc")
}

func (s *Server) dial() (err error) {
	var conn net.Conn
	conn, err = net.Dial("tcp", s.addr)
	if err != nil {
		return errors.Wrap(err, "tcp dail")
	}

	defer func() {
		if err != nil {
			conn.Close()
		}
	}()

	if s.hook != nil {
		if err := s.hook.OnAccept(&s.addr, &conn); err != nil {
			return errors.Wrap(err, "on accept hook")
		}
	}

	session, err := yamux.Server(conn, s.config)
	if err != nil {
		return errors.Wrap(err, "build sess")
	} else if _, err = session.Ping(); err != nil {
		return errors.Wrap(err, "ping session")
	}

	s.sess = &Session{Session: session, Name: s.addr}

	if s.hook != nil {
		if err := s.hook.OnBuild(&s.addr, s.sess); err != nil {
			return errors.Wrap(err, "build hook")
		}
	}

	go func() {
		<-s.sess.CloseChan()
		if s.hook != nil {
			s.hook.OnClose(s.addr, s.sess)
		}
	}()

	return nil
}
