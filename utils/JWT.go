package utils

import (
	"fmt"
	"getway/consts"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yanlihongaichila/framework/nacos"
	"gopkg.in/yaml.v2"
)

func getSecret() (string, error) {
	var conf struct {
		App struct {
			Secret string `yaml:"secret"`
		} `yaml:"app"`
	}
	cnfStr, err := nacos.GetConfig("DEFAULT_GROUP", consts.SERVICENAME)
	if err != nil {
		return "", err
	}
	err = yaml.Unmarshal([]byte(cnfStr), &conf)
	if err != nil {
		return "", err
	}
	return conf.App.Secret, nil
}

func GenToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	secret, err := getSecret()
	if err != nil {
		return "", err
	}

	return token.SignedString([]byte(secret))
}

func CheckToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		secret, err := getSecret()
		if err != nil {
			return nil, err
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	tokenMap := token.Claims.(jwt.MapClaims)
	fmt.Println("***********tokenMap")
	fmt.Println(tokenMap)
	if _, ok := tokenMap["user_id"]; !ok {
		return 0, fmt.Errorf("token invalid")
	}
	return int64(tokenMap["user_id"].(float64)), err
}
