package paths

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Declare_Primary_Paths(input *gin.Engine) {

	///ILK API UCU DENEMESI
	input.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "TEBRİKLER SUNUCU AÇILDI",
		})
	})

	///TOPLAMA İŞLEMİ
	//DİNAMİK DEĞİL
	input.GET("/sumofelements/:number1/:number2", func(c *gin.Context) {

		number1 := c.Param("number1")
		number2 := c.Param("number2")

		var i int
		var j int

		i, err1 := strconv.Atoi(number1)
		if err1 != nil {
			// ... handle error
			panic(err1)
		}
		j, err2 := strconv.Atoi(number2)
		if err2 != nil {
			// ... handle error
			panic(err2)
		}

		sum := i + j

		c.JSON(http.StatusOK, gin.H{
			"message": sum,
		})
	})
}
