package controllers

import (
	"bytes"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"iScore-api/api/auth"
	"iScore-api/api/responses"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func (server *Server) GetImages(w http.ResponseWriter, r *http.Request) {

	userId, err := auth.ExtractTokenID(r)

	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	imagesS3, err := getImagesS3(int64(userId), r.Header["Location"])
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	}

	responses.JSON(w, http.StatusOK, imagesS3)
}

func (server *Server) Upload(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	loc := r.FormValue("loc")

	err = r.ParseMultipartForm(32 << 20)
	if err != nil {

		return
	}

	file, _, err := r.FormFile("image")

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, file); err != nil {

		return
	}

	err = uploadImage(int64(userId), buf.Bytes(), loc)
	if err != nil {
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK)
}

func uploadImage(accountID int64, imgURL []byte, loc string) error {
	id := (uuid.New()).String()
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {

		return errors.New("error: cant create item")
	}

	key := strconv.FormatInt(accountID, 10) + "/"

	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_USER_BUCKET")),
		Key:    aws.String(key + loc + "/" + id),
		Body:   aws.ReadSeekCloser(strings.NewReader(string(imgURL))),
	})

	if err != nil {

		return errors.New("error: cant create item")
	}

	return nil
}

func getImagesS3(accountID int64, loc []string) ([]string, error) {
	var userImages []string

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {

		return nil, errors.New("error: cant create item")
	}

	svc := s3.New(sess)

	key := strconv.FormatInt(accountID, 10) + "/"

	params := &s3.ListObjectsV2Input{
		Bucket:    aws.String(os.Getenv("S3_USER_BUCKET")),
		Prefix:    aws.String(key + loc[0] + "/"),
		Delimiter: aws.String("/"),
	}

	resp, err := svc.ListObjectsV2(params)

	if err != nil {

		return nil, err
	}

	// Create S3 service client

	if *resp.KeyCount == 0 {
		return nil, nil
	}

	for _, key := range resp.Contents {
		if *key.Size == 0 {
			continue
		}
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("S3_USER_BUCKET")),
			Key:    aws.String(*key.Key),
		})

		urlStr, err := req.Presign(15 * time.Minute)

		if err != nil {

		}

		userImages = append(userImages, urlStr)
	}

	return userImages, nil
}

func getImagesS3General(loc string) ([]string, error) {
	var userImages []string

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {

		return nil, errors.New("error: cant create item")
	}

	svc := s3.New(sess)

	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("S3_GENERAL_BUCKET")),
		Prefix: aws.String(loc),
		//Delimiter: aws.String("/"),
	}

	resp, err := svc.ListObjectsV2(params)

	if err != nil {

		return nil, err
	}

	// Create S3 service client

	if *resp.KeyCount == 0 {

		return nil, nil
	}

	for _, key := range resp.Contents {
		if *key.Size == 0 {
			continue
		}
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("S3_GENERAL_BUCKET")),
			Key:    aws.String(*key.Key),
		})

		urlStr, err := req.Presign(15 * time.Minute)

		if err != nil {

		}

		userImages = append(userImages, urlStr)
	}

	return userImages, nil
}
