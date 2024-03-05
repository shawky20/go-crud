package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"
)

var (
	gitLabClientID     = "eeece5a9f9291328a36b173dc3b626874281b08122017da8d541130d22f90747"
	gitLabClientSecret = "gloas-f12b27834ed01d022f0e50979e2ab4e0082681cf491cb4d9313290065f662ba7"
	gitLabRedirectURI  = "http://localhost:3000/gitlab/callback"
)

// HandleGitLabCallback handles the GitLab OAuth callback
func HandleGitLabCallback(c *gin.Context) {
	fmt.Println("Receiving GitLab OAuth callback")

	var requestData map[string]string
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, ok := requestData["code"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GitLab OAuth code not found in request"})
		return
	}

	token, err := exchangeGitLabCodeForToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a client using the access token
	client, err := gitlab.NewOAuthClient(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// List projects for the authenticated user
	projects, _, err := client.Projects.ListProjects(&gitlab.ListProjectsOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success, access token, and user projects
	c.JSON(http.StatusOK, gin.H{"success": true, "access_token": token.AccessToken, "projects": projects})
}

// exchangeGitLabCodeForToken exchanges the GitLab OAuth code for an access token
func exchangeGitLabCodeForToken(code string) (*oauth2.Token, error) {
	config := oauth2.Config{
		ClientID:     gitLabClientID,
		ClientSecret: gitLabClientSecret,
		RedirectURL:  gitLabRedirectURI,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://gitlab.com/oauth/authorize",
			TokenURL: "https://gitlab.com/oauth/token",
		},
		Scopes: []string{"read_api"},
	}

	return config.Exchange(oauth2.NoContext, code)
}
