package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandlers func()) (context.Context, error) {
	registerHandlers()
	ctx = startService(ctx, serviceName, host, port)
	return ctx, nil								
}

func startService(ctx context.Context, serviceName, host, port string) context.Context{
	ctx,cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port
	go func(){
		err := srv.ListenAndServe()
		if err != nil{
			log.Printf("%s failed: %v", serviceName, err)
			cancel()
		}
	}()

	go func(){
		fmt.Printf("%s started. Press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}