package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {

	if os.Getenv("PLUGIN_USER") == "" {
		fmt.Println("user must be supplied")
		os.Exit(1)
	}
	if os.Getenv("PLUGIN_TOKEN") == "" {
		fmt.Println("token must be supplied")
		os.Exit(1)
	}
	if os.Getenv("PLUGIN_TITLE") == "" {
		fmt.Println("title must be supplied")
		os.Exit(1)
	}
	if os.Getenv("PLUGIN_MESSAGE") == "" {
		fmt.Println("message must be supplied")
		os.Exit(1)
	}

	resp, err := http.PostForm(
		"https://api.pushover.net/1/messages.json",
		url.Values{
			"user":    {os.Getenv("PLUGIN_USER")},
			"token":   {os.Getenv("PLUGIN_TOKEN")},
			"title":   {os.Getenv("PLUGIN_TITLE")},
			"message": {os.Getenv("PLUGIN_MESSAGE")}})

	/*
		write proper formatting
	*/

	fmt.Println(resp.Status)

	if err != nil {
		os.Exit(2)
	}
}
