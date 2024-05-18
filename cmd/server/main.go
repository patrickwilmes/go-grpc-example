/*
 * SPDX-License-Identifier: CC-BY-NC-ND-4.0
 *
 * go-grpc-example
 * Copyright (c) 2024 Patrick Wilmes <p.wilmes89@gmail.com>
 *
 * This file is licensed under the Creative Commons Attribution-NonCommercial-NoDerivatives 4.0 International License.
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://creativecommons.org/licenses/by-nc-nd/4.0/
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"github.com/patrickwilmes/go-grpc-example/generated/protocols"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GeometryServer struct {
	go_grpc_example.UnsafeGeometryServiceServer
}

func (s *GeometryServer) AreaOf(ctx context.Context, in *go_grpc_example.Rectangle) (*go_grpc_example.Area, error) {
	return &go_grpc_example.Area{Area: in.ASide * in.BSide}, nil
}

func newServer() *GeometryServer {
	return &GeometryServer{}
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	go_grpc_example.RegisterGeometryServiceServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
