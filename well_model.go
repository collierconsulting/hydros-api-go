package hydros

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"io/ioutil"
	"net/http"
	"time"
)

// WellModel Well response payload
type WellModel struct {
	*DefaultModelBase
	Serial                            string                  `json:"serial,omitempty"`
	Name                              null.String             `json:"name,omitempty"`
	StateWellID                       null.String             `json:"stateWellId,omitempty"`
	TCEQ                              null.String             `json:"tceq,omitempty"`
	Approved                          bool                    `json:"approved,omitempty"`
	Status                            *StatusModel            `json:"status,omitempty"`
	SystemID                          uint                    `json:"systemId,omitempty"`
	System                            *SystemModel            `json:"system,omitempty"`
	Location                          *LocationModel          `json:"location,omitempty"`
	SecondaryStatuses                 []*SecondaryStatusModel `json:"secondaryStatuses,omitempty"`
	Contacts                          []*ContactModel         `json:"contacts"`
	Owner                             *ContactModel           `json:"owner,omitempty"`
	Applicant                         *ContactModel           `json:"applicant,omitempty"`
	Driller                           *DrillerModel           `json:"driller,omitempty"`
	DrillerIsContact                  bool                    `json:"drillerIsContact,omitempty"`
	PumpInstaller                     *DrillerModel           `json:"pumpInstaller,omitempty"`
	PumpInstallerIsContact            bool                    `json:"pumpInstallerIsContact,omitempty"`
	Tank                              *WellTankModel          `json:"tank,omitempty"`
	Construction                      *ConstructionModel      `json:"construction,omitempty"`
	WellReplacementID                 uint                    `json:"wellReplacementId,omitempty"`
	EstimatedDrillingDate             null.Time               `json:"estimatedDrillingDate,omitempty"`
	DrillingDate                      null.Time               `json:"drillingDate,omitempty"`
	CompletionDate                    null.Time               `json:"completionDate,omitempty"`
	ApprovedDate                      null.Time               `json:"approvedDate,omitempty"`
	Confidential                      bool                    `json:"confidential,omitempty"`
	Exempt                            bool                    `json:"exempt,omitempty"`
	ExemptionType                     string                  `json:"exemptionType,omitempty"`
	ApplicationType                   string                  `json:"applicationType,omitempty"`
	InstalledByDriller                bool                    `json:"installedByDriller,omitempty"`
	UsedByOtherThanOwner              bool                    `json:"usedByOtherThanOwner,omitempty"`
	TransportedOutOfGCD               bool                    `json:"transportedOutOfGCD,omitempty"`
	TransportedOutOfGCDDescription    null.String             `json:"transportedOutOfGCDDescription,omitempty"`
	UsedByPublicWaterSystem           bool                    `json:"usedByPublicWaterSystem,omitempty"`
	RequestedMonitoring               bool                    `json:"requestedMonitoring,omitempty"`
	RequestedGrandfathered            bool                    `json:"requestedGrandfathered,omitempty"`
	RequestedExtension                bool                    `json:"requestedExtension,omitempty"`
	UseBasedExemption                 bool                    `json:"useBasedExemption,omitempty"`
	PcbExemptionIndividual            bool                    `json:"pcbExemptionIndividual,omitempty"`
	PcbExemptionWellSystem            bool                    `json:"pcbExemptionWellSystem,omitempty"`
	CertifiedBeneficial               bool                    `json:"certifiedBeneficial,omitempty"`
	CertifiedRules                    bool                    `json:"certifiedRules,omitempty"`
	LocationTransported               null.String             `json:"locationTransported,omitempty"`
	PreferredPayment                  string                  `json:"preferredPayment,omitempty"`
	BeneficialUseAgreement            bool                    `json:"beneficialUseAgreement,omitempty"`
	DistrictRulesAgreement            bool                    `json:"districtRulesAgreement,omitempty"`
	AbideRules                        bool                    `json:"abideRules,omitempty"`
	NeedsProduction                   bool                    `json:"needsProduction,omitempty"`
	Notes                             null.String             `json:"notes,omitempty"`
	WellReportTrackingNumber          null.Int                `json:"wellReportTrackingNumber,omitempty"`
	PluggingReportTrackingNumber      null.Int                `json:"pluggingReportTrackingNumber,omitempty"`
	CertifiedInfoCorrect              bool                    `json:"certifiedInfoCorrect,omitempty"`
	MonitoringWellID                  null.String             `json:"monitoringWellId,omitempty"`
	ElevationInFeet                   null.Float              `json:"elevationInFeet,omitempty"`
	RequestedRescind                  bool                    `json:"requestedRescind,omitempty"`
	CertifiedMinTractSize             bool                    `json:"certifiedMinTractSize,omitempty"`
	CertifiedDistPropertyLine         bool                    `json:"certifiedDistPropertyLine,omitempty"`
	CertifiedDistExistingWaterWell    bool                    `json:"certifiedDistExistingWaterWell,omitempty"`
	CertifiedLocation                 bool                    `json:"certifiedLocation,omitempty"`
	CertifiedPluggedCappedGuidelines  bool                    `json:"certifiedPluggedCappedGuidelines,omitempty"`
	CertifiedProvideReports           bool                    `json:"certifiedProvideReports,omitempty"`
	EstimatedAnnualTransportedGallons null.Int                `json:"estimatedAnnualTransportedGallons,omitempty"`
	EstimatedAnnualWaterProduction    null.Int                `json:"estimatedAnnualWaterProduction,omitempty"`
	WellLogReceived                   null.Time               `json:"wellLogReceived,omitempty"`
	WellUses                          []WellUse               `json:"wellUses"`
	CreatedAt                         time.Time               `json:"createdAt,omitempty"`
	UpdatedAt                         time.Time               `json:"updatedAt,omitempty"`

	_Update        func(model *WellModel, JSONMergePatch []byte) (*WellModel, error)
	_Save          func(model *WellModel) (*WellModel, error)
	_Permits       func(model *WellModel) ([]*PermitModel, error)
	_Delete        func(model *WellModel) error
	_TriggerUpdate func(model *WellModel) (*WellModel, error)
}

