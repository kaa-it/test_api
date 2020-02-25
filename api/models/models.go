package models

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

//type Timestamp int64

func MarshalTimestamp(t int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatInt(t, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (int64, error) {
	if res, ok := v.(json.Number); ok {
		return res.Int64()
	}
	if res, ok := v.(string); ok {
		return json.Number(res).Int64()
	}
	if res, ok := v.(int64); ok {
		return res, nil
	}
	if res, ok := v.(*int64); ok {
		return *res, nil
	}
	return 0, fmt.Errorf("could not convert %v of type %T to Int64", v, v)
}

func MarshalID(mac string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, mac)
	})
}

func UnmarshalID(v interface{}) (string, error) {
	mac, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("macs must be strings")
	}

	return mac, nil
}
