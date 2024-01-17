// TBD
package logger

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger(c *gin.Context) {
	start := time.Now()

	crw := &CustomResponseWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = crw

	reqBody, _ := c.GetRawData()
	if len(reqBody) > 0 {
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	}
	fmt.Printf("[%s]|%s|%s|Request: %s\n", INFO, c.Request.Method, c.Request.RequestURI, reqBody)

	c.Next()

	end := time.Now()
	latency := end.Sub(start)
	respBody := crw.body.String()
	fmt.Printf("[%s]|%s|%s|Response: %s|(%v)\n", INFO, c.Request.Method, c.Request.RequestURI, respBody, latency)
}
