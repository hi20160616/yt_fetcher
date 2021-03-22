package data

import (
	"database/sql"
	"fmt"
	"html"
	"net/url"
	"regexp"
	"time"

	"github.com/hi20160616/exhtml"
	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	youtube "github.com/kkdai/youtube/v2"
	"github.com/pkg/errors"
)

func GetChannelFromSource(c *pb.Channel) (*pb.Channel, error) {
	return getChannelFromSource(c)
}

func getChannelFromSource(c *pb.Channel) (*pb.Channel, error) {
	if c.Id == "" || c.Id == "cid" {
		return nil, errors.New("nil or wrong channel id: " + c.Id)
	}
	// https://www.youtube.com/channel/UCMUnInmOkrWN4gof9KlhNmQ/videos
	u, err := url.Parse("https://www.youtube.com/channel/" + c.Id + "/videos")
	raw, _, err := exhtml.GetRawAndDoc(u, 1*time.Minute)
	if err != nil {
		return nil, err
	}
	// get channel name
	// c.Name = exhtml.ElementsByTag(doc, "title")[0].FirstChild.Data
	// c.Name = strings.Replace(c.Name, " - YouTube", "", -1)
	re := regexp.MustCompile(`<title>(.*) - YouTube</title>`)
	rs := re.FindAllSubmatch(raw, -1)
	if len(rs) == 0 {
		return nil, errors.New("getChannelFromSource get channel name match nothing:" + c.Id)
	}
	c.Name = html.UnescapeString(string(rs[0][1]))
	// get vids
	re = regexp.MustCompile(`"gridVideoRenderer":{"videoId":"(.*?)","thumbnail":{"thumbnails"`)
	rs = re.FindAllSubmatch(raw, -1)
	if len(rs) == 0 {
		return nil, errors.New("getChannelFromSource get vids match nothing")
	}
	for _, r := range rs {
		c.Vids = append(c.Vids, string(r[1]))
	}
	return c, nil
}

// getVidsFromSource get vids from raw html page.
func getVidsFromSource(cid string) ([]string, error) {
	vids := []string{}
	u, err := url.Parse("https://www.youtube.com/channel/" + cid + "/videos")
	raw, _, err := exhtml.GetRawAndDoc(u, 1*time.Minute)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`"gridVideoRenderer":{"videoId":"(.*?)","thumbnail":{"thumbnails"`)
	rs := re.FindAllSubmatch(raw, -1)
	for _, r := range rs {
		vids = append(vids, string(r[1]))
	}
	return vids, nil
}

// getCidFromSource get cid from youtube video api source
func getCidFromSource(vid string) (string, error) {
	video := "http://youtube.com/get_video_info?video_id=" + vid
	u, err := url.Parse(video)
	if err != nil {
		return "", err
	}
	raw, _, err := exhtml.GetRawAndDoc(u, 1*time.Minute)
	if err != nil {
		return "", err
	}
	r, err := url.QueryUnescape(string(raw))
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`","externalChannelId":"(.*?)","availableCountries":`)
	rs := re.FindAllStringSubmatch(r, -1)
	return rs[0][1], nil
}

func getVideoFromApi(dc *sql.DB, vid string) (*pb.Video, error) {
	v := &pb.Video{Id: vid}
	client := youtube.Client{}
	// TODO: get thumbnail info
	video, err := client.GetVideo("https://www.youtube.com/watch?v=" + v.Id)
	if err != nil {
		return nil, errors.WithMessage(err, "getVideoFromApi error on videoId: "+v.Id)
	}
	cid, err := getCid(dc, vid, true)
	if err != nil {
		return nil, errors.WithMessage(err, "getVideoFromApi error on videoId: "+v.Id)
	}
	t := getLastModified(video)
	v.Title = video.Title
	for _, thumbnail := range video.Thumbnails {
		w, h := int32(thumbnail.Width), int32(thumbnail.Height)
		v.Thumbnails = append(v.Thumbnails,
			&pb.Thumbnail{
				Id:     fmt.Sprintf("%s_w%d", v.Id, w),
				Width:  w,
				Height: h,
				URL:    thumbnail.URL,
				Vid:    v.Id,
			})
	}
	v.Description = video.Description
	v.Duration = video.Duration.String()
	v.Cid = cid
	v.Cname = video.Author
	v.LastUpdated = t
	return v, nil
}

func getLastModified(v *youtube.Video) (rt string) {
	fs := v.Formats
	for _, f := range fs {
		if f.LastModified != "" {
			return f.LastModified
		}
	}
	return time.Nanosecond.String()[:16]
}
