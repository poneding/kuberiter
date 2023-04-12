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

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/poneding/kuberiter/pkg/server"
	"github.com/poneding/kuberiter/pkg/server/routes"
)

func main() {
	apiServer := server.NewServer(":5000")

	apiServer.RegisterRoutes(
		routes.RegisterAPIV1Route,
	)

	go func() {
		if err := apiServer.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server start failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("server shutting down ...")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	if err := apiServer.Stop(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	<-ctx.Done()

	log.Println("server shutdown")
}
