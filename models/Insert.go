package models

func insert(todo Todo) (id int64, err error) {

	conn, err := banco.ConnectToDB
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `insert into todos (title, description, done ) VALUES ($1,$2,$3) RETURNING id`
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.done).Scan(&id)
	return
}
