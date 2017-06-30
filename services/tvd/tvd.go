package tvd

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type Version struct {
	//edited
	XMLName          xml.Name `xml:"ns:Version"`
	Action           string   `xml:"-"`

	PManufacturerKey string `xml:"ns:p_ManufacturerKey,omitempty"`
}

type VersionResponse struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 VersionResponse"`

	VersionResult string `xml:"VersionResult,omitempty"`
}

type WritePigEntryMovementV2 struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovementV2"`

	Model   *WritePigEntryMovementV2Request `xml:"Model,omitempty"`
}

type WritePigEntryMovementV2Response struct {
	XMLName                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovementV2Response"`

	WritePigEntryMovementV2Result *WriteMovementResult `xml:"WritePigEntryMovementV2Result,omitempty"`
}

type WritePigEntryMovement struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovement"`

	PManufacturerKey          string    `xml:"p_ManufacturerKey,omitempty"`
	PReportingTvdNumber       int32     `xml:"p_ReportingTvdNumber,omitempty"`
	PLCID                     int32     `xml:"p_LCID,omitempty"`
	PEventDate                time.Time `xml:"p_EventDate,omitempty"`
	PNumberOfAnimals          int32     `xml:"p_NumberOfAnimals,omitempty"`
	PSourceHusbandryTvdNumber int32     `xml:"p_SourceHusbandryTvdNumber,omitempty"`
}

type WritePigEntryMovementResponse struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovementResponse"`

	WritePigEntryMovementResult *WriteMovementResult `xml:"WritePigEntryMovementResult,omitempty"`
}

type WritePigSlaughterMovement struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigSlaughterMovement"`

	PManufacturerKey           string    `xml:"p_ManufacturerKey,omitempty"`
	PReportingTvdNumber        int32     `xml:"p_ReportingTvdNumber,omitempty"`
	PLCID                      int32     `xml:"p_LCID,omitempty"`
	PEventDate                 time.Time `xml:"p_EventDate,omitempty"`
	PSourceHusbandryTvdNumber  int32     `xml:"p_SourceHusbandryTvdNumber,omitempty"`
	PMessageID                 int64     `xml:"p_MessageID,omitempty"`
	PClassifierNumber          int16     `xml:"p_ClassifierNumber,omitempty"`
	PClassifierEquipmentID     string    `xml:"p_ClassifierEquipmentID,omitempty"`
	PContractorNumberSlaughter string    `xml:"p_ContractorNumberSlaughter,omitempty"`
	PMFA                       int16     `xml:"p_MFA,omitempty"`
	PWeight                    float64   `xml:"p_Weight,omitempty"`
	PSlaughterInitiator        int32     `xml:"p_SlaughterInitiator,omitempty"`
}

type WritePigSlaughterMovementResponse struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigSlaughterMovementResponse"`

	WritePigSlaughterMovementResult *WriteMovementResult `xml:"WritePigSlaughterMovementResult,omitempty"`
}

type WriteGroupSlaughterMovement struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovement"`

	PManufacturerKey          string    `xml:"p_ManufacturerKey,omitempty"`
	PReportingTvdNumber       int32     `xml:"p_ReportingTvdNumber,omitempty"`
	PLCID                     int32     `xml:"p_LCID,omitempty"`
	PEventDate                time.Time `xml:"p_EventDate,omitempty"`
	PNumberOfAnimals          int32     `xml:"p_NumberOfAnimals,omitempty"`
	PSourceHusbandryTvdNumber int32     `xml:"p_SourceHusbandryTvdNumber,omitempty"`
}

type WriteGroupSlaughterMovementResponse struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovementResponse"`

	WriteGroupSlaughterMovementResult *WriteMovementResult `xml:"WriteGroupSlaughterMovementResult,omitempty"`
}

type WriteGroupSlaughterMovementV2 struct {
	XMLName                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovementV2"`

	PManufacturerKey             string    `xml:"p_ManufacturerKey,omitempty"`
	PReportingTvdNumber          int32     `xml:"p_ReportingTvdNumber,omitempty"`
	PLCID                        int32     `xml:"p_LCID,omitempty"`
	PEventDate                   time.Time `xml:"p_EventDate,omitempty"`
	PGenus                       int32     `xml:"p_Genus,omitempty"`
	PNumberOfAnimals             int32     `xml:"p_NumberOfAnimals,omitempty"`
	PSourceHusbandryTvdNumber    int32     `xml:"p_SourceHusbandryTvdNumber,omitempty"`
	PSlaughterInitiatorTvdNumber int32     `xml:"p_SlaughterInitiatorTvdNumber,omitempty"`
}

type WriteGroupSlaughterMovementV2Response struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovementV2Response"`

	WriteGroupSlaughterMovementV2Result *WriteMovementResult `xml:"WriteGroupSlaughterMovementV2Result,omitempty"`
}

type GetEquidOwnershipList struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidOwnershipList"`

	PManufacturerKey           string                     `xml:"p_ManufacturerKey,omitempty"`
	PLCID                      int32                      `xml:"p_LCID,omitempty"`
	PEquidOwnershipListRequest *EquidOwnershipListRequest `xml:"p_EquidOwnershipListRequest,omitempty"`
}

type GetEquidOwnershipListResponse struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidOwnershipListResponse"`

	GetEquidOwnershipListResult *GetEquidOwnershipListResult `xml:"GetEquidOwnershipListResult,omitempty"`
}

type GetEquidLivestock struct {
	XMLName                xml.Name `xml:"ns:GetEquidLivestock"`

	Action                 string                 `xml:"-"`
	PManufacturerKey       string                 `xml:"ns:p_ManufacturerKey,omitempty"`
	PLCID                  int32                  `xml:"ns:p_LCID,omitempty"`
	PEquidLivestockRequest *EquidLivestockRequest `xml:"ns:p_EquidLivestockRequest,omitempty"`
}

type GetEquidLivestockResponse struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidLivestockResponse"`

	GetEquidLivestockResult *GetEquidLivestockResult `xml:"GetEquidLivestockResult,omitempty"`
}

type WriteEquidRelocationNotification struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidRelocationNotification"`

	PManufacturerKey        string                  `xml:"p_ManufacturerKey,omitempty"`
	PLCID                   int32                   `xml:"p_LCID,omitempty"`
	PEquidRelocationRequest *EquidRelocationRequest `xml:"p_EquidRelocationRequest,omitempty"`
}

type WriteEquidRelocationNotificationResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidRelocationNotificationResponse"`

	WriteEquidRelocationNotificationResult *WriteNotificationResult `xml:"WriteEquidRelocationNotificationResult,omitempty"`
}

type WriteEquidAcquireOwnershipNotification struct {
	XMLName                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidAcquireOwnershipNotification"`

	PManufacturerKey              string                        `xml:"p_ManufacturerKey,omitempty"`
	PLCID                         int32                         `xml:"p_LCID,omitempty"`
	PEquidAcquireOwnershipRequest *EquidAcquireOwnershipRequest `xml:"p_EquidAcquireOwnershipRequest,omitempty"`
}

type WriteEquidAcquireOwnershipNotificationResponse struct {
	XMLName                                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidAcquireOwnershipNotificationResponse"`

	WriteEquidAcquireOwnershipNotificationResult *WriteNotificationResult `xml:"WriteEquidAcquireOwnershipNotificationResult,omitempty"`
}

type WriteEquidCedeOwnershipNotification struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidCedeOwnershipNotification"`

	PManufacturerKey           string                     `xml:"p_ManufacturerKey,omitempty"`
	PLCID                      int32                      `xml:"p_LCID,omitempty"`
	PEquidCedeOwnershipRequest *EquidCedeOwnershipRequest `xml:"p_EquidCedeOwnershipRequest,omitempty"`
}

type WriteEquidCedeOwnershipNotificationResponse struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidCedeOwnershipNotificationResponse"`

	WriteEquidCedeOwnershipNotificationResult *WriteNotificationResult `xml:"WriteEquidCedeOwnershipNotificationResult,omitempty"`
}

type GetFarmers struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmers"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
}

type GetFarmersResponse struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmersResponse"`

	GetFarmersResult *PersonListResult `xml:"GetFarmersResult,omitempty"`
}

type GetFarmManager struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmManager"`

	Action           string `xml:"-"`
	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
}

type GetFarmManagerResponse struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmManagerResponse"`

	GetFarmManagerResult *FarmManagerResult `xml:"GetFarmManagerResult,omitempty"`
}

type GetAnimalHusbandryAddress struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddress"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
}

type GetAnimalHusbandryAddressResponse struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddressResponse"`

	GetAnimalHusbandryAddressResult *AnimalHusbandryAddressResult `xml:"GetAnimalHusbandryAddressResult,omitempty"`
}

type GetAnimalHusbandryAddressV2 struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddressV2"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
}

type GetAnimalHusbandryAddressV2Response struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddressV2Response"`

	GetAnimalHusbandryAddressV2Result *AnimalHusbandryAddressResultV2 `xml:"GetAnimalHusbandryAddressV2Result,omitempty"`
}

type GetCodes struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCodes"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PType            string `xml:"p_Type,omitempty"`
}

type GetCodesResponse struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCodesResponse"`

	GetCodesResult *CodesResult `xml:"GetCodesResult,omitempty"`
}

type GetPersonAddress struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonAddress"`

	Action           string `xml:"-"`
	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PAgateNumber     string `xml:"p_AgateNumber,omitempty"`
}

type GetPersonAddressResponse struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonAddressResponse" json:"-"`

	GetPersonAddressResult *PersonAddressResult `xml:"GetPersonAddressResult,omitempty"`
}

type GetPersonIdentifiers struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonIdentifiers"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
}

type GetPersonIdentifiersResponse struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonIdentifiersResponse"`

	GetPersonIdentifiersResult *PersonIdentifiersResult `xml:"GetPersonIdentifiersResult,omitempty"`
}

type GetRegisteredGenera struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetRegisteredGenera"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
}

type GetRegisteredGeneraResponse struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetRegisteredGeneraResponse"`

	GetRegisteredGeneraResult *GeneraResult `xml:"GetRegisteredGeneraResult,omitempty"`
}

type WriteCattleArrivalBatchNotification struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleArrivalBatchNotification"`

	PManufacturerKey string                  `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32                   `xml:"p_LCID,omitempty"`
	PTVDNumber       int32                   `xml:"p_TVDNumber,omitempty"`
	PArrivalData     *CattleArrivalDataArray `xml:"p_ArrivalData,omitempty"`
}

type WriteCattleArrivalBatchNotificationResponse struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleArrivalBatchNotificationResponse"`

	WriteCattleArrivalBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleArrivalBatchNotificationResult,omitempty"`
}

type WriteCattleDeceasedNotification struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeceasedNotification"`

	PCattleDeceasedRequest *CattleNotificationRequest `xml:"p_CattleDeceasedRequest,omitempty"`
}

type WriteCattleDeceasedNotificationResponse struct {
	XMLName                               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeceasedNotificationResponse"`

	WriteCattleDeceasedNotificationResult *CattleNotificationResult `xml:"WriteCattleDeceasedNotificationResult,omitempty"`
}

type WriteCattleYardSlaughterNotification struct {
	XMLName                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleYardSlaughterNotification"`

	PCattleYardSlaughteredRequest *CattleNotificationRequest `xml:"p_CattleYardSlaughteredRequest,omitempty"`
}

type WriteCattleYardSlaughterNotificationResponse struct {
	XMLName                                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleYardSlaughterNotificationResponse"`

	WriteCattleYardSlaughterNotificationResult *CattleNotificationResult `xml:"WriteCattleYardSlaughterNotificationResult,omitempty"`
}

type WriteCattleExportNotification struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleExportNotification"`

	PCattleExportRequest *CattleNotificationWithCountryRequest `xml:"p_CattleExportRequest,omitempty"`
}

type WriteCattleExportNotificationResponse struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleExportNotificationResponse"`

	WriteCattleExportNotificationResult *CattleNotificationResult `xml:"WriteCattleExportNotificationResult,omitempty"`
}

type WriteCattleBirthNotification struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleBirthNotification"`

	PManufacturerKey string           `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32            `xml:"p_LCID,omitempty"`
	PTVDNumber       int32            `xml:"p_TVDNumber,omitempty"`
	PBirthData       *CattleBirthData `xml:"p_BirthData,omitempty"`
}

type WriteCattleBirthNotificationResponse struct {
	XMLName                            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleBirthNotificationResponse"`

	WriteCattleBirthNotificationResult *CattleNotificationResult `xml:"WriteCattleBirthNotificationResult,omitempty"`
}

type WriteCattleDaystayBatchNotification struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDaystayBatchNotification"`

	PManufacturerKey string                  `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32                   `xml:"p_LCID,omitempty"`
	PTVDNumber       int32                   `xml:"p_TVDNumber,omitempty"`
	PDaystayData     *CattleDaystayDataArray `xml:"p_DaystayData,omitempty"`
}

type WriteCattleDaystayBatchNotificationResponse struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDaystayBatchNotificationResponse"`

	WriteCattleDaystayBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleDaystayBatchNotificationResult,omitempty"`
}

type WriteCattleLeavingBatchNotification struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleLeavingBatchNotification"`

	PManufacturerKey string                  `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32                   `xml:"p_LCID,omitempty"`
	PTVDNumber       int32                   `xml:"p_TVDNumber,omitempty"`
	PLeavingData     *CattleLeavingDataArray `xml:"p_LeavingData,omitempty"`
}

type WriteCattleLeavingBatchNotificationResponse struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleLeavingBatchNotificationResponse"`

	WriteCattleLeavingBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleLeavingBatchNotificationResult,omitempty"`
}

type GetAnimalHusbandryMemberships struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryMemberships"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
}

type GetAnimalHusbandryMembershipsResponse struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryMembershipsResponse"`

	GetAnimalHusbandryMembershipsResult *HusbandryMembershipResult `xml:"GetAnimalHusbandryMembershipsResult,omitempty"`
}

type GetEarTagOrders struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEarTagOrders"`

	PManufacturerKey string    `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32     `xml:"p_LCID,omitempty"`
	PTVDNumber       int32     `xml:"p_TVDNumber,omitempty"`
	PSearchDateFrom  time.Time `xml:"p_SearchDateFrom,omitempty"`
	PSearchDateTo    time.Time `xml:"p_SearchDateTo,omitempty"`
	PArticleFilter   *IntArray `xml:"p_ArticleFilter,omitempty"`
}

type GetEarTagOrdersResponse struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEarTagOrdersResponse"`

	GetEarTagOrdersResult *EarTagOrderResult `xml:"GetEarTagOrdersResult,omitempty"`
}

type WriteCattleDeformationNotification struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeformationNotification"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	PEarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
	PDeformationType int32  `xml:"p_DeformationType,omitempty"`
}

type WriteCattleDeformationNotificationResponse struct {
	XMLName                                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeformationNotificationResponse"`

	WriteCattleDeformationNotificationResult *CattleNotificationResult `xml:"WriteCattleDeformationNotificationResult,omitempty"`
}

type DeleteAnimalHusbandryMemberships struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteAnimalHusbandryMemberships"`

	PManufacturerKey      string    `xml:"p_ManufacturerKey,omitempty"`
	PLCID                 int32     `xml:"p_LCID,omitempty"`
	PIDPRG                int32     `xml:"p_ID_PRG,omitempty"`
	PDeleteMembershipData *IntArray `xml:"p_DeleteMembershipData,omitempty"`
}

type DeleteAnimalHusbandryMembershipsResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteAnimalHusbandryMembershipsResponse"`

	DeleteAnimalHusbandryMembershipsResult *DeleteAnimalHusbandryMembershipResult `xml:"DeleteAnimalHusbandryMembershipsResult,omitempty"`
}

type DeleteCattleNotifications struct {
	XMLName                            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteCattleNotifications"`

	PManufacturerKey                   string    `xml:"p_ManufacturerKey,omitempty"`
	PLCID                              int32     `xml:"p_LCID,omitempty"`
	PTVDNumber                         int32     `xml:"p_TVDNumber,omitempty"`
	PDeleteCattleNotificationDataArray *IntArray `xml:"p_DeleteCattleNotificationDataArray,omitempty"`
}

type DeleteCattleNotificationsResponse struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteCattleNotificationsResponse"`

	DeleteCattleNotificationsResult *WriteCattleBatchNotificationResult `xml:"DeleteCattleNotificationsResult,omitempty"`
}

type WriteCattleDeathBirthNotification struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeathBirthNotification"`

	PManufacturerKey string                `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32                 `xml:"p_LCID,omitempty"`
	PTVDNumber       int32                 `xml:"p_TVDNumber,omitempty"`
	PDeathBirthData  *CattleDeathBirthData `xml:"p_DeathBirthData,omitempty"`
}

type WriteCattleDeathBirthNotificationResponse struct {
	XMLName                                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeathBirthNotificationResponse"`

	WriteCattleDeathBirthNotificationResult *CattleDeathBirthNotificationResult `xml:"WriteCattleDeathBirthNotificationResult,omitempty"`
}

type WriteCattleSlaughterBatchNotification struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotification"`

	PManufacturerKey string                  `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32                   `xml:"p_LCID,omitempty"`
	PTVDNumber       int32                   `xml:"p_TVDNumber,omitempty"`
	PSlaughterData   *CattleArrivalDataArray `xml:"p_SlaughterData,omitempty"`
}

type WriteCattleSlaughterBatchNotificationResponse struct {
	XMLName                                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotificationResponse"`

	WriteCattleSlaughterBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleSlaughterBatchNotificationResult,omitempty"`
}

