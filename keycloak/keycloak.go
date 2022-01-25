package keycloak

import (
	"context"
	"github.com/Nerzal/gocloak/v10"
	"github.com/spf13/viper"
	"log"
)

type keycloak struct {
	client       gocloak.GoCloak
	Realm        string
	Host         string
	clientId     string
	clientSecret string
	ctx          context.Context
}

func new() *keycloak {

	keycloakHost := viper.GetString("keycloak.host")
	clientId := viper.GetString("keycloak.client-id")
	clientSecret := viper.GetString("keycloak.client-secret")
	realm := viper.GetString("keycloak.realm")

	return &keycloak{
		client:       gocloak.NewClient(keycloakHost),
		Realm:        realm,
		Host:         keycloakHost,
		clientId:     clientId,
		clientSecret: clientSecret,
		ctx:          context.Background(),
	}
}

func (kc keycloak) aquireToken() *gocloak.JWT {
	token, err := kc.client.LoginClient(kc.ctx, kc.clientId, kc.clientSecret, kc.Realm)
	if err != nil {
		log.Fatalf("Something wrong with login: %v", err)
	}

	return token
}

func KCUserViaDiscordID(discordID string) *gocloak.User {
	kc := new()

	users, err := kc.client.GetUsers(kc.ctx, kc.aquireToken().AccessToken, kc.Realm, gocloak.GetUsersParams{
		IDPUserID: gocloak.StringP(discordID),
		IDPAlias:  gocloak.StringP("discord"),
	})

	if err != nil {
		log.Fatalf("error finding users: %v", err)
	}

	return users[0]
}
