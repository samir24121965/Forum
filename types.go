package main

type User struct {
	Id          int64
	Name        string
	Email       string
	Password    string
	Is_disabled int
	Is_deleted  int
	Created_at  string
	Updated_at  string
}

type Post struct {
	Id          int64
	User_id     int
	Title       string
	Content     string
	Nb_likes    int
	Nb_dislikes int
	Created_at  string
	Updated_at  string
}

type Category struct {
	Id         int64
	Name       string
	Nb_posts   int
	Created_at string
	Updated_at string
}

type Post_category struct {
	Post_id     int64
	Category_id int64
}

type Comment struct {
	Id          int64
	User_id     int64
	Post_id     int64
	Parent_id   int64
	Content     string
	Nb_likes    int
	Nb_dislikes int
	Is_deleted  int
	Created_at  string
	Updated_at  string
}

type Post_reaction struct {
	Id         int64
	User_id    int64
	Post_id    int64
	Content    string
	Is_liked   int
	Created_at string
}

type Comment_reaction struct {
	Id         int64
	User_id    int64
	Comment_id int64
	Content    string
	Is_liked   int
	Created_at string
}
