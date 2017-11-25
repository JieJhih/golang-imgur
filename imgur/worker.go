package imgur

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type UploadImageForm struct {
	Title    string `form:"title" json:"title,omitempty"`
	Category string `form:"category" binding:"required" json:"category,omitempty"`
	Image    string `form:"image" json:"image,omitempty"`
}

func UploadImage(c *gin.Context) {
	var form UploadImageForm
	if err := c.Bind(&form); err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
	}
	var result ImgurResponse
	sync := make(chan bool, 0)
	client := &http.Client{}
	data := url.Values{}
	data.Set("title", form.Title)
	if form.Category == "file" {

		go func() {
			file, err := c.FormFile("image")
			if err != nil {
				ResponseError(c, http.StatusInternalServerError, err.Error())
				return
			}
			src, _ := file.Open()
			defer src.Close()

			buf := make([]byte, file.Size)
			fReader := bufio.NewReader(src)
			fReader.Read(buf)
			imgBase64Str := base64.StdEncoding.EncodeToString(buf)
			data.Set("image", imgBase64Str)
			data.Set("type", "base64")

			sync <- true
		}()

	} else {

		if form.Image == "" {
			ResponseError(c, http.StatusBadRequest, "Image field cannot null")
			return
		}
		data.Set("image", form.Image)
		data.Set("type", form.Category)
	}
	<-sync
	req, err := http.NewRequest("POST", "https://api.imgur.com/3/image", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Client-ID "+Conf.Auth.ClientID)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		logrus.Println(err)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data":    result.Data,
			"success": result.Success,
			"status":  result.Status,
		},
	)
}

func ResponseError(c *gin.Context, code int, err string) {
	c.AbortWithStatusJSON(
		code,
		gin.H{
			"error": err,
		},
	)
}
