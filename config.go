package twitterbot

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Keys struct {
	consumerPublic, consumerSecret, accessPublic, accessSecret string
}

func ReadConfig(config string) (*Keys, error) {
	file, err := os.Open(config)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	keys := Keys{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, errors.New("Malformed config line: " + line)
		}
		keyType := strings.TrimSpace(parts[0])
		keyValue := strings.TrimSpace(parts[1])
		if keyType == "ConsumerSecret" {
			keys.consumerSecret = keyValue
		} else if keyType == "ConsumerPublic" {
			keys.consumerPublic = keyValue
		} else if keyType == "AccessSecret" {
			keys.accessSecret = keyValue
		} else if keyType == "AccessPublic" {
			keys.accessPublic = keyValue
		}
	}

	if keys.consumerSecret == "" || keys.consumerPublic == "" ||
		keys.accessSecret == "" || keys.accessPublic == "" {
		return nil, errors.New("Missing key in config\nConsumerSecret, ConsumerPublic, AccessSecret, AccessPublic required.")
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &keys, nil
}
