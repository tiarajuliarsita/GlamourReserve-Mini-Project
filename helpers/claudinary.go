package helpers

import (
	"bytes"
	"context"
	"glamour_reserve/app/config"
	"io"
	"mime/multipart"
	"path"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func RemoveExtention(fileName string) string {
	return path.Base(fileName[:len(fileName)-len(path.Ext(fileName))])
}

func ConvertFile(FileHeader *multipart.FileHeader) (*bytes.Reader, error) {
	file, err := FileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, file)
	if err != nil {
		return nil, err
	}

	fileReader := bytes.NewReader(buffer.Bytes())
	return fileReader, nil

}

func UploadFile(fileHeader *multipart.FileHeader, fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudApiKey(), config.EnvCloudApiSecret())
	if err != nil {
		return "", err
	}

	fileReader, err := ConvertFile(fileHeader)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, fileReader, uploader.UploadParams{
		PublicID: fileName,
		Folder:   config.EnvCloudUploadFolder(),
	})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}
