package responses

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judascrow/go-api-starter/api/services"
)

var empty map[string]interface{}

func JSON(c *gin.Context, statusCode int, data interface{}, message interface{}) {
	c.JSON(statusCode, gin.H{
		"status":     statusCode,
		"statusText": http.StatusText(statusCode),
		"data":       data,
		"message":    message,
	})

}

func JSONLIST(c *gin.Context, statusCode int, data interface{}, p services.PageMeta) {

	pageSize := p.PageSize
	page := p.Page
	count := p.Count

	pageMeta := map[string]interface{}{}
	pageMeta["offset"] = (page - 1) * pageSize
	pageMeta["requestedPageSize"] = pageSize
	pageMeta["currentPageNumber"] = page
	pageMeta["currentItemsCount"] = count

	pageMeta["prevPageNumber"] = 1
	totalPagesCount := int(math.Ceil(float64(count) / float64(pageSize)))
	pageMeta["totalPagesCount"] = totalPagesCount

	if page < totalPagesCount {
		pageMeta["hasNextPage"] = true
		pageMeta["nextPageNumber"] = page + 1
	} else {
		pageMeta["hasNextPage"] = false
		pageMeta["nextPageNumber"] = 1
	}
	if page > 1 {
		pageMeta["prevPageNumber"] = page - 1
	} else {
		pageMeta["hasPrevPage"] = false
		pageMeta["prevPageNumber"] = 1
	}

	pageMeta["nextPageUrl"] = fmt.Sprintf("%v?page=%d&pageSize=%d", c.Request.URL.Path, pageMeta["nextPageNumber"], pageMeta["requestedPageSize"])
	pageMeta["prevPageUrl"] = fmt.Sprintf("%s?page=%d&pageSize=%d", c.Request.URL.Path, pageMeta["prevPageNumber"], pageMeta["requestedPageSize"])

	c.JSON(statusCode, gin.H{
		"status":     statusCode,
		"statusText": http.StatusText(statusCode),
		"data":       data,
		"pageMeta":   pageMeta,
		"message":    empty,
	})

}

func ERROR(c *gin.Context, statusCode int, errors interface{}) {

	c.JSON(statusCode, gin.H{
		"status":     statusCode,
		"statusText": http.StatusText(statusCode),
		"data":       empty,
		"message":    errors,
	})
}