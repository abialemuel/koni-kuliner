package utility

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func Btoi(boolean bool) int {
	if boolean {
		return 1
	}
	return 0
}

func Filter(r *http.Request, searchList []string) map[string][]string {
	f := map[string][]string{}
	for _, params := range searchList {
		if len(r.URL.Query()[params]) > 0 {
			f[params] = r.URL.Query()[params]
		}
	}
	return f
}

func AppendQuery(query string, f map[string][]string) (string, []interface{}) {
	args := []interface{}{}
	if f["address"] != nil {
		address := f["address"]
		addressLike := "%" + address[0] + "%"
		newQuery, newArgs, _ := sqlx.In(" AND address LIKE ? ", addressLike)
		args = append(args, newArgs...)
		query += newQuery
	}

	if f["id"] != nil {
		ids := f["id"]
		newQuery, newArgs, _ := sqlx.In(" AND ID IN (?) ", ids)
		args = append(args, newArgs...)
		query += newQuery
	}

	if f["brand_id"] != nil {
		ids := f["brand_id"]
		newQuery, newArgs, _ := sqlx.In(" AND brand_id IN (?) ", ids)
		args = append(args, newArgs...)
		query += newQuery
	}

	if f["outlet_id"] != nil {
		ids := f["outlet_id"]
		newQuery, newArgs, _ := sqlx.In(" AND outlet_id IN (?) ", ids)
		args = append(args, newArgs...)
		query += newQuery
	}

	if f["customer_id"] != nil {
		ids := f["customer_id"]
		newQuery, newArgs, _ := sqlx.In(" AND customer_id IN (?) ", ids)
		args = append(args, newArgs...)
		query += newQuery
	}

	if f["name"] != nil {
		name := f["name"]
		nameLike := "%" + name[0] + "%"
		newQuery, newArgs, _ := sqlx.In(" AND name LIKE ? ", nameLike)
		args = append(args, newArgs...)
		query += newQuery
	}

	if f["limit"] == nil {
		args = append(args, fmt.Sprint(20))
		f["limit"] = []string{"20"}
		query += " LIMIT ?"
	} else {
		args = append(args, f["limit"][0])
		query += " LIMIT ?"
	}

	if f["offset"] == nil {
		args = append(args, fmt.Sprint(0))
		f["offset"] = []string{"0"}
		query += " OFFSET ?"
	} else {
		args = append(args, f["offset"][0])
		query += " OFFSET ?"
	}

	return query, args

}
