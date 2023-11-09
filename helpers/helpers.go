package helpers

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// HashPwd hashes given password
func HashPwd(pwd string, salt string) string {
	cost := 7
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd+salt), cost)
	CheckError(err)
	return string(hashed)
}

// ComparePwd compares given password by user with stored password in the database
func ComparePwd(hashed string, pwd string, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd+salt))
	return err == nil
}

var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func SaltGenerator() string {
	// buat salt dengan panjang 5 karakter
	salt := make([]byte, 5)
	for i := range salt {
		// generate random number dari 0 sampai panjang charset
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		CheckError(err)
		// asign ke dalam array salt
		salt[i] = charset[idx.Int64()]
	}
	// konversi ke string
	return string(salt)
}