type WriteCattleSlaughterBatchNotificationV2 struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotificationV2"`

	PManufacturerKey string                    `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32                     `xml:"p_LCID,omitempty"`
	PTVDNumber       int32                     `xml:"p_TVDNumber,omitempty"`
	PSlaughterData   *CattleSlaughterDataArray `xml:"p_SlaughterData,omitempty"`
}

type WriteCattleSlaughterBatchNotificationV2Response struct {
	XMLName                                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotificationV2Response"`

	WriteCattleSlaughterBatchNotificationV2Result *WriteCattleBatchNotificationResult `xml:"WriteCattleSlaughterBatchNotificationV2Result,omitempty"`
}

type GetCattleEarTags struct {
	//edited
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleEarTags"`

	Action           string `xml:"-"`
	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	POnStock         bool   `xml:"p_OnStock"`
}

type GetCattleEarTagsResponse struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleEarTagsResponse"`

	GetCattleEarTagsResult *CattleEarTagsResult `xml:"GetCattleEarTagsResult,omitempty"`
}

type GetCattleMovements struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleMovements"`

	PManufacturerKey string             `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32              `xml:"p_LCID,omitempty"`
	PWorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`
	PEarTagNumber    string             `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleMovementsResponse struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleMovementsResponse"`

	GetCattleMovementsResult *CattleMovementsResult `xml:"GetCattleMovementsResult,omitempty"`
}

type GetCattleHistory struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleHistory"`

	PManufacturerKey string             `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32              `xml:"p_LCID,omitempty"`
	PWorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`
	PEarTagNumber    string             `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleHistoryResponse struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleHistoryResponse"`

	GetCattleHistoryResult *CattleHistoryResult `xml:"GetCattleHistoryResult,omitempty"`
}

type GetCattleLivestock struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestock"`

	PManufacturerKey string             `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32              `xml:"p_LCID,omitempty"`
	PWorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`
	PTVDNumber       int32              `xml:"p_TVDNumber,omitempty"`
	PSearchDateFrom  time.Time          `xml:"p_SearchDateFrom,omitempty"`
	PSearchDateTo    time.Time          `xml:"p_SearchDateTo,omitempty"`
}

type GetCattleLivestockResponse struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestockResponse"`

	GetCattleLivestockResult *CattleLivestockResult `xml:"GetCattleLivestockResult,omitempty"`
}

type GetCattleLivestockV2 struct {
	//edited
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestockV2"`

	Action           string             `xml:"-"`
	PManufacturerKey string             `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32              `xml:"p_LCID,omitempty"`
	PWorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`
	PTVDNumber       int32              `xml:"p_TVDNumber,omitempty"`
	PSearchDateFrom  time.Time          `xml:"p_SearchDateFrom,omitempty"`
	PSearchDateTo    time.Time          `xml:"p_SearchDateTo,omitempty"`
}

type GetCattleLivestockV2Response struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestockV2Response"`

	GetCattleLivestockV2Result *CattleLivestockResultV2 `xml:"GetCattleLivestockV2Result,omitempty"`
}

type GetCattleDetail struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetail"`

	PManufacturerKey string             `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32              `xml:"p_LCID,omitempty"`
	PWorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`
	PEarTagNumber    string             `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleDetailResponse struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetailResponse"`

	GetCattleDetailResult *CattleDetailResult `xml:"GetCattleDetailResult,omitempty"`
}

type GetCattleDetailV2 struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetailV2"`

	Action           string `xml:"-"`

	PManufacturerKey string             `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32              `xml:"p_LCID,omitempty"`
	PWorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`
	PEarTagNumber    string             `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleDetailV2Response struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetailV2Response"`

	GetCattleDetailV2Result *CattleDetailResultV2 `xml:"GetCattleDetailV2Result,omitempty"`
}

type GetCattleStatus struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatus"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PEarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleStatusResponse struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatusResponse"`

	GetCattleStatusResult *CattleStateExternalResult `xml:"GetCattleStatusResult,omitempty"`
}

type GetCattleStatusV2 struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatusV2"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PEarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleStatusV2Response struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatusV2Response"`

	GetCattleStatusV2Result *CattleStateExternalResultV2 `xml:"GetCattleStatusV2Result,omitempty"`
}

type WriteNewEarTagOrder struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewEarTagOrder"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	PEarTagType      int32  `xml:"p_EarTagType,omitempty"`
	PAmount          int32  `xml:"p_Amount,omitempty"`
}

type WriteNewEarTagOrderResponse struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewEarTagOrderResponse"`

	WriteNewEarTagOrderResult *NewEarTagOrderResult `xml:"WriteNewEarTagOrderResult,omitempty"`
}

type WriteNewLabelEarTagOrder struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewLabelEarTagOrder"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	PLabelEarTagType int32  `xml:"p_LabelEarTagType,omitempty"`
	PAmount          int32  `xml:"p_Amount,omitempty"`
}

type WriteNewLabelEarTagOrderResponse struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewLabelEarTagOrderResponse"`

	WriteNewLabelEarTagOrderResult *NewEarTagOrderResult `xml:"WriteNewLabelEarTagOrderResult,omitempty"`
}

type CheckCattleForDisposalContribution struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattleForDisposalContribution"`

	PManufacturerKey string             `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32              `xml:"p_LCID,omitempty"`
	PTVDNumber       int32              `xml:"p_TVDNumber,omitempty"`
	PSlaughterData   *CattleArrivalData `xml:"p_SlaughterData,omitempty"`
}

type CheckCattleForDisposalContributionResponse struct {
	XMLName                                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattleForDisposalContributionResponse"`

	CheckCattleForDisposalContributionResult *DisposalContributionResult `xml:"CheckCattleForDisposalContributionResult,omitempty"`
}

type WriteAnimalClassificationNotification struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotification"`

	PManufacturerKey          string                    `xml:"p_ManufacturerKey,omitempty"`
	PLCID                     int32                     `xml:"p_LCID,omitempty"`
	PTVDNumber                int32                     `xml:"p_TVDNumber,omitempty"`
	PAnimalClassificationData *AnimalClassificationData `xml:"p_AnimalClassificationData,omitempty"`
}

type WriteAnimalClassificationNotificationResponse struct {
	XMLName                                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotificationResponse"`

	WriteAnimalClassificationNotificationResult *NotificationResult `xml:"WriteAnimalClassificationNotificationResult,omitempty"`
}

type WriteAnimalClassificationNotificationV2 struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotificationV2"`

	PManufacturerKey          string                      `xml:"p_ManufacturerKey,omitempty"`
	PLCID                     int32                       `xml:"p_LCID,omitempty"`
	PTVDNumber                int32                       `xml:"p_TVDNumber,omitempty"`
	PAnimalClassificationData *AnimalClassificationDataV2 `xml:"p_AnimalClassificationData,omitempty"`
}

type WriteAnimalClassificationNotificationV2Response struct {
	XMLName                                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotificationV2Response"`

	WriteAnimalClassificationNotificationV2Result *NotificationResult `xml:"WriteAnimalClassificationNotificationV2Result,omitempty"`
}

type EnableDataAccess struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EnableDataAccess"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	PIDPRG           int32  `xml:"p_ID_PRG,omitempty"`
}

type EnableDataAccessResponse struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EnableDataAccessResponse"`

	EnableDataAccessResult *DataAccessResult `xml:"EnableDataAccessResult,omitempty"`
}

type DisableDataAccess struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DisableDataAccess"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	PIDPRG           int32  `xml:"p_ID_PRG,omitempty"`
}

type DisableDataAccessResponse struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DisableDataAccessResponse"`

	DisableDataAccessResult *DataAccessResult `xml:"DisableDataAccessResult,omitempty"`
}

type CheckCattlesForPassport struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattlesForPassport"`

	PManufacturerKey string       `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32        `xml:"p_LCID,omitempty"`
	PIssueDate       time.Time    `xml:"p_IssueDate,omitempty"`
	PEarTagNumbers   *StringArray `xml:"p_EarTagNumbers,omitempty"`
}

type CheckCattlesForPassportResponse struct {
	XMLName                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattlesForPassportResponse"`

	CheckCattlesForPassportResult *CattlePassportResult `xml:"CheckCattlesForPassportResult,omitempty"`
}

type WriteCattlePassportOrders struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattlePassportOrders"`

	PManufacturerKey  string       `xml:"p_ManufacturerKey,omitempty"`
	PLCID             int32        `xml:"p_LCID,omitempty"`
	PPassportLanguage string       `xml:"p_PassportLanguage,omitempty"`
	PIssueDate        time.Time    `xml:"p_IssueDate,omitempty"`
	PEarTagNumbers    *StringArray `xml:"p_EarTagNumbers,omitempty"`
}

type WriteCattlePassportOrdersResponse struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattlePassportOrdersResponse"`

	WriteCattlePassportOrdersResult *CattlePassportResult `xml:"WriteCattlePassportOrdersResult,omitempty"`
}

type DeleteEarTagOrder struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteEarTagOrder"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	PNotificationID  int32  `xml:"p_NotificationID,omitempty"`
}

type DeleteEarTagOrderResponse struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteEarTagOrderResponse"`

	DeleteEarTagOrderResult *NotificationResult `xml:"DeleteEarTagOrderResult,omitempty"`
}

type WriteCattleChangeNameNotification struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleChangeNameNotification"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	PEarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
	PName            string `xml:"p_Name,omitempty"`
}

type WriteCattleChangeNameNotificationResponse struct {
	XMLName                                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleChangeNameNotificationResponse"`

	WriteCattleChangeNameNotificationResult *NotificationResult `xml:"WriteCattleChangeNameNotificationResult,omitempty"`
}

type GetLabelEarTagOrders struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetLabelEarTagOrders"`

	PManufacturerKey string    `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32     `xml:"p_LCID,omitempty"`
	PTVDNumber       int32     `xml:"p_TVDNumber,omitempty"`
	PSearchDateFrom  time.Time `xml:"p_SearchDateFrom,omitempty"`
	PSearchDateTo    time.Time `xml:"p_SearchDateTo,omitempty"`
	PArticleFilter   *IntArray `xml:"p_ArticleFilter,omitempty"`
}

type GetLabelEarTagOrdersResponse struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetLabelEarTagOrdersResponse"`

	GetLabelEarTagOrdersResult *EarTagOrderResult `xml:"GetLabelEarTagOrdersResult,omitempty"`
}

type WriteCattleTypeOfUseNotification struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleTypeOfUseNotification"`

	PManufacturerKey string    `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32     `xml:"p_LCID,omitempty"`
	PTVDNumber       int32     `xml:"p_TVDNumber,omitempty"`
	PEarTagNumber    string    `xml:"p_EarTagNumber,omitempty"`
	PCattleTypeOfUse int32     `xml:"p_CattleTypeOfUse,omitempty"`
	PEventDate       time.Time `xml:"p_EventDate,omitempty"`
}

type WriteCattleTypeOfUseNotificationResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleTypeOfUseNotificationResponse"`

	WriteCattleTypeOfUseNotificationResult *NotificationResult `xml:"WriteCattleTypeOfUseNotificationResult,omitempty"`
}

type DeleteLabelEarTagOrder struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteLabelEarTagOrder"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	PNotificationID  int32  `xml:"p_NotificationID,omitempty"`
}

type DeleteLabelEarTagOrderResponse struct {
	XMLName                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteLabelEarTagOrderResponse"`

	DeleteLabelEarTagOrderResult *NotificationResult `xml:"DeleteLabelEarTagOrderResult,omitempty"`
}

type WriteReplacementEarTagOrder struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementEarTagOrder"`

	PManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`
	PLCID            int32  `xml:"p_LCID,omitempty"`
	PTVDNumber       int32  `xml:"p_TVDNumber,omitempty"`
	PGenus           int32  `xml:"p_Genus,omitempty"`
	PEarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
	PLeftEarTag      bool   `xml:"p_LeftEarTag,omitempty"`
	PRightEarTag     bool   `xml:"p_RightEarTag,omitempty"`
}

type WriteReplacementEarTagOrderResponse struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementEarTagOrderResponse"`

	WriteReplacementEarTagOrderResult *NewEarTagOrderResult `xml:"WriteReplacementEarTagOrderResult,omitempty"`
}

type GetCattleOffsprings struct {
	XMLName             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleOffsprings"`

	PManufacturerKey    string             `xml:"p_ManufacturerKey,omitempty"`
	PLCID               int32              `xml:"p_LCID,omitempty"`
	PWorkingFocus       *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`
	PEarTagNumberMother string             `xml:"p_EarTagNumberMother,omitempty"`
}

type GetCattleOffspringsResponse struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleOffspringsResponse"`

	GetCattleOffspringsResult *CattleOffspringResult `xml:"GetCattleOffspringsResult,omitempty"`
}

type WritePoultrySlaughterNotification struct {
	XMLName             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultrySlaughterNotification"`

	PManufacturerKey    string           `xml:"p_ManufacturerKey,omitempty"`
	PLCID               int32            `xml:"p_LCID,omitempty"`
	PTVDNumber          int32            `xml:"p_TVDNumber,omitempty"`
	PEventDate          time.Time        `xml:"p_EventDate,omitempty"`
	PWeight             int32            `xml:"p_Weight,omitempty"`
	PPoultryType        *EnumPoultryType `xml:"p_PoultryType,omitempty"`
	PSourceTVDNumber    int32            `xml:"p_SourceTVDNumber,omitempty"`
	PDeliveryFileNumber string           `xml:"p_DeliveryFileNumber,omitempty"`
}

type WritePoultrySlaughterNotificationResponse struct {
	XMLName                                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultrySlaughterNotificationResponse"`

	WritePoultrySlaughterNotificationResult *ProcessingResult `xml:"WritePoultrySlaughterNotificationResult,omitempty"`
}

type WritePoultryBarnInNotification struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultryBarnInNotification"`

	PManufacturerKey           string                     `xml:"p_ManufacturerKey,omitempty"`
	PLCID                      int32                      `xml:"p_LCID,omitempty"`
	PTVDNumber                 int32                      `xml:"p_TVDNumber,omitempty"`
	PPoultryBarnInNotification *PoultryBarnInNotification `xml:"p_PoultryBarnInNotification,omitempty"`
}

type WritePoultryBarnInNotificationResponse struct {
	XMLName                              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultryBarnInNotificationResponse"`

	WritePoultryBarnInNotificationResult *ProcessingResult `xml:"WritePoultryBarnInNotificationResult,omitempty"`
}

type GetPoultryBarnInNotifications struct {
	XMLName                                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPoultryBarnInNotifications"`

	PSearchPoultryBarnInNotificationRequest *SearchPoultryBarnInNotificationRequest `xml:"p_SearchPoultryBarnInNotificationRequest,omitempty"`
}

type GetPoultryBarnInNotificationsResponse struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPoultryBarnInNotificationsResponse"`

	GetPoultryBarnInNotificationsResult *GetPoultryBarnInNotificationResult `xml:"GetPoultryBarnInNotificationsResult,omitempty"`
}

type GetAnimalHusbandryDetail struct {
	XMLName                          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryDetail"`

	PGetAnimalHusbandryDetailRequest *GetAnimalHusbandryDetailRequest `xml:"p_GetAnimalHusbandryDetailRequest,omitempty"`
}

type GetAnimalHusbandryDetailResponse struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryDetailResponse"`

	GetAnimalHusbandryDetailResult *GetAnimalHusbandryDetailResult `xml:"GetAnimalHusbandryDetailResult,omitempty"`
}

type GetPigArrivalNotificationForBreedingOrganisation struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPigArrivalNotificationForBreedingOrganisation"`

	PSearchSmallAnimalMovementRequest *SearchSmallAnimalMovementRequest `xml:"p_SearchSmallAnimalMovementRequest,omitempty"`
}

type GetPigArrivalNotificationForBreedingOrganisationResponse struct {
	XMLName                                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPigArrivalNotificationForBreedingOrganisationResponse"`

	GetPigArrivalNotificationForBreedingOrganisationResult *PigArrivalNotificationResult `xml:"GetPigArrivalNotificationForBreedingOrganisationResult,omitempty"`
}

type GetMembershipForOrganisation struct {
	XMLName  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetMembershipForOrganisation"`

	PRequest *GetMembershipForOrganisationRequest `xml:"p_Request,omitempty"`
}

type GetMembershipForOrganisationResponse struct {
	XMLName                            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetMembershipForOrganisationResponse"`

	GetMembershipForOrganisationResult *GetMembershipForOrganisationResult `xml:"GetMembershipForOrganisationResult,omitempty"`
}

type WriteReplacementBatchEarTagOrder struct {
	XMLName                                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementBatchEarTagOrder"`

	PManufacturerKey                         string                                   `xml:"p_ManufacturerKey,omitempty"`
	PLCID                                    int32                                    `xml:"p_LCID,omitempty"`
	PWriteReplacementBatchEarTagOrderRequest *WriteReplacementBatchEarTagOrderRequest `xml:"p_WriteReplacementBatchEarTagOrderRequest,omitempty"`
}

type WriteReplacementBatchEarTagOrderResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementBatchEarTagOrderResponse"`

	WriteReplacementBatchEarTagOrderResult *ReplacementEarTagOrdersResult `xml:"WriteReplacementBatchEarTagOrderResult,omitempty"`
}

type GetCattlesPerLeavingDate struct {
	XMLName                          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDate"`

	PManufacturerKey                 string                           `xml:"p_ManufacturerKey,omitempty"`
	PLCID                            int32                            `xml:"p_LCID,omitempty"`
	PGetCattlesPerLeavingDateRequest *GetCattlesPerLeavingDateRequest `xml:"p_GetCattlesPerLeavingDateRequest,omitempty"`
}

type GetCattlesPerLeavingDateResponse struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDateResponse"`

	GetCattlesPerLeavingDateResult *GetCattlesPerLeavingDateResult `xml:"GetCattlesPerLeavingDateResult,omitempty"`
}

type WritePigEntryMovementV2Request struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovementV2Request"`

	*BaseRequest

	ReportingTVDNumber       int32            `xml:"ReportingTVDNumber,omitempty"`
	EventDate                time.Time        `xml:"EventDate,omitempty"`
	NumberOfAnimals          int32            `xml:"NumberOfAnimals,omitempty"`
	SourceHusbandryTVDNumber int32            `xml:"SourceHusbandryTVDNumber,omitempty"`
	PigCategory              *EnumPigCategory `xml:"PigCategory,omitempty"`
}

