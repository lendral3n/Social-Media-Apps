package postings

import (
	"BE-Sosmed/features/comments"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type Posting struct {
	ID       uint
	Artikel  string
	Gambar   string
	UserID   uint
	Username string
	Image    string
}

type Handler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByPostID() echo.HandlerFunc
	GetByUsername() echo.HandlerFunc
}

type Service interface {
	TambahPosting(token *jwt.Token, newPosting Posting) (Posting, error)
	SemuaPosting() ([]Posting, error)
	AmbilComment(PostID uint) ([]comments.Comment, error)
	UpdatePosting(token *jwt.Token, updatePosting Posting) (Posting, error)
	DeletePosting(token *jwt.Token, postID uint) error
	AmbilPostingByPostID(PostID uint) (Posting, error)
	AmbilPostingByUsername(Username string) ([]Posting, error)
}

type Repository interface {
	InsertPosting(userID uint, newPosting Posting) (Posting, error)
	GetAllPost() ([]Posting, error)
	GetComment(PostID uint) ([]comments.Comment, error)
	UpdatePost(userID uint, updatePosting Posting) (Posting, error)
	DeletePost(userID uint, postID uint) error
	GetPostByPostID(PostID uint) (Posting, error)
	GetPostByUsername(Username string) ([]Posting, error)
}
