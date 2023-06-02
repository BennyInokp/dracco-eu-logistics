package utilities

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func GenerateUniqueID(baseString string) string {
	// Generate a random number
	randomNum := rand.Intn(10000)

	// Get the current timestamp
	timestamp := time.Now().Unix()

	// Concatenate the base string, timestamp, and random number
	uniqueID := fmt.Sprintf("%s-%d-%d", baseString, timestamp, randomNum)

	return uniqueID
}

// PrettyPrintStruct Prints structures with indents
func PrettyPrintStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func ValidateEmail(email string) bool {
	// Regular expression pattern for email validation
	// This pattern follows the RFC 5322 standard for email addresses
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Match the email against the pattern
	return regex.MatchString(email)
}

func RandomNumber() int {
	// Generate a random number
	number := rand.Intn(9999999999)
	return number
}
