package utilities

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
