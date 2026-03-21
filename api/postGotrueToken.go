package api

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/matthewhartstonge/argon2"
	"supadash/database"
	"time"
)

type GotrueToken struct {
	Email              string `json:"email"`
	Password           string `json:"password"`
	GotrueMetaSecurity struct {
		CaptchaToken string `json:"captcha_token"`
	} `json:"gotrue_meta_security"`
}

func (a *Api) postGotrueToken(c *gin.Context) {
	var body GotrueToken
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	account, err := a.queries.GetAccountByEmail(c.Request.Context(), body.Email)
	if err != nil {
		c.JSON(404, gin.H{"error": "Account not found"})
		return
	}

	if verified, err := argon2.VerifyEncoded([]byte(body.Password), []byte(account.PasswordHash)); err != nil || !verified {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	// 1 Hour Expiration for Access Token
	claims := jwt.RegisteredClaims{
		Issuer:    "supamanager.io",
		Subject:   account.GotrueID,
		Audience:  []string{"supamanager.io"},
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(1 * time.Hour)},
		NotBefore: &jwt.NumericDate{Time: time.Now()},
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// TODO: use a real secret from config
	signedJwt, err := token.SignedString([]byte(a.config.JwtSecret))
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// Generate secure refresh token
	refreshBytes := make([]byte, 32)
	if _, err := rand.Read(refreshBytes); err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate refresh token"})
		return
	}
	refreshToken := base64.URLEncoding.EncodeToString(refreshBytes)

	// Save refresh token to DB (valid for 30 days)
	_, err = a.queries.InsertRefreshToken(c.Request.Context(), database.InsertRefreshTokenParams{
		AccountID: account.ID,
		Token:     refreshToken,
		ExpiresAt: pgtype.Timestamptz{Time: time.Now().AddDate(0, 0, 30), Valid: true},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save refresh token"})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  signedJwt,
		"token_type":    "Bearer",
		"expires_in":    3600,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":    account.ID,
			"email": account.Email,
			"app_metadata": gin.H{
				"provider": "email",
			},
		},
	})
}
