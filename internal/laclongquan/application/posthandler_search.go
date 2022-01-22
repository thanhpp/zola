package application

import "context"

func (p PostHandler) Search(ctx context.Context, requestorID, keyword string, offset, limit int) ([]*GetPostResult, error) {
	requestor, err := p.userRepo.GetByID(ctx, requestorID)
	if err != nil {
		return nil, err
	}

	postIDs, err := p.esClient.SearchPost(keyword, offset, limit)
	if err != nil {
		return nil, err
	}

	var listRes = make([]*GetPostResult, 0, len(postIDs))
	for i := range postIDs {
		postRes, err := p.GetPost(ctx, requestor.ID().String(), postIDs[i])
		if err != nil {
			continue
		}
		listRes = append(listRes, postRes)
	}

	return listRes, nil
}
