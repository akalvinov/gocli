package ops

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/joho/godotenv"
)

func SaveSecret(secretName string, filName string, secretFiles map[string]string) error {
	secretJSON, err := GetSecret(secretName)
	if err != nil {
		return err
	}
	secretEnvs, err := ParseSecret(secretJSON)
	if err != nil {
		return err
	}
	err = WriteEnvFile(secretEnvs, filName)
	if err != nil {
		return err
	}

	for file, secret := range secretFiles {
		awsSecret, err := GetSecret(secret)
		if err != nil {
			return err
		}
		if _, err := os.Stat(file); err == nil {
			return errors.New("File " + file + " already exists. Can not overwrite")
		} else {
			err = os.WriteFile(file, []byte(awsSecret), 0444)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ValidateEnvs(requiredEnvs []string) error {
	var failed []string
	var failedFormat []string
	var validValue = regexp.MustCompile(`^[\pL\pM\pN\pP\pS]+$`)
	for _, value := range requiredEnvs {
		_, present := os.LookupEnv(value)
		if !present {
			failed = append(failed, value)
		} else if !validValue.MatchString(os.Getenv(value)) {
			failedFormat = append(failedFormat, value)
		}
	}
	if len(failed) > 0 {
		return errors.New("Failed to validate: no_value: " + strings.Join(failed, ", ") + "; not_printable: " + strings.Join(failedFormat, ", ")) //TODO: rewrite with error type
	}
	return nil
}

func WriteEnvFile(values map[string]string, fileName string) error {
	//TODO: validate k\v
	//TODO: parametrize file
	err := godotenv.Write(values, fileName)
	return err
}

func ParseSecret(rawSecret string) (map[string]string, error) {
	var vals map[string]string
	err := json.Unmarshal([]byte(rawSecret), &vals)
	return vals, err
}

func GetSecret(secretName string) (string, error) {
	region := "eu-central-1"
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return "", err
	}
	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}
	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return "", errors.Join(err, errors.New("Failed to retrieve secret: "+secretName))
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString

	return secretString, nil
}
