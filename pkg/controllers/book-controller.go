package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/rellie/go-bookstore/pkg/utils"
	"github.com/rellie/go-bookstore/pkg/models"
)

var NewBook models.Book

