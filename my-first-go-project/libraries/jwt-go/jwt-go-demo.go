package jwt_go

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Demo() {

	// 创建一个 JWT 的 Claims
	var key = []byte("console.xiaoduoai")

	claims := jwt.MapClaims{
		"username": "john.doe",
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // 设置 JWT 的过期时间为 1 小时
	}

	// 使用 HS256 算法创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对 JWT 进行签名
	// 这里的密钥可以是任意字符串，用于保护 JWT 的完整性
	// 在验证 JWT 时，需要使用相同的密钥进行签名验证
	// 注意：请不要将密钥硬编码在代码中，应使用安全的方式存储和获取密钥
	// 在生产环境中，建议使用环境变量或配置文件来管理密钥
	// 这里只是示例，为了方便演示，将密钥直接写在代码中
	signedToken, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Failed to sign JWT:", err)
		return
	}

	fmt.Println("Signed JWT:", signedToken)

	// 解析和验证 JWT
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		// 校验 JWT 使用的签名算法是否为 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		fmt.Println("Failed to parse JWT:", err)
		return
	}

	// 验证 JWT 的有效性
	if !parsedToken.Valid {
		fmt.Println("Invalid JWT")
		return
	}

	// 获取 JWT 中的声明
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid JWT claims")
		return
	}

	// 获取用户名声明
	username, ok := claims["username"].(string)
	if !ok {
		fmt.Println("Invalid username claim")
		return
	}

	fmt.Println("Valid JWT")
	fmt.Println("Username:", username)
}

func Demo1() {
	var key = []byte("console.xiaoduoai")
	const TokenTTL = time.Hour * 24 * 2

	type TokenUserInfo struct {
		TenantId   string `json:"tenant_id"`
		UserId     string `json:"user_id"`
		IsMainUser bool   `json:"is_main_user"`
	}

	type CustomClaims struct {
		TokenUserInfo
		jwt.StandardClaims
	}

	claims := CustomClaims{
		TokenUserInfo: TokenUserInfo{
			TenantId:   "q-6217343bd740a8402e35a196",
			UserId:     "q-6217343bd740a8402e35a196",
			IsMainUser: true,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			Issuer:    "console.xiaoduoai.com",
		},
	}

	jsonData, err := json.Marshal(claims)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonData))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)

	if err != nil {
		fmt.Println("Invalid claim: ", token)
		return
	}

	fmt.Println(signedToken)
}
