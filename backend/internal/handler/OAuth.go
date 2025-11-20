package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	OAuth "github.com/edge-aware-cyberSecurity/internal/OAuthConfig"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/oauth2"
)

type GoogleUserInfo struct {
	Email string      `json:"email"`
	Name  string      `json:"name"`
	Photo pgtype.Text `json:"picture"`
}

type GithubUserInfo struct {
	Login string      `json:"login"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Photo pgtype.Text `json:"avatar_url"`
}

type OAuthHandler struct {
	DB *db.Queries
}

func (h *OAuthHandler) GitHubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	ctx := r.Context()
	token, err := OAuth.GithubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// fetch user info
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "token "+token.AccessToken)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	var githubUser struct {
		Login  string      `json:"login"`
		Name   string      `json:"name"`
		Avatar pgtype.Text `json:"avatar_url"`
	}
	json.NewDecoder(resp.Body).Decode(&githubUser)

	// fetch emails
	reqEmails, _ := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	reqEmails.Header.Set("Authorization", "token "+token.AccessToken)
	respEmails, _ := client.Do(reqEmails)
	defer respEmails.Body.Close()

	var emails []struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
	json.NewDecoder(respEmails.Body).Decode(&emails)
	primaryEmail := ""
	for _, e := range emails {
		if e.Primary && e.Verified {
			primaryEmail = e.Email
			break
		}
	}

	user, err := h.DB.GetUserByEmail(ctx, primaryEmail)
	newUser := db.CreateUserParams{
		Name:  githubUser.Name,
		Email: primaryEmail,
		Photo: githubUser.Avatar,
	}

	// redirect to the dashboard if user exists
	if user.Email != "" {
		utils.WriteJSON(w, http.StatusOK, map[string]string{
			"message": "user login successfully",
		})
		return
	}

	if errors.Is(err, sql.ErrNoRows) && user.Email == "" {
		err := h.DB.CreateUser(ctx, newUser)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		utils.WriteJSON(w, http.StatusOK, map[string]string{
			"message": "user registered successfully",
		})

	}

}

func (h *OAuthHandler) GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	ctx := r.Context()
	token, err := OAuth.GoogleOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	client := oauth2.NewClient(r.Context(), oauth2.StaticTokenSource(token))
	resp, _ := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	defer resp.Body.Close()

	var GoogleUserInfo GoogleUserInfo
	json.NewDecoder(resp.Body).Decode(&GoogleUserInfo)
	user, err := h.DB.GetUserByEmail(ctx, GoogleUserInfo.Email)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	// redirect to the dashboard if user exists
	if user.Email != "" {
		utils.WriteJSON(w, http.StatusOK, map[string]string{
			"message": "user login successfully",
		})
		return
	}

	if errors.Is(err, sql.ErrNoRows) && user.Email == "" {
		newUser := db.CreateUserParams{
			Name:  GoogleUserInfo.Name,
			Email: GoogleUserInfo.Email,
			Photo: GoogleUserInfo.Photo,
		}

		// register new user
		err := h.DB.CreateUser(ctx, newUser)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		utils.WriteJSON(w, http.StatusOK, map[string]string{
			"message": "user registered successfully",
		})

	}
}