// WellSearchResults total and result list of found wells
type WellSearchResults struct {
	Total   int          `json:"total"`
	Results []*WellModel `json:"results"`
}

// Init Initializes spec and default backing functions for model instance
func (model *WellModel) Init(spec *ServiceSpec) *WellModel {
	model.Spec = spec

	if serviceMock, ok := spec.ModelServiceCallMocks["Save"]; ok {
		model._Save = serviceMock.MockFunc.(func(model *WellModel) (*WellModel, error))
	} else {
		model._Save = func(model *WellModel) (*WellModel, error) {
			return nil, errors.New("not implemented")
		}
	}

	if serviceMock, ok := spec.ModelServiceCallMocks["Permits"]; ok {
		model._Permits = serviceMock.MockFunc.(func(model *WellModel) ([]*PermitModel, error))
	} else {
		model._Permits = func(model *WellModel) ([]*PermitModel, error) {
			uri := fmt.Sprintf("%s/%s/%d/permits.json", model.Spec.Client.URL.String(), model.Spec.ServiceName, model.ID)
			req, err := http.NewRequest("GET", uri, nil)
			headers := model.Spec.Client.CreateHeadersFunc()
			for h := 0; h < len(headers); h++ {
				req.Header.Add(headers[h].Key, headers[h].Value)
			}

			resp, err := model.Spec.Client.HTTPClient.Do(req)
			if err != nil {
				return nil, err
			}

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			if resp.StatusCode != 200 {
				var errorResponse ErrorResponse
				err = json.Unmarshal(bodyBytes, &errorResponse)
				if err == nil && errorResponse.Message != "" {
					return nil, fmt.Errorf("%s: %s", errorResponse.Message, errorResponse.Description)
				}
				return nil, fmt.Errorf("%d error: %s", resp.StatusCode, string(bodyBytes))
			}

			var permits []*PermitModel
			err = json.Unmarshal(bodyBytes, &permits)
			if err != nil {
				return nil, err
			}
			initializedPermits := make([]*PermitModel, len(permits))
			for i, permit := range permits {
				initializedPermits[i] = permit.Init(spec.Client.Permit._ServiceSpec())
			}
			return initializedPermits, nil
		}
	}

	if serviceMock, ok := spec.ModelServiceCallMocks["Update"]; ok {
		model._Update = serviceMock.MockFunc.(func(model *WellModel, JSONMergePatch []byte) (*WellModel, error))
	} else {
		model._Update = func(model *WellModel, JSONMergePatch []byte) (*WellModel, error) {
			return nil, errors.New("not implemented")
		}
	}

	if serviceMock, ok := spec.ModelServiceCallMocks["TriggerUpdate"]; ok {
		model._TriggerUpdate = serviceMock.MockFunc.(func(model *WellModel) (*WellModel, error))
	} else {
		model._TriggerUpdate = func(model *WellModel) (*WellModel, error) {

			jsonStr, err := json.Marshal(model)
			if err != nil {
				return nil, err
			}

			uri := fmt.Sprintf("%s/%s/%d/triggerUpdate.json",
				model.Spec.Client.URL.String(), model.Spec.ServiceName, model.ID)
			req, err := http.NewRequest("PUT", uri, bytes.NewBuffer(jsonStr))
			headers := model.Spec.Client.CreateHeadersFunc()
			for h := 0; h < len(headers); h++ {
				req.Header.Add(headers[h].Key, headers[h].Value)
			}

			resp, err := model.Spec.Client.HTTPClient.Do(req)
			if err != nil {
				return nil, err
			}

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			if resp.StatusCode != 200 {
				var errorResponse ErrorResponse
				err = json.Unmarshal(bodyBytes, &errorResponse)
				if err == nil && errorResponse.Message != "" {
					return nil, fmt.Errorf("%s: %s", errorResponse.Message, errorResponse.Description)
				}
				return nil, fmt.Errorf("%d error: %s", resp.StatusCode, string(bodyBytes))
			}

			var well WellModel
			err = json.Unmarshal(bodyBytes, &well)
			if err != nil {
				return nil, err
			}
			return well.Init(model.Spec), nil
		}
	}

	if serviceMock, ok := spec.ModelServiceCallMocks["Delete"]; ok {
		model._Delete = serviceMock.MockFunc.(func(model *WellModel) error)
	} else {
		model._Delete = func(model *WellModel) error {
			return errors.New("not implemented")
		}
	}
	return model
}

