package main

import (
	"encoding/xml"
)

type urlset struct {
	Xmlns      string `xml:"xmlns,attr"`
	XmlnsXhtml string `xml:"xmlns:xhtml,attr"`
	URLs       []*url `xml:"url"`
}

type url struct {
	Loc        string  `xml:"loc"`
	ChangeFreq string  `xml:"changefreq"`
	Priority   float32 `xml:"priority"`
	Links      []*link `xml:"xhtml:link"`
}

type link struct {
	XMLName  xml.Name `xml:"xhtml:link"`
	Rel      string   `xml:"rel,attr"`
	HrefLang string   `xml:"hreflang,attr"`
	Href     string   `xml:"href,attr"`
}

func newURLSet(urlPairs ...[]*url) *urlset {
	urlset := &urlset{
		Xmlns:      "http://www.sitemaps.org/schemas/sitemap/0.9",
		XmlnsXhtml: "http://www.w3.org/1999/xhtml",
	}
	for _, urlPair := range urlPairs {
		urlset.URLs = append(urlset.URLs, urlPair...)
	}
	return urlset
}
