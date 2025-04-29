package infrustructure

import (
	"circle/pkg/server/domain"
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type IPostInterface interface {
	GetAll(posts *[]domain.Post) error
}

type PostInfrastructure struct {
	conn *dynamodb.Client
}

func NewPostInfrastructure(conn *dynamodb.Client) IPostInterface {
	return &PostInfrastructure{
		conn: conn,
	}
}

func (pi *PostInfrastructure) GetAll(posts *[]domain.Post) error {
	input := &dynamodb.ScanInput{
		TableName: aws.String("i.maker_Users"),
	}

	resp, err := pi.conn.Scan(context.TODO(), input)
	if err != nil {
		return err
	}

	for i, item := range resp.Items {
		slog.Info("Raw item", "index", i, "item", item)

		// 'post' フィールドだけ取り出す
		rawPost, ok := item["post"]
		if !ok {
			slog.Error("Missing post field", "index", i)
			continue
		}

		// 'post' を unmarshall
		var post domain.Post
		if err := attributevalue.Unmarshal(rawPost, &post); err != nil {
			slog.Error("Unmarshal post failed", "error", err, "index", i)
			return err
		}

		slog.Info("Parsed post", "index", i, "post", post)

		// postsに追加
		*posts = append(*posts, post)
	}

	return nil
}
