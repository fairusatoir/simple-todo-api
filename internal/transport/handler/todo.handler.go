package handler

import (
	"net/http"
	"simple-to-do/internal/model"
	"simple-to-do/internal/services"
	"simple-to-do/internal/transport/datatransfer"
	"simple-to-do/internal/utils/constants"
	"simple-to-do/pkg/logger"
	"simple-to-do/pkg/validator"
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
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	datatransfer.Write(w, datatransfer.Response(http.StatusOK, t))

}

func (th *TodoHandler) Get(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	t, err := th.S.FindByID(re.Context(), id)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusNotFound, err))
		return
	}

	datatransfer.Write(w, datatransfer.Response(http.StatusOK, t))

}

func (th *TodoHandler) Post(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	t := model.Task{}

	err := datatransfer.Bind(re, &t)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	if err := validator.ValidatePayloads(t); err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	t, err = th.S.Create(re.Context(), t)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	datatransfer.Write(w, datatransfer.Response(http.StatusCreated, t))
}

func (th *TodoHandler) Put(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	t := model.Task{}
	err = datatransfer.Bind(re, &t)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	if err := validator.ValidatePayloads(t); err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	t.SetId(id)

	t, err = th.S.Update(re.Context(), t)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		switch err {
		case constants.Err404:
			datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusNotFound, err))
		default:
			datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	datatransfer.Write(w, datatransfer.Response(http.StatusOK, t))
}

func (th *TodoHandler) Delete(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	err = th.S.Delete(re.Context(), id)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		switch err {
		case constants.Err404:
			datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusNotFound, err))
		default:
			datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	datatransfer.Write(w, datatransfer.Response(http.StatusOK, nil))
}

func (th *TodoHandler) SetStatus(w http.ResponseWriter, re *http.Request, p httprouter.Params) {
	s, err := strconv.ParseBool(re.URL.Query().Get("set"))
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	t, err := th.S.UpdateStatus(re.Context(), id, s)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
		datatransfer.Write(w, datatransfer.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	datatransfer.Write(w, datatransfer.Response(http.StatusOK, t))
}
