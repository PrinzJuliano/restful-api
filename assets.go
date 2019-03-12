package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<html>\r\n<head>\r\n    <meta charset=\"utf-8\">\r\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\r\n    <title>Hello</title>\r\n    <link rel=\"stylesheet\" href=\"/assets/css/bootstrap.min.css\">\r\n</head>\r\n<body>\r\n<h1>It works!</h1>\r\n<div class=\"display-1\">Really!</div>\r\n<span>Here is some data: {{.a}}</span>\r\n<script src=\"/assets/js/jquery.min.js\"></script>\r\n<script src=\"/assets/js/popper.min.js\"></script>\r\n<script src=\"/assets/js/bootstrap.min.js\"></script>\r\n</body>\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1552397421, 1552397421099368600),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1552397438, 1552397438915038600),
		Data:     nil,
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1552397438, 1552397438901067400),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}}, "")
