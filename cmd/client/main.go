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
	go_grpc_example "github.com/patrickwilmes/go-grpc-example/generated/protocols"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := go_grpc_example.NewGeometryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := client.AreaOf(ctx, &go_grpc_example.Rectangle{
		ASide: 10,
		BSide: 10,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Area is %d", result.Area)
}
