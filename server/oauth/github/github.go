package github

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/utils"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GithubUserInfo struct {
	Login                   string      `json:"login"`
	ID                      int         `json:"id"`
	NodeID                  string      `json:"node_id"`
	AvatarURL               string      `json:"avatar_url"`
	GravatarID              string      `json:"gravatar_id"`
	URL                     string      `json:"url"`
	HTMLURL                 string      `json:"html_url"`
	FollowersURL            string      `json:"followers_url"`
	FollowingURL            string      `json:"following_url"`
	GistsURL                string      `json:"gists_url"`
	StarredURL              string      `json:"starred_url"`
	SubscriptionsURL        string      `json:"subscriptions_url"`
	OrganizationsURL        string      `json:"organizations_url"`
	ReposURL                string      `json:"repos_url"`
	EventsURL               string      `json:"events_url"`
	ReceivedEventsURL       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    string      `json:"name"`
	Company                 interface{} `json:"company"`
	Blog                    string      `json:"blog"`
	Location                interface{} `json:"location"`
	Email                   string      `json:"email"`
	Hireable                interface{} `json:"hireable"`
	Bio                     string      `json:"bio"`
	TwitterUsername         interface{} `json:"twitter_username"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}

type GithubOAuth struct {
	clientId     string
	clientSecret string
}

func NewGithubOAuth(clientId, clientSecret string) *GithubOAuth {
	return &GithubOAuth{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (g *GithubOAuth) GetLoginUrl(c *gin.Context) {
	response.OkWithData("https://github.com/login/oauth/authorize?client_id="+g.clientId, c)
}

func (g *GithubOAuth) GetGithubUserInfo(code string) (GithubUserInfo, error) {
	header := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	body := map[string]string{
		"client_id":     g.clientId,
		"client_secret": g.clientSecret,
		"code":          code,
	}
	marshal, _ := json.Marshal(body)
	respBody, err := utils.SendHTTPRequest("POST", "https://github.com/login/oauth/access_token", header, marshal)
	if err != nil {
		log.Println("err:", err)
		return GithubUserInfo{}, err
	}

	var githubToken Token
	_ = json.Unmarshal(respBody, &githubToken)
	accessToken := githubToken.AccessToken

	header = map[string]string{
		"Authorization": "bearer " + accessToken,
		"Content-Type":  "application/json",
	}
	respBody, err = utils.SendHTTPRequest("GET", "https://api.github.com/user", header, nil)
	if err != nil {
		return GithubUserInfo{}, err
	}
	var githubUser GithubUserInfo
	_ = json.Unmarshal(respBody, &githubUser)
	return githubUser, nil
}
