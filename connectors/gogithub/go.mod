module github.com/ditrit/gandalf/connectors/gogithub

go 1.13

require (
	github.com/ditrit/gandalf/connectors/go v0.0.0-20210324150802-22782409f699
	github.com/ditrit/gandalf/core v0.0.0-20210309085845-3de4f1d2c53d
	github.com/ditrit/gandalf/libraries/goclient v0.0.0-20210216134342-40c7d10bd6c4
	github.com/ditrit/shoset v0.0.0-20201026092509-225b8a4a5276
	github.com/google/go-github/v33 v33.0.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
)

replace github.com/ditrit/gandalf/core => ../../core
