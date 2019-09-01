package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
)

type slugFields []*struct{ Slug string }

func (s slugFields) Slugs() []string {
	var slugs []string
	for _, sf := range s {
		slugs = append(slugs, sf.Slug)
	}
	sort.Strings(slugs)
	return slugs
}

func getSlugs() ([]string, []string, error) {
	resp, err := http.Post(endpoint, "application/json", bytes.NewBufferString(query))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	r := struct {
		Data struct {
			Clients slugFields
			Authors slugFields
		}
	}{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, nil, err
	}
	return r.Data.Authors.Slugs(), r.Data.Clients.Slugs(), nil
}
