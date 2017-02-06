package main

type webhook struct {
	Before, After, Ref, User_name string
	User_id, Project_id           int
	Repository                    gitlabRepository
	Commits                       []commit
	Total_commits_count           int
}

type commit struct {
	Id, Message, Timestamp, Url string
	Author                      author
}

type author struct {
	Name, Email string
}

type gitlabRepository struct {
	Name, Url, Description, Home string
}
