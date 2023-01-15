package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wb/rest-api/internal/config"
	"wb/rest-api/internal/storage/database"
	"wb/rest-api/pkg/logging"
)

const (
	success = "success"
)

type Server struct {
	logger *logging.Logger
	DB     database.Storage
}

func (s *Server) Run(cfg config.Server) {
	s.logger.Infof("run server (%s:%s)", cfg.Host, cfg.Port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), nil); err != nil {
		s.logger.Fatal()
	}
}

func NewServer(database database.Storage, logger *logging.Logger) *Server {
	logger.Info("init server")

	server := &Server{
		logger: logger,
		DB:     database,
	}

	server.InitRoutes()

	return server
}

func (s *Server) InitRoutes() {
	http.HandleFunc("/client/list", s.ClientList)
	http.HandleFunc("/client/create", s.ClientCreate)
	http.HandleFunc("/client/update", s.ClientUpdate)
	http.HandleFunc("/client/delete", s.ClientDelete)
	http.HandleFunc("/market/list", s.MarketList)
	http.HandleFunc("/market/create", s.MarketCreate)
	http.HandleFunc("/market/update", s.MarketUpdate)
	http.HandleFunc("/market/delete", s.MarketDelete)
}

func (s *Server) ClientList(w http.ResponseWriter, r *http.Request) {
	var client database.Client
	request, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unable to read request body", s.logger)
		return
	}

	if err = json.Unmarshal(request, &client); err != nil {
		writeError(w, http.StatusBadRequest, "wrong json", s.logger)
		return
	}

	if err = client.ValidateForList(request, s.logger); err != nil {
		writeError(w, http.StatusBadRequest, "validation fail", s.logger)
		return
	}

	list, err := s.DB.GetList(client)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "get list error", s.logger)
		return
	}

	response := make([]byte, 0)
	for _, mdl := range list {
		bytesModel, err := mdl.Marshal(s.logger)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "unable to marshal model", s.logger)
			return
		}
		response = append(response, bytesModel...)
	}

	w.Write(response)
}

func (s *Server) ClientCreate(w http.ResponseWriter, r *http.Request) {
	var client database.Client
	request, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unable to read request body", s.logger)
		return
	}

	if err = json.Unmarshal(request, &client); err != nil {
		writeError(w, http.StatusBadRequest, "wrong json", s.logger)
		return
	}

	if err = client.ValidateForCreate(request, s.logger); err != nil {
		writeError(w, http.StatusBadRequest, "validation fail", s.logger)
		return
	}

	id, err := s.DB.Insert(client)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "insert error", s.logger)
		return
	}

	response, err := json.Marshal(setResponseId(id))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "unable to marshal response", s.logger)
		return
	}

	w.Write(response)
}

func (s *Server) ClientUpdate(w http.ResponseWriter, r *http.Request) {
	var client database.Client
	request, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unable to read request body", s.logger)
		return
	}

	if err = json.Unmarshal(request, &client); err != nil {
		writeError(w, http.StatusBadRequest, "wrong json", s.logger)
		return
	}

	if err = client.ValidateForUpdate(request, s.logger); err != nil {
		writeError(w, http.StatusBadRequest, "validation fail", s.logger)
		return
	}

	err = s.DB.Update(client)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "update error", s.logger)
		return
	}

	response, err := json.Marshal(setStatus(success))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "unable to marshal response", s.logger)
		return
	}

	w.Write(response)
}

func (s *Server) ClientDelete(w http.ResponseWriter, r *http.Request) {
	var client database.Client
	request, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unable to read request body", s.logger)
		return
	}

	if err = json.Unmarshal(request, &client); err != nil {
		writeError(w, http.StatusBadRequest, "wrong json", s.logger)
		return
	}

	if err = client.ValidateForDelete(request, s.logger); err != nil {
		writeError(w, http.StatusBadRequest, "validation fail", s.logger)
		return
	}

	err = s.DB.Delete(client)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "delete error", s.logger)
		return
	}

	response, err := json.Marshal(setStatus(success))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "unable to marshal response", s.logger)
		return
	}

	w.Write(response)
}

func (s *Server) MarketList(w http.ResponseWriter, r *http.Request) {
	var market database.Market
	request, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unable to read request body", s.logger)
		return
	}

	if err = json.Unmarshal(request, &market); err != nil {
		writeError(w, http.StatusBadRequest, "wrong json", s.logger)
		return
	}

	if err = market.ValidateForList(request, s.logger); err != nil {
		writeError(w, http.StatusBadRequest, "validation fail", s.logger)
		return
	}

	list, err := s.DB.GetList(market)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "get list error", s.logger)
		return
	}

	response := make([]byte, 0)
	for _, mdl := range list {
		bytesModel, err := mdl.Marshal(s.logger)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "unable to marshal model", s.logger)
			return
		}
		response = append(response, bytesModel...)
	}

	w.Write(response)
}

func (s *Server) MarketCreate(w http.ResponseWriter, r *http.Request) {
	var market database.Market
	request, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unable to read request body", s.logger)
		return
	}

	if err = json.Unmarshal(request, &market); err != nil {
		writeError(w, http.StatusBadRequest, "wrong json", s.logger)
		return
	}

	if err = market.ValidateForCreate(request, s.logger); err != nil {
		writeError(w, http.StatusBadRequest, "validation fail", s.logger)
		return
	}

	id, err := s.DB.Insert(market)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "insert error", s.logger)
		return
	}

	response, err := json.Marshal(setResponseId(id))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "unable to marshal response", s.logger)
		return
	}

	w.Write(response)
}

func (s *Server) MarketUpdate(w http.ResponseWriter, r *http.Request) {
	var market database.Market
	request, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unable to read request body", s.logger)
		return
	}

	if err = json.Unmarshal(request, &market); err != nil {
		writeError(w, http.StatusBadRequest, "wrong json", s.logger)
		return
	}

	if err = market.ValidateForUpdate(request, s.logger); err != nil {
		writeError(w, http.StatusBadRequest, "validation fail", s.logger)
		return
	}

	err = s.DB.Update(market)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "update error", s.logger)
		return
	}

	response, err := json.Marshal(setStatus(success))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "unable to marshal response", s.logger)
		return
	}

	w.Write(response)
}

func (s *Server) MarketDelete(w http.ResponseWriter, r *http.Request) {
	var market database.Market
	request, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unable to read request body", s.logger)
		return
	}

	if err = json.Unmarshal(request, &market); err != nil {
		writeError(w, http.StatusBadRequest, "wrong json", s.logger)
		return
	}

	if err = market.ValidateForDelete(request, s.logger); err != nil {
		writeError(w, http.StatusBadRequest, "validation fail", s.logger)
		return
	}

	err = s.DB.Delete(market)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "delete error", s.logger)
		return
	}

	response, err := json.Marshal(setStatus(success))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "unable to marshal response", s.logger)
		return
	}

	w.Write(response)
}

func writeError(w http.ResponseWriter, status int, msg string, logger *logging.Logger) {
	logger.Warningf("error: status-[%d]; msg-[%s]", status, msg)
	w.Write([]byte(msg))
	w.WriteHeader(status)
}

func setResponseId(id string) map[string]string {
	responseMap := make(map[string]string)
	responseMap["id"] = id
	return responseMap
}

func setStatus(status string) map[string]string {
	responseMap := make(map[string]string)
	responseMap["status"] = status
	return responseMap
}
