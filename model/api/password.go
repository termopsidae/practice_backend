package api

import (
	"crypto/md5"
	"encoding/hex"
	"paractice/config"
)

func Get256Pw(inputPassword string) string {
	hash := md5.New()
	hash.Write([]byte(inputPassword + config.Config("SALT")))
	return hex.EncodeToString(hash.Sum(nil))
}
