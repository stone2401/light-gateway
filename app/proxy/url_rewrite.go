package proxy

import (
	"regexp"
	"strings"
	"sync"

	"github.com/stone2401/light-gateway-kernel/pcore"
)

type UrlRewrite struct {
	rw       sync.RWMutex
	rewrites map[string]rewrite
}

type rewrite struct {
	urlRegexp   *regexp.Regexp
	fromRewrite string
}

func NewUrlRewrite() *UrlRewrite {
	return &UrlRewrite{
		rw:       sync.RWMutex{},
		rewrites: map[string]rewrite{},
	}
}

func (u *UrlRewrite) Add(rewriteStr string) {
	u.rw.Lock()
	defer u.rw.Unlock()
	rewriteArr := strings.Split(rewriteStr, "\n")
	for _, rewrites := range rewriteArr {
		if rewrites == "" {
			continue
		}
		toRewrite, fromRewrite := strings.Split(rewrites, " ")[0], strings.Split(rewrites, " ")[1]
		u.rewrites[toRewrite] = rewrite{
			urlRegexp:   regexp.MustCompile(toRewrite),
			fromRewrite: fromRewrite,
		}
	}
}

func (u *UrlRewrite) RewriteHandler(ctx *pcore.Context) {
	u.rw.RLock()
	defer u.rw.RUnlock()
	for _, rewrite := range u.rewrites {
		if rewrite.urlRegexp.MatchString(ctx.Request.URL.Path) {
			ctx.Request.URL.Path = rewrite.urlRegexp.ReplaceAllString(ctx.Request.URL.Path, rewrite.fromRewrite)
			break
		}
	}
	ctx.Next()
}

func (u *UrlRewrite) Reset() {
	u.rw.Lock()
	defer u.rw.Unlock()
	u.rewrites = map[string]rewrite{}

}