// Update old model with patch data
func (model *WellModel) Update(JSONMergePatch []byte) (*WellModel, error) {
	return model._Update(model, JSONMergePatch)
}

// TriggerUpdate update well entry in search DB and take snapshot of well state
func (model *WellModel) TriggerUpdate() (*WellModel, error) {
	return model._TriggerUpdate(model)
}

// Save changed model
func (model *WellModel) Save() (*WellModel, error) {
	return model._Save(model)
}

// Permits fetch permits
func (model *WellModel) Permits() ([]*PermitModel, error) {
	return model._Permits(model)
}

// Delete model
func (model *WellModel) Delete() error {
	return model._Delete(model)
}

// StatusModel status model for well association
type StatusModel struct {
	ID     uint   `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

// LocationModel location model for well association
type LocationModel struct {
	ID                          uint        `json:"id,omitempty"`
	Latitude                    null.Float  `json:"latitude,omitempty"`
	Longitude                   null.Float  `json:"longitude,omitempty"`
	Section                     null.String `json:"section,omitempty"`
	Block                       null.String `json:"block,omitempty"`
	Elevation                   null.Float  `json:"elevation,omitempty"`
	Address1                    null.String `json:"address1,omitempty"`
	Address2                    null.String `json:"address2,omitempty"`
	County                      null.String `json:"county,omitempty"`
	City                        null.String `json:"city,omitempty"`
	State                       null.String `json:"state,omitempty"`
	PostalCode                  null.String `json:"postalCode,omitempty"`
	GpsManufacturer             null.String `json:"gpsManufacturer,omitempty"`
	GpsModel                    null.String `json:"gpsModel,omitempty"`
	QuarterQuad                 null.String `json:"quarterQuad,omitempty"`
	DistanceToPropertyLine1     null.Int    `json:"distanceToPropertyLine1,omitempty"`
	DistanceToPropertyLine2     null.Int    `json:"distanceToPropertyLine2,omitempty"`
	DistanceToPropertyLine1Type null.String `json:"distanceToPropertyLine1Type,omitempty"`
	DistanceToPropertyLine2Type null.String `json:"distanceToPropertyLine2Type,omitempty"`
	ContinuousAcredTotal        null.Float  `json:"continuousAcredTotal,omitempty"`
	DistNearestWellOnProperty   null.Float  `json:"distNearestWellOnProperty,omitempty"`
}

// ConstructionModel construction model for well association
type ConstructionModel struct {
	ID                  uint            `json:"id,omitempty"`
	CasingSize          null.Float      `json:"casingSize,omitempty"`
	CasingMaterial      null.String     `json:"casingMaterial,omitempty"`
	InsideDiameter      null.Float      `json:"insideDiameter,omitempty"`
	Depth               null.Float      `json:"depth,omitempty"`
	MaxPumpProduction   null.Int        `json:"maxPumpProduction,omitempty"`
	WithdrawalMethod    null.String     `json:"withdrawalMethod,omitempty"`
	PumpMotorSize       null.String     `json:"pumpMotorSize,omitempty"`
	PumpPowerSource     null.String     `json:"pumpPowerSource,omitempty"`
	PumpBowlSize        null.Float      `json:"pumpBowlSize,omitempty"`
	PumpBowlStages      null.Int        `json:"pumpBowlNumStages,omitempty"`
	PumpColumnLength    null.Float      `json:"pumpColumnLength,omitempty"`
	PumpDepth           null.Float      `json:"pumpDepth,omitempty"`
	ServiceConnections  null.Int        `json:"serviceConnections,omitempty"`
	IndividualsServiced null.Int        `json:"individualsServiced,omitempty"`
	DaysServicedPerYear null.Int        `json:"daysServicedPerYear,omitempty"`
	Confined            bool            `json:"confined,omitempty"`
	Screens             []*ScreenRecord `json:"screens,omitempty"`
	GamLayerAlias       *GamLayerAlias  `json:"gamLayerAlias,omitempty"`
	GamLayer            *GamLayerRecord `json:"gamLayer,omitempty"`
}

// ScreenRecord model
type ScreenRecord struct {
	ID          uint       `json:"id,omitempty"`
	TopDepth    null.Float `json:"topDepth,omitempty"`
	BottomDepth null.Float `json:"bottomDepth,omitempty"`
}

// GamLayerRecord model
type GamLayerRecord struct {
	ID   uint        `json:"id,omitempty"`
	Name null.String `json:"name,omitempty"`
}

// GamLayerAlias payload
type GamLayerAlias struct {
	ID        uint   `json:"id,omitempty"`
	LayerID   uint   `json:"layerId,omitempty"`
	Alias     string `json:"alias,omitempty"`
	LongAlias string `json:"longAlias,omitempty"`
}

// WellTankModel model
type WellTankModel struct {
	ID     uint        `json:"id,omitempty"`
	Size   null.Int    `json:"size,omitempty"`
	Volume null.Float  `json:"volume,omitempty"`
	Design null.String `json:"design,omitempty"`
}

// SystemModel model
type SystemModel struct {
	ID          uint        `json:"id,omitempty"`
	Name        null.String `json:"name,omitempty"`
	Description null.String `json:"description,omitempty"`
	CreatedAt   time.Time   `json:"createdAt,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time  `json:"deletedAt,omitempty"`
}

// SecondaryStatusModel model
type SecondaryStatusModel struct {
	ID              uint   `json:"id,omitempty"`
	SecondaryStatus string `json:"secondaryStatus,omitempty"`
}

// WellUse model
type WellUse struct {
	ID      uint   `json:"id"`
	WellUse string `json:"wellUse"`
	Exempt  bool   `json:"exempt"`
}
