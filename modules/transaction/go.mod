module moneypro.kamontat.net/models-transaction

go 1.15

replace moneypro.kamontat.net/models-common => ../models

replace moneypro.kamontat.net/connection-common => ../connection

replace moneypro.kamontat.net/models-currency => ../currency

replace moneypro.kamontat.net/models-profile => ../profile

replace moneypro.kamontat.net/utils-common => ../utils

replace moneypro.kamontat.net/utils-error => ../error

replace moneypro.kamontat.net/utils-logger => ../logger

replace moneypro.kamontat.net/utils-measure => ../measure

require (
	moneypro.kamontat.net/models-common v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/models-currency v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/utils-error v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/utils-logger v0.0.0-00010101000000-000000000000
)
