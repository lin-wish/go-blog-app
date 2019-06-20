package main

import "errors"

type article struct {
	ID      int
	Title   string
	Content string
}

var articleList = []article{
	article{ID: 1, Title: "Good Day", Content: "Today is a good day."},
	article{ID: 2, Title: "Bad Day", Content: "Today is a bad day."},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func createNewArticle(title, content string) (*article, error) {
	a := article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, a)

	return &a, nil
}
