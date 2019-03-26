package riotapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Object struct {
}

func Get(url string) (obj Object, err error) {
	// HTTPリクエスト
	res, err := http.Get(url)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return obj, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return obj, err
	}

	if err = json.Unmarshal(body, &obj); err != nil {
		log.Println(err)
		return obj, err
	}
	return obj, nil
}
