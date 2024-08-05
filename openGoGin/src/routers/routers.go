package routers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	// database of choice
	_ "github.com/go-sql-driver/mysql"
)

// Request = simple struct to pass data
type Request struct {
	RequestID int       `json:"requestid"`
	DataRows  []DataRow `json:"widgets"`
}

// DataRow = simple struct to return multiple rows
type DataRow struct {
	Column1 string `json:"column1"`
	Column2 string `json:"column2"`
	Column3 string `json:"column3"`
	Column4 int    `json:"column4"`
}

// RegisterRouters - serve up webpages and APIs
func RegisterRouters() *gin.Engine {
	var db *sql.DB

	// Database Open for AWS
	//db, err := sql.Open("mysql", "rootuser:password@tcp(address.of.us-west-2.rds.amazonaws.com:3306)/dbname")

	// Database Open for mysql
	db, err := sql.Open("mysql", "openuser:openpass@tcp(127.0.0.1:3306)/opendb")

	if err != nil {
		panic(err.Error())
	}

	// eventually close the database connection
	defer db.Close()

	// always a good idea to ping the database to check for access
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// base gin structure
	router := gin.Default()

	// load the assets and html pages
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "assets")
	router.StaticFile("/favicon.ico", "assets/favicon.ico")

	// GETs for webpages
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to OpenTable",
		})
	})

	router.GET("/column2/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create_column2.html", gin.H{
			"title": "Create Column2",
		})
	})

	router.GET("/column2/show/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "column2.html", gin.H{
			"title": "View Column2",
		})
	})

	router.GET("/column1/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "column1.html", gin.H{
			"title": "Column1 Information",
		})
	})

	//
	// API section
	//

	// GET details for a given column1
	router.GET("/api/column1/:value", func(c *gin.Context) {
		value := c.Param("value")
		row, err := RowInfo(value, db)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK, row)
	})

	// GET - Pull all rows based on parameters - all are optional or can be wildcarded
	router.GET("/api/column1", func(c *gin.Context) {
		column2Parm := c.DefaultQuery("column2", "%")
		column3Parm := c.DefaultQuery("column3", "%")

		requestList, err := RowListInfo(column2Parm, column3Parm, db)
		if err != nil {
			fmt.Println(err.Error()) // could also panic here, but it would be a crash
		}
		c.JSON(http.StatusOK, requestList)
	})

	//
	// PUT - Create a row
	router.POST("/api/column1", func(c *gin.Context) {
		var request DataRow

		var x []byte
		x, _ = ioutil.ReadAll(c.Request.Body)
		err := json.Unmarshal(x, &request)
		if err != nil {
			fmt.Println(err.Error())
		}

		requestResponse, err := CreateRow(request, db)
		if err != nil {
			fmt.Println(err.Error()) // could also panic here, but it would be a crash
		}
		c.JSON(http.StatusOK, requestResponse)
	})

	// PUT - update an column2 request
	router.PUT("/api/row/:value", func(c *gin.Context) {
		value := c.Param("value")
		fmt.Println(value)
		var x []byte
		x, _ = ioutil.ReadAll(c.Request.Body)
		var row DataRow
		err := json.Unmarshal(x, &row)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusNotAcceptable, 0)

		} else {
			updateResults, err := UpdateRow(row, db)
			if err != nil {
				fmt.Println(err.Error()) // could also panic here, but it would be a crash
				c.JSON(http.StatusNotAcceptable, 0)
			}
			c.JSON(http.StatusOK, updateResults)
		}

	})

	//return to main program
	return router
}
