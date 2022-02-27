package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type keycloak struct {
	client       gocloak.GoCloak
	Realm        string
	Host         string
	clientId     string
	clientSecret string
	ctx          context.Context
}

func newKc() *keycloak {

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

func (kc keycloak) acquireToken() *gocloak.JWT {
	token, err := kc.client.LoginClient(kc.ctx, kc.clientId, kc.clientSecret, kc.Realm)
	if err != nil {
		log.Errorf("Something wrong with login: %v", err)
	}

	return token
}

func KCUserViaDiscordID(discordID string) (*gocloak.User, error) {
	kc := newKc()

	users, err := kc.client.GetUsers(kc.ctx, kc.acquireToken().AccessToken, kc.Realm, gocloak.GetUsersParams{
		IDPUserID: gocloak.StringP(discordID),
		IDPAlias:  gocloak.StringP("discord"),
	})

	if err != nil {

		log.Errorf("error finding users: %v", err)
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}

	return users[0], err
}
