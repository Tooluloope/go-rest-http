package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/tooluloope/go-rest-http/internal/comment"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(cmtRow CommentRow) comment.Comment {
	return comment.Comment{
		ID:     cmtRow.ID,
		Slug:   cmtRow.Slug.String,
		Body:   cmtRow.Body.String,
		Author: cmtRow.Author.String,
	}
}

func (d *Database) GetComment(
	ctx context.Context,
	uuid string,
) (comment.Comment, error) {
	var cmtRow CommentRow

	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author FROM comments WHERE id = $1`,
		uuid,
	)

	err := row.Scan(
		&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}

	return convertCommentRowToComment(
		cmtRow,
	), nil
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.New().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	// rows, err := d.Client.QueryContext(
	// 	ctx,
	// 	`INSERT INTO comments (id, slug, body, author) VALUES (:ID, :Slug, :Body, :Author)`, postRow,
	// )

	rows, err := d.Client.QueryContext(
		ctx,
		`INSERT INTO comments (id, slug, body, author) VALUES ($1, $2, $3, $4)`,
		postRow.ID, postRow.Slug, postRow.Body, postRow.Author,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error inserting comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing rows: %w", err)
	}

	return cmt, nil
}

func (d *Database) DeleteComment(ctx context.Context, uuid string) error {
	_, err := d.Client.QueryContext(
		ctx,
		`DELETE FROM comments WHERE id = $1`,
		uuid,
	)

	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}

	return nil

}

func (d *Database) UpdateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {

	cmtRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	rows, err := d.Client.QueryContext(
		ctx,
		`UPDATE comments SET slug = $1, body = $2, author = $3 WHERE id = $4 RETURNING id, slug, body, author`,
		cmtRow.Slug, cmtRow.Body, cmtRow.Author, cmtRow.ID,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error updating comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing rows: %w", err)
	}

	return convertCommentRowToComment(
		cmtRow,
	), nil
}
