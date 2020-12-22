module kamontat.net/money-pro-cli

go 1.15

replace kamontat.net/money-pro-datasource => ../datasource

replace kamontat.net/money-pro-writer => ../writer

replace kamontat.net/money-pro-models => ../models

replace kamontat.net/money-pro-utils => ../utils

require (
	kamontat.net/money-pro-datasource v0.0.0-00010101000000-000000000000
	kamontat.net/money-pro-writer v0.0.0-00010101000000-000000000000
)
