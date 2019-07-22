module modules

go 1.12

replace aaa.com => ./pkg/src/aaa.com

require aaa.com v0.0.0-00010101000000-000000000000 // indirect
