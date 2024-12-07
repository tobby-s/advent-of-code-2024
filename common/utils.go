package common

import (
	"io"
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

func LoadData(url string) ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	cookie, err := LoadCookie()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie",cookie)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body , err := io.ReadAll(res.Body)
	strArr := strings.Split(string(body), "\n")
	return strArr, nil
}