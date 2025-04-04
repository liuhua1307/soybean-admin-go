// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"soybean-admin-go/db/model"
)

func newGood(db *gorm.DB, opts ...gen.DOOption) good {
	_good := good{}

	_good.goodDo.UseDB(db, opts...)
	_good.goodDo.UseModel(&model.Good{})

	tableName := _good.goodDo.TableName()
	_good.ALL = field.NewAsterisk(tableName)
	_good.ID = field.NewInt64(tableName, "id")
	_good.Name = field.NewString(tableName, "name")
	_good.Repo = field.NewString(tableName, "repo")
	_good.Class = field.NewString(tableName, "class")
	_good.Inventory = field.NewString(tableName, "inventory")
	_good.Weight = field.NewString(tableName, "weight")
	_good.Desc = field.NewString(tableName, "desc")
	_good.Orders = goodManyToManyOrders{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Orders", "model.Order"),
		Goods: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Orders.Goods", "model.Good"),
		},
		CustomerInfo: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Orders.CustomerInfo", "model.CustomerInfo"),
		},
	}

	_good.GoodOrders = goodHasOneGoodOrders{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("GoodOrders", "model.GoodOrder"),
	}

	_good.fillFieldMap()

	return _good
}

type good struct {
	goodDo goodDo

	ALL       field.Asterisk
	ID        field.Int64
	Name      field.String
	Repo      field.String
	Class     field.String
	Inventory field.String
	Weight    field.String
	Desc      field.String
	Orders    goodManyToManyOrders

	GoodOrders goodHasOneGoodOrders

	fieldMap map[string]field.Expr
}

func (g good) Table(newTableName string) *good {
	g.goodDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g good) As(alias string) *good {
	g.goodDo.DO = *(g.goodDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *good) updateTableName(table string) *good {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.Name = field.NewString(table, "name")
	g.Repo = field.NewString(table, "repo")
	g.Class = field.NewString(table, "class")
	g.Inventory = field.NewString(table, "inventory")
	g.Weight = field.NewString(table, "weight")
	g.Desc = field.NewString(table, "desc")

	g.fillFieldMap()

	return g
}

func (g *good) WithContext(ctx context.Context) IGoodDo { return g.goodDo.WithContext(ctx) }

func (g good) TableName() string { return g.goodDo.TableName() }

func (g good) Alias() string { return g.goodDo.Alias() }

func (g good) Columns(cols ...field.Expr) gen.Columns { return g.goodDo.Columns(cols...) }

func (g *good) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *good) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 9)
	g.fieldMap["id"] = g.ID
	g.fieldMap["name"] = g.Name
	g.fieldMap["repo"] = g.Repo
	g.fieldMap["class"] = g.Class
	g.fieldMap["inventory"] = g.Inventory
	g.fieldMap["weight"] = g.Weight
	g.fieldMap["desc"] = g.Desc

}

func (g good) clone(db *gorm.DB) good {
	g.goodDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g good) replaceDB(db *gorm.DB) good {
	g.goodDo.ReplaceDB(db)
	return g
}

type goodManyToManyOrders struct {
	db *gorm.DB

	field.RelationField

	Goods struct {
		field.RelationField
	}
	CustomerInfo struct {
		field.RelationField
	}
}

func (a goodManyToManyOrders) Where(conds ...field.Expr) *goodManyToManyOrders {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a goodManyToManyOrders) WithContext(ctx context.Context) *goodManyToManyOrders {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a goodManyToManyOrders) Session(session *gorm.Session) *goodManyToManyOrders {
	a.db = a.db.Session(session)
	return &a
}

func (a goodManyToManyOrders) Model(m *model.Good) *goodManyToManyOrdersTx {
	return &goodManyToManyOrdersTx{a.db.Model(m).Association(a.Name())}
}

type goodManyToManyOrdersTx struct{ tx *gorm.Association }

func (a goodManyToManyOrdersTx) Find() (result []*model.Order, err error) {
	return result, a.tx.Find(&result)
}

func (a goodManyToManyOrdersTx) Append(values ...*model.Order) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a goodManyToManyOrdersTx) Replace(values ...*model.Order) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a goodManyToManyOrdersTx) Delete(values ...*model.Order) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a goodManyToManyOrdersTx) Clear() error {
	return a.tx.Clear()
}

func (a goodManyToManyOrdersTx) Count() int64 {
	return a.tx.Count()
}

type goodHasOneGoodOrders struct {
	db *gorm.DB

	field.RelationField
}

func (a goodHasOneGoodOrders) Where(conds ...field.Expr) *goodHasOneGoodOrders {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a goodHasOneGoodOrders) WithContext(ctx context.Context) *goodHasOneGoodOrders {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a goodHasOneGoodOrders) Session(session *gorm.Session) *goodHasOneGoodOrders {
	a.db = a.db.Session(session)
	return &a
}

func (a goodHasOneGoodOrders) Model(m *model.Good) *goodHasOneGoodOrdersTx {
	return &goodHasOneGoodOrdersTx{a.db.Model(m).Association(a.Name())}
}

