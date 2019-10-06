# CLI for search data from json file

## List commands

```
go run app.go -h
NAME:
  Simple Search - CLI for search data

USAGE:
  app [global options] command [command options] [arguments...]

VERSION:
  1.0.0

AUTHOR:
  Navcs

COMMANDS:
  search, s    Search data
  describe, s  get all search fields
  help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
  --lang value   language for the greeting (default: "english")
  --help, -h     show help
  --version, -v  print the version
```

## Describe

```
NAME:
  Simple Search describe - get all search fields

USAGE:
  Simple Search describe command [command options] [arguments...]

COMMANDS:
  organizations  search fields for organizations
  users          search fields for users
  tickets        search fields for tickets

OPTIONS:
  --help, -h  show help
```

## Search

The search command require argument clause query. It has format `field=value`. Field is one in describe if model. Value is value of field you want query.

```
go run app.go search -h
NAME:
  Simple Search search - Search data

USAGE:
  Simple Search search command [command options] [arguments...]

COMMANDS:
   organizations  search data organizations
   users          search data users
   tickets        search data tickets

OPTIONS:
   --help, -h  show help
```

## Examples

### describe

```
go run app.go describe users

USERS can be searched by any fields below

_id
url
external_id
name
alias
created_at
active
verified
shared
locale
timezone
last_login_at
email
phone
signature
organization_id
tags
suspended
role
```

### search

```
go run app.go search organizations 'tags=West'
{
  "number_of_result": 1,
  "search_result": [
    {
      "_id": 101,
      "created_at": "2016-05-21T11:10:28 -10:00",
      "details": "MegaCorp",
      "domain_names": [
        "kage.com",
        "ecratic.com",
        "endipin.com",
        "zentix.com"
      ],
      "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
      "name": "Enthaze",
      "shared_tickets": false,
      "tags": [
        "Fulton",
        "West",
        "Rodriguez",
        "Farley"
      ],
      "tickets": [
        "A Drama in Portugal",
        "A Problem in Ethiopia",
        "A Problem in Turks and Caicos Islands",
        "A Problem in Guyana"
      ],
      "url": "http://initech.tokoin.io.com/api/v2/organizations/101.json",
      "users": [
        "Loraine Pittman",
        "Francis Bailey",
        "Haley Farmer",
        "Herrera Norman"
      ]
    }
  ]
}
```
