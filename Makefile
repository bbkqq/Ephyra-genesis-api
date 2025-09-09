
api:
	@#mv go.mod.bak go.mod
	@for file in $(shell find idl -maxdepth 1 -name "*.proto" ! -name "api.proto"); do \
		cwgo server --type HTTP --server_name Ephyra-genesis-api --module Ephyra-genesis-api --idl $$file; \
	done
	@find hertz_gen -type f -print0 | xargs -0 -I{} sh -c 'sed -i "" "s/,omitempty//g" "{}"'
	@#mv go.mod go.mod.bak
	@echo "gen-service done"


db:
	@#
	cwgo  model --db_type mysql --out_dir storage/mysql/dao --sql_dir storage/sql --unittest true --type_tag true --nullable true
	@echo "gen-db done"

all: api db