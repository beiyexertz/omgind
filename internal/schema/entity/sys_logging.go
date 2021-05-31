package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/wanhello/omgind/internal/schema/mixin"
)

/*

// LogItem 存储日志项
type LogItem struct {
	ID         uint      `gorm:"column:id;primary_key;auto_increment;"` // id
	Level      string    `gorm:"column:level;size:20;index;"`           // 日志级别
	TraceID    string    `gorm:"column:trace_id;size:128;index;"`       // 跟踪ID
	UserID     string    `gorm:"column:user_id;size:36;index;"`         // 用户ID
	Tag        string    `gorm:"column:tag;size:128;index;"`            // Tag
	Version    string    `gorm:"column:version;index;size:64;"`         // 版本号
	Message    string    `gorm:"column:message;size:1024;"`             // 消息
	Data       string    `gorm:"column:data;type:text;"`                // 日志数据(json)
	ErrorStack string    `gorm:"column:error_stack;type:text;"`         // Error Stack
	CreatedAt  time.Time `gorm:"column:created_at;index"`               // 创建时间
}

*/

type SysLogging struct {
	ent.Schema
}

func (SysLogging) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.MemoMixin{},
	}
}

// Fields of the SysLogging.
func (SysLogging) Fields() []ent.Field {
	return []ent.Field{

		field.String("level").MaxLen(32).Comment("日志级别"),
		field.String("trace_id").MaxLen(128).Comment("跟踪ID"),
		field.String("user_id").MaxLen(128).Comment("用户ID"),
		field.String("tag").MaxLen(128).Comment("Tag"),

		field.String("version").MaxLen(64).Comment("版本号"),
		field.String("message").Comment("消息"),
		field.String("data").Comment("日志数据(json string)"),
		field.String("error_stack").Comment("日志数据(json string)"),

		mixin.FieldCreateAt(),
	}
}

func (SysLogging) Indexes() []ent.Index {
	
	return []ent.Index{
		index.Fields("level"),
		index.Fields("trace_id"),
		index.Fields("user_id"),
		index.Fields("tag"),
		mixin.IndexCreateAt(),
	}
}


