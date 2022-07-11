package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

var (
	plainPassword = []byte("qwerty")
	salt          = []byte{0xd7, 0xc2, 0xf2, 0x51, 0xaa, 0x6a, 0x4e, 0x7b}
)

func PasswordMD5(plainPassword []byte) []byte {
	tmp := md5.Sum(plainPassword)
	return tmp[:]
}

func PasswordBcrypt(plainPassword []byte) []byte {
	passBcrypt, _ := bcrypt.GenerateFromPassword(plainPassword, 10)
	return passBcrypt
}

func PasswordPBKDF2(plainPassword []byte) []byte {
	return pbkdf2.Key(plainPassword, salt, 4096, 32, sha1.New)
}

func PasswordScrypt(plainPassword []byte) []byte {
	passScrypt, _ := scrypt.Key(plainPassword, salt, 1<<15, 8, 1, 32)
	return passScrypt
}

func PasswordArgon2(plainPassword []byte) []byte {
	return argon2.IDKey(plainPassword, salt, 1, 64*1024, 4, 32)
}

func main() {
	fmt.Printf("PasswordMD5: %x\n", PasswordMD5(plainPassword))
	fmt.Printf("PasswordBcrypt: %x\n", PasswordBcrypt(plainPassword))
	fmt.Printf("PasswordPBKDF2: %x\n", PasswordPBKDF2(plainPassword))
	fmt.Printf("PasswordScrypt: %x\n", PasswordScrypt(plainPassword))
	fmt.Printf("PasswordArgon2: %x\n", PasswordArgon2(plainPassword))
}
