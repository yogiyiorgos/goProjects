package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yogiyiorgos/goProjects/bookstore-API/models"
	"github.com/yogiyiorgos/goProjects/bookstore-API/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book
