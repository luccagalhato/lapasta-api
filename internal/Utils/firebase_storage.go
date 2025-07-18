package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App
var bucketName string

func InicializarFirebase() error {
	configPath := "config/firebase-config.json"

	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("erro ao ler firebase-config.json: %v", err)
	}

	var parsedConfig struct {
		StorageBucket string `json:"storageBucket"`
	}

	if err := json.Unmarshal(jsonData, &parsedConfig); err != nil {
		return fmt.Errorf("erro ao parsear firebase-config.json: %v", err)
	}

	if parsedConfig.StorageBucket == "" {
		return fmt.Errorf("campo storageBucket não encontrado no firebase-config.json")
	}

	bucketName = parsedConfig.StorageBucket

	opt := option.WithCredentialsFile(configPath)
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		StorageBucket: bucketName,
	}, opt)
	if err != nil {
		return fmt.Errorf("erro ao iniciar Firebase: %v", err)
	}

	firebaseApp = app
	return nil
}

func UploadImagemFirebase(fileBytes []byte, fileName string, pasta string) (string, error) {
	ctx := context.Background()

	if firebaseApp == nil {
		if err := InicializarFirebase(); err != nil {
			return "", err
		}
	}

	client, err := firebaseApp.Storage(ctx)
	if err != nil {
		return "", fmt.Errorf("erro ao obter storage client: %v", err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("erro ao obter bucket: %v", err)
	}

	writer := bucket.Object(fmt.Sprintf("%s/%s", pasta, fileName)).NewWriter(ctx)
	writer.ContentType = "image/png"

	if _, err := io.Copy(writer, bytes.NewReader(fileBytes)); err != nil {
		return "", fmt.Errorf("erro ao escrever no bucket: %v", err)
	}
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("erro ao fechar writer: %v", err)
	}

	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s%%2F%s?alt=media", bucketName, pasta, fileName)
	return url, nil
}
