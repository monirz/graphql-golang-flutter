package dataloaders

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/monirz/gql/api"
	"github.com/monirz/gql/api/dbl"
)

type ctxKeyType struct{ name string }

var CtxKey = ctxKeyType{"dataloaderctx"}

type Loaders struct {
	UserByID *UserLoader
}

func DataloaderMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			wait:     1 * time.Millisecond,
			maxBatch: 100,
			fetch: func(ids []int) ([]*api.User, []error) {
				var sqlQuery string
				if len(ids) == 1 {
					sqlQuery = "SELECT id, name, email from users WHERE id = ?"
				} else {
					sqlQuery = "SELECT id, name, email from users WHERE id IN (?)"
				}
				sqlQuery, arguments, err := sqlx.In(sqlQuery, ids)
				if err != nil {
					log.Println(err)
				}
				sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
				rows, err := dbl.LogAndQuery(db, sqlQuery, arguments...)
				defer rows.Close()
				if err != nil {
					log.Println(err)
				}
				userById := map[string]*api.User{}

				for rows.Next() {
					user := api.User{}
					if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
						fmt.Println(err)
						return nil, []error{errors.New("Internal error")}
					}
					userById[user.ID] = &user
				}

				users := make([]*api.User, len(ids))
				for i, id := range ids {
					users[i] = userById[string(id)]
					i++
				}

				return users, nil
			},
		}
		ctx := context.WithValue(r.Context(), CtxKey, &userloader)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
