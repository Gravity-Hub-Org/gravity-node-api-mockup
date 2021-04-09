package main

import (
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/migrations/common"
	"github.com/Gravity-Tech/gravity-node-api/model"
	"github.com/go-pg/migrations"
)

func init () {
	tableName := model.DefaultExtendedDBTableNames.CommitTable

	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			fmt.Printf("creating %v table...\n", tableName)
			_, err := db.Exec(fmt.Sprintf(
				`
				CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

				CREATE TABLE %[1]v (
					data_id serial primary key,
					tx_id integer references transactions (tx_id),
					nebula_id text,
					pulse_id integer,
                    tc_height integer,
					commit text,
					oracle_pub_key text
				);
				%v;
				`, tableName, common.CreateMaterializedViewQuery(tableName)))
			return err
		},
		func(db migrations.DB) error {
			fmt.Printf("dropping %v table...\n", tableName)
			_, err := db.Exec(fmt.Sprintf(`%v; DROP TABLE %v;`, common.DropMaterializedViewQuery(tableName), tableName))
			return err
		},
	)
}
