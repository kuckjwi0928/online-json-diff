package service

import "github.com/samber/lo"

type HttpService interface {
	MultiGet(urls []string, headers map[string]string) (interface{}, error)
}

func NewHttpService() HttpService {
	return &HttpServiceImpl{}
}

type HttpServiceImpl struct {
}

type httpChannelResult struct {
	response []byte
	err      error
}

func (h *HttpServiceImpl) MultiGet(urls []string, headers map[string]string) (interface{}, error) {
	ch := make(chan httpChannelResult, len(urls))

	for _, url := range urls {
		go func(ch chan<- httpChannelResult, url string) {
			client, err := NewHttpClient(
				WithURL(url),
				WithHeaders(headers),
			)

			if err != nil {
				ch <- httpChannelResult{
					err: err,
				}
				return
			}

			res, err := client.Get()

			if err != nil {
				ch <- httpChannelResult{
					err: err,
				}
				return
			}

			ch <- httpChannelResult{
				response: res,
				err:      nil,
			}
		}(ch, url)
	}

	results := make([]httpChannelResult, len(urls))

	for i := 0; i < len(urls); i++ {
		results[i] = <-ch
	}

	close(ch)

	for _, result := range results {
		if result.err != nil {
			return nil, result.err
		}
	}

	return lo.Map(results, func(result httpChannelResult, _ int) []byte {
		return result.response
	}), nil
}
