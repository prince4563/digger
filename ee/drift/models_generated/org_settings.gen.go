// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/ee/drift/model"
)

func newOrgSetting(db *gorm.DB, opts ...gen.DOOption) orgSetting {
	_orgSetting := orgSetting{}

	_orgSetting.orgSettingDo.UseDB(db, opts...)
	_orgSetting.orgSettingDo.UseModel(&model.OrgSetting{})

	tableName := _orgSetting.orgSettingDo.TableName()
	_orgSetting.ALL = field.NewAsterisk(tableName)
	_orgSetting.CreatedAt = field.NewTime(tableName, "created_at")
	_orgSetting.ScheduleType = field.NewString(tableName, "schedule_type")
	_orgSetting.Schedule = field.NewString(tableName, "schedule")
	_orgSetting.SlackNotificationURL = field.NewString(tableName, "slack_notification_url")
	_orgSetting.OrgID = field.NewString(tableName, "org_id")
	_orgSetting.ExternalOrgID = field.NewString(tableName, "external_org_id")
	_orgSetting.ID = field.NewString(tableName, "id")

	_orgSetting.fillFieldMap()

	return _orgSetting
}

type orgSetting struct {
	orgSettingDo

	ALL                  field.Asterisk
	CreatedAt            field.Time
	ScheduleType         field.String
	Schedule             field.String
	SlackNotificationURL field.String
	OrgID                field.String
	ExternalOrgID        field.String
	ID                   field.String

	fieldMap map[string]field.Expr
}

func (o orgSetting) Table(newTableName string) *orgSetting {
	o.orgSettingDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o orgSetting) As(alias string) *orgSetting {
	o.orgSettingDo.DO = *(o.orgSettingDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *orgSetting) updateTableName(table string) *orgSetting {
	o.ALL = field.NewAsterisk(table)
	o.CreatedAt = field.NewTime(table, "created_at")
	o.ScheduleType = field.NewString(table, "schedule_type")
	o.Schedule = field.NewString(table, "schedule")
	o.SlackNotificationURL = field.NewString(table, "slack_notification_url")
	o.OrgID = field.NewString(table, "org_id")
	o.ExternalOrgID = field.NewString(table, "external_org_id")
	o.ID = field.NewString(table, "id")

	o.fillFieldMap()

	return o
}

func (o *orgSetting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *orgSetting) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 7)
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["schedule_type"] = o.ScheduleType
	o.fieldMap["schedule"] = o.Schedule
	o.fieldMap["slack_notification_url"] = o.SlackNotificationURL
	o.fieldMap["org_id"] = o.OrgID
	o.fieldMap["external_org_id"] = o.ExternalOrgID
	o.fieldMap["id"] = o.ID
}

func (o orgSetting) clone(db *gorm.DB) orgSetting {
	o.orgSettingDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o orgSetting) replaceDB(db *gorm.DB) orgSetting {
	o.orgSettingDo.ReplaceDB(db)
	return o
}

type orgSettingDo struct{ gen.DO }

