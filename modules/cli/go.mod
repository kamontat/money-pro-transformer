module moneypro.kamontat.net/commandline

go 1.15

replace moneypro.kamontat.net/connection-common => ../connection

replace moneypro.kamontat.net/connection-csv => ../csv

replace moneypro.kamontat.net/utils-logger => ../logger

replace moneypro.kamontat.net/utils-error => ../error

replace moneypro.kamontat.net/utils-measure => ../measure

replace moneypro.kamontat.net/models-common => ../models

replace moneypro.kamontat.net/models-profile => ../profile

replace moneypro.kamontat.net/models-transaction => ../transaction

replace moneypro.kamontat.net/models-currency => ../currency

replace moneypro.kamontat.net/utils-common => ../utils

require (
	moneypro.kamontat.net/connection-common v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/models-profile v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/models-transaction v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/connection-csv v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/utils-error v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/utils-logger v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/utils-measure v0.0.0-00010101000000-000000000000
)
