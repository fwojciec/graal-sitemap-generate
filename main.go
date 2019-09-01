package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	siteURL        = "https://graalagency.com"
	endpoint       = "http://localhost:4000/query"
	query          = `{ "query": "{ clients: clients { slug } authors: authors { slug } }" }`
	fileName       = "./sitemap.xml"
	includeAuthors = false
)

func main() {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = writeSitemap(f)
	if err != nil {
		log.Fatal(err)
	}
}

func makeStaticPages(includeAuthors bool) []string {
	var staticPages []string
	if includeAuthors {
		staticPages = []string{"/", "/authors", "/clients", "/mailing-list", "/about-us", "/contact"}
	} else {
		staticPages = []string{"/", "/clients", "/mailing-list", "/about-us", "/contact"}
	}
	return staticPages
}

func buildStatic(staticPages []string) []*url {
	urls := make([]*url, 0)
	for _, page := range staticPages {
		locEn := strings.TrimRight(fmt.Sprintf("%s/en%s", siteURL, page), "/")
		locPl := strings.TrimRight(fmt.Sprintf("%s/pl%s", siteURL, page), "/")
		links := []*link{
			&link{Rel: "alternate", HrefLang: "en", Href: locEn},
			&link{Rel: "alternate", HrefLang: "pl", Href: locPl},
		}
		urlEN := &url{Loc: locEn, ChangeFreq: "monthly", Priority: 0.5, Links: links}
		urlPL := &url{Loc: locPl, ChangeFreq: "monthly", Priority: 0.5, Links: links}
		urls = append(urls, urlEN, urlPL)
	}
	return urls
}

func buildDynamic(prefix string, slugs []string) []*url {
	urls := make([]*url, 0)
	for _, slug := range slugs {
		locEn := fmt.Sprintf("%s/en/%s/%s", siteURL, prefix, slug)
		locPl := fmt.Sprintf("%s/pl/%s/%s", siteURL, prefix, slug)
		links := []*link{
			&link{Rel: "alternate", HrefLang: "en", Href: locEn},
			&link{Rel: "alternate", HrefLang: "pl", Href: locPl},
		}
		urlEN := &url{Loc: locEn, ChangeFreq: "weekly", Priority: 0.7, Links: links}
		urlPL := &url{Loc: locPl, ChangeFreq: "weekly", Priority: 0.7, Links: links}
		urls = append(urls, urlEN, urlPL)
	}
	return urls
}

func writeSitemap(w io.Writer) error {
	as, cs, err := getSlugs()
	if err != nil {
		return err
	}
	sp := makeStaticPages(len(as) > 0 && includeAuthors)
	us := newURLSet(
		buildStatic(sp),
		buildDynamic("clients", cs),
	)
	if includeAuthors {
		us.URLs = append(us.URLs, buildDynamic("authors", as)...)
	}
	output, err := xml.MarshalIndent(&us, "  ", "    ")
	if err != nil {
		return err
	}
	w.Write([]byte(xml.Header))
	w.Write(output)
	w.Write([]byte("\n"))
	return nil
}
