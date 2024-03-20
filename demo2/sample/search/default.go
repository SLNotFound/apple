package search

type defaultMatcher struct {
}

func init() {
	var matcher Matcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
