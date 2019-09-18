module git.intra.weibo.com/adx/logging

require (
	github.com/labstack/gommon v0.2.8
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v0.0.0-20170224212429-dcecefd839c4 // indirect
)

replace (
	golang.org/x/crypto v0.0.0-20181106171534-e4dc69e5b2fd => github.com/golang/crypto v0.0.0-20181106171534-e4dc69e5b2fd
	golang.org/x/sys v0.0.0-20181107165924-66b7b1311ac8 => github.com/golang/sys v0.0.0-20181107165924-66b7b1311ac8
)
