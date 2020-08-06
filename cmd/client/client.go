package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"micro_service/internal/app/client/model"
	"micro_service/internal/app/client/myendpoint"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /api/v1

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

var address string = "localhost:9000"

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var svc model.Service
	svc = model.ServiceStruct{}
	ep := myendpoint.MakePingEndpoint(svc)

	r := myendpoint.MakeHttpHandler(ctx, ep)
	go func() {
		log.Println("Http Server start at port", address)
		handler := r
		errChan <- http.ListenAndServe(address, handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	log.Println(<-errChan)
}
