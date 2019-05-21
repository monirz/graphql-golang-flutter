package api

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Review struct {
	ID          int       `json:"id"`
	VideoID     int       `json:"videoId"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Screenshot struct {
	ID      int    `json:"id"`
	VideoID int    `json:"videoId"`
	URL     string `json:"url"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Video struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      string    `json:"-"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Lets redefine the base ID type to use an id from an external library
func MarshalID(id int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

// And the same for the unmarshaler
func UnmarshalID(v interface{}) (int64, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("ids must be strings")
	}
	i, e := strconv.Atoi(id)
	return int64(i), e
}

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, errors.New("Timestamp parsing error")
}
