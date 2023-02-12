package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/cheeeasy2501/auth-id"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	grpcConn *grpc.ClientConn
}

func NewHanler(grpcConn *grpc.ClientConn) *Handler {
	return &Handler{
		grpcConn: grpcConn,
	}
}

func (h *Handler) login(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("test")
}

func main() {
	// conn, err := grpc.Dial("localhost:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	//    log.Fatal(err)
	// }
	// log.Println("GRPC connection is open")
	// defer conn.Close()

	//  h := NewHanler(conn)
	//  http.HandleFunc("/login", h.login)
	// //  http.HandleFunc("/register", register)
	// //  http.HandleFunc("/check-token", checkToken)

	// err = http.ListenAndServe(":9000", nil); if err != nil {
	// 	log.Fatal(err)
	// }
}

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := auth_id_gen.RegisterGatewayHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
