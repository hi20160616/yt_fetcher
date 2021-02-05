# Youtube Fetcher


# gRPC
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
api/yt_fetcher/api/server.proto
```

# Refer

https://hkgoldenmra.blogspot.com/2013/05/youtube.html

要获取 http://www.youtube.com/watch?v=fEcnrA6RIzo 的信息:

访问: http://www.youtube.com/get_video_info?video_id=fEcnrA6RIzo

然后通过 URL decode online 网站得到具体信息：

- `hl` 為預設語言  
- `author` 為影片上載者  
- `iurlsd` 為封面圖片  
- `thumbnail_url` 為封面縮圖  
- `length_seconds` 為影片長度，以秒計算  
- `title` 為影片標題  
- `url_encoded_fmt_stream_map` 為另一串 query string 保存著影片的來源資訊，而來源資訊以 `,` 分類再將 `url_encoded_fmt_stream_map` 拆解  

quality 為影片品質，分別有：  
- smail 為 240p  
- medium 為 360p  
- large 為 480p
- hd720 為 720p
- hd 1080 為 1080p

sig 為用作許可影片播放的「簽名」  

type 為影片類型，分別有：  

- video/3gpp 為 3gp 格式
- video/mp4 為 mp4 格式
- video/webm 為 webm 格式
- video/x-flv 為 flv 格式
- url 為影片來源，都是一種 query string

要下載一個 Youtube 影片，需要將 url 及 sig 以 signature 連接才能夠下載
即 `<url>&signature=<sig>` 的超連結
