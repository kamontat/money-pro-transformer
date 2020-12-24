module moneypro.kamontat.net/writer

go 1.15

replace moneypro.kamontat.net/models-common => ../models

replace moneypro.kamontat.net/models-profile => ../profile

replace moneypro.kamontat.net/utils-common => ../utils

replace moneypro.kamontat.net/utils-logger => ../logger

replace moneypro.kamontat.net/utils-error => ../error

require (
	moneypro.kamontat.net/models-common v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/models-profile v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/utils-error v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/utils-logger v0.0.0-00010101000000-000000000000
)
