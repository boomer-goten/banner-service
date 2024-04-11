package model

type ValueCache struct {
	IsActive bool
	Banner   ContentBanner
}

type ContentBanner struct {
	Content map[string]interface{} `json:"content,omitempty"`
}
