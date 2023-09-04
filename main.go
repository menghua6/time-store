package main

import(
	"github.com/gin-gonic/gin"
)
func main()  {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		response := "" +
		"  | 09/04 | 09/05 | 09/06 | 09/07 | 09/08 | 09/09 | 09.10 | \n" +
		"00|       |       |       |       |       |       |       | \n" +
		"01|       |       |       |       |       |       |       | \n" +
		"02|       |       |       |       |       |       |       | \n" +
		"03|       |       |       |       |       |       |       | \n" +
		"04|       |       |       |       |       |       |       | \n" +
		"05|       |       |       |       |       |       |       | \n" +
		"06|       |       |       |       |       |       |       | \n" +
		"07|       |       |       |       |       |       |       | \n" +
		"08|       |       |       |       |       |       |       | \n" +
		"09|       |       |       |       |       |       |       | \n" +
		"10|   *   |   *   |   *   |   *   |   *   |       |       | \n" +
		"11|   *   |   *   |   *   |   *   |   *   |       |       | \n" +
		"12|   *   |   *   |   *   |   *   |   *   |       |       | \n" +
		"13|   *   |   *   |   *   |   *   |   *   |       |       | \n" +
		"14|   *   |   *   |   *   |   *   |   *   |       |       | \n" +
		"15|   *   |   *   |   *   |   *   |   *   |       |       | \n" +
		"16|   *   |   *   |   *   |   *   |   *   |       |       | \n" +
		"17|   *   |   *   |   *   |   *   |   *   |       |       | \n" +
		"18|   *   |   *   |   *   |   *   |   *   |       |       | \n" +
		"19|       |       |       |       |   *   |       |       | \n" +
		"20|       |       |       |       |   *   |       |       | \n" +
		"21|       |       |       |       |       |       |       | \n" +
		"22|       |       |       |       |       |       |       | \n" +
		"23|       |       |       |       |       |       |       | \n"
		context.String(200, response)
	})

	r.Run(":8081")
}