package controller

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
)

var (
	claimsKey = "claims"
)

var (
	ErrClaimsNotExist   = errors.New("claims not exist")
	ErrNotClaims        = errors.New("not claims")
	ErrInvalidPostID    = errors.New("invalid post id")
	ErrInvalidUserID    = errors.New("invalid user id")
	ErrInvalidCommentID = errors.New("invalid comment id")
	ErrInvalidMediaID   = errors.New("invalid media id")
)

func getClaimsFromCtx(c *gin.Context) (*auth.Claims, error) {
	claimsItf, ok := c.Get(claimsKey)
	if !ok {
		return nil, ErrClaimsNotExist
	}

	claims, ok := claimsItf.(auth.Claims)
	if !ok {
		return nil, ErrNotClaims
	}

	return &claims, nil
}

func getUserUUIDFromClaims(c *gin.Context) (uuid.UUID, error) {
	claims, err := getClaimsFromCtx(c)
	if err != nil {
		return uuid.Nil, err
	}

	userUUID, err := uuid.Parse(claims.User.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return userUUID, nil
}

func getUserUUID(c *gin.Context) string {
	claims, err := getClaimsFromCtx(c)
	if err != nil {
		return ""
	}

	return claims.User.ID
}

func getUserUUIDFromParam(c *gin.Context) (uuid.UUID, error) {
	userID := c.Param("userid")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, ErrInvalidUserID
	}

	return userUUID, nil
}

func getPostID(c *gin.Context) (uuid.UUID, error) {
	postID := c.Param("postid")
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return uuid.Nil, ErrInvalidPostID
	}

	return postUUID, nil
}

func getMediaID(c *gin.Context) (uuid.UUID, error) {
	mediaID := c.Param("mediaid")
	mediaUUID, err := uuid.Parse(mediaID)
	if err != nil {
		return uuid.Nil, ErrInvalidMediaID
	}

	return mediaUUID, nil
}

func getCommentID(c *gin.Context) (uuid.UUID, error) {
	commentID := c.Param("commentid")
	commentUUID, err := uuid.Parse(commentID)
	if err != nil {
		return uuid.Nil, ErrInvalidCommentID
	}

	return commentUUID, nil
}

func genMultipartOpts(c *gin.Context) []application.MultipartOption {
	if !strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
		return nil
	}

	form, err := c.MultipartForm()
	if err != nil {
		return nil
	}
	images := form.File["image"]
	video, err := c.FormFile("video")
	if err != nil {
		return nil
	}

	return []application.MultipartOption{application.WithImagesMultipart(images), application.WithVideoMultipart(video)}
}

const source = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func genRandomString(length int) string {
	seedRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var strB = new(strings.Builder)
	strB.Grow(length)
	for i := 0; i < length; i++ {
		strB.WriteByte(source[seedRand.Intn(len(source))])
	}

	return strB.String()
}
