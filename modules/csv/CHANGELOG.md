# CSV Connection

## v0.2.0

- Change module name from `utils-csv` to `connection-csv`
- Add `NewReader()` for read csv string
- Remove `CsvKey` constants, move to models module (v1.0.0) instead

## v0.1.1

- Add 3 more columns. 'Auto Amount', 'Expense', and 'Income'

## v0.1.0

- Initial version
- Add csv key constants
- Add transaction transformer to csv format
- Add csv writer for write input as csv format
- Heavily depend on writer module (v1.0.0) for write data to file
