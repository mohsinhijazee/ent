// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number", Type: field.TypeString},
		{Name: "owner_id", Type: field.TypeInt, Default: 0},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cards_users_cards",
				Columns:    []*schema.Column{CardsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "currency", Type: field.TypeEnum, Enums: []string{"USD", "EUR", "ILS"}},
		{Name: "time", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "completed", "failed"}},
		{Name: "card_id", Type: field.TypeInt},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payments_cards_payments",
				Columns:    []*schema.Column{PaymentsColumns[6]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "payment_status_time",
				Unique:  false,
				Columns: []*schema.Column{PaymentsColumns[5], PaymentsColumns[3]},
			},
		},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "best_friend_id", Type: field.TypeUUID, Unique: true, Nullable: true, Default: "00000000-0000-0000-0000-000000000000"},
		{Name: "owner_id", Type: field.TypeInt, Default: 0},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pets_pets_best_friend",
				Columns:    []*schema.Column{PetsColumns[2]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "pets_users_owner",
				Columns:    []*schema.Column{PetsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "pet_name_owner_id",
				Unique:  true,
				Columns: []*schema.Column{PetsColumns[1], PetsColumns[3]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeFloat64},
		{Name: "name", Type: field.TypeString},
		{Name: "tags", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardsTable,
		PaymentsTable,
		PetsTable,
		UsersTable,
	}
)

func init() {
	CardsTable.ForeignKeys[0].RefTable = UsersTable
	CardsTable.Annotation = &entsql.Annotation{}
	CardsTable.Annotation.Checks = map[string]string{
		"number_length": "(LENGTH(`number`) = 16)",
	}
	PaymentsTable.ForeignKeys[0].RefTable = CardsTable
	PaymentsTable.Annotation = &entsql.Annotation{}
	PaymentsTable.Annotation.Checks = map[string]string{
		"amount_positive": "(`amount` > 0)",
	}
	PetsTable.ForeignKeys[0].RefTable = PetsTable
	PetsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.Annotation = &entsql.Annotation{
		Check: "age > 0",
	}
	UsersTable.Annotation.Checks = map[string]string{
		"name_not_empty": "name <> ''",
	}
}
