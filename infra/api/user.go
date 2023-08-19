package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/infra/api/parameter"
	"github.com/game-connect/gc-server/infra/api/output"
	"github.com/game-connect/gc-server/infra/api/response"
)

func LoginUser(user *parameter.LoginUser) (*model.User, error) {
	request, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/login_user", os.Getenv("GC_AUTH_URL")), bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &response.Success{
		Items: &output.LoginUser{},
	}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	userResult := &model.User{}
	if loginUser, ok := response.Items.(*output.LoginUser); ok {
		userResult.UserKey = loginUser.UserKey
		userResult.Name = loginUser.Name
		userResult.Email = loginUser.Email
		userResult.Token = loginUser.Token
		userResult.ImagePath = loginUser.ImagePath
	}

	return userResult, nil
}
