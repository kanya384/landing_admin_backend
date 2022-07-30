package helpers

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/disintegration/imaging"
	formatter "github.com/fabienm/go-logrus-formatters"
	satori "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserIdFromHandler(input interface{}) (userID primitive.ObjectID, err error) {
	res, ok := input.(map[string]interface{})
	if !ok {
		return userID, errors.New("no value")
	}
	id, ok := res["id"].(string)
	if !ok {
		return userID, errors.New("no value")
	}
	return primitive.ObjectIDFromHex(id)
}

func SetValueToCookies(w http.ResponseWriter, name string, value string, ttl time.Duration, secured bool) {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: time.Now().Add(ttl),
		Secure:  secured,
		Path:    "/",
	})
}

//check if interface is image and saves image (for create and update handlers)
func ProcessImage(file interface{}, fileStorePath string, IMAGE_WIDTH int, IMAGE_HEIGHT int) (filename string, err error) {
	fileIn, ok := file.(io.ReadCloser)
	if !ok {
		return "", errors.New("no file")
	}
	defer fileIn.Close()
	image, err := ResizeImage(fileIn, IMAGE_WIDTH, IMAGE_HEIGHT)
	if err != nil {
		return "", err
	}
	filename, err = SaveImageFile(image, "tmg.jpg", fileStorePath)
	if err != nil {
		return
	}
	return
}

func ProcessImages(file interface{}, fileStorePath string, IMAGE_WIDTH int, IMAGE_HEIGHT int, IMAGE_WIDTH_1 int, IMAGE_HEIGHT_1 int) (filename string, filename2 string, err error) {
	fileIn, ok := file.(io.ReadCloser)
	if !ok {
		return "", "", errors.New("no file")
	}
	defer fileIn.Close()
	imageIn, err := imaging.Decode(fileIn)
	if err != nil {
		fmt.Println(err)
	}
	image1 := imaging.Fill(imageIn, IMAGE_WIDTH, IMAGE_HEIGHT, imaging.Center, imaging.Lanczos)
	image2 := imaging.Fill(imageIn, IMAGE_WIDTH_1, IMAGE_HEIGHT_1, imaging.Center, imaging.Lanczos)
	if err != nil {
		return "", "", err
	}
	filename, err = SaveImageFile(image1, "tmg.jpg", fileStorePath)
	if err != nil {
		return
	}
	filename2, err = SaveImageFile(image2, "tmg.jpg", fileStorePath)
	if err != nil {
		return
	}
	return filename, filename2, err
}

func SavePlan(file interface{}, fileStorePath string, IMAGE_WIDTH int, IMAGE_HEIGHT int) (filename string, err error) {
	fileIn, ok := file.(io.ReadCloser)
	if !ok {
		return "", errors.New("no file")
	}
	defer fileIn.Close()

	body, err := ioutil.ReadAll(fileIn)
	if err != nil {
		return "", errors.New("no file")
	}

	filename, err = SaveFile(body, "tmg.png", fileStorePath)
	if err != nil {
		return
	}
	return
}

func ResizeImage(file io.ReadCloser, width int, height int) (image image.Image, err error) {
	imageIn, err := imaging.Decode(file)
	if err != nil {
		fmt.Println(err)
	}
	if height != 0 {
		image = imaging.Fill(imageIn, width, height, imaging.Center, imaging.Lanczos)
	} else {
		image = imaging.Resize(imageIn, width, 0, imaging.Lanczos)
	}
	return
}

func SaveImageFile(image image.Image, fileNameIn string, path string) (filename string, err error) {
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, image, nil)
	if err != nil {
		return
	}
	send_s3 := buf.Bytes()
	filename, err = SaveFile(send_s3, fileNameIn, path)
	return
}

func SaveFile(fileContent []byte, fileNameIn string, path string) (string, error) {
	fileSet := strings.Split(fileNameIn, ".")[len(strings.Split(fileNameIn, "."))-1]
	id := satori.NewV4()
	fileName := id.String() + "." + fileSet

	err := ioutil.WriteFile(path+"/"+fileName, fileContent, 0644)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func DeleteFile(filename string, path string) (err error) {
	return os.Remove(fmt.Sprintf("%s/%s", path, filename))
}

func ReadCsvFile(file io.ReadCloser) (rows [][]string, err error) {
	defer file.Close()
	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1
	rows, err = csvReader.ReadAll()
	rows = rows[1:]
	return
}

func ConfigureLogger(serviceName string, logLevel uint32, file *os.File) (*logrus.Entry, error) {
	log := logrus.New()
	log.SetLevel(logrus.Level(logLevel))
	gelFmt := formatter.NewGelf(serviceName)
	runtimeFormatter := &runtime.Formatter{ChildFormatter: gelFmt}
	log.SetFormatter(runtimeFormatter)

	log.SetOutput(file)
	//hook := graylog.NewGraylogHook(greyLogHost, map[string]interface{}{})
	//log.AddHook(hook)

	return log.WithField("service_name", "backend"), nil
}

func GenerateHashPassword(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
