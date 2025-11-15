package handler

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var (
	githubOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("YOUR_GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("YOUR_GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("YOUR_GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("YOUR_GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
)

func GitHubLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := githubOauthConfig.AuthCodeURL("random-state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("random-state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
