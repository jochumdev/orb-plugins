module github.com/go-orb/plugins/codecs/form

go 1.20

replace github.com/go-orb/plugins/codecs/proto => ../proto

require (
	github.com/go-orb/go-orb v0.0.0-20230709084536-48ca79fd6450
	github.com/go-playground/form/v4 v4.2.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20230711153332-06a737ee72cb // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require github.com/stretchr/testify v1.8.4
