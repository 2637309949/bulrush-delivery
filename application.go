/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [delivery plugin]
 */
// All non-path-specified requests that cannot be intercepted in the subpath
// are handled as middleware, and all requests that need to be intercepted
// need to be implemented on the httpProxy instance

package delivery

import (
	"net/http"

	"github.com/2637309949/bulrush"
	"github.com/gin-gonic/gin"
)

// Delivery service static files
type Delivery struct {
	bulrush.PNBase
	Path      string
	URLPrefix string
}

// Plugin for gin
func (delivery *Delivery) Plugin() bulrush.PNRet {
	return func(httpProxy *gin.Engine) {
		lf := localFile(delivery.Path, false)
		fileserver := http.FileServer(lf)
		if delivery.URLPrefix != "" {
			fileserver = http.StripPrefix(delivery.URLPrefix, fileserver)
		}
		httpProxy.GET(delivery.URLPrefix+"/*any", func(c *gin.Context) {
			if lf.Exists(delivery.URLPrefix, c.Request.URL.Path) {
				fileserver.ServeHTTP(c.Writer, c.Request)
				c.Abort()
			}
		})
	}
}
