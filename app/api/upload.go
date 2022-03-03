package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
)

func UploadBinary(c *gin.Context) {
	// ?act=addPicInfo
	act := c.Query("act")
	f, err := c.FormFile("resource")
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("UploadBinary---->>", f.Filename, act)

		dst := path.Join("./", f.Filename)
		_ = c.SaveUploadedFile(f, dst)

	}

}
