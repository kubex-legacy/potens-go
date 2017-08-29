package webui

import (
	"strconv"
	"github.com/kubex/proto-go/application"
	"fmt"
	"hash/crc32"
)

//SetPageTitle set the page title on the response
func SetPageTitle(response *application.HTTPResponse, PageTitle string) {
	response.Headers["x-cube-title"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{PageTitle}}
}

//SetBackPath set the back path on the response, relative to app route
func SetBackPath(response *application.HTTPResponse, BackPath string) {
	response.Headers["x-cube-back-path"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{BackPath}}
}

//SetPageIcon set the icon url/code on the response
func SetPageIcon(response *application.HTTPResponse, Icon string) {
	response.Headers["x-cube-icon"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{Icon}}
}

//SetPageFID set the FID for the entity being shown on the page
func SetPageFID(response *application.HTTPResponse, FID string) {
	response.Headers["x-cube-page-fid"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{FID}}
}

//SetCacheSeconds set the response to cache for X seconds
func SetCacheSeconds(response *application.HTTPResponse, Seconds int64) {
	response.Headers["x-cache-seconds"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{strconv.FormatInt(Seconds, 10)}}
}

//SetCachePublic set the response to cache in public caches
func SetCachePublic(response *application.HTTPResponse) {
	response.Headers["x-cache-scope"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{"public"}}
}

//SetCacheRevalidate set the response to re-validate cache
func SetCacheRevalidate(response *application.HTTPResponse) {
	response.Headers["x-cache-revalidate"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{"must-revalidate"}}
}

//SetCacheETag set the etag for the response
func SetCacheETag(response *application.HTTPResponse, ETag string) {
	response.Headers["x-cache-etag"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{ETag}}
}

//BuildETag Create an etag for caching
func BuildETag(name string, data []byte) string {
	crc := crc32.ChecksumIEEE(data)
	return fmt.Sprintf(`"%s-%d-%08X"`, name, len(data), crc)
}
