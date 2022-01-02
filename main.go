package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		return
	}
}

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		opts...,
	)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	errTaskLinkedinLogin := chromedp.Run(ctx, LinkedinLogin())
	if errTaskLinkedinLogin != nil {
		log.Fatal(errTaskLinkedinLogin)
	}

	skill := "php"
	location := "hanoi"
	errTaskSearchProfile := chromedp.Run(ctx, SearchProfile(skill, location))
	if errTaskSearchProfile != nil {
		log.Fatal(errTaskSearchProfile)
	}
}

func LinkedinLogin() chromedp.Tasks {
	url := "https://www.linkedin.com"
	emailField := "//*[@id='session_key']"
	passwordField := "//*[@id='session_password']"
	submitButton := "//*[@id='main-content']/section[1]/div/div/form/button"

	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.SendKeys(emailField, ""),
		chromedp.SendKeys(passwordField, ""),
		chromedp.Submit(submitButton),
		chromedp.Sleep(3 * time.Second),
	}
}

func SearchProfile(skill, location string) chromedp.Tasks {
	url := fmt.Sprintf("https://www.google.com/search?q=site:linkedin.com/in/ AND '%s' AND '%s'", skill, location)
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(3 * time.Second),
	}
}
