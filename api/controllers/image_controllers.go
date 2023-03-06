package controllers

import (
	"bytes"
	"errors"
	"fmt"
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
	fmt.Println("Called GetImages: ", r)
	userId, err := auth.ExtractTokenID(r)
	fmt.Println("Extract: ", userId, err)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	fmt.Println("Attemped Ext")

	fmt.Println(r.Header["Location"])
	fmt.Println("IDDDD: ", userId)

	imagesS3, err := getImagesS3(int64(userId), r.Header["Location"])
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	}
	fmt.Println("Image Array: ", imagesS3)
	responses.JSON(w, http.StatusOK, imagesS3)
}

func (server *Server) Upload(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	fmt.Println("1")
	loc := r.FormValue("loc")
	fmt.Println(loc)
	fmt.Println("Request", r)
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("2")
	file, _, err := r.FormFile("image")
	fmt.Println("3")

	buf := bytes.NewBuffer(nil)
	fmt.Println("4: ", file)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Starting upload")
	err = uploadImage(int64(userId), buf.Bytes(), loc)
	if err != nil {
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	fmt.Println("Upload Complete")

	responses.JSON(w, http.StatusOK)
}

func uploadImage(accountID int64, imgURL []byte, loc string) error {
	id := (uuid.New()).String()
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Printf("There was an issue uploading to s3: %s", err.Error())
		return errors.New("error: cant create item")
	}

	key := strconv.FormatInt(accountID, 10) + "/"
	fmt.Println("Starting Put")
	t, err := s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_USER_BUCKET")),
		Key:    aws.String(key + loc + "/" + id),
		Body:   aws.ReadSeekCloser(strings.NewReader(string(imgURL))),
	})

	fmt.Println("Ending Put")

	fmt.Println(key)

	fmt.Println("T: ", t)

	if err != nil {
		log.Printf("There was an issue uploading to s3: %s", err.Error())
		return errors.New("error: cant create item")
	}

	log.Printf("Successfully uploaded")
	return nil
}

func getImagesS3(accountID int64, loc []string) ([]string, error) {
	var userImages []string

	fmt.Println("LOCATION: ", loc)
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Printf("There was an issue uploading to s3: %s", err.Error())
		return nil, errors.New("error: cant create item")
	}

	svc := s3.New(sess)

	key := strconv.FormatInt(accountID, 10) + "/"

	params := &s3.ListObjectsV2Input{
		Bucket:    aws.String(os.Getenv("S3_USER_BUCKET")),
		Prefix:    aws.String(key + loc[0] + "/"),
		Delimiter: aws.String("/"),
	}

	fmt.Println("Params: ", params)

	resp, err := svc.ListObjectsV2(params)
	fmt.Println("List objects Error: ", err)
	fmt.Println(resp)

	if err != nil {
		fmt.Println("List objects Error: ", err)
		fmt.Println(resp)
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
			log.Println("Failed to sign request", err)
		}

		log.Println("The URL is", urlStr)
		userImages = append(userImages, urlStr)
	}

	return userImages, nil

	//key := strconv.FormatInt(accountID, 10) + "/"
	//t, err := s3.New(sess).PutObject(&s3.PutObjectInput{
	//	Bucket: aws.String(os.Getenv("S3_USER_BUCKET")),
	//	Key:    aws.String(key + loc + "/" + id),
	//	Body:   aws.ReadSeekCloser(strings.NewReader(string(imgURL))),
	//})
	//
	//fmt.Println(key)
	//
	//fmt.Println("T: ", t)
	//
	//if err != nil {
	//	log.Printf("There was an issue uploading to s3: %s", err.Error())
	//	return errors.New("error: cant create item")
	//}
	//
	//log.Printf("Successfully uploaded")
	//return nil
}
