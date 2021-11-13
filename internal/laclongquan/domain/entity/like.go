package entity

type Like struct {
	PostID  string
	Creator string
}

type LikeFactory interface {
	NewLike(post *Post, creator string) (*Like, error)
}

func NewLikeFactory() LikeFactory {
	return &likeFactoryImpl{}
}

type likeFactoryImpl struct{}

func (fac likeFactoryImpl) NewLike(post *Post, creator string) (*Like, error) {
	if post.IsLocked() {
		return nil, ErrLockedPost
	}

	return &Like{
		PostID:  post.ID(),
		Creator: creator,
	}, nil
}
