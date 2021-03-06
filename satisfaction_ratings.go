package zendesk

import (
	"fmt"
	"time"

	resty "gopkg.in/resty.v0"
)

type SatisfactionRatingApiHandler struct {
	client Client
}

func (s SatisfactionRatingApiHandler) SatisfactionRatings(score string, startDate *time.Time, stopDate *time.Time) (SatisfactionRatings, error) {
	urlString := fmt.Sprintf("/satisfaction_ratings.json?start_time=%d&end_time=%d", startDate.Unix(), stopDate.Unix())

	if len(score) > 0 {
		urlString = fmt.Sprintf(urlString+"&score=%s", score)
	}

	response, err := s.client.get(
		urlString,
		nil,
	)

	if err != nil {

	}

	return s.parseResults(response), err
}

func (s SatisfactionRatingApiHandler) parseResults(response *resty.Response) SatisfactionRatings {
	var object SatisfactionRatings

	s.client.parseResponseToInterface(response, &object)

	return object
}
