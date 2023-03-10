package v1

import "github.com/gin-gonic/gin"

type Favorite struct {
}

func NewFavorite() Favorite {
	return Favorite{}
}

func (f Favorite) List(c *gin.Context) {

}

func (f Favorite) Action(c *gin.Context) {

}
