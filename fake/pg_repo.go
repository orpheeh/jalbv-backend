package fake

import (
	"fmt"

	"github.com/orpheeh/jalbv-backend/config/database"
)

func getAllFakes() ([]Fake, error) {
	var fakes []Fake

	rows, err := database.Postgres.Query(`SELECT * FROM "Fake"`)
	if err != nil {
		return nil, fmt.Errorf("fakes : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var fake Fake
		if err := rows.Scan(&fake.ID, &fake.Name); err != nil {
			return nil, fmt.Errorf("fakes: %v", err)
		}
		fakes = append(fakes, fake)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("fakes: %v", err)
	}
	return fakes, nil
}

func addFake(fake Fake) (int64, error) {
	result, err := database.Postgres.Exec(fmt.Sprintf(`INSERT INTO "Fake" (name) VALUES ('%v')`, fake.Name))
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("addFake: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("addFake: %v", err)
	}
	return id, nil
}
