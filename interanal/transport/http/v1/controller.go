package http

import (
	"github.com/cheeeasy2501/go-gateway/internal/service"
)

type Controller struct {
	s service.Services
}

func NewController(s service.Services) *Controller {
	return &Controller{
		s: s,
	}
}
