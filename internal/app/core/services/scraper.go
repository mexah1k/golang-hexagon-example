package services

import (
	"github.com/gocolly/colly/v2"
)

type Scraper struct {
	collector *colly.Collector
}

func NewScraper() *Scraper {
	c := colly.NewCollector(
		colly.Async(true),
	)

	return &Scraper{collector: c}
}

func (s *Scraper) GetMetaTags(url string) ([]string, error) {
	var metaTags []string

	s.collector.OnHTML("meta", func(e *colly.HTMLElement) {
		metaTag := "<meta"
		for _, attr := range e.DOM.Nodes[0].Attr {
			metaTag += " " + attr.Key + "=\"" + attr.Val + "\""
		}
		metaTag += ">"
		metaTags = append(metaTags, metaTag)
	})

	err := s.collector.Visit(url)
	if err != nil {
		return nil, err
	}

	// Wait until the collector finishes
	s.collector.Wait()

	return metaTags, nil
}
