package jwt

import (
	"douyin_service/pkg/errcode"
	"douyin_service/pkg/util"
	"douyin_service/services/user/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(config.JwtCfg.Secret)
}

func GenerateToken(aud string) (string, error) {
	appKey := config.JwtCfg.Key
	appSecret := config.JwtCfg.Secret
	nowTime := time.Now()
	expireTime := nowTime.Add(config.JwtCfg.Expire * time.Second)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.JwtCfg.Issuer,
			Audience:  aud, // token的受众
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// ParseToken 解析Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func CheckToken(token string, userId string) (bool, *errcode.Error) {
	var (
		ecode = errcode.Success
	)

	if token == "" {
		ecode = errcode.InvalidParams
	} else {
		claim, err := ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				ecode = errcode.UnauthorizedTokenTimeout
			default:
				ecode = errcode.UnauthorizedTokenError
			}
		} else if userId != errcode.SkipCheckUserID && claim.Audience != userId { // 防止使用别人的Token
			ecode = errcode.IllegalToken
		}
	}

	if ecode != errcode.Success {
		return false, ecode
	}
	return true, ecode
}
