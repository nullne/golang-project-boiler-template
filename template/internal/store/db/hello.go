package db

import (
	"context"

	"database/sql"

	"{{GoModule}}/internal/domain"
)

type hello struct {
	ID   int            `db:"id"`
	Who  sql.NullString `db:"who"`
	Days int            `db:"days"`
}

func (h hello) to() *domain.Hello {
	return &domain.Hello{
		Who:  h.Who.String,
		Days: h.Days,
	}
}

func (m *Manager) GetHello(ctx context.Context, who string) (*domain.Hello, error) {
	s := "SELECT * FROM hello_tab WHERE who = ?"
	var data hello
	if err := m.core.GetContext(ctx, &data, s, who); err != nil {
		return nil, err
	}
	return data.to(), nil
}
