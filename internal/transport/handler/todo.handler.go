package handler

import (
	"net/http"
	"simple-to-do/internal/model"
	"simple-to-do/internal/services"
	"simple-to-do/internal/transport/datatransfer"
	"simple-to-do/internal/utils/constants"
	"simple-to-do/pkg/logger"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type TodoHandler struct {
	S services.Service
}

func InitalizedTodoHandler(s services.Service) Handler {
	return &TodoHandler{
		S: s,
	}
}

func (th *TodoHandler) All(w http.ResponseWriter, re *http.Request, _ httprouter.Params) {
	t, err := th.S.FindAll(re.Context())
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusInternalServerError, nil, err)
	}
	datatransfer.Write(w, http.StatusOK, t, nil)
}

func (th *TodoHandler) Get(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusBadRequest, nil, err)
	}

	t, err := th.S.FindByID(re.Context(), id)
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusInternalServerError, nil, err)
	}

	datatransfer.Write(w, http.StatusOK, t, nil)
}

func (th *TodoHandler) Post(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	t := model.Task{}
	err := datatransfer.Bind(re, &t)
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusInternalServerError, nil, err)
	}

	t, err = th.S.Create(re.Context(), t)
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusInternalServerError, nil, err)
	}

	datatransfer.Write(w, http.StatusCreated, t, nil)
}

func (th *TodoHandler) Put(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusBadRequest, nil, err)
	}

	t := model.Task{}
	err = datatransfer.Bind(re, &t)
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusBadRequest, nil, err)
	}

	t.SetId(id)
	t, err = th.S.Update(re.Context(), t)
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusInternalServerError, nil, err)
	}

	datatransfer.Write(w, http.StatusAccepted, t, nil)
}

func (th *TodoHandler) Delete(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusBadRequest, nil, err)
	}

	err = th.S.Delete(re.Context(), id)
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusInternalServerError, nil, err)
	}

	datatransfer.Write(w, http.StatusOK, nil, nil)
}

func (th *TodoHandler) SetStatus(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	s, err := strconv.ParseBool(re.URL.Query().Get("set"))
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusBadRequest, nil, err)
	}

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusBadRequest, nil, err)
	}

	t, err := th.S.UpdateStatus(re.Context(), id, s)
	if err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, http.StatusInternalServerError, nil, err)
	}

	datatransfer.Write(w, http.StatusOK, t, nil)
}
