package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vfluxus/mailservice/core"
	"github.com/vfluxus/mailservice/repository/entity"
	"github.com/vfluxus/mailservice/webserver/dto"
)

// -------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- PARAM ----------------------------------------------------------

var (
	errEmptyID = errors.New("Empty id param")
)

func getIDFromQuery(c *gin.Context) (id uint32, err error) {
	idStr := c.Query("id")
	if len(idStr) == 0 {
		return 0, errEmptyID
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	if idInt <= 0 {
		return 0, errors.New("Invalid id")
	}

	return uint32(idInt), nil
}

func getPageSizeFromQuery(c *gin.Context) (page uint, size uint, err error) {
	pageStr := c.Query("page")
	if len(pageStr) == 0 {
		return 0, 0, nil
	}
	pageInt, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, 0, err
	}
	if pageInt <= 0 {
		return 0, 0, errors.New("Invalid page")
	}
	page = uint(pageInt)

	sizeStr := c.Query("size")
	if len(pageStr) == 0 {
		return 0, 0, nil
	}
	sizeInt, err := strconv.Atoi(sizeStr)
	if err != nil {
		return 0, 0, err
	}
	if sizeInt <= 0 {
		return 0, 0, errors.New("Invalid size")
	}
	size = uint(sizeInt)

	return page, size, nil
}

// ----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- RESPONSE ----------------------------------------------------------

func ginRespErrAbort(c *gin.Context, code int, msg string) {
	resp := new(dto.RespErr)
	resp.SetCodeMessage(code, msg)

	c.AbortWithStatusJSON(code, resp)
}

// ----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- TEMPLATE ----------------------------------------------------------

// ------------------------------
// mergeTemplate marshal variables to byte and pass to template
func mergeTemplate(template *entity.Template, variables []*entity.Variable) (err error) {
	if template == nil {
		return errors.New("nil template")
	}

	data, err := json.Marshal(variables)
	if err != nil {
		return err
	}

	template.Variables = data
	return nil
}

// ----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- PASSWORD ----------------------------------------------------------

// for more encrypt information https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/09.6.html
func encryptPassword(pass string) (hashed string, err error) {
	c, err := aes.NewCipher([]byte(core.GetConfig().Key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	hashedRaw := gcm.Seal(nonce, nonce, []byte(pass), nil)
	hashed = string(hex.EncodeToString(hashedRaw))
	return hashed, nil
}

func decryptPassword(hashed string) (pass string, err error) {
	hashedHex, err := hex.DecodeString(hashed)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher([]byte(core.GetConfig().Key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(hashed) < nonceSize {
		return "", errors.New("hashed too short")
	}

	nonce := hashedHex[:nonceSize]
	hashedHex = hashedHex[nonceSize:]

	decryptedPass, err := gcm.Open(nil, []byte(nonce), []byte(hashedHex), nil)
	if err != nil {
		return "", err
	}

	pass = string(decryptedPass)

	return pass, nil
}
