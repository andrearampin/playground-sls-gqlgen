package main

import (
	"context"
	"net/http"

	graphhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/andrearampin/menyoo/internal/datastore"
	"github.com/andrearampin/menyoo/internal/graph/generated"
	"github.com/andrearampin/menyoo/internal/graph/resolvers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

const (
	_graphqlPath = "/graphql"
)

var lambdaRequestAdapter *httpadapter.HandlerAdapter

func init() {
	datastore, err := datastore.NewDB()
	if err != nil {
		panic(err)
	}

	srv := graphhandler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			Datastore: datastore,
		},
	}))
	http.Handle(_graphqlPath, srv)
	lambdaRequestAdapter = httpadapter.New(srv)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return lambdaRequestAdapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
