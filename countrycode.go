package gql_test

import (
	"errors"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"golang.org/x/text/language"
)

func MarshalCountryCode(r language.Region) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(r.String()))
	})
}

func UnmarshalCountryCode(v interface{}) (language.Region, error) {
	if tmpStr, ok := v.(string); ok {
		return language.ParseRegion(tmpStr)
	}
	return language.Region{}, errors.New("duration should be a ISO8601 formatted string")
}