type BaseRequest struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 BaseRequest"`

	ManufacturerKey string `xml:"ManufacturerKey,omitempty"`
	LCID            int32  `xml:"LCID,omitempty"`
}

type WriteMovementResult struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteMovementResult"`

	MovementId       int32             `xml:"MovementId,omitempty"`
	ProcessingResult *ProcessingResult `xml:"ProcessingResult,omitempty"`
}

type ProcessingResult struct {
	XMLName     xml.Name

	Code        int32  `xml:"Code,omitempty"`
	Description string `xml:"Description,omitempty"`
	Status      int32  `xml:"Status,omitempty"`
}

type EquidOwnershipListRequest struct {
	XMLName                                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidOwnershipListRequest"`

	OnBehalfOfAgateNumber                      string    `xml:"OnBehalfOfAgateNumber,omitempty"`
	SearchDateFrom                             time.Time `xml:"SearchDateFrom,omitempty"`
	SearchDateTo                               time.Time `xml:"SearchDateTo,omitempty"`
	ShowOnlyEquidsWhichWereLivingInQueryPeriod bool      `xml:"ShowOnlyEquidsWhichWereLivingInQueryPeriod,omitempty"`
}

type GetEquidOwnershipListResult struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidOwnershipListResult"`

	Result    *ProcessingResult `xml:"Result,omitempty"`
	EquidList *ArrayOfEquidItem `xml:"EquidList,omitempty"`
}

type ArrayOfEquidItem struct {
	XMLName   xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidItem"`

	EquidItem []*EquidItem `xml:"EquidItem,omitempty"`
}

type EquidItem struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidItem"`

	OwnerFirstName    string                                          `xml:"OwnerFirstName,omitempty"`
	OwnerLastName     string                                          `xml:"OwnerLastName,omitempty"`
	OwnerAgateNumber  string                                          `xml:"OwnerAgateNumber,omitempty"`
	UELN              string                                          `xml:"UELN,omitempty"`
	Name              string                                          `xml:"Name,omitempty"`
	OriginTVDNumber   string                                          `xml:"OriginTVDNumber,omitempty"`
	Husbandry         *HusbandryResult                                `xml:"Husbandry,omitempty"`
	NotificationState *TranslatedEnumTypeOfEnumEquidNotificationState `xml:"NotificationState,omitempty"`
	TypeOfUse         *TranslatedEnumTypeOfEnumEquidTypeOfUse         `xml:"TypeOfUse,omitempty"`
	Gender            *TranslatedEnumTypeOfEnumGender                 `xml:"Gender,omitempty"`
	Breed             *TranslatedEnumTypeOfEnumEquidBreed             `xml:"Breed,omitempty"`
	WithersClass      *TranslatedEnumTypeOfEnumEquidWithersClass      `xml:"WithersClass,omitempty"`
	IsPassPresent     bool                                            `xml:"IsPassPresent,omitempty"`
	ColorFreeText     string                                          `xml:"ColorFreeText,omitempty"`
	BirthDate         string                                          `xml:"BirthDate,omitempty"`
	DeathDate         string                                          `xml:"DeathDate,omitempty"`
	ArrivalDate       string                                          `xml:"ArrivalDate,omitempty"`
	LeavingDate       string                                          `xml:"LeavingDate,omitempty"`
}

type HusbandryResult struct {
	XMLName   xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryResult"`

	TVDNumber string `xml:"TVDNumber,omitempty"`
	Name      string `xml:"Name,omitempty"`
	Street    string `xml:"Street,omitempty"`
	PostCode  string `xml:"PostCode,omitempty"`
	City      string `xml:"City,omitempty"`
}

type EquidLivestockRequest struct {
	XMLName        xml.Name

	TVDNumber      int32     `xml:"ns:TVDNumber,omitempty"`
	SearchDateFrom time.Time `xml:"ns:SearchDateFrom,omitempty"`
	SearchDateTo   time.Time `xml:"ns:SearchDateTo,omitempty"`
}

type GetEquidLivestockResult struct {
	XMLName   xml.Name

	Result    *ProcessingResult `xml:"Result,omitempty"`
	EquidList *ArrayOfEquidItem `xml:"EquidList,omitempty"`
}

type EquidRelocationRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidRelocationRequest"`

	OnBehalfOfAgateNumber string                       `xml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                  string                       `xml:"Ueln,omitempty"`
	EventDate             time.Time                    `xml:"EventDate,omitempty"`
	LocationChangeType    *EnumEquidLocationChangeType `xml:"LocationChangeType,omitempty"`
	NewTVDNumber          int32                        `xml:"NewTVDNumber,omitempty"`
	OriginTVDNumber       int32                        `xml:"OriginTVDNumber,omitempty"`
	ImportOrExportCountry *EnumCountry                 `xml:"ImportOrExportCountry,omitempty"`
}

type WriteNotificationResult struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNotificationResult"`

	Result  *ProcessingResult `xml:"Result,omitempty"`
}

type EquidAcquireOwnershipRequest struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidAcquireOwnershipRequest"`

	OnBehalfOfAgateNumber  string    `xml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                   string    `xml:"Ueln,omitempty"`
	EventDate              time.Time `xml:"EventDate,omitempty"`
	ActualOwnerAgateNumber string    `xml:"ActualOwnerAgateNumber,omitempty"`
}

type EquidCedeOwnershipRequest struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidCedeOwnershipRequest"`

	OnBehalfOfAgateNumber       string    `xml:"OnBehalfOfAgateNumber,omitempty"`
	Ueln                        string    `xml:"Ueln,omitempty"`
	EventDate                   time.Time `xml:"EventDate,omitempty"`
	IsNewEquidOwnerLivingAbroad bool      `xml:"IsNewEquidOwnerLivingAbroad,omitempty"`
	CededToPersonAgateNumber    string    `xml:"CededToPersonAgateNumber,omitempty"`
}

type PersonListResult struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonListResult"`

	Result          *ProcessingResult `xml:"Result,omitempty"`
	PersonDataItems *PersonDataArray  `xml:"PersonDataItems,omitempty"`
}

type PersonDataArray struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonDataArray"`

	PersonDataItem []*PersonResult `xml:"PersonDataItem,omitempty" json:"-"`
}

type PersonResult struct {
	XMLName               xml.Name `json:"-"` //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonResult"`

	AgateNumber           string       `xml:"AgateNumber,omitempty" json:"agateNumber"`
	LastName              string       `xml:"LastName,omitempty" json:"lastName"`
	FirstName             string       `xml:"FirstName,omitempty" json:"firstName"`
	Street                string       `xml:"Street,omitempty" json:"street"`
	PostCode              string       `xml:"PostCode,omitempty" json:"postCode"`
	City                  string       `xml:"City,omitempty" json:"city"`
	EmailAddress          string       `xml:"EmailAddress,omitempty" json:"emailAddress"`
	PhoneNumbers          *StringArray `xml:"PhoneNumbers,omitempty" json:"phoneNumbers"`
	PreferredLanguageLCID int32        `xml:"PreferredLanguageLCID,omitempty" json:"preferredLanguageLcid"`
}

type StringArray struct {
	XMLName    xml.Name `json:"-"` //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 StringArray"`

	StringItem []string `xml:"StringItem,omitempty" json:"stringItem"`
}

type FarmManagerResult struct {
	XMLName               xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 FarmManagerResult"`

	AgateNumber           string            `xml:"AgateNumber,omitempty"`
	Title                 string            `xml:"Title,omitempty"`
	LastName              string            `xml:"LastName,omitempty"`
	FirstName             string            `xml:"FirstName,omitempty"`
	Street                string            `xml:"Street,omitempty"`
	PostCode              string            `xml:"PostCode,omitempty"`
	City                  string            `xml:"City,omitempty"`
	EmailAddress          string            `xml:"EmailAddress,omitempty"`
	PhoneNumbers          *StringArray      `xml:"PhoneNumbers,omitempty"`
	PreferredLanguageLCID int32             `xml:"PreferredLanguageLCID,omitempty"`
	Result                *ProcessingResult `xml:"Result,omitempty"`
}

type AnimalHusbandryAddressResult struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 AnimalHusbandryAddressResult"`

	Result                   *ProcessingResult `xml:"Result,omitempty"`
	PostData                 *HusbandryResult  `xml:"PostData,omitempty"`
	IsActive                 bool              `xml:"IsActive,omitempty"`
	MunicipalityName         string            `xml:"MunicipalityName,omitempty"`
	CantonShortname          string            `xml:"CantonShortname,omitempty"`
	CoordinateX              int32             `xml:"CoordinateX,omitempty"`
	CoordinateY              int32             `xml:"CoordinateY,omitempty"`
	Altitude                 int32             `xml:"Altitude,omitempty"`
	CantonAnimalHusbandryKey string            `xml:"CantonAnimalHusbandryKey,omitempty"`
	CantonPersonKey          string            `xml:"CantonPersonKey,omitempty"`
	BurNumber                string            `xml:"BurNumber,omitempty"`
	AnimalHusbandryType      int32             `xml:"AnimalHusbandryType,omitempty"`
	MunicipalityNumber       int32             `xml:"MunicipalityNumber,omitempty"`
	TypeOfUse                int32             `xml:"TypeOfUse,omitempty"`
}

type AnimalHusbandryAddressResultV2 struct {
	XMLName             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 AnimalHusbandryAddressResultV2"`

	Result              *ProcessingResult `xml:"Result,omitempty"`
	PostData            *HusbandryResult  `xml:"PostData,omitempty"`
	IsActive            bool              `xml:"IsActive,omitempty"`
	MunicipalityName    string            `xml:"MunicipalityName,omitempty"`
	CantonShortname     string            `xml:"CantonShortname,omitempty"`
	CoordinateX         int32             `xml:"CoordinateX,omitempty"`
	CoordinateY         int32             `xml:"CoordinateY,omitempty"`
	Altitude                 int32             `xml:"Altitude,omitempty"`
	CantonAnimalHusbandryKey string            `xml:"CantonAnimalHusbandryKey,omitempty"`
	CantonPersonKey          string            `xml:"CantonPersonKey,omitempty"`
	BurNumber           string            `xml:"BurNumber,omitempty"`
	AnimalHusbandryType int32             `xml:"AnimalHusbandryType,omitempty"`
	MunicipalityNumber  int32             `xml:"MunicipalityNumber,omitempty"`
	TypeOfUse           int32             `xml:"TypeOfUse,omitempty"`
	IsMountain          bool              `xml:"IsMountain,omitempty"`
}

type CodesResult struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CodesResult"`

	Result  *ProcessingResult `xml:"Result,omitempty"`
	Codes   *CodeArray        `xml:"Codes,omitempty"`
}

type CodeArray struct {
	XMLName  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CodeArray"`

	CodeItem []*Code `xml:"CodeItem,omitempty"`
}

type Code struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 Code"`

	Key     int32  `xml:"Key,omitempty"`
	Textde  string `xml:"Text_de,omitempty"`
	Textfr  string `xml:"Text_fr,omitempty"`
	Textit  string `xml:"Text_it,omitempty"`
}

type PersonAddressResult struct {
	XMLName     xml.Name `json:"-"` //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonAddressResult"`

	Result      *ProcessingResult `xml:"Result,omitempty" json:"-"`
	Title       string            `xml:"Title,omitempty" json:"title"`
	PostAddress *PersonResult     `xml:"PostAddress,omitempty" json:"postAddress"`
}

type PersonIdentifiersResult struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonIdentifiersResult"`

	Result          *ProcessingResult `xml:"Result,omitempty"`
	CantonPersonKey string            `xml:"CantonPersonKey,omitempty"`
	IDAGIS          int32             `xml:"ID_AGIS,omitempty"`
}

type GeneraResult struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GeneraResult"`

	Result  *ProcessingResult `xml:"Result,omitempty"`
	Genera  *IntArray         `xml:"Genera,omitempty"`
}

type IntArray struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 IntArray"`

	IntItem []int32 `xml:"IntItem,omitempty"`
}

type CattleArrivalDataArray struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleArrivalDataArray"`

	CattleArrivalDataItem []*CattleArrivalData `xml:"CattleArrivalDataItem,omitempty"`
}

type CattleArrivalData struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleArrivalData"`

	EarTagNumber    string    `xml:"EarTagNumber,omitempty"`
	EventDate       time.Time `xml:"EventDate,omitempty"`
	TVDNumberOrigin int32     `xml:"TVDNumberOrigin,omitempty"`
}

type WriteCattleBatchNotificationResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleBatchNotificationResult"`

	Result        *ProcessingResult              `xml:"Result,omitempty"`
	Resultdetails *CattleNotificationResultArray `xml:"Resultdetails,omitempty"`
}

type CattleNotificationResultArray struct {
	XMLName                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleNotificationResultArray"`

	CattleNotificationResultItem []*CattleNotificationResult `xml:"CattleNotificationResultItem,omitempty"`
}

type CattleNotificationResult struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleNotificationResult"`

	EarTagNumber   string            `xml:"EarTagNumber,omitempty"`
	NotificationID int32             `xml:"NotificationID,omitempty"`
	Result         *ProcessingResult `xml:"Result,omitempty"`
}

type CattleNotificationRequest struct {
	XMLName      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleNotificationRequest"`

	*BaseRequest

	TVDNumber    int32     `xml:"TVDNumber,omitempty"`
	EarTagNumber string    `xml:"EarTagNumber,omitempty"`
	EventDate    time.Time `xml:"EventDate,omitempty"`
}

type CattleNotificationWithCountryRequest struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleNotificationWithCountryRequest"`

	*CattleNotificationRequest

	Country *EnumCountry `xml:"Country,omitempty"`
}

type CattleBirthData struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleBirthData"`

	EarTagNumber              string    `xml:"EarTagNumber,omitempty"`
	IsMultipleBirth           bool      `xml:"IsMultipleBirth,omitempty"`
	RaceColor                 int32     `xml:"RaceColor,omitempty"`
	Gender                    int32     `xml:"Gender,omitempty"`
	BirthDate                 time.Time `xml:"BirthDate,omitempty"`
	Race                      int32     `xml:"Race,omitempty"`
	EarTagNumberFather        string    `xml:"EarTagNumberFather,omitempty"`
	EarTagNumberMother        string    `xml:"EarTagNumberMother,omitempty"`
	BreedingOrganisation      int32     `xml:"BreedingOrganisation,omitempty"`
	Name                      string    `xml:"Name,omitempty"`
	BirthProcess              int32     `xml:"BirthProcess,omitempty"`
	IsCastrated               bool      `xml:"IsCastrated,omitempty"`
	InseminationDate          time.Time `xml:"InseminationDate,omitempty"`
	EarTagNumberGeneticMother string    `xml:"EarTagNumberGeneticMother,omitempty"`
	BirthWeight               int32     `xml:"BirthWeight,omitempty"`
	IsPassportDesired         bool      `xml:"IsPassportDesired,omitempty"`
}

type CattleDaystayDataArray struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDaystayDataArray"`

	CattleDaystayDataItem []*CattleDaystayData `xml:"CattleDaystayDataItem,omitempty"`
}

type CattleDaystayData struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDaystayData"`

	EarTagNumber    string    `xml:"EarTagNumber,omitempty"`
	EventDate       time.Time `xml:"EventDate,omitempty"`
	TVDNumberOrigin int32     `xml:"TVDNumberOrigin,omitempty"`
}

type CattleLeavingDataArray struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLeavingDataArray"`

	CattleLeavingDataItem []*CattleLeavingData `xml:"CattleLeavingDataItem,omitempty"`
}

type CattleLeavingData struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLeavingData"`

	EarTagNumber                string    `xml:"EarTagNumber,omitempty"`
	EventDate                   time.Time `xml:"EventDate,omitempty"`
	LeavingReason               int32     `xml:"LeavingReason,omitempty"`
	MainLeavingReasonBreeding   int32     `xml:"MainLeavingReasonBreeding,omitempty"`
	SecondLeavingReasonBreeding int32     `xml:"SecondLeavingReasonBreeding,omitempty"`
}

type HusbandryMembershipResult struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryMembershipResult"`

	Result      *ProcessingResult         `xml:"Result,omitempty"`
	Memberships *HusbandryMembershipArray `xml:"Memberships,omitempty"`
}

type HusbandryMembershipArray struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryMembershipArray"`

	HusbandryMembershipItem []*HusbandryMembership `xml:"HusbandryMembershipItem,omitempty"`
}

type HusbandryMembership struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryMembership"`

	IDPRG       int32  `xml:"ID_PRG,omitempty"`
	AgateNumber string `xml:"AgateNumber,omitempty"`
	LastName    string `xml:"LastName,omitempty"`
	FirstName   string `xml:"FirstName,omitempty"`
	Role        string `xml:"Role,omitempty"`
	Genus       int32  `xml:"Genus,omitempty"`
}

type EarTagOrderResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EarTagOrderResult"`

	Result        *ProcessingResult     `xml:"Result,omitempty"`
	Resultdetails *EarTagOrderDataArray `xml:"Resultdetails,omitempty"`
}

type EarTagOrderDataArray struct {
	XMLName             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EarTagOrderDataArray"`

	EarTagOrderDataItem []*EarTagOrderData `xml:"EarTagOrderDataItem,omitempty"`
}

type EarTagOrderData struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EarTagOrderData"`

	NotificationID   int32     `xml:"NotificationID,omitempty"`
	EarTagType       int32     `xml:"EarTagType,omitempty"`
	Amount           int32     `xml:"Amount,omitempty"`
	IsExpress        bool      `xml:"IsExpress,omitempty"`
	OrderStatus      int32     `xml:"OrderStatus,omitempty"`
	OrderStatusDate  time.Time `xml:"OrderStatusDate,omitempty"`
	EarTagNumberFrom string    `xml:"EarTagNumberFrom,omitempty"`
	EarTagNumberTo   string    `xml:"EarTagNumberTo,omitempty"`
	Text1            string    `xml:"Text1,omitempty"`
	Text2            string    `xml:"Text2,omitempty"`
}

type DeleteAnimalHusbandryMembershipResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteAnimalHusbandryMembershipResult"`

	Result        *ProcessingResult                 `xml:"Result,omitempty"`
	Resultdetails *HusbandryNotificationResultArray `xml:"Resultdetails,omitempty"`
}

type HusbandryNotificationResultArray struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryNotificationResultArray"`

	HusbandryNotificationResultItem []*HusbandryNotificationResult `xml:"HusbandryNotificationResultItem,omitempty"`
}

type HusbandryNotificationResult struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryNotificationResult"`

	TVDNumber      int32             `xml:"TVDNumber,omitempty"`
	NotificationID int32             `xml:"NotificationID,omitempty"`
	Result         *ProcessingResult `xml:"Result,omitempty"`
}

type CattleDeathBirthData struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDeathBirthData"`

	DeathBirthTime            int32     `xml:"DeathBirthTime,omitempty"`
	IsMultipleBirth           bool      `xml:"IsMultipleBirth,omitempty"`
	RaceColor                 int32     `xml:"RaceColor,omitempty"`
	Gender                    int32     `xml:"Gender,omitempty"`
	BirthDate                 time.Time `xml:"BirthDate,omitempty"`
	Race                      int32     `xml:"Race,omitempty"`
	EarTagNumberFather        string    `xml:"EarTagNumberFather,omitempty"`
	EarTagNumberMother        string    `xml:"EarTagNumberMother,omitempty"`
	BreedingOrganisation      int32     `xml:"BreedingOrganisation,omitempty"`
	Name                      string    `xml:"Name,omitempty"`
	BirthProcess              int32     `xml:"BirthProcess,omitempty"`
	IsCastrated               bool      `xml:"IsCastrated,omitempty"`
	InseminationDate          time.Time `xml:"InseminationDate,omitempty"`
	EarTagNumberGeneticMother string    `xml:"EarTagNumberGeneticMother,omitempty"`
	BirthWeight               int32     `xml:"BirthWeight,omitempty"`
	IsPassportDesired         bool      `xml:"IsPassportDesired,omitempty"`
}

type CattleDeathBirthNotificationResult struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDeathBirthNotificationResult"`

	NotificationID int32             `xml:"NotificationID,omitempty"`
	Result         *ProcessingResult `xml:"Result,omitempty"`
}

type CattleSlaughterDataArray struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleSlaughterDataArray"`

	CattleSlaughterDataItem []*CattleSlaughterData `xml:"CattleSlaughterDataItem,omitempty"`
}

type CattleSlaughterData struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleSlaughterData"`

	EarTagNumber                string    `xml:"EarTagNumber,omitempty"`
	EventDate                   time.Time `xml:"EventDate,omitempty"`
	TVDNumberOrigin             int32     `xml:"TVDNumberOrigin,omitempty"`
	TVDNumberSlaughterInitiator int32     `xml:"TVDNumberSlaughterInitiator,omitempty"`
}

type CattleEarTagsResult struct {
	XMLName       xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleEarTagsResult"`

	Result        *ProcessingResult      `xml:"Result,omitempty"`
	CattleEarTags *CattleEarTagDataArray `xml:"CattleEarTags,omitempty"`
}

type CattleEarTagDataArray struct {
	XMLName              xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleEarTagDataArray"`

	CattleEarTagDataItem []*CattleEarTagData `xml:"CattleEarTagDataItem,omitempty"`
}

type CattleEarTagData struct {
	XMLName      xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleEarTagData"`

	EarTagNumber string `xml:"EarTagNumber,omitempty"`
	OrderDate    string `xml:"OrderDate,omitempty"`
	EarTagState  int32  `xml:"EarTagState,omitempty"`
}

type WorkingFocusArray struct {
	XMLName          xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WorkingFocusArray"`

	WorkingFocusItem *[]WorkingFocus `xml:"WorkingFocusItem,omitempty"`
}

type WorkingFocus struct {
	XMLName          xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WorkingFocus"`

	WorkingFocusType int32  `xml:"WorkingFocusType,omitempty"`
	TVDNumber        int32  `xml:"TVDNumber,omitempty"`
	MandateGiver     string `xml:"MandateGiver,omitempty"`
}

type CattleMovementsResult struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleMovementsResult"`

	Result             *ProcessingResult        `xml:"Result,omitempty"`
	CattleHistoryState int32                    `xml:"CattleHistoryState,omitempty"`
	EarTagNumber       string                   `xml:"EarTagNumber,omitempty"`
	CattleMovements    *CattleMovementDataArray `xml:"CattleMovements,omitempty"`
}

type CattleMovementDataArray struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleMovementDataArray"`

	CattleMovementDataItem []*CattleMovementData `xml:"CattleMovementDataItem,omitempty"`
}

type CattleMovementData struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleMovementData"`

	NotificationID    int32     `xml:"NotificationID,omitempty"`
	EventDate         time.Time `xml:"EventDate,omitempty"`
	NotificationDate  time.Time `xml:"NotificationDate,omitempty"`
	NotificationType  int32     `xml:"NotificationType,omitempty"`
	TVDNumberNotifier int32     `xml:"TVDNumberNotifier,omitempty"`
	TVDNumberOrigin   int32     `xml:"TVDNumberOrigin,omitempty"`
}

type CattleHistoryResult struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleHistoryResult"`

	Result             *ProcessingResult    `xml:"Result,omitempty"`
	EarTagNumber       string               `xml:"EarTagNumber,omitempty"`
	RemainingQuota     int32                `xml:"RemainingQuota,omitempty"`
	CattleHistoryState int32                `xml:"CattleHistoryState,omitempty"`
	CattleStays        *CattleStayDataArray `xml:"CattleStays,omitempty"`
}

type CattleStayDataArray struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStayDataArray"`

	CattleStayDataItem []*CattleStayData `xml:"CattleStayDataItem,omitempty"`
}

type CattleStayData struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStayData"`

	ArrivalDate             time.Time `xml:"ArrivalDate,omitempty"`
	ArrivalNotificationDate time.Time `xml:"ArrivalNotificationDate,omitempty"`
	ArrivalNotificationType int32     `xml:"ArrivalNotificationType,omitempty"`
	CountryOrigin           string    `xml:"CountryOrigin,omitempty"`
	TVDNumberOrigin         int32     `xml:"TVDNumberOrigin,omitempty"`
	TVDNumberStay           int32     `xml:"TVDNumberStay,omitempty"`
	StayAddress             string    `xml:"StayAddress,omitempty"`
	LeavingDate             time.Time `xml:"LeavingDate,omitempty"`
	LeavingNotificationDate time.Time `xml:"LeavingNotificationDate,omitempty"`
	LeavingNotificationType int32     `xml:"LeavingNotificationType,omitempty"`
	CattleStayState         int32     `xml:"CattleStayState,omitempty"`
}

type CattleLivestockResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockResult"`

	Result        *ProcessingResult         `xml:"Result,omitempty"`
	TVDNumber     int32                     `xml:"TVDNumber,omitempty"`
	Resultdetails *CattleLivestockDataArray `xml:"Resultdetails,omitempty"`
}

type CattleLivestockDataArray struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockDataArray"`

	CattleLivestockDataItem []*CattleLivestockData `xml:"CattleLivestockDataItem,omitempty"`
}

type CattleLivestockData struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockData"`

	EarTagNumber       string    `xml:"EarTagNumber,omitempty"`
	Name               string    `xml:"Name,omitempty"`
	Gender             int32     `xml:"Gender,omitempty"`
	BirthDate          time.Time `xml:"BirthDate,omitempty"`
	Race               int32     `xml:"Race,omitempty"`
	RaceColor          int32     `xml:"RaceColor,omitempty"`
	ArrivalDate        time.Time `xml:"ArrivalDate,omitempty"`
	LeavingDate        time.Time `xml:"LeavingDate,omitempty"`
	PendulumHusbandry  int32     `xml:"PendulumHusbandry,omitempty"`
	BvdState           string    `xml:"BvdState,omitempty"`
	BtState            string    `xml:"BtState,omitempty"`
	CattleTypeOfUse    int32     `xml:"CattleTypeOfUse,omitempty"`
	TypeOfUseDate      time.Time `xml:"TypeOfUseDate,omitempty"`
	FirstCalvingDate   time.Time `xml:"FirstCalvingDate,omitempty"`
	CattleHistoryState int32     `xml:"CattleHistoryState,omitempty"`
}

type CattleLivestockResultV2 struct {
	XMLName       xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockResultV2"`

	Result        *ProcessingResult           `xml:"Result,omitempty"`
	TVDNumber     int32                       `xml:"TVDNumber,omitempty"`
	Resultdetails *CattleLivestockDataArrayV2 `xml:"Resultdetails,omitempty"`
}

type CattleLivestockDataArrayV2 struct {
	XMLName                 xml.Name                 //`xml:"-"`//`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockDataArrayV2"`

	CattleLivestockDataItem []*CattleLivestockDataV2 //`xml:"CattleLivestockDataItem,omitempty"`
}

type CattleLivestockDataV2 struct {
					       //edited
	XMLName            xml.Name `json:"-"` //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockDataV2"`

	EarTagNumber       string `xml:"EarTagNumber,omitempty" json:"earTagNumber"`
	Name               string `xml:"Name,omitempty" json:"name"`
	Gender             string `xml:"Gender,omitempty" json:"gender"`
	BirthDate          string `xml:"BirthDate,omitempty" json:"birthDay"`
	Race               string `xml:"Race,omitempty" json:"race"`
	RaceColor          string `xml:"RaceColor,omitempty" json:"raceColor"`
	ArrivalDate        string `xml:"ArrivalDate,omitempty" json:"arrivalDate"`
	LeavingDate        string `xml:"LeavingDate,omitempty" json:"leavingDate"`
	PendulumHusbandry  string `xml:"PendulumHusbandry,omitempty" json:"pendulumHusbandry"`
	BvdState           string `xml:"BvdState,omitempty" json:"bvdState"`
	BtState            string `xml:"BtState,omitempty" json:"btState"`
	CattleTypeOfUse    string `xml:"CattleTypeOfUse,omitempty" json:"cattleTypeOfUse"`
	TypeOfUseDate      string `xml:"TypeOfUseDate,omitempty" json:"typeOfUseDate"`
	FirstCalvingDate   string `xml:"FirstCalvingDate,omitempty" json:"firstCalvingDate"`
	CattleHistoryState string `xml:"CattleHistoryState,omitempty" json:"cattleHistoryState"`
	LastCalvingDate    string `xml:"LastCalvingDate,omitempty" json:"lastCalvingDate"`
}

type CattleDetailResult struct {
	XMLName      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDetailResult"`

	Result       *ProcessingResult `xml:"Result,omitempty"`
	CattleDetail *CattleDetailData `xml:"CattleDetail,omitempty"`
}

type CattleDetailData struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDetailData"`

	BirthNotificationData *CattleBirthData `xml:"BirthNotificationData,omitempty"`
	LongName              string           `xml:"LongName,omitempty"`
	RaceFather            int32            `xml:"RaceFather,omitempty"`
	RaceColorFather       int32            `xml:"RaceColorFather,omitempty"`
	NameFather            string           `xml:"NameFather,omitempty"`
	RaceMother            int32            `xml:"RaceMother,omitempty"`
	RaceColorMother       int32            `xml:"RaceColorMother,omitempty"`
	NameMother            string           `xml:"NameMother,omitempty"`
	Deformations          *IntArray        `xml:"Deformations,omitempty"`
	DeathDate             time.Time        `xml:"DeathDate,omitempty"`
	BvdState              string           `xml:"BvdState,omitempty"`
	BtState               string           `xml:"BtState,omitempty"`
	CattleTypeOfUse       int32            `xml:"CattleTypeOfUse,omitempty"`
	TypeOfUseDate         time.Time        `xml:"TypeOfUseDate,omitempty"`
	FirstCalvingDate      time.Time        `xml:"FirstCalvingDate,omitempty"`
	CurrentHusbandry      int32            `xml:"CurrentHusbandry,omitempty"`
	AllYearHusbandry      int32            `xml:"AllYearHusbandry,omitempty"`
	PendulumHusbandry     int32            `xml:"PendulumHusbandry,omitempty"`
	Beefiness             string           `xml:"Beefiness,omitempty"`
	SlaughterCategory     string           `xml:"SlaughterCategory,omitempty"`
	FatTissue             string           `xml:"FatTissue,omitempty"`
}

type CattleDetailResultV2 struct {
	XMLName      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDetailResultV2"`

	Result       *ProcessingResult   `xml:"Result,omitempty"`
	CattleDetail *CattleDetailDataV2 `xml:"CattleDetail,omitempty"`
}

type CattleDetailDataV2 struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDetailDataV2"`

	*CattleDetailData

	LValue  float64 `xml:"LValue,omitempty"`
}

type CattleStateExternalResult struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStateExternalResult"`

	Result      *ProcessingResult `xml:"Result,omitempty"`
	CattleState *CattleStateData  `xml:"CattleState,omitempty"`
}

type CattleStateData struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStateData"`

	EarTagNumber       string `xml:"EarTagNumber,omitempty"`
	Name               string `xml:"Name,omitempty"`
	BvdState           string `xml:"BvdState,omitempty"`
	CattleHistoryState int32  `xml:"CattleHistoryState,omitempty"`
}

type CattleStateExternalResultV2 struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStateExternalResultV2"`

	Result      *ProcessingResult  `xml:"Result,omitempty"`
	CattleState *CattleStateDataV2 `xml:"CattleState,omitempty"`
}

type CattleStateDataV2 struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStateDataV2"`

	EarTagNumber       string    `xml:"EarTagNumber,omitempty"`
	Name               string    `xml:"Name,omitempty"`
	BvdState           string    `xml:"BvdState,omitempty"`
	CattleHistoryState int32     `xml:"CattleHistoryState,omitempty"`
	CattleBirthDate    time.Time `xml:"CattleBirthDate,omitempty"`
	CattleAgeInDays    int32     `xml:"CattleAgeInDays,omitempty"`
}

type NewEarTagOrderResult struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 NewEarTagOrderResult"`

	Result      *ProcessingResult `xml:"Result,omitempty"`
	EarTagOrder *EarTagOrderData  `xml:"EarTagOrder,omitempty"`
}

type DisposalContributionResult struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DisposalContributionResult"`

	DisposalContributionState bool              `xml:"DisposalContributionState,omitempty"`
	Result                    *ProcessingResult `xml:"Result,omitempty"`
}

type AnimalClassificationData struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 AnimalClassificationData"`

	MessageID                   int64     `xml:"MessageID,omitempty"`
	EarTagNumber                string    `xml:"EarTagNumber,omitempty"`
	TVDNumberOrigin             int32     `xml:"TVDNumberOrigin,omitempty"`
	SlaughterDate               time.Time `xml:"SlaughterDate,omitempty"`
	TVDNumberSlaughterInitiator int32     `xml:"TVDNumberSlaughterInitiator,omitempty"`
	Genus                       int32     `xml:"Genus,omitempty"`
	Weight                      float64   `xml:"Weight,omitempty"`
	ClassifierNumber            int32     `xml:"ClassifierNumber,omitempty"`
	ClassifierEquipmentID       string    `xml:"ClassifierEquipmentID,omitempty"`
	ContractorNumberSlaughter   string    `xml:"ContractorNumberSlaughter,omitempty"`
	SlaughterCategory           string    `xml:"SlaughterCategory,omitempty"`
	Beefiness                   string    `xml:"Beefiness,omitempty"`
	FatTissue                   string    `xml:"FatTissue,omitempty"`
	MFA                         int32     `xml:"MFA,omitempty"`
}

type NotificationResult struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 NotificationResult"`

	Result         *ProcessingResult `xml:"Result,omitempty"`
	NotificationID int32             `xml:"NotificationID,omitempty"`
}

type AnimalClassificationDataV2 struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 AnimalClassificationDataV2"`

	*AnimalClassificationData

	LValue  float64 `xml:"LValue,omitempty"`
}

type DataAccessResult struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DataAccessResult"`

	Result  *ProcessingResult `xml:"Result,omitempty"`
	IDPRG   int32             `xml:"ID_PRG,omitempty"`
}

type CattlePassportResult struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattlePassportResult"`

	Result          *ProcessingResult    `xml:"Result,omitempty"`
	CattlePassports *CattlePassportArray `xml:"CattlePassports,omitempty"`
}

type CattlePassportArray struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattlePassportArray"`

	CattlePassportItem []*CattlePassport `xml:"CattlePassportItem,omitempty"`
}

type CattlePassport struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattlePassport"`

	EarTagNumber      string            `xml:"EarTagNumber,omitempty"`
	PassportNumber    int32             `xml:"PassportNumber,omitempty"`
	PassportIssueDate time.Time         `xml:"PassportIssueDate,omitempty"`
	Result            *ProcessingResult `xml:"Result,omitempty"`
}

type CattleOffspringResult struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleOffspringResult"`

	Result             *ProcessingResult         `xml:"Result,omitempty"`
	EarTagNumberMother string                    `xml:"EarTagNumberMother,omitempty"`
	CattleOffsprings   *CattleOffspringDataArray `xml:"CattleOffsprings,omitempty"`
}

type CattleOffspringDataArray struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleOffspringDataArray"`

	CattleOffspringDataItem []*CattleOffspringData `xml:"CattleOffspringDataItem,omitempty"`
}

type CattleOffspringData struct {
	XMLName      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleOffspringData"`

	EarTagNumber string    `xml:"EarTagNumber,omitempty"`
	BirthDate    time.Time `xml:"BirthDate,omitempty"`
	Gender       int32     `xml:"Gender,omitempty"`
	Name         string    `xml:"Name,omitempty"`
	Race         int32     `xml:"Race,omitempty"`
}

