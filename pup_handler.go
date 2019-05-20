package main

import (
	"fmt"
	"net/http"
)

func pupHandler(c *gin.Context) {
	dogs := []string{
		"http://static3.businessinsider.com/image/5484d9d1eab8ea3017b17e29/9-science-backed-reasons-to-own-a-dog.jpg",
		"https://yt3.ggpht.com/a-/AAuE7mADdvvvDy2jJr4ajyHsv0Q26gTURdLpWQbJgg=s900-mo-c-c0xffffffff-rj-k-no",
		"https://s7d1.scene7.com/is/image/PETCO/dog-category-090617-369w-269h-hero-cutout-d?fmt=png-alpha",
		"https://mymodernmet.com/wp/wp-content/uploads/2018/11/dogs-catching-treats-christian-vieler-17.jpg",
		"https://pngimg.com/uploads/dog/dog_PNG50382.png",
		"https://4.bp.blogspot.com/-5whUDkuwejU/T89th9wreqI/AAAAAAAI79E/NYbhjBSrR2g/s1600/crazy+dog+pictures+%2826%29.jpg",
		"https://media1.s-nbcnews.com/j/newscms/2016_37/1157735/dog-halloween-costume-turtle-today-16-09-13_3948f79ce3b0c6eb795493c5a671ff13.today-inline-large.png",
	}
	randDog := randStr(dogs)
	return c.String(http.StatusOK, fmt.Sprintf(`<html>
	<head></head>
	<body>
	<img src="%s" />
	</body>
	</html>`, randDog))
}
