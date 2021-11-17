package generic

import (
	"context"
	"github.com/pydio/cells/v4/common/server"
	"net"
)

type Server struct {
	*server.ServerImpl

	ctx      context.Context
	handlers []func() error
}

type Handler interface {
	Start() error
	Stop() error
}

var (
	Default = New(context.Background())
)

func Register(s *Server) {
	Default = s
}

func New(ctx context.Context) *Server {
	return &Server{
		ctx: ctx,
		ServerImpl: &server.ServerImpl{},
	}
}

func (s *Server) RegisterHandler(h Handler) {
	s.Handle(h.Start)
	s.RegisterAfterServe(h.Stop)
}

func (s *Server) Handle(h func() error) {
	s.handlers = append(s.handlers, h)
}

func (s *Server) Serve(l net.Listener) error {
	if err := s.BeforeServe(); err != nil {
		return err
	}

	for _, handler := range s.handlers {
		go handler()
	}

	// Block until the context is over
	select {
	case <-s.ctx.Done():
	}

	if err := s.AfterServe(); err != nil {
		return err
	}

	return nil
}

func (s *Server) As(i interface{}) bool {
	p, ok := i.(**Server)
	if !ok {
			return false
		}

	*p = s
	return true
}