package service

import (
	"errors"
)

type HttpService interface {
	MultiRequest(urls []string, method string, body []byte, headers map[string]string) map[string]HttpChannelResult
}

func NewHttpService() HttpService {
	return &HttpServiceImpl{}
}

type HttpServiceImpl struct {
}

type HttpChannelResult struct {
	url      string
	Response []byte
	Err      error
}

func (h *HttpServiceImpl) MultiRequest(urls []string, method string, body []byte, headers map[string]string) map[string]HttpChannelResult {
	ch := make(chan HttpChannelResult, len(urls))
	client, _ := NewHttpClient(
		WithHeaders(headers),
	)
	for _, url := range urls {
		go func(ch chan<- HttpChannelResult, url string) {
			var (
				res []byte
				err error
			)

			if method == "GET" {
				res, err = client.Get(url)
			} else if method == "POST" {
				res, err = client.Post(url, body)
			} else {
				ch <- HttpChannelResult{
					url: url,
					Err: errors.New("unsupported method"),
				}
				return
			}

			if err != nil {
				ch <- HttpChannelResult{
					url: url,
					Err: err,
				}
				return
			}

			ch <- HttpChannelResult{
				url:      url,
				Response: res,
				Err:      nil,
			}
		}(ch, url)
	}

	results := make(map[string]HttpChannelResult)

	for i := 0; i < len(urls); i++ {
		channelResult := <-ch
		results[channelResult.url] = channelResult
	}

	close(ch)

	return results
}
