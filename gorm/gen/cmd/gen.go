package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"os"
)

var MysqlDsn = os.Getenv("Mysql")

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./db/gen",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(ConnectDB())
	users := g.GenerateModel("users", gen.FieldRelate(field.Many2Many, "Roles", g.GenerateModel("roles"), &field.RelateConfig{
		RelateSlice: true,
		GORMTag: map[string][]string{
			"many2many": {"user_roles"},
		}}))

	roles := g.GenerateModel("roles", gen.FieldRelate(field.Many2Many, "Users", users, &field.RelateConfig{
		RelateSlice: true,
		GORMTag: map[string][]string{
			"many2many": {"user_roles"},
		},
	}))
	userRoles := g.GenerateModel("user_roles", gen.FieldRelate(field.BelongsTo, "User", users, &field.RelateConfig{
		RelatePointer: true,
		GORMTag: map[string][]string{
			"foreignKey": {"user_id"},
		},
	}), gen.FieldRelate(field.BelongsTo, "Role", roles, &field.RelateConfig{
		RelatePointer: true,
		GORMTag: map[string][]string{
			"foreignKey": {"role_id"},
		},
	}))

	orders := g.GenerateModel("orders", gen.FieldRelate(field.Many2Many, "Goods", g.GenerateModel("goods"), &field.RelateConfig{
		RelateSlice: true,
		GORMTag: map[string][]string{
			"many2many": {"good_orders"},
		},
	}), gen.FieldRelate(field.HasOne, "CustomerInfo", g.GenerateModel("customer_info"), &field.RelateConfig{
		RelatePointer: true,
	}))
	goods := g.GenerateModel("goods", gen.FieldRelate(field.Many2Many, "Orders", orders, &field.RelateConfig{
		RelateSlice: true,
		GORMTag: map[string][]string{
			"many2many": {"good_orders"},
		},
	}), gen.FieldRelate(field.HasOne, "GoodOrders", g.GenerateModel("good_orders"), &field.RelateConfig{
		GORMTag: map[string][]string{
			"foreignKey": {"good_id"},
		},
	}))

	customerInfo := g.GenerateModel("customer_info")
	goodOrders := g.GenerateModel("good_orders", gen.FieldRelate(field.BelongsTo, "Order", orders, &field.RelateConfig{
		RelatePointer: true,
		GORMTag: map[string][]string{
			"foreignKey": {"order_id"},
		},
	}), gen.FieldRelate(field.BelongsTo, "Good", goods, &field.RelateConfig{
		RelatePointer: true,
		GORMTag: map[string][]string{
			"foreignKey": {"good_id"},
		},
	}))

	g.ApplyBasic(roles, users, orders, goods, customerInfo, goodOrders, userRoles)
	g.Execute()
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	return db
}
