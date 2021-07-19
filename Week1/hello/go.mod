module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	test v0.0.0-00010101000000-000000000000
)

replace test => ../test
