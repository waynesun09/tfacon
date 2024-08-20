// Package common contains all shared structs(data structures) required for all connectors
package common

// GeneralUpdatedList is an updated list of object, with the prediction from TFA classifier
// each connector should have it's own UpdatedList structure and implement the
// GeneralUpdatedList interface.
type GeneralUpdatedList interface {
	GetSelf() GeneralUpdatedList
}

// TFAConfig is the tfa config stuct
type TFAConfig struct {
	Concurrent                bool
	AddAttributes             bool
	Re                        bool
	AutoFinalizeDefectType    bool
	AutoFinalizationThreshold float32
	RetryTimes                int
	Verbose                   bool
}

type (
	// TFAModel is the data structure for describing
	// the request body for TFA Classifer.
	TFAModel map[string]TFAInput
	// TFAInput is the data structure for describing
	// the three input params for TFA Classifier.
	TFAInput struct {
		ID                     string  `json:"id"`
		Project                string  `json:"project"`
		Messages               string  `json:"messages"`
		AutoFinalizeDefectType bool    `json:"finalization_flag"`
		FinalizationThreshold  float32 `json:"finalization_threshold"`
	}
)

// TFADefectTypeToSubType is map of defect type to sub type
var TFADefectTypeToSubType = map[string]PredictedSubType{
	"Automation Bug": PredictedAutomationBug,
	"Product Bug":    PredictedProductBug,
	"System Issue":   PredictedSystemBug,
	"No Defect":      PredictedNoDefect,
}

// TFADefectType is map of defect type to main support type
var TFADefectType = map[string]MainDefectType{
	"Automation Bug": AutomationBug,
	"Product Bug":    ProductBug,
	"System Issue":   SystemBug,
	"No Defect":      NoDefect,
}

// PredictedSubType is map of predicted sub type
type PredictedSubType map[string]string

// MainDefectType is map of defect type
type MainDefectType map[string]string

// PredictedSubTypes is the map of predicted types
var PredictedSubTypes = map[string]PredictedSubType{
	"PREDICTED_AUTOMATION_BUG": PredictedAutomationBug,
	"PREDICTED_SYSTEM_BUG":     PredictedSystemBug,
	"PREDICTED_PRODUCT_BUG":    PredictedProductBug,
	"PREDICTED_NO_DEFECT":      PredictedNoDefect,
}

// MainDefectTypes is map of defect types
var MainDefectTypes = map[string]MainDefectType{
	"AUTOMATION_BUG": AutomationBug,
	"SYSTEM_BUG":     SystemBug,
	"PRODUCT_BUG":    ProductBug,
	"NO_DEFECT":      NoDefect,
}

// NoDefect is map of the NO_DEFECT type
var NoDefect = MainDefectType{
	"locator":   "nd001",
	"typeRef":   "NO_DEFECT",
	"longName":  "No Defect",
	"shortName": "ND",
	"color":     "#777777",
}

// AutomationBug is map of the AUTOMATION_BUG type
var AutomationBug = MainDefectType{
	"locator":   "ab001",
	"typeRef":   "AUTOMATION_BUG",
	"longName":  "Automation Bug",
	"shortName": "AB",
	"color":     "#f7d63e",
}

// ProductBug is map of the PRODUCT_BUG type
var ProductBug = MainDefectType{
	"locator":   "pb001",
	"typeRef":   "PRODUCT_BUG",
	"longName":  "Product Bug",
	"shortName": "PB",
	"color":     "#ec3900",
}

// SystemBug is map of the SYSTEM_BUG type
var SystemBug = MainDefectType{
	"locator":   "si001",
	"typeRef":   "SYSTEM_ISSUE",
	"longName":  "System Issue",
	"shortName": "SI",
	"color":     "#0274d1",
}

// PredictedAutomationBug is map of the Predicted Automation Bug type
var PredictedAutomationBug = PredictedSubType{
	"typeRef":   "TO_INVESTIGATE",
	"longName":  "Predicted Automation Bug",
	"shortName": "TIA",
	"color":     "#ffeeaa",
}

// PredictedSystemBug is map of the Predicted System Issue type
var PredictedSystemBug = PredictedSubType{
	"typeRef":   "TO_INVESTIGATE",
	"longName":  "Predicted System Issue",
	"shortName": "TIS",
	"color":     "#aaaaff",
}

// PredictedProductBug is map of the Predicted Product Bug type
var PredictedProductBug = PredictedSubType{
	"typeRef":   "TO_INVESTIGATE",
	"longName":  "Predicted Product Bug",
	"shortName": "TIP",
	"color":     "#ffaaaa",
}

// PredictedNoDefect is map of the Predicted No Defect type
var PredictedNoDefect = PredictedSubType{
	"typeRef":   "TO_INVESTIGATE",
	"longName":  "Predicted No Defect",
	"shortName": "TND",
	"color":     "#C1BCB4",
}
