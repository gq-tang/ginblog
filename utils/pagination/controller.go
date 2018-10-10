package pagination

import (
	"github.com/gin-gonic/gin"
)

func SetPaginator(ctx *gin.Context, per int, nums interface{}) (paginator *Paginator) {
	paginator = NewPaginator(ctx.Request, per, nums)
	ctx.Set("paginator", paginator)
	return
}
