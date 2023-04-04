package main

import "fmt"

type author struct {
	firstName string
	lastName string
	role string
	bio string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

type website struct {
	blogPostList []blogPost
}

func (b blogPost) details(){
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.fullName())
	fmt.Println("Bio: ", b.bio);
}

func (w website) printList(){
	for _, blog := range w.blogPostList {
		blog.details();
		fmt.Println();
	}
}

func main(){
	// Create the author
	newAuthor := author{
		"Hemant", "Soni", "Golang Developer", "Loves the the golang project",
	}

	// Create the blog
	newBlog := blogPost{
		"Composition in golang",
		"Go does not have inheritance, but have composition",
		newAuthor,
	}
	newBlog1 := blogPost{
		"Composition in golang",
		"Go does not have inheritance, but have composition",
		newAuthor,
	}
	newBlog2 := blogPost{
		"Composition in golang",
		"Go does not have inheritance, but have composition",
		newAuthor,
	}

	w := website{
		blogPostList: []blogPost{newBlog, newBlog1, newBlog2},
	}

	w.printList();
}