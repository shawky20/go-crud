package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

var (
	clientID     = "a941851b22ce289687b1"
	clientSecret = "5a17c0bbf9456b40023ebb87267e8dc3d98d6789"
	redirectURI  = "http://localhost:3000/github/callback"
)


func HandleGitHubCallback(c *gin.Context) {
	fmt.Println("recieviing ")
	var requestData map[string]string	
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, ok := requestData["code"]
	if !ok {	
		c.JSON(http.StatusBadRequest, gin.H{"error": "GitHub OAuth code not found in request"})
		return
	}

	token, err := exchangeGitHubCodeForToken(code)
	fmt.Println(token);
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a client using the access token
	client := github.NewClient(oauth2.NewClient(oauth2.NoContext, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})))

	// List repositories for the authenticated user
	repos, _, err := client.Repositories.List(oauth2.NoContext, "", nil)
	if err != nil {
		return 
	}

	// Respond with success or redirect the user to a success page
	c.JSON(http.StatusOK, gin.H{"success": true, "code": code, "repos": repos})

}


func exchangeGitHubCodeForToken(code string) (string, error) {
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURI,
		Endpoint:     oauth2.Endpoint{TokenURL: "https://github.com/login/oauth/access_token"},
	}

	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

