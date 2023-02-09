package v1

import (
	"fmt"

	"github.com/chenliu1993/go-learning/pkg/constant"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func EchoHello(c *gin.Context) {
	// The rsequestsedurl /api/v1/hello?token=*******&msg=*
	// Better use postman to do it
	sender := c.PostForm("test")
	customMsg := c.Query("msg")
	c.JSON(constant.StatusOK, Response{
		Code: 200,
		Msg:  "test",
		Data: fmt.Sprintf("Hello From gin %s from %s", customMsg, sender),
	})
}
