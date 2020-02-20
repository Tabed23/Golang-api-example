package router

import (
	"fmt"
	"gin_pratice_api/database"
	"gin_pratice_api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)


func GetAllEmployee(c * gin.Context){
	db ,err := database.GetMongo()

	if err!= nil{
		c.JSON(http.StatusOK, gin.H{"status":http.StatusOK,
			"massage":"Database error ",
		})
		return

	}
	fmt.Println("Mongo server is running ",db.Name)
	var emp = models.Employees
	err =  db.C("employee").Find(bson.M{}).All(&emp)

	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status":http.StatusOK,
			"massage":"Database error Get All Employee",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":http.StatusOK,
		"employee":&emp,
	})
}


func GetEmployee(c *gin.Context){
	db ,err := database.GetMongo()

	if err!= nil{
		c.JSON(http.StatusOK, gin.H{"status":http.StatusOK,
			"massage":"Database error ",
		})
		return

	}
	fmt.Println("Mongo server is running ", db.Name)
	id := c.Param("id")
	idStr , idError :=  strconv.Atoi(id)
	if idError != nil{
		c.JSON(http.StatusOK, gin.H{"status":http.StatusOK,
			"massage": "error Parsing id"},)
		return
	}
	emp  :=  models.Employees
	err =  db.C("employee").Find(bson.M{"id":idStr}).One(&emp)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"status":http.StatusOK,
			"massage": "Employee not found"},)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status":http.StatusOK,
		"employee":&emp,
	})

}
func CreateUser(c *gin.Context) {
	db,err := database.GetMongo()
	fmt.Println("MONGO RUNNING: ", db.Name)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"status":http.StatusOK,
			"massage": "Database error"},
			)
		return
	}
	emp := models.Employee{}
	err = c.Bind(&emp)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	err = db.C("employee").Insert(emp)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Insert User",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success Insert User",
		"employee":    &emp,
	})
}
// Delete
func DeleteEmployee(c * gin.Context){
	db ,err := database.GetMongo()

	if err!= nil{
		c.JSON(http.StatusOK, gin.H{"status":http.StatusOK,
			"massage":"Database error ",
		})
		return

	}
	fmt.Println("MONGO RUNNING: ", db.Name)

	id := c.Param("id")
	idParse, errParse := strconv.Atoi(id)
	if errParse != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Error Parsing id",
		})
		return
	}

	err = db.C("employee").Remove(bson.M{"id": &idParse})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Error Delete User",
			"status" :http.StatusOK,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully  Delete User",
	})
}

func UpdateEmployee(c *gin.Context){
	db, err :=  database.GetMongo()
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"status":http.StatusOK,
			"massage":"Database error ",
		})
		return

	}
	id :=  c.Param("id")
	fmt.Println("MONGO RUNNING: ", db.Name)
	idParse, errParse := strconv.Atoi(id)
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}
	emp :=  models.Employee{}
	err =  c.Bind(&emp)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "Error Get Body",
		})
		return
	}
	emp.ID = int32(idParse)
	err = db.C("employee").Update(bson.M{"id":&idParse},emp)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":http.StatusOK,
			"message": "Error Update User",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Update User",
		"status":http.StatusOK,
		"employee":    &emp,
	})


}
