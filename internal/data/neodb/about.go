package neodb

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/recovery-flow/news-radar/internal/app/models"
	"github.com/recovery-flow/news-radar/internal/config"
)

type AboutImpl struct {
	driver neo4j.Driver
}

func NewAbout(cfg config.Config) (*AboutImpl, error) {
	driver, err := neo4j.NewDriver(cfg.Database.Neo4j.URI, neo4j.BasicAuth(cfg.Database.Neo4j.Username, cfg.Database.Neo4j.Password, ""))
	if err != nil {
		return nil, fmt.Errorf("failed to create neo4j driver: %w", err)
	}

	if err = driver.VerifyConnectivity(); err != nil {
		return nil, fmt.Errorf("failed to verify connectivity: %w", err)
	}

	return &AboutImpl{
		driver: driver,
	}, nil
}

func (a *AboutImpl) Create(ctx context.Context, articleID uuid.UUID, theme string) error {
	session, err := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		cypher := `
			MATCH (art:Article { id: $articleID })
			MATCH (th:Theme { name: $theme })
			MERGE (art)-[r:ABOUT]->(th)
		`
		params := map[string]any{
			"articleID": articleID.String(),
			"theme":     theme,
		}

		_, err := tx.Run(cypher, params)
		if err != nil {
			return nil, fmt.Errorf("failed to create ABOUT relationship: %w", err)
		}
		return nil, nil
	})

	return err
}

func (a *AboutImpl) Delete(ctx context.Context, articleID uuid.UUID, theme string) error {
	session, err := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		cypher := `
			MATCH (art:Article { id: $articleID })-[r:ABOUT]->(th:Theme { name: $themeName })
			DELETE r
		`
		params := map[string]any{
			"articleID": articleID.String(),
			"themeName": theme,
		}

		_, err := tx.Run(cypher, params)
		if err != nil {
			return nil, fmt.Errorf("failed to delete ABOUT relationship: %w", err)
		}
		return nil, nil
	})

	return err
}

func (a *AboutImpl) SetForArticle(ctx context.Context, articleID uuid.UUID, themes []string) error {
	session, err := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		deleteCypher := `
			MATCH (a:Article { id: $articleID })-[r:ABOUT]->(:Theme)
			DELETE r
		`
		params := map[string]any{"articleID": articleID.String()}
		_, err := tx.Run(deleteCypher, params)
		if err != nil {
			return nil, fmt.Errorf("failed to delete existing ABOUT relationships: %w", err)
		}

		createCypher := `
			MATCH (a:Article { id: $articleID })
			FOREACH (themeName IN $themes |
				MATCH (th:Theme { name: themeName })
				MERGE (a)-[:ABOUT]->(th)
			)
		`
		params["themes"] = themes
		_, err = tx.Run(createCypher, params)
		if err != nil {
			return nil, fmt.Errorf("failed to create new ABOUT relationships: %w", err)
		}
		return nil, nil
	})
	return err
}

func (a *AboutImpl) GetForArticle(ctx context.Context, articleID uuid.UUID) ([]ThemeModels, error) {
	session, err := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	if err != nil {
		return nil, err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (any, error) {
		cypher := `
			MATCH (a:Article { id: $articleID })-[:ABOUT]->(th:Theme)
			RETURN th
		`
		params := map[string]any{
			"articleID": articleID.String(),
		}

		records, err := tx.Run(cypher, params)
		if err != nil {
			return nil, err
		}

		var themesList []ThemeModels
		for records.Next() {
			record := records.Record()
			node, ok := record.Get("th")
			if !ok {
				continue
			}
			props := node.(neo4j.Node).Props()
			status, err := models.ParseThemeStatus(props["status"].(string))
			if err != nil {
				return nil, err
			}
			theme := ThemeModels{
				Name:   props["name"].(string),
				Status: status,
			}
			if createdAtStr, ok := props["created_at"].(string); ok {
				if parsedTime, err := time.Parse(time.RFC3339, createdAtStr); err == nil {
					theme.CreatedAt = parsedTime
				}
			}
			themesList = append(themesList, theme)
		}
		return themesList, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]ThemeModels), nil
}
