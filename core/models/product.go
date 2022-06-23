package models

import (


	"github.com/google/uuid"
	// "github.com/jinzhu/gorm"
)

type Product struct {
	Model
	Name             string `gorm:"unique;not null"`
	ShortDescription string
	Description      string
	Logo             string
	
	DomainID         uuid.UUID `gorm:"type:uuid"`
	Domain           *Domain

	Authorizations   []Authorization
	Libraries        []Library `gorm:"many2many:product_libraries;"`
	Tags             []Tag     `gorm:"many2many:product_tags;"`
	Environments     []Environment

	GitServerURL string `gorm:"not null"`
	GitPersonalAccessToken string `gorm:"not null"`
	GitOrganization string

	GitRepoName string
}

func (p *Product) GetGitServerURL() string {
	if p.GitServerURL[len(p.GitServerURL) - 1:] != "/" {
		p.GitServerURL = p.GitServerURL + "/"
	}
	return p.GitServerURL
}

func (p *Product) GetGitOrganization() string {
	if p.GitOrganization[len(p.GitOrganization) - 1:] != "/" {
		p.GitOrganization = p.GitOrganization + "/"
	}
	return p.GitOrganization
}

func (p *Product) GetGitRepoName() string {
	if p.GitRepoName == "" {
		return p.Name
	}
	return p.GitRepoName
}

func (p *Product) GetRepoURL() string {
	return p.GetGitServerURL() + p.GetGitOrganization() + p.GetGitRepoName() + ".git"
}