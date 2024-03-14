package internal

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"

	cache "course.project/authentication/internal/common/cache"
	dao "course.project/authentication/internal/dao"
	pb "course.project/authentication/proto"
)

var aeskey = []byte("8fvh9cok87bjefgh")

var jwtKey = []byte("9vjs9ifed")

type AuthenticationServerImpl struct{}

func (a *AuthenticationServerImpl) Login(ctx context.Context, userInfo *pb.UserInfo) (*pb.LoginStatus, error) {
	decoded, err := passwordDecode(userInfo.EncodedPassword)
	if err != nil {
		return nil, err
	}
	if !dao.IsValidUser(userInfo.UserName, decoded) {
		return &pb.LoginStatus{
			Status: int32(pb.Constant_FAILED),
		}, nil
	}
	cacheToken, err := cache.Ca.Get(ctx, userInfo.UserName)
	if err != nil {
		return nil, err
	}
	if len(cacheToken) > 0 {
		return &pb.LoginStatus{
			Status: int32(pb.Constant_ONLINE),
			LoginToken: &pb.LoginToken{
				Token: cacheToken,
			},
		}, nil
	}
	token, err := createToken(userInfo.UserName)
	if err != nil {
		return &pb.LoginStatus{
			Status: int32(pb.Constant_FAILED),
		}, err
	}
	err = cache.Ca.Set(ctx, userInfo.UserName, token, time.Minute*30)
	if err != nil {
		return &pb.LoginStatus{
			Status: int32(pb.Constant_FAILED),
		}, err
	}
	return &pb.LoginStatus{
		Status: int32(pb.Constant_NORMAL),
		LoginToken: &pb.LoginToken{
			Token: token,
		},
	}, nil
}

func passwordDecode(encoded string) (string, error) {
	bytesPass, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Println(err)
		return "", err
	}
	tpass, err := aesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(tpass), nil
}

func passwordEncode(password string) (string, error) {
	xpass, err := aesEncrypt([]byte(password), aeskey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	pass64 := base64.StdEncoding.EncodeToString(xpass)
	return pass64, nil
}

// decode helper
func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func aesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS5UnPadding(origData)
	return origData, nil
}

// encode helper
func pKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

func aesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

type UserJWT struct {
	Username string
	jwt.MapClaims
}

func createToken(username string) (string, error) {
	claims := UserJWT{
		username,
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatalf("login token creating error: %v\n", err)
		return "", err
	}

	return tokenString, nil
}

func (a *AuthenticationServerImpl) Logout(ctx context.Context, loginToken *pb.LoginToken) (*pb.LoginStatus, error) {
	username, err := parseToken(loginToken.Token)
	if err != nil {
		return nil, err
	}
	cache.Ca.Del(ctx, username)
	return &pb.LoginStatus{
		Status: int32(pb.Constant_NORMAL),
	}, nil
}

func parseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserJWT{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		log.Fatalf("login token parse error: %v\n", err)
		return "", err
	}
	claims, _ := token.Claims.(*UserJWT)
	return claims.Username, nil
}
