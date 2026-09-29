package curl

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func conn(req *http.Request) (string, error) {
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-CC-Version", "2018-03-22")
	req.Header.Add("X-CC-Api-Key", "0063512a-ea50-453e-abd3-347e7ab23e3b")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Post(url string, queryParams string) (string, error) {
	params := strings.NewReader(queryParams)

	req, err := http.NewRequest("POST", url, params)
	if err != nil {
		return "", err
	}

	connRes, connErr := conn(req)
	if connErr != nil {
		return "", connErr
	}

	return connRes, nil
}

func Get(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	connRes, connErr := conn(req)
	if connErr != nil {
		return "", connErr
	}

	return connRes, nil
}
