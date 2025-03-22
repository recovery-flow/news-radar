package neodb

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Authorship interface {
	Create(ctx context.Context, articleID uuid.UUID, authorID uuid.UUID) error
	Delete(ctx context.Context, articleID uuid.UUID, authorID uuid.UUID) error

	SetForArticle(ctx context.Context, ID uuid.UUID, author []uuid.UUID) error
	GetForArticle(ctx context.Context, ID uuid.UUID) ([]uuid.UUID, error)

	GetForAuthor(ctx context.Context, ID uuid.UUID) ([]uuid.UUID, error)
}

type authorship struct {
	driver neo4j.Driver
}

func NewAuthorship(uri, username, password string) (Authorship, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, fmt.Errorf("failed to create neo4j driver: %w", err)
	}

	if err = driver.VerifyConnectivity(); err != nil {
		return nil, fmt.Errorf("failed to verify connectivity: %w", err)
	}

	return &authorship{
		driver: driver,
	}, nil
}

func (a *authorship) Create(ctx context.Context, articleID uuid.UUID, authorID uuid.UUID) error {
	session, err := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		cypher := `
            MATCH (art:Article { id: $articleID })
            MATCH (auth:Author { id: $authorID })
            MERGE (art)-[:AUTHORED_BY]->(auth)
        `
		params := map[string]any{
			"articleID": articleID.String(),
			"authorID":  authorID.String(),
		}
		_, err := tx.Run(cypher, params)
		if err != nil {
			return nil, fmt.Errorf("failed to create authorship relationship: %w", err)
		}
		return nil, nil
	})
	return err
}

func (a *authorship) Delete(ctx context.Context, articleID uuid.UUID, authorID uuid.UUID) error {
	session, err := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		cypher := `
			MATCH (art:Article { id: $articleID })-[r:AUTHORED_BY]->(auth:Author { id: $authorID })
			DELETE r
		`
		params := map[string]any{
			"articleID": articleID.String(),
			"authorID":  authorID,
		}

		_, err := tx.Run(cypher, params)
		if err != nil {
			return nil, fmt.Errorf("failed to delete AUTHOR relationship: %w", err)
		}
		return nil, nil
	})

	return err
}

func (a *authorship) SetForArticle(ctx context.Context, articleID uuid.UUID, authors []uuid.UUID) error {
	session, err := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		deleteCypher := `
			MATCH (a:Article { id: $id })-[r:AUTHORED_BY]->(:Author)
			DELETE r
		`
		params := map[string]any{
			"id": articleID.String(),
		}
		_, err := tx.Run(deleteCypher, params)
		if err != nil {
			return nil, fmt.Errorf("failed to delete existing authorship relationships: %w", err)
		}

		authorIDs := make([]string, len(authors))
		for i, authID := range authors {
			authorIDs[i] = authID.String()
		}

		createCypher := `
			MATCH (a:Article { id: $id })
			FOREACH (authorId IN $authors |
				MATCH (au:Author { id: authorId })
				MERGE (a)-[:AUTHORED_BY]->(au)
			)
		`
		params["authors"] = authorIDs

		_, err = tx.Run(createCypher, params)
		if err != nil {
			return nil, fmt.Errorf("failed to create new authorship relationships: %w", err)
		}

		return nil, nil
	})
	return err
}

func (a *authorship) GetForArticle(ctx context.Context, articleID uuid.UUID) ([]uuid.UUID, error) {
	session, err := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	if err != nil {
		return nil, err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		cypher := `
			MATCH (a:Article { id: $id })-[:AUTHORED_BY]->(au:Author)
			RETURN au.id AS authorID
		`
		params := map[string]any{
			"id": articleID.String(),
		}
		records, err := tx.Run(cypher, params)
		if err != nil {
			return nil, err
		}
		var authorIDs []uuid.UUID
		for records.Next() {
			record := records.Record()
			authorIDVal, ok := record.Get("authorID")
			if !ok {
				continue
			}

			if idStr, ok := authorIDVal.(string); ok {
				uid, err := uuid.Parse(idStr)
				if err != nil {
					continue
				}
				authorIDs = append(authorIDs, uid)
			}
		}
		return authorIDs, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]uuid.UUID), nil
}

func (a *authorship) GetForAuthor(ctx context.Context, authorID uuid.UUID) ([]uuid.UUID, error) {
	session, err := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	if err != nil {
		return nil, err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		cypher := `
			MATCH (au:Author { id: $id })<-[:AUTHORED_BY]-(art:Article)
			RETURN art.id AS articleID
		`
		params := map[string]any{
			"id": authorID.String(),
		}
		records, err := tx.Run(cypher, params)
		if err != nil {
			return nil, err
		}
		var articleIDs []uuid.UUID
		for records.Next() {
			record := records.Record()
			articleIDVal, ok := record.Get("articleID")
			if !ok {
				continue
			}

			articleIDStr, ok := articleIDVal.(string)
			if !ok {
				continue
			}
			parsedID, err := uuid.Parse(articleIDStr)
			if err != nil {
				continue
			}
			articleIDs = append(articleIDs, parsedID)
		}
		return articleIDs, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]uuid.UUID), nil
}