type PoultryBarnInNotification struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PoultryBarnInNotification"`

	Amount                    int32                   `xml:"Amount,omitempty"`
	EventDate                 time.Time               `xml:"EventDate,omitempty"`
	SourceTvdNumber           int32                   `xml:"SourceTvdNumber,omitempty"`
	PoultryUsageReason        *EnumPoultryUsageReason `xml:"PoultryUsageReason,omitempty"`
	AllowSeveralNotifications bool                    `xml:"AllowSeveralNotifications,omitempty"`
}

type SearchPoultryBarnInNotificationRequest struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SearchPoultryBarnInNotificationRequest"`

	*BaseRequest

	TVDNumber int32     `xml:"TVDNumber,omitempty"`
	FromDate  time.Time `xml:"FromDate,omitempty"`
	ToDate    time.Time `xml:"ToDate,omitempty"`
}

type GetPoultryBarnInNotificationResult struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPoultryBarnInNotificationResult"`

	Result               *ProcessingResult                                                      `xml:"Result,omitempty"`
	ArrivalNotifications *SmallAnimalNotificationDataArrayOfTypeSearchPoultryBarnInNotification `xml:"ArrivalNotifications,omitempty"`
}

type SmallAnimalNotificationDataArrayOfTypeSearchPoultryBarnInNotification struct {
	XMLName                              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SmallAnimalNotificationDataArrayOfTypeSearchPoultryBarnInNotification"`

	SmallAnimalNotificationDataArrayItem []*SearchPoultryBarnInNotification `xml:"SmallAnimalNotificationDataArrayItem,omitempty"`
}

type SearchPoultryBarnInNotification struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SearchPoultryBarnInNotification"`

	Amount                     int32                                       `xml:"Amount,omitempty"`
	EventDate                  time.Time                                   `xml:"EventDate,omitempty"`
	SourceTvdNumber            int32                                       `xml:"SourceTvdNumber,omitempty"`
	PoultryUsageReason         *TranslatedEnumTypeOfEnumPoultryUsageReason `xml:"PoultryUsageReason,omitempty"`
	NotificationDate           time.Time                                   `xml:"NotificationDate,omitempty"`
	HerdenIdentificationNumber string                                      `xml:"HerdenIdentificationNumber,omitempty"`
	CreatedBy                  string                                      `xml:"CreatedBy,omitempty"`
	DeletedBy                  string                                      `xml:"DeletedBy,omitempty"`
	IsDeleted                  bool                                        `xml:"IsDeleted,omitempty"`
}

type GetAnimalHusbandryDetailRequest struct {
	XMLName   xml.Name //`xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryDetailRequest"`

	*BaseRequest

	TVDNumber int32 `xml:"TVDNumber,omitempty"`
}

type GetAnimalHusbandryDetailResult struct {
	XMLName                  xml.Name // `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryDetailResult"`

	Result                   *ProcessingResult                                 `xml:"Result,omitempty"`
	PostData                 *HusbandryResult                                  `xml:"PostData,omitempty"`
	IsActive                 bool                                              `xml:"IsActive,omitempty"`
	MunicipalityName         string                                            `xml:"MunicipalityName,omitempty"`
	CantonShortname          string                                            `xml:"CantonShortname,omitempty"`
	CoordinateX              int32                                             `xml:"CoordinateX,omitempty"`
	CoordinateY              int32                                             `xml:"CoordinateY,omitempty"`
	Altitude                 int32                                             `xml:"Altitude,omitempty"`
	CantonAnimalHusbandryKey string                                            `xml:"CantonAnimalHusbandryKey,omitempty"`
	CantonPersonKey          string                                            `xml:"CantonPersonKey,omitempty"`
	BurNumber                string                                            `xml:"BurNumber,omitempty"`
	AnimalHusbandryType      *TranslatedEnumTypeOfEnumAnimalHusbandryType      `xml:"AnimalHusbandryType,omitempty"`
	MunicipalityNumber       int32                                             `xml:"MunicipalityNumber,omitempty"`
	TypeOfUse                *TranslatedEnumTypeOfEnumAnimalHusbandryTypeOfUse `xml:"TypeOfUse,omitempty"`
	IsMountain               bool                                              `xml:"IsMountain,omitempty"`
	Zone                     *TranslatedEnumTypeOfEnumZone                     `xml:"Zone,omitempty"`
	Area                     *TranslatedEnumTypeOfEnumArea                     `xml:"Area,omitempty"`
}

type SearchSmallAnimalMovementRequest struct {
	XMLName  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SearchSmallAnimalMovementRequest"`

	*BaseRequest

	FromDate time.Time `xml:"FromDate,omitempty"`
	ToDate   time.Time `xml:"ToDate,omitempty"`
}

type PigArrivalNotificationResult struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PigArrivalNotificationResult"`

	Result               *ProcessingResult                                             `xml:"Result,omitempty"`
	ArrivalNotifications *SmallAnimalNotificationDataArrayOfTypePigArrivalNotification `xml:"ArrivalNotifications,omitempty"`
}

type SmallAnimalNotificationDataArrayOfTypePigArrivalNotification struct {
	XMLName                              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SmallAnimalNotificationDataArrayOfTypePigArrivalNotification"`

	SmallAnimalNotificationDataArrayItem []*PigArrivalNotification `xml:"SmallAnimalNotificationDataArrayItem,omitempty"`
}

type PigArrivalNotification struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PigArrivalNotification"`

	TVDNumber       int32                                `xml:"TVDNumber,omitempty"`
	SourceTVDNumber int32                                `xml:"SourceTVDNumber,omitempty"`
	Amount          int32                                `xml:"Amount,omitempty"`
	DeletionDate    time.Time                            `xml:"DeletionDate,omitempty"`
	CreationDate    time.Time                            `xml:"CreationDate,omitempty"`
	EventDate       time.Time                            `xml:"EventDate,omitempty"`
	PigCategory     *TranslatedEnumTypeOfEnumPigCategory `xml:"PigCategory,omitempty"`
}

type GetMembershipForOrganisationRequest struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetMembershipForOrganisationRequest"`

	*BaseRequest

	TVDNumber   int32  `xml:"TVDNumber,omitempty"`
	AgateNumber string `xml:"AgateNumber,omitempty"`
}

type GetMembershipForOrganisationResult struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetMembershipForOrganisationResult"`

	Result                 *ProcessingResult `xml:"Result,omitempty"`
	MembershipOrganisation *GenusElements    `xml:"MembershipOrganisation,omitempty"`
	BreedingOrganisation   *GenusElements    `xml:"BreedingOrganisation,omitempty"`
}

type GenusElements struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GenusElements"`

	Genus   []string `xml:"Genus,omitempty"`
}

type WriteReplacementBatchEarTagOrderRequest struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementBatchEarTagOrderRequest"`

	TVDNumber            int32                 `xml:"TVDNumber,omitempty"`
	Genus                *EnumGenus            `xml:"Genus,omitempty"`
	EartagOrderItemArray *EartagOrderItemArray `xml:"EartagOrderItemArray,omitempty"`
}

type EartagOrderItemArray struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EartagOrderItemArray"`

	EartagOrderArrayItem []*EartagOrderItem `xml:"EartagOrderArrayItem,omitempty"`
}

type EartagOrderItem struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EartagOrderItem"`

	EarTagNumber     string `xml:"EarTagNumber,omitempty"`
	IsExpress        bool   `xml:"IsExpress,omitempty"`
	IsLeftSideOrder  bool   `xml:"IsLeftSideOrder,omitempty"`
	IsRightSideOrder bool   `xml:"IsRightSideOrder,omitempty"`
}

type ReplacementEarTagOrdersResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ReplacementEarTagOrdersResult"`

	Result        *ProcessingResult                      `xml:"Result,omitempty"`
	ResultDetails *ArrayOfReplacementEarTagOrderDataItem `xml:"ResultDetails,omitempty"`
}

type ArrayOfReplacementEarTagOrderDataItem struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfReplacementEarTagOrderDataItem"`

	ReplacementEarTagOrderDataItem []*ReplacementEarTagOrderDataItem `xml:"ReplacementEarTagOrderDataItem,omitempty"`
}

type ReplacementEarTagOrderDataItem struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ReplacementEarTagOrderDataItem"`

	NotificationID  int32                                `xml:"NotificationID,omitempty"`
	CombiArticle    int32                                `xml:"CombiArticle,omitempty"`
	OrderStatus     *TranslatedEnumTypeOfEnumOrderStatus `xml:"OrderStatus,omitempty"`
	OrderStatusDate time.Time                            `xml:"OrderStatusDate,omitempty"`
	EarTagNumber    string                               `xml:"EarTagNumber,omitempty"`
	IsExpress       bool                                 `xml:"IsExpress,omitempty"`
	Text            string                               `xml:"Text,omitempty"`
}

type GetCattlesPerLeavingDateRequest struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDateRequest"`

	TVDNumber   int32     `xml:"TVDNumber,omitempty"`
	LeavingDate time.Time `xml:"LeavingDate,omitempty"`
}

type GetCattlesPerLeavingDateResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDateResult"`

	Result        *ProcessingResult                    `xml:"Result,omitempty"`
	ResultDetails *ArrayOfGetCattlesPerLeavingDateData `xml:"ResultDetails,omitempty"`
}

type ArrayOfGetCattlesPerLeavingDateData struct {
	XMLName                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfGetCattlesPerLeavingDateData"`

	GetCattlesPerLeavingDateData []*GetCattlesPerLeavingDateData `xml:"GetCattlesPerLeavingDateData,omitempty"`
}

type GetCattlesPerLeavingDateData struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDateData"`

	EarTagNumber  string                              `xml:"EarTagNumber,omitempty"`
	Name          string                              `xml:"Name,omitempty"`
	BirthDate     time.Time                           `xml:"BirthDate,omitempty"`
	LeavingDate   time.Time                           `xml:"LeavingDate,omitempty"`
	TvdNumberGJTH string                              `xml:"TvdNumberGJTH,omitempty"`
	Gender        *TranslatedEnumTypeOfEnumGender     `xml:"Gender,omitempty"`
	Race          *TranslatedEnumTypeOfEnumCattleRace `xml:"Race,omitempty"`
}

type Char int32

const ()

type Duration *Duration

const ()

type Guid string

const ()

type EnumPigCategory string

const (
	EnumPigCategoryOther EnumPigCategory = "Other"

	EnumPigCategoryWeaner EnumPigCategory = "Weaner"

	EnumPigCategoryFeederPig EnumPigCategory = "FeederPig"

	EnumPigCategoryPigForSlaughter EnumPigCategory = "PigForSlaughter"

	EnumPigCategorySow EnumPigCategory = "Sow"

	EnumPigCategoryBoar EnumPigCategory = "Boar"

	EnumPigCategoryGilt EnumPigCategory = "Gilt"
)

type EnumEquidNotificationState string

const (
	EnumEquidNotificationStateOk EnumEquidNotificationState = "Ok"

	EnumEquidNotificationStateLocationChangePending EnumEquidNotificationState = "LocationChangePending"

	EnumEquidNotificationStateOwnershipAcquisitionPending EnumEquidNotificationState = "OwnershipAcquisitionPending"

	EnumEquidNotificationStateLocationChangeAndAcquisitionPending EnumEquidNotificationState = "LocationChangeAndAcquisitionPending"
)

type EnumEquidTypeOfUse string

const (
	EnumEquidTypeOfUseUndefined EnumEquidTypeOfUse = "Undefined"

	EnumEquidTypeOfUseNutztier EnumEquidTypeOfUse = "Nutztier"

	EnumEquidTypeOfUseHeimtier EnumEquidTypeOfUse = "Heimtier"
)

type EnumGender string

const (
	EnumGenderFemale EnumGender = "Female"

	EnumGenderMale EnumGender = "Male"
)

type EnumEquidBreed string

