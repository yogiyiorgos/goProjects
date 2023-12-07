package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/yogiyiorgos/bookstore-API/utils"
	"github.com/yogiyiorgos/bookstore-API/models"
)

var NewBook models.Book
