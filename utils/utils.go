package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

)

func LoadCookie() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {return "", err}
	return os.Getenv("COOKIE"), nil
}

func LoadData(url string) ([]string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	cookie, err := LoadCookie()
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cookie",cookie)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body , err := io.ReadAll(res.Body)
	strArr := strings.Split(string(body), "\n")
	return strArr
}