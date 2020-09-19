package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/db"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/models"

	"github.com/graphql-go/graphql"
	"github.com/sharmarajdaksh/authenticated-gql-api-gateway/config"
)

var githubAuthMutationInputType = graphql.FieldConfigArgument{
	"code": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var githubAuthMutation = &graphql.Field{
	Type:        graphql.String,
	Description: "Sign up or Log in user via Github OAuth",
	Args:        githubAuthMutationInputType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		code, ok := p.Args["code"].(string)
		if !ok {
			return nil, errors.New("Could not parse required argument")
		}

		loginOrSignupGithubUser(code)

		return nil, nil
	},
}

func getGithubAccessToken(code string) (string, error) {

	reqMap := map[string]string{
		"client_id":     config.C.Auth.Github.GithubClientID,
		"client_secret": config.C.Auth.Github.GithubClientSecret,
		"code":          code,
	}
	reqJSON, e := json.Marshal(reqMap)
	if e != nil {
		return "", errors.New("Could not parse required argument")
	}

	req, e := http.NewRequest("POST", config.C.Auth.Github.GithubOauthTokenURL, bytes.NewBuffer(reqJSON))
	if e != nil {
		return "", errors.New("Failed to obtain access token from Github")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, e := http.DefaultClient.Do(req)
	if e != nil {
		return "", errors.New("Failed to obtain access token from Github")
	}

	resB, _ := ioutil.ReadAll(res.Body)
	if e != nil {
		return "", errors.New("Failed to parse response from Github")
	}

	var resData struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	e = json.Unmarshal(resB, &resData)
	if e != nil {
		return "", errors.New("Failed to parse response from Github")
	}

	return resData.AccessToken, nil
}

func getGithubUserdataB(accessToken string) ([]byte, error) {

	req, e := http.NewRequest("GET", config.C.Auth.Github.GithubOauthUserdataURL, nil)
	if e != nil {
		return nil, errors.New("Failed to fetch user data from Github")
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	res, e := http.DefaultClient.Do(req)
	if e != nil {
		return nil, errors.New("Failed to fetch user data from Github")
	}

	resB, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return nil, errors.New("Failed to parse response from Github")
	}

	return resB, nil
}

func handleUserManagement(userdataB []byte) (interface{}, error) {

	var userdata struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
	}
	if e := json.Unmarshal(userdataB, &userdata); e != nil {
		return "", errors.New("Failed to parse response from Github")
	}

	var baseUser models.User
	if e := db.DB.FirstOrCreate(&baseUser, models.User{Email: userdata.Email}); e.Error != nil {
		log.Println(e.Error)
		return nil, e.Error
	}

	githubUser := models.GithubUser{
		GithubID:  userdata.ID,
		Name:      userdata.Name,
		AvatarURL: userdata.AvatarURL,
		User:      baseUser,
	}

	if e := db.DB.FirstOrCreate(&githubUser, githubUser); e.Error != nil {
		log.Println(e.Error)
		return nil, e.Error
	}

	return &githubUser, nil
}

func generateAuthToken() {
	return
}

func loginOrSignupGithubUser(code string) (interface{}, error) {

	accessToken, e := getGithubAccessToken(code)
	if e != nil {
		return nil, e
	}

	userdataB, e := getGithubUserdataB(accessToken)
	if e != nil {
		return nil, e
	}

	user, e := handleUserManagement(userdataB)
	if e != nil {
		return nil, e
	}

	return user, nil
}

// GithubUserTable
// UserTable
// FK: GithubID
