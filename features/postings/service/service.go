package service

import (
	"BE-Sosmed/features/comments"
	"BE-Sosmed/features/postings"
	"BE-Sosmed/features/users"
	"BE-Sosmed/helper/jwt"
	"errors"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type PostingService struct {
	m    postings.Repository
	user users.Service
}

func New(model postings.Repository, user users.Service) postings.Service {
	return &PostingService{
		m:    model,
		user: user,
	}
}

func (ps *PostingService) TambahPosting(token *golangjwt.Token, newPosting postings.Posting) (postings.Posting, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return postings.Posting{}, err
	}

	result, err := ps.m.InsertPosting(userID, newPosting)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return postings.Posting{}, errors.New("posting sudah ada pada sistem")
		}
		return postings.Posting{}, errors.New("terjadi kesalahan server")
	}

	return result, nil
}

func (ps *PostingService) AmbilComment(PostID uint) ([]comments.Comment, error) {
	result, err := ps.m.GetComment(PostID)

	if err != nil {
		return nil, errors.New("terjadi kesalahan server")
	}

	for i, post := range result {
		user, err := ps.user.GetUserById(post.UserID)
		if err != nil {
			return nil, err
		}

		result[i].Username = user.Username
		result[i].Image = user.Image

	}

	return result, nil
}

func (ps *PostingService) SemuaPosting(page int64, pageSize int64) ([]postings.Posting, postings.Pagination, error) {
	posts, pagination, err := ps.m.GetAllPost(page, pageSize)

	if err != nil {
		return nil, postings.Pagination{}, err
	}

	for i, post := range posts {
		user, err := ps.user.GetUserById(post.UserID)
		if err != nil {
			return nil, postings.Pagination{}, err
		}

		posts[i].Username = user.Username
		posts[i].Image = user.Image

	}

	return posts, pagination, nil
}

func (ps *PostingService) UpdatePosting(token *golangjwt.Token, updatePosting postings.Posting) (postings.Posting, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return postings.Posting{}, err
	}

	updatedPost, err := ps.m.UpdatePost(userID, updatePosting)
	if err != nil {
		return postings.Posting{}, err
	}

	return updatedPost, nil
}

func (ps *PostingService) DeletePosting(token *golangjwt.Token, postID uint) error {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return err
	}

	err = ps.m.DeletePost(userID, postID)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PostingService) AmbilPostingByPostID(PostID uint) (postings.Posting, error) {
	post, err := ps.m.GetPostByPostID(PostID)

	if err != nil {
		return postings.Posting{}, err
	}

	user, err := ps.user.GetUserById(post.UserID)
	if err != nil {
		return postings.Posting{}, err
	}

	post.Username = user.Username
	post.Image = user.Image

	return post, nil
}

func (ps *PostingService) AmbilPostingByUsername(Username string) ([]postings.Posting, error) {
	posts, err := ps.m.GetPostByUsername(Username)

	if err != nil {
		return nil, err
	}

	for i, post := range posts {
		user, err := ps.user.GetUserById(post.UserID)
		if err != nil {
			return nil, err
		}

		posts[i].Username = user.Username
		posts[i].Image = user.Image

	}

	return posts, nil
}

func (ps *PostingService) LikePosting(postID uint) error {
	return ps.m.LikePost(postID)
}
