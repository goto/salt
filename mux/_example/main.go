package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/goto/salt/common"
	"github.com/goto/salt/mux"
	commonv1 "github.com/goto/salt/server/example/proto/gotocompany/common/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	grpcServer := grpc.NewServer()
	grpcGateway := runtime.NewServeMux()

	commonSvc := common.New(&commonv1.Version{})
	grpcServer.RegisterService(&commonv1.CommonService_ServiceDesc, commonSvc)
	if err := commonv1.RegisterCommonServiceHandlerServer(ctx, grpcGateway, commonSvc); err != nil {
		panic(err)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/api/", http.StripPrefix("/api", grpcGateway))
	httpMux.Handle("/ping", http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		for i := 0; i < 5; i++ {
			log.Printf("dooing stuff")
			time.Sleep(1 * time.Second)
		}
	}))

	log.Fatalf("server exited: %v", mux.Serve(ctx, "localhost:8080",
		mux.WithHTTP(httpMux),
		mux.WithGRPC(grpcServer),
		mux.WithGracePeriod(5*time.Second),
	))
}
