package hydros

import "time"

// WellModel Well response payload
type WellModel struct {
	*DefaultModelBase
	Serial                           *string                 `json:"serial,omitempty"`
	Name                             *string                 `json:"name,omitempty"`
	StateWellID                      *string                 `json:"stateWellId,omitempty"`
	TCEQ                             *string                 `json:"tceq,omitempty"`
	Approved                         *bool                   `json:"approved,omitempty"`
	Status                           *StatusModel            `json:"status,omitempty"`
	SystemID                         *uint                   `json:"systemId,omitempty"`
	System                           *SystemModel            `json:"system,omitempty,omitempty"`
	Location                         *LocationModel          `json:"location,omitempty"`
	SecondaryStatuses                []*SecondaryStatusModel `json:"secondaryStatuses,omitempty"`
	Owner                            *ContactModel           `json:"owner,omitempty"`
	Applicant                        *ContactModel           `json:"applicant,omitempty,omitempty"`
	Driller                          *DrillerModel           `json:"driller,omitempty,omitempty"`
	DrillerIsContact                 *bool                   `json:"drillerIsContact,omitempty"`
	PumpInstaller                    *DrillerModel           `json:"pumpInstaller,omitempty,omitempty"`
	PumpInstallerIsContact           *bool                   `json:"pumpInstallerIsContact,omitempty"`
	Tank                             *WellTankModel          `json:"tank,omitempty,omitempty"`
	Construction                     *ConstructionModel      `json:"construction,omitempty,omitempty"`
	WellReplacementID                *uint                   `json:"wellReplacementId,omitempty"`
	EstimatedDrillingDate            *time.Time              `json:"estimatedDrillingDate,omitempty"`
	DrillingDate                     *time.Time              `json:"drillingDate,omitempty"`
	CompletionDate                   *time.Time              `json:"completionDate,omitempty"`
	ApprovedDate                     *time.Time              `json:"approvedDate,omitempty"`
	Confidential                     *bool                   `json:"confidential,omitempty"`
	Exempt                           *bool                   `json:"exempt,omitempty"`
	ExemptionType                    string                  `json:"exemptionType,omitempty"`
	ApplicationType                  string                  `json:"applicationType,omitempty"`
	InstalledByDriller               *bool                   `json:"installedByDriller,omitempty"`
	UsedByOtherThanOwner             *bool                   `json:"usedByOtherThanOwner,omitempty"`
	TransportedOutOfGCD              *bool                   `json:"transportedOutOfGCD,omitempty"`
	TransportedOutOfGCDDescription   *string                 `json:"transportedOutOfGCDDescription,omitempty"`
	UsedByPublicWaterSystem          *bool                   `json:"usedByPublicWaterSystem,omitempty"`
	RequestedMonitoring              *bool                   `json:"requestedMonitoring,omitempty"`
	RequestedGrandfathered           *bool                   `json:"requestedGrandfathered,omitempty"`
	RequestedExtension               *bool                   `json:"requestedExtension,omitempty"`
	UseBasedExemption                *bool                   `json:"useBasedExemption,omitempty"`
	PcbExemptionIndividual           *bool                   `json:"pcbExemptionIndividual,omitempty"`
	PcbExemptionWellSystem           *bool                   `json:"pcbExemptionWellSystem,omitempty"`
	CertifiedBeneficial              *bool                   `json:"certifiedBeneficial,omitempty"`
	CertifiedRules                   *bool                   `json:"certifiedRules,omitempty"`
	LocationTransported              *string                 `json:"locationTransported,omitempty"`
	PreferredPayment                 string                  `json:"preferredPayment,omitempty"`
	BeneficialUseAgreement           *bool                   `json:"beneficialUseAgreement,omitempty"`
	DistrictRulesAgreement           *bool                   `json:"districtRulesAgreement,omitempty"`
	AbideRules                       *bool                   `json:"abideRules,omitempty"`
	NeedsProduction                  *bool                   `json:"needsProduction,omitempty"`
	Notes                            *string                 `json:"notes,omitempty"`
	WellReportTrackingNumber         *int                    `json:"wellReportTrackingNumber,omitempty"`
	PluggingReportTrackingNumber     *int                    `json:"pluggingReportTrackingNumber,omitempty"`
	CertifiedInfoCorrect             *bool                   `json:"certifiedInfoCorrect,omitempty"`
	MonitoringWellID                 *string                 `json:"monitoringWellId,omitempty"`
	ElevationInFeet                  *float64                `json:"elevationInFeet,omitempty"`
	RequestedRescind                 *bool                   `json:"requestedRescind,omitempty"`
	CertifiedMinTractSize            *bool                   `json:"certifiedMinTractSize,omitempty"`
	CertifiedDistPropertyLine        *bool                   `json:"certifiedDistPropertyLine,omitempty"`
	CertifiedDistExistingWaterWell   *bool                   `json:"certifiedDistExistingWaterWell,omitempty"`
	CertifiedLocation                *bool                   `json:"certifiedLocation,omitempty"`
	CertifiedPluggedCappedGuidelines *bool                   `json:"certifiedPluggedCappedGuidelines,omitempty"`
	CertifiedProvideReports          *bool                   `json:"certifiedProvideReports,omitempty"`
	EstimatedAnnualWaterProduction   *int                    `json:"estimatedAnnualWaterProduction,omitempty"`
	WellLogReceived                  *time.Time              `json:"wellLogReceived,omitempty"`
	CreatedAt                        *time.Time              `json:"createdAt,omitempty"`
	UpdatedAt                        *time.Time              `json:"updatedAt,omitempty"`

	_Update func(model *WellModel) (*WellModel, error)
	_Delete func() error
}

