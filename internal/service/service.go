package service

import "github.com/nokamoto/shogithink/pkg/api/v1alpha1/v1alpha1connect"

// Shogi implements the v1alpha1connect.ShogiServiceHandler interface.
type Shogi struct {
	v1alpha1connect.UnimplementedShogiServiceHandler
}

// New creates a new instance of the Shogi service handler.
func New() v1alpha1connect.ShogiServiceHandler {
	return &Shogi{}
}
