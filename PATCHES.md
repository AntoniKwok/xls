# Fork of extrame/xls v0.0.1 with fixes

Based on `github.com/extrame/xls` v0.0.1 (the version nuel-furniture-api used) plus:
- `Row.Col` returns "" for a MULBLANK record with no cells instead of panicking
- `Row` returns nil for a missing row
- no `os.Exit`; `OpenReader` returns an error when the Workbook stream is missing; `WorkBook.Close`
- allocations sized from corrupt SST counts / hyperlink lengths are bounded

Upstream PRs: extrame/xls#105, #106, #107.
