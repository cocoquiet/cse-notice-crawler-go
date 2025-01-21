package notice

type Notice struct {
	num       int
	link      string
	title     string
	category  string
	content   string
	createdAt string
}

func NewNotice(num int, link, title, category, content, createdAt string) *Notice {
	return &Notice{
		num:       num,
		link:      link,
		title:     title,
		category:  category,
		content:   content,
		createdAt: createdAt,
	}
}

func ToDict(n *Notice) map[string]interface{} {
	return map[string]interface{}{
		"num":       n.num,
		"link":      n.link,
		"title":     n.title,
		"category":  n.category,
		"content":   n.content,
		"createdAt": n.createdAt,
	}
}
