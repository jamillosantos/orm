package orm

import "context"

type HookBeforeInsert interface {
	BeforeInsert(ctx context.Context, fields ...SchemaField) error
}

type HookAfterInsert interface {
	AfterInsert(ctx context.Context, err error, fields ...SchemaField)
}

type HookBeforeUpdate interface {
	BeforeUpdate(ctx context.Context, fields ...SchemaField) error
}

type HookAfterUpdate interface {
	AfterUpdate(ctx context.Context, err error, fields ...SchemaField)
}

type HookBeforeSave interface {
	BeforeSave(ctx context.Context, fields ...SchemaField) error
}

type HookAfterSave interface {
	AfterSave(ctx context.Context, err error, fields ...SchemaField)
}

type HookBeforeDelete interface {
	BeforeDelete(ctx context.Context) error
}

type HookAfterDelete interface {
	AfterDelete(ctx context.Context, err error)
}
