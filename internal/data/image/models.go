package image

import (
	"gorm.io/gorm"
)

type BaseImage struct {
	Name   string
	Tag    string
	Digest string
}

type Image struct {
	gorm.Model
	Name    string
	Tag     string
	Digests []ImageDigest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ImageDigest struct {
	gorm.Model
	ImageID uint
	Digest  string
}