const (
	EnumEquidBreedUndefined EnumEquidBreed = "Undefined"

	EnumEquidBreedAchalTekkiner EnumEquidBreed = "AchalTekkiner"

	EnumEquidBreedAchalTekkinerPartbred EnumEquidBreed = "AchalTekkinerPartbred"

	EnumEquidBreedAegidienberger EnumEquidBreed = "Aegidienberger"

	EnumEquidBreedAmericanMiniatureHorse EnumEquidBreed = "AmericanMiniatureHorse"

	EnumEquidBreedAmericanSaddlebred EnumEquidBreed = "AmericanSaddlebred"

	EnumEquidBreedAndalusier EnumEquidBreed = "Andalusier"

	EnumEquidBreedAngloarabVollblutVorbuch EnumEquidBreed = "AngloarabVollblutVorbuch"

	EnumEquidBreedAngloaraber EnumEquidBreed = "Angloaraber"

	EnumEquidBreedAngloAraberVorbuch EnumEquidBreed = "AngloAraberVorbuch"

	EnumEquidBreedAngloarabischesHalbblut EnumEquidBreed = "AngloarabischesHalbblut"

	EnumEquidBreedAngloarabischesHalbblutVorbuch EnumEquidBreed = "AngloarabischesHalbblutVorbuch"

	EnumEquidBreedAngloarabischesVollblut EnumEquidBreed = "AngloarabischesVollblut"

	EnumEquidBreedAngloNormaenner EnumEquidBreed = "AngloNormaenner"

	EnumEquidBreedAppaloosa EnumEquidBreed = "Appaloosa"

	EnumEquidBreedAraberBerber EnumEquidBreed = "AraberBerber"

	EnumEquidBreedArabischesVollblut EnumEquidBreed = "ArabischesVollblut"

	EnumEquidBreedAraboFriesen EnumEquidBreed = "AraboFriesen"

	EnumEquidBreedBadenWuerttemberg EnumEquidBreed = "BadenWuerttemberg"

	EnumEquidBreedBayerischesWarmblut EnumEquidBreed = "BayerischesWarmblut"

	EnumEquidBreedBelgischesWarmblut EnumEquidBreed = "BelgischesWarmblut"

	EnumEquidBreedBerber EnumEquidBreed = "Berber"

	EnumEquidBreedBERMA EnumEquidBreed = "BERMA"

	EnumEquidBreedBosniaken EnumEquidBreed = "Bosniaken"

	EnumEquidBreedBrandenburger EnumEquidBreed = "Brandenburger"

	EnumEquidBreedBritishSporthorse EnumEquidBreed = "BritishSporthorse"

	EnumEquidBreedBritishSpottedPony EnumEquidBreed = "BritishSpottedPony"

	EnumEquidBreedCamargue EnumEquidBreed = "Camargue"

	EnumEquidBreedCHKleinpferd EnumEquidBreed = "CHKleinpferd"

	EnumEquidBreedCHSportpony EnumEquidBreed = "CHSportpony"

	EnumEquidBreedClydesdale EnumEquidBreed = "Clydesdale"

	EnumEquidBreedCobNormand EnumEquidBreed = "CobNormand"

	EnumEquidBreedConnemara EnumEquidBreed = "Connemara"

	EnumEquidBreedCreamColorSchweiz EnumEquidBreed = "CreamColorSchweiz"

	EnumEquidBreedCremeColors EnumEquidBreed = "CremeColors"

	EnumEquidBreedCriaCaballar EnumEquidBreed = "CriaCaballar"

	EnumEquidBreedCriollo EnumEquidBreed = "Criollo"

	EnumEquidBreedDartmoor EnumEquidBreed = "Dartmoor"

	EnumEquidBreedDeutschesClassicPony EnumEquidBreed = "DeutschesClassicPony"

	EnumEquidBreedDeutschesReitpony EnumEquidBreed = "DeutschesReitpony"

	EnumEquidBreedDeutschesWarmblut EnumEquidBreed = "DeutschesWarmblut"

	EnumEquidBreedDuelmener EnumEquidBreed = "Duelmener"

	EnumEquidBreedEnglischesVollblut EnumEquidBreed = "EnglischesVollblut"

	EnumEquidBreedAndereEselrasse EnumEquidBreed = "AndereEselrasse"

	EnumEquidBreedExmoor EnumEquidBreed = "Exmoor"

	EnumEquidBreedFellPony EnumEquidBreed = "FellPony"

	EnumEquidBreedFinnischesWarmblut EnumEquidBreed = "FinnischesWarmblut"

	EnumEquidBreedFjord EnumEquidBreed = "Fjord"

	EnumEquidBreedFriese EnumEquidBreed = "Friese"

	EnumEquidBreedHaflingerMischrasse EnumEquidBreed = "HaflingerMischrasse"

	EnumEquidBreedHannoveraner EnumEquidBreed = "Hannoveraner"

	EnumEquidBreedHessen EnumEquidBreed = "Hessen"

	EnumEquidBreedHighlandPony EnumEquidBreed = "HighlandPony"

	EnumEquidBreedHighlandPonyCarron EnumEquidBreed = "HighlandPonyCarron"

	EnumEquidBreedHispano EnumEquidBreed = "Hispano"

	EnumEquidBreedHolsteiner EnumEquidBreed = "Holsteiner"

	EnumEquidBreedIrlaender EnumEquidBreed = "Irlaender"

	EnumEquidBreedIslandpferd EnumEquidBreed = "Islandpferd"

	EnumEquidBreedKladruber EnumEquidBreed = "Kladruber"

	EnumEquidBreedKnabstrupper EnumEquidBreed = "Knabstrupper"

	EnumEquidBreedKreuzungZVCH EnumEquidBreed = "KreuzungZVCH"

	EnumEquidBreedLettischesWarmblut EnumEquidBreed = "LettischesWarmblut"

	EnumEquidBreedLewitzer EnumEquidBreed = "Lewitzer"

	EnumEquidBreedLewitzschecke EnumEquidBreed = "Lewitzschecke"

	EnumEquidBreedLipizzaner EnumEquidBreed = "Lipizzaner"

	EnumEquidBreedLittauischesWarmblut EnumEquidBreed = "LittauischesWarmblut"

	EnumEquidBreedLusitano EnumEquidBreed = "Lusitano"

	EnumEquidBreedLuxemburgischesWarmblut EnumEquidBreed = "LuxemburgischesWarmblut"

	EnumEquidBreedMaulesel EnumEquidBreed = "Maulesel"

	EnumEquidBreedMaultier EnumEquidBreed = "Maultier"

	EnumEquidBreedMazedonier EnumEquidBreed = "Mazedonier"

	EnumEquidBreedMazedonierPartbred EnumEquidBreed = "MazedonierPartbred"

	EnumEquidBreedMecklenburger EnumEquidBreed = "Mecklenburger"

	EnumEquidBreedMerens EnumEquidBreed = "Merens"

	EnumEquidBreedMissouriFoxtrotter EnumEquidBreed = "MissouriFoxtrotter"

	EnumEquidBreedMorgan EnumEquidBreed = "Morgan"

	EnumEquidBreedNewForest EnumEquidBreed = "NewForest"

	EnumEquidBreedNoriker EnumEquidBreed = "Noriker"

	EnumEquidBreedOesterrWarmblut EnumEquidBreed = "OesterrWarmblut"

	EnumEquidBreedOldenburgerPferd EnumEquidBreed = "OldenburgerPferd"

	EnumEquidBreedOrlow EnumEquidBreed = "Orlow"

	EnumEquidBreedOstfriese EnumEquidBreed = "Ostfriese"

	EnumEquidBreedPaint EnumEquidBreed = "Paint"

	EnumEquidBreedPaintHorse EnumEquidBreed = "PaintHorse"

	EnumEquidBreedPalomino EnumEquidBreed = "Palomino"

	EnumEquidBreedPartbred EnumEquidBreed = "Partbred"

	EnumEquidBreedPartbredaraber EnumEquidBreed = "Partbredaraber"

	EnumEquidBreedPaso EnumEquidBreed = "Paso"

	EnumEquidBreedPasoColombiano EnumEquidBreed = "PasoColombiano"

	EnumEquidBreedPasoFino EnumEquidBreed = "PasoFino"

	EnumEquidBreedPasoPeruano EnumEquidBreed = "PasoPeruano"

	EnumEquidBreedPercheron EnumEquidBreed = "Percheron"

	EnumEquidBreedPinto EnumEquidBreed = "Pinto"

	EnumEquidBreedPleven EnumEquidBreed = "Pleven"

	EnumEquidBreedPolenTrakehner EnumEquidBreed = "PolenTrakehner"

	EnumEquidBreedPottok EnumEquidBreed = "Pottok"

	EnumEquidBreedPrzewalski EnumEquidBreed = "Przewalski"

	EnumEquidBreedPurarazaespanola EnumEquidBreed = "Purarazaespanola"

	EnumEquidBreedQuarterHorse EnumEquidBreed = "QuarterHorse"

	EnumEquidBreedRheinland EnumEquidBreed = "Rheinland"

	EnumEquidBreedRussischesWarmblut EnumEquidBreed = "RussischesWarmblut"

	EnumEquidBreedSachsen EnumEquidBreed = "Sachsen"

	EnumEquidBreedSachsenAnhaltiner EnumEquidBreed = "SachsenAnhaltiner"

	EnumEquidBreedSelleFrancais EnumEquidBreed = "SelleFrancais"

	EnumEquidBreedShagyaAraber EnumEquidBreed = "ShagyaAraber"

	EnumEquidBreedShetlandpony EnumEquidBreed = "Shetlandpony"

	EnumEquidBreedShireHorse EnumEquidBreed = "ShireHorse"

	EnumEquidBreedSlowenischesWarmblut EnumEquidBreed = "SlowenischesWarmblut"

	EnumEquidBreedSpecialColorSchweiz EnumEquidBreed = "SpecialColorSchweiz"

	EnumEquidBreedSporthorseBrasilien EnumEquidBreed = "SporthorseBrasilien"

	EnumEquidBreedSporthorseMexico EnumEquidBreed = "SporthorseMexico"

	EnumEquidBreedTennesseeWalkingHorse EnumEquidBreed = "TennesseeWalkingHorse"

	EnumEquidBreedThueringer EnumEquidBreed = "Thueringer"

	EnumEquidBreedTinker EnumEquidBreed = "Tinker"

	EnumEquidBreedTraber EnumEquidBreed = "Traber"

	EnumEquidBreedTraitComtois EnumEquidBreed = "TraitComtois"

	EnumEquidBreedTrakehner EnumEquidBreed = "Trakehner"

	EnumEquidBreedWarmblutRheinlandPfalzSaar EnumEquidBreed = "WarmblutRheinlandPfalzSaar"

	EnumEquidBreedWelsh EnumEquidBreed = "Welsh"

	EnumEquidBreedWelshCobSektionDWD EnumEquidBreed = "WelshCobSektionDWD"

	EnumEquidBreedWelshMountainPonyWA EnumEquidBreed = "WelshMountainPonyWA"

	EnumEquidBreedWelshPartbredWK EnumEquidBreed = "WelshPartbredWK"

	EnumEquidBreedWelshPonyCobTypWC EnumEquidBreed = "WelshPonyCobTypWC"

	EnumEquidBreedWelshRidingPonyWB EnumEquidBreed = "WelshRidingPonyWB"

	EnumEquidBreedWestfale EnumEquidBreed = "Westfale"

	EnumEquidBreedWuerttemberger EnumEquidBreed = "Wuerttemberger"

	EnumEquidBreedZangersheide EnumEquidBreed = "Zangersheide"

	EnumEquidBreedZweibruecken EnumEquidBreed = "Zweibruecken"

	EnumEquidBreedFreiberger EnumEquidBreed = "Freiberger"

	EnumEquidBreedHaflinger EnumEquidBreed = "Haflinger"

	EnumEquidBreedSchweizerWarmblut EnumEquidBreed = "SchweizerWarmblut"

	EnumEquidBreedAraber EnumEquidBreed = "Araber"

	EnumEquidBreedPony EnumEquidBreed = "Pony"

	EnumEquidBreedWarmblut EnumEquidBreed = "Warmblut"

	EnumEquidBreedVollblut EnumEquidBreed = "Vollblut"

	EnumEquidBreedKaltblut EnumEquidBreed = "Kaltblut"

	EnumEquidBreedKreuzung EnumEquidBreed = "Kreuzung"

	EnumEquidBreedAndere EnumEquidBreed = "Andere"

	EnumEquidBreedHollaendischesWarmblutKWPN EnumEquidBreed = "HollaendischesWarmblutKWPN"

	EnumEquidBreedDaenischesWarmblut EnumEquidBreed = "DaenischesWarmblut"

	EnumEquidBreedPolnischesWarmblut EnumEquidBreed = "PolnischesWarmblut"

	EnumEquidBreedUngarischesWarmblut EnumEquidBreed = "UngarischesWarmblut"

	EnumEquidBreedTschechischesWarmblut EnumEquidBreed = "TschechischesWarmblut"

	EnumEquidBreedItalienischesWarmblut EnumEquidBreed = "ItalienischesWarmblut"

	EnumEquidBreedCHSportpferd EnumEquidBreed = "CHSportpferd"

	EnumEquidBreedHalbblut EnumEquidBreed = "Halbblut"

	EnumEquidBreedSchwedischesHalbblut EnumEquidBreed = "SchwedischesHalbblut"

	EnumEquidBreedDaenischesReitpony EnumEquidBreed = "DaenischesReitpony"

	EnumEquidBreedBardigiano EnumEquidBreed = "Bardigiano"

	EnumEquidBreedKabardiner EnumEquidBreed = "Kabardiner"

	EnumEquidBreedArgentinischesPolopony EnumEquidBreed = "ArgentinischesPolopony"

	EnumEquidBreedUngarischesVollblut EnumEquidBreed = "UngarischesVollblut"

	EnumEquidBreedHalfSaddlebred EnumEquidBreed = "HalfSaddlebred"

	EnumEquidBreedTigerscheck EnumEquidBreed = "Tigerscheck"

	EnumEquidBreedTigerscheckShetlandtyp EnumEquidBreed = "TigerscheckShetlandtyp"

	EnumEquidBreedMiniShetlandpony EnumEquidBreed = "MiniShetlandpony"

	EnumEquidBreedIrishCob EnumEquidBreed = "IrishCob"

	EnumEquidBreedCurlyHorse EnumEquidBreed = "CurlyHorse"

	EnumEquidBreedCruzadoIberico EnumEquidBreed = "CruzadoIberico"

	EnumEquidBreedDalesPony EnumEquidBreed = "DalesPony"

	EnumEquidBreedDeutschesPartbredShetlandPony EnumEquidBreed = "DeutschesPartbredShetlandPony"

	EnumEquidBreedDeutschesSportpferd EnumEquidBreed = "DeutschesSportpferd"

	EnumEquidBreedEnglischesReitpony EnumEquidBreed = "EnglischesReitpony"

	EnumEquidBreedHannoverschersHalbblut EnumEquidBreed = "HannoverschersHalbblut"

	EnumEquidBreedInternationalesOldenburgerSpringpferd EnumEquidBreed = "InternationalesOldenburgerSpringpferd"

	EnumEquidBreedKigerMustang EnumEquidBreed = "KigerMustang"

	EnumEquidBreedKleinesDeutschesPony EnumEquidBreed = "KleinesDeutschesPony"

	EnumEquidBreedKleinesDeutschesReitpferd EnumEquidBreed = "KleinesDeutschesReitpferd"

	EnumEquidBreedLeonharder EnumEquidBreed = "Leonharder"

	EnumEquidBreedLeutstettenerPferd EnumEquidBreed = "LeutstettenerPferd"

	EnumEquidBreedMangalargaMarchadores EnumEquidBreed = "MangalargaMarchadores"

	EnumEquidBreedNRWReitpferd EnumEquidBreed = "NRWReitpferd"

	EnumEquidBreedOstfrieseAltOldenburger EnumEquidBreed = "OstfrieseAltOldenburger"

	EnumEquidBreedPasoIberoamericano EnumEquidBreed = "PasoIberoamericano"

	EnumEquidBreedPortugiesischesSportpferd EnumEquidBreed = "PortugiesischesSportpferd"

	EnumEquidBreedRheinischDeutschesKaltblut EnumEquidBreed = "RheinischDeutschesKaltblut"

	EnumEquidBreedRottaler EnumEquidBreed = "Rottaler"

	EnumEquidBreedSchwarzwaelderKaltblut EnumEquidBreed = "SchwarzwaelderKaltblut"

	EnumEquidBreedSpanischesSportpferd EnumEquidBreed = "SpanischesSportpferd"

	EnumEquidBreedSueddeutschesKaltblut EnumEquidBreed = "SueddeutschesKaltblut"

	EnumEquidBreedWarlander EnumEquidBreed = "Warlander"

	EnumEquidBreedPoitouEsel EnumEquidBreed = "PoitouEsel"

	EnumEquidBreedGrandNoirduBerry EnumEquidBreed = "GrandNoirduBerry"

	EnumEquidBreedAndalusischerRiesenesel EnumEquidBreed = "AndalusischerRiesenesel"

	EnumEquidBreedKatalanischerRiesenesel EnumEquidBreed = "KatalanischerRiesenesel"

	EnumEquidBreedAmericanMiniaturEsel EnumEquidBreed = "AmericanMiniaturEsel"

	EnumEquidBreedRagusana EnumEquidBreed = "Ragusana"

	EnumEquidBreedAmiataEsel EnumEquidBreed = "AmiataEsel"

	EnumEquidBreedMartinaFrancaEsel EnumEquidBreed = "MartinaFrancaEsel"

	EnumEquidBreedContentinEsel EnumEquidBreed = "ContentinEsel"

	EnumEquidBreedNormandieEsel EnumEquidBreed = "NormandieEsel"

	EnumEquidBreedHausesel EnumEquidBreed = "Hausesel"

	EnumEquidBreedSchweizerZuchtpferd EnumEquidBreed = "SchweizerZuchtpferd"
)

type EnumEquidWithersClass string

const (
	EnumEquidWithersClassLessOrEqualThan148cm EnumEquidWithersClass = "LessOrEqualThan148cm"

	EnumEquidWithersClassGreaterThan148cm EnumEquidWithersClass = "GreaterThan148cm"
)

type EnumCountry string

