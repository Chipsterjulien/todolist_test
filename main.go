package main

import (
	"github.com/WnP/todolist_test/db"
	"github.com/gin-gonic/gin"
)

var dbmap = db.InitDb()

func indexHandler(c *gin.Context) {
	messages := []db.Messages{}

	dbmap.Find(&messages)

	if len(messages) == 0 {
		obj := gin.H{"error": "No message in the database"}
		c.HTML(200, "index.html", obj)
	} else {
		obj := gin.H{"title": messages}
		c.HTML(200, "index.html", obj)
	}
}

func postHandler(c *gin.Context) {
	// On récupère le formulaire
	c.Request.ParseForm()
	// On récupère la valeur de title
	title := c.Request.Form.Get("title")

	if title != "" {
		data := &db.Messages{Title: title}
		dbmap.Save(&data)
	}

	c.Redirect(301, "/")
}

func deleteHandler(c *gin.Context) {
	c.Request.ParseForm()
	id := c.Request.Form.Get("id")

	if id != "" {
		dbmap.Where("id = ?", id).Delete(db.Messages{})
	}
	//id := c.Params.ByName("id")
	//dbmap.Where("id = ?", id).Delete(db.Messages{})

	c.Redirect(301, "/")
}

func main() {
	// go get github.com/gin-gonic/gin
	// go get github.com/jinzhu/gorm
	// go get github.com/mattn/go-sqlite3

	// Initialisation de gin
	r := gin.Default()
	// Définit où se trouve les templates html
	r.LoadHTMLGlob("views/*")
	// Déclare la première route, quand l'utilisateur arrive sur /
	r.GET("/", indexHandler)
	// Déclare une route quand l'utilisateur se connecte sur /submit
	r.POST("/submit", postHandler)
	//r.GET("/delete/:id", deleteHandler)
	r.POST("/delete", deleteHandler)

	r.Run(":3000")
}
