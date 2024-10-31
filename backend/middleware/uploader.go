package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileUploadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set max file size (3MB)
		const maxSize = 3 * 1024 * 1024 

		if err := c.Request.ParseMultipartForm(maxSize); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "File size exceeds 5MB limit",
			})
			return
		}

		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad request",
			})
			return
		}
		defer file.Close()

		if header.Size > maxSize {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "File size exceeds 3MB limit",
			})
			return
		}

		buffer := make([]byte, 512)
		if _, err := file.Read(buffer); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Could not read file",
			})
			return
		}

		if _, err := file.Seek(0, 0); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Could not process file",
			})
			return
		}

		fileType := http.DetectContentType(buffer)
		allowedTypes := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/gif":  true,
			"image/webp": true,
			"image/jpg":  true,
		}

		if !allowedTypes[fileType] {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "File type not allowed. Please upload JPEG, PNG or GIF",
			})
			return
		}

		c.Set("filePath", header.Filename)
		c.Set("file", file)

		c.Next()
	}
}
