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

	fmt.Printf("::::::::::TOKEN STRING GENERATED SUCCESSFULLY::::::::::\n%s\n\n", tokenString)
}

func TestVerifyToken(t *testing.T) {
	tokenString := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTcyMzgxOTc0MywiaWF0IjoxNzIzNTYwNTQzLCJuYW1lIjoiVGFubW95IFNhaGEiLCJzdWIiOiJ0YW5tb3kuc2FoYSJ9.Hh68zIzgYC4mWFVKWOWGdRWVssu-cNNNYoCRxI-AhST0ae90mu9Vapz7ZUHuTPesbghMloCBkC043LsNHobNuT9CAOC54hwI0YRaHz6j_XsUZmDkZKiOSfSy1HakpU1FOfZZj7Qp7rGaD4eF9FAiHWuFDI_XoG2zCRyXjpQeTPfguzeoqBJPqzmQhJDdCD2xLBTYS2HvopFNL0ETOLVO5GpFzpnqww4gm1X6yLllYz1fLajGqg3uQ9Iewo4DJB40bHKFjW5JLbNyV3Jb7U4sAggrTfODUpIEe0FGBgtyu2CI5k9IKDZ6k9vIkb4Ea9_AbsvDeg9Z1cKRG3jZ6pPWEA`

	jwtToken, err := auth.VerifyJWTToken(tokenString)
	if err != nil {
		fmt.Printf("ERROR ENCOUNTERED:: %v", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok && !jwtToken.Valid {
		fmt.Printf("cannot verify the token! please give a proper token!")
	}

	fmt.Printf("::::::::::TOKEN VERIFIED SUCCESSFULLY::::::::::\n%s\n\n%s\t\t%s\n\n", jwtToken.Raw, claims["sub"], claims["name"])
}
