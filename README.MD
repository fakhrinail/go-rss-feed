# Go RSS Feeds

> This repo is made for the purpose of learning Go

## Table of contents

-   [General info](#general-info)
-   [Technologies](#technologies)
-   [How to Install](#how-to-install)
-   [How to Run](#how-to-run)
-   [Features](#features)

## General info

This is a backend app for learning to build REST API in Go. This app will collect RSS feeds every so often and save it for the users to see.

## Technologies

- Go
- PostgreSQL
- SQLC

## How to Install

```
go mod vendor
```

## How to Run

First, copy the contents in .env.example and paste it in your own `.env` file, change it to suite your database setup

Next, run the migrations using goose by using the command below

```
goose postgres [YOUR_CONNECTION_STRING] up
```

Now you can run the app and try it for yourself!

```
go build && ./go-rss-feed
```

## Upcoming Features
- Unit testing will be added soon!
