package constants

import "os"

var PG_HOST = os.Getenv("PG_PORT")
var PG_PORT = os.Getenv("PG_PORT")
var PG_USER = os.Getenv("PG_USER")
var PG_PASSWORD = os.Getenv("PG_PASSWORD")
var JWT_SECRET = os.Getenv("JWT_SECRET")
