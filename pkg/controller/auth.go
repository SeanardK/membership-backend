package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/SeanardK/web-profile/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

type loginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	baseURL := strings.TrimRight(utils.GetEnv("KEYCLOAK_BASE_URL", ""), "/")
	realm := utils.GetEnv("KEYCLOAK_REALM", "")
	clientID := utils.GetEnv("CLIENT_ID", "")
	clientSecret := utils.GetEnv("CLIENT_SECRET", "")

	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", baseURL, realm)

	formData := url.Values{
		"grant_type": {"password"},
		"client_id":  {clientID},
		"username":   {req.Username},
		"password":   {req.Password},
	}
	if clientSecret != "" {
		formData.Set("client_secret", clientSecret)
	}

	resp, err := http.PostForm(tokenURL, formData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to reach auth server"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to read auth response"})
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse auth response"})
		return
	}

	if resp.StatusCode != http.StatusOK {
		errDesc, _ := result["error_description"].(string)
		c.JSON(resp.StatusCode, gin.H{"message": errDesc})
		return
	}

	c.JSON(http.StatusOK, result)
}
