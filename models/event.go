// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/tikasan/eventory/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": Models
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"golang.org/x/net/context"
)

// イベント
type Event struct {
	ID               int `gorm:"primary_key"` // primary key
	APIType          string
	Title            string
	Accept           int
	Address          string
	DataHash         string
	Description      string
	EventGenres      []EventGenre // has many EventGenres
	Identifier       string
	Limits           int
	PrefID           int // has many Event
	URL              string
	UserFollowEvents []UserFollowEvent // has many UserFollowEvents
	Wait             int
	CreatedAt        time.Time  // timestamp
	DeletedAt        *time.Time // nullable timestamp (soft delete)
	EndAt            time.Time  // timestamp
	StartAt          time.Time  // timestamp
	UpdatedAt        time.Time  // timestamp
	Pref             Pref
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Event) TableName() string {
	return "events"

}

// EventDB is the implementation of the storage interface for
// Event.
type EventDB struct {
	Db *gorm.DB
}

// NewEventDB creates a new storage type.
func NewEventDB(db *gorm.DB) *EventDB {
	return &EventDB{Db: db}
}

// DB returns the underlying database.
func (m *EventDB) DB() interface{} {
	return m.Db
}

// EventStorage represents the storage interface.
type EventStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Event, error)
	Get(ctx context.Context, id int) (*Event, error)
	Add(ctx context.Context, event *Event) error
	Update(ctx context.Context, event *Event) error
	Delete(ctx context.Context, id int) error

	ListEvent(ctx context.Context, prefID int) []*app.Event
	OneEvent(ctx context.Context, id int, prefID int) (*app.Event, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *EventDB) TableName() string {
	return "events"

}

// Belongs To Relationships

// EventFilterByPref is a gorm filter for a Belongs To relationship.
func EventFilterByPref(prefID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if prefID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("pref_id = ?", prefID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Event as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *EventDB) Get(ctx context.Context, id int) (*Event, error) {
	defer goa.MeasureSince([]string{"goa", "db", "event", "get"}, time.Now())

	var native Event
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Event
func (m *EventDB) List(ctx context.Context) ([]*app.Event, error) {
	defer goa.MeasureSince([]string{"goa", "db", "event", "list"}, time.Now())

	var objs []*app.Event
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// List returns an array of Event
func (m *EventDB) ListByQ(ctx context.Context, q string, sort string, page int) ([]*app.Event, error) {
	defer goa.MeasureSince([]string{"goa", "db", "event", "list"}, time.Now())

	var objs []*app.Event
	err := m.Db.Table(m.TableName()).
		Find(&objs).
		Scopes(
			CreatePagingQuery(page),
			CreateSortQuery(sort),
			CreateLikeQuery(q, "description")).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *EventDB) Add(ctx context.Context, model *Event) error {
	defer goa.MeasureSince([]string{"goa", "db", "event", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Event", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *EventDB) Update(ctx context.Context, model *Event) error {
	defer goa.MeasureSince([]string{"goa", "db", "event", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Event", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *EventDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "event", "delete"}, time.Now())

	var obj Event

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Event", "error", err.Error())
		return err
	}

	return nil
}
