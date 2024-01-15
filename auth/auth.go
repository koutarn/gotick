package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
)

const redirectURL = ""
var AccessToken = ""

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func redirectToAuthorizationPage(w http.ResponseWriter, r *http.Request) (err error){
	clientID, _ ,err := getEnv()
	if err != nil {
		return err
	}

    params := url.Values{}
    params.Add("client_id", clientID)
    params.Add("scope", "tasks:write tasks:read")
    params.Add("state", "YOUR_STATE") // 状態を固有の値に設定
    params.Add("redirect_uri", redirectURL)
    params.Add("response_type", "code")

    authorizationURL := "https://ticktick.com/oauth/authorize?" + params.Encode()
    http.Redirect(w, r, authorizationURL, http.StatusFound)
}

func exchangeCodeForToken(code string) (*TokenResponse, error) {
	clientID,clientSecret,err := getEnv()
	if err != nil {
		return nil,err
	}
	
    params := url.Values{}
    params.Add("client_id", clientID)
    params.Add("client_secret", clientSecret)
    params.Add("code", code)
    params.Add("grant_type", "authorization_code")
    params.Add("redirect_uri", redirectURL)

    req, err := http.NewRequest("POST", "https://ticktick.com/oauth/token", bytes.NewBufferString(params.Encode()))
    if err != nil {
        return nil, err
    }

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var tokenResponse TokenResponse
    if err := json.Unmarshal(body, &tokenResponse); err != nil {
        return nil, err
    }

    return &tokenResponse, nil
}

// 環境情報を読み込む(API Key and Secret)
func getEnv()(clientID string,clientSecret string,err error){
	clientID = os.Getenv("TICKTICK_API")
	if clientID == "" {
		return "","",errors.New("Ticktick API is not found")
	} 
	clientSecret = os.Getenv("TICKTICK_SECRET")
	if clientSecret == "" {
		return "","",errors.New("Ticktick Secret is not found")
	}

	return clientID,clientSecret,nil
}