package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main(){
	fmt.Println("Please enter url")
	var url string
	fmt.Scanln(&url)
	response, err := getTheResponse(url)
	if err != nil{
		fmt.Println("There some errors", err)
	}
	fmt.Println("This is your content", response)
}

func getTheResponse(url string) (string, error){
	response, err := http.Get(url)
	if err != nil{
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200{
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil{
		return "", err
	}
	content := doc.Find("body").Text()
	return content, nil
}
