package agents

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// Crawler contains information about one crawler.
type Crawler struct {
	// Regexp of User Agent of the crawler.
	Pattern string `json:"pattern"`

	// Official url of the robot.
	URL string `json:"url"`

	// Examples of full User Agent strings.
	Instances []string `json:"instances"`
}

// Private time needed to convert addition_date from/to the format used in JSON.
type jsonCrawler struct {
	Pattern      string   `json:"pattern"`
	AdditionDate string   `json:"addition_date"` //nolint:tagliatelle
	URL          string   `json:"url"`
	Instances    []string `json:"instances"`
}

// MarshalJSON implements json.Marshaler interface.
func (c *Crawler) MarshalJSON() ([]byte, error) {
	jc := jsonCrawler{
		Pattern:   c.Pattern,
		URL:       c.URL,
		Instances: c.Instances,
	}
	return json.Marshal(jc)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (c *Crawler) UnmarshalJSON(b []byte) error {
	var jc jsonCrawler
	if err := json.Unmarshal(b, &jc); err != nil {
		return err
	}

	c.Pattern = jc.Pattern
	c.URL = jc.URL
	c.Instances = jc.Instances

	if c.Pattern == "" {
		return fmt.Errorf("empty pattern in record %s", string(b))
	}

	return nil
}

// Crawlers the list of crawlers, built from contents of crawler-user-agents.json.
var Crawlers = func() []Crawler { //nolint:gochecknoglobals
	var crawlers []Crawler
	if err := json.Unmarshal([]byte(crawlersJSON), &crawlers); err != nil {
		panic(err)
	}
	return crawlers
}()

var regexps = func() []*regexp.Regexp { //nolint:gochecknoglobals
	rgx := make([]*regexp.Regexp, len(Crawlers))
	for i, crawler := range Crawlers {
		rgx[i] = regexp.MustCompile(crawler.Pattern)
	}
	return rgx
}()

// IsCrawler returns if User Agent string matches any of crawler patterns.
func IsCrawler(userAgent string) bool {
	for _, re := range regexps {
		if re.MatchString(userAgent) {
			return true
		}
	}
	return false
}

// MatchingCrawlers finds all crawlers matching the User Agent and returns the list of their indices in Crawlers.
func MatchingCrawlers(userAgent string) []int {
	indices := []int{}
	for i, re := range regexps {
		if re.MatchString(userAgent) {
			indices = append(indices, i)
		}
	}
	return indices
}
