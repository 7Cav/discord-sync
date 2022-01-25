package cavAPI

import (
	"context"
	"crypto/tls"
	"github.com/7cav/api/proto"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

type api struct {
	token string
}

func new() *api {
	token := viper.GetString("cav.api-token")

	return &api{
		token: token,
	}
}

func GetUserViaKCID(keycloakID *string) *proto.Profile {
	profile, err := new().client().GetUserViaKeycloakId(context.Background(), &proto.KeycloakIdRequest{KeycloakId: *keycloakID})
	if err != nil {
		panic(err)
	}
	return profile
}

func (api api) client() proto.MilpacServiceClient {
	rpcCreds := oauth.NewOauthAccess(&oauth2.Token{AccessToken: api.token})

	// use TLS config to auto detect SSL/TLS cert from Api Host
	config := &tls.Config{
		InsecureSkipVerify: false,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(config)),
		grpc.WithPerRPCCredentials(rpcCreds),
		grpc.WithBlock(),
	}

	conn, _ := grpc.Dial("api.7cav.us:443", opts...)
	return proto.NewMilpacServiceClient(conn)
}
