// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LogsLogSummary struct {
	Method     *string  `json:"method,omitempty"`
	StatusCode *float64 `json:"status_code,omitempty"`
	URL        *string  `json:"url,omitempty"`
}

type Logs struct {
	DashboardView *string         `json:"dashboard_view,omitempty"`
	LogID         *string         `json:"log_id,omitempty"`
	LogSummary    *LogsLogSummary `json:"log_summary,omitempty"`
}
