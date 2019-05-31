package gql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/markbates/going/randx"
	"github.com/monirz/gql/api"
	"github.com/monirz/gql/api/dbl"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

const (
	DB_USER     = "root"
	DB_PASSWORD = "password"
	DB_NAME     = "gqldemo"
)

var DB *sql.DB

var videoPublishedChannel map[string]chan *api.Video
var userCreatedChannel map[string]chan *api.User

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:password@/gqldemo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	videoPublishedChannel = map[string]chan *api.Video{}
}

func NewRootResolvers(db *sql.DB) Config {

	c := Config{
		Resolvers: &Resolver{
			db: db,
		},
	}

	// Complexity
	countComplexity := func(childComplexity int, limit *int, offset *int) int {
		return *limit * childComplexity
	}
	c.Complexity.Query.Videos = countComplexity
	c.Complexity.Video.Related = countComplexity

	// // Schema Directive
	// c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	// 	ctxUserID := ctx.Value(UserIDCtxKey)
	// 	if ctxUserID != nil {
	// 		return next(ctx)
	// 	} else {
	// 		return nil, errors.UnauthorisedError
	// 	}
	// }
	return c
}

type Resolver struct {
	db *sql.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Review() ReviewResolver {
	return &reviewResolver{r}
}
func (r *Resolver) Screenshot() ScreenshotResolver {
	return &screenshotResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}
func (r *Resolver) Video() VideoResolver {
	return &videoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*api.User, error) {
	NewUser := &api.User{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}

	r.db = DB
	rows, err := dbl.LogAndQuery(r.db, "INSERT INTO users (id ,name, email) VALUES(?, ?, ?)",
		input.ID, input.Name, input.Email)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		if err := rows.Scan(&NewUser.ID); err != nil {
			log.Println(err)

			return nil, err
		}
	}

	defer rows.Close()

	for _, observer := range userCreatedChannel {
		observer <- NewUser
	}
	log.Println("new user > ", NewUser)

	return NewUser, err
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input NewVideo) (*api.Video, error) {
	newVideo := &api.Video{
		URL:         input.URL,
		Description: input.Description,
		Name:        input.Name,
		CreatedAt:   time.Now().UTC(),
		UserID:      "1",
	}

	r.db = DB
	rows, err := dbl.LogAndQuery(r.db, "INSERT INTO videos (name, description, url, user_id, created_at) VALUES(?, ?, ?, ?, ?)",
		input.Name, input.Description, input.URL, "1", newVideo.CreatedAt)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		if err := rows.Scan(&newVideo.ID); err != nil {
			log.Println(err)

			return nil, err
		}
	}

	defer rows.Close()

	for _, observer := range videoPublishedChannel {
		observer <- newVideo
	}
	log.Println("new video > ", newVideo)

	return newVideo, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Videos(ctx context.Context, limit *int, offset *int) ([]*api.Video, error) {
	var video api.Video
	var videos []*api.Video

	// r.db = DB
	// fmt.Println(r.db)
	rows, err := dbl.LogAndQuery(r.db, "SELECT id, name, description, url, created_at, user_id FROM videos ORDER BY created_at desc limit ? offset ?", 10, 0)
	defer rows.Close()
	fmt.Println("debug ------------------ ", rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&video.ID, &video.Name, &video.Description, &video.URL, &video.CreatedAt, &video.UserID)

		fmt.Println("vidoe name ---------------------------------------", video.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		videos = append(videos, &video)
	}

	return videos, nil
}

type reviewResolver struct{ *Resolver }

func (r *reviewResolver) ID(ctx context.Context, obj *api.Review) (int64, error) {
	panic("not implemented id")
}
func (r *reviewResolver) VideoID(ctx context.Context, obj *api.Review) (int64, error) {
	panic("not implemented video id")
}
func (r *reviewResolver) User(ctx context.Context, obj *api.Review) (*api.User, error) {
	panic("not implemented user")
}

type screenshotResolver struct{ *Resolver }

func (r *screenshotResolver) ID(ctx context.Context, obj *api.Screenshot) (int64, error) {
	panic("not implemented id screen")
}
func (r *screenshotResolver) VideoID(ctx context.Context, obj *api.Screenshot) (int64, error) {
	panic("not implemented video screen id ")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) VideoPublished(ctx context.Context) (<-chan *api.Video, error) {
	id := randx.String(8)

	videoEvent := make(chan *api.Video, 1)
	go func() {
		<-ctx.Done()
	}()
	videoPublishedChannel[id] = videoEvent
	return videoEvent, nil
}

func (r *subscriptionResolver) UserCreated(ctx context.Context) (<-chan *api.User, error) {
	id := randx.String(8)

	userEvent := make(chan *api.User, 1)
	go func() {
		<-ctx.Done()
	}()
	userCreatedChannel[id] = userEvent
	return userEvent, nil

}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *api.User) (int64, error) {
	panic("not implemented user resolver id")
}

type videoResolver struct{ *Resolver }

func (r *videoResolver) User(ctx context.Context, obj *api.Video) (*api.User, error) {

	user := &api.User{ID: "1", Name: "john doe", Email: "john@mail.com"}

	return user, nil
}
func (r *videoResolver) Screenshots(ctx context.Context, obj *api.Video) ([]*api.Screenshot, error) {
	panic("not implemented video reolver screenshots")
}
func (r *videoResolver) Related(ctx context.Context, obj *api.Video, limit *int, offset *int) ([]*api.Video, error) {
	panic("not implemented related ")
}
