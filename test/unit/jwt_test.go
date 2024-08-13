package unit

import (
	"fmt"
	"testing"

	"github.com/TanmoyTSSaha/GoBase/internal/auth"
	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
	username := "tanmoy.saha"
	fullname := "Tanmoy Saha"

	tokenString, err := auth.GenerateJWTToken(username, fullname)
	if err != nil {
		fmt.Printf("ERROR ENCOUNTERED:: %v", err)
	}

	fmt.Printf("::::::::::TOKEN STRING GENERATED SUCCESSFULLY::::::::::\n\n%s\n\n", tokenString)
}

func TestVerifyToken(t *testing.T) {
	tokenString := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTcyMzU3MTgzOCwiaWF0IjoxNzIzNTcxNzc4LCJuYW1lIjoiVGFubW95IFNhaGEiLCJzdWIiOiJ0YW5tb3kuc2FoYSJ9.ecggGBg_-WrNEWdE80KxGKTGOw-tLoY5N0PNjEK1Z2Wyv01odpY1WyhvUi8vIfiwqoMMrRuoVsn97mcn-aLPXwL-OLZrgoP2d47h4b7XO7XVlsLS8zdif6vzeaCLzl-0vtTMdM8UQp1CV5hfsa9x5eI7RtKspjGmhLKdWILuTZ3bQgoq3NaM-W7mcIFH2yKhYvmloRTFztp27Ntn3ejOIKPS480wIRs-ehbeYYfF4othrXe5f5btl1LkwxkJY_cVN80wG7P80eNPLw33rUWGD_RkFWmeDpkgZEEuv9vpgfgJKXU6rBRu4uapt0z4YLIz8e9gqNMZtTyY9_8_MqRBAg`

	jwtToken, err := auth.VerifyJWTToken(tokenString)
	if err != nil {
		fmt.Printf("ERROR ENCOUNTERED:: %v", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok && !jwtToken.Valid {
		fmt.Printf("cannot verify the token! please give a proper token!")
	}

	fmt.Printf("::::::::::TOKEN VERIFIED SUCCESSFULLY::::::::::\n\n%s\n\n%s\t\t%s\n\n", jwtToken.Raw, claims["sub"], claims["name"])
}
