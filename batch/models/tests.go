package models

func UpdateTest() error {
	sql := `update test set created_at = now()`
	_, err := Db.Exec(sql)
	return err
}
