package util

import (
	"fmt"
	"reflect"

	"github.com/orpheeh/jalbv-backend/config/database"
)

func Create(tableName string, datas map[string]string) (int64, error) {
	var keys, values string = "", ""
	for k, v := range datas {
		if keys == "" {
			keys = k
		} else {
			keys = fmt.Sprintf("%v,%v", keys, k)
		}
		if values == "" {
			values = fmt.Sprintf(`'%v'`, v)
		} else {
			values = fmt.Sprintf(`%v,'%v'`, values, v)
		}
	}
	str := fmt.Sprintf(`INSERT INTO "%v" (%v) VALUES (%v)`, tableName, keys, values)
	result, err := database.Postgres.Exec(str)
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("add%v: %v", tableName, err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("add%v: %v", tableName, err)
	}
	return id, nil
}

func ReadAll(tableName string, variables []interface{}, keys []string) ([]map[string]string, error) {
	var datas []map[string]string
	rows, err := database.Postgres.Query(fmt.Sprintf(`SELECT * FROM "%v"`, tableName))
	if err != nil {
		return nil, fmt.Errorf("%v : %v", tableName, err)
	}
	defer rows.Close()
	for rows.Next() {
		var data map[string]string = make(map[string]string)
		if err := rows.Scan(variables...); err != nil {
			return nil, fmt.Errorf("%v: %v", tableName, err)
		}
		for i, value := range variables {
			rv := reflect.ValueOf(value)
			if rv.Kind() == reflect.Ptr {
				data[keys[i]] = rv.Elem().String()
			}
		}
		datas = append(datas, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%v: %v", tableName, err)
	}
	return datas, nil
}

func Update() {

}

func Delete() {

}
