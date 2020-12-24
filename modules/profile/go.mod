module moneypro.kamontat.net/models-profile

go 1.15

replace moneypro.kamontat.net/models-common => ../models

replace moneypro.kamontat.net/utils-logger => ../logger

require (
	moneypro.kamontat.net/models-common v0.0.0-00010101000000-000000000000
	moneypro.kamontat.net/utils-logger v0.0.0-00010101000000-000000000000
)
