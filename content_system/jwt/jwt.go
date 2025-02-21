package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 加密的密钥
var key = "abcdefg123456"

type MyClaims struct {
	Username string `json:"user_name"` //自定义Payload有效载荷字段
	//Age      string `json:"age"`
	jwt.RegisteredClaims //提供标准验证功能,固定写法
}

// 生成token,生成签名，生成jwt
func SetToken(username string) (string, error) {
	SetClaims := MyClaims{
		Username: username,
		//Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), //有效时间(过期时间)，持续2个小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                    //生效时间，立即生效
			//Issuer:    os.Getenv("JWT_ISSUER"),                            //签发人
			//Subject:   "somebody",                                         //主题
			//ID:        "1",                                                //JWT ID用于标识该JWT
			//Audience:  []string{"somebody_else"},                          //用户
		},
	}
	//使用指定的加密方式(HS256签名算法)，和声明类型创建新令牌
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	//获得完整的、签名的令牌
	token, err := tokenStruct.SignedString([]byte(key)) //加密的密钥
	if err != nil {
		return "", err
	}
	return token, nil
}

// 验证token，验签，验证签名，解析jwt
func CheckToken(token string) (*MyClaims, error) {
	//解析、验证并返回token。
	tokenObj, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil //key 为自定义的密钥
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenObj.Claims.(*MyClaims); ok && tokenObj.Valid {
		fmt.Printf("%v %v\n", claims.Username, claims.RegisteredClaims)
		return claims, nil
	} else {
		return nil, err
	}
}
