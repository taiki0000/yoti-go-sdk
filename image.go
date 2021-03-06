package yoti

import (
	"encoding/base64"
)

// Deprecated: Will be removed in v3.0.0 - use attribute.ContentType instead. ImageType Image format
type ImageType int

const (
	// ImageTypeJpeg JPEG format
	ImageTypeJpeg ImageType = 1 + iota
	// ImageTypePng PNG format
	ImageTypePng
	// ImageTypeOther Other image formats
	ImageTypeOther
)

// Deprecated: Will be removed in v3.0.0 - use attribute.Image instead. ImageType struct containing
// the type of the image and the data in bytes.
type Image struct {
	Type ImageType
	Data []byte
}

// Deprecated: Will be removed in v3.0.0, please use image.GetMIMEType instead. GetContentType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func (image *Image) GetContentType() string {
	switch image.Type {
	case ImageTypeJpeg:
		return "image/jpeg"

	case ImageTypePng:
		return "image/png"

	default:
		return ""
	}
}

// Deprecated: Will be removed in v3.0.0, please use image.Base64URL() instead. URL Image encoded in a base64 URL
func (image *Image) URL() string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(image.Data)
	return "data:" + image.GetContentType() + ";base64;," + base64EncodedImage
}
