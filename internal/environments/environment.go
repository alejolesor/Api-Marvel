package environments

import (
	"crypto/md5"
	"fmt"
)

var urlMarvel = "http://gateway.marvel.com/v1/"
var apikeyPublic = "4767bd441b9c524b7c9db29e8b2c3b16"
var apikeyPrivate = "edcae0e40a18c8d19b91454202a3b17aeaa5b049"
var ts = "1"
var hash string

//DigestString Genera hash para consumir api marvell
func DigestString(t string, publicKey string, privateKey string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(t+privateKey+publicKey)))
}

//GetURLMarvell retorna endpoint para consumir api ...
func GetURLMarvell(methodComics string) string {
	hash = DigestString(ts, apikeyPublic, apikeyPrivate)
	urlConcat := urlMarvel + methodComics + "apikey=" + apikeyPublic + "&" + "ts=" + ts + "&" + "hash=" + hash
	return urlConcat
}
