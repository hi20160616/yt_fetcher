module github.com/hi20160616/yt_fetcher

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/hi20160616/exhtml v0.0.0-20210202082235-6cd27713f638
	github.com/kkdai/youtube/v2 v2.5.0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20210316092652-d523dce5a7f4 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/sys v0.0.0-20210316164454-77fc1eacc6aa // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200605160147-a5ece683394c // indirect
)

replace github.com/kkdai/youtube/v2 v2.5.0 => github.com/hi20160616/youtube/v2 v2.5.1
