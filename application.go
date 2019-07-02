// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// All non-path-specified requests that cannot be intercepted in the subpath
// are handled as middleware, and all requests that need to be intercepted
// need to be implemented on the httpProxy instance

package delivery

import (
	"net/http"
	"path"

	"github.com/2637309949/bulrush"
	"github.com/gin-gonic/gin"
)

// Delivery service static files
type Delivery struct {
	bulrush.PNBase
	Path      string
	URLPrefix string
}

// New return Delivery with default property
func New() *Delivery {
	del := &Delivery{
		URLPrefix: "/public",
		Path:      path.Join("assets/public", ""),
	}
	return del
}

// Plugin for gin
func (delivery *Delivery) Plugin() interface{} {
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
