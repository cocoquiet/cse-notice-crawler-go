package notice

import "fmt"

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

func ToDict(n *Notice) string {
	return fmt.Sprintf("{'num': %d, 'link': '%s', 'title': '%s', 'category': '%s', content: '%s', 'created_at': '%s'}", n.num, n.link, n.title, n.category, n.content, n.createdAt)
}
