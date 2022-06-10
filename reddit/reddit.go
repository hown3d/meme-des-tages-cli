package reddit

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Client struct {
	redditClient *reddit.Client
}

func NewClient() Client {
	redditclient := reddit.DefaultClient()
	return Client{
		redditClient: redditclient,
	}
}

const time_day = "day"

func (c Client) GetTopPost(ctx context.Context, subreddit string) (*url.URL, error) {
	posts, _, err := c.redditClient.Subreddit.TopPosts(ctx, subreddit, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{},
		Time:        time_day,
	})
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		u := post.URL
		uri, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		res, err := isPicture(uri)
		if err != nil {
			return nil, fmt.Errorf("checking if url %v is picture: %w", u, err)
		}
		if res {
			return uri, nil
		}
	}
	return nil, errors.New("no picture found on top today")
}

var pictureFileExtensions []string = []string{"jpg", "jpeg", "png", "gif", "svg"}

func isPicture(u *url.URL) (bool, error) {
	p := u.Path
	queryExt := filepath.Ext(p)
	// remove leading dot
	if queryExt != "" {
		queryExt = queryExt[1:]
	}
	for _, ext := range pictureFileExtensions {
		if queryExt == ext {
			return true, nil
		}
	}
	return false, nil
}
