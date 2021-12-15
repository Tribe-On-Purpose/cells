package mux

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	pb "github.com/pydio/cells/v4/common/proto/registry"

	caddy "github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/pydio/cells/v4/common/registry"
	servercontext "github.com/pydio/cells/v4/common/server/context"
)

func RegisterServerMux(ctx context.Context, s *http.ServeMux) {
	caddy.RegisterModule(Middleware{
		r: servercontext.GetRegistry(ctx),
		s: s,
	})
	httpcaddyfile.RegisterHandlerDirective("mux", parseCaddyfile)
}

type Middleware struct {
	r registry.Registry
	s *http.ServeMux
}

// CaddyModule returns the Caddy module information.
func (m Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.mux",
		New: func() caddy.Module { return &m },
	}
}

// Provision adds routes to the main server
func (m Middleware) Provision(ctx caddy.Context) error {
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// try to find it in the current mux
	_, pattern := m.s.Handler(r)
	if len(pattern) > 0 && (pattern != "/" || r.URL.Path == "/") {
		m.s.ServeHTTP(w, r)
		return nil
	}

	// Couldn't find it in the mux, we go through the registered endpoints
	nodes, err := m.r.List(registry.WithType(pb.ItemType_NODE))
	if err != nil {
		return err
	}

	for _, item := range nodes {
		node, ok := item.(registry.Node)
		if !ok {
			continue
		}
		for _, endpoint := range node.Endpoints() {
			if endpoint == "/" {
				continue
			}
			ok, err := regexp.Match(endpoint, []byte(r.URL.Path))
			if err != nil {
				return err
			}

			if ok {
				// TODO v4 - proxy should be set once when watching the node
				u, err := url.Parse("http://" + strings.Replace(node.Address()[0], "[::]", "", -1))
				if err != nil {
					return err
				}
				proxy := httputil.NewSingleHostReverseProxy(u)
				proxy.ServeHTTP(w, r)
				return nil
			}
		}
	}

	// no matching filter found
	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	/*for d.Next() {
		if !d.Args(&m.Output) {
			return d.ArgErr()
		}
	}*/
	return nil
}

func (m Middleware) WrapListener(ln net.Listener) net.Listener {
	fmt.Println("The address is ? ", ln.Addr())
	return ln
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}