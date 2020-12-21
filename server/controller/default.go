package controller
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tokenAuthGolang/model"
	"tokenAuthGolang/services"
	"github.com/jackc/pgx/v4"
)

func UpAndRunningResponse(c *gin.Context){
	c.JSON(200, gin.H{
		"msg": "server up and running",
	})
}

func ProcessLogin(context *gin.Context) {
	var outputModel model.FormLogin
	err := services.GetDatabaseInstance().QueryRow(context, "SELECT e.employee_id, e.login FROM employee e WHERE e.login=$1 AND e.password = $2", "admin","123").Scan(outputModel.Login, outputModel.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{
				"data": "Not Found",
			})
		} else {
			databaseErrormsg := "Failed database: "+ err.Error()
			context.JSON(http.StatusInternalServerError, gin.H{"errors": databaseErrormsg})
		}
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data": outputModel,
	})
}

func ValidateFormLoginInput(c *gin.Context) {
	var inputData model.FormLogin
	if err := c.ShouldBindJSON(&inputData); err != nil {
		validationErrorMsg := "login and password fields are required"
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrorMsg})
		return
	}
}