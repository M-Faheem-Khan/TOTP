package totp

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

type Totp struct {
	Id     int64
	Secret string
	TOTP   string
	URL    string
}

// GenerateTOTP calculates the 6-digit TOTP token for a given secret byte slice.
func GenerateTOTP(secret []byte, timeStep int64, digits int) (Totp, error) {
	var totp Totp

	// Step 1: Calculate the Time Counter (C)
	// We use the current Unix time and divide by the time step (usually 30 seconds)
	currentTime := time.Now().Unix()
	counter := currentTime / timeStep

	// Convert the 64-bit integer counter into an 8-byte big-endian byte array
	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(counter))

	// Step 2: Generate the HMAC-SHA1 Hash
	mac := hmac.New(sha1.New, secret)
	_, err := mac.Write(counterBytes)
	if err != nil {
		return totp, err
	}
	hash := mac.Sum(nil)

	// Step 3: Dynamic Truncation
	// Get the last 4 bits of the 20-byte hash to determine the offset (0 to 15)
	offset := hash[len(hash)-1] & 0x0F

	// Extract 4 bytes starting from the offset
	// Mask the first bit to prevent signed/unsigned confusion (& 0x7F)
	binaryCode := int32(hash[offset]&0x7F)<<24 |
		int32(hash[offset+1]&0xFF)<<16 |
		int32(hash[offset+2]&0xFF)<<8 |
		int32(hash[offset+3]&0xFF)

	// Compute modulo to get a 6-digit (or specified length) integer
	mod := int32(math.Pow10(digits))
	token := binaryCode % mod

	// Format with leading zeros if necessary (e.g., "023456")
	formatStr := fmt.Sprintf("%%0%dd", digits)
	totp.Secret = string(secret)
	totp.TOTP = fmt.Sprintf(formatStr, token)
	return totp, nil
}

// GenerateSecret creates a 16-character Base32 secret compatible with Google Authenticator.
// 16 characters in Base32 equates to exactly 10 bytes (80 bits) of entropy.
func GenerateSecret() (string, error) {
	// 10 bytes * 8 bits = 80 bits of entropy
	randomBytes := make([]byte, 10)

	// Use crypto/rand for cryptographically secure randomness
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode to Base32 string (No padding '=' per Google's preference)
	secret := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	return secret, nil
}
