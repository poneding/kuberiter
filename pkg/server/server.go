/*
Copyright 2023 The Kuberiter Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRoute func(router gin.IRouter)

type Server struct {
	Addr   string
	router *gin.Engine
	*http.Server
}

func NewServer(addr string) *Server {
	router := gin.Default()
	return &Server{
		Addr:   addr,
		router: router,
		Server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (s *Server) RegisterRoutes(RegisterRoutes ...RegisterRoute) {
	for _, registerRoute := range RegisterRoutes {
		registerRoute(s.router)
	}
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
