package unit

import (
	"fmt"
	"testing"

	"github.com/TanmoyTSSaha/GoBase/internal/auth"
)

func TestGenerateHashFromPassword(t *testing.T) {
	passwordString := "Test@123"
	params := auth.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := auth.GenerateHashFromPassword(passwordString, &params)
	if err != nil {
		fmt.Printf("ERROR ENCOUNTERED:: %v", err)
	}

	fmt.Printf("::::::::::PRINTING HASH::::::::::\n\n%s\n\n", hash)
}


func TestVerifyHashPassword(t *testing.T) {
	passwordHash := `$argon2id$v=19$m=65536,t=3,p=2$Bf1fRV+7T3Kx+JyGsn2VvQ$iHgCCpirx2Ndxz9TbxwKI75Ke6b0W2Xi/RGPlJrWhzE`
	passwordString := "Test@123"

	hash, err := auth.VerifyHashPassword(passwordString, passwordHash)
	if err != nil {
		fmt.Printf("ERROR ENCOUNTERED:: %v\n\n", err)
	}

	if hash {
		fmt.Printf("\n::::::::::HASHED PASSWORD MATCHED SUCCESSFULLY::::::::::\n\n")
	} else {
		fmt.Printf("\n::::::::::HASHED PASSWORD DID NOT MATCHED::::::::::\n\n")
	}
}
