/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [delivery plugin]
 */
package delivery

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// index html
const index = "index.html"

type localFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func localFile(root string, indexes bool) *localFileSystem {
	return &localFileSystem{
		FileSystem: gin.Dir(root, indexes),
		indexes:    indexes,
		root:       root,
	}
}

// Exists detect the presence of files
func (local *localFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		name := path.Join(local.root, p)
		stats, err := os.Stat(name)
		if err != nil {
			return false
		}
		if stats.IsDir() {
			if !local.indexes {
				index := path.Join(name, index)
				_, err := os.Stat(index)
				if err != nil {
					return false
				}
			}
		}
		return true
	}
	return false
}