// Init Initializes spec and default backing functions for model instance
func (model *WellModel) Init(spec *ServiceSpec) *WellModel {
	model.Spec = spec

	model._Update = func(model *WellModel) (*WellModel, error) {
		return nil, nil
	}
	model._Delete = func() error {
		return nil
	}
	return model
}

// Update old model with new
func (model *WellModel) Update(updatedModel *WellModel) (*WellModel, error) {
	return model._Update(updatedModel)
}

// Delete model
func (model *WellModel) Delete() error {
	return model._Delete()
}

// StatusModel status model for well association
type StatusModel struct {
	ID        uint       `json:"id,omitempty"`
	Status    *string    `json:"status,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// LocationModel location model for well association
type LocationModel struct {
	ID                          uint       `json:"id,omitempty"`
	Latitude                    *float64   `json:"latitude,omitempty"`
	Longitude                   *float64   `json:"longitude,omitempty"`
	Section                     *string    `json:"section,omitempty"`
	Block                       *string    `json:"block,omitempty"`
	Elevation                   *float64   `json:"elevation,omitempty"`
	Address1                    *string    `json:"address1,omitempty"`
	Address2                    *string    `json:"address2,omitempty"`
	County                      *string    `json:"county,omitempty"`
	City                        *string    `json:"city,omitempty"`
	State                       *string    `json:"state,omitempty"`
	PostalCode                  *string    `json:"postalCode,omitempty"`
	GpsManufacturer             *string    `json:"gpsManufacturer,omitempty"`
	GpsModel                    *string    `json:"gpsModel,omitempty"`
	QuarterQuad                 *string    `json:"quarterQuad,omitempty"`
	DistanceToPropertyLine1     *int       `json:"distanceToPropertyLine1,omitempty"`
	DistanceToPropertyLine2     *int       `json:"distanceToPropertyLine2,omitempty"`
	DistanceToPropertyLine1Type *string    `json:"distanceToPropertyLine1Type,omitempty"`
	DistanceToPropertyLine2Type *string    `json:"distanceToPropertyLine2Type,omitempty"`
	ContinuousAcredTotal        *float64   `json:"continuousAcredTotal,omitempty"`
	DistNearestWellOnProperty   *float64   `json:"distNearestWellOnProperty,omitempty"`
	CreatedAt                   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt                   *time.Time `json:"updatedAt,omitempty"`
}

// ConstructionModel construction model for well association
type ConstructionModel struct {
	ID                  uint            `json:"id,omitempty"`
	CasingSize          *float64        `json:"casingSize,omitempty"`
	CasingMaterial      *string         `json:"casingMaterial,omitempty"`
	InsideDiameter      *float64        `json:"insideDiameter,omitempty"`
	Depth               *float64        `json:"depth,omitempty"`
	MaxPumpProduction   *int            `json:"maxPumpProduction,omitempty"`
	WithdrawalMethod    *string         `json:"withdrawalMethod,omitempty"`
	PumpMotorSize       *string         `json:"pumpMotorSize,omitempty"`
	PumpPowerSource     *string         `json:"pumpPowerSource,omitempty"`
	PumpBowlSize        *float64        `json:"pumpBowlSize,omitempty"`
	PumpBowlStages      *int            `json:"pumpBowlNumStages,omitempty"`
	PumpColumnLength    *float64        `json:"pumpColumnLength,omitempty"`
	PumpDepth           *float64        `json:"pumpDepth,omitempty"`
	ServiceConnections  *int            `json:"serviceConnections,omitempty"`
	IndividualsServiced *int            `json:"individualsServiced,omitempty"`
	DaysServicedPerYear *int            `json:"daysServicedPerYear,omitempty"`
	Confined            *bool           `json:"confined,omitempty"`
	Screens             []*ScreenRecord `json:"screens,omitempty"`
	GamLayerAlias       *GamLayerAlias  `json:"gamLayerAlias,omitempty,omitempty"`
	GamLayer            *GamLayerRecord `json:"gamLayer,omitempty,omitempty"`
	CreatedAt           *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt           *time.Time      `json:"updatedAt,omitempty"`
}

// ScreenRecord db model
type ScreenRecord struct {
	ID          uint       `json:"id,omitempty"`
	TopDepth    *float64   `json:"topDepth,omitempty"`
	BottomDepth *float64   `json:"bottomDepth,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type GamLayerRecord struct {
	ID   uint    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GamLayerAlias payload
type GamLayerAlias struct {
	ID        uint       `json:"id,omitempty"`
	LayerID   uint       `json:"layerId,omitempty"`
	Alias     *string    `json:"alias,omitempty"`
	LongAlias *string    `json:"longAlias,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type WellTankModel struct {
	ID        uint       `json:"id,omitempty"`
	Size      *int       `json:"size,omitempty"`
	Volume    *float64   `json:"volume,omitempty"`
	Design    *string    `json:"design,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// SystemModel db model
type SystemModel struct {
	ID          uint       `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

// SecondaryStatusModel db model
type SecondaryStatusModel struct {
	ID              uint       `json:"id,omitempty"`
	SecondaryStatus *string    `json:"secondaryStatus,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`
}
