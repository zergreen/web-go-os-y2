package info

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

//รับ INPUT แปลงค่า
func (ep *Endpoint) CreateUserEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	var request userRequest //model รับ input จาก body

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	Id := request.Id
	FirstName := request.FirstName
	LastName := request.LastName
	District := request.District

	fmt.Printf("\n CreateUserInfo... ")
	result, err := createUserInfo(Id, FirstName, LastName, District)
	if err != nil {
		fmt.Printf("\n CreateUserInfo fail ")
		c.JSON(http.StatusBadRequest, err)
	}
	fmt.Printf("\n CreateUserInfo succeed")
	c.JSON(http.StatusOK, result)
}

//รับ INPUT แปลงค่า
func (ep *Endpoint) InfoEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	var request userRequest //model รับ input จาก body

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	Id := request.Id
	FirstName := request.FirstName
	LastName := request.LastName
	District := request.District
	fmt.Printf("\n firstName:%v",FirstName)
	fmt.Printf("\n checkUserInfo... ")
	result, err := checkInfo(Id, FirstName, LastName, District)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)

}

func (ep *Endpoint) ShowGeographyEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()
	fmt.Printf("\n showAllGeography... ")
	result, err := showGeography()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)

}
