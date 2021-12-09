package cli

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// ProductService :
type ProductService struct {
	client *Client
}

// List :
func (as *ProductService) List(token string) ([]models.Product, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/product/", token, nil)
	if err != nil {
		return nil, err
	}
	var products []models.Product
	err = as.client.do(req, &products)
	return products, err
}

// Create :
func (as *ProductService) Create(token string, product models.Product, parentProductName string) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/product/"+parentProductName, token, product)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *ProductService) Read(token string, id uuid.UUID) (*models.Product, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/product/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var product models.Product
	err = as.client.do(req, &product)
	return &product, err
}

// Read :
func (as *ProductService) ReadByName(token string, name string) (*models.Product, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/product/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var product models.Product
	err = as.client.do(req, &product)
	return &product, err
}

// Update :
func (as *ProductService) Update(token string, id uuid.UUID, product models.Product) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/product/"+id.String(), token, product)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ProductService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/product/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