type IOrgSettingDo interface {
	gen.SubQuery
	Debug() IOrgSettingDo
	WithContext(ctx context.Context) IOrgSettingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrgSettingDo
	WriteDB() IOrgSettingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrgSettingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrgSettingDo
	Not(conds ...gen.Condition) IOrgSettingDo
	Or(conds ...gen.Condition) IOrgSettingDo
	Select(conds ...field.Expr) IOrgSettingDo
	Where(conds ...gen.Condition) IOrgSettingDo
	Order(conds ...field.Expr) IOrgSettingDo
	Distinct(cols ...field.Expr) IOrgSettingDo
	Omit(cols ...field.Expr) IOrgSettingDo
	Join(table schema.Tabler, on ...field.Expr) IOrgSettingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrgSettingDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrgSettingDo
	Group(cols ...field.Expr) IOrgSettingDo
	Having(conds ...gen.Condition) IOrgSettingDo
	Limit(limit int) IOrgSettingDo
	Offset(offset int) IOrgSettingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrgSettingDo
	Unscoped() IOrgSettingDo
	Create(values ...*model.OrgSetting) error
	CreateInBatches(values []*model.OrgSetting, batchSize int) error
	Save(values ...*model.OrgSetting) error
	First() (*model.OrgSetting, error)
	Take() (*model.OrgSetting, error)
	Last() (*model.OrgSetting, error)
	Find() ([]*model.OrgSetting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrgSetting, err error)
	FindInBatches(result *[]*model.OrgSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OrgSetting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrgSettingDo
	Assign(attrs ...field.AssignExpr) IOrgSettingDo
	Joins(fields ...field.RelationField) IOrgSettingDo
	Preload(fields ...field.RelationField) IOrgSettingDo
	FirstOrInit() (*model.OrgSetting, error)
	FirstOrCreate() (*model.OrgSetting, error)
	FindByPage(offset int, limit int) (result []*model.OrgSetting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrgSettingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o orgSettingDo) Debug() IOrgSettingDo {
	return o.withDO(o.DO.Debug())
}

func (o orgSettingDo) WithContext(ctx context.Context) IOrgSettingDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orgSettingDo) ReadDB() IOrgSettingDo {
	return o.Clauses(dbresolver.Read)
}

func (o orgSettingDo) WriteDB() IOrgSettingDo {
	return o.Clauses(dbresolver.Write)
}

func (o orgSettingDo) Session(config *gorm.Session) IOrgSettingDo {
	return o.withDO(o.DO.Session(config))
}

func (o orgSettingDo) Clauses(conds ...clause.Expression) IOrgSettingDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orgSettingDo) Returning(value interface{}, columns ...string) IOrgSettingDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orgSettingDo) Not(conds ...gen.Condition) IOrgSettingDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orgSettingDo) Or(conds ...gen.Condition) IOrgSettingDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orgSettingDo) Select(conds ...field.Expr) IOrgSettingDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orgSettingDo) Where(conds ...gen.Condition) IOrgSettingDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orgSettingDo) Order(conds ...field.Expr) IOrgSettingDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orgSettingDo) Distinct(cols ...field.Expr) IOrgSettingDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orgSettingDo) Omit(cols ...field.Expr) IOrgSettingDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orgSettingDo) Join(table schema.Tabler, on ...field.Expr) IOrgSettingDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orgSettingDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrgSettingDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orgSettingDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrgSettingDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orgSettingDo) Group(cols ...field.Expr) IOrgSettingDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orgSettingDo) Having(conds ...gen.Condition) IOrgSettingDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orgSettingDo) Limit(limit int) IOrgSettingDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orgSettingDo) Offset(offset int) IOrgSettingDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orgSettingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrgSettingDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orgSettingDo) Unscoped() IOrgSettingDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orgSettingDo) Create(values ...*model.OrgSetting) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orgSettingDo) CreateInBatches(values []*model.OrgSetting, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orgSettingDo) Save(values ...*model.OrgSetting) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orgSettingDo) First() (*model.OrgSetting, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrgSetting), nil
	}
}

func (o orgSettingDo) Take() (*model.OrgSetting, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrgSetting), nil
	}
}

func (o orgSettingDo) Last() (*model.OrgSetting, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrgSetting), nil
	}
}

func (o orgSettingDo) Find() ([]*model.OrgSetting, error) {
	result, err := o.DO.Find()
	return result.([]*model.OrgSetting), err
}

func (o orgSettingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrgSetting, err error) {
	buf := make([]*model.OrgSetting, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orgSettingDo) FindInBatches(result *[]*model.OrgSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orgSettingDo) Attrs(attrs ...field.AssignExpr) IOrgSettingDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orgSettingDo) Assign(attrs ...field.AssignExpr) IOrgSettingDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orgSettingDo) Joins(fields ...field.RelationField) IOrgSettingDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orgSettingDo) Preload(fields ...field.RelationField) IOrgSettingDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orgSettingDo) FirstOrInit() (*model.OrgSetting, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrgSetting), nil
	}
}

func (o orgSettingDo) FirstOrCreate() (*model.OrgSetting, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrgSetting), nil
	}
}

func (o orgSettingDo) FindByPage(offset int, limit int) (result []*model.OrgSetting, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o orgSettingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orgSettingDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orgSettingDo) Delete(models ...*model.OrgSetting) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orgSettingDo) withDO(do gen.Dao) *orgSettingDo {
	o.DO = *do.(*gen.DO)
	return o
}
