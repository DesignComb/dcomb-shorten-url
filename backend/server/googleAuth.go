package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"main/config"
	"main/model"
	jwt "main/pkg/jwt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func access(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"url": oauthURL(),
	})
}

func oauthURL() string {
	u := "https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&response_type=code&scope=%s %s&redirect_uri=%s"

	return fmt.Sprintf(u, config.Val.GoogleClientID, "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email", config.Val.Host+"/validate/login")
}

func login(c *gin.Context) {
	code := c.Query("code")

	token, err := accessToken(code)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("accessToken error")
		c.Redirect(http.StatusFound, "/")
		return
	}

	googleUserId, email, name, picture, err := getGoogleUserInfo(token)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("getGoogleUserInfo error")
		c.Redirect(http.StatusFound, "/")
		return
	}

	// 不是使用者，new user
	user, _ := model.GetUser(googleUserId)
	if (user == model.User{}) {
		var user model.User
		user.GoogleUserId = googleUserId
		user.GoogleUserEmail = email
		user.GoogleUserName = name
		user.GoogleUserPicture = picture
		user, err = model.CreateUser(user)
		if err != nil {
			panic(err)
		}
		log.Infof("userid: %v created", user.ID)
	}
	user, _ = model.GetUser(googleUserId)

	// create jwt token
	jwtToken, err := jwt.GenerateToken(strconv.FormatUint(user.ID, 10), googleUserId, email, name, picture)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("GenerateToken error")

		c.Redirect(http.StatusFound, "/")
		return
	}
	//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Content-Type, Set-Cookie")
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Set-Cookie")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	fmt.Println(jwtToken)
	c.SetSameSite(4)
	c.SetCookie(jwt.Key, jwtToken, config.Val.JWTTokenLife, "", "http://localhost:8001", true, false)

	log.Infof("userid: %v logged in", user.ID)
}

func accessToken(code string) (token string, err error) {
	u := "https://www.googleapis.com/oauth2/v4/token"

	data := url.Values{
		"code":          {code},
		"client_id":     {config.Val.GoogleClientID},
		"client_secret": {config.Val.GoogleSecretKey},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {config.Val.Host + "/validate/login"},
	}
	body := strings.NewReader(data.Encode())

	resp, err := http.Post(u, "application/x-www-form-urlencoded", body)
	if err != nil {
		return token, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	token = gjson.GetBytes(b, "access_token").String()

	return token, nil
}

func getGoogleUserInfo(token string) (googleUserId, email, name, picture string, err error) {
	u := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?alt=json&access_token=%s", token)
	resp, err := http.Get(u)
	if err != nil {
		return googleUserId, email, name, picture, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return googleUserId, email, name, picture, err
	}

	googleUserId = gjson.GetBytes(body, "id").String()
	email = gjson.GetBytes(body, "email").String()
	name = gjson.GetBytes(body, "name").String()
	picture = gjson.GetBytes(body, "picture").String()

	return googleUserId, email, name, picture, nil
}
