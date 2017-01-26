package scrape

import (
  "log"
  "github.com/PuerkitoBio/goquery"
  "../models"
)

func GetPrice(link string, selector string) string {
  document, err := goquery.NewDocument(link)
  if err != nil {
    log.Fatal(err)
  }

  price := document.Find(selector).First().Text()

  return price
}

func getText(document *goquery.Document, selector string) string {
  return document.Find(selector).First().Text()
}

func getImagesLink(document *goquery.Document, selector string) []string {
  var imagesLinks []string

  document.Find(selector).Each(func(i int, selection *goquery.Selection) {
    link, _ := selection.Find("img").Attr("src")
    imagesLinks = append(imagesLinks, link)
  })

  return imagesLinks
}

func GetProductInfo(link string, rule *models.Rule) models.Product {
  document, err := goquery.NewDocument(link)

  if err != nil {
    log.Fatal(err)
  }

  product := models.Product {}

  product.Title = getText(document, rule.Title)
  product.Price = getText(document, rule.Price)
  product.Images = getImagesLink(document, rule.Images)

  return product
}
