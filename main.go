package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jwt <your-jwt-token>")
		return
	}

	token := os.Args[1]

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid JWT format")
		return
	}

	fmt.Println("--- HEADER ---")
	if header, err := decode(parts[0]); err != nil {
		fmt.Printf("Invalid header (%s)\n", err)
	} else {
		fmt.Println(header)
	}

	fmt.Println("--- PAYLOAD ---")
	if payload, err := decode(parts[1]); err != nil {
		fmt.Printf("Invalid payload (%s)\n", err)
	} else {
		fmt.Println(payload)
	}
}

func decode(encodedJson string) (string, error) {
	decodedBytes, err := base64.RawURLEncoding.DecodeString(encodedJson)
	if err != nil {
		return "", err
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(decodedBytes, &obj); err != nil {
		return "", err
	}

	formatted, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}
	return string(formatted), nil
}
