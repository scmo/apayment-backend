package tvd_old

import (
	"encoding/xml"
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1"

// NewAnimalTracingPortType creates an initializes a AnimalTracingPortType.
func NewAnimalTracingPortType(cli *soap.Client) AnimalTracingPortType {
	return &animalTracingPortType{cli}
}

// AnimalTracingPortType was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AnimalTracingPortType interface {
	// CheckCattleForDisposalContribution was auto-generated from WSDL.
	CheckCattleForDisposalContribution(α *CheckCattleForDisposalContribution) (β *CheckCattleForDisposalContributionResponse, err error)

	// CheckCattlesForPassport was auto-generated from WSDL.
	CheckCattlesForPassport(α *CheckCattlesForPassport) (β *CheckCattlesForPassportResponse, err error)

	// DeleteAnimalHusbandryMemberships was auto-generated from WSDL.
	DeleteAnimalHusbandryMemberships(α *DeleteAnimalHusbandryMemberships) (β *DeleteAnimalHusbandryMembershipsResponse, err error)

	// DeleteCattleNotifications was auto-generated from WSDL.
	DeleteCattleNotifications(α *DeleteCattleNotifications) (β *DeleteCattleNotificationsResponse, err error)

	// DeleteEarTagOrder was auto-generated from WSDL.
	DeleteEarTagOrder(α *DeleteEarTagOrder) (β *DeleteEarTagOrderResponse, err error)

	// DeleteLabelEarTagOrder was auto-generated from WSDL.
	DeleteLabelEarTagOrder(α *DeleteLabelEarTagOrder) (β *DeleteLabelEarTagOrderResponse, err error)

	// DisableDataAccess was auto-generated from WSDL.
	DisableDataAccess(α *DisableDataAccess) (β *DisableDataAccessResponse, err error)

	// EnableDataAccess was auto-generated from WSDL.
	EnableDataAccess(α *EnableDataAccess) (β *EnableDataAccessResponse, err error)

	// GetAnimalHusbandryAddress was auto-generated from WSDL.
	GetAnimalHusbandryAddress(α *GetAnimalHusbandryAddress) (β *GetAnimalHusbandryAddressResponse, err error)

	// GetAnimalHusbandryAddressV2 was auto-generated from WSDL.
	GetAnimalHusbandryAddressV2(α *GetAnimalHusbandryAddressV2) (β *GetAnimalHusbandryAddressV2Response, err error)

	// GetAnimalHusbandryDetail was auto-generated from WSDL.
	GetAnimalHusbandryDetail(α *GetAnimalHusbandryDetail) (β *GetAnimalHusbandryDetailResponse, err error)

	// GetAnimalHusbandryMemberships was auto-generated from WSDL.
	GetAnimalHusbandryMemberships(α *GetAnimalHusbandryMemberships) (β *GetAnimalHusbandryMembershipsResponse, err error)

	// GetCattleDetail was auto-generated from WSDL.
	GetCattleDetail(α *GetCattleDetail) (β *GetCattleDetailResponse, err error)

	// GetCattleDetailV2 was auto-generated from WSDL.
	GetCattleDetailV2(α *GetCattleDetailV2) (β *GetCattleDetailV2Response, err error)

	// GetCattleEarTags was auto-generated from WSDL.
	GetCattleEarTags(α *GetCattleEarTags) (β *GetCattleEarTagsResponse, err error)

	// GetCattleHistory was auto-generated from WSDL.
	GetCattleHistory(α *GetCattleHistory) (β *GetCattleHistoryResponse, err error)

	// GetCattleHistoryV2 was auto-generated from WSDL.
	GetCattleHistoryV2(α *GetCattleHistoryV2) (β *GetCattleHistoryV2Response, err error)

	// GetCattleLivestock was auto-generated from WSDL.
	GetCattleLivestock(α *GetCattleLivestock) (β *GetCattleLivestockResponse, err error)

	// GetCattleLivestockV2 was auto-generated from WSDL.
	GetCattleLivestockV2(α *GetCattleLivestockV2) (β *GetCattleLivestockV2Response, err error)

	// GetCattleMovements was auto-generated from WSDL.
	GetCattleMovements(α *GetCattleMovements) (β *GetCattleMovementsResponse, err error)

	// GetCattleOffsprings was auto-generated from WSDL.
	GetCattleOffsprings(α *GetCattleOffsprings) (β *GetCattleOffspringsResponse, err error)

	// GetCattleStatus was auto-generated from WSDL.
	GetCattleStatus(α *GetCattleStatus) (β *GetCattleStatusResponse, err error)

	// GetCattleStatusV2 was auto-generated from WSDL.
	GetCattleStatusV2(α *GetCattleStatusV2) (β *GetCattleStatusV2Response, err error)

	// GetCattlesPerLeavingDate was auto-generated from WSDL.
	GetCattlesPerLeavingDate(α *GetCattlesPerLeavingDate) (β *GetCattlesPerLeavingDateResponse, err error)

	// GetCodes was auto-generated from WSDL.
	GetCodes(α *GetCodes) (β *GetCodesResponse, err error)

	// GetEarTagOrders was auto-generated from WSDL.
	GetEarTagOrders(α *GetEarTagOrders) (β *GetEarTagOrdersResponse, err error)

	// GetEquidDetail was auto-generated from WSDL.
	GetEquidDetail(α *GetEquidDetail) (β *GetEquidDetailResponse, err error)

	// GetEquidLivestock was auto-generated from WSDL.
	GetEquidLivestock(α *GetEquidLivestock) (β *GetEquidLivestockResponse, err error)

	// GetEquidOwnershipList was auto-generated from WSDL.
	GetEquidOwnershipList(α *GetEquidOwnershipList) (β *GetEquidOwnershipListResponse, err error)

	// GetEquidSlaughterList was auto-generated from WSDL.
	GetEquidSlaughterList(α *GetEquidSlaughterList) (β *GetEquidSlaughterListResponse, err error)

	// GetEquidsWithNotificationsOfMO was auto-generated from WSDL.
	GetEquidsWithNotificationsOfMO(α *GetEquidsWithNotificationsOfMO) (β *GetEquidsWithNotificationsOfMOResponse, err error)

	// GetFarmManager was auto-generated from WSDL.
	GetFarmManager(α *GetFarmManager) (β *GetFarmManagerResponse, err error)

	// GetFarmers was auto-generated from WSDL.
	GetFarmers(α *GetFarmers) (β *GetFarmersResponse, err error)

	// GetLabelEarTagOrders was auto-generated from WSDL.
	GetLabelEarTagOrders(α *GetLabelEarTagOrders) (β *GetLabelEarTagOrdersResponse, err error)

	// GetMembershipForOrganisation was auto-generated from WSDL.
	GetMembershipForOrganisation(α *GetMembershipForOrganisation) (β *GetMembershipForOrganisationResponse, err error)

	// GetPersonAddress was auto-generated from WSDL.
	GetPersonAddress(α *GetPersonAddress) (β *GetPersonAddressResponse, err error)

	// GetPersonIdentifiers was auto-generated from WSDL.
	GetPersonIdentifiers(α *GetPersonIdentifiers) (β *GetPersonIdentifiersResponse, err error)

	// GetPigArrivalNotificationForBreedingOrganisation was auto-generated
	// from WSDL.
	GetPigArrivalNotificationForBreedingOrganisation(α *GetPigArrivalNotificationForBreedingOrganisation) (β *GetPigArrivalNotificationForBreedingOrganisationResponse, err error)

	// GetPoultryBarnInNotifications was auto-generated from WSDL.
	GetPoultryBarnInNotifications(α *GetPoultryBarnInNotifications) (β *GetPoultryBarnInNotificationsResponse, err error)

	// GetRegisteredGenera was auto-generated from WSDL.
	GetRegisteredGenera(α *GetRegisteredGenera) (β *GetRegisteredGeneraResponse, err error)

	// GetSmallAnimalSlaughterList was auto-generated from WSDL.
	GetSmallAnimalSlaughterList(α *GetSmallAnimalSlaughterList) (β *GetSmallAnimalSlaughterListResponse, err error)

	// SearchEquid was auto-generated from WSDL.
	SearchEquid(α *SearchEquid) (β *SearchEquidResponse, err error)

	// Version was auto-generated from WSDL.
	Version(α *Version) (β *VersionResponse, err error)

	// WriteAnimalClassificationNotification was auto-generated from
	// WSDL.
	WriteAnimalClassificationNotification(α *WriteAnimalClassificationNotification) (β *WriteAnimalClassificationNotificationResponse, err error)

	// WriteAnimalClassificationNotificationV2 was auto-generated from
	// WSDL.
	WriteAnimalClassificationNotificationV2(α *WriteAnimalClassificationNotificationV2) (β *WriteAnimalClassificationNotificationV2Response, err error)

	// WriteCattleArrivalBatchNotification was auto-generated from
	// WSDL.
	WriteCattleArrivalBatchNotification(α *WriteCattleArrivalBatchNotification) (β *WriteCattleArrivalBatchNotificationResponse, err error)

	// WriteCattleBirthNotification was auto-generated from WSDL.
	WriteCattleBirthNotification(α *WriteCattleBirthNotification) (β *WriteCattleBirthNotificationResponse, err error)

	// WriteCattleChangeNameNotification was auto-generated from WSDL.
	WriteCattleChangeNameNotification(α *WriteCattleChangeNameNotification) (β *WriteCattleChangeNameNotificationResponse, err error)

	// WriteCattleDaystayBatchNotification was auto-generated from
	// WSDL.
	WriteCattleDaystayBatchNotification(α *WriteCattleDaystayBatchNotification) (β *WriteCattleDaystayBatchNotificationResponse, err error)

	// WriteCattleDeathBirthNotification was auto-generated from WSDL.
	WriteCattleDeathBirthNotification(α *WriteCattleDeathBirthNotification) (β *WriteCattleDeathBirthNotificationResponse, err error)

	// WriteCattleDeceasedNotification was auto-generated from WSDL.
	WriteCattleDeceasedNotification(α *WriteCattleDeceasedNotification) (β *WriteCattleDeceasedNotificationResponse, err error)

	// WriteCattleDeformationNotification was auto-generated from WSDL.
	WriteCattleDeformationNotification(α *WriteCattleDeformationNotification) (β *WriteCattleDeformationNotificationResponse, err error)

	// WriteCattleExportNotification was auto-generated from WSDL.
	WriteCattleExportNotification(α *WriteCattleExportNotification) (β *WriteCattleExportNotificationResponse, err error)

	// WriteCattleImportAfterExportNotification was auto-generated
	// from WSDL.
	WriteCattleImportAfterExportNotification(α *WriteCattleImportAfterExportNotification) (β *WriteCattleImportAfterExportNotificationResponse, err error)

	// WriteCattleLeavingBatchNotification was auto-generated from
	// WSDL.
	WriteCattleLeavingBatchNotification(α *WriteCattleLeavingBatchNotification) (β *WriteCattleLeavingBatchNotificationResponse, err error)

	// WriteCattlePassportOrders was auto-generated from WSDL.
	WriteCattlePassportOrders(α *WriteCattlePassportOrders) (β *WriteCattlePassportOrdersResponse, err error)

	// WriteCattleSlaughterBatchNotification was auto-generated from
	// WSDL.
	WriteCattleSlaughterBatchNotification(α *WriteCattleSlaughterBatchNotification) (β *WriteCattleSlaughterBatchNotificationResponse, err error)

	// WriteCattleSlaughterBatchNotificationV2 was auto-generated from
	// WSDL.
	WriteCattleSlaughterBatchNotificationV2(α *WriteCattleSlaughterBatchNotificationV2) (β *WriteCattleSlaughterBatchNotificationV2Response, err error)

	// WriteCattleTypeOfUseNotification was auto-generated from WSDL.
	WriteCattleTypeOfUseNotification(α *WriteCattleTypeOfUseNotification) (β *WriteCattleTypeOfUseNotificationResponse, err error)

	// WriteCattleYardSlaughterNotification was auto-generated from
	// WSDL.
	WriteCattleYardSlaughterNotification(α *WriteCattleYardSlaughterNotification) (β *WriteCattleYardSlaughterNotificationResponse, err error)

	// WriteEquidAcquireOwnershipNotification was auto-generated from
	// WSDL.
	WriteEquidAcquireOwnershipNotification(α *WriteEquidAcquireOwnershipNotification) (β *WriteEquidAcquireOwnershipNotificationResponse, err error)

	// WriteEquidBasicDataNotification was auto-generated from WSDL.
	WriteEquidBasicDataNotification(α *WriteEquidBasicDataNotification) (β *WriteEquidBasicDataNotificationResponse, err error)

	// WriteEquidBirthNotification was auto-generated from WSDL.
	WriteEquidBirthNotification(α *WriteEquidBirthNotification) (β *WriteEquidBirthNotificationResponse, err error)

	// WriteEquidCastrationNotification was auto-generated from WSDL.
	WriteEquidCastrationNotification(α *WriteEquidCastrationNotification) (β *WriteEquidCastrationNotificationResponse, err error)

	// WriteEquidCedeOwnershipNotification was auto-generated from
	// WSDL.
	WriteEquidCedeOwnershipNotification(α *WriteEquidCedeOwnershipNotification) (β *WriteEquidCedeOwnershipNotificationResponse, err error)

	// WriteEquidDeathNotification was auto-generated from WSDL.
	WriteEquidDeathNotification(α *WriteEquidDeathNotification) (β *WriteEquidDeathNotificationResponse, err error)

	// WriteEquidImportAfterExportNotification was auto-generated from
	// WSDL.
	WriteEquidImportAfterExportNotification(α *WriteEquidImportAfterExportNotification) (β *WriteEquidImportAfterExportNotificationResponse, err error)

	// WriteEquidImportNotification was auto-generated from WSDL.
	WriteEquidImportNotification(α *WriteEquidImportNotification) (β *WriteEquidImportNotificationResponse, err error)

	// WriteEquidMembershipOrganisationsNotification was auto-generated
	// from WSDL.
	WriteEquidMembershipOrganisationsNotification(α *WriteEquidMembershipOrganisationsNotification) (β *WriteEquidMembershipOrganisationsNotificationResponse, err error)

	// WriteEquidNewChipNotification was auto-generated from WSDL.
	WriteEquidNewChipNotification(α *WriteEquidNewChipNotification) (β *WriteEquidNewChipNotificationResponse, err error)

	// WriteEquidOrderBasePassNotification was auto-generated from
	// WSDL.
	WriteEquidOrderBasePassNotification(α *WriteEquidOrderBasePassNotification) (β *WriteEquidOrderBasePassNotificationResponse, err error)

	// WriteEquidPassIssuedNotification was auto-generated from WSDL.
	WriteEquidPassIssuedNotification(α *WriteEquidPassIssuedNotification) (β *WriteEquidPassIssuedNotificationResponse, err error)

	// WriteEquidPassIssuerPermissionNotification was auto-generated
	// from WSDL.
	WriteEquidPassIssuerPermissionNotification(α *WriteEquidPassIssuerPermissionNotification) (β *WriteEquidPassIssuerPermissionNotificationResponse, err error)

	// WriteEquidRelocationNotification was auto-generated from WSDL.
	WriteEquidRelocationNotification(α *WriteEquidRelocationNotification) (β *WriteEquidRelocationNotificationResponse, err error)

	// WriteEquidSlaughterNotification was auto-generated from WSDL.
	WriteEquidSlaughterNotification(α *WriteEquidSlaughterNotification) (β *WriteEquidSlaughterNotificationResponse, err error)

	// WriteEquidTypeOfUseNotification was auto-generated from WSDL.
	WriteEquidTypeOfUseNotification(α *WriteEquidTypeOfUseNotification) (β *WriteEquidTypeOfUseNotificationResponse, err error)

	// WriteEquidWithersClassNotification was auto-generated from WSDL.
	WriteEquidWithersClassNotification(α *WriteEquidWithersClassNotification) (β *WriteEquidWithersClassNotificationResponse, err error)

	// WriteGroupSlaughterMovement was auto-generated from WSDL.
	WriteGroupSlaughterMovement(α *WriteGroupSlaughterMovement) (β *WriteGroupSlaughterMovementResponse, err error)

	// WriteGroupSlaughterMovementV2 was auto-generated from WSDL.
	WriteGroupSlaughterMovementV2(α *WriteGroupSlaughterMovementV2) (β *WriteGroupSlaughterMovementV2Response, err error)

	// WriteNewEarTagOrder was auto-generated from WSDL.
	WriteNewEarTagOrder(α *WriteNewEarTagOrder) (β *WriteNewEarTagOrderResponse, err error)

	// WriteNewLabelEarTagOrder was auto-generated from WSDL.
	WriteNewLabelEarTagOrder(α *WriteNewLabelEarTagOrder) (β *WriteNewLabelEarTagOrderResponse, err error)

	// WritePigEntryMovement was auto-generated from WSDL.
	WritePigEntryMovement(α *WritePigEntryMovement) (β *WritePigEntryMovementResponse, err error)

	// WritePigEntryMovementV2 was auto-generated from WSDL.
	WritePigEntryMovementV2(α *WritePigEntryMovementV2) (β *WritePigEntryMovementV2Response, err error)

	// WritePigSlaughterMovement was auto-generated from WSDL.
	WritePigSlaughterMovement(α *WritePigSlaughterMovement) (β *WritePigSlaughterMovementResponse, err error)

	// WritePoultryBarnInNotification was auto-generated from WSDL.
	WritePoultryBarnInNotification(α *WritePoultryBarnInNotification) (β *WritePoultryBarnInNotificationResponse, err error)

	// WritePoultrySlaughterNotification was auto-generated from WSDL.
	WritePoultrySlaughterNotification(α *WritePoultrySlaughterNotification) (β *WritePoultrySlaughterNotificationResponse, err error)

	// WriteReplacementBatchEarTagOrder was auto-generated from WSDL.
	WriteReplacementBatchEarTagOrder(α *WriteReplacementBatchEarTagOrder) (β *WriteReplacementBatchEarTagOrderResponse, err error)

	// WriteReplacementEarTagOrder was auto-generated from WSDL.
	WriteReplacementEarTagOrder(α *WriteReplacementEarTagOrder) (β *WriteReplacementEarTagOrderResponse, err error)
}

// DateTime in WSDL format.
type DateTime string

// EnumAnimalHusbandryType was auto-generated from WSDL.
type EnumAnimalHusbandryType string

