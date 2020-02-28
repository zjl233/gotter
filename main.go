// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml
//
// The code under api/petstore/ has been generated from that specification.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent"
	"github.com/zjl233/gotter/ent/migrate"
	"github.com/zjl233/gotter/srv"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("start")

	// load env
	_ = godotenv.Load()

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Fatal("parse port error")
	}
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	//  ==============db================
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_DB"), os.Getenv("PG_PASSWORD"))
	fmt.Println("++++++++++++++ db address: ", connString, "++++++++++++++++++++++++++")
	client, err := ent.Open("postgres", connString)
	if err != nil {
		log.Fatalf("failed connecting to psql: %v", err)
	}
	//defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// =================db==================

	// Create an instance of our handler which satisfies the generated interface
	petStore := srv.NewPostSrv(client)

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	api.RegisterHandlers(e, petStore)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", port)))
}
