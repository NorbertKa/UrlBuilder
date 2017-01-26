package UrlBuilder

func (u Url) Generate() []byte {
	url := make([]byte, 9)
	count := 0
	count = copy(url[count:], u.protocol)
	count = copy(url[count:], path)
	count = copy(url[count:], u.host)
	for i, p := range u.path {
		if i == 0 {
			count = copy(url[count:], p)
		} else {
			count = copy(url[count:], slash)
			count = copy(url[count:], p)
		}
	}
	initializedQuery := false
	for k, q := range u.queryParams {
		if !initializedQuery {
			count = copy(url[count:], questionMark)
			count = copy(url[count:], k)
			count = copy(url[count:], q)
			initializedQuery = true
		} else {
			count = copy(url[count:], ampersand)
			count = copy(url[count:], k)
			count = copy(url[count:], q)
		}
	}
	return url
}
