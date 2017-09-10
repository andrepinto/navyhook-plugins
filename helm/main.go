package main

import (
	"github.com/andrepinto/navyhook/plugins/base/proto"
	"github.com/andrepinto/navyhook/plugins/base/plug"
	"github.com/andrepinto/navyhook-plugins/helm/pkg"
)

const PLUGIN_NAME  = "helm-plugin"

type VendorPluginDemo struct{}

func (b VendorPluginDemo) Executor(req *proto.NavyHookVendorPluginRequest) (*proto.NavyHookVendorPluginResponse, error) {

	var err error

	switch req.Action {
	case "build":
		err := pkg.Build(req.Configs)
		return &proto.NavyHookVendorPluginResponse{},err
	default:
		return &proto.NavyHookVendorPluginResponse{},err

	}

	return &proto.NavyHookVendorPluginResponse{},err
}

func main() {
	plug.RunPlugin(&VendorPluginDemo{})
}


