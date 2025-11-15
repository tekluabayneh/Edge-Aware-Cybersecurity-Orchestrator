package oauth

import (
	"fmt"
	"net/http"

	OAuth "github.com/edge-aware-cyberSecurity/internal/OAuthConfig"
	"golang.org/x/oauth2"
)

func GitHubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	token, err := OAuth.GithubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Now you can use token to fetch user info from GitHub API
	fmt.Fprintf(w, "Access token: %s", token.AccessToken)

	client := oauth2.NewClient(r.Context(), oauth2.StaticTokenSource(token))
	resp, _ := client.Get("https://api.github.com/user/emails")

	fmt.Println(resp)
	//TODO: parse JSON to get user email

}
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	token, err := OAuth.GithubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Access token: %s", token.AccessToken)
	client := oauth2.NewClient(r.Context(), oauth2.StaticTokenSource(token))
	resp, _ := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	fmt.Println(resp)
	//TODO: parse JSON to get user email
}
