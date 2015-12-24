package main

import (
	"net/url"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	DefaultURL = "https://github.com"
	DefaultAPI = "https://api.github.com"
)

// GitHub (and Enterprise) API options.
type Client struct {
	// Repository name
	Owner string
	Repo string

	// API endpoint (Default: https://api.github.com)
	API *url.URL

	// GitHub API token.
	Token string
}

// NewClientOAuth returns a new GitHub API client with OAuth.
func (c *Client) newClientOAuth(token string) *github.Client {
	// OAuth2 client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	// GitHub API Client with OAuth
	client:= github.NewClient(tc)
	return client
}
