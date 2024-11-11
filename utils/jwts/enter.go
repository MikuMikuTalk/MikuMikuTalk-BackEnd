package jwts

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtPayload struct {
	UserID   uint   `json:"userID"`
	Nickname string `json:"nickname"`
	Role     int8   `json:"role"` // 1. 管理员 2.普通用户
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims
}

// 构建GTI，为唯一值
func generateJTI() string {
	return uuid.NewString()
}

func ExtractJTI(claims *CustomClaims) string {
	return claims.ID
}

func GenerateJwtToken(payload JwtPayload, accessSecret string, expires int) (string, error) {
	cliams := CustomClaims{
		JwtPayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        generateJTI(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	return token.SignedString([]byte(accessSecret))
}

func ParseToken(jwtToken string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("非法jwt token")
}

func ValidateToken(jwtToken string, accessSecret string, redisClient *redis.Client) (bool, error) {
	claims, err := ParseToken(jwtToken, accessSecret)
	if err != nil {
		return false, err // Token 无效
	}
	// 提取 JTI
	jti := ExtractJTI(claims)
	key := "logout_" + jti
	// 如果已经存在在redis中，那么说明这个jwt token已经是被注销了，不能再用了
	exists, err := redisClient.Exists(key).Result()
	if err != nil {
		return false, errors.New("Redis 校验失败")
	}
	// 如果 Redis 中存在 JTI 标记，则认为 Token 已注销
	if exists > 0 {
		return false, errors.New("Token 已被注销")
	}
	// 校验通过
	return true, nil
}