// Validate validates EnumAnimalHusbandryType.
func (v EnumAnimalHusbandryType) Validate() bool {
	for _, vv := range []string{
		"Unknown",
		"ProductionHusbandry",
		"HerdsmanHusbandry",
		"CoOperationPastureHusbandry",
		"SummeringHusbandry",
		"CoOperationHusbandry",
		"LiveStockGroup",
		"Population",
		"LiveStockDealerEnterprise",
		"MigratoryHerd",
		"MedicalCenter",
		"SlaughterEnterprise",
		"MarketAuctionExhibition",
		"CoOperationBranchHusbandry",
		"NonComercial",
		"OeLNCommunity",
		"AllYearHusbandry",
		"EnterpriseHusbandry",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumAnimalHusbandryTypeOfUse was auto-generated from WSDL.
type EnumAnimalHusbandryTypeOfUse string

// Validate validates EnumAnimalHusbandryTypeOfUse.
func (v EnumAnimalHusbandryTypeOfUse) Validate() bool {
	for _, vv := range []string{
		"Mixed",
		"Cow",
		"MilkCow",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumArea was auto-generated from WSDL.
type EnumArea string

// Validate validates EnumArea.
func (v EnumArea) Validate() bool {
	for _, vv := range []string{
		"Unknown",
		"Mountain",
		"Summering",
		"Valley",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumCattleRace was auto-generated from WSDL.
type EnumCattleRace string

// Validate validates EnumCattleRace.
func (v EnumCattleRace) Validate() bool {
	for _, vv := range []string{
		"Undefined",
		"Holstein",
		"Rotfleckvieh",
		"Jersey",
		"Braunvieh",
		"Angler",
		"Original_Braunvieh",
		"Red_Holstein",
		"Kreuzung",
		"Hinterwaelder",
		"Charolais",
		"Limousin",
		"Weissblaue_Belgier",
		"Blonde_d_Aquitaine",
		"Maine_Anjou",
		"Salers",
		"Montbeliard",
		"Aubrac",
		"Gasconne",
		"Piemontese",
		"Chianina",
		"Romagnola",
		"Marchigiana",
		"INRA95",
		"Angus",
		"Hereford",
		"Highland_Cattle",
		"Galloway",
		"Guernsey",
		"Swiss_Fleckvieh",
		"Luing",
		"Kiwicross",
		"Normande",
		"Ayrshire",
		"Abondance",
		"Grauvieh",
		"Dexter",
		"Bazadaise",
		"Tuxer",
		"Murnau_Werdenfelser",
		"Daenische_Rotbunte",
		"Schwedische_Rotbunte",
		"Norwegische_Rotbunte",
		"Pinzgauer",
		"Simmental",
		"Dahomey",
		"Tarentaise",
		"Vosgienne",
		"Texas_Longhorn",
		"Gelbvieh",
		"Wasserbueffel",
		"Bison",
		"Yak",
		"Eringer",
		"Evolene",
		"Zebu",
		"Wagyu",
		"Shorthorn",
		"Parthenaise",
		"Rotes_Hoehenvieh",
		"Valdostana",
		"Zwergzebu",
		"Other",
		"Bordelaise",
		"Pustertaler_Sprinzen",
		"LowLine",
		"Welsh_Black",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumCountry was auto-generated from WSDL.
type EnumCountry string

// Validate validates EnumCountry.
func (v EnumCountry) Validate() bool {
	for _, vv := range []string{
		"Undefined",
		"ABW",
		"AFG",
		"AGO",
		"AIA",
		"ALA",
		"ALB",
		"AND",
		"ANT",
		"ARE",
		"ARG",
		"ARM",
		"ASM",
		"ATA",
		"ATG",
		"AUS",
		"AUT",
		"AZE",
		"BDI",
		"BEL",
		"BEN",
		"BFA",
		"BGD",
		"BGR",
		"BHR",
		"BHS",
		"BIH",
		"BLM",
		"BLR",
		"BLZ",
		"BMU",
		"BOL",
		"BRA",
		"BRB",
		"BRN",
		"BTN",
		"BWA",
		"CAF",
		"CAN",
		"CHE",
		"CHL",
		"CHN",
		"CIV",
		"CMR",
		"COD",
		"COG",
		"COK",
		"COL",
		"COM",
		"CPV",
		"CRI",
		"CUB",
		"CXR",
		"CYM",
		"CYP",
		"CZE",
		"DEU",
		"DJI",
		"DMA",
		"DNK",
		"DOM",
		"DZA",
		"ECU",
		"EGY",
		"ERI",
		"ESH",
		"ESP",
		"EST",
		"ETH",
		"FIN",
		"FJI",
		"FLK",
		"FRA",
		"FRO",
		"FSM",
		"GAB",
		"GBR",
		"GEO",
		"GGY",
		"GHA",
		"GIB",
		"GIN",
		"GLP",
		"GMB",
		"GNB",
		"GNQ",
		"GRC",
		"GRD",
		"GRL",
		"GTM",
		"GUF",
		"GUM",
		"GUY",
		"HKG",
		"HND",
		"HRV",
		"HTI",
		"HUN",
		"IDN",
		"IMN",
		"IND",
		"IRL",
		"IRN",
		"IRQ",
		"ISL",
		"ISR",
		"ITA",
		"JAM",
		"JEY",
		"JOR",
		"JPN",
		"KAZ",
		"KEN",
		"KGZ",
		"KHM",
		"KIR",
		"KNA",
		"KOR",
		"KWT",
		"LAO",
		"LBN",
		"LBR",
		"LBY",
		"LCA",
		"LIE",
		"LKA",
		"LSO",
		"LTU",
		"LUX",
		"LVA",
		"MAC",
		"MAF",
		"MAR",
		"MCO",
		"MDA",
		"MDG",
		"MDV",
		"MEX",
		"MHL",
		"MKD",
		"MLI",
		"MLT",
		"MMR",
		"MNE",
		"MNG",
		"MNP",
		"MOZ",
		"MRT",
		"MSR",
		"MTQ",
		"MUS",
		"MWI",
		"MYS",
		"MYT",
		"NAM",
		"NCL",
		"NER",
		"NFK",
		"NGA",
		"NIC",
		"NIU",
		"NLD",
		"NOR",
		"NPL",
		"NRU",
		"NZL",
		"OMN",
		"PAK",
		"PAN",
		"PCN",
		"PER",
		"PHL",
		"PLW",
		"PNG",
		"POL",
		"PRI",
		"PRK",
		"PRT",
		"PRY",
		"PSE",
		"PYF",
		"QAT",
		"REU",
		"ROU",
		"RUS",
		"RWA",
		"SAU",
		"SDN",
		"SEN",
		"SGP",
		"SHN",
		"SJM",
		"SLB",
		"SLE",
		"SLV",
		"SMR",
		"SOM",
		"SPM",
		"SRB",
		"STP",
		"SUR",
		"SVK",
		"SVN",
		"SWE",
		"SWZ",
		"SYC",
		"SYR",
		"TCA",
		"TCD",
		"TGO",
		"THA",
		"TJK",
		"TKL",
		"TKM",
		"TLS",
		"TON",
		"TTO",
		"TUN",
		"TUR",
		"TUV",
		"TZA",
		"UGA",
		"UKR",
		"URY",
		"USA",
		"UZB",
		"VAT",
		"VCT",
		"VEN",
		"VGB",
		"VIR",
		"VNM",
		"VUT",
		"WLF",
		"WSM",
		"YEM",
		"ZAF",
		"ZMB",
		"ZWE",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidBreed was auto-generated from WSDL.
type EnumEquidBreed string

// Validate validates EnumEquidBreed.
func (v EnumEquidBreed) Validate() bool {
	for _, vv := range []string{
		"Undefined",
		"Achal_Tekkiner",
		"Achal_Tekkiner_Partbred",
		"Aegidienberger",
		"American_Miniature_Horse",
		"American_Saddlebred",
		"Andalusier",
		"Angloarab_Vollblut_Vorbuch",
		"Angloaraber",
		"Anglo_Araber_Vorbuch",
		"Angloarabisches_Halbblut",
		"Angloarabisches_Halbblut_Vorbuch",
		"Angloarabisches_Vollblut",
		"Anglo_Normaenner",
		"Appaloosa",
		"Araber_Berber",
		"Arabisches_Vollblut",
		"Arabo_Friesen",
		"Baden_Wuerttemberg",
		"Bayerisches_Warmblut",
		"Belgisches_Warmblut",
		"Berber",
		"BER_MA",
		"Bosniaken",
		"Brandenburger",
		"British_Sporthorse",
		"British_Spotted_Pony",
		"Camargue",
		"CH_Kleinpferd",
		"CH_Sportpony",
		"Clydesdale",
		"Cob_Normand",
		"Connemara",
		"Cream_Color_Schweiz",
		"Creme_Colors",
		"Cria_Caballar",
		"Criollo",
		"Dartmoor",
		"Deutsches_Classic_Pony",
		"Deutsches_Reitpony",
		"Deutsches_Warmblut",
		"Duelmener",
		"Englisches_Vollblut",
		"Andere_Eselrasse",
		"Exmoor",
		"Fell_Pony",
		"Finnisches_Warmblut",
		"Fjord",
		"Friese",
		"Haflinger_Mischrasse",
		"Hannoveraner",
		"Hessen",
		"Highland_Pony",
		"Highland_Pony_Carron",
		"Hispano",
		"Holsteiner",
		"Irlaender",
		"Islandpferd",
		"Kladruber",
		"Knabstrupper",
		"Kreuzung_ZVCH",
		"Lettisches_Warmblut",
		"Lewitzer",
		"Lewitzschecke",
		"Lipizzaner",
		"Littauisches_Warmblut",
		"Lusitano",
		"Luxemburgisches_Warmblut",
		"Maulesel",
		"Maultier",
		"Mazedonier_",
		"Mazedonier_Partbred",
		"Mecklenburger",
		"Merens",
		"Missouri_Foxtrotter",
		"Morgan",
		"New_Forest",
		"Noriker",
		"Oesterr_Warmblut",
		"Oldenburger_Pferd",
		"Orlow",
		"Ostfriese",
		"Paint",
		"Paint_Horse",
		"Palomino",
		"Partbred",
		"Partbredaraber",
		"Paso",
		"Paso_Colombiano",
		"Paso_Fino",
		"Paso_Peruano",
		"Percheron",
		"Pinto",
		"Pleven",
		"Polen_Trakehner",
		"Pottok",
		"Przewalski",
		"Pura_raza_espanola",
		"Quarter_Horse",
		"Rheinland",
		"Russisches_Warmblut",
		"Sachsen",
		"Sachsen_Anhaltiner",
		"Selle_Francais",
		"Shagya_Araber",
		"Shetlandpony",
		"Shire_Horse",
		"Slowenisches_Warmblut",
		"Special_Color_Schweiz",
		"Sporthorse_Brasilien",
		"Sporthorse_Mexico",
		"Tennessee_Walking_Horse",
		"Thueringer",
		"Tinker",
		"Traber",
		"Trait_Comtois",
		"Trakehner",
		"Warmblut_Rheinland_Pfalz_Saar",
		"Welsh",
		"Welsh_Cob_Sektion_D_WD",
		"Welsh_Mountain_Pony_WA",
		"Welsh_Partbred_WK",
		"Welsh_Pony_Cob_Typ_WC",
		"Welsh_Riding_Pony_WB",
		"Westfale",
		"Wuerttemberger",
		"Zangersheide",
		"Zweibruecken",
		"Freiberger",
		"Haflinger",
		"Schweizer_Warmblut",
		"Araber",
		"Pony",
		"Warmblut",
		"Vollblut",
		"Kaltblut",
		"Kreuzung",
		"Andere",
		"Hollaendisches_Warmblut_KWPN",
		"Daenisches_Warmblut",
		"Polnisches_Warmblut",
		"Ungarisches_Warmblut",
		"Tschechisches_Warmblut",
		"Italienisches_Warmblut",
		"CH_Sportpferd",
		"Halbblut",
		"Schwedisches_Halbblut",
		"Daenisches_Reitpony",
		"Bardigiano",
		"Kabardiner",
		"Argentinisches_Polopony",
		"Ungarisches_Vollblut",
		"Half_Saddlebred",
		"Tigerscheck",
		"Tigerscheck_Shetlandtyp",
		"Mini_Shetlandpony",
		"Irish_Cob",
		"Curly_Horse",
		"Cruzado_Iberico",
		"Dales_Pony",
		"Deutsches_Partbred_Shetland_Pony",
		"Deutsches_Sportpferd",
		"Englisches_Reitpony",
		"Hannoverschers_Halbblut",
		"Internationales_Oldenburger_Springpferd",
		"Kiger_Mustang",
		"Kleines_Deutsches_Pony",
		"Kleines_Deutsches_Reitpferd",
		"Leonharder",
		"Leutstettener_Pferd",
		"Mangalarga_Marchadores",
		"NRW_Reitpferd",
		"Ostfriese_Alt_Oldenburger",
		"Paso_Iberoamericano",
		"Portugiesisches_Sport_pferd",
		"Rheinisch_Deutsches_Kaltblut",
		"Rottaler",
		"Schwarzwaelder_Kaltblut",
		"Spanisches_Sportpferd",
		"Sueddeutsches_Kaltblut",
		"Warlander",
		"Poitou_Esel",
		"Grand_Noir_du_Berry",
		"Andalusischer_Riesenesel",
		"Katalanischer_Riesenesel",
		"American_Miniatur_Esel",
		"Ragusana",
		"Amiata_Esel",
		"Martina_Franca_Esel",
		"Contentin_Esel",
		"Normandie_Esel",
		"Hausesel",
		"Schweizer_Zuchtpferd",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidColor was auto-generated from WSDL.
type EnumEquidColor string

// Validate validates EnumEquidColor.
func (v EnumEquidColor) Validate() bool {
	for _, vv := range []string{
		"Unclassified",
		"Rappe",
		"Braun",
		"Fuchs",
		"Schimmel",
		"Schecke",
		"Grau",
		"OtherColor",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidDeathNotificationType was auto-generated from WSDL.
type EnumEquidDeathNotificationType string

// Validate validates EnumEquidDeathNotificationType.
func (v EnumEquidDeathNotificationType) Validate() bool {
	for _, vv := range []string{
		"Decease",
		"Euthanize",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidLocationChangeType was auto-generated from WSDL.
type EnumEquidLocationChangeType string

// Validate validates EnumEquidLocationChangeType.
func (v EnumEquidLocationChangeType) Validate() bool {
	for _, vv := range []string{
		"InSwitzerland",
		"OutOfSwitzerland",
		"IntoSwitzerland",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidNotificationState was auto-generated from WSDL.
type EnumEquidNotificationState string

// Validate validates EnumEquidNotificationState.
func (v EnumEquidNotificationState) Validate() bool {
	for _, vv := range []string{
		"Ok",
		"LocationChangePending",
		"OwnershipAcquisitionPending",
		"LocationChangeAndAcquisitionPending",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidNotificationType was auto-generated from WSDL.
type EnumEquidNotificationType string

// Validate validates EnumEquidNotificationType.
func (v EnumEquidNotificationType) Validate() bool {
	for _, vv := range []string{
		"Undefined",
		"BasicDataChange",
		"Slaughter",
		"FirstRegistration",
		"Birth",
		"Import",
		"LegacyVerbalDescription",
		"GraphicalDescription",
		"OwnershipCession_National",
		"OwnershipAcquisition",
		"ChipImplantation",
		"Pass",
		"LocationChange",
		"NewUeln",
		"TypeOfUsageChange",
		"DataRetrieval",
		"EquidMembership",
		"LocationChangePending",
		"PassOrder",
		"Decease_Euthanize",
		"Castration",
		"PassIssuerPermission",
		"Withers",
		"OwnershipCession_International",
		"ImportAfterExport",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidPassIssuingState was auto-generated from WSDL.
type EnumEquidPassIssuingState string

// Validate validates EnumEquidPassIssuingState.
func (v EnumEquidPassIssuingState) Validate() bool {
	for _, vv := range []string{
		"NotStarted",
		"Ordered",
		"OrderedAndIssued",
		"JustIssued",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidPassOrderType was auto-generated from WSDL.
type EnumEquidPassOrderType string

// Validate validates EnumEquidPassOrderType.
func (v EnumEquidPassOrderType) Validate() bool {
	for _, vv := range []string{
		"Not_Migratable",
		"Hardcopy",
		"PDF",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidPassType was auto-generated from WSDL.
type EnumEquidPassType string

// Validate validates EnumEquidPassType.
func (v EnumEquidPassType) Validate() bool {
	for _, vv := range []string{
		"Unknown",
		"FirstPass",
		"Substitute",
		"Duplicate",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidSpecies was auto-generated from WSDL.
type EnumEquidSpecies string

// Validate validates EnumEquidSpecies.
func (v EnumEquidSpecies) Validate() bool {
	for _, vv := range []string{
		"Undefined",
		"Pferd",
		"Esel",
		"Maultier",
		"Maulesel",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidTypeOfUse was auto-generated from WSDL.
type EnumEquidTypeOfUse string

// Validate validates EnumEquidTypeOfUse.
func (v EnumEquidTypeOfUse) Validate() bool {
	for _, vv := range []string{
		"Undefined",
		"Nutztier",
		"Heimtier",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumEquidWithersClass was auto-generated from WSDL.
type EnumEquidWithersClass string

// Validate validates EnumEquidWithersClass.
func (v EnumEquidWithersClass) Validate() bool {
	for _, vv := range []string{
		"LessOrEqualThan148cm",
		"GreaterThan148cm",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumGender was auto-generated from WSDL.
type EnumGender string

// Validate validates EnumGender.
func (v EnumGender) Validate() bool {
	for _, vv := range []string{
		"Female",
		"Male",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumGenus was auto-generated from WSDL.
type EnumGenus string

// Validate validates EnumGenus.
func (v EnumGenus) Validate() bool {
	for _, vv := range []string{
		"Unknow",
		"Cattle",
		"Pig",
		"Sheep",
		"Goat",
		"Camelid",
		"Game",
		"Equid",
		"Bee",
		"Fish",
		"Poultry",
		"Others",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumOrderStatus was auto-generated from WSDL.
type EnumOrderStatus string

// Validate validates EnumOrderStatus.
func (v EnumOrderStatus) Validate() bool {
	for _, vv := range []string{
		"New",
		"InDevelopment",
		"OrderedAtSupplier",
		"OrderedFromStock",
		"Send",
		"Closed",
		"Deleted",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumPigCategory was auto-generated from WSDL.
type EnumPigCategory string

// Validate validates EnumPigCategory.
func (v EnumPigCategory) Validate() bool {
	for _, vv := range []string{
		"Other",
		"Weaner",
		"FeederPig",
		"PigForSlaughter",
		"Sow",
		"Boar",
		"Gilt",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumPoultryType was auto-generated from WSDL.
type EnumPoultryType string

// Validate validates EnumPoultryType.
func (v EnumPoultryType) Validate() bool {
	for _, vv := range []string{
		"Other",
		"Chicken",
		"Hen",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumPoultryUsageReason was auto-generated from WSDL.
type EnumPoultryUsageReason string

// Validate validates EnumPoultryUsageReason.
func (v EnumPoultryUsageReason) Validate() bool {
	for _, vv := range []string{
		"Undefined",
		"LayingHen",
		"MastPoultry",
		"MastTurkey",
		"BreedingAnimalLayingLine",
		"BreedingAnimalMastLine",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumRole was auto-generated from WSDL.
type EnumRole string

// Validate validates EnumRole.
func (v EnumRole) Validate() bool {
	for _, vv := range []string{
		"System",
		"Administrator",
		"Identitas",
		"Farmer",
		"Owner",
		"IdentityVerifier",
		"PassportIssuer",
		"Government",
		"SlaughterHouse",
		"BreedingOrganisation",
		"MemberOrganisation",
		"MandateTaker",
		"Guest",
		"ChipImplanter",
		"Accountant",
		"CattlePassportIssuer",
		"ScientificalDataRetriever",
		"SlaughterInitiator",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// EnumZone was auto-generated from WSDL.
type EnumZone string

// Validate validates EnumZone.
func (v EnumZone) Validate() bool {
	for _, vv := range []string{
		"Unknown",
		"Hills",
		"Mountain01",
		"Mountain02",
		"Mountain03",
		"Mountain04",
		"Summering",
		"Valley",
		"ForeignAncestralZoneForSummeringAnh",
		"ForeignNonAncestralZoneForSummeringAnh",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// Char was auto-generated from WSDL.
type char int

// Duration was auto-generated from WSDL.
type duration string

// Guid was auto-generated from WSDL.
type guid string

// AnimalClassificationData was auto-generated from WSDL.
type AnimalClassificationData struct {
	MessageID                   int64    `xml:"MessageID,omitempty" json:"MessageID,omitempty" yaml:"MessageID,omitempty"`
	EarTagNumber                string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	TVDNumberOrigin             int      `xml:"TVDNumberOrigin,omitempty" json:"TVDNumberOrigin,omitempty" yaml:"TVDNumberOrigin,omitempty"`
	SlaughterDate               DateTime `xml:"SlaughterDate,omitempty" json:"SlaughterDate,omitempty" yaml:"SlaughterDate,omitempty"`
	TVDNumberSlaughterInitiator int      `xml:"TVDNumberSlaughterInitiator,omitempty" json:"TVDNumberSlaughterInitiator,omitempty" yaml:"TVDNumberSlaughterInitiator,omitempty"`
	Genus                       int      `xml:"Genus,omitempty" json:"Genus,omitempty" yaml:"Genus,omitempty"`
	Weight                      float64  `xml:"Weight,omitempty" json:"Weight,omitempty" yaml:"Weight,omitempty"`
	ClassifierNumber            int      `xml:"ClassifierNumber,omitempty" json:"ClassifierNumber,omitempty" yaml:"ClassifierNumber,omitempty"`
	ClassifierEquipmentID       string   `xml:"ClassifierEquipmentID,omitempty" json:"ClassifierEquipmentID,omitempty" yaml:"ClassifierEquipmentID,omitempty"`
	ContractorNumberSlaughter   string   `xml:"ContractorNumberSlaughter,omitempty" json:"ContractorNumberSlaughter,omitempty" yaml:"ContractorNumberSlaughter,omitempty"`
	SlaughterCategory           string   `xml:"SlaughterCategory,omitempty" json:"SlaughterCategory,omitempty" yaml:"SlaughterCategory,omitempty"`
	Beefiness                   string   `xml:"Beefiness,omitempty" json:"Beefiness,omitempty" yaml:"Beefiness,omitempty"`
	FatTissue                   string   `xml:"FatTissue,omitempty" json:"FatTissue,omitempty" yaml:"FatTissue,omitempty"`
	MFA                         int      `xml:"MFA,omitempty" json:"MFA,omitempty" yaml:"MFA,omitempty"`
}

// AnimalClassificationDataV2 was auto-generated from WSDL.
type AnimalClassificationDataV2 struct {
	MessageID                   int64    `xml:"MessageID,omitempty" json:"MessageID,omitempty" yaml:"MessageID,omitempty"`
	EarTagNumber                string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	TVDNumberOrigin             int      `xml:"TVDNumberOrigin,omitempty" json:"TVDNumberOrigin,omitempty" yaml:"TVDNumberOrigin,omitempty"`
	SlaughterDate               DateTime `xml:"SlaughterDate,omitempty" json:"SlaughterDate,omitempty" yaml:"SlaughterDate,omitempty"`
	TVDNumberSlaughterInitiator int      `xml:"TVDNumberSlaughterInitiator,omitempty" json:"TVDNumberSlaughterInitiator,omitempty" yaml:"TVDNumberSlaughterInitiator,omitempty"`
	Genus                       int      `xml:"Genus,omitempty" json:"Genus,omitempty" yaml:"Genus,omitempty"`
	Weight                      float64  `xml:"Weight,omitempty" json:"Weight,omitempty" yaml:"Weight,omitempty"`
	ClassifierNumber            int      `xml:"ClassifierNumber,omitempty" json:"ClassifierNumber,omitempty" yaml:"ClassifierNumber,omitempty"`
	ClassifierEquipmentID       string   `xml:"ClassifierEquipmentID,omitempty" json:"ClassifierEquipmentID,omitempty" yaml:"ClassifierEquipmentID,omitempty"`
	ContractorNumberSlaughter   string   `xml:"ContractorNumberSlaughter,omitempty" json:"ContractorNumberSlaughter,omitempty" yaml:"ContractorNumberSlaughter,omitempty"`
	SlaughterCategory           string   `xml:"SlaughterCategory,omitempty" json:"SlaughterCategory,omitempty" yaml:"SlaughterCategory,omitempty"`
	Beefiness                   string   `xml:"Beefiness,omitempty" json:"Beefiness,omitempty" yaml:"Beefiness,omitempty"`
	FatTissue                   string   `xml:"FatTissue,omitempty" json:"FatTissue,omitempty" yaml:"FatTissue,omitempty"`
	MFA                         int      `xml:"MFA,omitempty" json:"MFA,omitempty" yaml:"MFA,omitempty"`
	LValue                      float64  `xml:"LValue,omitempty" json:"LValue,omitempty" yaml:"LValue,omitempty"`
}

// AnimalHusbandryAddressResult was auto-generated from WSDL.
type AnimalHusbandryAddressResult struct {
	Result                   *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	PostData                 *HusbandryResult  `xml:"PostData,omitempty" json:"PostData,omitempty" yaml:"PostData,omitempty"`
	IsActive                 bool              `xml:"IsActive,omitempty" json:"IsActive,omitempty" yaml:"IsActive,omitempty"`
	MunicipalityName         string            `xml:"MunicipalityName,omitempty" json:"MunicipalityName,omitempty" yaml:"MunicipalityName,omitempty"`
	CantonShortname          string            `xml:"CantonShortname,omitempty" json:"CantonShortname,omitempty" yaml:"CantonShortname,omitempty"`
	CoordinateX              int               `xml:"CoordinateX,omitempty" json:"CoordinateX,omitempty" yaml:"CoordinateX,omitempty"`
	CoordinateY              int               `xml:"CoordinateY,omitempty" json:"CoordinateY,omitempty" yaml:"CoordinateY,omitempty"`
	Altitude                 int               `xml:"Altitude,omitempty" json:"Altitude,omitempty" yaml:"Altitude,omitempty"`
	CantonAnimalHusbandryKey string            `xml:"CantonAnimalHusbandryKey,omitempty" json:"CantonAnimalHusbandryKey,omitempty" yaml:"CantonAnimalHusbandryKey,omitempty"`
	CantonPersonKey          string            `xml:"CantonPersonKey,omitempty" json:"CantonPersonKey,omitempty" yaml:"CantonPersonKey,omitempty"`
	BurNumber                string            `xml:"BurNumber,omitempty" json:"BurNumber,omitempty" yaml:"BurNumber,omitempty"`
	AnimalHusbandryType      int               `xml:"AnimalHusbandryType,omitempty" json:"AnimalHusbandryType,omitempty" yaml:"AnimalHusbandryType,omitempty"`
	MunicipalityNumber       int               `xml:"MunicipalityNumber,omitempty" json:"MunicipalityNumber,omitempty" yaml:"MunicipalityNumber,omitempty"`
	TypeOfUse                int               `xml:"TypeOfUse,omitempty" json:"TypeOfUse,omitempty" yaml:"TypeOfUse,omitempty"`
}

// AnimalHusbandryAddressResultV2 was auto-generated from WSDL.
type AnimalHusbandryAddressResultV2 struct {
	Result                   *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	PostData                 *HusbandryResult  `xml:"PostData,omitempty" json:"PostData,omitempty" yaml:"PostData,omitempty"`
	IsActive                 bool              `xml:"IsActive,omitempty" json:"IsActive,omitempty" yaml:"IsActive,omitempty"`
	MunicipalityName         string            `xml:"MunicipalityName,omitempty" json:"MunicipalityName,omitempty" yaml:"MunicipalityName,omitempty"`
	CantonShortname          string            `xml:"CantonShortname,omitempty" json:"CantonShortname,omitempty" yaml:"CantonShortname,omitempty"`
	CoordinateX              int               `xml:"CoordinateX,omitempty" json:"CoordinateX,omitempty" yaml:"CoordinateX,omitempty"`
	CoordinateY              int               `xml:"CoordinateY,omitempty" json:"CoordinateY,omitempty" yaml:"CoordinateY,omitempty"`
	Altitude                 int               `xml:"Altitude,omitempty" json:"Altitude,omitempty" yaml:"Altitude,omitempty"`
	CantonAnimalHusbandryKey string            `xml:"CantonAnimalHusbandryKey,omitempty" json:"CantonAnimalHusbandryKey,omitempty" yaml:"CantonAnimalHusbandryKey,omitempty"`
	CantonPersonKey          string            `xml:"CantonPersonKey,omitempty" json:"CantonPersonKey,omitempty" yaml:"CantonPersonKey,omitempty"`
	BurNumber                string            `xml:"BurNumber,omitempty" json:"BurNumber,omitempty" yaml:"BurNumber,omitempty"`
	AnimalHusbandryType      int               `xml:"AnimalHusbandryType,omitempty" json:"AnimalHusbandryType,omitempty" yaml:"AnimalHusbandryType,omitempty"`
	MunicipalityNumber       int               `xml:"MunicipalityNumber,omitempty" json:"MunicipalityNumber,omitempty" yaml:"MunicipalityNumber,omitempty"`
	TypeOfUse                int               `xml:"TypeOfUse,omitempty" json:"TypeOfUse,omitempty" yaml:"TypeOfUse,omitempty"`
	IsMountain               bool              `xml:"IsMountain,omitempty" json:"IsMountain,omitempty" yaml:"IsMountain,omitempty"`
}

// ArrayOfEquidChipData was auto-generated from WSDL.
type ArrayOfEquidChipData struct {
	EquidChipData []*EquidChipData `xml:"EquidChipData,omitempty" json:"EquidChipData,omitempty" yaml:"EquidChipData,omitempty"`
}

// ArrayOfEquidDescendantData was auto-generated from WSDL.
type ArrayOfEquidDescendantData struct {
	EquidDescendantData []*EquidDescendantData `xml:"EquidDescendantData,omitempty" json:"EquidDescendantData,omitempty" yaml:"EquidDescendantData,omitempty"`
}

// ArrayOfEquidHusbandryHistoryData was auto-generated from WSDL.
type ArrayOfEquidHusbandryHistoryData struct {
	EquidHusbandryHistoryData []*EquidHusbandryHistoryData `xml:"EquidHusbandryHistoryData,omitempty" json:"EquidHusbandryHistoryData,omitempty" yaml:"EquidHusbandryHistoryData,omitempty"`
}

// ArrayOfEquidItem was auto-generated from WSDL.
type ArrayOfEquidItem struct {
	EquidItem []*EquidItem `xml:"EquidItem,omitempty" json:"EquidItem,omitempty" yaml:"EquidItem,omitempty"`
}

// ArrayOfEquidItemMO was auto-generated from WSDL.
type ArrayOfEquidItemMO struct {
	EquidItemMO []*EquidItemMO `xml:"EquidItemMO,omitempty" json:"EquidItemMO,omitempty" yaml:"EquidItemMO,omitempty"`
}

// ArrayOfEquidMembershipData was auto-generated from WSDL.
type ArrayOfEquidMembershipData struct {
	EquidMembershipData []*EquidMembershipData `xml:"EquidMembershipData,omitempty" json:"EquidMembershipData,omitempty" yaml:"EquidMembershipData,omitempty"`
}

// ArrayOfEquidNotificationItem was auto-generated from WSDL.
type ArrayOfEquidNotificationItem struct {
	EquidNotificationItem []*EquidNotificationItem `xml:"EquidNotificationItem,omitempty" json:"EquidNotificationItem,omitempty" yaml:"EquidNotificationItem,omitempty"`
}

// ArrayOfEquidOwnershipHistoryData was auto-generated from WSDL.
type ArrayOfEquidOwnershipHistoryData struct {
	EquidOwnershipHistoryData []*EquidOwnershipHistoryData `xml:"EquidOwnershipHistoryData,omitempty" json:"EquidOwnershipHistoryData,omitempty" yaml:"EquidOwnershipHistoryData,omitempty"`
}

// ArrayOfEquidPassData was auto-generated from WSDL.
type ArrayOfEquidPassData struct {
	EquidPassData []*EquidPassData `xml:"EquidPassData,omitempty" json:"EquidPassData,omitempty" yaml:"EquidPassData,omitempty"`
}

// ArrayOfEquidSearchResultData was auto-generated from WSDL.
type ArrayOfEquidSearchResultData struct {
	EquidSearchResultData []*EquidSearchResultData `xml:"EquidSearchResultData,omitempty" json:"EquidSearchResultData,omitempty" yaml:"EquidSearchResultData,omitempty"`
}

// ArrayOfEquidSlaughterListData was auto-generated from WSDL.
type ArrayOfEquidSlaughterListData struct {
	EquidSlaughterListData []*EquidSlaughterListData `xml:"EquidSlaughterListData,omitempty" json:"EquidSlaughterListData,omitempty" yaml:"EquidSlaughterListData,omitempty"`
}

// ArrayOfEquidUelnData was auto-generated from WSDL.
type ArrayOfEquidUelnData struct {
	EquidUelnData []*EquidUelnData `xml:"EquidUelnData,omitempty" json:"EquidUelnData,omitempty" yaml:"EquidUelnData,omitempty"`
}

// ArrayOfEquidUelnItem was auto-generated from WSDL.
type ArrayOfEquidUelnItem struct {
	EquidUelnItem []*EquidUelnItem `xml:"EquidUelnItem,omitempty" json:"EquidUelnItem,omitempty" yaml:"EquidUelnItem,omitempty"`
}

// ArrayOfGetCattlesPerLeavingDateData was auto-generated from
// WSDL.
type ArrayOfGetCattlesPerLeavingDateData struct {
	GetCattlesPerLeavingDateData []*GetCattlesPerLeavingDateData `xml:"GetCattlesPerLeavingDateData,omitempty" json:"GetCattlesPerLeavingDateData,omitempty" yaml:"GetCattlesPerLeavingDateData,omitempty"`
}

// ArrayOfReplacementEarTagOrderDataItem was auto-generated from
// WSDL.
type ArrayOfReplacementEarTagOrderDataItem struct {
	ReplacementEarTagOrderDataItem []*ReplacementEarTagOrderDataItem `xml:"ReplacementEarTagOrderDataItem,omitempty" json:"ReplacementEarTagOrderDataItem,omitempty" yaml:"ReplacementEarTagOrderDataItem,omitempty"`
}

// ArrayOfSmallAnimalSlaughterListData was auto-generated from
// WSDL.
type ArrayOfSmallAnimalSlaughterListData struct {
	SmallAnimalSlaughterListData []*SmallAnimalSlaughterListData `xml:"SmallAnimalSlaughterListData,omitempty" json:"SmallAnimalSlaughterListData,omitempty" yaml:"SmallAnimalSlaughterListData,omitempty"`
}

// ArrayOfstring was auto-generated from WSDL.
type ArrayOfstring struct {
	String []string `xml:"string,omitempty" json:"string,omitempty" yaml:"string,omitempty"`
}

// BaseRequest was auto-generated from WSDL.
type BaseRequest struct {
	ManufacturerKey string `xml:"ManufacturerKey,omitempty" json:"ManufacturerKey,omitempty" yaml:"ManufacturerKey,omitempty"`
	LCID            int    `xml:"LCID,omitempty" json:"LCID,omitempty" yaml:"LCID,omitempty"`
}

// CattleArrivalData was auto-generated from WSDL.
type CattleArrivalData struct {
	EarTagNumber    string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	EventDate       DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	TVDNumberOrigin int      `xml:"TVDNumberOrigin,omitempty" json:"TVDNumberOrigin,omitempty" yaml:"TVDNumberOrigin,omitempty"`
}

// CattleArrivalDataArray was auto-generated from WSDL.
type CattleArrivalDataArray struct {
	CattleArrivalDataItem []*CattleArrivalData `xml:"CattleArrivalDataItem,omitempty" json:"CattleArrivalDataItem,omitempty" yaml:"CattleArrivalDataItem,omitempty"`
}

// CattleBirthData was auto-generated from WSDL.
type CattleBirthData struct {
	EarTagNumber              string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	IsMultipleBirth           bool     `xml:"IsMultipleBirth,omitempty" json:"IsMultipleBirth,omitempty" yaml:"IsMultipleBirth,omitempty"`
	RaceColor                 int      `xml:"RaceColor,omitempty" json:"RaceColor,omitempty" yaml:"RaceColor,omitempty"`
	Gender                    int      `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	BirthDate                 DateTime `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	Race                      int      `xml:"Race,omitempty" json:"Race,omitempty" yaml:"Race,omitempty"`
	EarTagNumberFather        string   `xml:"EarTagNumberFather,omitempty" json:"EarTagNumberFather,omitempty" yaml:"EarTagNumberFather,omitempty"`
	EarTagNumberMother        string   `xml:"EarTagNumberMother,omitempty" json:"EarTagNumberMother,omitempty" yaml:"EarTagNumberMother,omitempty"`
	BreedingOrganisation      int      `xml:"BreedingOrganisation,omitempty" json:"BreedingOrganisation,omitempty" yaml:"BreedingOrganisation,omitempty"`
	Name                      string   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	BirthProcess              int      `xml:"BirthProcess,omitempty" json:"BirthProcess,omitempty" yaml:"BirthProcess,omitempty"`
	IsCastrated               bool     `xml:"IsCastrated,omitempty" json:"IsCastrated,omitempty" yaml:"IsCastrated,omitempty"`
	InseminationDate          DateTime `xml:"InseminationDate,omitempty" json:"InseminationDate,omitempty" yaml:"InseminationDate,omitempty"`
	EarTagNumberGeneticMother string   `xml:"EarTagNumberGeneticMother,omitempty" json:"EarTagNumberGeneticMother,omitempty" yaml:"EarTagNumberGeneticMother,omitempty"`
	BirthWeight               int      `xml:"BirthWeight,omitempty" json:"BirthWeight,omitempty" yaml:"BirthWeight,omitempty"`
	IsPassportDesired         bool     `xml:"IsPassportDesired,omitempty" json:"IsPassportDesired,omitempty" yaml:"IsPassportDesired,omitempty"`
}

// CattleDaystayData was auto-generated from WSDL.
type CattleDaystayData struct {
	EarTagNumber    string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	EventDate       DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	TVDNumberOrigin int      `xml:"TVDNumberOrigin,omitempty" json:"TVDNumberOrigin,omitempty" yaml:"TVDNumberOrigin,omitempty"`
}

// CattleDaystayDataArray was auto-generated from WSDL.
type CattleDaystayDataArray struct {
	CattleDaystayDataItem []*CattleDaystayData `xml:"CattleDaystayDataItem,omitempty" json:"CattleDaystayDataItem,omitempty" yaml:"CattleDaystayDataItem,omitempty"`
}

// CattleDeathBirthData was auto-generated from WSDL.
type CattleDeathBirthData struct {
	DeathBirthTime            int      `xml:"DeathBirthTime,omitempty" json:"DeathBirthTime,omitempty" yaml:"DeathBirthTime,omitempty"`
	IsMultipleBirth           bool     `xml:"IsMultipleBirth,omitempty" json:"IsMultipleBirth,omitempty" yaml:"IsMultipleBirth,omitempty"`
	RaceColor                 int      `xml:"RaceColor,omitempty" json:"RaceColor,omitempty" yaml:"RaceColor,omitempty"`
	Gender                    int      `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	BirthDate                 DateTime `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	Race                      int      `xml:"Race,omitempty" json:"Race,omitempty" yaml:"Race,omitempty"`
	EarTagNumberFather        string   `xml:"EarTagNumberFather,omitempty" json:"EarTagNumberFather,omitempty" yaml:"EarTagNumberFather,omitempty"`
	EarTagNumberMother        string   `xml:"EarTagNumberMother,omitempty" json:"EarTagNumberMother,omitempty" yaml:"EarTagNumberMother,omitempty"`
	BreedingOrganisation      int      `xml:"BreedingOrganisation,omitempty" json:"BreedingOrganisation,omitempty" yaml:"BreedingOrganisation,omitempty"`
	Name                      string   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	BirthProcess              int      `xml:"BirthProcess,omitempty" json:"BirthProcess,omitempty" yaml:"BirthProcess,omitempty"`
	IsCastrated               bool     `xml:"IsCastrated,omitempty" json:"IsCastrated,omitempty" yaml:"IsCastrated,omitempty"`
	InseminationDate          DateTime `xml:"InseminationDate,omitempty" json:"InseminationDate,omitempty" yaml:"InseminationDate,omitempty"`
	EarTagNumberGeneticMother string   `xml:"EarTagNumberGeneticMother,omitempty" json:"EarTagNumberGeneticMother,omitempty" yaml:"EarTagNumberGeneticMother,omitempty"`
	BirthWeight               int      `xml:"BirthWeight,omitempty" json:"BirthWeight,omitempty" yaml:"BirthWeight,omitempty"`
	IsPassportDesired         bool     `xml:"IsPassportDesired,omitempty" json:"IsPassportDesired,omitempty" yaml:"IsPassportDesired,omitempty"`
}

// CattleDeathBirthNotificationResult was auto-generated from WSDL.
type CattleDeathBirthNotificationResult struct {
	NotificationID int               `xml:"NotificationID,omitempty" json:"NotificationID,omitempty" yaml:"NotificationID,omitempty"`
	Result         *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

// CattleDetailData was auto-generated from WSDL.
type CattleDetailData struct {
	BirthNotificationData *CattleBirthData `xml:"BirthNotificationData,omitempty" json:"BirthNotificationData,omitempty" yaml:"BirthNotificationData,omitempty"`
	LongName              string           `xml:"LongName,omitempty" json:"LongName,omitempty" yaml:"LongName,omitempty"`
	RaceFather            int              `xml:"RaceFather,omitempty" json:"RaceFather,omitempty" yaml:"RaceFather,omitempty"`
	RaceColorFather       int              `xml:"RaceColorFather,omitempty" json:"RaceColorFather,omitempty" yaml:"RaceColorFather,omitempty"`
	NameFather            string           `xml:"NameFather,omitempty" json:"NameFather,omitempty" yaml:"NameFather,omitempty"`
	RaceMother            int              `xml:"RaceMother,omitempty" json:"RaceMother,omitempty" yaml:"RaceMother,omitempty"`
	RaceColorMother       int              `xml:"RaceColorMother,omitempty" json:"RaceColorMother,omitempty" yaml:"RaceColorMother,omitempty"`
	NameMother            string           `xml:"NameMother,omitempty" json:"NameMother,omitempty" yaml:"NameMother,omitempty"`
	Deformations          *IntArray        `xml:"Deformations,omitempty" json:"Deformations,omitempty" yaml:"Deformations,omitempty"`
	DeathDate             DateTime         `xml:"DeathDate,omitempty" json:"DeathDate,omitempty" yaml:"DeathDate,omitempty"`
	BvdState              string           `xml:"BvdState,omitempty" json:"BvdState,omitempty" yaml:"BvdState,omitempty"`
	BtState               string           `xml:"BtState,omitempty" json:"BtState,omitempty" yaml:"BtState,omitempty"`
	CattleTypeOfUse       int              `xml:"CattleTypeOfUse,omitempty" json:"CattleTypeOfUse,omitempty" yaml:"CattleTypeOfUse,omitempty"`
	TypeOfUseDate         DateTime         `xml:"TypeOfUseDate,omitempty" json:"TypeOfUseDate,omitempty" yaml:"TypeOfUseDate,omitempty"`
	FirstCalvingDate      DateTime         `xml:"FirstCalvingDate,omitempty" json:"FirstCalvingDate,omitempty" yaml:"FirstCalvingDate,omitempty"`
	CurrentHusbandry      int              `xml:"CurrentHusbandry,omitempty" json:"CurrentHusbandry,omitempty" yaml:"CurrentHusbandry,omitempty"`
	AllYearHusbandry      int              `xml:"AllYearHusbandry,omitempty" json:"AllYearHusbandry,omitempty" yaml:"AllYearHusbandry,omitempty"`
	PendulumHusbandry     int              `xml:"PendulumHusbandry,omitempty" json:"PendulumHusbandry,omitempty" yaml:"PendulumHusbandry,omitempty"`
	Beefiness             string           `xml:"Beefiness,omitempty" json:"Beefiness,omitempty" yaml:"Beefiness,omitempty"`
	SlaughterCategory     string           `xml:"SlaughterCategory,omitempty" json:"SlaughterCategory,omitempty" yaml:"SlaughterCategory,omitempty"`
	FatTissue             string           `xml:"FatTissue,omitempty" json:"FatTissue,omitempty" yaml:"FatTissue,omitempty"`
}

// CattleDetailDataV2 was auto-generated from WSDL.
type CattleDetailDataV2 struct {
	BirthNotificationData *CattleBirthData `xml:"BirthNotificationData,omitempty" json:"BirthNotificationData,omitempty" yaml:"BirthNotificationData,omitempty"`
	LongName              string           `xml:"LongName,omitempty" json:"LongName,omitempty" yaml:"LongName,omitempty"`
	RaceFather            int              `xml:"RaceFather,omitempty" json:"RaceFather,omitempty" yaml:"RaceFather,omitempty"`
	RaceColorFather       int              `xml:"RaceColorFather,omitempty" json:"RaceColorFather,omitempty" yaml:"RaceColorFather,omitempty"`
	NameFather            string           `xml:"NameFather,omitempty" json:"NameFather,omitempty" yaml:"NameFather,omitempty"`
	RaceMother            int              `xml:"RaceMother,omitempty" json:"RaceMother,omitempty" yaml:"RaceMother,omitempty"`
	RaceColorMother       int              `xml:"RaceColorMother,omitempty" json:"RaceColorMother,omitempty" yaml:"RaceColorMother,omitempty"`
	NameMother            string           `xml:"NameMother,omitempty" json:"NameMother,omitempty" yaml:"NameMother,omitempty"`
	Deformations          *IntArray        `xml:"Deformations,omitempty" json:"Deformations,omitempty" yaml:"Deformations,omitempty"`
	DeathDate             DateTime         `xml:"DeathDate,omitempty" json:"DeathDate,omitempty" yaml:"DeathDate,omitempty"`
	BvdState              string           `xml:"BvdState,omitempty" json:"BvdState,omitempty" yaml:"BvdState,omitempty"`
	BtState               string           `xml:"BtState,omitempty" json:"BtState,omitempty" yaml:"BtState,omitempty"`
	CattleTypeOfUse       int              `xml:"CattleTypeOfUse,omitempty" json:"CattleTypeOfUse,omitempty" yaml:"CattleTypeOfUse,omitempty"`
	TypeOfUseDate         DateTime         `xml:"TypeOfUseDate,omitempty" json:"TypeOfUseDate,omitempty" yaml:"TypeOfUseDate,omitempty"`
	FirstCalvingDate      DateTime         `xml:"FirstCalvingDate,omitempty" json:"FirstCalvingDate,omitempty" yaml:"FirstCalvingDate,omitempty"`
	CurrentHusbandry      int              `xml:"CurrentHusbandry,omitempty" json:"CurrentHusbandry,omitempty" yaml:"CurrentHusbandry,omitempty"`
	AllYearHusbandry      int              `xml:"AllYearHusbandry,omitempty" json:"AllYearHusbandry,omitempty" yaml:"AllYearHusbandry,omitempty"`
	PendulumHusbandry     int              `xml:"PendulumHusbandry,omitempty" json:"PendulumHusbandry,omitempty" yaml:"PendulumHusbandry,omitempty"`
	Beefiness             string           `xml:"Beefiness,omitempty" json:"Beefiness,omitempty" yaml:"Beefiness,omitempty"`
	SlaughterCategory     string           `xml:"SlaughterCategory,omitempty" json:"SlaughterCategory,omitempty" yaml:"SlaughterCategory,omitempty"`
	FatTissue             string           `xml:"FatTissue,omitempty" json:"FatTissue,omitempty" yaml:"FatTissue,omitempty"`
	LValue                float64          `xml:"LValue,omitempty" json:"LValue,omitempty" yaml:"LValue,omitempty"`
}

// CattleDetailResult was auto-generated from WSDL.
type CattleDetailResult struct {
	Result       *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	CattleDetail *CattleDetailData `xml:"CattleDetail,omitempty" json:"CattleDetail,omitempty" yaml:"CattleDetail,omitempty"`
}

// CattleDetailResultV2 was auto-generated from WSDL.
type CattleDetailResultV2 struct {
	Result       *ProcessingResult   `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	CattleDetail *CattleDetailDataV2 `xml:"CattleDetail,omitempty" json:"CattleDetail,omitempty" yaml:"CattleDetail,omitempty"`
}

// CattleEarTagData was auto-generated from WSDL.
type CattleEarTagData struct {
	EarTagNumber string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	OrderDate    DateTime `xml:"OrderDate,omitempty" json:"OrderDate,omitempty" yaml:"OrderDate,omitempty"`
	EarTagState  int      `xml:"EarTagState,omitempty" json:"EarTagState,omitempty" yaml:"EarTagState,omitempty"`
}

// CattleEarTagDataArray was auto-generated from WSDL.
type CattleEarTagDataArray struct {
	CattleEarTagDataItem []*CattleEarTagData `xml:"CattleEarTagDataItem,omitempty" json:"CattleEarTagDataItem,omitempty" yaml:"CattleEarTagDataItem,omitempty"`
}

// CattleEarTagsResult was auto-generated from WSDL.
type CattleEarTagsResult struct {
	Result        *ProcessingResult      `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	CattleEarTags *CattleEarTagDataArray `xml:"CattleEarTags,omitempty" json:"CattleEarTags,omitempty" yaml:"CattleEarTags,omitempty"`
}

// CattleHistoryResult was auto-generated from WSDL.
type CattleHistoryResult struct {
	Result             *ProcessingResult    `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	EarTagNumber       string               `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	RemainingQuota     int                  `xml:"RemainingQuota,omitempty" json:"RemainingQuota,omitempty" yaml:"RemainingQuota,omitempty"`
	CattleHistoryState int                  `xml:"CattleHistoryState,omitempty" json:"CattleHistoryState,omitempty" yaml:"CattleHistoryState,omitempty"`
	CattleStays        *CattleStayDataArray `xml:"CattleStays,omitempty" json:"CattleStays,omitempty" yaml:"CattleStays,omitempty"`
}

// CattleHistoryResultV2 was auto-generated from WSDL.
type CattleHistoryResultV2 struct {
	Result             *ProcessingResult      `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	EarTagNumber       string                 `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	RemainingQuota     int                    `xml:"RemainingQuota,omitempty" json:"RemainingQuota,omitempty" yaml:"RemainingQuota,omitempty"`
	CattleHistoryState int                    `xml:"CattleHistoryState,omitempty" json:"CattleHistoryState,omitempty" yaml:"CattleHistoryState,omitempty"`
	CattleStays        *CattleStayDataArrayV2 `xml:"CattleStays,omitempty" json:"CattleStays,omitempty" yaml:"CattleStays,omitempty"`
}

// CattleLeavingData was auto-generated from WSDL.
type CattleLeavingData struct {
	EarTagNumber                string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	EventDate                   DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	LeavingReason               int      `xml:"LeavingReason,omitempty" json:"LeavingReason,omitempty" yaml:"LeavingReason,omitempty"`
	MainLeavingReasonBreeding   int      `xml:"MainLeavingReasonBreeding,omitempty" json:"MainLeavingReasonBreeding,omitempty" yaml:"MainLeavingReasonBreeding,omitempty"`
	SecondLeavingReasonBreeding int      `xml:"SecondLeavingReasonBreeding,omitempty" json:"SecondLeavingReasonBreeding,omitempty" yaml:"SecondLeavingReasonBreeding,omitempty"`
}

// CattleLeavingDataArray was auto-generated from WSDL.
type CattleLeavingDataArray struct {
	CattleLeavingDataItem []*CattleLeavingData `xml:"CattleLeavingDataItem,omitempty" json:"CattleLeavingDataItem,omitempty" yaml:"CattleLeavingDataItem,omitempty"`
}

// CattleLivestockData was auto-generated from WSDL.
type CattleLivestockData struct {
	EarTagNumber       string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	Name               string   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Gender             int      `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	BirthDate          DateTime `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	Race               int      `xml:"Race,omitempty" json:"Race,omitempty" yaml:"Race,omitempty"`
	RaceColor          int      `xml:"RaceColor,omitempty" json:"RaceColor,omitempty" yaml:"RaceColor,omitempty"`
	ArrivalDate        DateTime `xml:"ArrivalDate,omitempty" json:"ArrivalDate,omitempty" yaml:"ArrivalDate,omitempty"`
	LeavingDate        DateTime `xml:"LeavingDate,omitempty" json:"LeavingDate,omitempty" yaml:"LeavingDate,omitempty"`
	PendulumHusbandry  int      `xml:"PendulumHusbandry,omitempty" json:"PendulumHusbandry,omitempty" yaml:"PendulumHusbandry,omitempty"`
	BvdState           string   `xml:"BvdState,omitempty" json:"BvdState,omitempty" yaml:"BvdState,omitempty"`
	BtState            string   `xml:"BtState,omitempty" json:"BtState,omitempty" yaml:"BtState,omitempty"`
	CattleTypeOfUse    int      `xml:"CattleTypeOfUse,omitempty" json:"CattleTypeOfUse,omitempty" yaml:"CattleTypeOfUse,omitempty"`
	TypeOfUseDate      DateTime `xml:"TypeOfUseDate,omitempty" json:"TypeOfUseDate,omitempty" yaml:"TypeOfUseDate,omitempty"`
	FirstCalvingDate   DateTime `xml:"FirstCalvingDate,omitempty" json:"FirstCalvingDate,omitempty" yaml:"FirstCalvingDate,omitempty"`
	CattleHistoryState int      `xml:"CattleHistoryState,omitempty" json:"CattleHistoryState,omitempty" yaml:"CattleHistoryState,omitempty"`
}

// CattleLivestockDataArray was auto-generated from WSDL.
type CattleLivestockDataArray struct {
	CattleLivestockDataItem []*CattleLivestockData `xml:"CattleLivestockDataItem,omitempty" json:"CattleLivestockDataItem,omitempty" yaml:"CattleLivestockDataItem,omitempty"`
}

// CattleLivestockDataArrayV2 was auto-generated from WSDL.
type CattleLivestockDataArrayV2 struct {
	CattleLivestockDataItem []*CattleLivestockDataV2 `xml:"CattleLivestockDataItem,omitempty" json:"CattleLivestockDataItem,omitempty" yaml:"CattleLivestockDataItem,omitempty"`
}

// CattleLivestockDataV2 was auto-generated from WSDL.
type CattleLivestockDataV2 struct {
	EarTagNumber       string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	Name               string   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Gender             int      `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	BirthDate          DateTime `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	Race               int      `xml:"Race,omitempty" json:"Race,omitempty" yaml:"Race,omitempty"`
	RaceColor          int      `xml:"RaceColor,omitempty" json:"RaceColor,omitempty" yaml:"RaceColor,omitempty"`
	ArrivalDate        DateTime `xml:"ArrivalDate,omitempty" json:"ArrivalDate,omitempty" yaml:"ArrivalDate,omitempty"`
	LeavingDate        DateTime `xml:"LeavingDate,omitempty" json:"LeavingDate,omitempty" yaml:"LeavingDate,omitempty"`
	PendulumHusbandry  int      `xml:"PendulumHusbandry,omitempty" json:"PendulumHusbandry,omitempty" yaml:"PendulumHusbandry,omitempty"`
	BvdState           string   `xml:"BvdState,omitempty" json:"BvdState,omitempty" yaml:"BvdState,omitempty"`
	BtState            string   `xml:"BtState,omitempty" json:"BtState,omitempty" yaml:"BtState,omitempty"`
	CattleTypeOfUse    int      `xml:"CattleTypeOfUse,omitempty" json:"CattleTypeOfUse,omitempty" yaml:"CattleTypeOfUse,omitempty"`
	TypeOfUseDate      DateTime `xml:"TypeOfUseDate,omitempty" json:"TypeOfUseDate,omitempty" yaml:"TypeOfUseDate,omitempty"`
	FirstCalvingDate   DateTime `xml:"FirstCalvingDate,omitempty" json:"FirstCalvingDate,omitempty" yaml:"FirstCalvingDate,omitempty"`
	CattleHistoryState int      `xml:"CattleHistoryState,omitempty" json:"CattleHistoryState,omitempty" yaml:"CattleHistoryState,omitempty"`
	LastCalvingDate    DateTime `xml:"LastCalvingDate,omitempty" json:"LastCalvingDate,omitempty" yaml:"LastCalvingDate,omitempty"`
}

// CattleLivestockResult was auto-generated from WSDL.
type CattleLivestockResult struct {
	Result        *ProcessingResult         `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	TVDNumber     int                       `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	Resultdetails *CattleLivestockDataArray `xml:"Resultdetails,omitempty" json:"Resultdetails,omitempty" yaml:"Resultdetails,omitempty"`
}

// CattleLivestockResultV2 was auto-generated from WSDL.
type CattleLivestockResultV2 struct {
	Result        *ProcessingResult           `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	TVDNumber     int                         `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	Resultdetails *CattleLivestockDataArrayV2 `xml:"Resultdetails,omitempty" json:"Resultdetails,omitempty" yaml:"Resultdetails,omitempty"`
}

// CattleMovementData was auto-generated from WSDL.
type CattleMovementData struct {
	NotificationID    int      `xml:"NotificationID,omitempty" json:"NotificationID,omitempty" yaml:"NotificationID,omitempty"`
	EventDate         DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	NotificationDate  DateTime `xml:"NotificationDate,omitempty" json:"NotificationDate,omitempty" yaml:"NotificationDate,omitempty"`
	NotificationType  int      `xml:"NotificationType,omitempty" json:"NotificationType,omitempty" yaml:"NotificationType,omitempty"`
	TVDNumberNotifier int      `xml:"TVDNumberNotifier,omitempty" json:"TVDNumberNotifier,omitempty" yaml:"TVDNumberNotifier,omitempty"`
	TVDNumberOrigin   int      `xml:"TVDNumberOrigin,omitempty" json:"TVDNumberOrigin,omitempty" yaml:"TVDNumberOrigin,omitempty"`
}

// CattleMovementDataArray was auto-generated from WSDL.
type CattleMovementDataArray struct {
	CattleMovementDataItem []*CattleMovementData `xml:"CattleMovementDataItem,omitempty" json:"CattleMovementDataItem,omitempty" yaml:"CattleMovementDataItem,omitempty"`
}

// CattleMovementsResult was auto-generated from WSDL.
type CattleMovementsResult struct {
	Result             *ProcessingResult        `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	CattleHistoryState int                      `xml:"CattleHistoryState,omitempty" json:"CattleHistoryState,omitempty" yaml:"CattleHistoryState,omitempty"`
	EarTagNumber       string                   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	CattleMovements    *CattleMovementDataArray `xml:"CattleMovements,omitempty" json:"CattleMovements,omitempty" yaml:"CattleMovements,omitempty"`
}

// CattleNotificationRequest was auto-generated from WSDL.
type CattleNotificationRequest struct {
	ManufacturerKey string   `xml:"ManufacturerKey,omitempty" json:"ManufacturerKey,omitempty" yaml:"ManufacturerKey,omitempty"`
	LCID            int      `xml:"LCID,omitempty" json:"LCID,omitempty" yaml:"LCID,omitempty"`
	TVDNumber       int      `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	EarTagNumber    string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	EventDate       DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
}

// CattleNotificationResult was auto-generated from WSDL.
type CattleNotificationResult struct {
	EarTagNumber   string            `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	NotificationID int               `xml:"NotificationID,omitempty" json:"NotificationID,omitempty" yaml:"NotificationID,omitempty"`
	Result         *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

// CattleNotificationResultArray was auto-generated from WSDL.
type CattleNotificationResultArray struct {
	CattleNotificationResultItem []*CattleNotificationResult `xml:"CattleNotificationResultItem,omitempty" json:"CattleNotificationResultItem,omitempty" yaml:"CattleNotificationResultItem,omitempty"`
}

// CattleNotificationWithCountryRequest was auto-generated from
// WSDL.
type CattleNotificationWithCountryRequest struct {
	ManufacturerKey string      `xml:"ManufacturerKey,omitempty" json:"ManufacturerKey,omitempty" yaml:"ManufacturerKey,omitempty"`
	LCID            int         `xml:"LCID,omitempty" json:"LCID,omitempty" yaml:"LCID,omitempty"`
	TVDNumber       int         `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	EarTagNumber    string      `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	EventDate       DateTime    `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	Country         EnumCountry `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
}

// CattleOffspringData was auto-generated from WSDL.
type CattleOffspringData struct {
	EarTagNumber string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	BirthDate    DateTime `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	Gender       int      `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	Name         string   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Race         int      `xml:"Race,omitempty" json:"Race,omitempty" yaml:"Race,omitempty"`
}

// CattleOffspringDataArray was auto-generated from WSDL.
type CattleOffspringDataArray struct {
	CattleOffspringDataItem []*CattleOffspringData `xml:"CattleOffspringDataItem,omitempty" json:"CattleOffspringDataItem,omitempty" yaml:"CattleOffspringDataItem,omitempty"`
}

// CattleOffspringResult was auto-generated from WSDL.
type CattleOffspringResult struct {
	Result             *ProcessingResult         `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	EarTagNumberMother string                    `xml:"EarTagNumberMother,omitempty" json:"EarTagNumberMother,omitempty" yaml:"EarTagNumberMother,omitempty"`
	CattleOffsprings   *CattleOffspringDataArray `xml:"CattleOffsprings,omitempty" json:"CattleOffsprings,omitempty" yaml:"CattleOffsprings,omitempty"`
}

// CattlePassport was auto-generated from WSDL.
type CattlePassport struct {
	EarTagNumber      string            `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	PassportNumber    int               `xml:"PassportNumber,omitempty" json:"PassportNumber,omitempty" yaml:"PassportNumber,omitempty"`
	PassportIssueDate DateTime          `xml:"PassportIssueDate,omitempty" json:"PassportIssueDate,omitempty" yaml:"PassportIssueDate,omitempty"`
	Result            *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

// CattlePassportArray was auto-generated from WSDL.
type CattlePassportArray struct {
	CattlePassportItem []*CattlePassport `xml:"CattlePassportItem,omitempty" json:"CattlePassportItem,omitempty" yaml:"CattlePassportItem,omitempty"`
}

// CattlePassportResult was auto-generated from WSDL.
type CattlePassportResult struct {
	Result          *ProcessingResult    `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	CattlePassports *CattlePassportArray `xml:"CattlePassports,omitempty" json:"CattlePassports,omitempty" yaml:"CattlePassports,omitempty"`
}

// CattleSlaughterData was auto-generated from WSDL.
type CattleSlaughterData struct {
	EarTagNumber                string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	EventDate                   DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	TVDNumberOrigin             int      `xml:"TVDNumberOrigin,omitempty" json:"TVDNumberOrigin,omitempty" yaml:"TVDNumberOrigin,omitempty"`
	TVDNumberSlaughterInitiator int      `xml:"TVDNumberSlaughterInitiator,omitempty" json:"TVDNumberSlaughterInitiator,omitempty" yaml:"TVDNumberSlaughterInitiator,omitempty"`
}

// CattleSlaughterDataArray was auto-generated from WSDL.
type CattleSlaughterDataArray struct {
	CattleSlaughterDataItem []*CattleSlaughterData `xml:"CattleSlaughterDataItem,omitempty" json:"CattleSlaughterDataItem,omitempty" yaml:"CattleSlaughterDataItem,omitempty"`
}

// CattleStateData was auto-generated from WSDL.
type CattleStateData struct {
	EarTagNumber       string `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	Name               string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	BvdState           string `xml:"BvdState,omitempty" json:"BvdState,omitempty" yaml:"BvdState,omitempty"`
	CattleHistoryState int    `xml:"CattleHistoryState,omitempty" json:"CattleHistoryState,omitempty" yaml:"CattleHistoryState,omitempty"`
}

// CattleStateDataV2 was auto-generated from WSDL.
type CattleStateDataV2 struct {
	EarTagNumber       string   `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	Name               string   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	BvdState           string   `xml:"BvdState,omitempty" json:"BvdState,omitempty" yaml:"BvdState,omitempty"`
	CattleHistoryState int      `xml:"CattleHistoryState,omitempty" json:"CattleHistoryState,omitempty" yaml:"CattleHistoryState,omitempty"`
	CattleBirthDate    DateTime `xml:"CattleBirthDate,omitempty" json:"CattleBirthDate,omitempty" yaml:"CattleBirthDate,omitempty"`
	CattleAgeInDays    int      `xml:"CattleAgeInDays,omitempty" json:"CattleAgeInDays,omitempty" yaml:"CattleAgeInDays,omitempty"`
}

// CattleStateExternalResult was auto-generated from WSDL.
type CattleStateExternalResult struct {
	Result      *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	CattleState *CattleStateData  `xml:"CattleState,omitempty" json:"CattleState,omitempty" yaml:"CattleState,omitempty"`
}

// CattleStateExternalResultV2 was auto-generated from WSDL.
type CattleStateExternalResultV2 struct {
	Result      *ProcessingResult  `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	CattleState *CattleStateDataV2 `xml:"CattleState,omitempty" json:"CattleState,omitempty" yaml:"CattleState,omitempty"`
}

// CattleStayData was auto-generated from WSDL.
type CattleStayData struct {
	ArrivalDate             DateTime `xml:"ArrivalDate,omitempty" json:"ArrivalDate,omitempty" yaml:"ArrivalDate,omitempty"`
	ArrivalNotificationDate DateTime `xml:"ArrivalNotificationDate,omitempty" json:"ArrivalNotificationDate,omitempty" yaml:"ArrivalNotificationDate,omitempty"`
	ArrivalNotificationType int      `xml:"ArrivalNotificationType,omitempty" json:"ArrivalNotificationType,omitempty" yaml:"ArrivalNotificationType,omitempty"`
	CountryOrigin           string   `xml:"CountryOrigin,omitempty" json:"CountryOrigin,omitempty" yaml:"CountryOrigin,omitempty"`
	TVDNumberOrigin         int      `xml:"TVDNumberOrigin,omitempty" json:"TVDNumberOrigin,omitempty" yaml:"TVDNumberOrigin,omitempty"`
	TVDNumberStay           int      `xml:"TVDNumberStay,omitempty" json:"TVDNumberStay,omitempty" yaml:"TVDNumberStay,omitempty"`
	StayAddress             string   `xml:"StayAddress,omitempty" json:"StayAddress,omitempty" yaml:"StayAddress,omitempty"`
	LeavingDate             DateTime `xml:"LeavingDate,omitempty" json:"LeavingDate,omitempty" yaml:"LeavingDate,omitempty"`
	LeavingNotificationDate DateTime `xml:"LeavingNotificationDate,omitempty" json:"LeavingNotificationDate,omitempty" yaml:"LeavingNotificationDate,omitempty"`
	LeavingNotificationType int      `xml:"LeavingNotificationType,omitempty" json:"LeavingNotificationType,omitempty" yaml:"LeavingNotificationType,omitempty"`
	CattleStayState         int      `xml:"CattleStayState,omitempty" json:"CattleStayState,omitempty" yaml:"CattleStayState,omitempty"`
}

// CattleStayDataArray was auto-generated from WSDL.
type CattleStayDataArray struct {
	CattleStayDataItem []*CattleStayData `xml:"CattleStayDataItem,omitempty" json:"CattleStayDataItem,omitempty" yaml:"CattleStayDataItem,omitempty"`
}

// CattleStayDataArrayV2 was auto-generated from WSDL.
type CattleStayDataArrayV2 struct {
	CattleStayDataItemV2 []*CattleStayDataV2 `xml:"CattleStayDataItemV2,omitempty" json:"CattleStayDataItemV2,omitempty" yaml:"CattleStayDataItemV2,omitempty"`
}

// CattleStayDataV2 was auto-generated from WSDL.
type CattleStayDataV2 struct {
	ArrivalDate             DateTime `xml:"ArrivalDate,omitempty" json:"ArrivalDate,omitempty" yaml:"ArrivalDate,omitempty"`
	ArrivalNotificationDate DateTime `xml:"ArrivalNotificationDate,omitempty" json:"ArrivalNotificationDate,omitempty" yaml:"ArrivalNotificationDate,omitempty"`
	ArrivalNotificationType int      `xml:"ArrivalNotificationType,omitempty" json:"ArrivalNotificationType,omitempty" yaml:"ArrivalNotificationType,omitempty"`
	CountryOrigin           string   `xml:"CountryOrigin,omitempty" json:"CountryOrigin,omitempty" yaml:"CountryOrigin,omitempty"`
	TVDNumberOrigin         int      `xml:"TVDNumberOrigin,omitempty" json:"TVDNumberOrigin,omitempty" yaml:"TVDNumberOrigin,omitempty"`
	TVDNumberStay           int      `xml:"TVDNumberStay,omitempty" json:"TVDNumberStay,omitempty" yaml:"TVDNumberStay,omitempty"`
	AnhName                 string   `xml:"AnhName,omitempty" json:"AnhName,omitempty" yaml:"AnhName,omitempty"`
	Street                  string   `xml:"Street,omitempty" json:"Street,omitempty" yaml:"Street,omitempty"`
	ZipCode                 int      `xml:"ZipCode,omitempty" json:"ZipCode,omitempty" yaml:"ZipCode,omitempty"`
	City                    string   `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	LeavingDate             DateTime `xml:"LeavingDate,omitempty" json:"LeavingDate,omitempty" yaml:"LeavingDate,omitempty"`
	LeavingNotificationDate DateTime `xml:"LeavingNotificationDate,omitempty" json:"LeavingNotificationDate,omitempty" yaml:"LeavingNotificationDate,omitempty"`
	LeavingNotificationType int      `xml:"LeavingNotificationType,omitempty" json:"LeavingNotificationType,omitempty" yaml:"LeavingNotificationType,omitempty"`
	CattleStayState         int      `xml:"CattleStayState,omitempty" json:"CattleStayState,omitempty" yaml:"CattleStayState,omitempty"`
}

// CheckCattleForDisposalContribution was auto-generated from WSDL.
type CheckCattleForDisposalContribution struct {
	XMLName           xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattleForDisposalContribution" json:"-" yaml:"-"`
	P_ManufacturerKey string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int                `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_SlaughterData   *CattleArrivalData `xml:"p_SlaughterData" json:"p_SlaughterData" yaml:"p_SlaughterData"`
}

// CheckCattleForDisposalContributionResponse was auto-generated
// from WSDL.
type CheckCattleForDisposalContributionResponse struct {
	CheckCattleForDisposalContributionResult *DisposalContributionResult `xml:"CheckCattleForDisposalContributionResult,omitempty" json:"CheckCattleForDisposalContributionResult,omitempty" yaml:"CheckCattleForDisposalContributionResult,omitempty"`
}

// CheckCattlesForPassport was auto-generated from WSDL.
type CheckCattlesForPassport struct {
	XMLName           xml.Name     `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattlesForPassport" json:"-" yaml:"-"`
	P_ManufacturerKey string       `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int          `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_IssueDate       DateTime     `xml:"p_IssueDate" json:"p_IssueDate" yaml:"p_IssueDate"`
	P_EarTagNumbers   *StringArray `xml:"p_EarTagNumbers" json:"p_EarTagNumbers" yaml:"p_EarTagNumbers"`
}

// CheckCattlesForPassportResponse was auto-generated from WSDL.
type CheckCattlesForPassportResponse struct {
	CheckCattlesForPassportResult *CattlePassportResult `xml:"CheckCattlesForPassportResult,omitempty" json:"CheckCattlesForPassportResult,omitempty" yaml:"CheckCattlesForPassportResult,omitempty"`
}

// Code was auto-generated from WSDL.
type Code struct {
	Key     int    `xml:"Key,omitempty" json:"Key,omitempty" yaml:"Key,omitempty"`
	Text_de string `xml:"Text_de,omitempty" json:"Text_de,omitempty" yaml:"Text_de,omitempty"`
	Text_fr string `xml:"Text_fr,omitempty" json:"Text_fr,omitempty" yaml:"Text_fr,omitempty"`
	Text_it string `xml:"Text_it,omitempty" json:"Text_it,omitempty" yaml:"Text_it,omitempty"`
}

// CodeArray was auto-generated from WSDL.
type CodeArray struct {
	CodeItem []*Code `xml:"CodeItem,omitempty" json:"CodeItem,omitempty" yaml:"CodeItem,omitempty"`
}

// CodesResult was auto-generated from WSDL.
type CodesResult struct {
	Result *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	Codes  *CodeArray        `xml:"Codes,omitempty" json:"Codes,omitempty" yaml:"Codes,omitempty"`
}

// DataAccessResult was auto-generated from WSDL.
type DataAccessResult struct {
	Result *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	ID_PRG int               `xml:"ID_PRG,omitempty" json:"ID_PRG,omitempty" yaml:"ID_PRG,omitempty"`
}

// DeleteAnimalHusbandryMembershipResult was auto-generated from
// WSDL.
type DeleteAnimalHusbandryMembershipResult struct {
	Result        *ProcessingResult                 `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	Resultdetails *HusbandryNotificationResultArray `xml:"Resultdetails,omitempty" json:"Resultdetails,omitempty" yaml:"Resultdetails,omitempty"`
}

// DeleteAnimalHusbandryMemberships was auto-generated from WSDL.
type DeleteAnimalHusbandryMemberships struct {
	XMLName                xml.Name  `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteAnimalHusbandryMemberships" json:"-" yaml:"-"`
	P_ManufacturerKey      string    `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                 int       `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_ID_PRG               int       `xml:"p_ID_PRG" json:"p_ID_PRG" yaml:"p_ID_PRG"`
	P_DeleteMembershipData *IntArray `xml:"p_DeleteMembershipData" json:"p_DeleteMembershipData" yaml:"p_DeleteMembershipData"`
}

// DeleteAnimalHusbandryMembershipsResponse was auto-generated
// from WSDL.
type DeleteAnimalHusbandryMembershipsResponse struct {
	DeleteAnimalHusbandryMembershipsResult *DeleteAnimalHusbandryMembershipResult `xml:"DeleteAnimalHusbandryMembershipsResult,omitempty" json:"DeleteAnimalHusbandryMembershipsResult,omitempty" yaml:"DeleteAnimalHusbandryMembershipsResult,omitempty"`
}

// DeleteCattleNotifications was auto-generated from WSDL.
type DeleteCattleNotifications struct {
	XMLName                             xml.Name  `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteCattleNotifications" json:"-" yaml:"-"`
	P_ManufacturerKey                   string    `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                              int       `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber                         int       `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_DeleteCattleNotificationDataArray *IntArray `xml:"p_DeleteCattleNotificationDataArray" json:"p_DeleteCattleNotificationDataArray" yaml:"p_DeleteCattleNotificationDataArray"`
}

// DeleteCattleNotificationsResponse was auto-generated from WSDL.
type DeleteCattleNotificationsResponse struct {
	DeleteCattleNotificationsResult *WriteCattleBatchNotificationResult `xml:"DeleteCattleNotificationsResult,omitempty" json:"DeleteCattleNotificationsResult,omitempty" yaml:"DeleteCattleNotificationsResult,omitempty"`
}

// DeleteEarTagOrder was auto-generated from WSDL.
type DeleteEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteEarTagOrder" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_NotificationID  int      `xml:"p_NotificationID" json:"p_NotificationID" yaml:"p_NotificationID"`
}

// DeleteEarTagOrderResponse was auto-generated from WSDL.
type DeleteEarTagOrderResponse struct {
	DeleteEarTagOrderResult *NotificationResult `xml:"DeleteEarTagOrderResult,omitempty" json:"DeleteEarTagOrderResult,omitempty" yaml:"DeleteEarTagOrderResult,omitempty"`
}

// DeleteLabelEarTagOrder was auto-generated from WSDL.
type DeleteLabelEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteLabelEarTagOrder" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_NotificationID  int      `xml:"p_NotificationID" json:"p_NotificationID" yaml:"p_NotificationID"`
}

// DeleteLabelEarTagOrderResponse was auto-generated from WSDL.
type DeleteLabelEarTagOrderResponse struct {
	DeleteLabelEarTagOrderResult *NotificationResult `xml:"DeleteLabelEarTagOrderResult,omitempty" json:"DeleteLabelEarTagOrderResult,omitempty" yaml:"DeleteLabelEarTagOrderResult,omitempty"`
}

// DisableDataAccess was auto-generated from WSDL.
type DisableDataAccess struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DisableDataAccess" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_ID_PRG          int      `xml:"p_ID_PRG" json:"p_ID_PRG" yaml:"p_ID_PRG"`
}

// DisableDataAccessResponse was auto-generated from WSDL.
type DisableDataAccessResponse struct {
	DisableDataAccessResult *DataAccessResult `xml:"DisableDataAccessResult,omitempty" json:"DisableDataAccessResult,omitempty" yaml:"DisableDataAccessResult,omitempty"`
}

// DisposalContributionResult was auto-generated from WSDL.
type DisposalContributionResult struct {
	DisposalContributionState bool              `xml:"DisposalContributionState,omitempty" json:"DisposalContributionState,omitempty" yaml:"DisposalContributionState,omitempty"`
	Result                    *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

// EarTagOrderData was auto-generated from WSDL.
type EarTagOrderData struct {
	NotificationID   int      `xml:"NotificationID,omitempty" json:"NotificationID,omitempty" yaml:"NotificationID,omitempty"`
	EarTagType       int      `xml:"EarTagType,omitempty" json:"EarTagType,omitempty" yaml:"EarTagType,omitempty"`
	Amount           int      `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
	IsExpress        bool     `xml:"IsExpress,omitempty" json:"IsExpress,omitempty" yaml:"IsExpress,omitempty"`
	OrderStatus      int      `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty" yaml:"OrderStatus,omitempty"`
	OrderStatusDate  DateTime `xml:"OrderStatusDate,omitempty" json:"OrderStatusDate,omitempty" yaml:"OrderStatusDate,omitempty"`
	EarTagNumberFrom string   `xml:"EarTagNumberFrom,omitempty" json:"EarTagNumberFrom,omitempty" yaml:"EarTagNumberFrom,omitempty"`
	EarTagNumberTo   string   `xml:"EarTagNumberTo,omitempty" json:"EarTagNumberTo,omitempty" yaml:"EarTagNumberTo,omitempty"`
	Text1            string   `xml:"Text1,omitempty" json:"Text1,omitempty" yaml:"Text1,omitempty"`
	Text2            string   `xml:"Text2,omitempty" json:"Text2,omitempty" yaml:"Text2,omitempty"`
}

// EarTagOrderDataArray was auto-generated from WSDL.
type EarTagOrderDataArray struct {
	EarTagOrderDataItem []*EarTagOrderData `xml:"EarTagOrderDataItem,omitempty" json:"EarTagOrderDataItem,omitempty" yaml:"EarTagOrderDataItem,omitempty"`
}

// EarTagOrderResult was auto-generated from WSDL.
type EarTagOrderResult struct {
	Result        *ProcessingResult     `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	Resultdetails *EarTagOrderDataArray `xml:"Resultdetails,omitempty" json:"Resultdetails,omitempty" yaml:"Resultdetails,omitempty"`
}

// EartagOrderItem was auto-generated from WSDL.
type EartagOrderItem struct {
	EarTagNumber     string `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	IsExpress        bool   `xml:"IsExpress,omitempty" json:"IsExpress,omitempty" yaml:"IsExpress,omitempty"`
	IsLeftSideOrder  bool   `xml:"IsLeftSideOrder,omitempty" json:"IsLeftSideOrder,omitempty" yaml:"IsLeftSideOrder,omitempty"`
	IsRightSideOrder bool   `xml:"IsRightSideOrder,omitempty" json:"IsRightSideOrder,omitempty" yaml:"IsRightSideOrder,omitempty"`
}

// EartagOrderItemArray was auto-generated from WSDL.
type EartagOrderItemArray struct {
	EartagOrderArrayItem []*EartagOrderItem `xml:"EartagOrderArrayItem,omitempty" json:"EartagOrderArrayItem,omitempty" yaml:"EartagOrderArrayItem,omitempty"`
}

// EnableDataAccess was auto-generated from WSDL.
type EnableDataAccess struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EnableDataAccess" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_ID_PRG          int      `xml:"p_ID_PRG" json:"p_ID_PRG" yaml:"p_ID_PRG"`
}

// EnableDataAccessResponse was auto-generated from WSDL.
type EnableDataAccessResponse struct {
	EnableDataAccessResult *DataAccessResult `xml:"EnableDataAccessResult,omitempty" json:"EnableDataAccessResult,omitempty" yaml:"EnableDataAccessResult,omitempty"`
}

// EquidAcquireOwnershipRequest was auto-generated from WSDL.
type EquidAcquireOwnershipRequest struct {
	OnBehalfOfAgateNumber  string   `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                   string   `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	EventDate              DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	ActualOwnerAgateNumber string   `xml:"ActualOwnerAgateNumber,omitempty" json:"ActualOwnerAgateNumber,omitempty" yaml:"ActualOwnerAgateNumber,omitempty"`
}

// EquidAdditonalInformationData was auto-generated from WSDL.
type EquidAdditonalInformationData struct {
	IsCastrated                       bool                                            `xml:"IsCastrated,omitempty" json:"IsCastrated,omitempty" yaml:"IsCastrated,omitempty"`
	EquidTypeOfUsage                  *TranslatedEnumTypeOfEnumEquidTypeOfUse         `xml:"EquidTypeOfUsage,omitempty" json:"EquidTypeOfUsage,omitempty" yaml:"EquidTypeOfUsage,omitempty"`
	DeathDate                         DateTime                                        `xml:"DeathDate,omitempty" json:"DeathDate,omitempty" yaml:"DeathDate,omitempty"`
	IsPassIssuerPermittedToChangeData bool                                            `xml:"IsPassIssuerPermittedToChangeData,omitempty" json:"IsPassIssuerPermittedToChangeData,omitempty" yaml:"IsPassIssuerPermittedToChangeData,omitempty"`
	EquidWithersClass                 *TranslatedEnumTypeOfEnumEquidWithersClass      `xml:"EquidWithersClass,omitempty" json:"EquidWithersClass,omitempty" yaml:"EquidWithersClass,omitempty"`
	EquidNotificationState            *TranslatedEnumTypeOfEnumEquidNotificationState `xml:"EquidNotificationState,omitempty" json:"EquidNotificationState,omitempty" yaml:"EquidNotificationState,omitempty"`
	EquidPassIssuingState             *TranslatedEnumTypeOfEnumEquidPassIssuingState  `xml:"EquidPassIssuingState,omitempty" json:"EquidPassIssuingState,omitempty" yaml:"EquidPassIssuingState,omitempty"`
}

// EquidBasicData was auto-generated from WSDL.
type EquidBasicData struct {
	Name              string                                `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	SportBreedingName string                                `xml:"SportBreedingName,omitempty" json:"SportBreedingName,omitempty" yaml:"SportBreedingName,omitempty"`
	BirthDate         DateTime                              `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	HerdBookNumber    string                                `xml:"HerdBookNumber,omitempty" json:"HerdBookNumber,omitempty" yaml:"HerdBookNumber,omitempty"`
	PlaceOfBirth      string                                `xml:"PlaceOfBirth,omitempty" json:"PlaceOfBirth,omitempty" yaml:"PlaceOfBirth,omitempty"`
	ColorFreeText     string                                `xml:"ColorFreeText,omitempty" json:"ColorFreeText,omitempty" yaml:"ColorFreeText,omitempty"`
	EquidColor        EnumEquidColor                        `xml:"EquidColor,omitempty" json:"EquidColor,omitempty" yaml:"EquidColor,omitempty"`
	EquidGender       *TranslatedEnumTypeOfEnumGender       `xml:"EquidGender,omitempty" json:"EquidGender,omitempty" yaml:"EquidGender,omitempty"`
	EquidSpecies      *TranslatedEnumTypeOfEnumEquidSpecies `xml:"EquidSpecies,omitempty" json:"EquidSpecies,omitempty" yaml:"EquidSpecies,omitempty"`
	EquidBreed        *TranslatedEnumTypeOfEnumEquidBreed   `xml:"EquidBreed,omitempty" json:"EquidBreed,omitempty" yaml:"EquidBreed,omitempty"`
	UelnMother        string                                `xml:"UelnMother,omitempty" json:"UelnMother,omitempty" yaml:"UelnMother,omitempty"`
	UelnGeneticMother string                                `xml:"UelnGeneticMother,omitempty" json:"UelnGeneticMother,omitempty" yaml:"UelnGeneticMother,omitempty"`
}

// EquidBasicDataRequest was auto-generated from WSDL.
type EquidBasicDataRequest struct {
	OnBehalfOfAgateNumber string                `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string                `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	EquidChangeBasicData  *EquidChangeBasicData `xml:"EquidChangeBasicData,omitempty" json:"EquidChangeBasicData,omitempty" yaml:"EquidChangeBasicData,omitempty"`
}

// EquidBirthRequest was auto-generated from WSDL.
type EquidBirthRequest struct {
	OnBehalfOfAgateNumber       string                           `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	BirthDate                   DateTime                         `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	HerdBookNumber              string                           `xml:"HerdBookNumber,omitempty" json:"HerdBookNumber,omitempty" yaml:"HerdBookNumber,omitempty"`
	TVDNumber                   int                              `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	IsMultipleBirth             bool                             `xml:"IsMultipleBirth,omitempty" json:"IsMultipleBirth,omitempty" yaml:"IsMultipleBirth,omitempty"`
	Name                        string                           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	EquidSpecies                EnumEquidSpecies                 `xml:"EquidSpecies,omitempty" json:"EquidSpecies,omitempty" yaml:"EquidSpecies,omitempty"`
	EquidBreed                  EnumEquidBreed                   `xml:"EquidBreed,omitempty" json:"EquidBreed,omitempty" yaml:"EquidBreed,omitempty"`
	EquidColor                  EnumEquidColor                   `xml:"EquidColor,omitempty" json:"EquidColor,omitempty" yaml:"EquidColor,omitempty"`
	Gender                      EnumGender                       `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	UelnMother                  string                           `xml:"UelnMother,omitempty" json:"UelnMother,omitempty" yaml:"UelnMother,omitempty"`
	UelnGeneticMother           string                           `xml:"UelnGeneticMother,omitempty" json:"UelnGeneticMother,omitempty" yaml:"UelnGeneticMother,omitempty"`
	EquidRudimentaryDescription *EquidRudimentaryDescriptionData `xml:"EquidRudimentaryDescription,omitempty" json:"EquidRudimentaryDescription,omitempty" yaml:"EquidRudimentaryDescription,omitempty"`
	IsPassAvailable             bool                             `xml:"IsPassAvailable,omitempty" json:"IsPassAvailable,omitempty" yaml:"IsPassAvailable,omitempty"`
}

// EquidCastrationRequest was auto-generated from WSDL.
type EquidCastrationRequest struct {
	OnBehalfOfAgateNumber string   `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string   `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	EventDate             DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
}

// EquidCedeOwnershipRequest was auto-generated from WSDL.
type EquidCedeOwnershipRequest struct {
	OnBehalfOfAgateNumber       string   `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                        string   `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	EventDate                   DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	IsNewEquidOwnerLivingAbroad bool     `xml:"IsNewEquidOwnerLivingAbroad,omitempty" json:"IsNewEquidOwnerLivingAbroad,omitempty" yaml:"IsNewEquidOwnerLivingAbroad,omitempty"`
	CededToPersonAgateNumber    string   `xml:"CededToPersonAgateNumber,omitempty" json:"CededToPersonAgateNumber,omitempty" yaml:"CededToPersonAgateNumber,omitempty"`
}

// EquidChangeBasicData was auto-generated from WSDL.
type EquidChangeBasicData struct {
	Name              string           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	SportBreedingName string           `xml:"SportBreedingName,omitempty" json:"SportBreedingName,omitempty" yaml:"SportBreedingName,omitempty"`
	HerdBookNumber    string           `xml:"HerdBookNumber,omitempty" json:"HerdBookNumber,omitempty" yaml:"HerdBookNumber,omitempty"`
	BirthDate         DateTime         `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	PlaceOfBirth      string           `xml:"PlaceOfBirth,omitempty" json:"PlaceOfBirth,omitempty" yaml:"PlaceOfBirth,omitempty"`
	ColorFreeText     string           `xml:"ColorFreeText,omitempty" json:"ColorFreeText,omitempty" yaml:"ColorFreeText,omitempty"`
	EquidColor        EnumEquidColor   `xml:"EquidColor,omitempty" json:"EquidColor,omitempty" yaml:"EquidColor,omitempty"`
	EquidGender       EnumGender       `xml:"EquidGender,omitempty" json:"EquidGender,omitempty" yaml:"EquidGender,omitempty"`
	EquidSpecies      EnumEquidSpecies `xml:"EquidSpecies,omitempty" json:"EquidSpecies,omitempty" yaml:"EquidSpecies,omitempty"`
	EquidBreed        EnumEquidBreed   `xml:"EquidBreed,omitempty" json:"EquidBreed,omitempty" yaml:"EquidBreed,omitempty"`
	UelnMother        string           `xml:"UelnMother,omitempty" json:"UelnMother,omitempty" yaml:"UelnMother,omitempty"`
	UelnGeneticMother string           `xml:"UelnGeneticMother,omitempty" json:"UelnGeneticMother,omitempty" yaml:"UelnGeneticMother,omitempty"`
}

// EquidChipData was auto-generated from WSDL.
type EquidChipData struct {
	ImplementationDateTime DateTime `xml:"ImplementationDateTime,omitempty" json:"ImplementationDateTime,omitempty" yaml:"ImplementationDateTime,omitempty"`
	ChipNumber             string   `xml:"ChipNumber,omitempty" json:"ChipNumber,omitempty" yaml:"ChipNumber,omitempty"`
	ImplementationLocation string   `xml:"ImplementationLocation,omitempty" json:"ImplementationLocation,omitempty" yaml:"ImplementationLocation,omitempty"`
	AgateNumber            string   `xml:"AgateNumber,omitempty" json:"AgateNumber,omitempty" yaml:"AgateNumber,omitempty"`
	FirstName              string   `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`
	LastName               string   `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
}

// EquidDeathRequest was auto-generated from WSDL.
type EquidDeathRequest struct {
	OnBehalfOfAgateNumber string                         `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string                         `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	DeathNotificationType EnumEquidDeathNotificationType `xml:"DeathNotificationType,omitempty" json:"DeathNotificationType,omitempty" yaml:"DeathNotificationType,omitempty"`
	EventDate             DateTime                       `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
}

// EquidDescendantData was auto-generated from WSDL.
type EquidDescendantData struct {
	Name             string                                  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Ueln             string                                  `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	BirthDate        DateTime                                `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	Gender           *TranslatedEnumTypeOfEnumGender         `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	EquidBreed       *TranslatedEnumTypeOfEnumEquidBreed     `xml:"EquidBreed,omitempty" json:"EquidBreed,omitempty" yaml:"EquidBreed,omitempty"`
	EquidTypeOfUsage *TranslatedEnumTypeOfEnumEquidTypeOfUse `xml:"EquidTypeOfUsage,omitempty" json:"EquidTypeOfUsage,omitempty" yaml:"EquidTypeOfUsage,omitempty"`
}

// EquidDetailData was auto-generated from WSDL.
type EquidDetailData struct {
	EquidBasicData                  *EquidBasicData                   `xml:"EquidBasicData,omitempty" json:"EquidBasicData,omitempty" yaml:"EquidBasicData,omitempty"`
	EquidAdditonalInformationData   *EquidAdditonalInformationData    `xml:"EquidAdditonalInformationData,omitempty" json:"EquidAdditonalInformationData,omitempty" yaml:"EquidAdditonalInformationData,omitempty"`
	EquidOwnershipHistoryDataList   *ArrayOfEquidOwnershipHistoryData `xml:"EquidOwnershipHistoryDataList,omitempty" json:"EquidOwnershipHistoryDataList,omitempty" yaml:"EquidOwnershipHistoryDataList,omitempty"`
	EquidHusbandryHistoryDataList   *ArrayOfEquidHusbandryHistoryData `xml:"EquidHusbandryHistoryDataList,omitempty" json:"EquidHusbandryHistoryDataList,omitempty" yaml:"EquidHusbandryHistoryDataList,omitempty"`
	EquidDescendantDataList         *ArrayOfEquidDescendantData       `xml:"EquidDescendantDataList,omitempty" json:"EquidDescendantDataList,omitempty" yaml:"EquidDescendantDataList,omitempty"`
	EquidPassDataList               *ArrayOfEquidPassData             `xml:"EquidPassDataList,omitempty" json:"EquidPassDataList,omitempty" yaml:"EquidPassDataList,omitempty"`
	EquidUelnDataList               *ArrayOfEquidUelnData             `xml:"EquidUelnDataList,omitempty" json:"EquidUelnDataList,omitempty" yaml:"EquidUelnDataList,omitempty"`
	EquidChipDataList               *ArrayOfEquidChipData             `xml:"EquidChipDataList,omitempty" json:"EquidChipDataList,omitempty" yaml:"EquidChipDataList,omitempty"`
	EquidMembershipDataList         *ArrayOfEquidMembershipData       `xml:"EquidMembershipDataList,omitempty" json:"EquidMembershipDataList,omitempty" yaml:"EquidMembershipDataList,omitempty"`
	EquidRudimentaryDescriptionData *EquidRudimentaryDescriptionData  `xml:"EquidRudimentaryDescriptionData,omitempty" json:"EquidRudimentaryDescriptionData,omitempty" yaml:"EquidRudimentaryDescriptionData,omitempty"`
}

// EquidDetailRequest was auto-generated from WSDL.
type EquidDetailRequest struct {
	OnBehalfOfAgateNumber string   `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	UserRole              EnumRole `xml:"UserRole,omitempty" json:"UserRole,omitempty" yaml:"UserRole,omitempty"`
	Ueln                  string   `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
}

// EquidDetailResult was auto-generated from WSDL.
type EquidDetailResult struct {
	Result          *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	EquidDetailData *EquidDetailData  `xml:"EquidDetailData,omitempty" json:"EquidDetailData,omitempty" yaml:"EquidDetailData,omitempty"`
}

// EquidHusbandryHistoryData was auto-generated from WSDL.
type EquidHusbandryHistoryData struct {
	FromDate              DateTime                                       `xml:"FromDate,omitempty" json:"FromDate,omitempty" yaml:"FromDate,omitempty"`
	ToDate                DateTime                                       `xml:"ToDate,omitempty" json:"ToDate,omitempty" yaml:"ToDate,omitempty"`
	Name                  string                                         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	StreetAndStreetNumber string                                         `xml:"StreetAndStreetNumber,omitempty" json:"StreetAndStreetNumber,omitempty" yaml:"StreetAndStreetNumber,omitempty"`
	Postcode              string                                         `xml:"Postcode,omitempty" json:"Postcode,omitempty" yaml:"Postcode,omitempty"`
	City                  string                                         `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	TVDNumber             string                                         `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	AnimalHusbandryType   *TranslatedEnumTypeOfEnumAnimalHusbandryType   `xml:"AnimalHusbandryType,omitempty" json:"AnimalHusbandryType,omitempty" yaml:"AnimalHusbandryType,omitempty"`
	NotificationType      *TranslatedEnumTypeOfEnumEquidNotificationType `xml:"NotificationType,omitempty" json:"NotificationType,omitempty" yaml:"NotificationType,omitempty"`
}

// EquidImportAfterExportRequest was auto-generated from WSDL.
type EquidImportAfterExportRequest struct {
	OnBehalfOfAgateNumber string      `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string      `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	EventDate             DateTime    `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	ImportCountry         EnumCountry `xml:"ImportCountry,omitempty" json:"ImportCountry,omitempty" yaml:"ImportCountry,omitempty"`
	TVDNumber             int         `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
}

// EquidImportRequest was auto-generated from WSDL.
type EquidImportRequest struct {
	OnBehalfOfAgateNumber string             `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string             `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	EventDate             DateTime           `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	ChipNumber            string             `xml:"ChipNumber,omitempty" json:"ChipNumber,omitempty" yaml:"ChipNumber,omitempty"`
	HerdBookNumber        string             `xml:"HerdBookNumber,omitempty" json:"HerdBookNumber,omitempty" yaml:"HerdBookNumber,omitempty"`
	Country               EnumCountry        `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	TVDNumber             int                `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	Name                  string             `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	SportBreedingName     string             `xml:"SportBreedingName,omitempty" json:"SportBreedingName,omitempty" yaml:"SportBreedingName,omitempty"`
	BirthDate             DateTime           `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	EquidSpecies          EnumEquidSpecies   `xml:"EquidSpecies,omitempty" json:"EquidSpecies,omitempty" yaml:"EquidSpecies,omitempty"`
	EquidBreed            EnumEquidBreed     `xml:"EquidBreed,omitempty" json:"EquidBreed,omitempty" yaml:"EquidBreed,omitempty"`
	EquidColor            string             `xml:"EquidColor,omitempty" json:"EquidColor,omitempty" yaml:"EquidColor,omitempty"`
	Gender                EnumGender         `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	IsCastrated           bool               `xml:"IsCastrated,omitempty" json:"IsCastrated,omitempty" yaml:"IsCastrated,omitempty"`
	EquidTypeOfUse        EnumEquidTypeOfUse `xml:"EquidTypeOfUse,omitempty" json:"EquidTypeOfUse,omitempty" yaml:"EquidTypeOfUse,omitempty"`
	IsPassAvailable       bool               `xml:"IsPassAvailable,omitempty" json:"IsPassAvailable,omitempty" yaml:"IsPassAvailable,omitempty"`
	UelnGeneticMother     string             `xml:"UelnGeneticMother,omitempty" json:"UelnGeneticMother,omitempty" yaml:"UelnGeneticMother,omitempty"`
}

// EquidInitialResult was auto-generated from WSDL.
type EquidInitialResult struct {
	Ueln   string            `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	Result *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

// EquidItem was auto-generated from WSDL.
type EquidItem struct {
	OwnerFirstName    string                                          `xml:"OwnerFirstName,omitempty" json:"OwnerFirstName,omitempty" yaml:"OwnerFirstName,omitempty"`
	OwnerLastName     string                                          `xml:"OwnerLastName,omitempty" json:"OwnerLastName,omitempty" yaml:"OwnerLastName,omitempty"`
	OwnerAgateNumber  string                                          `xml:"OwnerAgateNumber,omitempty" json:"OwnerAgateNumber,omitempty" yaml:"OwnerAgateNumber,omitempty"`
	UELN              string                                          `xml:"UELN,omitempty" json:"UELN,omitempty" yaml:"UELN,omitempty"`
	Name              string                                          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	OriginTVDNumber   int                                             `xml:"OriginTVDNumber,omitempty" json:"OriginTVDNumber,omitempty" yaml:"OriginTVDNumber,omitempty"`
	Husbandry         *HusbandryResult                                `xml:"Husbandry,omitempty" json:"Husbandry,omitempty" yaml:"Husbandry,omitempty"`
	NotificationState *TranslatedEnumTypeOfEnumEquidNotificationState `xml:"NotificationState,omitempty" json:"NotificationState,omitempty" yaml:"NotificationState,omitempty"`
	TypeOfUse         *TranslatedEnumTypeOfEnumEquidTypeOfUse         `xml:"TypeOfUse,omitempty" json:"TypeOfUse,omitempty" yaml:"TypeOfUse,omitempty"`
	Gender            *TranslatedEnumTypeOfEnumGender                 `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	Breed             *TranslatedEnumTypeOfEnumEquidBreed             `xml:"Breed,omitempty" json:"Breed,omitempty" yaml:"Breed,omitempty"`
	WithersClass      *TranslatedEnumTypeOfEnumEquidWithersClass      `xml:"WithersClass,omitempty" json:"WithersClass,omitempty" yaml:"WithersClass,omitempty"`
	IsPassPresent     bool                                            `xml:"IsPassPresent,omitempty" json:"IsPassPresent,omitempty" yaml:"IsPassPresent,omitempty"`
	ColorFreeText     string                                          `xml:"ColorFreeText,omitempty" json:"ColorFreeText,omitempty" yaml:"ColorFreeText,omitempty"`
	BirthDate         DateTime                                        `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	DeathDate         DateTime                                        `xml:"DeathDate,omitempty" json:"DeathDate,omitempty" yaml:"DeathDate,omitempty"`
	ArrivalDate       DateTime                                        `xml:"ArrivalDate,omitempty" json:"ArrivalDate,omitempty" yaml:"ArrivalDate,omitempty"`
	LeavingDate       DateTime                                        `xml:"LeavingDate,omitempty" json:"LeavingDate,omitempty" yaml:"LeavingDate,omitempty"`
}

// EquidItemMO was auto-generated from WSDL.
type EquidItemMO struct {
	EquidUelnItemList *ArrayOfEquidUelnItem         `xml:"EquidUelnItemList,omitempty" json:"EquidUelnItemList,omitempty" yaml:"EquidUelnItemList,omitempty"`
	NotificationList  *ArrayOfEquidNotificationItem `xml:"NotificationList,omitempty" json:"NotificationList,omitempty" yaml:"NotificationList,omitempty"`
}

// EquidLivestockRequest was auto-generated from WSDL.
type EquidLivestockRequest struct {
	TVDNumber      int      `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	SearchDateFrom DateTime `xml:"SearchDateFrom,omitempty" json:"SearchDateFrom,omitempty" yaml:"SearchDateFrom,omitempty"`
	SearchDateTo   DateTime `xml:"SearchDateTo,omitempty" json:"SearchDateTo,omitempty" yaml:"SearchDateTo,omitempty"`
}

// EquidMembershipData was auto-generated from WSDL.
type EquidMembershipData struct {
	MembershipOrganisationAgateNumber string `xml:"MembershipOrganisationAgateNumber,omitempty" json:"MembershipOrganisationAgateNumber,omitempty" yaml:"MembershipOrganisationAgateNumber,omitempty"`
	MembershipOrganisationName        string `xml:"MembershipOrganisationName,omitempty" json:"MembershipOrganisationName,omitempty" yaml:"MembershipOrganisationName,omitempty"`
}

// EquidMembershipOrganisationRequest was auto-generated from WSDL.
type EquidMembershipOrganisationRequest struct {
	OnBehalfOfAgateNumber              string         `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                               string         `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	MembershipOrganisationAgateNumbers *ArrayOfstring `xml:"MembershipOrganisationAgateNumbers,omitempty" json:"MembershipOrganisationAgateNumbers,omitempty" yaml:"MembershipOrganisationAgateNumbers,omitempty"`
}

// EquidNewChipRequest was auto-generated from WSDL.
type EquidNewChipRequest struct {
	OnBehalfOfAgateNumber  string   `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                   string   `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	ChipNumber             string   `xml:"ChipNumber,omitempty" json:"ChipNumber,omitempty" yaml:"ChipNumber,omitempty"`
	ImplementationDate     DateTime `xml:"ImplementationDate,omitempty" json:"ImplementationDate,omitempty" yaml:"ImplementationDate,omitempty"`
	ImplementationLocation string   `xml:"ImplementationLocation,omitempty" json:"ImplementationLocation,omitempty" yaml:"ImplementationLocation,omitempty"`
}

// EquidNotificationItem was auto-generated from WSDL.
type EquidNotificationItem struct {
	NotificationType *TranslatedEnumTypeOfEnumEquidNotificationType `xml:"NotificationType,omitempty" json:"NotificationType,omitempty" yaml:"NotificationType,omitempty"`
	NotificationDate DateTime                                       `xml:"NotificationDate,omitempty" json:"NotificationDate,omitempty" yaml:"NotificationDate,omitempty"`
	EventDate        DateTime                                       `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	DeletionDate     DateTime                                       `xml:"DeletionDate,omitempty" json:"DeletionDate,omitempty" yaml:"DeletionDate,omitempty"`
}

// EquidNotificationsForMembershipOrganisationRequest was auto-generated
// from WSDL.
type EquidNotificationsForMembershipOrganisationRequest struct {
	MembershipOrganisationAgateNumber string   `xml:"MembershipOrganisationAgateNumber,omitempty" json:"MembershipOrganisationAgateNumber,omitempty" yaml:"MembershipOrganisationAgateNumber,omitempty"`
	DateFrom                          DateTime `xml:"DateFrom,omitempty" json:"DateFrom,omitempty" yaml:"DateFrom,omitempty"`
	DateTo                            DateTime `xml:"DateTo,omitempty" json:"DateTo,omitempty" yaml:"DateTo,omitempty"`
}

// EquidOrderBasePassRequest was auto-generated from WSDL.
type EquidOrderBasePassRequest struct {
	OnBehalfOfAgateNumber string                 `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string                 `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	EquidPassOrderType    EnumEquidPassOrderType `xml:"EquidPassOrderType,omitempty" json:"EquidPassOrderType,omitempty" yaml:"EquidPassOrderType,omitempty"`
}

// EquidOwnershipHistoryData was auto-generated from WSDL.
type EquidOwnershipHistoryData struct {
	FromDate              DateTime `xml:"FromDate,omitempty" json:"FromDate,omitempty" yaml:"FromDate,omitempty"`
	ToDate                DateTime `xml:"ToDate,omitempty" json:"ToDate,omitempty" yaml:"ToDate,omitempty"`
	Name                  string   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	StreetAndStreetNumber string   `xml:"StreetAndStreetNumber,omitempty" json:"StreetAndStreetNumber,omitempty" yaml:"StreetAndStreetNumber,omitempty"`
	Postcode              string   `xml:"Postcode,omitempty" json:"Postcode,omitempty" yaml:"Postcode,omitempty"`
	City                  string   `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	AgateNumber           string   `xml:"AgateNumber,omitempty" json:"AgateNumber,omitempty" yaml:"AgateNumber,omitempty"`
	AGIS                  int      `xml:"AGIS,omitempty" json:"AGIS,omitempty" yaml:"AGIS,omitempty"`
}

// EquidOwnershipListRequest was auto-generated from WSDL.
type EquidOwnershipListRequest struct {
	OnBehalfOfAgateNumber                      string   `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	SearchDateFrom                             DateTime `xml:"SearchDateFrom,omitempty" json:"SearchDateFrom,omitempty" yaml:"SearchDateFrom,omitempty"`
	SearchDateTo                               DateTime `xml:"SearchDateTo,omitempty" json:"SearchDateTo,omitempty" yaml:"SearchDateTo,omitempty"`
	ShowOnlyEquidsWhichWereLivingInQueryPeriod bool     `xml:"ShowOnlyEquidsWhichWereLivingInQueryPeriod,omitempty" json:"ShowOnlyEquidsWhichWereLivingInQueryPeriod,omitempty" yaml:"ShowOnlyEquidsWhichWereLivingInQueryPeriod,omitempty"`
}

// EquidPassData was auto-generated from WSDL.
type EquidPassData struct {
	PassType         *TranslatedEnumTypeOfEnumEquidPassType `xml:"PassType,omitempty" json:"PassType,omitempty" yaml:"PassType,omitempty"`
	IssueDate        DateTime                               `xml:"IssueDate,omitempty" json:"IssueDate,omitempty" yaml:"IssueDate,omitempty"`
	AgateNumber      string                                 `xml:"AgateNumber,omitempty" json:"AgateNumber,omitempty" yaml:"AgateNumber,omitempty"`
	PassportIssuer   string                                 `xml:"PassportIssuer,omitempty" json:"PassportIssuer,omitempty" yaml:"PassportIssuer,omitempty"`
	NotificationDate DateTime                               `xml:"NotificationDate,omitempty" json:"NotificationDate,omitempty" yaml:"NotificationDate,omitempty"`
}

// EquidPassIssuedRequest was auto-generated from WSDL.
type EquidPassIssuedRequest struct {
	OnBehalfOfAgateNumber string            `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string            `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	EquidPassType         EnumEquidPassType `xml:"EquidPassType,omitempty" json:"EquidPassType,omitempty" yaml:"EquidPassType,omitempty"`
	EventDate             DateTime          `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
}

// EquidPassIssuerPermissionRequest was auto-generated from WSDL.
type EquidPassIssuerPermissionRequest struct {
	OnBehalfOfAgateNumber    string `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                     string `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	IsEquidPassIssuerAllowed bool   `xml:"IsEquidPassIssuerAllowed,omitempty" json:"IsEquidPassIssuerAllowed,omitempty" yaml:"IsEquidPassIssuerAllowed,omitempty"`
}

// EquidRelocationRequest was auto-generated from WSDL.
type EquidRelocationRequest struct {
	OnBehalfOfAgateNumber string                      `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string                      `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	EventDate             DateTime                    `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	LocationChangeType    EnumEquidLocationChangeType `xml:"LocationChangeType,omitempty" json:"LocationChangeType,omitempty" yaml:"LocationChangeType,omitempty"`
	NewTVDNumber          int                         `xml:"NewTVDNumber,omitempty" json:"NewTVDNumber,omitempty" yaml:"NewTVDNumber,omitempty"`
	OriginTVDNumber       int                         `xml:"OriginTVDNumber,omitempty" json:"OriginTVDNumber,omitempty" yaml:"OriginTVDNumber,omitempty"`
	ImportOrExportCountry EnumCountry                 `xml:"ImportOrExportCountry,omitempty" json:"ImportOrExportCountry,omitempty" yaml:"ImportOrExportCountry,omitempty"`
}

// EquidRudimentaryDescriptionData was auto-generated from WSDL.
type EquidRudimentaryDescriptionData struct {
	HasWhiteOnHead          bool `xml:"HasWhiteOnHead,omitempty" json:"HasWhiteOnHead,omitempty" yaml:"HasWhiteOnHead,omitempty"`
	HasWhiteOnLegFrontLeft  bool `xml:"HasWhiteOnLegFrontLeft,omitempty" json:"HasWhiteOnLegFrontLeft,omitempty" yaml:"HasWhiteOnLegFrontLeft,omitempty"`
	HasWhiteOnLegFrontRight bool `xml:"HasWhiteOnLegFrontRight,omitempty" json:"HasWhiteOnLegFrontRight,omitempty" yaml:"HasWhiteOnLegFrontRight,omitempty"`
	HasWhiteOnLegBackLeft   bool `xml:"HasWhiteOnLegBackLeft,omitempty" json:"HasWhiteOnLegBackLeft,omitempty" yaml:"HasWhiteOnLegBackLeft,omitempty"`
	HasWhiteOnLegBackRight  bool `xml:"HasWhiteOnLegBackRight,omitempty" json:"HasWhiteOnLegBackRight,omitempty" yaml:"HasWhiteOnLegBackRight,omitempty"`
}

// EquidSearchRequest was auto-generated from WSDL.
type EquidSearchRequest struct {
	OnBehalfOfAgateNumber    string         `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	UserRole                 EnumRole       `xml:"UserRole,omitempty" json:"UserRole,omitempty" yaml:"UserRole,omitempty"`
	Ueln                     string         `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	ChipNumber               string         `xml:"ChipNumber,omitempty" json:"ChipNumber,omitempty" yaml:"ChipNumber,omitempty"`
	AnimalHusbandryTVDNumber int            `xml:"AnimalHusbandryTVDNumber,omitempty" json:"AnimalHusbandryTVDNumber,omitempty" yaml:"AnimalHusbandryTVDNumber,omitempty"`
	Name                     string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	HerdBookNumber           string         `xml:"HerdBookNumber,omitempty" json:"HerdBookNumber,omitempty" yaml:"HerdBookNumber,omitempty"`
	OwnerAgateNumber         string         `xml:"OwnerAgateNumber,omitempty" json:"OwnerAgateNumber,omitempty" yaml:"OwnerAgateNumber,omitempty"`
	SportBreedingName        string         `xml:"SportBreedingName,omitempty" json:"SportBreedingName,omitempty" yaml:"SportBreedingName,omitempty"`
	BirthDateFrom            DateTime       `xml:"BirthDateFrom,omitempty" json:"BirthDateFrom,omitempty" yaml:"BirthDateFrom,omitempty"`
	BirthDateTo              DateTime       `xml:"BirthDateTo,omitempty" json:"BirthDateTo,omitempty" yaml:"BirthDateTo,omitempty"`
	EquidBreed               EnumEquidBreed `xml:"EquidBreed,omitempty" json:"EquidBreed,omitempty" yaml:"EquidBreed,omitempty"`
	ShowOnlyLivingEquids     bool           `xml:"ShowOnlyLivingEquids,omitempty" json:"ShowOnlyLivingEquids,omitempty" yaml:"ShowOnlyLivingEquids,omitempty"`
}

// EquidSearchResult was auto-generated from WSDL.
type EquidSearchResult struct {
	Result        *ProcessingResult             `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	SearchResults *ArrayOfEquidSearchResultData `xml:"SearchResults,omitempty" json:"SearchResults,omitempty" yaml:"SearchResults,omitempty"`
}

// EquidSearchResultData was auto-generated from WSDL.
type EquidSearchResultData struct {
	Name              string                                  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Ueln              string                                  `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	ChipNumber        string                                  `xml:"ChipNumber,omitempty" json:"ChipNumber,omitempty" yaml:"ChipNumber,omitempty"`
	OwnerAgateNumber  string                                  `xml:"OwnerAgateNumber,omitempty" json:"OwnerAgateNumber,omitempty" yaml:"OwnerAgateNumber,omitempty"`
	HerdBookNumber    string                                  `xml:"HerdBookNumber,omitempty" json:"HerdBookNumber,omitempty" yaml:"HerdBookNumber,omitempty"`
	TVDNumber         string                                  `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	BirthDate         DateTime                                `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	SportBreedingName string                                  `xml:"SportBreedingName,omitempty" json:"SportBreedingName,omitempty" yaml:"SportBreedingName,omitempty"`
	DeathDate         DateTime                                `xml:"DeathDate,omitempty" json:"DeathDate,omitempty" yaml:"DeathDate,omitempty"`
	EquidTypeOfUse    *TranslatedEnumTypeOfEnumEquidTypeOfUse `xml:"EquidTypeOfUse,omitempty" json:"EquidTypeOfUse,omitempty" yaml:"EquidTypeOfUse,omitempty"`
}

// EquidSlaughterListData was auto-generated from WSDL.
type EquidSlaughterListData struct {
	Name                              string   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Ueln                              string   `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	ChipNumber                        string   `xml:"ChipNumber,omitempty" json:"ChipNumber,omitempty" yaml:"ChipNumber,omitempty"`
	OriginAnimalHusbandryTVDNumber    int      `xml:"OriginAnimalHusbandryTVDNumber,omitempty" json:"OriginAnimalHusbandryTVDNumber,omitempty" yaml:"OriginAnimalHusbandryTVDNumber,omitempty"`
	BirthDate                         DateTime `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	SlaughterDate                     DateTime `xml:"SlaughterDate,omitempty" json:"SlaughterDate,omitempty" yaml:"SlaughterDate,omitempty"`
	HasDisposalFee                    bool     `xml:"HasDisposalFee,omitempty" json:"HasDisposalFee,omitempty" yaml:"HasDisposalFee,omitempty"`
	AnimalHusbandrySlaughterInitiator int      `xml:"AnimalHusbandrySlaughterInitiator,omitempty" json:"AnimalHusbandrySlaughterInitiator,omitempty" yaml:"AnimalHusbandrySlaughterInitiator,omitempty"`
}

// EquidSlaughterListRequest was auto-generated from WSDL.
type EquidSlaughterListRequest struct {
	TVDNumber int      `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	DateFrom  DateTime `xml:"DateFrom,omitempty" json:"DateFrom,omitempty" yaml:"DateFrom,omitempty"`
	DateTo    DateTime `xml:"DateTo,omitempty" json:"DateTo,omitempty" yaml:"DateTo,omitempty"`
}

// EquidSlaughterListResult was auto-generated from WSDL.
type EquidSlaughterListResult struct {
	Result                     *ProcessingResult              `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	EquidSlaughterListDataList *ArrayOfEquidSlaughterListData `xml:"EquidSlaughterListDataList,omitempty" json:"EquidSlaughterListDataList,omitempty" yaml:"EquidSlaughterListDataList,omitempty"`
}

// EquidSlaughterRequest was auto-generated from WSDL.
type EquidSlaughterRequest struct {
	TVDNumber                      int      `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	Ueln                           string   `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	ChipNumber                     string   `xml:"ChipNumber,omitempty" json:"ChipNumber,omitempty" yaml:"ChipNumber,omitempty"`
	EventDate                      DateTime `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	OriginAnimalHusbandryTVDNumber int      `xml:"OriginAnimalHusbandryTVDNumber,omitempty" json:"OriginAnimalHusbandryTVDNumber,omitempty" yaml:"OriginAnimalHusbandryTVDNumber,omitempty"`
	SlaughterInitiatorTVDNumber    int      `xml:"SlaughterInitiatorTVDNumber,omitempty" json:"SlaughterInitiatorTVDNumber,omitempty" yaml:"SlaughterInitiatorTVDNumber,omitempty"`
}

// EquidUelnData was auto-generated from WSDL.
type EquidUelnData struct {
	Ueln             string   `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	NotificationDate DateTime `xml:"NotificationDate,omitempty" json:"NotificationDate,omitempty" yaml:"NotificationDate,omitempty"`
	AgateNumber      string   `xml:"AgateNumber,omitempty" json:"AgateNumber,omitempty" yaml:"AgateNumber,omitempty"`
	FirstName        string   `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`
	LastName         string   `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
}

// EquidUelnItem was auto-generated from WSDL.
type EquidUelnItem struct {
	Ueln     string `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	IsActive bool   `xml:"IsActive,omitempty" json:"IsActive,omitempty" yaml:"IsActive,omitempty"`
}

// EquidWithersClassRequest was auto-generated from WSDL.
type EquidWithersClassRequest struct {
	OnBehalfOfAgateNumber string                `xml:"OnBehalfOfAgateNumber,omitempty" json:"OnBehalfOfAgateNumber,omitempty" yaml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string                `xml:"Ueln,omitempty" json:"Ueln,omitempty" yaml:"Ueln,omitempty"`
	WithersClass          EnumEquidWithersClass `xml:"WithersClass,omitempty" json:"WithersClass,omitempty" yaml:"WithersClass,omitempty"`
}

// EquidsWithNotificationsOfMOResult was auto-generated from WSDL.
type EquidsWithNotificationsOfMOResult struct {
	Result    *ProcessingResult   `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	EquidList *ArrayOfEquidItemMO `xml:"EquidList,omitempty" json:"EquidList,omitempty" yaml:"EquidList,omitempty"`
}

// FarmManagerResult was auto-generated from WSDL.
type FarmManagerResult struct {
	AgateNumber           string            `xml:"AgateNumber,omitempty" json:"AgateNumber,omitempty" yaml:"AgateNumber,omitempty"`
	Title                 string            `xml:"Title,omitempty" json:"Title,omitempty" yaml:"Title,omitempty"`
	LastName              string            `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
	FirstName             string            `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`
	Street                string            `xml:"Street,omitempty" json:"Street,omitempty" yaml:"Street,omitempty"`
	PostCode              string            `xml:"PostCode,omitempty" json:"PostCode,omitempty" yaml:"PostCode,omitempty"`
	City                  string            `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	EmailAddress          string            `xml:"EmailAddress,omitempty" json:"EmailAddress,omitempty" yaml:"EmailAddress,omitempty"`
	PhoneNumbers          *StringArray      `xml:"PhoneNumbers,omitempty" json:"PhoneNumbers,omitempty" yaml:"PhoneNumbers,omitempty"`
	PreferredLanguageLCID int               `xml:"PreferredLanguageLCID,omitempty" json:"PreferredLanguageLCID,omitempty" yaml:"PreferredLanguageLCID,omitempty"`
	Result                *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

// GeneraResult was auto-generated from WSDL.
type GeneraResult struct {
	Result *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	Genera *IntArray         `xml:"Genera,omitempty" json:"Genera,omitempty" yaml:"Genera,omitempty"`
}

// GenusElements was auto-generated from WSDL.
type GenusElements struct {
	Genus []*TranslatedEnumTypeOfEnumGenus `xml:"Genus,omitempty" json:"Genus,omitempty" yaml:"Genus,omitempty"`
}

// GetAnimalHusbandryAddress was auto-generated from WSDL.
type GetAnimalHusbandryAddress struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddress" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
}

// GetAnimalHusbandryAddressResponse was auto-generated from WSDL.
type GetAnimalHusbandryAddressResponse struct {
	GetAnimalHusbandryAddressResult *AnimalHusbandryAddressResult `xml:"GetAnimalHusbandryAddressResult,omitempty" json:"GetAnimalHusbandryAddressResult,omitempty" yaml:"GetAnimalHusbandryAddressResult,omitempty"`
}

// GetAnimalHusbandryAddressV2 was auto-generated from WSDL.
type GetAnimalHusbandryAddressV2 struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddressV2" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
}

// GetAnimalHusbandryAddressV2Response was auto-generated from
// WSDL.
type GetAnimalHusbandryAddressV2Response struct {
	GetAnimalHusbandryAddressV2Result *AnimalHusbandryAddressResultV2 `xml:"GetAnimalHusbandryAddressV2Result,omitempty" json:"GetAnimalHusbandryAddressV2Result,omitempty" yaml:"GetAnimalHusbandryAddressV2Result,omitempty"`
}

// GetAnimalHusbandryDetail was auto-generated from WSDL.
type GetAnimalHusbandryDetail struct {
	XMLName                           xml.Name                         `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryDetail" json:"-" yaml:"-"`
	P_GetAnimalHusbandryDetailRequest *GetAnimalHusbandryDetailRequest `xml:"p_GetAnimalHusbandryDetailRequest" json:"p_GetAnimalHusbandryDetailRequest" yaml:"p_GetAnimalHusbandryDetailRequest"`
}

// GetAnimalHusbandryDetailRequest was auto-generated from WSDL.
type GetAnimalHusbandryDetailRequest struct {
	ManufacturerKey string `xml:"ManufacturerKey,omitempty" json:"ManufacturerKey,omitempty" yaml:"ManufacturerKey,omitempty"`
	LCID            int    `xml:"LCID,omitempty" json:"LCID,omitempty" yaml:"LCID,omitempty"`
	TVDNumber       int    `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
}

// GetAnimalHusbandryDetailResponse was auto-generated from WSDL.
type GetAnimalHusbandryDetailResponse struct {
	GetAnimalHusbandryDetailResult *GetAnimalHusbandryDetailResult `xml:"GetAnimalHusbandryDetailResult,omitempty" json:"GetAnimalHusbandryDetailResult,omitempty" yaml:"GetAnimalHusbandryDetailResult,omitempty"`
}

// GetAnimalHusbandryDetailResult was auto-generated from WSDL.
type GetAnimalHusbandryDetailResult struct {
	Result                   *ProcessingResult                                 `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	PostData                 *HusbandryResult                                  `xml:"PostData,omitempty" json:"PostData,omitempty" yaml:"PostData,omitempty"`
	IsActive                 bool                                              `xml:"IsActive,omitempty" json:"IsActive,omitempty" yaml:"IsActive,omitempty"`
	MunicipalityName         string                                            `xml:"MunicipalityName,omitempty" json:"MunicipalityName,omitempty" yaml:"MunicipalityName,omitempty"`
	CantonShortname          string                                            `xml:"CantonShortname,omitempty" json:"CantonShortname,omitempty" yaml:"CantonShortname,omitempty"`
	CoordinateX              int                                               `xml:"CoordinateX,omitempty" json:"CoordinateX,omitempty" yaml:"CoordinateX,omitempty"`
	CoordinateY              int                                               `xml:"CoordinateY,omitempty" json:"CoordinateY,omitempty" yaml:"CoordinateY,omitempty"`
	Altitude                 int                                               `xml:"Altitude,omitempty" json:"Altitude,omitempty" yaml:"Altitude,omitempty"`
	CantonAnimalHusbandryKey string                                            `xml:"CantonAnimalHusbandryKey,omitempty" json:"CantonAnimalHusbandryKey,omitempty" yaml:"CantonAnimalHusbandryKey,omitempty"`
	CantonPersonKey          string                                            `xml:"CantonPersonKey,omitempty" json:"CantonPersonKey,omitempty" yaml:"CantonPersonKey,omitempty"`
	BurNumber                string                                            `xml:"BurNumber,omitempty" json:"BurNumber,omitempty" yaml:"BurNumber,omitempty"`
	AnimalHusbandryType      *TranslatedEnumTypeOfEnumAnimalHusbandryType      `xml:"AnimalHusbandryType,omitempty" json:"AnimalHusbandryType,omitempty" yaml:"AnimalHusbandryType,omitempty"`
	MunicipalityNumber       int                                               `xml:"MunicipalityNumber,omitempty" json:"MunicipalityNumber,omitempty" yaml:"MunicipalityNumber,omitempty"`
	TypeOfUse                *TranslatedEnumTypeOfEnumAnimalHusbandryTypeOfUse `xml:"TypeOfUse,omitempty" json:"TypeOfUse,omitempty" yaml:"TypeOfUse,omitempty"`
	IsMountain               bool                                              `xml:"IsMountain,omitempty" json:"IsMountain,omitempty" yaml:"IsMountain,omitempty"`
	Zone                     *TranslatedEnumTypeOfEnumZone                     `xml:"Zone,omitempty" json:"Zone,omitempty" yaml:"Zone,omitempty"`
	Area                     *TranslatedEnumTypeOfEnumArea                     `xml:"Area,omitempty" json:"Area,omitempty" yaml:"Area,omitempty"`
}

// GetAnimalHusbandryMemberships was auto-generated from WSDL.
type GetAnimalHusbandryMemberships struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryMemberships" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
}

// GetAnimalHusbandryMembershipsResponse was auto-generated from
// WSDL.
type GetAnimalHusbandryMembershipsResponse struct {
	GetAnimalHusbandryMembershipsResult *HusbandryMembershipResult `xml:"GetAnimalHusbandryMembershipsResult,omitempty" json:"GetAnimalHusbandryMembershipsResult,omitempty" yaml:"GetAnimalHusbandryMembershipsResult,omitempty"`
}

// GetCattleDetail was auto-generated from WSDL.
type GetCattleDetail struct {
	XMLName           xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetail" json:"-" yaml:"-"`
	P_ManufacturerKey string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus" json:"p_WorkingFocus" yaml:"p_WorkingFocus"`
	P_EarTagNumber    string             `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
}

// GetCattleDetailResponse was auto-generated from WSDL.
type GetCattleDetailResponse struct {
	GetCattleDetailResult *CattleDetailResult `xml:"GetCattleDetailResult,omitempty" json:"GetCattleDetailResult,omitempty" yaml:"GetCattleDetailResult,omitempty"`
}

// GetCattleDetailV2 was auto-generated from WSDL.
type GetCattleDetailV2 struct {
	XMLName           xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetailV2" json:"-" yaml:"-"`
	P_ManufacturerKey string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus" json:"p_WorkingFocus" yaml:"p_WorkingFocus"`
	P_EarTagNumber    string             `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
}

// GetCattleDetailV2Response was auto-generated from WSDL.
type GetCattleDetailV2Response struct {
	GetCattleDetailV2Result *CattleDetailResultV2 `xml:"GetCattleDetailV2Result,omitempty" json:"GetCattleDetailV2Result,omitempty" yaml:"GetCattleDetailV2Result,omitempty"`
}

// GetCattleEarTags was auto-generated from WSDL.
type GetCattleEarTags struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleEarTags" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_OnStock         bool     `xml:"p_OnStock" json:"p_OnStock" yaml:"p_OnStock"`
}

// GetCattleEarTagsResponse was auto-generated from WSDL.
type GetCattleEarTagsResponse struct {
	GetCattleEarTagsResult *CattleEarTagsResult `xml:"GetCattleEarTagsResult,omitempty" json:"GetCattleEarTagsResult,omitempty" yaml:"GetCattleEarTagsResult,omitempty"`
}

// GetCattleHistory was auto-generated from WSDL.
type GetCattleHistory struct {
	XMLName           xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleHistory" json:"-" yaml:"-"`
	P_ManufacturerKey string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus" json:"p_WorkingFocus" yaml:"p_WorkingFocus"`
	P_EarTagNumber    string             `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
}

// GetCattleHistoryResponse was auto-generated from WSDL.
type GetCattleHistoryResponse struct {
	GetCattleHistoryResult *CattleHistoryResult `xml:"GetCattleHistoryResult,omitempty" json:"GetCattleHistoryResult,omitempty" yaml:"GetCattleHistoryResult,omitempty"`
}

// GetCattleHistoryV2 was auto-generated from WSDL.
type GetCattleHistoryV2 struct {
	XMLName           xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleHistoryV2" json:"-" yaml:"-"`
	P_ManufacturerKey string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus" json:"p_WorkingFocus" yaml:"p_WorkingFocus"`
	P_EarTagNumber    string             `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
}

// GetCattleHistoryV2Response was auto-generated from WSDL.
type GetCattleHistoryV2Response struct {
	GetCattleHistoryV2Result *CattleHistoryResultV2 `xml:"GetCattleHistoryV2Result,omitempty" json:"GetCattleHistoryV2Result,omitempty" yaml:"GetCattleHistoryV2Result,omitempty"`
}

// GetCattleLivestock was auto-generated from WSDL.
type GetCattleLivestock struct {
	XMLName           xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestock" json:"-" yaml:"-"`
	P_ManufacturerKey string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus" json:"p_WorkingFocus" yaml:"p_WorkingFocus"`
	P_TVDNumber       int                `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_SearchDateFrom  DateTime           `xml:"p_SearchDateFrom" json:"p_SearchDateFrom" yaml:"p_SearchDateFrom"`
	P_SearchDateTo    DateTime           `xml:"p_SearchDateTo" json:"p_SearchDateTo" yaml:"p_SearchDateTo"`
}

// GetCattleLivestockResponse was auto-generated from WSDL.
type GetCattleLivestockResponse struct {
	GetCattleLivestockResult *CattleLivestockResult `xml:"GetCattleLivestockResult,omitempty" json:"GetCattleLivestockResult,omitempty" yaml:"GetCattleLivestockResult,omitempty"`
}

// GetCattleLivestockV2 was auto-generated from WSDL.
type GetCattleLivestockV2 struct {
	XMLName           xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestockV2" json:"-" yaml:"-"`
	P_ManufacturerKey string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus" json:"p_WorkingFocus" yaml:"p_WorkingFocus"`
	P_TVDNumber       int                `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_SearchDateFrom  DateTime           `xml:"p_SearchDateFrom" json:"p_SearchDateFrom" yaml:"p_SearchDateFrom"`
	P_SearchDateTo    DateTime           `xml:"p_SearchDateTo" json:"p_SearchDateTo" yaml:"p_SearchDateTo"`
}

// GetCattleLivestockV2Response was auto-generated from WSDL.
type GetCattleLivestockV2Response struct {
	GetCattleLivestockV2Result *CattleLivestockResultV2 `xml:"GetCattleLivestockV2Result,omitempty" json:"GetCattleLivestockV2Result,omitempty" yaml:"GetCattleLivestockV2Result,omitempty"`
}

// GetCattleMovements was auto-generated from WSDL.
type GetCattleMovements struct {
	XMLName           xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleMovements" json:"-" yaml:"-"`
	P_ManufacturerKey string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus" json:"p_WorkingFocus" yaml:"p_WorkingFocus"`
	P_EarTagNumber    string             `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
}

// GetCattleMovementsResponse was auto-generated from WSDL.
type GetCattleMovementsResponse struct {
	GetCattleMovementsResult *CattleMovementsResult `xml:"GetCattleMovementsResult,omitempty" json:"GetCattleMovementsResult,omitempty" yaml:"GetCattleMovementsResult,omitempty"`
}

// GetCattleOffsprings was auto-generated from WSDL.
type GetCattleOffsprings struct {
	XMLName              xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleOffsprings" json:"-" yaml:"-"`
	P_ManufacturerKey    string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID               int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WorkingFocus       *WorkingFocusArray `xml:"p_WorkingFocus" json:"p_WorkingFocus" yaml:"p_WorkingFocus"`
	P_EarTagNumberMother string             `xml:"p_EarTagNumberMother,omitempty" json:"p_EarTagNumberMother,omitempty" yaml:"p_EarTagNumberMother,omitempty"`
}

// GetCattleOffspringsResponse was auto-generated from WSDL.
type GetCattleOffspringsResponse struct {
	GetCattleOffspringsResult *CattleOffspringResult `xml:"GetCattleOffspringsResult,omitempty" json:"GetCattleOffspringsResult,omitempty" yaml:"GetCattleOffspringsResult,omitempty"`
}

// GetCattleStatus was auto-generated from WSDL.
type GetCattleStatus struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatus" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EarTagNumber    string   `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
}

// GetCattleStatusResponse was auto-generated from WSDL.
type GetCattleStatusResponse struct {
	GetCattleStatusResult *CattleStateExternalResult `xml:"GetCattleStatusResult,omitempty" json:"GetCattleStatusResult,omitempty" yaml:"GetCattleStatusResult,omitempty"`
}

// GetCattleStatusV2 was auto-generated from WSDL.
type GetCattleStatusV2 struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatusV2" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EarTagNumber    string   `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
}

// GetCattleStatusV2Response was auto-generated from WSDL.
type GetCattleStatusV2Response struct {
	GetCattleStatusV2Result *CattleStateExternalResultV2 `xml:"GetCattleStatusV2Result,omitempty" json:"GetCattleStatusV2Result,omitempty" yaml:"GetCattleStatusV2Result,omitempty"`
}

// GetCattlesPerLeavingDate was auto-generated from WSDL.
type GetCattlesPerLeavingDate struct {
	XMLName                           xml.Name                         `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDate" json:"-" yaml:"-"`
	P_ManufacturerKey                 string                           `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                            int                              `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_GetCattlesPerLeavingDateRequest *GetCattlesPerLeavingDateRequest `xml:"p_GetCattlesPerLeavingDateRequest" json:"p_GetCattlesPerLeavingDateRequest" yaml:"p_GetCattlesPerLeavingDateRequest"`
}

// GetCattlesPerLeavingDateData was auto-generated from WSDL.
type GetCattlesPerLeavingDateData struct {
	EarTagNumber          string                              `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	Name                  string                              `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	BirthDate             DateTime                            `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	LeavingDate           DateTime                            `xml:"LeavingDate,omitempty" json:"LeavingDate,omitempty" yaml:"LeavingDate,omitempty"`
	Gender                *TranslatedEnumTypeOfEnumGender     `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	Race                  *TranslatedEnumTypeOfEnumCattleRace `xml:"Race,omitempty" json:"Race,omitempty" yaml:"Race,omitempty"`
	TVDNumber             int                                 `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	HusbandryName         string                              `xml:"HusbandryName,omitempty" json:"HusbandryName,omitempty" yaml:"HusbandryName,omitempty"`
	StreetAndStreetNumber string                              `xml:"StreetAndStreetNumber,omitempty" json:"StreetAndStreetNumber,omitempty" yaml:"StreetAndStreetNumber,omitempty"`
	City                  string                              `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	PostCode              int                                 `xml:"PostCode,omitempty" json:"PostCode,omitempty" yaml:"PostCode,omitempty"`
}

// GetCattlesPerLeavingDateRequest was auto-generated from WSDL.
type GetCattlesPerLeavingDateRequest struct {
	TVDNumber   int      `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	LeavingDate DateTime `xml:"LeavingDate,omitempty" json:"LeavingDate,omitempty" yaml:"LeavingDate,omitempty"`
}

// GetCattlesPerLeavingDateResponse was auto-generated from WSDL.
type GetCattlesPerLeavingDateResponse struct {
	GetCattlesPerLeavingDateResult *GetCattlesPerLeavingDateResult `xml:"GetCattlesPerLeavingDateResult,omitempty" json:"GetCattlesPerLeavingDateResult,omitempty" yaml:"GetCattlesPerLeavingDateResult,omitempty"`
}

// GetCattlesPerLeavingDateResult was auto-generated from WSDL.
type GetCattlesPerLeavingDateResult struct {
	Result        *ProcessingResult                    `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	ResultDetails *ArrayOfGetCattlesPerLeavingDateData `xml:"ResultDetails,omitempty" json:"ResultDetails,omitempty" yaml:"ResultDetails,omitempty"`
}

// GetCodes was auto-generated from WSDL.
type GetCodes struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCodes" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_Type            string   `xml:"p_Type,omitempty" json:"p_Type,omitempty" yaml:"p_Type,omitempty"`
}

// GetCodesResponse was auto-generated from WSDL.
type GetCodesResponse struct {
	GetCodesResult *CodesResult `xml:"GetCodesResult,omitempty" json:"GetCodesResult,omitempty" yaml:"GetCodesResult,omitempty"`
}

// GetEarTagOrders was auto-generated from WSDL.
type GetEarTagOrders struct {
	XMLName           xml.Name  `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEarTagOrders" json:"-" yaml:"-"`
	P_ManufacturerKey string    `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int       `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int       `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_SearchDateFrom  DateTime  `xml:"p_SearchDateFrom" json:"p_SearchDateFrom" yaml:"p_SearchDateFrom"`
	P_SearchDateTo    DateTime  `xml:"p_SearchDateTo" json:"p_SearchDateTo" yaml:"p_SearchDateTo"`
	P_ArticleFilter   *IntArray `xml:"p_ArticleFilter" json:"p_ArticleFilter" yaml:"p_ArticleFilter"`
}

// GetEarTagOrdersResponse was auto-generated from WSDL.
type GetEarTagOrdersResponse struct {
	GetEarTagOrdersResult *EarTagOrderResult `xml:"GetEarTagOrdersResult,omitempty" json:"GetEarTagOrdersResult,omitempty" yaml:"GetEarTagOrdersResult,omitempty"`
}

// GetEquidDetail was auto-generated from WSDL.
type GetEquidDetail struct {
	XMLName              xml.Name            `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidDetail" json:"-" yaml:"-"`
	P_ManufacturerKey    string              `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID               int                 `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidDetailRequest *EquidDetailRequest `xml:"p_EquidDetailRequest" json:"p_EquidDetailRequest" yaml:"p_EquidDetailRequest"`
}

// GetEquidDetailResponse was auto-generated from WSDL.
type GetEquidDetailResponse struct {
	GetEquidDetailResult *EquidDetailResult `xml:"GetEquidDetailResult,omitempty" json:"GetEquidDetailResult,omitempty" yaml:"GetEquidDetailResult,omitempty"`
}

// GetEquidLivestock was auto-generated from WSDL.
type GetEquidLivestock struct {
	XMLName                 xml.Name               `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidLivestock" json:"-" yaml:"-"`
	P_ManufacturerKey       string                 `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                  int                    `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidLivestockRequest *EquidLivestockRequest `xml:"p_EquidLivestockRequest" json:"p_EquidLivestockRequest" yaml:"p_EquidLivestockRequest"`
}

// GetEquidLivestockResponse was auto-generated from WSDL.
type GetEquidLivestockResponse struct {
	GetEquidLivestockResult *GetEquidLivestockResult `xml:"GetEquidLivestockResult,omitempty" json:"GetEquidLivestockResult,omitempty" yaml:"GetEquidLivestockResult,omitempty"`
}

// GetEquidLivestockResult was auto-generated from WSDL.
type GetEquidLivestockResult struct {
	Result    *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	EquidList *ArrayOfEquidItem `xml:"EquidList,omitempty" json:"EquidList,omitempty" yaml:"EquidList,omitempty"`
}

// GetEquidOwnershipList was auto-generated from WSDL.
type GetEquidOwnershipList struct {
	XMLName                     xml.Name                   `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidOwnershipList" json:"-" yaml:"-"`
	P_ManufacturerKey           string                     `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                      int                        `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidOwnershipListRequest *EquidOwnershipListRequest `xml:"p_EquidOwnershipListRequest" json:"p_EquidOwnershipListRequest" yaml:"p_EquidOwnershipListRequest"`
}

// GetEquidOwnershipListResponse was auto-generated from WSDL.
type GetEquidOwnershipListResponse struct {
	GetEquidOwnershipListResult *GetEquidOwnershipListResult `xml:"GetEquidOwnershipListResult,omitempty" json:"GetEquidOwnershipListResult,omitempty" yaml:"GetEquidOwnershipListResult,omitempty"`
}

// GetEquidOwnershipListResult was auto-generated from WSDL.
type GetEquidOwnershipListResult struct {
	Result    *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	EquidList *ArrayOfEquidItem `xml:"EquidList,omitempty" json:"EquidList,omitempty" yaml:"EquidList,omitempty"`
}

// GetEquidSlaughterList was auto-generated from WSDL.
type GetEquidSlaughterList struct {
	XMLName                     xml.Name                   `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidSlaughterList" json:"-" yaml:"-"`
	P_ManufacturerKey           string                     `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                      int                        `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidSlaughterListRequest *EquidSlaughterListRequest `xml:"p_EquidSlaughterListRequest" json:"p_EquidSlaughterListRequest" yaml:"p_EquidSlaughterListRequest"`
}

// GetEquidSlaughterListResponse was auto-generated from WSDL.
type GetEquidSlaughterListResponse struct {
	GetEquidSlaughterListResult *EquidSlaughterListResult `xml:"GetEquidSlaughterListResult,omitempty" json:"GetEquidSlaughterListResult,omitempty" yaml:"GetEquidSlaughterListResult,omitempty"`
}

// GetEquidsWithNotificationsOfMO was auto-generated from WSDL.
type GetEquidsWithNotificationsOfMO struct {
	XMLName                              xml.Name                                            `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidsWithNotificationsOfMO" json:"-" yaml:"-"`
	P_ManufacturerKey                    string                                              `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                               int                                                 `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidsWithNotificationsOfMoRequest *EquidNotificationsForMembershipOrganisationRequest `xml:"p_EquidsWithNotificationsOfMoRequest" json:"p_EquidsWithNotificationsOfMoRequest" yaml:"p_EquidsWithNotificationsOfMoRequest"`
}

// GetEquidsWithNotificationsOfMOResponse was auto-generated from
// WSDL.
type GetEquidsWithNotificationsOfMOResponse struct {
	GetEquidsWithNotificationsOfMOResult *EquidsWithNotificationsOfMOResult `xml:"GetEquidsWithNotificationsOfMOResult,omitempty" json:"GetEquidsWithNotificationsOfMOResult,omitempty" yaml:"GetEquidsWithNotificationsOfMOResult,omitempty"`
}

// GetFarmManager was auto-generated from WSDL.
type GetFarmManager struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmManager" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
}

// GetFarmManagerResponse was auto-generated from WSDL.
type GetFarmManagerResponse struct {
	GetFarmManagerResult *FarmManagerResult `xml:"GetFarmManagerResult,omitempty" json:"GetFarmManagerResult,omitempty" yaml:"GetFarmManagerResult,omitempty"`
}

// GetFarmers was auto-generated from WSDL.
type GetFarmers struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmers" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
}

// GetFarmersResponse was auto-generated from WSDL.
type GetFarmersResponse struct {
	GetFarmersResult *PersonListResult `xml:"GetFarmersResult,omitempty" json:"GetFarmersResult,omitempty" yaml:"GetFarmersResult,omitempty"`
}

// GetLabelEarTagOrders was auto-generated from WSDL.
type GetLabelEarTagOrders struct {
	XMLName           xml.Name  `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetLabelEarTagOrders" json:"-" yaml:"-"`
	P_ManufacturerKey string    `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int       `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int       `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_SearchDateFrom  DateTime  `xml:"p_SearchDateFrom" json:"p_SearchDateFrom" yaml:"p_SearchDateFrom"`
	P_SearchDateTo    DateTime  `xml:"p_SearchDateTo" json:"p_SearchDateTo" yaml:"p_SearchDateTo"`
	P_ArticleFilter   *IntArray `xml:"p_ArticleFilter" json:"p_ArticleFilter" yaml:"p_ArticleFilter"`
}

// GetLabelEarTagOrdersResponse was auto-generated from WSDL.
type GetLabelEarTagOrdersResponse struct {
	GetLabelEarTagOrdersResult *EarTagOrderResult `xml:"GetLabelEarTagOrdersResult,omitempty" json:"GetLabelEarTagOrdersResult,omitempty" yaml:"GetLabelEarTagOrdersResult,omitempty"`
}

// GetMembershipForOrganisation was auto-generated from WSDL.
type GetMembershipForOrganisation struct {
	XMLName   xml.Name                             `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetMembershipForOrganisation" json:"-" yaml:"-"`
	P_Request *GetMembershipForOrganisationRequest `xml:"p_Request" json:"p_Request" yaml:"p_Request"`
}

// GetMembershipForOrganisationRequest was auto-generated from
// WSDL.
type GetMembershipForOrganisationRequest struct {
	ManufacturerKey string `xml:"ManufacturerKey,omitempty" json:"ManufacturerKey,omitempty" yaml:"ManufacturerKey,omitempty"`
	LCID            int    `xml:"LCID,omitempty" json:"LCID,omitempty" yaml:"LCID,omitempty"`
	TVDNumber       int    `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	AgateNumber     string `xml:"AgateNumber,omitempty" json:"AgateNumber,omitempty" yaml:"AgateNumber,omitempty"`
}

// GetMembershipForOrganisationResponse was auto-generated from
// WSDL.
type GetMembershipForOrganisationResponse struct {
	GetMembershipForOrganisationResult *GetMembershipForOrganisationResult `xml:"GetMembershipForOrganisationResult,omitempty" json:"GetMembershipForOrganisationResult,omitempty" yaml:"GetMembershipForOrganisationResult,omitempty"`
}

// GetMembershipForOrganisationResult was auto-generated from WSDL.
type GetMembershipForOrganisationResult struct {
	Result                 *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	MembershipOrganisation *GenusElements    `xml:"MembershipOrganisation,omitempty" json:"MembershipOrganisation,omitempty" yaml:"MembershipOrganisation,omitempty"`
	BreedingOrganisation   *GenusElements    `xml:"BreedingOrganisation,omitempty" json:"BreedingOrganisation,omitempty" yaml:"BreedingOrganisation,omitempty"`
}

// GetPersonAddress was auto-generated from WSDL.
type GetPersonAddress struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonAddress" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_AgateNumber     string   `xml:"p_AgateNumber,omitempty" json:"p_AgateNumber,omitempty" yaml:"p_AgateNumber,omitempty"`
}

// GetPersonAddressResponse was auto-generated from WSDL.
type GetPersonAddressResponse struct {
	GetPersonAddressResult *PersonAddressResult `xml:"GetPersonAddressResult,omitempty" json:"GetPersonAddressResult,omitempty" yaml:"GetPersonAddressResult,omitempty"`
}

// GetPersonIdentifiers was auto-generated from WSDL.
type GetPersonIdentifiers struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonIdentifiers" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
}

// GetPersonIdentifiersResponse was auto-generated from WSDL.
type GetPersonIdentifiersResponse struct {
	GetPersonIdentifiersResult *PersonIdentifiersResult `xml:"GetPersonIdentifiersResult,omitempty" json:"GetPersonIdentifiersResult,omitempty" yaml:"GetPersonIdentifiersResult,omitempty"`
}

// GetPigArrivalNotificationForBreedingOrganisation was auto-generated
// from WSDL.
type GetPigArrivalNotificationForBreedingOrganisation struct {
	XMLName                            xml.Name                          `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPigArrivalNotificationForBreedingOrganisation" json:"-" yaml:"-"`
	P_SearchSmallAnimalMovementRequest *SearchSmallAnimalMovementRequest `xml:"p_SearchSmallAnimalMovementRequest" json:"p_SearchSmallAnimalMovementRequest" yaml:"p_SearchSmallAnimalMovementRequest"`
}

// GetPigArrivalNotificationForBreedingOrganisationResponse was
// auto-generated from WSDL.
type GetPigArrivalNotificationForBreedingOrganisationResponse struct {
	GetPigArrivalNotificationForBreedingOrganisationResult *PigArrivalNotificationResult `xml:"GetPigArrivalNotificationForBreedingOrganisationResult,omitempty" json:"GetPigArrivalNotificationForBreedingOrganisationResult,omitempty" yaml:"GetPigArrivalNotificationForBreedingOrganisationResult,omitempty"`
}

// GetPoultryBarnInNotificationResult was auto-generated from WSDL.
type GetPoultryBarnInNotificationResult struct {
	Result               *ProcessingResult                                                      `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	ArrivalNotifications *SmallAnimalNotificationDataArrayOfTypeSearchPoultryBarnInNotification `xml:"ArrivalNotifications,omitempty" json:"ArrivalNotifications,omitempty" yaml:"ArrivalNotifications,omitempty"`
}

// GetPoultryBarnInNotifications was auto-generated from WSDL.
type GetPoultryBarnInNotifications struct {
	XMLName                                  xml.Name                                `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPoultryBarnInNotifications" json:"-" yaml:"-"`
	P_SearchPoultryBarnInNotificationRequest *SearchPoultryBarnInNotificationRequest `xml:"p_SearchPoultryBarnInNotificationRequest" json:"p_SearchPoultryBarnInNotificationRequest" yaml:"p_SearchPoultryBarnInNotificationRequest"`
}

// GetPoultryBarnInNotificationsResponse was auto-generated from
// WSDL.
type GetPoultryBarnInNotificationsResponse struct {
	GetPoultryBarnInNotificationsResult *GetPoultryBarnInNotificationResult `xml:"GetPoultryBarnInNotificationsResult,omitempty" json:"GetPoultryBarnInNotificationsResult,omitempty" yaml:"GetPoultryBarnInNotificationsResult,omitempty"`
}

// GetRegisteredGenera was auto-generated from WSDL.
type GetRegisteredGenera struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetRegisteredGenera" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
}

// GetRegisteredGeneraResponse was auto-generated from WSDL.
type GetRegisteredGeneraResponse struct {
	GetRegisteredGeneraResult *GeneraResult `xml:"GetRegisteredGeneraResult,omitempty" json:"GetRegisteredGeneraResult,omitempty" yaml:"GetRegisteredGeneraResult,omitempty"`
}

// GetSmallAnimalSlaughterList was auto-generated from WSDL.
type GetSmallAnimalSlaughterList struct {
	XMLName           xml.Name                     `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetSmallAnimalSlaughterList" json:"-" yaml:"-"`
	P_ManufacturerKey string                       `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                          `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_Request         *SmallAnimalSlaughterRequest `xml:"p_Request" json:"p_Request" yaml:"p_Request"`
}

// GetSmallAnimalSlaughterListResponse was auto-generated from
// WSDL.
type GetSmallAnimalSlaughterListResponse struct {
	GetSmallAnimalSlaughterListResult *SmallAnimalSlaughterResult `xml:"GetSmallAnimalSlaughterListResult,omitempty" json:"GetSmallAnimalSlaughterListResult,omitempty" yaml:"GetSmallAnimalSlaughterListResult,omitempty"`
}

// HusbandryMembership was auto-generated from WSDL.
type HusbandryMembership struct {
	ID_PRG      int    `xml:"ID_PRG,omitempty" json:"ID_PRG,omitempty" yaml:"ID_PRG,omitempty"`
	AgateNumber string `xml:"AgateNumber,omitempty" json:"AgateNumber,omitempty" yaml:"AgateNumber,omitempty"`
	LastName    string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
	FirstName   string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`
	Role        string `xml:"Role,omitempty" json:"Role,omitempty" yaml:"Role,omitempty"`
	Genus       int    `xml:"Genus,omitempty" json:"Genus,omitempty" yaml:"Genus,omitempty"`
}

// HusbandryMembershipArray was auto-generated from WSDL.
type HusbandryMembershipArray struct {
	HusbandryMembershipItem []*HusbandryMembership `xml:"HusbandryMembershipItem,omitempty" json:"HusbandryMembershipItem,omitempty" yaml:"HusbandryMembershipItem,omitempty"`
}

// HusbandryMembershipResult was auto-generated from WSDL.
type HusbandryMembershipResult struct {
	Result      *ProcessingResult         `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	Memberships *HusbandryMembershipArray `xml:"Memberships,omitempty" json:"Memberships,omitempty" yaml:"Memberships,omitempty"`
}

// HusbandryNotificationResult was auto-generated from WSDL.
type HusbandryNotificationResult struct {
	TVDNumber      int               `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	NotificationID int               `xml:"NotificationID,omitempty" json:"NotificationID,omitempty" yaml:"NotificationID,omitempty"`
	Result         *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

// HusbandryNotificationResultArray was auto-generated from WSDL.
type HusbandryNotificationResultArray struct {
	HusbandryNotificationResultItem []*HusbandryNotificationResult `xml:"HusbandryNotificationResultItem,omitempty" json:"HusbandryNotificationResultItem,omitempty" yaml:"HusbandryNotificationResultItem,omitempty"`
}

// HusbandryResult was auto-generated from WSDL.
type HusbandryResult struct {
	TVDNumber string `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	Name      string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Street    string `xml:"Street,omitempty" json:"Street,omitempty" yaml:"Street,omitempty"`
	PostCode  string `xml:"PostCode,omitempty" json:"PostCode,omitempty" yaml:"PostCode,omitempty"`
	City      string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
}

// IntArray was auto-generated from WSDL.
type IntArray struct {
	IntItem []int `xml:"IntItem,omitempty" json:"IntItem,omitempty" yaml:"IntItem,omitempty"`
}

// NewEarTagOrderResult was auto-generated from WSDL.
type NewEarTagOrderResult struct {
	Result      *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	EarTagOrder *EarTagOrderData  `xml:"EarTagOrder,omitempty" json:"EarTagOrder,omitempty" yaml:"EarTagOrder,omitempty"`
}

// NotificationResult was auto-generated from WSDL.
type NotificationResult struct {
	Result         *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	NotificationID int               `xml:"NotificationID,omitempty" json:"NotificationID,omitempty" yaml:"NotificationID,omitempty"`
}

// PersonAddressResult was auto-generated from WSDL.
type PersonAddressResult struct {
	Result      *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	Title       string            `xml:"Title,omitempty" json:"Title,omitempty" yaml:"Title,omitempty"`
	PostAddress *PersonResult     `xml:"PostAddress,omitempty" json:"PostAddress,omitempty" yaml:"PostAddress,omitempty"`
}

// PersonDataArray was auto-generated from WSDL.
type PersonDataArray struct {
	PersonDataItem []*PersonResult `xml:"PersonDataItem,omitempty" json:"PersonDataItem,omitempty" yaml:"PersonDataItem,omitempty"`
}

// PersonIdentifiersResult was auto-generated from WSDL.
type PersonIdentifiersResult struct {
	Result          *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	CantonPersonKey string            `xml:"CantonPersonKey,omitempty" json:"CantonPersonKey,omitempty" yaml:"CantonPersonKey,omitempty"`
	ID_AGIS         int               `xml:"ID_AGIS,omitempty" json:"ID_AGIS,omitempty" yaml:"ID_AGIS,omitempty"`
}

// PersonListResult was auto-generated from WSDL.
type PersonListResult struct {
	Result          *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	PersonDataItems *PersonDataArray  `xml:"PersonDataItems,omitempty" json:"PersonDataItems,omitempty" yaml:"PersonDataItems,omitempty"`
}

// PersonResult was auto-generated from WSDL.
type PersonResult struct {
	AgateNumber           string       `xml:"AgateNumber,omitempty" json:"AgateNumber,omitempty" yaml:"AgateNumber,omitempty"`
	LastName              string       `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
	FirstName             string       `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`
	Street                string       `xml:"Street,omitempty" json:"Street,omitempty" yaml:"Street,omitempty"`
	PostCode              string       `xml:"PostCode,omitempty" json:"PostCode,omitempty" yaml:"PostCode,omitempty"`
	City                  string       `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	EmailAddress          string       `xml:"EmailAddress,omitempty" json:"EmailAddress,omitempty" yaml:"EmailAddress,omitempty"`
	PhoneNumbers          *StringArray `xml:"PhoneNumbers,omitempty" json:"PhoneNumbers,omitempty" yaml:"PhoneNumbers,omitempty"`
	PreferredLanguageLCID int          `xml:"PreferredLanguageLCID,omitempty" json:"PreferredLanguageLCID,omitempty" yaml:"PreferredLanguageLCID,omitempty"`
}

// PigArrivalNotification was auto-generated from WSDL.
type PigArrivalNotification struct {
	TVDNumber       int                                  `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	SourceTVDNumber int                                  `xml:"SourceTVDNumber,omitempty" json:"SourceTVDNumber,omitempty" yaml:"SourceTVDNumber,omitempty"`
	Amount          int                                  `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
	DeletionDate    DateTime                             `xml:"DeletionDate,omitempty" json:"DeletionDate,omitempty" yaml:"DeletionDate,omitempty"`
	CreationDate    DateTime                             `xml:"CreationDate,omitempty" json:"CreationDate,omitempty" yaml:"CreationDate,omitempty"`
	EventDate       DateTime                             `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	PigCategory     *TranslatedEnumTypeOfEnumPigCategory `xml:"PigCategory,omitempty" json:"PigCategory,omitempty" yaml:"PigCategory,omitempty"`
}

// PigArrivalNotificationResult was auto-generated from WSDL.
type PigArrivalNotificationResult struct {
	Result               *ProcessingResult                                             `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	ArrivalNotifications *SmallAnimalNotificationDataArrayOfTypePigArrivalNotification `xml:"ArrivalNotifications,omitempty" json:"ArrivalNotifications,omitempty" yaml:"ArrivalNotifications,omitempty"`
}

// PoultryBarnInNotification was auto-generated from WSDL.
type PoultryBarnInNotification struct {
	Amount                    int                    `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
	EventDate                 DateTime               `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	SourceTvdNumber           int                    `xml:"SourceTvdNumber,omitempty" json:"SourceTvdNumber,omitempty" yaml:"SourceTvdNumber,omitempty"`
	PoultryUsageReason        EnumPoultryUsageReason `xml:"PoultryUsageReason,omitempty" json:"PoultryUsageReason,omitempty" yaml:"PoultryUsageReason,omitempty"`
	AllowSeveralNotifications bool                   `xml:"AllowSeveralNotifications,omitempty" json:"AllowSeveralNotifications,omitempty" yaml:"AllowSeveralNotifications,omitempty"`
}

// ProcessingResult was auto-generated from WSDL.
type ProcessingResult struct {
	Code        int    `xml:"Code,omitempty" json:"Code,omitempty" yaml:"Code,omitempty"`
	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Status      int    `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
}

// ReplacementEarTagOrderDataItem was auto-generated from WSDL.
type ReplacementEarTagOrderDataItem struct {
	NotificationID  int                                  `xml:"NotificationID,omitempty" json:"NotificationID,omitempty" yaml:"NotificationID,omitempty"`
	CombiArticle    int                                  `xml:"CombiArticle,omitempty" json:"CombiArticle,omitempty" yaml:"CombiArticle,omitempty"`
	OrderStatus     *TranslatedEnumTypeOfEnumOrderStatus `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty" yaml:"OrderStatus,omitempty"`
	OrderStatusDate DateTime                             `xml:"OrderStatusDate,omitempty" json:"OrderStatusDate,omitempty" yaml:"OrderStatusDate,omitempty"`
	EarTagNumber    string                               `xml:"EarTagNumber,omitempty" json:"EarTagNumber,omitempty" yaml:"EarTagNumber,omitempty"`
	IsExpress       bool                                 `xml:"IsExpress,omitempty" json:"IsExpress,omitempty" yaml:"IsExpress,omitempty"`
	Text            string                               `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`
}

// ReplacementEarTagOrdersResult was auto-generated from WSDL.
type ReplacementEarTagOrdersResult struct {
	Result        *ProcessingResult                      `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	ResultDetails *ArrayOfReplacementEarTagOrderDataItem `xml:"ResultDetails,omitempty" json:"ResultDetails,omitempty" yaml:"ResultDetails,omitempty"`
}

// SearchEquid was auto-generated from WSDL.
type SearchEquid struct {
	XMLName              xml.Name            `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SearchEquid" json:"-" yaml:"-"`
	P_ManufacturerKey    string              `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID               int                 `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidSearchRequest *EquidSearchRequest `xml:"p_EquidSearchRequest" json:"p_EquidSearchRequest" yaml:"p_EquidSearchRequest"`
}

// SearchEquidResponse was auto-generated from WSDL.
type SearchEquidResponse struct {
	SearchEquidResult *EquidSearchResult `xml:"SearchEquidResult,omitempty" json:"SearchEquidResult,omitempty" yaml:"SearchEquidResult,omitempty"`
}

// SearchPoultryBarnInNotification was auto-generated from WSDL.
type SearchPoultryBarnInNotification struct {
	Amount                     int                                         `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
	EventDate                  DateTime                                    `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	SourceTvdNumber            int                                         `xml:"SourceTvdNumber,omitempty" json:"SourceTvdNumber,omitempty" yaml:"SourceTvdNumber,omitempty"`
	PoultryUsageReason         *TranslatedEnumTypeOfEnumPoultryUsageReason `xml:"PoultryUsageReason,omitempty" json:"PoultryUsageReason,omitempty" yaml:"PoultryUsageReason,omitempty"`
	NotificationDate           DateTime                                    `xml:"NotificationDate,omitempty" json:"NotificationDate,omitempty" yaml:"NotificationDate,omitempty"`
	HerdenIdentificationNumber string                                      `xml:"HerdenIdentificationNumber,omitempty" json:"HerdenIdentificationNumber,omitempty" yaml:"HerdenIdentificationNumber,omitempty"`
	CreatedBy                  string                                      `xml:"CreatedBy,omitempty" json:"CreatedBy,omitempty" yaml:"CreatedBy,omitempty"`
	DeletedBy                  string                                      `xml:"DeletedBy,omitempty" json:"DeletedBy,omitempty" yaml:"DeletedBy,omitempty"`
	IsDeleted                  bool                                        `xml:"IsDeleted,omitempty" json:"IsDeleted,omitempty" yaml:"IsDeleted,omitempty"`
}

// SearchPoultryBarnInNotificationRequest was auto-generated from
// WSDL.
type SearchPoultryBarnInNotificationRequest struct {
	ManufacturerKey string   `xml:"ManufacturerKey,omitempty" json:"ManufacturerKey,omitempty" yaml:"ManufacturerKey,omitempty"`
	LCID            int      `xml:"LCID,omitempty" json:"LCID,omitempty" yaml:"LCID,omitempty"`
	TVDNumber       int      `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	FromDate        DateTime `xml:"FromDate,omitempty" json:"FromDate,omitempty" yaml:"FromDate,omitempty"`
	ToDate          DateTime `xml:"ToDate,omitempty" json:"ToDate,omitempty" yaml:"ToDate,omitempty"`
}

// SearchSmallAnimalMovementRequest was auto-generated from WSDL.
type SearchSmallAnimalMovementRequest struct {
	ManufacturerKey string   `xml:"ManufacturerKey,omitempty" json:"ManufacturerKey,omitempty" yaml:"ManufacturerKey,omitempty"`
	LCID            int      `xml:"LCID,omitempty" json:"LCID,omitempty" yaml:"LCID,omitempty"`
	FromDate        DateTime `xml:"FromDate,omitempty" json:"FromDate,omitempty" yaml:"FromDate,omitempty"`
	ToDate          DateTime `xml:"ToDate,omitempty" json:"ToDate,omitempty" yaml:"ToDate,omitempty"`
}

// SmallAnimalNotificationDataArrayOfTypePigArrivalNotification
// was auto-generated from WSDL.
type SmallAnimalNotificationDataArrayOfTypePigArrivalNotification struct {
	SmallAnimalNotificationDataArrayItem []*PigArrivalNotification `xml:"SmallAnimalNotificationDataArrayItem,omitempty" json:"SmallAnimalNotificationDataArrayItem,omitempty" yaml:"SmallAnimalNotificationDataArrayItem,omitempty"`
}

// SmallAnimalNotificationDataArrayOfTypeSearchPoultryBarnInNotification
// was auto-generated from WSDL.
type SmallAnimalNotificationDataArrayOfTypeSearchPoultryBarnInNotification struct {
	SmallAnimalNotificationDataArrayItem []*SearchPoultryBarnInNotification `xml:"SmallAnimalNotificationDataArrayItem,omitempty" json:"SmallAnimalNotificationDataArrayItem,omitempty" yaml:"SmallAnimalNotificationDataArrayItem,omitempty"`
}

// SmallAnimalSlaughterListData was auto-generated from WSDL.
type SmallAnimalSlaughterListData struct {
	EventDate              DateTime  `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	NotificationDate       DateTime  `xml:"NotificationDate,omitempty" json:"NotificationDate,omitempty" yaml:"NotificationDate,omitempty"`
	Amount                 int       `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
	Genus                  EnumGenus `xml:"Genus,omitempty" json:"Genus,omitempty" yaml:"Genus,omitempty"`
	SourceTVDNumber        string    `xml:"SourceTVDNumber,omitempty" json:"SourceTVDNumber,omitempty" yaml:"SourceTVDNumber,omitempty"`
	SlaughterInitTVDNumber string    `xml:"SlaughterInitTVDNumber,omitempty" json:"SlaughterInitTVDNumber,omitempty" yaml:"SlaughterInitTVDNumber,omitempty"`
	CreatorBZVAccount      string    `xml:"CreatorBZVAccount,omitempty" json:"CreatorBZVAccount,omitempty" yaml:"CreatorBZVAccount,omitempty"`
	DeleterBZVAccount      string    `xml:"DeleterBZVAccount,omitempty" json:"DeleterBZVAccount,omitempty" yaml:"DeleterBZVAccount,omitempty"`
}

// SmallAnimalSlaughterRequest was auto-generated from WSDL.
type SmallAnimalSlaughterRequest struct {
	FromDate           DateTime  `xml:"FromDate,omitempty" json:"FromDate,omitempty" yaml:"FromDate,omitempty"`
	ToDate             DateTime  `xml:"ToDate,omitempty" json:"ToDate,omitempty" yaml:"ToDate,omitempty"`
	SlaughterTVDNumber int       `xml:"SlaughterTVDNumber,omitempty" json:"SlaughterTVDNumber,omitempty" yaml:"SlaughterTVDNumber,omitempty"`
	Genus              EnumGenus `xml:"Genus,omitempty" json:"Genus,omitempty" yaml:"Genus,omitempty"`
}

// SmallAnimalSlaughterResult was auto-generated from WSDL.
type SmallAnimalSlaughterResult struct {
	Result        *ProcessingResult                    `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	SlaughterList *ArrayOfSmallAnimalSlaughterListData `xml:"SlaughterList,omitempty" json:"SlaughterList,omitempty" yaml:"SlaughterList,omitempty"`
}

// StringArray was auto-generated from WSDL.
type StringArray struct {
	StringItem []string `xml:"StringItem,omitempty" json:"StringItem,omitempty" yaml:"StringItem,omitempty"`
}

// TranslatedEnumTypeOfEnumAnimalHusbandryType was auto-generated
// from WSDL.
type TranslatedEnumTypeOfEnumAnimalHusbandryType struct {
	EnumValue            EnumAnimalHusbandryType `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string                  `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumAnimalHusbandryTypeOfUse was auto-generated
// from WSDL.
type TranslatedEnumTypeOfEnumAnimalHusbandryTypeOfUse struct {
	EnumValue            EnumAnimalHusbandryTypeOfUse `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string                       `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumArea was auto-generated from WSDL.
type TranslatedEnumTypeOfEnumArea struct {
	EnumValue            EnumArea `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string   `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumCattleRace was auto-generated from WSDL.
type TranslatedEnumTypeOfEnumCattleRace struct {
	EnumValue            EnumCattleRace `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string         `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumEquidBreed was auto-generated from WSDL.
type TranslatedEnumTypeOfEnumEquidBreed struct {
	EnumValue            EnumEquidBreed `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string         `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumEquidNotificationState was auto-generated
// from WSDL.
type TranslatedEnumTypeOfEnumEquidNotificationState struct {
	EnumValue            EnumEquidNotificationState `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string                     `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumEquidNotificationType was auto-generated
// from WSDL.
type TranslatedEnumTypeOfEnumEquidNotificationType struct {
	EnumValue            EnumEquidNotificationType `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string                    `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumEquidPassIssuingState was auto-generated
// from WSDL.
type TranslatedEnumTypeOfEnumEquidPassIssuingState struct {
	EnumValue            EnumEquidPassIssuingState `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string                    `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumEquidPassType was auto-generated from
// WSDL.
type TranslatedEnumTypeOfEnumEquidPassType struct {
	EnumValue            EnumEquidPassType `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string            `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumEquidSpecies was auto-generated from
// WSDL.
type TranslatedEnumTypeOfEnumEquidSpecies struct {
	EnumValue            EnumEquidSpecies `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string           `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumEquidTypeOfUse was auto-generated from
// WSDL.
type TranslatedEnumTypeOfEnumEquidTypeOfUse struct {
	EnumValue            EnumEquidTypeOfUse `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string             `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumEquidWithersClass was auto-generated
// from WSDL.
type TranslatedEnumTypeOfEnumEquidWithersClass struct {
	EnumValue            EnumEquidWithersClass `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string                `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumGender was auto-generated from WSDL.
type TranslatedEnumTypeOfEnumGender struct {
	EnumValue            EnumGender `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string     `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumGenus was auto-generated from WSDL.
type TranslatedEnumTypeOfEnumGenus struct {
	EnumValue            EnumGenus `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string    `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumOrderStatus was auto-generated from
// WSDL.
type TranslatedEnumTypeOfEnumOrderStatus struct {
	EnumValue            EnumOrderStatus `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string          `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumPigCategory was auto-generated from
// WSDL.
type TranslatedEnumTypeOfEnumPigCategory struct {
	EnumValue            EnumPigCategory `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string          `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumPoultryUsageReason was auto-generated
// from WSDL.
type TranslatedEnumTypeOfEnumPoultryUsageReason struct {
	EnumValue            EnumPoultryUsageReason `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string                 `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// TranslatedEnumTypeOfEnumZone was auto-generated from WSDL.
type TranslatedEnumTypeOfEnumZone struct {
	EnumValue            EnumZone `xml:"EnumValue,omitempty" json:"EnumValue,omitempty" yaml:"EnumValue,omitempty"`
	RequestedTranslation string   `xml:"RequestedTranslation,omitempty" json:"RequestedTranslation,omitempty" yaml:"RequestedTranslation,omitempty"`
}

// Version was auto-generated from WSDL.
type Version struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 Version" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
}

// VersionResponse was auto-generated from WSDL.
type VersionResponse struct {
	VersionResult string `xml:"VersionResult,omitempty" json:"VersionResult,omitempty" yaml:"VersionResult,omitempty"`
}

// WorkingFocus was auto-generated from WSDL.
type WorkingFocus struct {
	WorkingFocusType int    `xml:"WorkingFocusType,omitempty" json:"WorkingFocusType,omitempty" yaml:"WorkingFocusType,omitempty"`
	TVDNumber        int    `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	MandateGiver     string `xml:"MandateGiver,omitempty" json:"MandateGiver,omitempty" yaml:"MandateGiver,omitempty"`
}

// WorkingFocusArray was auto-generated from WSDL.
type WorkingFocusArray struct {
	WorkingFocusItem []*WorkingFocus `xml:"WorkingFocusItem,omitempty" json:"WorkingFocusItem,omitempty" yaml:"WorkingFocusItem,omitempty"`
}

// WriteAnimalClassificationNotification was auto-generated from
// WSDL.
type WriteAnimalClassificationNotification struct {
	XMLName                    xml.Name                  `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotification" json:"-" yaml:"-"`
	P_ManufacturerKey          string                    `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                     int                       `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber                int                       `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_AnimalClassificationData *AnimalClassificationData `xml:"p_AnimalClassificationData" json:"p_AnimalClassificationData" yaml:"p_AnimalClassificationData"`
}

// WriteAnimalClassificationNotificationResponse was auto-generated
// from WSDL.
type WriteAnimalClassificationNotificationResponse struct {
	WriteAnimalClassificationNotificationResult *NotificationResult `xml:"WriteAnimalClassificationNotificationResult,omitempty" json:"WriteAnimalClassificationNotificationResult,omitempty" yaml:"WriteAnimalClassificationNotificationResult,omitempty"`
}

// WriteAnimalClassificationNotificationV2 was auto-generated from
// WSDL.
type WriteAnimalClassificationNotificationV2 struct {
	XMLName                    xml.Name                    `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotificationV2" json:"-" yaml:"-"`
	P_ManufacturerKey          string                      `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                     int                         `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber                int                         `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_AnimalClassificationData *AnimalClassificationDataV2 `xml:"p_AnimalClassificationData" json:"p_AnimalClassificationData" yaml:"p_AnimalClassificationData"`
}

// WriteAnimalClassificationNotificationV2Response was auto-generated
// from WSDL.
type WriteAnimalClassificationNotificationV2Response struct {
	WriteAnimalClassificationNotificationV2Result *NotificationResult `xml:"WriteAnimalClassificationNotificationV2Result,omitempty" json:"WriteAnimalClassificationNotificationV2Result,omitempty" yaml:"WriteAnimalClassificationNotificationV2Result,omitempty"`
}

// WriteCattleArrivalBatchNotification was auto-generated from
// WSDL.
type WriteCattleArrivalBatchNotification struct {
	XMLName           xml.Name                `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleArrivalBatchNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string                  `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                     `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int                     `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_ArrivalData     *CattleArrivalDataArray `xml:"p_ArrivalData" json:"p_ArrivalData" yaml:"p_ArrivalData"`
}

// WriteCattleArrivalBatchNotificationResponse was auto-generated
// from WSDL.
type WriteCattleArrivalBatchNotificationResponse struct {
	WriteCattleArrivalBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleArrivalBatchNotificationResult,omitempty" json:"WriteCattleArrivalBatchNotificationResult,omitempty" yaml:"WriteCattleArrivalBatchNotificationResult,omitempty"`
}

// WriteCattleBatchNotificationResult was auto-generated from WSDL.
type WriteCattleBatchNotificationResult struct {
	Result        *ProcessingResult              `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
	Resultdetails *CattleNotificationResultArray `xml:"Resultdetails,omitempty" json:"Resultdetails,omitempty" yaml:"Resultdetails,omitempty"`
}

// WriteCattleBirthNotification was auto-generated from WSDL.
type WriteCattleBirthNotification struct {
	XMLName           xml.Name         `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleBirthNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string           `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int              `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int              `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_BirthData       *CattleBirthData `xml:"p_BirthData" json:"p_BirthData" yaml:"p_BirthData"`
}

// WriteCattleBirthNotificationResponse was auto-generated from
// WSDL.
type WriteCattleBirthNotificationResponse struct {
	WriteCattleBirthNotificationResult *CattleNotificationResult `xml:"WriteCattleBirthNotificationResult,omitempty" json:"WriteCattleBirthNotificationResult,omitempty" yaml:"WriteCattleBirthNotificationResult,omitempty"`
}

// WriteCattleChangeNameNotification was auto-generated from WSDL.
type WriteCattleChangeNameNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleChangeNameNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_EarTagNumber    string   `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
	P_Name            string   `xml:"p_Name,omitempty" json:"p_Name,omitempty" yaml:"p_Name,omitempty"`
}

// WriteCattleChangeNameNotificationResponse was auto-generated
// from WSDL.
type WriteCattleChangeNameNotificationResponse struct {
	WriteCattleChangeNameNotificationResult *NotificationResult `xml:"WriteCattleChangeNameNotificationResult,omitempty" json:"WriteCattleChangeNameNotificationResult,omitempty" yaml:"WriteCattleChangeNameNotificationResult,omitempty"`
}

// WriteCattleDaystayBatchNotification was auto-generated from
// WSDL.
type WriteCattleDaystayBatchNotification struct {
	XMLName           xml.Name                `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDaystayBatchNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string                  `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                     `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int                     `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_DaystayData     *CattleDaystayDataArray `xml:"p_DaystayData" json:"p_DaystayData" yaml:"p_DaystayData"`
}

// WriteCattleDaystayBatchNotificationResponse was auto-generated
// from WSDL.
type WriteCattleDaystayBatchNotificationResponse struct {
	WriteCattleDaystayBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleDaystayBatchNotificationResult,omitempty" json:"WriteCattleDaystayBatchNotificationResult,omitempty" yaml:"WriteCattleDaystayBatchNotificationResult,omitempty"`
}

// WriteCattleDeathBirthNotification was auto-generated from WSDL.
type WriteCattleDeathBirthNotification struct {
	XMLName           xml.Name              `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeathBirthNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string                `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                   `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int                   `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_DeathBirthData  *CattleDeathBirthData `xml:"p_DeathBirthData" json:"p_DeathBirthData" yaml:"p_DeathBirthData"`
}

// WriteCattleDeathBirthNotificationResponse was auto-generated
// from WSDL.
type WriteCattleDeathBirthNotificationResponse struct {
	WriteCattleDeathBirthNotificationResult *CattleDeathBirthNotificationResult `xml:"WriteCattleDeathBirthNotificationResult,omitempty" json:"WriteCattleDeathBirthNotificationResult,omitempty" yaml:"WriteCattleDeathBirthNotificationResult,omitempty"`
}

// WriteCattleDeceasedNotification was auto-generated from WSDL.
type WriteCattleDeceasedNotification struct {
	XMLName                 xml.Name                   `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeceasedNotification" json:"-" yaml:"-"`
	P_CattleDeceasedRequest *CattleNotificationRequest `xml:"p_CattleDeceasedRequest" json:"p_CattleDeceasedRequest" yaml:"p_CattleDeceasedRequest"`
}

// WriteCattleDeceasedNotificationResponse was auto-generated from
// WSDL.
type WriteCattleDeceasedNotificationResponse struct {
	WriteCattleDeceasedNotificationResult *CattleNotificationResult `xml:"WriteCattleDeceasedNotificationResult,omitempty" json:"WriteCattleDeceasedNotificationResult,omitempty" yaml:"WriteCattleDeceasedNotificationResult,omitempty"`
}

// WriteCattleDeformationNotification was auto-generated from WSDL.
type WriteCattleDeformationNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeformationNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_EarTagNumber    string   `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
	P_DeformationType int      `xml:"p_DeformationType" json:"p_DeformationType" yaml:"p_DeformationType"`
}

// WriteCattleDeformationNotificationResponse was auto-generated
// from WSDL.
type WriteCattleDeformationNotificationResponse struct {
	WriteCattleDeformationNotificationResult *CattleNotificationResult `xml:"WriteCattleDeformationNotificationResult,omitempty" json:"WriteCattleDeformationNotificationResult,omitempty" yaml:"WriteCattleDeformationNotificationResult,omitempty"`
}

// WriteCattleExportNotification was auto-generated from WSDL.
type WriteCattleExportNotification struct {
	XMLName               xml.Name                              `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleExportNotification" json:"-" yaml:"-"`
	P_CattleExportRequest *CattleNotificationWithCountryRequest `xml:"p_CattleExportRequest" json:"p_CattleExportRequest" yaml:"p_CattleExportRequest"`
}

// WriteCattleExportNotificationResponse was auto-generated from
// WSDL.
type WriteCattleExportNotificationResponse struct {
	WriteCattleExportNotificationResult *CattleNotificationResult `xml:"WriteCattleExportNotificationResult,omitempty" json:"WriteCattleExportNotificationResult,omitempty" yaml:"WriteCattleExportNotificationResult,omitempty"`
}

// WriteCattleImportAfterExportNotification was auto-generated
// from WSDL.
type WriteCattleImportAfterExportNotification struct {
	XMLName                          xml.Name                              `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleImportAfterExportNotification" json:"-" yaml:"-"`
	P_CattleImportAfterExportRequest *CattleNotificationWithCountryRequest `xml:"p_CattleImportAfterExportRequest" json:"p_CattleImportAfterExportRequest" yaml:"p_CattleImportAfterExportRequest"`
}

// WriteCattleImportAfterExportNotificationResponse was auto-generated
// from WSDL.
type WriteCattleImportAfterExportNotificationResponse struct {
	WriteCattleImportAfterExportNotificationResult *CattleNotificationResult `xml:"WriteCattleImportAfterExportNotificationResult,omitempty" json:"WriteCattleImportAfterExportNotificationResult,omitempty" yaml:"WriteCattleImportAfterExportNotificationResult,omitempty"`
}

// WriteCattleLeavingBatchNotification was auto-generated from
// WSDL.
type WriteCattleLeavingBatchNotification struct {
	XMLName           xml.Name                `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleLeavingBatchNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string                  `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                     `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int                     `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_LeavingData     *CattleLeavingDataArray `xml:"p_LeavingData" json:"p_LeavingData" yaml:"p_LeavingData"`
}

// WriteCattleLeavingBatchNotificationResponse was auto-generated
// from WSDL.
type WriteCattleLeavingBatchNotificationResponse struct {
	WriteCattleLeavingBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleLeavingBatchNotificationResult,omitempty" json:"WriteCattleLeavingBatchNotificationResult,omitempty" yaml:"WriteCattleLeavingBatchNotificationResult,omitempty"`
}

// WriteCattlePassportOrders was auto-generated from WSDL.
type WriteCattlePassportOrders struct {
	XMLName            xml.Name     `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattlePassportOrders" json:"-" yaml:"-"`
	P_ManufacturerKey  string       `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID             int          `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_PassportLanguage string       `xml:"p_PassportLanguage,omitempty" json:"p_PassportLanguage,omitempty" yaml:"p_PassportLanguage,omitempty"`
	P_IssueDate        DateTime     `xml:"p_IssueDate" json:"p_IssueDate" yaml:"p_IssueDate"`
	P_EarTagNumbers    *StringArray `xml:"p_EarTagNumbers" json:"p_EarTagNumbers" yaml:"p_EarTagNumbers"`
}

// WriteCattlePassportOrdersResponse was auto-generated from WSDL.
type WriteCattlePassportOrdersResponse struct {
	WriteCattlePassportOrdersResult *CattlePassportResult `xml:"WriteCattlePassportOrdersResult,omitempty" json:"WriteCattlePassportOrdersResult,omitempty" yaml:"WriteCattlePassportOrdersResult,omitempty"`
}

// WriteCattleSlaughterBatchNotification was auto-generated from
// WSDL.
type WriteCattleSlaughterBatchNotification struct {
	XMLName           xml.Name                `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string                  `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                     `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int                     `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_SlaughterData   *CattleArrivalDataArray `xml:"p_SlaughterData" json:"p_SlaughterData" yaml:"p_SlaughterData"`
}

// WriteCattleSlaughterBatchNotificationResponse was auto-generated
// from WSDL.
type WriteCattleSlaughterBatchNotificationResponse struct {
	WriteCattleSlaughterBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleSlaughterBatchNotificationResult,omitempty" json:"WriteCattleSlaughterBatchNotificationResult,omitempty" yaml:"WriteCattleSlaughterBatchNotificationResult,omitempty"`
}

// WriteCattleSlaughterBatchNotificationV2 was auto-generated from
// WSDL.
type WriteCattleSlaughterBatchNotificationV2 struct {
	XMLName           xml.Name                  `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotificationV2" json:"-" yaml:"-"`
	P_ManufacturerKey string                    `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int                       `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int                       `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_SlaughterData   *CattleSlaughterDataArray `xml:"p_SlaughterData" json:"p_SlaughterData" yaml:"p_SlaughterData"`
}

// WriteCattleSlaughterBatchNotificationV2Response was auto-generated
// from WSDL.
type WriteCattleSlaughterBatchNotificationV2Response struct {
	WriteCattleSlaughterBatchNotificationV2Result *WriteCattleBatchNotificationResult `xml:"WriteCattleSlaughterBatchNotificationV2Result,omitempty" json:"WriteCattleSlaughterBatchNotificationV2Result,omitempty" yaml:"WriteCattleSlaughterBatchNotificationV2Result,omitempty"`
}

// WriteCattleTypeOfUseNotification was auto-generated from WSDL.
type WriteCattleTypeOfUseNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleTypeOfUseNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_EarTagNumber    string   `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
	P_CattleTypeOfUse int      `xml:"p_CattleTypeOfUse" json:"p_CattleTypeOfUse" yaml:"p_CattleTypeOfUse"`
	P_EventDate       DateTime `xml:"p_EventDate" json:"p_EventDate" yaml:"p_EventDate"`
}

// WriteCattleTypeOfUseNotificationResponse was auto-generated
// from WSDL.
type WriteCattleTypeOfUseNotificationResponse struct {
	WriteCattleTypeOfUseNotificationResult *NotificationResult `xml:"WriteCattleTypeOfUseNotificationResult,omitempty" json:"WriteCattleTypeOfUseNotificationResult,omitempty" yaml:"WriteCattleTypeOfUseNotificationResult,omitempty"`
}

// WriteCattleYardSlaughterNotification was auto-generated from
// WSDL.
type WriteCattleYardSlaughterNotification struct {
	XMLName                        xml.Name                   `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleYardSlaughterNotification" json:"-" yaml:"-"`
	P_CattleYardSlaughteredRequest *CattleNotificationRequest `xml:"p_CattleYardSlaughteredRequest" json:"p_CattleYardSlaughteredRequest" yaml:"p_CattleYardSlaughteredRequest"`
}

// WriteCattleYardSlaughterNotificationResponse was auto-generated
// from WSDL.
type WriteCattleYardSlaughterNotificationResponse struct {
	WriteCattleYardSlaughterNotificationResult *CattleNotificationResult `xml:"WriteCattleYardSlaughterNotificationResult,omitempty" json:"WriteCattleYardSlaughterNotificationResult,omitempty" yaml:"WriteCattleYardSlaughterNotificationResult,omitempty"`
}

// WriteEquidAcquireOwnershipNotification was auto-generated from
// WSDL.
type WriteEquidAcquireOwnershipNotification struct {
	XMLName                        xml.Name                      `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidAcquireOwnershipNotification" json:"-" yaml:"-"`
	P_ManufacturerKey              string                        `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                         int                           `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidAcquireOwnershipRequest *EquidAcquireOwnershipRequest `xml:"p_EquidAcquireOwnershipRequest" json:"p_EquidAcquireOwnershipRequest" yaml:"p_EquidAcquireOwnershipRequest"`
}

// WriteEquidAcquireOwnershipNotificationResponse was auto-generated
// from WSDL.
type WriteEquidAcquireOwnershipNotificationResponse struct {
	WriteEquidAcquireOwnershipNotificationResult *WriteNotificationResult `xml:"WriteEquidAcquireOwnershipNotificationResult,omitempty" json:"WriteEquidAcquireOwnershipNotificationResult,omitempty" yaml:"WriteEquidAcquireOwnershipNotificationResult,omitempty"`
}

// WriteEquidBasicDataNotification was auto-generated from WSDL.
type WriteEquidBasicDataNotification struct {
	XMLName                      xml.Name               `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidBasicDataNotification" json:"-" yaml:"-"`
	P_ManufacturerKey            string                 `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                       int                    `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WriteEquidBasicDataRequest *EquidBasicDataRequest `xml:"p_WriteEquidBasicDataRequest" json:"p_WriteEquidBasicDataRequest" yaml:"p_WriteEquidBasicDataRequest"`
}

// WriteEquidBasicDataNotificationResponse was auto-generated from
// WSDL.
type WriteEquidBasicDataNotificationResponse struct {
	WriteEquidBasicDataNotificationResult *WriteNotificationResult `xml:"WriteEquidBasicDataNotificationResult,omitempty" json:"WriteEquidBasicDataNotificationResult,omitempty" yaml:"WriteEquidBasicDataNotificationResult,omitempty"`
}

// WriteEquidBirthNotification was auto-generated from WSDL.
type WriteEquidBirthNotification struct {
	XMLName             xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidBirthNotification" json:"-" yaml:"-"`
	P_ManufacturerKey   string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID              int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidBirthRequest *EquidBirthRequest `xml:"p_EquidBirthRequest" json:"p_EquidBirthRequest" yaml:"p_EquidBirthRequest"`
}

// WriteEquidBirthNotificationResponse was auto-generated from
// WSDL.
type WriteEquidBirthNotificationResponse struct {
	WriteEquidBirthNotificationResult *EquidInitialResult `xml:"WriteEquidBirthNotificationResult,omitempty" json:"WriteEquidBirthNotificationResult,omitempty" yaml:"WriteEquidBirthNotificationResult,omitempty"`
}

// WriteEquidCastrationNotification was auto-generated from WSDL.
type WriteEquidCastrationNotification struct {
	XMLName                  xml.Name                `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidCastrationNotification" json:"-" yaml:"-"`
	P_ManufacturerKey        string                  `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                   int                     `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidCastrationRequest *EquidCastrationRequest `xml:"p_EquidCastrationRequest" json:"p_EquidCastrationRequest" yaml:"p_EquidCastrationRequest"`
}

// WriteEquidCastrationNotificationResponse was auto-generated
// from WSDL.
type WriteEquidCastrationNotificationResponse struct {
	WriteEquidCastrationNotificationResult *WriteNotificationResult `xml:"WriteEquidCastrationNotificationResult,omitempty" json:"WriteEquidCastrationNotificationResult,omitempty" yaml:"WriteEquidCastrationNotificationResult,omitempty"`
}

// WriteEquidCedeOwnershipNotification was auto-generated from
// WSDL.
type WriteEquidCedeOwnershipNotification struct {
	XMLName                     xml.Name                   `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidCedeOwnershipNotification" json:"-" yaml:"-"`
	P_ManufacturerKey           string                     `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                      int                        `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidCedeOwnershipRequest *EquidCedeOwnershipRequest `xml:"p_EquidCedeOwnershipRequest" json:"p_EquidCedeOwnershipRequest" yaml:"p_EquidCedeOwnershipRequest"`
}

// WriteEquidCedeOwnershipNotificationResponse was auto-generated
// from WSDL.
type WriteEquidCedeOwnershipNotificationResponse struct {
	WriteEquidCedeOwnershipNotificationResult *WriteNotificationResult `xml:"WriteEquidCedeOwnershipNotificationResult,omitempty" json:"WriteEquidCedeOwnershipNotificationResult,omitempty" yaml:"WriteEquidCedeOwnershipNotificationResult,omitempty"`
}

// WriteEquidDeathNotification was auto-generated from WSDL.
type WriteEquidDeathNotification struct {
	XMLName             xml.Name           `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidDeathNotification" json:"-" yaml:"-"`
	P_ManufacturerKey   string             `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID              int                `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidDeathRequest *EquidDeathRequest `xml:"p_EquidDeathRequest" json:"p_EquidDeathRequest" yaml:"p_EquidDeathRequest"`
}

// WriteEquidDeathNotificationResponse was auto-generated from
// WSDL.
type WriteEquidDeathNotificationResponse struct {
	WriteEquidDeathNotificationResult *WriteNotificationResult `xml:"WriteEquidDeathNotificationResult,omitempty" json:"WriteEquidDeathNotificationResult,omitempty" yaml:"WriteEquidDeathNotificationResult,omitempty"`
}

// WriteEquidImportAfterExportNotification was auto-generated from
// WSDL.
type WriteEquidImportAfterExportNotification struct {
	XMLName                         xml.Name                       `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidImportAfterExportNotification" json:"-" yaml:"-"`
	P_ManufacturerKey               string                         `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                          int                            `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidImportAfterExportRequest *EquidImportAfterExportRequest `xml:"p_EquidImportAfterExportRequest" json:"p_EquidImportAfterExportRequest" yaml:"p_EquidImportAfterExportRequest"`
}

// WriteEquidImportAfterExportNotificationResponse was auto-generated
// from WSDL.
type WriteEquidImportAfterExportNotificationResponse struct {
	WriteEquidImportAfterExportNotificationResult *WriteNotificationResult `xml:"WriteEquidImportAfterExportNotificationResult,omitempty" json:"WriteEquidImportAfterExportNotificationResult,omitempty" yaml:"WriteEquidImportAfterExportNotificationResult,omitempty"`
}

// WriteEquidImportNotification was auto-generated from WSDL.
type WriteEquidImportNotification struct {
	XMLName              xml.Name            `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidImportNotification" json:"-" yaml:"-"`
	P_ManufacturerKey    string              `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID               int                 `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidImportRequest *EquidImportRequest `xml:"p_EquidImportRequest" json:"p_EquidImportRequest" yaml:"p_EquidImportRequest"`
}

// WriteEquidImportNotificationResponse was auto-generated from
// WSDL.
type WriteEquidImportNotificationResponse struct {
	WriteEquidImportNotificationResult *EquidInitialResult `xml:"WriteEquidImportNotificationResult,omitempty" json:"WriteEquidImportNotificationResult,omitempty" yaml:"WriteEquidImportNotificationResult,omitempty"`
}

// WriteEquidMembershipOrganisationsNotification was auto-generated
// from WSDL.
type WriteEquidMembershipOrganisationsNotification struct {
	XMLName                         xml.Name                            `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidMembershipOrganisationsNotification" json:"-" yaml:"-"`
	P_ManufacturerKey               string                              `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                          int                                 `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_MembershipOrganisationRequest *EquidMembershipOrganisationRequest `xml:"p_MembershipOrganisationRequest" json:"p_MembershipOrganisationRequest" yaml:"p_MembershipOrganisationRequest"`
}

// WriteEquidMembershipOrganisationsNotificationResponse was auto-generated
// from WSDL.
type WriteEquidMembershipOrganisationsNotificationResponse struct {
	WriteEquidMembershipOrganisationsNotificationResult *WriteNotificationResult `xml:"WriteEquidMembershipOrganisationsNotificationResult,omitempty" json:"WriteEquidMembershipOrganisationsNotificationResult,omitempty" yaml:"WriteEquidMembershipOrganisationsNotificationResult,omitempty"`
}

// WriteEquidNewChipNotification was auto-generated from WSDL.
type WriteEquidNewChipNotification struct {
	XMLName               xml.Name             `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidNewChipNotification" json:"-" yaml:"-"`
	P_ManufacturerKey     string               `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                int                  `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidNewChipRequest *EquidNewChipRequest `xml:"p_EquidNewChipRequest" json:"p_EquidNewChipRequest" yaml:"p_EquidNewChipRequest"`
}

// WriteEquidNewChipNotificationResponse was auto-generated from
// WSDL.
type WriteEquidNewChipNotificationResponse struct {
	WriteEquidNewChipNotificationResult *WriteNotificationResult `xml:"WriteEquidNewChipNotificationResult,omitempty" json:"WriteEquidNewChipNotificationResult,omitempty" yaml:"WriteEquidNewChipNotificationResult,omitempty"`
}

// WriteEquidOrderBasePassNotification was auto-generated from
// WSDL.
type WriteEquidOrderBasePassNotification struct {
	XMLName                     xml.Name                   `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidOrderBasePassNotification" json:"-" yaml:"-"`
	P_ManufacturerKey           string                     `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                      int                        `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidOrderBasePassRequest *EquidOrderBasePassRequest `xml:"p_EquidOrderBasePassRequest" json:"p_EquidOrderBasePassRequest" yaml:"p_EquidOrderBasePassRequest"`
}

// WriteEquidOrderBasePassNotificationResponse was auto-generated
// from WSDL.
type WriteEquidOrderBasePassNotificationResponse struct {
	WriteEquidOrderBasePassNotificationResult *WriteNotificationResult `xml:"WriteEquidOrderBasePassNotificationResult,omitempty" json:"WriteEquidOrderBasePassNotificationResult,omitempty" yaml:"WriteEquidOrderBasePassNotificationResult,omitempty"`
}

// WriteEquidPassIssuedNotification was auto-generated from WSDL.
type WriteEquidPassIssuedNotification struct {
	XMLName                  xml.Name                `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidPassIssuedNotification" json:"-" yaml:"-"`
	P_ManufacturerKey        string                  `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                   int                     `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidPassIssuedRequest *EquidPassIssuedRequest `xml:"p_EquidPassIssuedRequest" json:"p_EquidPassIssuedRequest" yaml:"p_EquidPassIssuedRequest"`
}

// WriteEquidPassIssuedNotificationResponse was auto-generated
// from WSDL.
type WriteEquidPassIssuedNotificationResponse struct {
	WriteEquidPassIssuedNotificationResult *WriteNotificationResult `xml:"WriteEquidPassIssuedNotificationResult,omitempty" json:"WriteEquidPassIssuedNotificationResult,omitempty" yaml:"WriteEquidPassIssuedNotificationResult,omitempty"`
}

// WriteEquidPassIssuerPermissionNotification was auto-generated
// from WSDL.
type WriteEquidPassIssuerPermissionNotification struct {
	XMLName                            xml.Name                          `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidPassIssuerPermissionNotification" json:"-" yaml:"-"`
	P_ManufacturerKey                  string                            `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                             int                               `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidPassIssuerPermissionRequest *EquidPassIssuerPermissionRequest `xml:"p_EquidPassIssuerPermissionRequest" json:"p_EquidPassIssuerPermissionRequest" yaml:"p_EquidPassIssuerPermissionRequest"`
}

// WriteEquidPassIssuerPermissionNotificationResponse was auto-generated
// from WSDL.
type WriteEquidPassIssuerPermissionNotificationResponse struct {
	WriteEquidPassIssuerPermissionNotificationResult *WriteNotificationResult `xml:"WriteEquidPassIssuerPermissionNotificationResult,omitempty" json:"WriteEquidPassIssuerPermissionNotificationResult,omitempty" yaml:"WriteEquidPassIssuerPermissionNotificationResult,omitempty"`
}

// WriteEquidRelocationNotification was auto-generated from WSDL.
type WriteEquidRelocationNotification struct {
	XMLName                  xml.Name                `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidRelocationNotification" json:"-" yaml:"-"`
	P_ManufacturerKey        string                  `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                   int                     `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidRelocationRequest *EquidRelocationRequest `xml:"p_EquidRelocationRequest" json:"p_EquidRelocationRequest" yaml:"p_EquidRelocationRequest"`
}

// WriteEquidRelocationNotificationResponse was auto-generated
// from WSDL.
type WriteEquidRelocationNotificationResponse struct {
	WriteEquidRelocationNotificationResult *WriteNotificationResult `xml:"WriteEquidRelocationNotificationResult,omitempty" json:"WriteEquidRelocationNotificationResult,omitempty" yaml:"WriteEquidRelocationNotificationResult,omitempty"`
}

// WriteEquidSlaughterNotification was auto-generated from WSDL.
type WriteEquidSlaughterNotification struct {
	XMLName                 xml.Name               `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidSlaughterNotification" json:"-" yaml:"-"`
	P_ManufacturerKey       string                 `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                  int                    `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidSlaughterRequest *EquidSlaughterRequest `xml:"p_EquidSlaughterRequest" json:"p_EquidSlaughterRequest" yaml:"p_EquidSlaughterRequest"`
}

// WriteEquidSlaughterNotificationResponse was auto-generated from
// WSDL.
type WriteEquidSlaughterNotificationResponse struct {
	WriteEquidSlaughterNotificationResult *WriteNotificationResult `xml:"WriteEquidSlaughterNotificationResult,omitempty" json:"WriteEquidSlaughterNotificationResult,omitempty" yaml:"WriteEquidSlaughterNotificationResult,omitempty"`
}

// WriteEquidTypeOfUseNotification was auto-generated from WSDL.
type WriteEquidTypeOfUseNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidTypeOfUseNotification" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_Ueln            string   `xml:"p_Ueln,omitempty" json:"p_Ueln,omitempty" yaml:"p_Ueln,omitempty"`
}

// WriteEquidTypeOfUseNotificationResponse was auto-generated from
// WSDL.
type WriteEquidTypeOfUseNotificationResponse struct {
	WriteEquidTypeOfUseNotificationResult *WriteNotificationResult `xml:"WriteEquidTypeOfUseNotificationResult,omitempty" json:"WriteEquidTypeOfUseNotificationResult,omitempty" yaml:"WriteEquidTypeOfUseNotificationResult,omitempty"`
}

// WriteEquidWithersClassNotification was auto-generated from WSDL.
type WriteEquidWithersClassNotification struct {
	XMLName                    xml.Name                  `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidWithersClassNotification" json:"-" yaml:"-"`
	P_ManufacturerKey          string                    `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                     int                       `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EquidWithersClassRequest *EquidWithersClassRequest `xml:"p_EquidWithersClassRequest" json:"p_EquidWithersClassRequest" yaml:"p_EquidWithersClassRequest"`
}

// WriteEquidWithersClassNotificationResponse was auto-generated
// from WSDL.
type WriteEquidWithersClassNotificationResponse struct {
	WriteEquidWithersClassNotificationResult *WriteNotificationResult `xml:"WriteEquidWithersClassNotificationResult,omitempty" json:"WriteEquidWithersClassNotificationResult,omitempty" yaml:"WriteEquidWithersClassNotificationResult,omitempty"`
}

// WriteGroupSlaughterMovement was auto-generated from WSDL.
type WriteGroupSlaughterMovement struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovement" json:"-" yaml:"-"`
	P_ManufacturerKey          string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_ReportingTvdNumber       int      `xml:"p_ReportingTvdNumber" json:"p_ReportingTvdNumber" yaml:"p_ReportingTvdNumber"`
	P_LCID                     int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EventDate                DateTime `xml:"p_EventDate" json:"p_EventDate" yaml:"p_EventDate"`
	P_NumberOfAnimals          int      `xml:"p_NumberOfAnimals" json:"p_NumberOfAnimals" yaml:"p_NumberOfAnimals"`
	P_SourceHusbandryTvdNumber int      `xml:"p_SourceHusbandryTvdNumber" json:"p_SourceHusbandryTvdNumber" yaml:"p_SourceHusbandryTvdNumber"`
}

// WriteGroupSlaughterMovementResponse was auto-generated from
// WSDL.
type WriteGroupSlaughterMovementResponse struct {
	WriteGroupSlaughterMovementResult *WriteMovementResult `xml:"WriteGroupSlaughterMovementResult,omitempty" json:"WriteGroupSlaughterMovementResult,omitempty" yaml:"WriteGroupSlaughterMovementResult,omitempty"`
}

// WriteGroupSlaughterMovementV2 was auto-generated from WSDL.
type WriteGroupSlaughterMovementV2 struct {
	XMLName                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovementV2" json:"-" yaml:"-"`
	P_ManufacturerKey             string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_ReportingTvdNumber          int      `xml:"p_ReportingTvdNumber" json:"p_ReportingTvdNumber" yaml:"p_ReportingTvdNumber"`
	P_LCID                        int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EventDate                   DateTime `xml:"p_EventDate" json:"p_EventDate" yaml:"p_EventDate"`
	P_Genus                       int      `xml:"p_Genus" json:"p_Genus" yaml:"p_Genus"`
	P_NumberOfAnimals             int      `xml:"p_NumberOfAnimals" json:"p_NumberOfAnimals" yaml:"p_NumberOfAnimals"`
	P_SourceHusbandryTvdNumber    int      `xml:"p_SourceHusbandryTvdNumber" json:"p_SourceHusbandryTvdNumber" yaml:"p_SourceHusbandryTvdNumber"`
	P_SlaughterInitiatorTvdNumber int      `xml:"p_SlaughterInitiatorTvdNumber,omitempty" json:"p_SlaughterInitiatorTvdNumber,omitempty" yaml:"p_SlaughterInitiatorTvdNumber,omitempty"`
}

// WriteGroupSlaughterMovementV2Response was auto-generated from
// WSDL.
type WriteGroupSlaughterMovementV2Response struct {
	WriteGroupSlaughterMovementV2Result *WriteMovementResult `xml:"WriteGroupSlaughterMovementV2Result,omitempty" json:"WriteGroupSlaughterMovementV2Result,omitempty" yaml:"WriteGroupSlaughterMovementV2Result,omitempty"`
}

// WriteMovementResult was auto-generated from WSDL.
type WriteMovementResult struct {
	MovementId       int               `xml:"MovementId,omitempty" json:"MovementId,omitempty" yaml:"MovementId,omitempty"`
	ProcessingResult *ProcessingResult `xml:"ProcessingResult,omitempty" json:"ProcessingResult,omitempty" yaml:"ProcessingResult,omitempty"`
}

// WriteNewEarTagOrder was auto-generated from WSDL.
type WriteNewEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewEarTagOrder" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_EarTagType      int      `xml:"p_EarTagType" json:"p_EarTagType" yaml:"p_EarTagType"`
	P_Amount          int      `xml:"p_Amount" json:"p_Amount" yaml:"p_Amount"`
}

// WriteNewEarTagOrderResponse was auto-generated from WSDL.
type WriteNewEarTagOrderResponse struct {
	WriteNewEarTagOrderResult *NewEarTagOrderResult `xml:"WriteNewEarTagOrderResult,omitempty" json:"WriteNewEarTagOrderResult,omitempty" yaml:"WriteNewEarTagOrderResult,omitempty"`
}

// WriteNewLabelEarTagOrder was auto-generated from WSDL.
type WriteNewLabelEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewLabelEarTagOrder" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_LabelEarTagType int      `xml:"p_LabelEarTagType" json:"p_LabelEarTagType" yaml:"p_LabelEarTagType"`
	P_Amount          int      `xml:"p_Amount" json:"p_Amount" yaml:"p_Amount"`
}

// WriteNewLabelEarTagOrderResponse was auto-generated from WSDL.
type WriteNewLabelEarTagOrderResponse struct {
	WriteNewLabelEarTagOrderResult *NewEarTagOrderResult `xml:"WriteNewLabelEarTagOrderResult,omitempty" json:"WriteNewLabelEarTagOrderResult,omitempty" yaml:"WriteNewLabelEarTagOrderResult,omitempty"`
}

// WriteNotificationResult was auto-generated from WSDL.
type WriteNotificationResult struct {
	Result *ProcessingResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

// WritePigEntryMovement was auto-generated from WSDL.
type WritePigEntryMovement struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovement" json:"-" yaml:"-"`
	P_ManufacturerKey          string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_ReportingTvdNumber       int      `xml:"p_ReportingTvdNumber" json:"p_ReportingTvdNumber" yaml:"p_ReportingTvdNumber"`
	P_LCID                     int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EventDate                DateTime `xml:"p_EventDate" json:"p_EventDate" yaml:"p_EventDate"`
	P_NumberOfAnimals          int      `xml:"p_NumberOfAnimals" json:"p_NumberOfAnimals" yaml:"p_NumberOfAnimals"`
	P_SourceHusbandryTvdNumber int      `xml:"p_SourceHusbandryTvdNumber" json:"p_SourceHusbandryTvdNumber" yaml:"p_SourceHusbandryTvdNumber"`
}

// WritePigEntryMovementResponse was auto-generated from WSDL.
type WritePigEntryMovementResponse struct {
	WritePigEntryMovementResult *WriteMovementResult `xml:"WritePigEntryMovementResult,omitempty" json:"WritePigEntryMovementResult,omitempty" yaml:"WritePigEntryMovementResult,omitempty"`
}

// WritePigEntryMovementV2 was auto-generated from WSDL.
type WritePigEntryMovementV2 struct {
	XMLName xml.Name                        `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovementV2" json:"-" yaml:"-"`
	Model   *WritePigEntryMovementV2Request `xml:"Model" json:"Model" yaml:"Model"`
}

// WritePigEntryMovementV2Request was auto-generated from WSDL.
type WritePigEntryMovementV2Request struct {
	ManufacturerKey          string          `xml:"ManufacturerKey,omitempty" json:"ManufacturerKey,omitempty" yaml:"ManufacturerKey,omitempty"`
	LCID                     int             `xml:"LCID,omitempty" json:"LCID,omitempty" yaml:"LCID,omitempty"`
	ReportingTVDNumber       int             `xml:"ReportingTVDNumber,omitempty" json:"ReportingTVDNumber,omitempty" yaml:"ReportingTVDNumber,omitempty"`
	EventDate                DateTime        `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
	NumberOfAnimals          int             `xml:"NumberOfAnimals,omitempty" json:"NumberOfAnimals,omitempty" yaml:"NumberOfAnimals,omitempty"`
	SourceHusbandryTVDNumber int             `xml:"SourceHusbandryTVDNumber,omitempty" json:"SourceHusbandryTVDNumber,omitempty" yaml:"SourceHusbandryTVDNumber,omitempty"`
	PigCategory              EnumPigCategory `xml:"PigCategory,omitempty" json:"PigCategory,omitempty" yaml:"PigCategory,omitempty"`
}

// WritePigEntryMovementV2Response was auto-generated from WSDL.
type WritePigEntryMovementV2Response struct {
	WritePigEntryMovementV2Result *WriteMovementResult `xml:"WritePigEntryMovementV2Result,omitempty" json:"WritePigEntryMovementV2Result,omitempty" yaml:"WritePigEntryMovementV2Result,omitempty"`
}

// WritePigSlaughterMovement was auto-generated from WSDL.
type WritePigSlaughterMovement struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigSlaughterMovement" json:"-" yaml:"-"`
	P_ManufacturerKey           string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_ReportingTvdNumber        int      `xml:"p_ReportingTvdNumber" json:"p_ReportingTvdNumber" yaml:"p_ReportingTvdNumber"`
	P_LCID                      int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_EventDate                 DateTime `xml:"p_EventDate" json:"p_EventDate" yaml:"p_EventDate"`
	P_SourceHusbandryTvdNumber  int      `xml:"p_SourceHusbandryTvdNumber" json:"p_SourceHusbandryTvdNumber" yaml:"p_SourceHusbandryTvdNumber"`
	P_MessageID                 int64    `xml:"p_MessageID" json:"p_MessageID" yaml:"p_MessageID"`
	P_ClassifierNumber          int16   `xml:"p_ClassifierNumber" json:"p_ClassifierNumber" yaml:"p_ClassifierNumber"`
	P_ClassifierEquipmentID     string   `xml:"p_ClassifierEquipmentID,omitempty" json:"p_ClassifierEquipmentID,omitempty" yaml:"p_ClassifierEquipmentID,omitempty"`
	P_ContractorNumberSlaughter string   `xml:"p_ContractorNumberSlaughter,omitempty" json:"p_ContractorNumberSlaughter,omitempty" yaml:"p_ContractorNumberSlaughter,omitempty"`
	P_MFA                       int16   `xml:"p_MFA" json:"p_MFA" yaml:"p_MFA"`
	P_Weight                    float64  `xml:"p_Weight" json:"p_Weight" yaml:"p_Weight"`
	P_SlaughterInitiator        int      `xml:"p_SlaughterInitiator" json:"p_SlaughterInitiator" yaml:"p_SlaughterInitiator"`
}

// WritePigSlaughterMovementResponse was auto-generated from WSDL.
type WritePigSlaughterMovementResponse struct {
	WritePigSlaughterMovementResult *WriteMovementResult `xml:"WritePigSlaughterMovementResult,omitempty" json:"WritePigSlaughterMovementResult,omitempty" yaml:"WritePigSlaughterMovementResult,omitempty"`
}

// WritePoultryBarnInNotification was auto-generated from WSDL.
type WritePoultryBarnInNotification struct {
	XMLName                     xml.Name                   `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultryBarnInNotification" json:"-" yaml:"-"`
	P_ManufacturerKey           string                     `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                      int                        `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber                 int                        `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_PoultryBarnInNotification *PoultryBarnInNotification `xml:"p_PoultryBarnInNotification" json:"p_PoultryBarnInNotification" yaml:"p_PoultryBarnInNotification"`
}

// WritePoultryBarnInNotificationResponse was auto-generated from
// WSDL.
type WritePoultryBarnInNotificationResponse struct {
	WritePoultryBarnInNotificationResult *ProcessingResult `xml:"WritePoultryBarnInNotificationResult,omitempty" json:"WritePoultryBarnInNotificationResult,omitempty" yaml:"WritePoultryBarnInNotificationResult,omitempty"`
}

// WritePoultrySlaughterNotification was auto-generated from WSDL.
type WritePoultrySlaughterNotification struct {
	XMLName              xml.Name        `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultrySlaughterNotification" json:"-" yaml:"-"`
	P_ManufacturerKey    string          `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID               int             `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber          int             `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_EventDate          DateTime        `xml:"p_EventDate" json:"p_EventDate" yaml:"p_EventDate"`
	P_Weight             int             `xml:"p_Weight" json:"p_Weight" yaml:"p_Weight"`
	P_PoultryType        EnumPoultryType `xml:"p_PoultryType" json:"p_PoultryType" yaml:"p_PoultryType"`
	P_SourceTVDNumber    int             `xml:"p_SourceTVDNumber" json:"p_SourceTVDNumber" yaml:"p_SourceTVDNumber"`
	P_DeliveryFileNumber string          `xml:"p_DeliveryFileNumber,omitempty" json:"p_DeliveryFileNumber,omitempty" yaml:"p_DeliveryFileNumber,omitempty"`
}

// WritePoultrySlaughterNotificationResponse was auto-generated
// from WSDL.
type WritePoultrySlaughterNotificationResponse struct {
	WritePoultrySlaughterNotificationResult *ProcessingResult `xml:"WritePoultrySlaughterNotificationResult,omitempty" json:"WritePoultrySlaughterNotificationResult,omitempty" yaml:"WritePoultrySlaughterNotificationResult,omitempty"`
}

// WriteReplacementBatchEarTagOrder was auto-generated from WSDL.
type WriteReplacementBatchEarTagOrder struct {
	XMLName                                   xml.Name                                 `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementBatchEarTagOrder" json:"-" yaml:"-"`
	P_ManufacturerKey                         string                                   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID                                    int                                      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_WriteReplacementBatchEarTagOrderRequest *WriteReplacementBatchEarTagOrderRequest `xml:"p_WriteReplacementBatchEarTagOrderRequest" json:"p_WriteReplacementBatchEarTagOrderRequest" yaml:"p_WriteReplacementBatchEarTagOrderRequest"`
}

// WriteReplacementBatchEarTagOrderRequest was auto-generated from
// WSDL.
type WriteReplacementBatchEarTagOrderRequest struct {
	TVDNumber            int                   `xml:"TVDNumber,omitempty" json:"TVDNumber,omitempty" yaml:"TVDNumber,omitempty"`
	Genus                EnumGenus             `xml:"Genus,omitempty" json:"Genus,omitempty" yaml:"Genus,omitempty"`
	EartagOrderItemArray *EartagOrderItemArray `xml:"EartagOrderItemArray,omitempty" json:"EartagOrderItemArray,omitempty" yaml:"EartagOrderItemArray,omitempty"`
}

// WriteReplacementBatchEarTagOrderResponse was auto-generated
// from WSDL.
type WriteReplacementBatchEarTagOrderResponse struct {
	WriteReplacementBatchEarTagOrderResult *ReplacementEarTagOrdersResult `xml:"WriteReplacementBatchEarTagOrderResult,omitempty" json:"WriteReplacementBatchEarTagOrderResult,omitempty" yaml:"WriteReplacementBatchEarTagOrderResult,omitempty"`
}

// WriteReplacementEarTagOrder was auto-generated from WSDL.
type WriteReplacementEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementEarTagOrder" json:"-" yaml:"-"`
	P_ManufacturerKey string   `xml:"p_ManufacturerKey,omitempty" json:"p_ManufacturerKey,omitempty" yaml:"p_ManufacturerKey,omitempty"`
	P_LCID            int      `xml:"p_LCID" json:"p_LCID" yaml:"p_LCID"`
	P_TVDNumber       int      `xml:"p_TVDNumber" json:"p_TVDNumber" yaml:"p_TVDNumber"`
	P_Genus           int      `xml:"p_Genus" json:"p_Genus" yaml:"p_Genus"`
	P_EarTagNumber    string   `xml:"p_EarTagNumber,omitempty" json:"p_EarTagNumber,omitempty" yaml:"p_EarTagNumber,omitempty"`
	P_LeftEarTag      bool     `xml:"p_LeftEarTag,omitempty" json:"p_LeftEarTag,omitempty" yaml:"p_LeftEarTag,omitempty"`
	P_RightEarTag     bool     `xml:"p_RightEarTag,omitempty" json:"p_RightEarTag,omitempty" yaml:"p_RightEarTag,omitempty"`
}

// WriteReplacementEarTagOrderResponse was auto-generated from
// WSDL.
type WriteReplacementEarTagOrderResponse struct {
	WriteReplacementEarTagOrderResult *NewEarTagOrderResult `xml:"WriteReplacementEarTagOrderResult,omitempty" json:"WriteReplacementEarTagOrderResult,omitempty" yaml:"WriteReplacementEarTagOrderResult,omitempty"`
}

// animalTracingPortType implements the AnimalTracingPortType interface.
type animalTracingPortType struct {
	cli *soap.Client
}

// CheckCattleForDisposalContribution was auto-generated from WSDL.
func (p *animalTracingPortType) CheckCattleForDisposalContribution(α *CheckCattleForDisposalContribution) (β *CheckCattleForDisposalContributionResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M CheckCattleForDisposalContributionResponse `xml:"CheckCattleForDisposalContributionResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// CheckCattlesForPassport was auto-generated from WSDL.
func (p *animalTracingPortType) CheckCattlesForPassport(α *CheckCattlesForPassport) (β *CheckCattlesForPassportResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M CheckCattlesForPassportResponse `xml:"CheckCattlesForPassportResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// DeleteAnimalHusbandryMemberships was auto-generated from WSDL.
func (p *animalTracingPortType) DeleteAnimalHusbandryMemberships(α *DeleteAnimalHusbandryMemberships) (β *DeleteAnimalHusbandryMembershipsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M DeleteAnimalHusbandryMembershipsResponse `xml:"DeleteAnimalHusbandryMembershipsResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// DeleteCattleNotifications was auto-generated from WSDL.
func (p *animalTracingPortType) DeleteCattleNotifications(α *DeleteCattleNotifications) (β *DeleteCattleNotificationsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M DeleteCattleNotificationsResponse `xml:"DeleteCattleNotificationsResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// DeleteEarTagOrder was auto-generated from WSDL.
func (p *animalTracingPortType) DeleteEarTagOrder(α *DeleteEarTagOrder) (β *DeleteEarTagOrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M DeleteEarTagOrderResponse `xml:"DeleteEarTagOrderResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// DeleteLabelEarTagOrder was auto-generated from WSDL.
func (p *animalTracingPortType) DeleteLabelEarTagOrder(α *DeleteLabelEarTagOrder) (β *DeleteLabelEarTagOrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M DeleteLabelEarTagOrderResponse `xml:"DeleteLabelEarTagOrderResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// DisableDataAccess was auto-generated from WSDL.
func (p *animalTracingPortType) DisableDataAccess(α *DisableDataAccess) (β *DisableDataAccessResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M DisableDataAccessResponse `xml:"DisableDataAccessResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// EnableDataAccess was auto-generated from WSDL.
func (p *animalTracingPortType) EnableDataAccess(α *EnableDataAccess) (β *EnableDataAccessResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M EnableDataAccessResponse `xml:"EnableDataAccessResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAnimalHusbandryAddress was auto-generated from WSDL.
func (p *animalTracingPortType) GetAnimalHusbandryAddress(α *GetAnimalHusbandryAddress) (β *GetAnimalHusbandryAddressResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetAnimalHusbandryAddressResponse `xml:"GetAnimalHusbandryAddressResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAnimalHusbandryAddressV2 was auto-generated from WSDL.
func (p *animalTracingPortType) GetAnimalHusbandryAddressV2(α *GetAnimalHusbandryAddressV2) (β *GetAnimalHusbandryAddressV2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetAnimalHusbandryAddressV2Response `xml:"GetAnimalHusbandryAddressV2Response"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAnimalHusbandryDetail was auto-generated from WSDL.
func (p *animalTracingPortType) GetAnimalHusbandryDetail(α *GetAnimalHusbandryDetail) (β *GetAnimalHusbandryDetailResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetAnimalHusbandryDetailResponse `xml:"GetAnimalHusbandryDetailResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAnimalHusbandryMemberships was auto-generated from WSDL.
func (p *animalTracingPortType) GetAnimalHusbandryMemberships(α *GetAnimalHusbandryMemberships) (β *GetAnimalHusbandryMembershipsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetAnimalHusbandryMembershipsResponse `xml:"GetAnimalHusbandryMembershipsResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleDetail was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleDetail(α *GetCattleDetail) (β *GetCattleDetailResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleDetailResponse `xml:"GetCattleDetailResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleDetailV2 was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleDetailV2(α *GetCattleDetailV2) (β *GetCattleDetailV2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleDetailV2Response `xml:"GetCattleDetailV2Response"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleEarTags was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleEarTags(α *GetCattleEarTags) (β *GetCattleEarTagsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleEarTagsResponse `xml:"GetCattleEarTagsResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleHistory was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleHistory(α *GetCattleHistory) (β *GetCattleHistoryResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleHistoryResponse `xml:"GetCattleHistoryResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleHistoryV2 was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleHistoryV2(α *GetCattleHistoryV2) (β *GetCattleHistoryV2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleHistoryV2Response `xml:"GetCattleHistoryV2Response"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleLivestock was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleLivestock(α *GetCattleLivestock) (β *GetCattleLivestockResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleLivestockResponse `xml:"GetCattleLivestockResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleLivestockV2 was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleLivestockV2(α *GetCattleLivestockV2) (β *GetCattleLivestockV2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleLivestockV2Response `xml:"GetCattleLivestockV2Response"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleMovements was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleMovements(α *GetCattleMovements) (β *GetCattleMovementsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleMovementsResponse `xml:"GetCattleMovementsResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleOffsprings was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleOffsprings(α *GetCattleOffsprings) (β *GetCattleOffspringsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleOffspringsResponse `xml:"GetCattleOffspringsResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleStatus was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleStatus(α *GetCattleStatus) (β *GetCattleStatusResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleStatusResponse `xml:"GetCattleStatusResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattleStatusV2 was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattleStatusV2(α *GetCattleStatusV2) (β *GetCattleStatusV2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattleStatusV2Response `xml:"GetCattleStatusV2Response"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCattlesPerLeavingDate was auto-generated from WSDL.
func (p *animalTracingPortType) GetCattlesPerLeavingDate(α *GetCattlesPerLeavingDate) (β *GetCattlesPerLeavingDateResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCattlesPerLeavingDateResponse `xml:"GetCattlesPerLeavingDateResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCodes was auto-generated from WSDL.
func (p *animalTracingPortType) GetCodes(α *GetCodes) (β *GetCodesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetCodesResponse `xml:"GetCodesResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetEarTagOrders was auto-generated from WSDL.
func (p *animalTracingPortType) GetEarTagOrders(α *GetEarTagOrders) (β *GetEarTagOrdersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetEarTagOrdersResponse `xml:"GetEarTagOrdersResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetEquidDetail was auto-generated from WSDL.
func (p *animalTracingPortType) GetEquidDetail(α *GetEquidDetail) (β *GetEquidDetailResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetEquidDetailResponse `xml:"GetEquidDetailResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetEquidLivestock was auto-generated from WSDL.
func (p *animalTracingPortType) GetEquidLivestock(α *GetEquidLivestock) (β *GetEquidLivestockResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetEquidLivestockResponse `xml:"GetEquidLivestockResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetEquidOwnershipList was auto-generated from WSDL.
func (p *animalTracingPortType) GetEquidOwnershipList(α *GetEquidOwnershipList) (β *GetEquidOwnershipListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetEquidOwnershipListResponse `xml:"GetEquidOwnershipListResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetEquidSlaughterList was auto-generated from WSDL.
func (p *animalTracingPortType) GetEquidSlaughterList(α *GetEquidSlaughterList) (β *GetEquidSlaughterListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetEquidSlaughterListResponse `xml:"GetEquidSlaughterListResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetEquidsWithNotificationsOfMO was auto-generated from WSDL.
func (p *animalTracingPortType) GetEquidsWithNotificationsOfMO(α *GetEquidsWithNotificationsOfMO) (β *GetEquidsWithNotificationsOfMOResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetEquidsWithNotificationsOfMOResponse `xml:"GetEquidsWithNotificationsOfMOResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetFarmManager was auto-generated from WSDL.
func (p *animalTracingPortType) GetFarmManager(α *GetFarmManager) (β *GetFarmManagerResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetFarmManagerResponse `xml:"GetFarmManagerResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetFarmers was auto-generated from WSDL.
func (p *animalTracingPortType) GetFarmers(α *GetFarmers) (β *GetFarmersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetFarmersResponse `xml:"GetFarmersResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetLabelEarTagOrders was auto-generated from WSDL.
func (p *animalTracingPortType) GetLabelEarTagOrders(α *GetLabelEarTagOrders) (β *GetLabelEarTagOrdersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetLabelEarTagOrdersResponse `xml:"GetLabelEarTagOrdersResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetMembershipForOrganisation was auto-generated from WSDL.
func (p *animalTracingPortType) GetMembershipForOrganisation(α *GetMembershipForOrganisation) (β *GetMembershipForOrganisationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetMembershipForOrganisationResponse `xml:"GetMembershipForOrganisationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetPersonAddress was auto-generated from WSDL.
func (p *animalTracingPortType) GetPersonAddress(α *GetPersonAddress) (β *GetPersonAddressResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetPersonAddressResponse `xml:"GetPersonAddressResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetPersonIdentifiers was auto-generated from WSDL.
func (p *animalTracingPortType) GetPersonIdentifiers(α *GetPersonIdentifiers) (β *GetPersonIdentifiersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetPersonIdentifiersResponse `xml:"GetPersonIdentifiersResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetPigArrivalNotificationForBreedingOrganisation was auto-generated
// from WSDL.
func (p *animalTracingPortType) GetPigArrivalNotificationForBreedingOrganisation(α *GetPigArrivalNotificationForBreedingOrganisation) (β *GetPigArrivalNotificationForBreedingOrganisationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetPigArrivalNotificationForBreedingOrganisationResponse `xml:"GetPigArrivalNotificationForBreedingOrganisationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetPoultryBarnInNotifications was auto-generated from WSDL.
func (p *animalTracingPortType) GetPoultryBarnInNotifications(α *GetPoultryBarnInNotifications) (β *GetPoultryBarnInNotificationsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetPoultryBarnInNotificationsResponse `xml:"GetPoultryBarnInNotificationsResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRegisteredGenera was auto-generated from WSDL.
func (p *animalTracingPortType) GetRegisteredGenera(α *GetRegisteredGenera) (β *GetRegisteredGeneraResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetRegisteredGeneraResponse `xml:"GetRegisteredGeneraResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetSmallAnimalSlaughterList was auto-generated from WSDL.
func (p *animalTracingPortType) GetSmallAnimalSlaughterList(α *GetSmallAnimalSlaughterList) (β *GetSmallAnimalSlaughterListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M GetSmallAnimalSlaughterListResponse `xml:"GetSmallAnimalSlaughterListResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SearchEquid was auto-generated from WSDL.
func (p *animalTracingPortType) SearchEquid(α *SearchEquid) (β *SearchEquidResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M SearchEquidResponse `xml:"SearchEquidResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Version was auto-generated from WSDL.
func (p *animalTracingPortType) Version(α *Version) (β *VersionResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M VersionResponse `xml:"VersionResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteAnimalClassificationNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteAnimalClassificationNotification(α *WriteAnimalClassificationNotification) (β *WriteAnimalClassificationNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteAnimalClassificationNotificationResponse `xml:"WriteAnimalClassificationNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteAnimalClassificationNotificationV2 was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteAnimalClassificationNotificationV2(α *WriteAnimalClassificationNotificationV2) (β *WriteAnimalClassificationNotificationV2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteAnimalClassificationNotificationV2Response `xml:"WriteAnimalClassificationNotificationV2Response"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleArrivalBatchNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteCattleArrivalBatchNotification(α *WriteCattleArrivalBatchNotification) (β *WriteCattleArrivalBatchNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleArrivalBatchNotificationResponse `xml:"WriteCattleArrivalBatchNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleBirthNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteCattleBirthNotification(α *WriteCattleBirthNotification) (β *WriteCattleBirthNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleBirthNotificationResponse `xml:"WriteCattleBirthNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleChangeNameNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteCattleChangeNameNotification(α *WriteCattleChangeNameNotification) (β *WriteCattleChangeNameNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleChangeNameNotificationResponse `xml:"WriteCattleChangeNameNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleDaystayBatchNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteCattleDaystayBatchNotification(α *WriteCattleDaystayBatchNotification) (β *WriteCattleDaystayBatchNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleDaystayBatchNotificationResponse `xml:"WriteCattleDaystayBatchNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleDeathBirthNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteCattleDeathBirthNotification(α *WriteCattleDeathBirthNotification) (β *WriteCattleDeathBirthNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleDeathBirthNotificationResponse `xml:"WriteCattleDeathBirthNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleDeceasedNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteCattleDeceasedNotification(α *WriteCattleDeceasedNotification) (β *WriteCattleDeceasedNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleDeceasedNotificationResponse `xml:"WriteCattleDeceasedNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleDeformationNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteCattleDeformationNotification(α *WriteCattleDeformationNotification) (β *WriteCattleDeformationNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleDeformationNotificationResponse `xml:"WriteCattleDeformationNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleExportNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteCattleExportNotification(α *WriteCattleExportNotification) (β *WriteCattleExportNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleExportNotificationResponse `xml:"WriteCattleExportNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleImportAfterExportNotification was auto-generated
// from WSDL.
func (p *animalTracingPortType) WriteCattleImportAfterExportNotification(α *WriteCattleImportAfterExportNotification) (β *WriteCattleImportAfterExportNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleImportAfterExportNotificationResponse `xml:"WriteCattleImportAfterExportNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleLeavingBatchNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteCattleLeavingBatchNotification(α *WriteCattleLeavingBatchNotification) (β *WriteCattleLeavingBatchNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleLeavingBatchNotificationResponse `xml:"WriteCattleLeavingBatchNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattlePassportOrders was auto-generated from WSDL.
func (p *animalTracingPortType) WriteCattlePassportOrders(α *WriteCattlePassportOrders) (β *WriteCattlePassportOrdersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattlePassportOrdersResponse `xml:"WriteCattlePassportOrdersResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleSlaughterBatchNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteCattleSlaughterBatchNotification(α *WriteCattleSlaughterBatchNotification) (β *WriteCattleSlaughterBatchNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleSlaughterBatchNotificationResponse `xml:"WriteCattleSlaughterBatchNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleSlaughterBatchNotificationV2 was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteCattleSlaughterBatchNotificationV2(α *WriteCattleSlaughterBatchNotificationV2) (β *WriteCattleSlaughterBatchNotificationV2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleSlaughterBatchNotificationV2Response `xml:"WriteCattleSlaughterBatchNotificationV2Response"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleTypeOfUseNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteCattleTypeOfUseNotification(α *WriteCattleTypeOfUseNotification) (β *WriteCattleTypeOfUseNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleTypeOfUseNotificationResponse `xml:"WriteCattleTypeOfUseNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteCattleYardSlaughterNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteCattleYardSlaughterNotification(α *WriteCattleYardSlaughterNotification) (β *WriteCattleYardSlaughterNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteCattleYardSlaughterNotificationResponse `xml:"WriteCattleYardSlaughterNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidAcquireOwnershipNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteEquidAcquireOwnershipNotification(α *WriteEquidAcquireOwnershipNotification) (β *WriteEquidAcquireOwnershipNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidAcquireOwnershipNotificationResponse `xml:"WriteEquidAcquireOwnershipNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidBasicDataNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidBasicDataNotification(α *WriteEquidBasicDataNotification) (β *WriteEquidBasicDataNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidBasicDataNotificationResponse `xml:"WriteEquidBasicDataNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidBirthNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidBirthNotification(α *WriteEquidBirthNotification) (β *WriteEquidBirthNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidBirthNotificationResponse `xml:"WriteEquidBirthNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidCastrationNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidCastrationNotification(α *WriteEquidCastrationNotification) (β *WriteEquidCastrationNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidCastrationNotificationResponse `xml:"WriteEquidCastrationNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidCedeOwnershipNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteEquidCedeOwnershipNotification(α *WriteEquidCedeOwnershipNotification) (β *WriteEquidCedeOwnershipNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidCedeOwnershipNotificationResponse `xml:"WriteEquidCedeOwnershipNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidDeathNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidDeathNotification(α *WriteEquidDeathNotification) (β *WriteEquidDeathNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidDeathNotificationResponse `xml:"WriteEquidDeathNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidImportAfterExportNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteEquidImportAfterExportNotification(α *WriteEquidImportAfterExportNotification) (β *WriteEquidImportAfterExportNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidImportAfterExportNotificationResponse `xml:"WriteEquidImportAfterExportNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidImportNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidImportNotification(α *WriteEquidImportNotification) (β *WriteEquidImportNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidImportNotificationResponse `xml:"WriteEquidImportNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidMembershipOrganisationsNotification was auto-generated
// from WSDL.
func (p *animalTracingPortType) WriteEquidMembershipOrganisationsNotification(α *WriteEquidMembershipOrganisationsNotification) (β *WriteEquidMembershipOrganisationsNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidMembershipOrganisationsNotificationResponse `xml:"WriteEquidMembershipOrganisationsNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidNewChipNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidNewChipNotification(α *WriteEquidNewChipNotification) (β *WriteEquidNewChipNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidNewChipNotificationResponse `xml:"WriteEquidNewChipNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidOrderBasePassNotification was auto-generated from
// WSDL.
func (p *animalTracingPortType) WriteEquidOrderBasePassNotification(α *WriteEquidOrderBasePassNotification) (β *WriteEquidOrderBasePassNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidOrderBasePassNotificationResponse `xml:"WriteEquidOrderBasePassNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidPassIssuedNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidPassIssuedNotification(α *WriteEquidPassIssuedNotification) (β *WriteEquidPassIssuedNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidPassIssuedNotificationResponse `xml:"WriteEquidPassIssuedNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidPassIssuerPermissionNotification was auto-generated
// from WSDL.
func (p *animalTracingPortType) WriteEquidPassIssuerPermissionNotification(α *WriteEquidPassIssuerPermissionNotification) (β *WriteEquidPassIssuerPermissionNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidPassIssuerPermissionNotificationResponse `xml:"WriteEquidPassIssuerPermissionNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidRelocationNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidRelocationNotification(α *WriteEquidRelocationNotification) (β *WriteEquidRelocationNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidRelocationNotificationResponse `xml:"WriteEquidRelocationNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidSlaughterNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidSlaughterNotification(α *WriteEquidSlaughterNotification) (β *WriteEquidSlaughterNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidSlaughterNotificationResponse `xml:"WriteEquidSlaughterNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidTypeOfUseNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidTypeOfUseNotification(α *WriteEquidTypeOfUseNotification) (β *WriteEquidTypeOfUseNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidTypeOfUseNotificationResponse `xml:"WriteEquidTypeOfUseNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteEquidWithersClassNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WriteEquidWithersClassNotification(α *WriteEquidWithersClassNotification) (β *WriteEquidWithersClassNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteEquidWithersClassNotificationResponse `xml:"WriteEquidWithersClassNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteGroupSlaughterMovement was auto-generated from WSDL.
func (p *animalTracingPortType) WriteGroupSlaughterMovement(α *WriteGroupSlaughterMovement) (β *WriteGroupSlaughterMovementResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteGroupSlaughterMovementResponse `xml:"WriteGroupSlaughterMovementResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteGroupSlaughterMovementV2 was auto-generated from WSDL.
func (p *animalTracingPortType) WriteGroupSlaughterMovementV2(α *WriteGroupSlaughterMovementV2) (β *WriteGroupSlaughterMovementV2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteGroupSlaughterMovementV2Response `xml:"WriteGroupSlaughterMovementV2Response"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteNewEarTagOrder was auto-generated from WSDL.
func (p *animalTracingPortType) WriteNewEarTagOrder(α *WriteNewEarTagOrder) (β *WriteNewEarTagOrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteNewEarTagOrderResponse `xml:"WriteNewEarTagOrderResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteNewLabelEarTagOrder was auto-generated from WSDL.
func (p *animalTracingPortType) WriteNewLabelEarTagOrder(α *WriteNewLabelEarTagOrder) (β *WriteNewLabelEarTagOrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteNewLabelEarTagOrderResponse `xml:"WriteNewLabelEarTagOrderResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WritePigEntryMovement was auto-generated from WSDL.
func (p *animalTracingPortType) WritePigEntryMovement(α *WritePigEntryMovement) (β *WritePigEntryMovementResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WritePigEntryMovementResponse `xml:"WritePigEntryMovementResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WritePigEntryMovementV2 was auto-generated from WSDL.
func (p *animalTracingPortType) WritePigEntryMovementV2(α *WritePigEntryMovementV2) (β *WritePigEntryMovementV2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WritePigEntryMovementV2Response `xml:"WritePigEntryMovementV2Response"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WritePigSlaughterMovement was auto-generated from WSDL.
func (p *animalTracingPortType) WritePigSlaughterMovement(α *WritePigSlaughterMovement) (β *WritePigSlaughterMovementResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WritePigSlaughterMovementResponse `xml:"WritePigSlaughterMovementResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WritePoultryBarnInNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WritePoultryBarnInNotification(α *WritePoultryBarnInNotification) (β *WritePoultryBarnInNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WritePoultryBarnInNotificationResponse `xml:"WritePoultryBarnInNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WritePoultrySlaughterNotification was auto-generated from WSDL.
func (p *animalTracingPortType) WritePoultrySlaughterNotification(α *WritePoultrySlaughterNotification) (β *WritePoultrySlaughterNotificationResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WritePoultrySlaughterNotificationResponse `xml:"WritePoultrySlaughterNotificationResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteReplacementBatchEarTagOrder was auto-generated from WSDL.
func (p *animalTracingPortType) WriteReplacementBatchEarTagOrder(α *WriteReplacementBatchEarTagOrder) (β *WriteReplacementBatchEarTagOrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteReplacementBatchEarTagOrderResponse `xml:"WriteReplacementBatchEarTagOrderResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// WriteReplacementEarTagOrder was auto-generated from WSDL.
func (p *animalTracingPortType) WriteReplacementEarTagOrder(α *WriteReplacementEarTagOrder) (β *WriteReplacementEarTagOrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
				M WriteReplacementEarTagOrderResponse `xml:"WriteReplacementEarTagOrderResponse"`
			}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

