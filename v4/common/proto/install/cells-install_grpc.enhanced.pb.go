// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: cells-install.proto

package install

import (
	context "context"
	fmt "fmt"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var (
	enhancedInstallServers = make(map[string]InstallEnhancedServer)
)

type NamedInstallServer interface {
	InstallServer
	Name() string
}
type InstallEnhancedServer map[string]NamedInstallServer

func (m InstallEnhancedServer) GetDefaults(ctx context.Context, r *GetDefaultsRequest) (*GetDefaultsResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method GetDefaults should have a context")
	}
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.GetDefaults(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaults not implemented")
}

func (m InstallEnhancedServer) Install(ctx context.Context, r *InstallRequest) (*InstallResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method Install should have a context")
	}
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.Install(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}

func (m InstallEnhancedServer) PerformCheck(ctx context.Context, r *PerformCheckRequest) (*PerformCheckResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method PerformCheck should have a context")
	}
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.PerformCheck(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method PerformCheck not implemented")
}
func (m InstallEnhancedServer) mustEmbedUnimplementedInstallServer() {}
func RegisterInstallEnhancedServer(s grpc.ServiceRegistrar, srv NamedInstallServer) {
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedInstallServers[addr]
	if !ok {
		m = InstallEnhancedServer{}
		enhancedInstallServers[addr] = m
		RegisterInstallServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterInstallEnhancedServer(s grpc.ServiceRegistrar, name string) {
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedInstallServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}