const (
	EnumCountryUndefined EnumCountry = "Undefined"

	EnumCountryABW EnumCountry = "ABW"

	EnumCountryAFG EnumCountry = "AFG"

	EnumCountryAGO EnumCountry = "AGO"

	EnumCountryAIA EnumCountry = "AIA"

	EnumCountryALA EnumCountry = "ALA"

	EnumCountryALB EnumCountry = "ALB"

	EnumCountryAND EnumCountry = "AND"

	EnumCountryANT EnumCountry = "ANT"

	EnumCountryARE EnumCountry = "ARE"

	EnumCountryARG EnumCountry = "ARG"

	EnumCountryARM EnumCountry = "ARM"

	EnumCountryASM EnumCountry = "ASM"

	EnumCountryATA EnumCountry = "ATA"

	EnumCountryATG EnumCountry = "ATG"

	EnumCountryAUS EnumCountry = "AUS"

	EnumCountryAUT EnumCountry = "AUT"

	EnumCountryAZE EnumCountry = "AZE"

	EnumCountryBDI EnumCountry = "BDI"

	EnumCountryBEL EnumCountry = "BEL"

	EnumCountryBEN EnumCountry = "BEN"

	EnumCountryBFA EnumCountry = "BFA"

	EnumCountryBGD EnumCountry = "BGD"

	EnumCountryBGR EnumCountry = "BGR"

	EnumCountryBHR EnumCountry = "BHR"

	EnumCountryBHS EnumCountry = "BHS"

	EnumCountryBIH EnumCountry = "BIH"

	EnumCountryBLM EnumCountry = "BLM"

	EnumCountryBLR EnumCountry = "BLR"

	EnumCountryBLZ EnumCountry = "BLZ"

	EnumCountryBMU EnumCountry = "BMU"

	EnumCountryBOL EnumCountry = "BOL"

	EnumCountryBRA EnumCountry = "BRA"

	EnumCountryBRB EnumCountry = "BRB"

	EnumCountryBRN EnumCountry = "BRN"

	EnumCountryBTN EnumCountry = "BTN"

	EnumCountryBWA EnumCountry = "BWA"

	EnumCountryCAF EnumCountry = "CAF"

	EnumCountryCAN EnumCountry = "CAN"

	EnumCountryCHE EnumCountry = "CHE"

	EnumCountryCHL EnumCountry = "CHL"

	EnumCountryCHN EnumCountry = "CHN"

	EnumCountryCIV EnumCountry = "CIV"

	EnumCountryCMR EnumCountry = "CMR"

	EnumCountryCOD EnumCountry = "COD"

	EnumCountryCOG EnumCountry = "COG"

	EnumCountryCOK EnumCountry = "COK"

	EnumCountryCOL EnumCountry = "COL"

	EnumCountryCOM EnumCountry = "COM"

	EnumCountryCPV EnumCountry = "CPV"

	EnumCountryCRI EnumCountry = "CRI"

	EnumCountryCUB EnumCountry = "CUB"

	EnumCountryCXR EnumCountry = "CXR"

	EnumCountryCYM EnumCountry = "CYM"

	EnumCountryCYP EnumCountry = "CYP"

	EnumCountryCZE EnumCountry = "CZE"

	EnumCountryDEU EnumCountry = "DEU"

	EnumCountryDJI EnumCountry = "DJI"

	EnumCountryDMA EnumCountry = "DMA"

	EnumCountryDNK EnumCountry = "DNK"

	EnumCountryDOM EnumCountry = "DOM"

	EnumCountryDZA EnumCountry = "DZA"

	EnumCountryECU EnumCountry = "ECU"

	EnumCountryEGY EnumCountry = "EGY"

	EnumCountryERI EnumCountry = "ERI"

	EnumCountryESH EnumCountry = "ESH"

	EnumCountryESP EnumCountry = "ESP"

	EnumCountryEST EnumCountry = "EST"

	EnumCountryETH EnumCountry = "ETH"

	EnumCountryFIN EnumCountry = "FIN"

	EnumCountryFJI EnumCountry = "FJI"

	EnumCountryFLK EnumCountry = "FLK"

	EnumCountryFRA EnumCountry = "FRA"

	EnumCountryFRO EnumCountry = "FRO"

	EnumCountryFSM EnumCountry = "FSM"

	EnumCountryGAB EnumCountry = "GAB"

	EnumCountryGBR EnumCountry = "GBR"

	EnumCountryGEO EnumCountry = "GEO"

	EnumCountryGGY EnumCountry = "GGY"

	EnumCountryGHA EnumCountry = "GHA"

	EnumCountryGIB EnumCountry = "GIB"

	EnumCountryGIN EnumCountry = "GIN"

	EnumCountryGLP EnumCountry = "GLP"

	EnumCountryGMB EnumCountry = "GMB"

	EnumCountryGNB EnumCountry = "GNB"

	EnumCountryGNQ EnumCountry = "GNQ"

	EnumCountryGRC EnumCountry = "GRC"

	EnumCountryGRD EnumCountry = "GRD"

	EnumCountryGRL EnumCountry = "GRL"

	EnumCountryGTM EnumCountry = "GTM"

	EnumCountryGUF EnumCountry = "GUF"

	EnumCountryGUM EnumCountry = "GUM"

	EnumCountryGUY EnumCountry = "GUY"

	EnumCountryHKG EnumCountry = "HKG"

	EnumCountryHND EnumCountry = "HND"

	EnumCountryHRV EnumCountry = "HRV"

	EnumCountryHTI EnumCountry = "HTI"

	EnumCountryHUN EnumCountry = "HUN"

	EnumCountryIDN EnumCountry = "IDN"

	EnumCountryIMN EnumCountry = "IMN"

	EnumCountryIND EnumCountry = "IND"

	EnumCountryIRL EnumCountry = "IRL"

	EnumCountryIRN EnumCountry = "IRN"

	EnumCountryIRQ EnumCountry = "IRQ"

	EnumCountryISL EnumCountry = "ISL"

	EnumCountryISR EnumCountry = "ISR"

	EnumCountryITA EnumCountry = "ITA"

	EnumCountryJAM EnumCountry = "JAM"

	EnumCountryJEY EnumCountry = "JEY"

	EnumCountryJOR EnumCountry = "JOR"

	EnumCountryJPN EnumCountry = "JPN"

	EnumCountryKAZ EnumCountry = "KAZ"

	EnumCountryKEN EnumCountry = "KEN"

	EnumCountryKGZ EnumCountry = "KGZ"

	EnumCountryKHM EnumCountry = "KHM"

	EnumCountryKIR EnumCountry = "KIR"

	EnumCountryKNA EnumCountry = "KNA"

	EnumCountryKOR EnumCountry = "KOR"

	EnumCountryKWT EnumCountry = "KWT"

	EnumCountryLAO EnumCountry = "LAO"

	EnumCountryLBN EnumCountry = "LBN"

	EnumCountryLBR EnumCountry = "LBR"

	EnumCountryLBY EnumCountry = "LBY"

	EnumCountryLCA EnumCountry = "LCA"

	EnumCountryLIE EnumCountry = "LIE"

	EnumCountryLKA EnumCountry = "LKA"

	EnumCountryLSO EnumCountry = "LSO"

	EnumCountryLTU EnumCountry = "LTU"

	EnumCountryLUX EnumCountry = "LUX"

	EnumCountryLVA EnumCountry = "LVA"

	EnumCountryMAC EnumCountry = "MAC"

	EnumCountryMAF EnumCountry = "MAF"

	EnumCountryMAR EnumCountry = "MAR"

	EnumCountryMCO EnumCountry = "MCO"

	EnumCountryMDA EnumCountry = "MDA"

	EnumCountryMDG EnumCountry = "MDG"

	EnumCountryMDV EnumCountry = "MDV"

	EnumCountryMEX EnumCountry = "MEX"

	EnumCountryMHL EnumCountry = "MHL"

	EnumCountryMKD EnumCountry = "MKD"

	EnumCountryMLI EnumCountry = "MLI"

	EnumCountryMLT EnumCountry = "MLT"

	EnumCountryMMR EnumCountry = "MMR"

	EnumCountryMNE EnumCountry = "MNE"

	EnumCountryMNG EnumCountry = "MNG"

	EnumCountryMNP EnumCountry = "MNP"

	EnumCountryMOZ EnumCountry = "MOZ"

	EnumCountryMRT EnumCountry = "MRT"

	EnumCountryMSR EnumCountry = "MSR"

	EnumCountryMTQ EnumCountry = "MTQ"

	EnumCountryMUS EnumCountry = "MUS"

	EnumCountryMWI EnumCountry = "MWI"

	EnumCountryMYS EnumCountry = "MYS"

	EnumCountryMYT EnumCountry = "MYT"

	EnumCountryNAM EnumCountry = "NAM"

	EnumCountryNCL EnumCountry = "NCL"

	EnumCountryNER EnumCountry = "NER"

	EnumCountryNFK EnumCountry = "NFK"

	EnumCountryNGA EnumCountry = "NGA"

	EnumCountryNIC EnumCountry = "NIC"

	EnumCountryNIU EnumCountry = "NIU"

	EnumCountryNLD EnumCountry = "NLD"

	EnumCountryNOR EnumCountry = "NOR"

	EnumCountryNPL EnumCountry = "NPL"

	EnumCountryNRU EnumCountry = "NRU"

	EnumCountryNZL EnumCountry = "NZL"

	EnumCountryOMN EnumCountry = "OMN"

	EnumCountryPAK EnumCountry = "PAK"

	EnumCountryPAN EnumCountry = "PAN"

	EnumCountryPCN EnumCountry = "PCN"

	EnumCountryPER EnumCountry = "PER"

	EnumCountryPHL EnumCountry = "PHL"

	EnumCountryPLW EnumCountry = "PLW"

	EnumCountryPNG EnumCountry = "PNG"

	EnumCountryPOL EnumCountry = "POL"

	EnumCountryPRI EnumCountry = "PRI"

	EnumCountryPRK EnumCountry = "PRK"

	EnumCountryPRT EnumCountry = "PRT"

	EnumCountryPRY EnumCountry = "PRY"

	EnumCountryPSE EnumCountry = "PSE"

	EnumCountryPYF EnumCountry = "PYF"

	EnumCountryQAT EnumCountry = "QAT"

	EnumCountryREU EnumCountry = "REU"

	EnumCountryROU EnumCountry = "ROU"

	EnumCountryRUS EnumCountry = "RUS"

	EnumCountryRWA EnumCountry = "RWA"

	EnumCountrySAU EnumCountry = "SAU"

	EnumCountrySDN EnumCountry = "SDN"

	EnumCountrySEN EnumCountry = "SEN"

	EnumCountrySGP EnumCountry = "SGP"

	EnumCountrySHN EnumCountry = "SHN"

	EnumCountrySJM EnumCountry = "SJM"

	EnumCountrySLB EnumCountry = "SLB"

	EnumCountrySLE EnumCountry = "SLE"

	EnumCountrySLV EnumCountry = "SLV"

	EnumCountrySMR EnumCountry = "SMR"

	EnumCountrySOM EnumCountry = "SOM"

	EnumCountrySPM EnumCountry = "SPM"

	EnumCountrySRB EnumCountry = "SRB"

	EnumCountrySTP EnumCountry = "STP"

	EnumCountrySUR EnumCountry = "SUR"

	EnumCountrySVK EnumCountry = "SVK"

	EnumCountrySVN EnumCountry = "SVN"

	EnumCountrySWE EnumCountry = "SWE"

	EnumCountrySWZ EnumCountry = "SWZ"

	EnumCountrySYC EnumCountry = "SYC"

	EnumCountrySYR EnumCountry = "SYR"

	EnumCountryTCA EnumCountry = "TCA"

	EnumCountryTCD EnumCountry = "TCD"

	EnumCountryTGO EnumCountry = "TGO"

	EnumCountryTHA EnumCountry = "THA"

	EnumCountryTJK EnumCountry = "TJK"

	EnumCountryTKL EnumCountry = "TKL"

	EnumCountryTKM EnumCountry = "TKM"

	EnumCountryTLS EnumCountry = "TLS"

	EnumCountryTON EnumCountry = "TON"

	EnumCountryTTO EnumCountry = "TTO"

	EnumCountryTUN EnumCountry = "TUN"

	EnumCountryTUR EnumCountry = "TUR"

	EnumCountryTUV EnumCountry = "TUV"

	EnumCountryTZA EnumCountry = "TZA"

	EnumCountryUGA EnumCountry = "UGA"

	EnumCountryUKR EnumCountry = "UKR"

	EnumCountryURY EnumCountry = "URY"

	EnumCountryUSA EnumCountry = "USA"

	EnumCountryUZB EnumCountry = "UZB"

	EnumCountryVAT EnumCountry = "VAT"

	EnumCountryVCT EnumCountry = "VCT"

	EnumCountryVEN EnumCountry = "VEN"

	EnumCountryVGB EnumCountry = "VGB"

	EnumCountryVIR EnumCountry = "VIR"

	EnumCountryVNM EnumCountry = "VNM"

	EnumCountryVUT EnumCountry = "VUT"

	EnumCountryWLF EnumCountry = "WLF"

	EnumCountryWSM EnumCountry = "WSM"

	EnumCountryYEM EnumCountry = "YEM"

	EnumCountryZAF EnumCountry = "ZAF"

	EnumCountryZMB EnumCountry = "ZMB"

	EnumCountryZWE EnumCountry = "ZWE"
)

type EnumPoultryType string

const (
	EnumPoultryTypeOther EnumPoultryType = "Other"

	EnumPoultryTypeChicken EnumPoultryType = "Chicken"

	EnumPoultryTypeHen EnumPoultryType = "Hen"
)

type EnumPoultryUsageReason string

const (
	EnumPoultryUsageReasonUndefined EnumPoultryUsageReason = "Undefined"

	EnumPoultryUsageReasonLayingHen EnumPoultryUsageReason = "LayingHen"

	EnumPoultryUsageReasonMastPoultry EnumPoultryUsageReason = "MastPoultry"

	EnumPoultryUsageReasonMastTurkey EnumPoultryUsageReason = "MastTurkey"

	EnumPoultryUsageReasonBreedingAnimalLayingLine EnumPoultryUsageReason = "BreedingAnimalLayingLine"

	EnumPoultryUsageReasonBreedingAnimalMastLine EnumPoultryUsageReason = "BreedingAnimalMastLine"
)

type EnumAnimalHusbandryType string

const (
	EnumAnimalHusbandryTypeUnknown EnumAnimalHusbandryType = "Unknown"

	EnumAnimalHusbandryTypeProductionHusbandry EnumAnimalHusbandryType = "ProductionHusbandry"

	EnumAnimalHusbandryTypeHerdsmanHusbandry EnumAnimalHusbandryType = "HerdsmanHusbandry"

	EnumAnimalHusbandryTypeCoOperationPastureHusbandry EnumAnimalHusbandryType = "CoOperationPastureHusbandry"

	EnumAnimalHusbandryTypeSummeringHusbandry EnumAnimalHusbandryType = "SummeringHusbandry"

	EnumAnimalHusbandryTypeCoOperationHusbandry EnumAnimalHusbandryType = "CoOperationHusbandry"

	EnumAnimalHusbandryTypeLiveStockGroup EnumAnimalHusbandryType = "LiveStockGroup"

	EnumAnimalHusbandryTypePopulation EnumAnimalHusbandryType = "Population"

	EnumAnimalHusbandryTypeLiveStockDealerEnterprise EnumAnimalHusbandryType = "LiveStockDealerEnterprise"

	EnumAnimalHusbandryTypeMigratoryHerd EnumAnimalHusbandryType = "MigratoryHerd"

	EnumAnimalHusbandryTypeMedicalCenter EnumAnimalHusbandryType = "MedicalCenter"

	EnumAnimalHusbandryTypeSlaughterEnterprise EnumAnimalHusbandryType = "SlaughterEnterprise"

	EnumAnimalHusbandryTypeMarketAuctionExhibition EnumAnimalHusbandryType = "MarketAuctionExhibition"

	EnumAnimalHusbandryTypeCoOperationBranchHusbandry EnumAnimalHusbandryType = "CoOperationBranchHusbandry"

	EnumAnimalHusbandryTypeNonComercial EnumAnimalHusbandryType = "NonComercial"

	EnumAnimalHusbandryTypeOeLNCommunity EnumAnimalHusbandryType = "OeLNCommunity"

	EnumAnimalHusbandryTypeAllYearHusbandry EnumAnimalHusbandryType = "AllYearHusbandry"

	EnumAnimalHusbandryTypeEnterpriseHusbandry EnumAnimalHusbandryType = "EnterpriseHusbandry"
)

type EnumAnimalHusbandryTypeOfUse string

const (
	EnumAnimalHusbandryTypeOfUseMixed EnumAnimalHusbandryTypeOfUse = "Mixed"

	EnumAnimalHusbandryTypeOfUseCow EnumAnimalHusbandryTypeOfUse = "Cow"

	EnumAnimalHusbandryTypeOfUseMilkCow EnumAnimalHusbandryTypeOfUse = "MilkCow"
)

type EnumZone string

const (
	EnumZoneUnknown EnumZone = "Unknown"

	EnumZoneHills EnumZone = "Hills"

	EnumZoneMountain01 EnumZone = "Mountain01"

	EnumZoneMountain02 EnumZone = "Mountain02"

	EnumZoneMountain03 EnumZone = "Mountain03"

	EnumZoneMountain04 EnumZone = "Mountain04"

	EnumZoneSummering EnumZone = "Summering"

	EnumZoneValley EnumZone = "Valley"

	EnumZoneForeignAncestralZoneForSummeringAnh EnumZone = "ForeignAncestralZoneForSummeringAnh"

	EnumZoneForeignNonAncestralZoneForSummeringAnh EnumZone = "ForeignNonAncestralZoneForSummeringAnh"
)

type EnumArea string

const (
	EnumAreaUnknown EnumArea = "Unknown"

	EnumAreaMountain EnumArea = "Mountain"

	EnumAreaSummering EnumArea = "Summering"

	EnumAreaValley EnumArea = "Valley"
)

type EnumGenus string

const (
	EnumGenusUnknow EnumGenus = "Unknow"

	EnumGenusCattle EnumGenus = "Cattle"

	EnumGenusPig EnumGenus = "Pig"

	EnumGenusSheep EnumGenus = "Sheep"

	EnumGenusGoat EnumGenus = "Goat"

	EnumGenusCamelid EnumGenus = "Camelid"

	EnumGenusGame EnumGenus = "Game"

	EnumGenusEquid EnumGenus = "Equid"

	EnumGenusBee EnumGenus = "Bee"

	EnumGenusFish EnumGenus = "Fish"

	EnumGenusPoultry EnumGenus = "Poultry"

	EnumGenusOthers EnumGenus = "Others"
)

type EnumOrderStatus string

const (
	EnumOrderStatusNew EnumOrderStatus = "New"

	EnumOrderStatusInDevelopment EnumOrderStatus = "InDevelopment"

	EnumOrderStatusOrderedAtSupplier EnumOrderStatus = "OrderedAtSupplier"

	EnumOrderStatusOrderedFromStock EnumOrderStatus = "OrderedFromStock"

	EnumOrderStatusSend EnumOrderStatus = "Send"

	EnumOrderStatusClosed EnumOrderStatus = "Closed"

	EnumOrderStatusDeleted EnumOrderStatus = "Deleted"
)

type EnumCattleRace string

const (
	EnumCattleRaceUndefined EnumCattleRace = "Undefined"

	EnumCattleRaceHolstein EnumCattleRace = "Holstein"

	EnumCattleRaceRotfleckvieh EnumCattleRace = "Rotfleckvieh"

	EnumCattleRaceJersey EnumCattleRace = "Jersey"

	EnumCattleRaceBraunvieh EnumCattleRace = "Braunvieh"

	EnumCattleRaceAngler EnumCattleRace = "Angler"

	EnumCattleRaceOriginalBraunvieh EnumCattleRace = "OriginalBraunvieh"

	EnumCattleRaceRedHolstein EnumCattleRace = "RedHolstein"

	EnumCattleRaceKreuzung EnumCattleRace = "Kreuzung"

	EnumCattleRaceHinterwaelder EnumCattleRace = "Hinterwaelder"

	EnumCattleRaceCharolais EnumCattleRace = "Charolais"

	EnumCattleRaceLimousin EnumCattleRace = "Limousin"

	EnumCattleRaceWeissblaueBelgier EnumCattleRace = "WeissblaueBelgier"

	EnumCattleRaceBlondedAquitaine EnumCattleRace = "BlondedAquitaine"

	EnumCattleRaceMaineAnjou EnumCattleRace = "MaineAnjou"

	EnumCattleRaceSalers EnumCattleRace = "Salers"

	EnumCattleRaceMontbeliard EnumCattleRace = "Montbeliard"

	EnumCattleRaceAubrac EnumCattleRace = "Aubrac"

	EnumCattleRaceGasconne EnumCattleRace = "Gasconne"

	EnumCattleRacePiemontese EnumCattleRace = "Piemontese"

	EnumCattleRaceChianina EnumCattleRace = "Chianina"

	EnumCattleRaceRomagnola EnumCattleRace = "Romagnola"

	EnumCattleRaceMarchigiana EnumCattleRace = "Marchigiana"

	EnumCattleRaceINRA95 EnumCattleRace = "INRA95"

	EnumCattleRaceAngus EnumCattleRace = "Angus"

	EnumCattleRaceHereford EnumCattleRace = "Hereford"

	EnumCattleRaceHighlandCattle EnumCattleRace = "HighlandCattle"

	EnumCattleRaceGalloway EnumCattleRace = "Galloway"

	EnumCattleRaceGuernsey EnumCattleRace = "Guernsey"

	EnumCattleRaceSwissFleckvieh EnumCattleRace = "SwissFleckvieh"

	EnumCattleRaceLuing EnumCattleRace = "Luing"

	EnumCattleRaceKiwicross EnumCattleRace = "Kiwicross"

	EnumCattleRaceNormande EnumCattleRace = "Normande"

	EnumCattleRaceAyrshire EnumCattleRace = "Ayrshire"

	EnumCattleRaceAbondance EnumCattleRace = "Abondance"

	EnumCattleRaceGrauvieh EnumCattleRace = "Grauvieh"

	EnumCattleRaceDexter EnumCattleRace = "Dexter"

	EnumCattleRaceBazadaise EnumCattleRace = "Bazadaise"

	EnumCattleRaceTuxer EnumCattleRace = "Tuxer"

	EnumCattleRaceMurnauWerdenfelser EnumCattleRace = "MurnauWerdenfelser"

	EnumCattleRaceDaenischeRotbunte EnumCattleRace = "DaenischeRotbunte"

	EnumCattleRaceSchwedischeRotbunte EnumCattleRace = "SchwedischeRotbunte"

	EnumCattleRaceNorwegischeRotbunte EnumCattleRace = "NorwegischeRotbunte"

	EnumCattleRacePinzgauer EnumCattleRace = "Pinzgauer"

	EnumCattleRaceSimmental EnumCattleRace = "Simmental"

	EnumCattleRaceDahomey EnumCattleRace = "Dahomey"

	EnumCattleRaceTarentaise EnumCattleRace = "Tarentaise"

	EnumCattleRaceVosgienne EnumCattleRace = "Vosgienne"

	EnumCattleRaceTexasLonghorn EnumCattleRace = "TexasLonghorn"

	EnumCattleRaceGelbvieh EnumCattleRace = "Gelbvieh"

	EnumCattleRaceWasserbueffel EnumCattleRace = "Wasserbueffel"

	EnumCattleRaceBison EnumCattleRace = "Bison"

	EnumCattleRaceYak EnumCattleRace = "Yak"

	EnumCattleRaceEringer EnumCattleRace = "Eringer"

	EnumCattleRaceEvolene EnumCattleRace = "Evolene"

	EnumCattleRaceZebu EnumCattleRace = "Zebu"

	EnumCattleRaceWagyu EnumCattleRace = "Wagyu"

	EnumCattleRaceShorthorn EnumCattleRace = "Shorthorn"

	EnumCattleRaceParthenaise EnumCattleRace = "Parthenaise"

	EnumCattleRaceRotesHoehenvieh EnumCattleRace = "RotesHoehenvieh"

	EnumCattleRaceValdostana EnumCattleRace = "Valdostana"

	EnumCattleRaceZwergzebu EnumCattleRace = "Zwergzebu"

	EnumCattleRaceOther EnumCattleRace = "Other"

	EnumCattleRaceBordelaise EnumCattleRace = "Bordelaise"

	EnumCattleRacePustertalerSprinzen EnumCattleRace = "PustertalerSprinzen"

	EnumCattleRaceLowLine EnumCattleRace = "LowLine"

	EnumCattleRaceWelshBlack EnumCattleRace = "WelshBlack"
)

