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
