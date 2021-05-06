package image

import (
	"github.com/kgalli/image-check/internal/data"
)

type ImageRepository struct {
	conn *data.Connection
}

func NewImageRepository(conn *data.Connection) *ImageRepository {
	return &ImageRepository{
		conn: conn,
	}
}

func (r *ImageRepository) Create(baseImage BaseImage) (Image, error) {
	image := Image{
		Name: baseImage.Name,
		Tag:  baseImage.Tag,
		Digests: []ImageDigest{
			{
				Digest: baseImage.Digest,
			},
		},
	}

	result := r.conn.DB.Create(&image)

	return image, result.Error
}

func (r *ImageRepository) AddDigest(name string, tag string, digest string) error {
	var image Image
	imageToFind := Image{Name: name, Tag: tag}
	if result := r.conn.DB.Where(&imageToFind).First(&image); result.Error != nil {
		return result.Error
	}

	imageDigest := ImageDigest{
		ImageID: image.ID,
		Digest:  digest,
	}

	return r.conn.DB.Create(&imageDigest).Error
}

func (r *ImageRepository) Migrate() error {
	return r.conn.DB.AutoMigrate(&Image{}, &ImageDigest{}).Error
}