type TranslatedEnumTypeOfEnumEquidNotificationState struct {
	XMLName              xml.Name //`xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidNotificationState"`

	EnumValue            *EnumEquidNotificationState `xml:"EnumValue,omitempty"`
	RequestedTranslation string                      `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidTypeOfUse struct {
	XMLName              xml.Name //`xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidTypeOfUse"`

	EnumValue            *EnumEquidTypeOfUse `xml:"EnumValue,omitempty"`
	RequestedTranslation string              `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumGender struct {
	XMLName              xml.Name //`xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumGender"`

	EnumValue            *EnumGender `xml:"EnumValue,omitempty"`
	RequestedTranslation string      `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidBreed struct {
	XMLName              xml.Name //`xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidBreed"`

	EnumValue            *EnumEquidBreed `xml:"EnumValue,omitempty"`
	RequestedTranslation string          `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidWithersClass struct {
	XMLName              xml.Name //`xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidWithersClass"`

	EnumValue            *EnumEquidWithersClass `xml:"EnumValue,omitempty"`
	RequestedTranslation string                 `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumPoultryUsageReason struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumPoultryUsageReason"`

	EnumValue            *EnumPoultryUsageReason `xml:"EnumValue,omitempty"`
	RequestedTranslation string                  `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumAnimalHusbandryType struct {
	XMLName              xml.Name // `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumAnimalHusbandryType"`

	EnumValue            *EnumAnimalHusbandryType `xml:"EnumValue,omitempty"`
	RequestedTranslation string                   `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumAnimalHusbandryTypeOfUse struct {
	XMLName              xml.Name // `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumAnimalHusbandryTypeOfUse"`

	EnumValue            *EnumAnimalHusbandryTypeOfUse `xml:"EnumValue,omitempty"`
	RequestedTranslation string                        `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumZone struct {
	XMLName              xml.Name // `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumZone"`

	EnumValue            *EnumZone `xml:"EnumValue,omitempty"`
	RequestedTranslation string    `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumArea struct {
	XMLName              xml.Name // `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumArea"`

	EnumValue            *EnumArea `xml:"EnumValue,omitempty"`
	RequestedTranslation string    `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumPigCategory struct {
	XMLName              xml.Name //`xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumPigCategory"`

	EnumValue            *EnumPigCategory `xml:"EnumValue,omitempty"`
	RequestedTranslation string           `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumOrderStatus struct {
	XMLName              xml.Name //`xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumOrderStatus"`

	EnumValue            *EnumOrderStatus `xml:"EnumValue,omitempty"`
	RequestedTranslation string           `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumCattleRace struct {
	XMLName              xml.Name //`xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumCattleRace"`

	EnumValue            *EnumCattleRace `xml:"EnumValue,omitempty"`
	RequestedTranslation string          `xml:"RequestedTranslation,omitempty"`
}

type EnumEquidLocationChangeType string

const (
	EnumEquidLocationChangeTypeInSwitzerland EnumEquidLocationChangeType = "InSwitzerland"

	EnumEquidLocationChangeTypeOutOfSwitzerland EnumEquidLocationChangeType = "OutOfSwitzerland"

	EnumEquidLocationChangeTypeIntoSwitzerland EnumEquidLocationChangeType = "IntoSwitzerland"
)

type AnimalTracingPortType struct {
	client *SOAPClient
}

func NewAnimalTracingPortType(url string, tls bool, auth *BasicAuth) *AnimalTracingPortType {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &AnimalTracingPortType{
		client: client,
	}
}

func (service *AnimalTracingPortType) Version(request *Version) (*VersionResponse, error) {
	response := new(VersionResponse)
	err := service.client.Call(request.Action, request, response)
	if err != nil {
		return nil, err
	}
	fmt.Println("res: ", response)
	return response, nil
}

func (service *AnimalTracingPortType) WritePigEntryMovementV2(request *WritePigEntryMovementV2) (*WritePigEntryMovementV2Response, error) {
	response := new(WritePigEntryMovementV2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WritePigEntryMovement(request *WritePigEntryMovement) (*WritePigEntryMovementResponse, error) {
	response := new(WritePigEntryMovementResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WritePigSlaughterMovement(request *WritePigSlaughterMovement) (*WritePigSlaughterMovementResponse, error) {
	response := new(WritePigSlaughterMovementResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteGroupSlaughterMovement(request *WriteGroupSlaughterMovement) (*WriteGroupSlaughterMovementResponse, error) {
	response := new(WriteGroupSlaughterMovementResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteGroupSlaughterMovementV2(request *WriteGroupSlaughterMovementV2) (*WriteGroupSlaughterMovementV2Response, error) {
	response := new(WriteGroupSlaughterMovementV2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetEquidOwnershipList(request *GetEquidOwnershipList) (*GetEquidOwnershipListResponse, error) {
	response := new(GetEquidOwnershipListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetEquidLivestock(request *GetEquidLivestock) (*GetEquidLivestockResponse, error) {
	response := new(GetEquidLivestockResponse)
	err := service.client.Call(request.Action, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidRelocationNotification(request *WriteEquidRelocationNotification) (*WriteEquidRelocationNotificationResponse, error) {
	response := new(WriteEquidRelocationNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidAcquireOwnershipNotification(request *WriteEquidAcquireOwnershipNotification) (*WriteEquidAcquireOwnershipNotificationResponse, error) {
	response := new(WriteEquidAcquireOwnershipNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidCedeOwnershipNotification(request *WriteEquidCedeOwnershipNotification) (*WriteEquidCedeOwnershipNotificationResponse, error) {
	response := new(WriteEquidCedeOwnershipNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetFarmers(request *GetFarmers) (*GetFarmersResponse, error) {
	response := new(GetFarmersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetFarmManager(request *GetFarmManager) (*GetFarmManagerResponse, error) {
	response := new(GetFarmManagerResponse)
	err := service.client.Call(request.Action, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetAnimalHusbandryAddress(request *GetAnimalHusbandryAddress) (*GetAnimalHusbandryAddressResponse, error) {
	response := new(GetAnimalHusbandryAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetAnimalHusbandryAddressV2(request *GetAnimalHusbandryAddressV2) (*GetAnimalHusbandryAddressV2Response, error) {
	response := new(GetAnimalHusbandryAddressV2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCodes(request *GetCodes) (*GetCodesResponse, error) {
	response := new(GetCodesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetPersonAddress(request *GetPersonAddress) (*GetPersonAddressResponse, error) {
	response := new(GetPersonAddressResponse)
	err := service.client.Call(request.Action, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetPersonIdentifiers(request *GetPersonIdentifiers) (*GetPersonIdentifiersResponse, error) {
	response := new(GetPersonIdentifiersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetRegisteredGenera(request *GetRegisteredGenera) (*GetRegisteredGeneraResponse, error) {
	response := new(GetRegisteredGeneraResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleArrivalBatchNotification(request *WriteCattleArrivalBatchNotification) (*WriteCattleArrivalBatchNotificationResponse, error) {
	response := new(WriteCattleArrivalBatchNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleDeceasedNotification(request *WriteCattleDeceasedNotification) (*WriteCattleDeceasedNotificationResponse, error) {
	response := new(WriteCattleDeceasedNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleYardSlaughterNotification(request *WriteCattleYardSlaughterNotification) (*WriteCattleYardSlaughterNotificationResponse, error) {
	response := new(WriteCattleYardSlaughterNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleExportNotification(request *WriteCattleExportNotification) (*WriteCattleExportNotificationResponse, error) {
	response := new(WriteCattleExportNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleBirthNotification(request *WriteCattleBirthNotification) (*WriteCattleBirthNotificationResponse, error) {
	response := new(WriteCattleBirthNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleDaystayBatchNotification(request *WriteCattleDaystayBatchNotification) (*WriteCattleDaystayBatchNotificationResponse, error) {
	response := new(WriteCattleDaystayBatchNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleLeavingBatchNotification(request *WriteCattleLeavingBatchNotification) (*WriteCattleLeavingBatchNotificationResponse, error) {
	response := new(WriteCattleLeavingBatchNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetAnimalHusbandryMemberships(request *GetAnimalHusbandryMemberships) (*GetAnimalHusbandryMembershipsResponse, error) {
	response := new(GetAnimalHusbandryMembershipsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetEarTagOrders(request *GetEarTagOrders) (*GetEarTagOrdersResponse, error) {
	response := new(GetEarTagOrdersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleDeformationNotification(request *WriteCattleDeformationNotification) (*WriteCattleDeformationNotificationResponse, error) {
	response := new(WriteCattleDeformationNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) DeleteAnimalHusbandryMemberships(request *DeleteAnimalHusbandryMemberships) (*DeleteAnimalHusbandryMembershipsResponse, error) {
	response := new(DeleteAnimalHusbandryMembershipsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) DeleteCattleNotifications(request *DeleteCattleNotifications) (*DeleteCattleNotificationsResponse, error) {
	response := new(DeleteCattleNotificationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleDeathBirthNotification(request *WriteCattleDeathBirthNotification) (*WriteCattleDeathBirthNotificationResponse, error) {
	response := new(WriteCattleDeathBirthNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleSlaughterBatchNotification(request *WriteCattleSlaughterBatchNotification) (*WriteCattleSlaughterBatchNotificationResponse, error) {
	response := new(WriteCattleSlaughterBatchNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleSlaughterBatchNotificationV2(request *WriteCattleSlaughterBatchNotificationV2) (*WriteCattleSlaughterBatchNotificationV2Response, error) {
	response := new(WriteCattleSlaughterBatchNotificationV2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleEarTags(request *GetCattleEarTags) (*GetCattleEarTagsResponse, error) {
	response := new(GetCattleEarTagsResponse)
	err := service.client.Call(request.Action, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleMovements(request *GetCattleMovements) (*GetCattleMovementsResponse, error) {
	response := new(GetCattleMovementsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleHistory(request *GetCattleHistory) (*GetCattleHistoryResponse, error) {
	response := new(GetCattleHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleLivestock(request *GetCattleLivestock) (*GetCattleLivestockResponse, error) {
	response := new(GetCattleLivestockResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleLivestockV2(request *GetCattleLivestockV2) (*GetCattleLivestockV2Response, error) {
	response := new(GetCattleLivestockV2Response)
	err := service.client.Call(request.Action, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleDetail(request *GetCattleDetail) (*GetCattleDetailResponse, error) {
	response := new(GetCattleDetailResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleDetailV2(request *GetCattleDetailV2) (*GetCattleDetailV2Response, error) {
	response := new(GetCattleDetailV2Response)
	err := service.client.Call(request.Action, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleStatus(request *GetCattleStatus) (*GetCattleStatusResponse, error) {
	response := new(GetCattleStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleStatusV2(request *GetCattleStatusV2) (*GetCattleStatusV2Response, error) {
	response := new(GetCattleStatusV2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteNewEarTagOrder(request *WriteNewEarTagOrder) (*WriteNewEarTagOrderResponse, error) {
	response := new(WriteNewEarTagOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteNewLabelEarTagOrder(request *WriteNewLabelEarTagOrder) (*WriteNewLabelEarTagOrderResponse, error) {
	response := new(WriteNewLabelEarTagOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) CheckCattleForDisposalContribution(request *CheckCattleForDisposalContribution) (*CheckCattleForDisposalContributionResponse, error) {
	response := new(CheckCattleForDisposalContributionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteAnimalClassificationNotification(request *WriteAnimalClassificationNotification) (*WriteAnimalClassificationNotificationResponse, error) {
	response := new(WriteAnimalClassificationNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteAnimalClassificationNotificationV2(request *WriteAnimalClassificationNotificationV2) (*WriteAnimalClassificationNotificationV2Response, error) {
	response := new(WriteAnimalClassificationNotificationV2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) EnableDataAccess(request *EnableDataAccess) (*EnableDataAccessResponse, error) {
	response := new(EnableDataAccessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) DisableDataAccess(request *DisableDataAccess) (*DisableDataAccessResponse, error) {
	response := new(DisableDataAccessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) CheckCattlesForPassport(request *CheckCattlesForPassport) (*CheckCattlesForPassportResponse, error) {
	response := new(CheckCattlesForPassportResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattlePassportOrders(request *WriteCattlePassportOrders) (*WriteCattlePassportOrdersResponse, error) {
	response := new(WriteCattlePassportOrdersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) DeleteEarTagOrder(request *DeleteEarTagOrder) (*DeleteEarTagOrderResponse, error) {
	response := new(DeleteEarTagOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleChangeNameNotification(request *WriteCattleChangeNameNotification) (*WriteCattleChangeNameNotificationResponse, error) {
	response := new(WriteCattleChangeNameNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetLabelEarTagOrders(request *GetLabelEarTagOrders) (*GetLabelEarTagOrdersResponse, error) {
	response := new(GetLabelEarTagOrdersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteCattleTypeOfUseNotification(request *WriteCattleTypeOfUseNotification) (*WriteCattleTypeOfUseNotificationResponse, error) {
	response := new(WriteCattleTypeOfUseNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) DeleteLabelEarTagOrder(request *DeleteLabelEarTagOrder) (*DeleteLabelEarTagOrderResponse, error) {
	response := new(DeleteLabelEarTagOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteReplacementEarTagOrder(request *WriteReplacementEarTagOrder) (*WriteReplacementEarTagOrderResponse, error) {
	response := new(WriteReplacementEarTagOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattleOffsprings(request *GetCattleOffsprings) (*GetCattleOffspringsResponse, error) {
	response := new(GetCattleOffspringsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WritePoultrySlaughterNotification(request *WritePoultrySlaughterNotification) (*WritePoultrySlaughterNotificationResponse, error) {
	response := new(WritePoultrySlaughterNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WritePoultryBarnInNotification(request *WritePoultryBarnInNotification) (*WritePoultryBarnInNotificationResponse, error) {
	response := new(WritePoultryBarnInNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetPoultryBarnInNotifications(request *GetPoultryBarnInNotifications) (*GetPoultryBarnInNotificationsResponse, error) {
	response := new(GetPoultryBarnInNotificationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetAnimalHusbandryDetail(request *GetAnimalHusbandryDetail) (*GetAnimalHusbandryDetailResponse, error) {
	response := new(GetAnimalHusbandryDetailResponse)
	err := service.client.Call("http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1/AnimalTracingPortType/GetAnimalHusbandryDetail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetPigArrivalNotificationForBreedingOrganisation(request *GetPigArrivalNotificationForBreedingOrganisation) (*GetPigArrivalNotificationForBreedingOrganisationResponse, error) {
	response := new(GetPigArrivalNotificationForBreedingOrganisationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetMembershipForOrganisation(request *GetMembershipForOrganisation) (*GetMembershipForOrganisationResponse, error) {
	response := new(GetMembershipForOrganisationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteReplacementBatchEarTagOrder(request *WriteReplacementBatchEarTagOrder) (*WriteReplacementBatchEarTagOrderResponse, error) {
	response := new(WriteReplacementBatchEarTagOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetCattlesPerLeavingDate(request *GetCattlesPerLeavingDate) (*GetCattlesPerLeavingDateResponse, error) {
	response := new(GetCattlesPerLeavingDateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XmlSOAP string   `xml:"xmlns:soap,attr"`
	XmlNS   string   `xml:"xmlns:ns,attr"`

	Header  SOAPHeader
	Body    SOAPBody
}

type SOAPResponseEnvelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`

	Body    SOAPResponseBody
}

type SOAPResponseBody struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPHeader struct {
	XMLName xml.Name    `xml:"soap:Header"`
	XmlWSA  string      `xml:"xmlns:wsa,attr"`
	Action  Action      `xml:"wsa:Action"`
	To      To          `xml:"wsa:To"`
	Content interface{} `xml:",omitempty"`
}

type Action struct {
	ActionUrl string `xml:",innerxml"`
}
type To struct {
	ToUrl string `xml:",innerxml"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"soap:Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Fault"`

	Code    string `xml:"faultcode,omitempty"`
	String  string `xml:"faultstring,omitempty"`
	Actor   string `xml:"faultactor,omitempty"`
	Detail  string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPResponseBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token xml.Token
		err error
		consumed bool
	)

	Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://www.w3.org/2003/05/soap-envelope" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {

	envelope := SOAPEnvelope{
		XmlSOAP: "http://www.w3.org/2003/05/soap-envelope",
		XmlNS:   "http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1",
		Header: SOAPHeader{
			XmlWSA: "http://www.w3.org/2005/08/addressing",
		},
	}
	envelope.Header.Action.ActionUrl = soapAction
	envelope.Header.To.ToUrl = "https://ws.wbf.admin.ch/Livestock/AnimalTracing/1"

	envelope.Body.Content = request

	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")

	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	//log.Println("SOAPResponse:", string(rawbody))
	respEnvelope := new(SOAPResponseEnvelope)
	respEnvelope.Body = SOAPResponseBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}
	return nil
}
