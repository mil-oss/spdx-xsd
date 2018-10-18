package xsdprov

// ProvEntry ... data provenance rreport item
type ProvEntry struct {
	EntryType string   `json:"entrytype"`
	Status    string   `json:"status"`
	Timestamp int64    `json:"timestamp,omitempty"`
	FilePath  string   `json:"filepath,omitempty"`
	XslPath   string   `json:"xslpath,omitempty"`
	XsdPath   string   `json:"xsdpath,omitempty"`
	Digest    string   `json:"digest,omitempty"`
	Valid     bool     `json:"valid,omitempty"`
	Message   string   `json:"message,omitempty"`
	Errors    []string `json:"errors[],omitempty"`
}

// ValidationData ... post data for validation
type ValidationData struct {
	XMLName   string `json:"xmlname,omitempty"`
	XMLPath   string `json:"xmlpath,omitempty"`
	XMLString string `json:"xmlstr"`
	XSDName   string `json:"xsdname,omitempty"`
	XSDPath   string `json:"xsdpath,omitempty"`
	XSDString string `json:"xsdstr,omitempty"`
	Valid     bool   `json:"valid,omitempty"`
}

// VerifyData ... post data for verification
type VerifyData struct {
	ID     string `json:"id"`
	Digest string `json:"digest"`
}

// TransformData ... post data for transformation
type TransformData struct {
	XMLName    string   `json:"xmlname,omitempty"`
	XMLPath    string   `json:"xmlpath,omitempty"`
	XMLString  string   `json:"xmlstr"`
	XSLName    string   `json:"xslname,omitempty"`
	XSLPath    string   `json:"xslpath,omitempty"`
	XSLString  string   `json:"xslstr,omitempty"`
	ResultPath string   `json:"resultpath,omitempty"`
	Params     []string `json:"params,omitempty"`
}

// ID ... internal id
type ID struct {
	ID string `json:"id,omitempty"`
}

// Success ... http erro response
type Success struct {
	Status  bool   `json:"status,omitempty"`
	Content string `json:"content,omitempty"`
}

// ValErr ... http error id
type ValErr struct {
	Status  bool   `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

// ValErrs ... list of Errors
type ValErrs []ValErr
