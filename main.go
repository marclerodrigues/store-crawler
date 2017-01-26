package main

import (
  "fmt"
  "./scrape"
  "./models"
)

func main() {
  link := "http://www.ricardoeletro.com.br/Produto/Celular-Smartphone-Alcatel-Pixi4-5-Preto-OT5045-Dual-Chip-4G-Tela-5-Camera-8MP-frontal-5MP-com-flash-Quad-Core-8GB-1GB-RAM-Android-60/44-491-516-607761?ch_pagetype=home&ch_feature=mostpopular"

  rule := &models.Rule {
    Title: "#ProdutoDetalhesNomeProduto",
    Price: "#ProdutoDetalhesPrecoComprarAgoraPrecoDePreco",
    Images: "#ProdutoDetalhesFotos" }

  productInfo := scrape.GetProductInfo(link, rule)

  fmt.Println(productInfo)
}
