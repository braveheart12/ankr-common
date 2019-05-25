package pgrpc

import (
	"context"
	"net"
	"time"

	"github.com/hashicorp/yamux"
	"google.golang.org/grpc"
)

type Session struct {
	*yamux.Session

	Name string
	Opts []grpc.DialOption
}

type Hook interface {
	OnAccept(key *string, conn *net.Conn) error
	OnBuild(key *string, conn *Session) error
	OnClose(key string, conn *Session)
}

type EmptyHook struct{}

func (*EmptyHook) OnAccept(key *string, conn *net.Conn) error { return nil }
func (*EmptyHook) OnBuild(key *string, conn *Session) error   { return nil }
func (*EmptyHook) OnClose(key string, conn *Session)          {}

// ping handle
type ping struct {
	addr string
	sess *Session
	hook Hook
}

func (p *ping) Ping(stream StreamPing_PingServer) error {
	go func() {
		for {
			stream.Recv()
		}
	}()

	for range time.Tick(5 * time.Second) {
		if err := stream.Send(&PingMessage{}); err != nil {
			p.sess.Close()
			return err
		}
	}
	return nil
}

func pgrpcKeepalive(sess *Session, logger yamux.Logger) {
	finalizer := func(err error) {
		sess.GoAway()
		if logger != nil {
			logger.Println("close session:", err)
		}
	}

	conn, err := grpc.DialContext(context.Background(), sess.Name,
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return sess.Open()
		}))
	if err != nil {
		finalizer(err)
		return
	}

	client, err := NewStreamPingClient(conn).Ping(context.Background())
	if err != nil {
		finalizer(err)
		return
	}

	go func() {
		for {
			client.Recv()
		}
	}()
	go func() {
		for range time.Tick(5 * time.Second) {
			if err := client.Send(&PingMessage{}); err != nil {
				finalizer(err)
				return
			}
		}
	}()
}
