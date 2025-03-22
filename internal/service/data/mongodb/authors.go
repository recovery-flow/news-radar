package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Author struct {
	ID        uuid.UUID  `json:"id" bson:"id"`
	Name      string     `json:"name" bson:"name"`
	Desc      *string    `json:"desc" bson:"desc"`
	Avatar    *string    `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Email     *string    `json:"email,omitempty" bson:"email,omitempty"`
	Telegram  *string    `json:"telegram,omitempty" bson:"telegram,omitempty"`
	Twitter   *string    `json:"twitter,omitempty" bson:"twitter,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
}

const (
	AuthorsCollection = "authors"
)

type Authors interface {
	New() Authors

	Insert(ctx context.Context, author *Author) (*Author, error)
	Delete(ctx context.Context) error
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]*Author, error)
	Get(ctx context.Context) (*Author, error)

	FiltersID(id uuid.UUID) Authors
	FiltersName(name string) Authors

	Update(ctx context.Context, fields map[string]any) (*Author, error)

	Limit(limit int64) Authors
	Skip(skip int64) Authors
	Sort(field string, ascending bool) Authors
}

type authors struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewAuthors(uri, dbName string) (Authors, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	coll := database.Collection(AuthorsCollection)

	return &authors{
		client:     client,
		database:   database,
		collection: coll,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (a *authors) New() Authors {
	return &authors{
		client:     a.client,
		database:   a.database,
		collection: a.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (a *authors) Insert(ctx context.Context, author *Author) (*Author, error) {
	_, err := a.collection.InsertOne(ctx, author)
	if err != nil {
		return nil, fmt.Errorf("failed to insert author: %w", err)
	}
	return author, nil
}

func (a *authors) Delete(ctx context.Context) error {
	_, err := a.collection.DeleteOne(ctx, a.filters)
	if err != nil {
		return fmt.Errorf("failed to delete authors: %w", err)
	}
	return nil
}

func (a *authors) Count(ctx context.Context) (int64, error) {
	return a.collection.CountDocuments(ctx, a.filters)
}

func (a *authors) Select(ctx context.Context) ([]*Author, error) {
	cursor, err := a.collection.Find(ctx, a.filters)
	if err != nil {
		return nil, fmt.Errorf("failed to select authors: %w", err)
	}
	defer cursor.Close(ctx)

	var aths []*Author
	if err = cursor.All(ctx, &aths); err != nil {
		return nil, fmt.Errorf("failed to decode authors: %w", err)
	}
	return aths, nil
}

func (a *authors) Get(ctx context.Context) (*Author, error) {
	var ath Author
	err := a.collection.FindOne(ctx, a.filters).Decode(&ath)
	if err != nil {
		return nil, fmt.Errorf("failed to get author: %w", err)
	}
	return &ath, nil
}

func (a *authors) FiltersID(id uuid.UUID) Authors {
	a.filters["_id"] = id
	return a
}

func (a *authors) FiltersName(name string) Authors {
	a.filters["name"] = bson.M{
		"$regex":   fmt.Sprintf(".*%s.*", name),
		"$options": "i",
	}
	return a
}

func (a *authors) Update(ctx context.Context, fields map[string]any) (*Author, error) {
	validFields := map[string]bool{
		"name":       true,
		"desc":       true,
		"avatar":     true,
		"email":      true,
		"telegram":   true,
		"twitter":    true,
		"updated_at": true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields[key] = value
		}
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated Author
	err := a.collection.FindOneAndUpdate(ctx, a.filters, bson.M{"$set": updateFields}, opts).Decode(&updated)
	if err != nil {
		return nil, fmt.Errorf("failed to update document: %w", err)
	}

	for key, value := range updateFields {
		if _, exists := a.filters[key]; exists {
			a.filters[key] = value
		}
	}

	return &updated, nil
}

func (a *authors) Limit(limit int64) Authors {
	a.limit = limit
	return a
}

func (a *authors) Skip(skip int64) Authors {
	a.skip = skip
	return a
}

func (a *authors) Sort(field string, ascending bool) Authors {
	order := 1
	if !ascending {
		order = -1
	}
	a.sort = bson.D{{field, order}}
	return a
}
