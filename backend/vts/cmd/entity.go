package main

// VIEW data
type RequestObject struct {
	APIName     string      `json:"apiName"`
	APIVersion  string      `json:"apiVersion"`
	RequestID   string      `json:"requestID"`
	MessageType string      `json:"messageType"`
	Data        interface{} `json:"data"`
}

type Response struct {
	APIName     string      `json:"apiName"`
	APIVersion  string      `json:"apiVersion"`
	Timestamp   int         `json:"timestamp"`
	RequestID   string      `json:"requestID"`
	MessageType string      `json:"messageType"`
	Data        interface{} `json:"data"`
}

type ErrorResponse ErrorData

// Internal data
type ErrorData struct {
	ErrorID int    `json:"errorID"`
	Message string `json:"message"`
}

type AuthenticationTokenRequestData struct {
	PluginName      string `json:"pluginName"`
	PluginDeveloper string `json:"pluginDeveloper"`
	PluginIcon      string `json:"pluginIcon"`
}

func (atrq *AuthenticationTokenRequestData) Validate() Response {
	if atrq.PluginDeveloper == "Justin" {
		return Response{}
	}
	return Response{}
}

type AuthenticationRequestData struct {
	PluginName          string `json:"pluginName"`
	PluginDeveloper     string `json:"pluginDeveloper"`
	AuthenticationToken string `json:"authenticationToken"`
}
