package domain

import (
	"context"
	"database/sql"
	"fmt"
)

// cordinate with database
type EmpolyeeRepoDB struct {
	db *sql.DB
}

func NewEmpolyeeDB(db *sql.DB) *EmpolyeeRepoDB {
	return &EmpolyeeRepoDB{db: db}
}

func (d *EmpolyeeRepoDB) CreateEmpolyee(ctx context.Context, empolyee Empolyee) (*Empolyee, error) {

	result, err := d.db.Exec("INSERT INTO employees (name, position, salary, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		empolyee.Name, empolyee.Position, empolyee.Salary, empolyee.CreatedAt, empolyee.UpdateAt)

	if err != nil {

		return nil, err
	}
	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	emp, _ := d.FindByEmpolyeeId(ctx, int(id))

	return emp, nil

}

func (d *EmpolyeeRepoDB) FindByEmpolyeeId(ctx context.Context, employeeID int) (*Empolyee, error) {

	query := `SELECT * FROM employees WHERE id = $1`
	row := d.db.QueryRowContext(ctx, query, employeeID) // Assuming employeeID is the ID of the employee you want to fetch

	var employee Empolyee

	err := row.Scan(&employee.ID, &employee.Name, &employee.Position, &employee.Salary, &employee.CreatedAt, &employee.UpdateAt)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (d *EmpolyeeRepoDB) AllEmpolyee(ctx context.Context, paination Pagination) (*EmpolyeeResponseDto, error) {

	var empolyee []*Empolyee
	baseQuery := `SELECT * FROM employees ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := d.db.QueryContext(ctx, baseQuery, paination.RecordPerPage, paination.CurrentPage)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var e Empolyee

		err := rows.Scan(&e.ID, &e.Name, &e.Position, &e.Salary, &e.CreatedAt, &e.UpdateAt)

		if err != nil {
			return nil, err
		}
		empolyee = append(empolyee, &e)

		if err := rows.Err(); err != nil {
			return nil, err
		}

	}

	// total count
	countQuery := `SELECT COUNT(*) FROM employees`
	var totalCount int

	if err = d.db.QueryRowContext(ctx, countQuery).Scan(&totalCount); err != nil {
		return nil, err
	}
	paination.TotalCount = totalCount
	paination.NextPage = totalCount > paination.RecordPerPage

	return &EmpolyeeResponseDto{Empolyee: empolyee, MetaData: paination}, nil
}

func (d *EmpolyeeRepoDB) DeleteEmpolyeeRecord(ctx context.Context, id int) error {

	query := `Delete from employees where id=$1`

	result, err := d.db.Exec(query, id)

	fmt.Println(result)

	if err != nil {
		return err
	}

	return nil
}

func (d *EmpolyeeRepoDB) UpdateEmpolyee(ctx context.Context, emp *EmpolyeeUpdateRequestDto) (*Empolyee, error) {

	query := "UPDATE employees SET"
	var args []interface{}

	if emp.Name != "" {
		query += " name = ?,"
		args = append(args, emp.Name)
	}
	if emp.Position != "" {
		query += " position = ?,"
		args = append(args, emp.Position)
	}
	if emp.Salary != 0 {
		query += " salary = ?,"
		args = append(args, emp.Salary)
	}

	query += " updated_at = ?,"
	args = append(args, emp.UpdateAt)

	// Remove the trailing comma
	query = query[:len(query)-1]
	query += " WHERE id = ?"
	args = append(args, emp.ID)

	_, err := d.db.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	// Fetch the updated record

	empolyee, _ := d.FindByEmpolyeeId(ctx, emp.ID)

	return empolyee, nil

}
