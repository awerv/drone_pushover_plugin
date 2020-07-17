package main

import (
  "net/http"
  "net/url"
  "fmt"
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
	if os.Getenv("PLUGIN_MESSGAE") == "" {
		fmt.Println("message must be supplied")
		os.Exit(1)
	}


	resp, err := http.PostForm(
	  "https://api.pushover.net/1/messages.json",
	  url.Values{
		"user": { os.Getenv("PLUGIN_USER") },
		"token": { os.Getenv("PLUGIN_TOKEN") },
		"title": { os.Getenv("PLUGIN_TITLE") },
		"message": { os.Getenv("PLUGIN_MESSAGE") } })

	/*
		write proper formatting
	*/

	fmt.Println("[%s] %s", resp.Status, resp.Body)

	if err != nil {
		os.Exit(2)
	}
}