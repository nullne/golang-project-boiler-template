package {{ServiceName}}

import (
	"context"

	"{{GoModule}}/internal/domain"
	"{{GoModule}}/internal/store"
)

type {{title ServiceName}} struct{
	store store.Interface
}

func New(s store.Interface) {{title ServiceName}} {
	return {{title ServiceName}}{
		store: s,
	}
}


func ({{slice ServiceName 0 1}} *{{title ServiceName}}) GetHello(ctx context.Context, who string) (*domain.Hello, error){
	return nil, nil
}