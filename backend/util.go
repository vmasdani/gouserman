package main

import (
	"errors"
	"io/fs"
	"math/rand"
	"os"
	"strings"

	"github.com/samber/lo"
)

func GetJwtSecret() (string, error) {
	runes := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	jwtSecret, err := os.ReadFile("jwt.secret")

	if err != nil {
		return "", errors.New("jwt.secret file does not exist")
	}

	if len(string(jwtSecret)) > 0 {
		return string(jwtSecret), nil
	}

	r := lo.Range(32)

	r1 := lo.Map(r, func(n int, i int) string {
		return string(runes[rand.Intn(len(runes))])
	})
	r = nil

	r2 := strings.Join(r1, "")
	r1 = nil

	os.WriteFile("jwt.secret", []byte(r2), fs.ModeAppend)

	return r2, nil
}
