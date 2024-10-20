package graph

import (
	"context"
	"crypto/rand"
	"math/big"
	mongo "root/mongo"
	"root/graph/model"
)

var (
	db = mongo.ConnectDB()
)

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	message := &model.Message{
		ID:   randNumber.String(),
		Text: input.Text,
	}

	// Вызов функции для вставки сообщения в базу данных
	if err := db.InsertMessage(message); err != nil {
		return nil, err // Возвращаем ошибку, если вставка не удалась
	}

	return message, nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return r.messages, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
