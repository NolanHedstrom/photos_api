package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cozy/goexif2/exif"
	"github.com/gin-gonic/gin"
)

//PhotoInfo is a struct for displaying photo metadata
type PhotoInfo struct {
	PhotoName string    `json:"Photo Name"`
	Camera    string    `json:"Camera Model"`
	TimeTaken time.Time `json:"Time Taken"`
	Lat       float64   `json:"Latitude"`
	Long      float64   `json:"Longitude"`
}

// Photos is a struct for displaying an array of photo metadata
type Photos struct {
	Photos []PhotoInfo `json:"Photos"`
}

//LoadPhoto will show a photo
func LoadPhoto(c *gin.Context) {

	var photodetails PhotoInfo
	file, err := os.Open("/Users/nolan/Downloads/Photos/IMG_7798.jpg") // For read access.
	if err != nil {
		log.Println(err.Error())
	}

	photodetails.PhotoName = filepath.Base(file.Name())

	//exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(file)
	if err != nil {
		fmt.Println("Erroring out")
		log.Fatal(err)
	}

	camModel, err := x.Get(exif.Model) // normally, don't ignore errors!
	if err != nil && camModel != nil {
		photodetails.Camera, _ = camModel.StringVal()
	}

	// Two convenience functions exist for date/time taken and GPS coords:
	photodetails.TimeTaken, _ = x.DateTime()

	photodetails.Lat, photodetails.Long, _ = x.LatLong()

	c.JSON(http.StatusOK, photodetails)

}

func LoadPhotos(c *gin.Context) {

	var photos []PhotoInfo

	dirname := "/Users/nolan/Downloads/Photos/"

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		photo, err := GetPhotoInfo(file.Name(), dirname)
		if err != nil {
			fmt.Println(err)
		} else {
			photos = append(photos, photo)
		}
	}

	var photolist Photos
	photolist.Photos = photos

	c.JSON(http.StatusOK, photolist)
}

func GetPhotoInfo(fileName string, dir string) (PhotoInfo, error) {
	var photoDetails PhotoInfo

	file, err := os.Open(dir + fileName) // For read access.
	if err != nil {
		log.Println(err.Error())
		return photoDetails, err
	}
	photoDetails.PhotoName = filepath.Base(file.Name())

	//exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(file)
	if err != nil {
		fmt.Println("Erroring out")
		log.Fatal(err)
	}

	camModel, err := x.Get(exif.Model) // normally, don't ignore errors!
	if err == nil && camModel != nil {
		photoDetails.Camera, _ = camModel.StringVal()
	} else {
		photoDetails.Camera = "PLEASE UPDATE METADATA"
	}

	// Two convenience functions exist for date/time taken and GPS coords:
	photoDetails.TimeTaken, _ = x.DateTime()

	photoDetails.Lat, photoDetails.Long, _ = x.LatLong()
	return photoDetails, nil
}
