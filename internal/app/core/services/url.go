package services

import (
	"golang-hexagon-example/internal/app/core/models"
	"golang-hexagon-example/internal/app/core/ports"
	"log"
)

type UrlService struct {
	Scraper    ports.ScraperService
	Repository ports.UrlRepository
}

func NewUrlService(scraper ports.ScraperService, repository ports.UrlRepository) *UrlService {
	return &UrlService{
		Scraper:    scraper,
		Repository: repository,
	}
}

func (dcs *UrlService) Analyze(urls []string) error {
	for _, link := range urls {
		metaTags, err := dcs.Scraper.GetMetaTags(link)
		if err != nil {
			log.Printf("Error scraping %s: %v", link, err)
			continue
		}

		analysisResult := models.UrlAnalysisResult{
			Link:     link,
			MetaTags: metaTags,
		}

		if err = dcs.Repository.Create(analysisResult); err != nil {
			log.Printf("Error saving result for %s: %v", link, err)
			continue
		}
	}

	return nil
}
