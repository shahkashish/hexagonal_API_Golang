package controllers

import (
	"api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)
type Data struct {
	Id string `json : "id"`
	Name string `json : "name"`
}


type Controller struct{
	repo repository.Repository
}
func NewController(repo repository.Repository) *Controller {
    return &Controller{repo: repo}
}

func PostToDB() string {
	return "this is data from db"
}
func (d *Controller) Deletedata(c *gin.Context) {
	name := c.Query("name")
	id := c.Query("id")
	s := d.repo.Deletedata(name,id)
	c.JSON(200, gin.H{
		"data": s,
	})
}
func (d *Controller) AddData(c *gin.Context) {
	var requestData Data
    if err := c.BindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	print("data " + requestData.Id)
	s := d.repo.AddData(requestData.Name,requestData.Id)
	c.JSON(200, gin.H{
		"data": s,
	})
}