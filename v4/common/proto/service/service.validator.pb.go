// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package service

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Query) Validate() error {
	for _, item := range this.SubQueries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SubQueries", err)
			}
		}
	}
	if this.ResourcePolicyQuery != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResourcePolicyQuery); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResourcePolicyQuery", err)
		}
	}
	return nil
}
func (this *ResourcePolicyQuery) Validate() error {
	return nil
}
func (this *ResourcePolicy) Validate() error {
	return nil
}
func (this *StopEvent) Validate() error {
	return nil
}
func (this *StatusResponse) Validate() error {
	return nil
}
func (this *ChangesArchiveQuery) Validate() error {
	return nil
}