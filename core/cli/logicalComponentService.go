package cli

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// LogicalComponentService :
type LogicalComponentService struct {
	client *Client
}

// List :
func (as *LogicalComponentService) List(token string) ([]models.LogicalComponent, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/logicalComponent", token, nil)
	if err != nil {
		return nil, err
	}
	var logicalComponents []models.LogicalComponent
	err = as.client.do(req, &logicalComponents)
	return logicalComponents, err
}

// Upload :
func (as *LogicalComponentService) Upload(token, tenant, typeName, fileToUpload string, params map[string]string) error {

	file, err := os.Open(fileToUpload)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("myFile", filepath.Base(fileToUpload))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := as.client.newRequestUpload("POST", "/ditrit/Gandalf/1.0.0/logicalComponent/upload/"+tenant+"/"+typeName, token, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *LogicalComponentService) Read(token string, id uuid.UUID) (*models.LogicalComponent, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/logicalComponent/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var logicalComponent models.LogicalComponent
	err = as.client.do(req, &logicalComponent)
	return &logicalComponent, err
}

// Read :
func (as *LogicalComponentService) ReadByName(token string, name string) (*models.LogicalComponent, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/logicalComponent/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var logicalComponent models.LogicalComponent
	err = as.client.do(req, &logicalComponent)
	return &logicalComponent, err
}
