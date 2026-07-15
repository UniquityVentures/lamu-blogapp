# lago-blogapp

A reference implementation of lago, with the expected directory structure for projects.

This implementation will be building a complete Blogging Web App which will demonstrate most of the features of the framework

This implementation will provide a sample for the following:
- Authentication
- CRUD of models
- Patching of Core and bundled plugins

At a glance, the file structure is supposed to look as follows:
```
<project root>
├── config.toml
├── go.mod
├── go.sum
├── main.go
└── plugins
    └── <plugin name>
        ├── templates
        │   └── <go html/template compatible files for rendering html directly>
        ├── migrations
        │   └── <migration sql files created by goose>
        ├── app.go
        ├── config.go
        ├── pages.go
        ├── migrations.go
        ├── routes.go
        ├── models.go
        ├── commands.go
        └── views.go
```
