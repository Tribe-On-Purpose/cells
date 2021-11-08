package service

import (
	"context"
	"crypto/tls"

	"github.com/google/uuid"
)

// ServiceOptions stores all options for a pydio service
type ServiceOptions struct {
	Name string
	ID   string
	Tags []string

	Version     string
	Description string
	Source      string

	Context context.Context
	Cancel  context.CancelFunc

	// DAO        func(dao.DAO) dao.DAO
	// Prefix     interface{}
	// Migrations []*Migration

	// Port      string
	TLSConfig *tls.Config

	// Micro micro.Service

	Dependencies []*dependency

	// Starting options
	AutoStart   bool
	AutoRestart bool
	Fork        bool
	Unique      bool

	// Registry registry.Registry

	// Regexp *regexp.Regexp
	// Flags  pflag.FlagSet

	// MinNumberOfNodes int

	// Before and After funcs
	// BeforeInit  []func(Service) error
	// AfterInit   []func(Service) error
	// BeforeStart []func(Service) error
	// AfterStart  []func(Service) error
	// BeforeStop  []func(Service) error
	// AfterStop   []func(Service) error

	// OnRegexpMatch func(Service, []string) error
}

type dependency struct {
	Name string
	Tag  []string
}

//
type ServiceOption func(*ServiceOptions)

// Name option for a service
func Name(n string) ServiceOption {
	return func(o *ServiceOptions) {
		o.Name = n
	}
}

// Tag option for a service
func Tag(t ...string) ServiceOption {
	return func(o *ServiceOptions) {
		o.Tags = append(o.Tags, t...)
	}
}

// Description option for a service
func Description(d string) ServiceOption {
	return func(o *ServiceOptions) {
		o.Description = d
	}
}

// Source option for a service
func Source(s string) ServiceOption {
	return func(o *ServiceOptions) {
		o.Source = s
	}
}

// Context option for a service
func Context(c context.Context) ServiceOption {
	return func(o *ServiceOptions) {
		o.Context = c
	}
}

// Cancel option for a service
func Cancel(c context.CancelFunc) ServiceOption {
	return func(o *ServiceOptions) {
		o.Cancel = c
	}
}

// WithTLSConfig option for a service
func WithTLSConfig(c *tls.Config) ServiceOption {
	return func(o *ServiceOptions) {
		o.TLSConfig = c
	}
}

// AutoStart option for a service
func AutoStart(b bool) ServiceOption {
	return func(o *ServiceOptions) {
		o.AutoStart = b
	}
}

// Fork option for a service
func Fork(b bool) ServiceOption {
	return func(o *ServiceOptions) {
		o.Fork = b
	}
}

// Unique option for a service
func Unique(b bool) ServiceOption {
	return func(o *ServiceOptions) {
		o.Unique = b
	}
}

// Dependency option for a service
func Dependency(n string, t []string) ServiceOption {
	return func(o *ServiceOptions) {
		o.Dependencies = append(o.Dependencies, &dependency{n, t})
	}
}

func newOptions(opts ...ServiceOption) ServiceOptions {
	opt := ServiceOptions{}

	opt.ID = uuid.New().String()

	for _, o := range opts {
		o(&opt)
	}

	return opt
}