package util

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/orpheeh/jalbv-backend/config/database"
)

func Create(tableName string, datas map[string]string) (int64, error) {
	var keys, values string = "", ""
	for k, v := range datas {
		if keys == "" {
			keys = fmt.Sprintf(`"%v"`, k)
		} else {
			keys = fmt.Sprintf(`%v,"%v"`, keys, k)
		}
		if values == "" {
			values = fmt.Sprintf(`'%v'`, v)
		} else {
			values = fmt.Sprintf(`%v,'%v'`, values, v)
		}
	}
	str := fmt.Sprintf(`INSERT INTO "%v" (%v) VALUES (%v) RETURNING id`, tableName, keys, values)
	fmt.Println(str)
	var id int64
	err := database.Postgres.QueryRow(str).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("add%v: %v", tableName, err)
	}
	return id, nil
}

func ReadAll(tableName string, variables []interface{}, keys []string, condition string) ([]map[string]interface{}, error) {
	var datas []map[string]interface{}
	params := ""
	for i, k := range keys {
		if i == 0 {
			params = fmt.Sprintf(`"%v"`, k)
		} else {
			params = fmt.Sprintf(`%v,"%v"`, params, k)
		}
	}
	rows, err := database.Postgres.Query(fmt.Sprintf(`SELECT %v FROM "%v" %v`, params, tableName, condition))
	if err != nil {
		return nil, fmt.Errorf("%v : %v", tableName, err)
	}
	defer rows.Close()
	for rows.Next() {
		var data map[string]interface{} = make(map[string]interface{})
		if err := rows.Scan(variables...); err != nil {
			return nil, fmt.Errorf("%v: %v", tableName, err)
		}
		for i, value := range variables {
			rv := reflect.ValueOf(value)
			if rv.Elem().Kind() == reflect.Bool {
				data[keys[i]] = rv.Elem().Bool()
			} else if rv.Elem().Kind() == reflect.String {
				data[keys[i]] = rv.Elem().String()
			} else if rv.Elem().Kind() == reflect.Int64 {
				data[keys[i]] = rv.Elem().Int()
			} else if rv.Elem().Kind() == reflect.Int {
				data[keys[i]] = rv.Elem().Int()
			}
		}
		datas = append(datas, data)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%v: %v", tableName, err)
	}
	return datas, nil
}

func ReadOne(tableName string, variables []interface{}, keys []string, condition string) (map[string]interface{}, error) {
	var data map[string]interface{} = make(map[string]interface{})
	params := ""
	for i, k := range keys {
		if i == 0 {
			params = fmt.Sprintf(`"%v"`, k)
		} else {
			params = fmt.Sprintf(`%v,"%v"`, params, k)
		}
	}
	str := fmt.Sprintf(`SELECT %v FROM "%v" %v`, params, tableName, condition)
	row := database.Postgres.QueryRow(str)
	if err := row.Scan(variables...); err != nil {
		if err == sql.ErrNoRows {
			return data, fmt.Errorf("No such found")
		}
		return data, fmt.Errorf("%vById: %v", tableName, err)
	}
	for i, value := range variables {
		rv := reflect.ValueOf(value)
		if rv.Elem().Kind() == reflect.Bool {
			data[keys[i]] = rv.Elem().Bool()
		} else if rv.Elem().Kind() == reflect.String {
			data[keys[i]] = rv.Elem().String()
		} else if rv.Elem().Kind() == reflect.Int64 {
			data[keys[i]] = rv.Elem().Int()
		} else if rv.Elem().Kind() == reflect.Int {
			data[keys[i]] = rv.Elem().Int()
		}
	}
	return data, nil
}

func Update(tableName string, datas map[string]string, condition string) (int64, error) {
	var updated string = ""
	for k, v := range datas {
		if updated == "" {
			updated = fmt.Sprintf(`"%v" = '%v'`, k, v)
		} else {
			updated = fmt.Sprintf(`%v,"%v" = '%v'`, updated, k, v)
		}
	}
	str := fmt.Sprintf(`UPDATE "%v" SET %v %v`, tableName, updated, condition)
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

func Delete(tableName string, condition string) (int64, error) {
	str := fmt.Sprintf(`DELETE FROM "%v" %v`, tableName, condition)
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
