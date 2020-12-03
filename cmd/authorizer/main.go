package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"errors"
)

func handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	//token := request.AuthorizationToken
	//tokenSlice := strings.Split(token, " ")
	//var bearerToken string
	//if len(tokenSlice) > 1 {
	//	bearerToken = tokenSlice[len(tokenSlice)-1]
	//}
	//
	//secret := os.Getenv("AUTHORIZER_SECRET")
	//if len(secret) == 0 || bearerToken != secret {
	//	return events.APIGatewayCustomAuthorizerResponse{}, errors.New("unauthorized")
	//}
	if request.AuthorizationToken == "" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("unauthorized")
	}

	return generatePolicy("user", "Allow", request.MethodArn), nil
}

func generatePolicy(principalID, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}
	return authResponse
}

func main() {
	lambda.Start(handler)
}
