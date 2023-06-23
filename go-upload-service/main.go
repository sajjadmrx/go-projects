package main

import (
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var folder = "storage/images"

func init() {

	_, err := os.Stat(folder)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Created")
	}

}

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "Hello World"})
	})

	router.POST("/upload/image", uploadFile)
	err := router.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}

}

func uploadFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err %s", err.Error()))
		return
	}

	format := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), format)

	now := time.Now()

	imageFile, _, err := image.Decode(file)

	if err != nil {
		log.Fatal(err)
	}
	width, height := 256, 256
	src := imaging.Resize(imageFile, width, height, imaging.Lanczos)

	filename :=
		strings.ReplaceAll(
			strings.ToLower(
				originalFileName), " ", "_") + "-" + fmt.Sprintf(
			"%d-%d-%v", width, height, now.Unix()) + format

	filePath := folder + "/" + filename

	err = imaging.Save(src, folder+"/"+filename)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"filepath": filePath,
	})

}
