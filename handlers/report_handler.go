package handlers

import "net/http"

type ReportHandler interface {
	GetReportToday(w http.ResponseWriter, r *http.Request)
	GetReportByDate(w http.ResponseWriter, r *http.Request)
}
