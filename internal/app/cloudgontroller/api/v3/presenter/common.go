package presenter

type Metadata struct {
	Labels      interface{} `json:"labels"`
	Annotations interface{} `json:"annotations"`
}

type Pagination struct {
	TotalResults int  `json:"total_results"`
	TotalPages   int  `json:"total_pages"`
	First        Link `json:"first"`
	Last         Link `json:"last"`
	Next         Link `json:"next"`
	Previous     Link `json:"previous"`
}

type Link struct {
	Href   string `json:"href"`
	Method string `json:"method,omitempty"`
}