type goodHasOneGoodOrdersTx struct{ tx *gorm.Association }

func (a goodHasOneGoodOrdersTx) Find() (result *model.GoodOrder, err error) {
	return result, a.tx.Find(&result)
}

func (a goodHasOneGoodOrdersTx) Append(values ...*model.GoodOrder) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a goodHasOneGoodOrdersTx) Replace(values ...*model.GoodOrder) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a goodHasOneGoodOrdersTx) Delete(values ...*model.GoodOrder) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a goodHasOneGoodOrdersTx) Clear() error {
	return a.tx.Clear()
}

func (a goodHasOneGoodOrdersTx) Count() int64 {
	return a.tx.Count()
}

type goodDo struct{ gen.DO }

type IGoodDo interface {
	gen.SubQuery
	Debug() IGoodDo
	WithContext(ctx context.Context) IGoodDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGoodDo
	WriteDB() IGoodDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGoodDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGoodDo
	Not(conds ...gen.Condition) IGoodDo
	Or(conds ...gen.Condition) IGoodDo
	Select(conds ...field.Expr) IGoodDo
	Where(conds ...gen.Condition) IGoodDo
	Order(conds ...field.Expr) IGoodDo
	Distinct(cols ...field.Expr) IGoodDo
	Omit(cols ...field.Expr) IGoodDo
	Join(table schema.Tabler, on ...field.Expr) IGoodDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGoodDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGoodDo
	Group(cols ...field.Expr) IGoodDo
	Having(conds ...gen.Condition) IGoodDo
	Limit(limit int) IGoodDo
	Offset(offset int) IGoodDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGoodDo
	Unscoped() IGoodDo
	Create(values ...*model.Good) error
	CreateInBatches(values []*model.Good, batchSize int) error
	Save(values ...*model.Good) error
	First() (*model.Good, error)
	Take() (*model.Good, error)
	Last() (*model.Good, error)
	Find() ([]*model.Good, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Good, err error)
	FindInBatches(result *[]*model.Good, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Good) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGoodDo
	Assign(attrs ...field.AssignExpr) IGoodDo
	Joins(fields ...field.RelationField) IGoodDo
	Preload(fields ...field.RelationField) IGoodDo
	FirstOrInit() (*model.Good, error)
	FirstOrCreate() (*model.Good, error)
	FindByPage(offset int, limit int) (result []*model.Good, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGoodDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g goodDo) Debug() IGoodDo {
	return g.withDO(g.DO.Debug())
}

func (g goodDo) WithContext(ctx context.Context) IGoodDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goodDo) ReadDB() IGoodDo {
	return g.Clauses(dbresolver.Read)
}

func (g goodDo) WriteDB() IGoodDo {
	return g.Clauses(dbresolver.Write)
}

func (g goodDo) Session(config *gorm.Session) IGoodDo {
	return g.withDO(g.DO.Session(config))
}

func (g goodDo) Clauses(conds ...clause.Expression) IGoodDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goodDo) Returning(value interface{}, columns ...string) IGoodDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g goodDo) Not(conds ...gen.Condition) IGoodDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goodDo) Or(conds ...gen.Condition) IGoodDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goodDo) Select(conds ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goodDo) Where(conds ...gen.Condition) IGoodDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goodDo) Order(conds ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goodDo) Distinct(cols ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goodDo) Omit(cols ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goodDo) Join(table schema.Tabler, on ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goodDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGoodDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goodDo) RightJoin(table schema.Tabler, on ...field.Expr) IGoodDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goodDo) Group(cols ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goodDo) Having(conds ...gen.Condition) IGoodDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goodDo) Limit(limit int) IGoodDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goodDo) Offset(offset int) IGoodDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goodDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGoodDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goodDo) Unscoped() IGoodDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goodDo) Create(values ...*model.Good) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goodDo) CreateInBatches(values []*model.Good, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goodDo) Save(values ...*model.Good) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goodDo) First() (*model.Good, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) Take() (*model.Good, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) Last() (*model.Good, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) Find() ([]*model.Good, error) {
	result, err := g.DO.Find()
	return result.([]*model.Good), err
}

func (g goodDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Good, err error) {
	buf := make([]*model.Good, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goodDo) FindInBatches(result *[]*model.Good, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goodDo) Attrs(attrs ...field.AssignExpr) IGoodDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goodDo) Assign(attrs ...field.AssignExpr) IGoodDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goodDo) Joins(fields ...field.RelationField) IGoodDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g goodDo) Preload(fields ...field.RelationField) IGoodDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g goodDo) FirstOrInit() (*model.Good, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) FirstOrCreate() (*model.Good, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) FindByPage(offset int, limit int) (result []*model.Good, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g goodDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g goodDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g goodDo) Delete(models ...*model.Good) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *goodDo) withDO(do gen.Dao) *goodDo {
	g.DO = *do.(*gen.DO)
	return g
}
