package main

import (
	"fmt"
	"net/http"
)

func kittyHandler(c *gin.Context) {
	cats = []string{
		"https://pbs.twimg.com/media/CZwt9t1WYAELLfX.jpg",
		"https://top13.net/wp-content/uploads/2015/10/perfectly-timed-cat-photos-funny-fb.jpg",
		"https://fthmb.tqn.com/cD0PNhMM0BxevlBvAgD1ntpQLac=/3558x2363/filters:fill(auto,1)/Cat-rolling-GettyImages-165893132-58ac5ef05f9b58a3c90a144f.jpg",
		"http://1.bp.blogspot.com/-cnxXgWn146M/UAcvCGngvDI/AAAAAAAAA9s/yCzePisIlUM/s1600/catss.jpg",
		"https://i.ytimg.com/vi/7OD55d6iRCQ/maxresdefault.jpg",
	}
	randCat := randStr(cats)
	return c.String(http.StatusOK, fmt.Sprintf(`<html>
	<head></head>
	<body>
	<img src="%s" />
	</body>
	</html>`, randCat))
}
