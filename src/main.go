package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redpanda-data/benthos/v4/public/service"
)

func main() {

	prodFunc, stream, err := BuildStream(streamSrc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, streamInput, err := BuildStream(streamSrcInput)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w := &Web{prodFunc}
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/", w.benthosInput)
	r.POST("/output", w.benthosOutput)
	go func() {

		r.Run(fmt.Sprintf("0.0.0.0:%d", 8080))
	}()

	go func() {
		fmt.Println("Stream stream with Input")
		if err := streamInput.Run(context.Background()); err != nil {

			fmt.Println(err.Error())
			return
		}

	}()
	fmt.Println("Stream stream with Producer function")
	if err := stream.Run(context.Background()); err != nil {
		fmt.Println(err.Error())
		return
	}

}

type Web struct {
	prodFunc service.MessageHandlerFunc
}

func (w *Web) benthosInput(c *gin.Context) {
	payload := map[string]string{
		"id": "123",
	}
	buf, _ := json.Marshal(payload)
	msg := service.NewMessage(buf)
	w.prodFunc(c.Request.Context(), msg)

	c.JSON(200, payload)
}

func (w *Web) benthosOutput(c *gin.Context) {
	var payload map[string]string
	c.BindJSON(&payload)
	buf, _ := json.Marshal(payload)
	fmt.Printf("Output : %s ", string(buf))
	c.JSON(200, payload)
}
