package tvd

//https://github.com/fiorix/wsdl2go
import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
	"github.com/astaxie/beego"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type Version struct {
	XMLName           xml.Name `xml:"ns:Version"`

	P_ManufacturerKey string `xml:"ns:p_ManufacturerKey,omitempty"`
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
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovement"`

	P_ManufacturerKey          string `xml:"p_ManufacturerKey,omitempty"`

	P_ReportingTvdNumber       int32 `xml:"p_ReportingTvdNumber,omitempty"`

	P_LCID                     int32 `xml:"p_LCID,omitempty"`

	P_EventDate                time.Time `xml:"p_EventDate,omitempty"`

	P_NumberOfAnimals          int32 `xml:"p_NumberOfAnimals,omitempty"`

	P_SourceHusbandryTvdNumber int32 `xml:"p_SourceHusbandryTvdNumber,omitempty"`
}

type WritePigEntryMovementResponse struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovementResponse"`

	WritePigEntryMovementResult *WriteMovementResult `xml:"WritePigEntryMovementResult,omitempty"`
}

type WritePigSlaughterMovement struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigSlaughterMovement"`

	P_ManufacturerKey           string `xml:"p_ManufacturerKey,omitempty"`

	P_ReportingTvdNumber        int32 `xml:"p_ReportingTvdNumber,omitempty"`

	P_LCID                      int32 `xml:"p_LCID,omitempty"`

	P_EventDate                 time.Time `xml:"p_EventDate,omitempty"`

	P_SourceHusbandryTvdNumber  int32 `xml:"p_SourceHusbandryTvdNumber,omitempty"`

	P_MessageID                 int64 `xml:"p_MessageID,omitempty"`

	P_ClassifierNumber          int16 `xml:"p_ClassifierNumber,omitempty"`

	P_ClassifierEquipmentID     string `xml:"p_ClassifierEquipmentID,omitempty"`

	P_ContractorNumberSlaughter string `xml:"p_ContractorNumberSlaughter,omitempty"`

	P_MFA                       int16 `xml:"p_MFA,omitempty"`

	P_Weight                    float64 `xml:"p_Weight,omitempty"`

	P_SlaughterInitiator        int32 `xml:"p_SlaughterInitiator,omitempty"`
}

type WritePigSlaughterMovementResponse struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigSlaughterMovementResponse"`

	WritePigSlaughterMovementResult *WriteMovementResult `xml:"WritePigSlaughterMovementResult,omitempty"`
}

type WriteGroupSlaughterMovement struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovement"`

	P_ManufacturerKey          string `xml:"p_ManufacturerKey,omitempty"`

	P_ReportingTvdNumber       int32 `xml:"p_ReportingTvdNumber,omitempty"`

	P_LCID                     int32 `xml:"p_LCID,omitempty"`

	P_EventDate                time.Time `xml:"p_EventDate,omitempty"`

	P_NumberOfAnimals          int32 `xml:"p_NumberOfAnimals,omitempty"`

	P_SourceHusbandryTvdNumber int32 `xml:"p_SourceHusbandryTvdNumber,omitempty"`
}

type WriteGroupSlaughterMovementResponse struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovementResponse"`

	WriteGroupSlaughterMovementResult *WriteMovementResult `xml:"WriteGroupSlaughterMovementResult,omitempty"`
}

type WriteGroupSlaughterMovementV2 struct {
	XMLName                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovementV2"`

	P_ManufacturerKey             string `xml:"p_ManufacturerKey,omitempty"`

	P_ReportingTvdNumber          int32 `xml:"p_ReportingTvdNumber,omitempty"`

	P_LCID                        int32 `xml:"p_LCID,omitempty"`

	P_EventDate                   time.Time `xml:"p_EventDate,omitempty"`

	P_Genus                       int32 `xml:"p_Genus,omitempty"`

	P_NumberOfAnimals             int32 `xml:"p_NumberOfAnimals,omitempty"`

	P_SourceHusbandryTvdNumber    int32 `xml:"p_SourceHusbandryTvdNumber,omitempty"`

	P_SlaughterInitiatorTvdNumber int32 `xml:"p_SlaughterInitiatorTvdNumber,omitempty"`
}

type WriteGroupSlaughterMovementV2Response struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteGroupSlaughterMovementV2Response"`

	WriteGroupSlaughterMovementV2Result *WriteMovementResult `xml:"WriteGroupSlaughterMovementV2Result,omitempty"`
}

type GetSmallAnimalSlaughterList struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetSmallAnimalSlaughterList"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_Request         *SmallAnimalSlaughterRequest `xml:"p_Request,omitempty"`
}

type GetSmallAnimalSlaughterListResponse struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetSmallAnimalSlaughterListResponse"`

	GetSmallAnimalSlaughterListResult *SmallAnimalSlaughterResult `xml:"GetSmallAnimalSlaughterListResult,omitempty"`
}

type GetEquidOwnershipList struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidOwnershipList"`

	P_ManufacturerKey           string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                      int32 `xml:"p_LCID,omitempty"`

	P_EquidOwnershipListRequest *EquidOwnershipListRequest `xml:"p_EquidOwnershipListRequest,omitempty"`
}

type GetEquidOwnershipListResponse struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidOwnershipListResponse"`

	GetEquidOwnershipListResult *GetEquidOwnershipListResult `xml:"GetEquidOwnershipListResult,omitempty"`
}

type GetEquidLivestock struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidLivestock"`

	P_ManufacturerKey       string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                  int32 `xml:"p_LCID,omitempty"`

	P_EquidLivestockRequest *EquidLivestockRequest `xml:"p_EquidLivestockRequest,omitempty"`
}

type GetEquidLivestockResponse struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidLivestockResponse"`

	GetEquidLivestockResult *GetEquidLivestockResult `xml:"GetEquidLivestockResult,omitempty"`
}

type WriteEquidRelocationNotification struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidRelocationNotification"`

	P_ManufacturerKey        string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                   int32 `xml:"p_LCID,omitempty"`

	P_EquidRelocationRequest *EquidRelocationRequest `xml:"p_EquidRelocationRequest,omitempty"`
}

type WriteEquidRelocationNotificationResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidRelocationNotificationResponse"`

	WriteEquidRelocationNotificationResult *WriteNotificationResult `xml:"WriteEquidRelocationNotificationResult,omitempty"`
}

type WriteEquidAcquireOwnershipNotification struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidAcquireOwnershipNotification"`

	P_ManufacturerKey              string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                         int32 `xml:"p_LCID,omitempty"`

	P_EquidAcquireOwnershipRequest *EquidAcquireOwnershipRequest `xml:"p_EquidAcquireOwnershipRequest,omitempty"`
}

type WriteEquidAcquireOwnershipNotificationResponse struct {
	XMLName                                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidAcquireOwnershipNotificationResponse"`

	WriteEquidAcquireOwnershipNotificationResult *WriteNotificationResult `xml:"WriteEquidAcquireOwnershipNotificationResult,omitempty"`
}

type WriteEquidCedeOwnershipNotification struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidCedeOwnershipNotification"`

	P_ManufacturerKey           string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                      int32 `xml:"p_LCID,omitempty"`

	P_EquidCedeOwnershipRequest *EquidCedeOwnershipRequest `xml:"p_EquidCedeOwnershipRequest,omitempty"`
}

type WriteEquidCedeOwnershipNotificationResponse struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidCedeOwnershipNotificationResponse"`

	WriteEquidCedeOwnershipNotificationResult *WriteNotificationResult `xml:"WriteEquidCedeOwnershipNotificationResult,omitempty"`
}

type WriteEquidNewChipNotification struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidNewChipNotification"`

	P_ManufacturerKey     string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                int32 `xml:"p_LCID,omitempty"`

	P_EquidNewChipRequest *EquidNewChipRequest `xml:"p_EquidNewChipRequest,omitempty"`
}

type WriteEquidNewChipNotificationResponse struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidNewChipNotificationResponse"`

	WriteEquidNewChipNotificationResult *WriteNotificationResult `xml:"WriteEquidNewChipNotificationResult,omitempty"`
}

type WriteEquidWithersClassNotification struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidWithersClassNotification"`

	P_ManufacturerKey          string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                     int32 `xml:"p_LCID,omitempty"`

	P_EquidWithersClassRequest *EquidWithersClassRequest `xml:"p_EquidWithersClassRequest,omitempty"`
}

type WriteEquidWithersClassNotificationResponse struct {
	XMLName                                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidWithersClassNotificationResponse"`

	WriteEquidWithersClassNotificationResult *WriteNotificationResult `xml:"WriteEquidWithersClassNotificationResult,omitempty"`
}

type GetEquidDetail struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidDetail"`

	P_ManufacturerKey    string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID               int32 `xml:"p_LCID,omitempty"`

	P_EquidDetailRequest *EquidDetailRequest `xml:"p_EquidDetailRequest,omitempty"`
}

type GetEquidDetailResponse struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidDetailResponse"`

	GetEquidDetailResult *EquidDetailResult `xml:"GetEquidDetailResult,omitempty"`
}

type WriteEquidCastrationNotification struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidCastrationNotification"`

	P_ManufacturerKey        string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                   int32 `xml:"p_LCID,omitempty"`

	P_EquidCastrationRequest *EquidCastrationRequest `xml:"p_EquidCastrationRequest,omitempty"`
}

type WriteEquidCastrationNotificationResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidCastrationNotificationResponse"`

	WriteEquidCastrationNotificationResult *WriteNotificationResult `xml:"WriteEquidCastrationNotificationResult,omitempty"`
}

type WriteEquidSlaughterNotification struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidSlaughterNotification"`

	P_ManufacturerKey       string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                  int32 `xml:"p_LCID,omitempty"`

	P_EquidSlaughterRequest *EquidSlaughterRequest `xml:"p_EquidSlaughterRequest,omitempty"`
}

type WriteEquidSlaughterNotificationResponse struct {
	XMLName                               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidSlaughterNotificationResponse"`

	WriteEquidSlaughterNotificationResult *WriteNotificationResult `xml:"WriteEquidSlaughterNotificationResult,omitempty"`
}

type WriteEquidDeathNotification struct {
	XMLName             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidDeathNotification"`

	P_ManufacturerKey   string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID              int32 `xml:"p_LCID,omitempty"`

	P_EquidDeathRequest *EquidDeathRequest `xml:"p_EquidDeathRequest,omitempty"`
}

type WriteEquidDeathNotificationResponse struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidDeathNotificationResponse"`

	WriteEquidDeathNotificationResult *WriteNotificationResult `xml:"WriteEquidDeathNotificationResult,omitempty"`
}

type WriteEquidImportNotification struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidImportNotification"`

	P_ManufacturerKey    string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID               int32 `xml:"p_LCID,omitempty"`

	P_EquidImportRequest *EquidImportRequest `xml:"p_EquidImportRequest,omitempty"`
}

type WriteEquidImportNotificationResponse struct {
	XMLName                            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidImportNotificationResponse"`

	WriteEquidImportNotificationResult *EquidInitialResult `xml:"WriteEquidImportNotificationResult,omitempty"`
}

type WriteEquidBirthNotification struct {
	XMLName             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidBirthNotification"`

	P_ManufacturerKey   string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID              int32 `xml:"p_LCID,omitempty"`

	P_EquidBirthRequest *EquidBirthRequest `xml:"p_EquidBirthRequest,omitempty"`
}

type WriteEquidBirthNotificationResponse struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidBirthNotificationResponse"`

	WriteEquidBirthNotificationResult *EquidInitialResult `xml:"WriteEquidBirthNotificationResult,omitempty"`
}

type WriteEquidPassIssuerPermissionNotification struct {
	XMLName                            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidPassIssuerPermissionNotification"`

	P_ManufacturerKey                  string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                             int32 `xml:"p_LCID,omitempty"`

	P_EquidPassIssuerPermissionRequest *EquidPassIssuerPermissionRequest `xml:"p_EquidPassIssuerPermissionRequest,omitempty"`
}

type WriteEquidPassIssuerPermissionNotificationResponse struct {
	XMLName                                          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidPassIssuerPermissionNotificationResponse"`

	WriteEquidPassIssuerPermissionNotificationResult *WriteNotificationResult `xml:"WriteEquidPassIssuerPermissionNotificationResult,omitempty"`
}

type WriteEquidOrderBasePassNotification struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidOrderBasePassNotification"`

	P_ManufacturerKey           string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                      int32 `xml:"p_LCID,omitempty"`

	P_EquidOrderBasePassRequest *EquidOrderBasePassRequest `xml:"p_EquidOrderBasePassRequest,omitempty"`
}

type WriteEquidOrderBasePassNotificationResponse struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidOrderBasePassNotificationResponse"`

	WriteEquidOrderBasePassNotificationResult *WriteNotificationResult `xml:"WriteEquidOrderBasePassNotificationResult,omitempty"`
}

type WriteEquidPassIssuedNotification struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidPassIssuedNotification"`

	P_ManufacturerKey        string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                   int32 `xml:"p_LCID,omitempty"`

	P_EquidPassIssuedRequest *EquidPassIssuedRequest `xml:"p_EquidPassIssuedRequest,omitempty"`
}

type WriteEquidPassIssuedNotificationResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidPassIssuedNotificationResponse"`

	WriteEquidPassIssuedNotificationResult *WriteNotificationResult `xml:"WriteEquidPassIssuedNotificationResult,omitempty"`
}

type WriteEquidMembershipOrganisationsNotification struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidMembershipOrganisationsNotification"`

	P_ManufacturerKey               string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                          int32 `xml:"p_LCID,omitempty"`

	P_MembershipOrganisationRequest *EquidMembershipOrganisationRequest `xml:"p_MembershipOrganisationRequest,omitempty"`
}

type WriteEquidMembershipOrganisationsNotificationResponse struct {
	XMLName                                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidMembershipOrganisationsNotificationResponse"`

	WriteEquidMembershipOrganisationsNotificationResult *WriteNotificationResult `xml:"WriteEquidMembershipOrganisationsNotificationResult,omitempty"`
}

type WriteEquidTypeOfUseNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidTypeOfUseNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_Ueln            string `xml:"p_Ueln,omitempty"`
}

type WriteEquidTypeOfUseNotificationResponse struct {
	XMLName                               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidTypeOfUseNotificationResponse"`

	WriteEquidTypeOfUseNotificationResult *WriteNotificationResult `xml:"WriteEquidTypeOfUseNotificationResult,omitempty"`
}

type WriteEquidBasicDataNotification struct {
	XMLName                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidBasicDataNotification"`

	P_ManufacturerKey            string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                       int32 `xml:"p_LCID,omitempty"`

	P_WriteEquidBasicDataRequest *EquidBasicDataRequest `xml:"p_WriteEquidBasicDataRequest,omitempty"`
}

type WriteEquidBasicDataNotificationResponse struct {
	XMLName                               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidBasicDataNotificationResponse"`

	WriteEquidBasicDataNotificationResult *WriteNotificationResult `xml:"WriteEquidBasicDataNotificationResult,omitempty"`
}

type SearchEquid struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SearchEquid"`

	P_ManufacturerKey    string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID               int32 `xml:"p_LCID,omitempty"`

	P_EquidSearchRequest *EquidSearchRequest `xml:"p_EquidSearchRequest,omitempty"`
}

type SearchEquidResponse struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SearchEquidResponse"`

	SearchEquidResult *EquidSearchResult `xml:"SearchEquidResult,omitempty"`
}

type WriteEquidImportAfterExportNotification struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidImportAfterExportNotification"`

	P_ManufacturerKey               string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                          int32 `xml:"p_LCID,omitempty"`

	P_EquidImportAfterExportRequest *EquidImportAfterExportRequest `xml:"p_EquidImportAfterExportRequest,omitempty"`
}

type WriteEquidImportAfterExportNotificationResponse struct {
	XMLName                                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteEquidImportAfterExportNotificationResponse"`

	WriteEquidImportAfterExportNotificationResult *WriteNotificationResult `xml:"WriteEquidImportAfterExportNotificationResult,omitempty"`
}

type GetEquidSlaughterList struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidSlaughterList"`

	P_ManufacturerKey           string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                      int32 `xml:"p_LCID,omitempty"`

	P_EquidSlaughterListRequest *EquidSlaughterListRequest `xml:"p_EquidSlaughterListRequest,omitempty"`
}

type GetEquidSlaughterListResponse struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidSlaughterListResponse"`

	GetEquidSlaughterListResult *EquidSlaughterListResult `xml:"GetEquidSlaughterListResult,omitempty"`
}

type GetEquidsWithNotificationsOfMO struct {
	XMLName                              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidsWithNotificationsOfMO"`

	P_ManufacturerKey                    string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                               int32 `xml:"p_LCID,omitempty"`

	P_EquidsWithNotificationsOfMoRequest *EquidNotificationsForMembershipOrganisationRequest `xml:"p_EquidsWithNotificationsOfMoRequest,omitempty"`
}

type GetEquidsWithNotificationsOfMOResponse struct {
	XMLName                              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidsWithNotificationsOfMOResponse"`

	GetEquidsWithNotificationsOfMOResult *EquidsWithNotificationsOfMOResult `xml:"GetEquidsWithNotificationsOfMOResult,omitempty"`
}

type GetFarmers struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmers"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`
}

type GetFarmersResponse struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmersResponse"`

	GetFarmersResult *PersonListResult `xml:"GetFarmersResult,omitempty"`
}

type GetFarmManager struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmManager"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`
}

type GetFarmManagerResponse struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetFarmManagerResponse"`

	GetFarmManagerResult *FarmManagerResult `xml:"GetFarmManagerResult,omitempty"`
}

type GetAnimalHusbandryAddress struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddress"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`
}

type GetAnimalHusbandryAddressResponse struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddressResponse"`

	GetAnimalHusbandryAddressResult *AnimalHusbandryAddressResult `xml:"GetAnimalHusbandryAddressResult,omitempty"`
}

type GetAnimalHusbandryAddressV2 struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddressV2"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`
}

type GetAnimalHusbandryAddressV2Response struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryAddressV2Response"`

	GetAnimalHusbandryAddressV2Result *AnimalHusbandryAddressResultV2 `xml:"GetAnimalHusbandryAddressV2Result,omitempty"`
}

type GetCodes struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCodes"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_Type            string `xml:"p_Type,omitempty"`
}

type GetCodesResponse struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCodesResponse"`

	GetCodesResult *CodesResult `xml:"GetCodesResult,omitempty"`
}

type GetPersonAddress struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonAddress"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_AgateNumber     string `xml:"p_AgateNumber,omitempty"`
}

type GetPersonAddressResponse struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonAddressResponse"`

	GetPersonAddressResult *PersonAddressResult `xml:"GetPersonAddressResult,omitempty"`
}

type GetPersonIdentifiers struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonIdentifiers"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`
}

type GetPersonIdentifiersResponse struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPersonIdentifiersResponse"`

	GetPersonIdentifiersResult *PersonIdentifiersResult `xml:"GetPersonIdentifiersResult,omitempty"`
}

type GetRegisteredGenera struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetRegisteredGenera"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`
}

type GetRegisteredGeneraResponse struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetRegisteredGeneraResponse"`

	GetRegisteredGeneraResult *GeneraResult `xml:"GetRegisteredGeneraResult,omitempty"`
}

type WriteCattleArrivalBatchNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleArrivalBatchNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_ArrivalData     *CattleArrivalDataArray `xml:"p_ArrivalData,omitempty"`
}

type WriteCattleArrivalBatchNotificationResponse struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleArrivalBatchNotificationResponse"`

	WriteCattleArrivalBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleArrivalBatchNotificationResult,omitempty"`
}

type WriteCattleDeceasedNotification struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeceasedNotification"`

	P_CattleDeceasedRequest *CattleNotificationRequest `xml:"p_CattleDeceasedRequest,omitempty"`
}

type WriteCattleDeceasedNotificationResponse struct {
	XMLName                               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeceasedNotificationResponse"`

	WriteCattleDeceasedNotificationResult *CattleNotificationResult `xml:"WriteCattleDeceasedNotificationResult,omitempty"`
}

type WriteCattleYardSlaughterNotification struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleYardSlaughterNotification"`

	P_CattleYardSlaughteredRequest *CattleNotificationRequest `xml:"p_CattleYardSlaughteredRequest,omitempty"`
}

type WriteCattleYardSlaughterNotificationResponse struct {
	XMLName                                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleYardSlaughterNotificationResponse"`

	WriteCattleYardSlaughterNotificationResult *CattleNotificationResult `xml:"WriteCattleYardSlaughterNotificationResult,omitempty"`
}

type WriteCattleExportNotification struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleExportNotification"`

	P_CattleExportRequest *CattleNotificationWithCountryRequest `xml:"p_CattleExportRequest,omitempty"`
}

type WriteCattleExportNotificationResponse struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleExportNotificationResponse"`

	WriteCattleExportNotificationResult *CattleNotificationResult `xml:"WriteCattleExportNotificationResult,omitempty"`
}

type WriteCattleImportAfterExportNotification struct {
	XMLName                          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleImportAfterExportNotification"`

	P_CattleImportAfterExportRequest *CattleNotificationWithCountryRequest `xml:"p_CattleImportAfterExportRequest,omitempty"`
}

type WriteCattleImportAfterExportNotificationResponse struct {
	XMLName                                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleImportAfterExportNotificationResponse"`

	WriteCattleImportAfterExportNotificationResult *CattleNotificationResult `xml:"WriteCattleImportAfterExportNotificationResult,omitempty"`
}

type WriteCattleBirthNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleBirthNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_BirthData       *CattleBirthData `xml:"p_BirthData,omitempty"`
}

type WriteCattleBirthNotificationResponse struct {
	XMLName                            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleBirthNotificationResponse"`

	WriteCattleBirthNotificationResult *CattleNotificationResult `xml:"WriteCattleBirthNotificationResult,omitempty"`
}

type WriteCattleDaystayBatchNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDaystayBatchNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_DaystayData     *CattleDaystayDataArray `xml:"p_DaystayData,omitempty"`
}

type WriteCattleDaystayBatchNotificationResponse struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDaystayBatchNotificationResponse"`

	WriteCattleDaystayBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleDaystayBatchNotificationResult,omitempty"`
}

type WriteCattleLeavingBatchNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleLeavingBatchNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_LeavingData     *CattleLeavingDataArray `xml:"p_LeavingData,omitempty"`
}

type WriteCattleLeavingBatchNotificationResponse struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleLeavingBatchNotificationResponse"`

	WriteCattleLeavingBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleLeavingBatchNotificationResult,omitempty"`
}

type GetAnimalHusbandryMemberships struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryMemberships"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`
}

type GetAnimalHusbandryMembershipsResponse struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryMembershipsResponse"`

	GetAnimalHusbandryMembershipsResult *HusbandryMembershipResult `xml:"GetAnimalHusbandryMembershipsResult,omitempty"`
}

type GetEarTagOrders struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEarTagOrders"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_SearchDateFrom  time.Time `xml:"p_SearchDateFrom,omitempty"`

	P_SearchDateTo    time.Time `xml:"p_SearchDateTo,omitempty"`

	P_ArticleFilter   *IntArray `xml:"p_ArticleFilter,omitempty"`
}

type GetEarTagOrdersResponse struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEarTagOrdersResponse"`

	GetEarTagOrdersResult *EarTagOrderResult `xml:"GetEarTagOrdersResult,omitempty"`
}

type WriteCattleDeformationNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeformationNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`

	P_DeformationType int32 `xml:"p_DeformationType,omitempty"`
}

type WriteCattleDeformationNotificationResponse struct {
	XMLName                                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeformationNotificationResponse"`

	WriteCattleDeformationNotificationResult *CattleNotificationResult `xml:"WriteCattleDeformationNotificationResult,omitempty"`
}

type DeleteAnimalHusbandryMemberships struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteAnimalHusbandryMemberships"`

	P_ManufacturerKey      string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                 int32 `xml:"p_LCID,omitempty"`

	P_ID_PRG               int32 `xml:"p_ID_PRG,omitempty"`

	P_DeleteMembershipData *IntArray `xml:"p_DeleteMembershipData,omitempty"`
}

type DeleteAnimalHusbandryMembershipsResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteAnimalHusbandryMembershipsResponse"`

	DeleteAnimalHusbandryMembershipsResult *DeleteAnimalHusbandryMembershipResult `xml:"DeleteAnimalHusbandryMembershipsResult,omitempty"`
}

type DeleteCattleNotifications struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteCattleNotifications"`

	P_ManufacturerKey                   string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                              int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber                         int32 `xml:"p_TVDNumber,omitempty"`

	P_DeleteCattleNotificationDataArray *IntArray `xml:"p_DeleteCattleNotificationDataArray,omitempty"`
}

type DeleteCattleNotificationsResponse struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteCattleNotificationsResponse"`

	DeleteCattleNotificationsResult *WriteCattleBatchNotificationResult `xml:"DeleteCattleNotificationsResult,omitempty"`
}

type WriteCattleDeathBirthNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeathBirthNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_DeathBirthData  *CattleDeathBirthData `xml:"p_DeathBirthData,omitempty"`
}

type WriteCattleDeathBirthNotificationResponse struct {
	XMLName                                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleDeathBirthNotificationResponse"`

	WriteCattleDeathBirthNotificationResult *CattleDeathBirthNotificationResult `xml:"WriteCattleDeathBirthNotificationResult,omitempty"`
}

type WriteCattleSlaughterBatchNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_SlaughterData   *CattleArrivalDataArray `xml:"p_SlaughterData,omitempty"`
}

type WriteCattleSlaughterBatchNotificationResponse struct {
	XMLName                                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotificationResponse"`

	WriteCattleSlaughterBatchNotificationResult *WriteCattleBatchNotificationResult `xml:"WriteCattleSlaughterBatchNotificationResult,omitempty"`
}

type WriteCattleSlaughterBatchNotificationV2 struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotificationV2"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_SlaughterData   *CattleSlaughterDataArray `xml:"p_SlaughterData,omitempty"`
}

type WriteCattleSlaughterBatchNotificationV2Response struct {
	XMLName                                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleSlaughterBatchNotificationV2Response"`

	WriteCattleSlaughterBatchNotificationV2Result *WriteCattleBatchNotificationResult `xml:"WriteCattleSlaughterBatchNotificationV2Result,omitempty"`
}

type GetCattleEarTags struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleEarTags"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_OnStock         bool `xml:"p_OnStock,omitempty"`
}

type GetCattleEarTagsResponse struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleEarTagsResponse"`

	GetCattleEarTagsResult *CattleEarTagsResult `xml:"GetCattleEarTagsResult,omitempty"`
}

type GetCattleMovements struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleMovements"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleMovementsResponse struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleMovementsResponse"`

	GetCattleMovementsResult *CattleMovementsResult `xml:"GetCattleMovementsResult,omitempty"`
}

type GetCattleHistory struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleHistory"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleHistoryResponse struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleHistoryResponse"`

	GetCattleHistoryResult *CattleHistoryResult `xml:"GetCattleHistoryResult,omitempty"`
}

type GetCattleHistoryV2 struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleHistoryV2"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleHistoryV2Response struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleHistoryV2Response"`

	GetCattleHistoryV2Result *CattleHistoryResultV2 `xml:"GetCattleHistoryV2Result,omitempty"`
}

type GetCattleLivestock struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestock"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_SearchDateFrom  time.Time `xml:"p_SearchDateFrom,omitempty"`

	P_SearchDateTo    time.Time `xml:"p_SearchDateTo,omitempty"`
}

type GetCattleLivestockResponse struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestockResponse"`

	GetCattleLivestockResult *CattleLivestockResult `xml:"GetCattleLivestockResult,omitempty"`
}

type GetCattleLivestockV2 struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestockV2"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_SearchDateFrom  time.Time `xml:"p_SearchDateFrom,omitempty"`

	P_SearchDateTo    time.Time `xml:"p_SearchDateTo,omitempty"`
}

type GetCattleLivestockV2Response struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleLivestockV2Response"`

	GetCattleLivestockV2Result *CattleLivestockResultV2 `xml:"GetCattleLivestockV2Result,omitempty"`
}

type GetCattleDetail struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetail"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleDetailResponse struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetailResponse"`

	GetCattleDetailResult *CattleDetailResult `xml:"GetCattleDetailResult,omitempty"`
}

type GetCattleDetailV2 struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetailV2"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_WorkingFocus    *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleDetailV2Response struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleDetailV2Response"`

	GetCattleDetailV2Result *CattleDetailResultV2 `xml:"GetCattleDetailV2Result,omitempty"`
}

type GetCattleStatus struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatus"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleStatusResponse struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatusResponse"`

	GetCattleStatusResult *CattleStateExternalResult `xml:"GetCattleStatusResult,omitempty"`
}

type GetCattleStatusV2 struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatusV2"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`
}

type GetCattleStatusV2Response struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleStatusV2Response"`

	GetCattleStatusV2Result *CattleStateExternalResultV2 `xml:"GetCattleStatusV2Result,omitempty"`
}

type WriteNewEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewEarTagOrder"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_EarTagType      int32 `xml:"p_EarTagType,omitempty"`

	P_Amount          int32 `xml:"p_Amount,omitempty"`
}

type WriteNewEarTagOrderResponse struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewEarTagOrderResponse"`

	WriteNewEarTagOrderResult *NewEarTagOrderResult `xml:"WriteNewEarTagOrderResult,omitempty"`
}

type WriteNewLabelEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewLabelEarTagOrder"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_LabelEarTagType int32 `xml:"p_LabelEarTagType,omitempty"`

	P_Amount          int32 `xml:"p_Amount,omitempty"`
}

type WriteNewLabelEarTagOrderResponse struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNewLabelEarTagOrderResponse"`

	WriteNewLabelEarTagOrderResult *NewEarTagOrderResult `xml:"WriteNewLabelEarTagOrderResult,omitempty"`
}

type CheckCattleForDisposalContribution struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattleForDisposalContribution"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_SlaughterData   *CattleArrivalData `xml:"p_SlaughterData,omitempty"`
}

type CheckCattleForDisposalContributionResponse struct {
	XMLName                                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattleForDisposalContributionResponse"`

	CheckCattleForDisposalContributionResult *DisposalContributionResult `xml:"CheckCattleForDisposalContributionResult,omitempty"`
}

type WriteAnimalClassificationNotification struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotification"`

	P_ManufacturerKey          string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                     int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber                int32 `xml:"p_TVDNumber,omitempty"`

	P_AnimalClassificationData *AnimalClassificationData `xml:"p_AnimalClassificationData,omitempty"`
}

type WriteAnimalClassificationNotificationResponse struct {
	XMLName                                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotificationResponse"`

	WriteAnimalClassificationNotificationResult *NotificationResult `xml:"WriteAnimalClassificationNotificationResult,omitempty"`
}

type WriteAnimalClassificationNotificationV2 struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotificationV2"`

	P_ManufacturerKey          string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                     int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber                int32 `xml:"p_TVDNumber,omitempty"`

	P_AnimalClassificationData *AnimalClassificationDataV2 `xml:"p_AnimalClassificationData,omitempty"`
}

type WriteAnimalClassificationNotificationV2Response struct {
	XMLName                                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteAnimalClassificationNotificationV2Response"`

	WriteAnimalClassificationNotificationV2Result *NotificationResult `xml:"WriteAnimalClassificationNotificationV2Result,omitempty"`
}

type EnableDataAccess struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EnableDataAccess"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_ID_PRG          int32 `xml:"p_ID_PRG,omitempty"`
}

type EnableDataAccessResponse struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EnableDataAccessResponse"`

	EnableDataAccessResult *DataAccessResult `xml:"EnableDataAccessResult,omitempty"`
}

type DisableDataAccess struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DisableDataAccess"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_ID_PRG          int32 `xml:"p_ID_PRG,omitempty"`
}

type DisableDataAccessResponse struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DisableDataAccessResponse"`

	DisableDataAccessResult *DataAccessResult `xml:"DisableDataAccessResult,omitempty"`
}

type CheckCattlesForPassport struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattlesForPassport"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_IssueDate       time.Time `xml:"p_IssueDate,omitempty"`

	P_EarTagNumbers   *StringArray `xml:"p_EarTagNumbers,omitempty"`
}

type CheckCattlesForPassportResponse struct {
	XMLName                       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CheckCattlesForPassportResponse"`

	CheckCattlesForPassportResult *CattlePassportResult `xml:"CheckCattlesForPassportResult,omitempty"`
}

type WriteCattlePassportOrders struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattlePassportOrders"`

	P_ManufacturerKey  string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID             int32 `xml:"p_LCID,omitempty"`

	P_PassportLanguage string `xml:"p_PassportLanguage,omitempty"`

	P_IssueDate        time.Time `xml:"p_IssueDate,omitempty"`

	P_EarTagNumbers    *StringArray `xml:"p_EarTagNumbers,omitempty"`
}

type WriteCattlePassportOrdersResponse struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattlePassportOrdersResponse"`

	WriteCattlePassportOrdersResult *CattlePassportResult `xml:"WriteCattlePassportOrdersResult,omitempty"`
}

type DeleteEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteEarTagOrder"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_NotificationID  int32 `xml:"p_NotificationID,omitempty"`
}

type DeleteEarTagOrderResponse struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteEarTagOrderResponse"`

	DeleteEarTagOrderResult *NotificationResult `xml:"DeleteEarTagOrderResult,omitempty"`
}

type WriteCattleChangeNameNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleChangeNameNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`

	P_Name            string `xml:"p_Name,omitempty"`
}

type WriteCattleChangeNameNotificationResponse struct {
	XMLName                                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleChangeNameNotificationResponse"`

	WriteCattleChangeNameNotificationResult *NotificationResult `xml:"WriteCattleChangeNameNotificationResult,omitempty"`
}

type GetLabelEarTagOrders struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetLabelEarTagOrders"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_SearchDateFrom  time.Time `xml:"p_SearchDateFrom,omitempty"`

	P_SearchDateTo    time.Time `xml:"p_SearchDateTo,omitempty"`

	P_ArticleFilter   *IntArray `xml:"p_ArticleFilter,omitempty"`
}

type GetLabelEarTagOrdersResponse struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetLabelEarTagOrdersResponse"`

	GetLabelEarTagOrdersResult *EarTagOrderResult `xml:"GetLabelEarTagOrdersResult,omitempty"`
}

type WriteCattleTypeOfUseNotification struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleTypeOfUseNotification"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`

	P_CattleTypeOfUse int32 `xml:"p_CattleTypeOfUse,omitempty"`

	P_EventDate       time.Time `xml:"p_EventDate,omitempty"`
}

type WriteCattleTypeOfUseNotificationResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleTypeOfUseNotificationResponse"`

	WriteCattleTypeOfUseNotificationResult *NotificationResult `xml:"WriteCattleTypeOfUseNotificationResult,omitempty"`
}

type DeleteLabelEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteLabelEarTagOrder"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_NotificationID  int32 `xml:"p_NotificationID,omitempty"`
}

type DeleteLabelEarTagOrderResponse struct {
	XMLName                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteLabelEarTagOrderResponse"`

	DeleteLabelEarTagOrderResult *NotificationResult `xml:"DeleteLabelEarTagOrderResult,omitempty"`
}

type WriteReplacementEarTagOrder struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementEarTagOrder"`

	P_ManufacturerKey string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID            int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber       int32 `xml:"p_TVDNumber,omitempty"`

	P_Genus           int32 `xml:"p_Genus,omitempty"`

	P_EarTagNumber    string `xml:"p_EarTagNumber,omitempty"`

	P_LeftEarTag      bool `xml:"p_LeftEarTag,omitempty"`

	P_RightEarTag     bool `xml:"p_RightEarTag,omitempty"`
}

type WriteReplacementEarTagOrderResponse struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementEarTagOrderResponse"`

	WriteReplacementEarTagOrderResult *NewEarTagOrderResult `xml:"WriteReplacementEarTagOrderResult,omitempty"`
}

type GetCattleOffsprings struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleOffsprings"`

	P_ManufacturerKey    string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID               int32 `xml:"p_LCID,omitempty"`

	P_WorkingFocus       *WorkingFocusArray `xml:"p_WorkingFocus,omitempty"`

	P_EarTagNumberMother string `xml:"p_EarTagNumberMother,omitempty"`
}

type GetCattleOffspringsResponse struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattleOffspringsResponse"`

	GetCattleOffspringsResult *CattleOffspringResult `xml:"GetCattleOffspringsResult,omitempty"`
}

type WritePoultrySlaughterNotification struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultrySlaughterNotification"`

	P_ManufacturerKey    string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID               int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber          int32 `xml:"p_TVDNumber,omitempty"`

	P_EventDate          time.Time `xml:"p_EventDate,omitempty"`

	P_Weight             int32 `xml:"p_Weight,omitempty"`

	P_PoultryType        *EnumPoultryType `xml:"p_PoultryType,omitempty"`

	P_SourceTVDNumber    int32 `xml:"p_SourceTVDNumber,omitempty"`

	P_DeliveryFileNumber string `xml:"p_DeliveryFileNumber,omitempty"`
}

type WritePoultrySlaughterNotificationResponse struct {
	XMLName                                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultrySlaughterNotificationResponse"`

	WritePoultrySlaughterNotificationResult *ProcessingResult `xml:"WritePoultrySlaughterNotificationResult,omitempty"`
}

type WritePoultryBarnInNotification struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultryBarnInNotification"`

	P_ManufacturerKey           string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                      int32 `xml:"p_LCID,omitempty"`

	P_TVDNumber                 int32 `xml:"p_TVDNumber,omitempty"`

	P_PoultryBarnInNotification *PoultryBarnInNotification `xml:"p_PoultryBarnInNotification,omitempty"`
}

type WritePoultryBarnInNotificationResponse struct {
	XMLName                              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePoultryBarnInNotificationResponse"`

	WritePoultryBarnInNotificationResult *ProcessingResult `xml:"WritePoultryBarnInNotificationResult,omitempty"`
}

type GetPoultryBarnInNotifications struct {
	XMLName                                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPoultryBarnInNotifications"`

	P_SearchPoultryBarnInNotificationRequest *SearchPoultryBarnInNotificationRequest `xml:"p_SearchPoultryBarnInNotificationRequest,omitempty"`
}

type GetPoultryBarnInNotificationsResponse struct {
	XMLName                             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPoultryBarnInNotificationsResponse"`

	GetPoultryBarnInNotificationsResult *GetPoultryBarnInNotificationResult `xml:"GetPoultryBarnInNotificationsResult,omitempty"`
}

type GetAnimalHusbandryDetail struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryDetail"`

	P_GetAnimalHusbandryDetailRequest *GetAnimalHusbandryDetailRequest `xml:"p_GetAnimalHusbandryDetailRequest,omitempty"`
}

type GetAnimalHusbandryDetailResponse struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryDetailResponse"`

	GetAnimalHusbandryDetailResult *GetAnimalHusbandryDetailResult `xml:"GetAnimalHusbandryDetailResult,omitempty"`
}

type GetPigArrivalNotificationForBreedingOrganisation struct {
	XMLName                            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPigArrivalNotificationForBreedingOrganisation"`

	P_SearchSmallAnimalMovementRequest *SearchSmallAnimalMovementRequest `xml:"p_SearchSmallAnimalMovementRequest,omitempty"`
}

type GetPigArrivalNotificationForBreedingOrganisationResponse struct {
	XMLName                                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPigArrivalNotificationForBreedingOrganisationResponse"`

	GetPigArrivalNotificationForBreedingOrganisationResult *PigArrivalNotificationResult `xml:"GetPigArrivalNotificationForBreedingOrganisationResult,omitempty"`
}

type GetMembershipForOrganisation struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetMembershipForOrganisation"`

	P_Request *GetMembershipForOrganisationRequest `xml:"p_Request,omitempty"`
}

type GetMembershipForOrganisationResponse struct {
	XMLName                            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetMembershipForOrganisationResponse"`

	GetMembershipForOrganisationResult *GetMembershipForOrganisationResult `xml:"GetMembershipForOrganisationResult,omitempty"`
}

type WriteReplacementBatchEarTagOrder struct {
	XMLName                                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementBatchEarTagOrder"`

	P_ManufacturerKey                         string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                                    int32 `xml:"p_LCID,omitempty"`

	P_WriteReplacementBatchEarTagOrderRequest *WriteReplacementBatchEarTagOrderRequest `xml:"p_WriteReplacementBatchEarTagOrderRequest,omitempty"`
}

type WriteReplacementBatchEarTagOrderResponse struct {
	XMLName                                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementBatchEarTagOrderResponse"`

	WriteReplacementBatchEarTagOrderResult *ReplacementEarTagOrdersResult `xml:"WriteReplacementBatchEarTagOrderResult,omitempty"`
}

type GetCattlesPerLeavingDate struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDate"`

	P_ManufacturerKey                 string `xml:"p_ManufacturerKey,omitempty"`

	P_LCID                            int32 `xml:"p_LCID,omitempty"`

	P_GetCattlesPerLeavingDateRequest *GetCattlesPerLeavingDateRequest `xml:"p_GetCattlesPerLeavingDateRequest,omitempty"`
}

type GetCattlesPerLeavingDateResponse struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDateResponse"`

	GetCattlesPerLeavingDateResult *GetCattlesPerLeavingDateResult `xml:"GetCattlesPerLeavingDateResult,omitempty"`
}

type WritePigEntryMovementV2Request struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WritePigEntryMovementV2Request"`

	*BaseRequest

	ReportingTVDNumber       int32 `xml:"ReportingTVDNumber,omitempty"`

	EventDate                time.Time `xml:"EventDate,omitempty"`

	NumberOfAnimals          int32 `xml:"NumberOfAnimals,omitempty"`

	SourceHusbandryTVDNumber int32 `xml:"SourceHusbandryTVDNumber,omitempty"`

	PigCategory              *EnumPigCategory `xml:"PigCategory,omitempty"`
}

type BaseRequest struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 BaseRequest"`

	ManufacturerKey string `xml:"ManufacturerKey,omitempty"`

	LCID            int32 `xml:"LCID,omitempty"`
}

type WriteMovementResult struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteMovementResult"`

	MovementId       int32 `xml:"MovementId,omitempty"`

	ProcessingResult *ProcessingResult `xml:"ProcessingResult,omitempty"`
}

type ProcessingResult struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ProcessingResult"`

	Code        int32 `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`

	Status      int32 `xml:"Status,omitempty"`
}

type SmallAnimalSlaughterRequest struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SmallAnimalSlaughterRequest"`

	FromDate           time.Time `xml:"FromDate,omitempty"`

	ToDate             time.Time `xml:"ToDate,omitempty"`

	SlaughterTVDNumber int32 `xml:"SlaughterTVDNumber,omitempty"`

	Genus              *EnumGenus `xml:"Genus,omitempty"`
}

type SmallAnimalSlaughterResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SmallAnimalSlaughterResult"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	SlaughterList *ArrayOfSmallAnimalSlaughterListData `xml:"SlaughterList,omitempty"`
}

type ArrayOfSmallAnimalSlaughterListData struct {
	XMLName                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfSmallAnimalSlaughterListData"`

	SmallAnimalSlaughterListData []*SmallAnimalSlaughterListData `xml:"SmallAnimalSlaughterListData,omitempty"`
}

type SmallAnimalSlaughterListData struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SmallAnimalSlaughterListData"`

	EventDate              time.Time `xml:"EventDate,omitempty"`

	NotificationDate       time.Time `xml:"NotificationDate,omitempty"`

	Amount                 int32 `xml:"Amount,omitempty"`

	Genus                  *EnumGenus `xml:"Genus,omitempty"`

	SourceTVDNumber        string `xml:"SourceTVDNumber,omitempty"`

	SlaughterInitTVDNumber string `xml:"SlaughterInitTVDNumber,omitempty"`

	CreatorBZVAccount      string `xml:"CreatorBZVAccount,omitempty"`

	DeleterBZVAccount      string `xml:"DeleterBZVAccount,omitempty"`
}

type EquidOwnershipListRequest struct {
	XMLName                                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidOwnershipListRequest"`

	OnBehalfOfAgateNumber                      string `xml:"OnBehalfOfAgateNumber,omitempty"`

	SearchDateFrom                             time.Time `xml:"SearchDateFrom,omitempty"`

	SearchDateTo                               time.Time `xml:"SearchDateTo,omitempty"`

	ShowOnlyEquidsWhichWereLivingInQueryPeriod bool `xml:"ShowOnlyEquidsWhichWereLivingInQueryPeriod,omitempty"`
}

type GetEquidOwnershipListResult struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidOwnershipListResult"`

	Result    *ProcessingResult `xml:"Result,omitempty"`

	EquidList *ArrayOfEquidItem `xml:"EquidList,omitempty"`
}

type ArrayOfEquidItem struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidItem"`

	EquidItem []*EquidItem `xml:"EquidItem,omitempty"`
}

type EquidItem struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidItem"`

	OwnerFirstName    string `xml:"OwnerFirstName,omitempty"`

	OwnerLastName     string `xml:"OwnerLastName,omitempty"`

	OwnerAgateNumber  string `xml:"OwnerAgateNumber,omitempty"`

	UELN              string `xml:"UELN,omitempty"`

	Name              string `xml:"Name,omitempty"`

	OriginTVDNumber   int32 `xml:"OriginTVDNumber,omitempty"`

	Husbandry         *HusbandryResult `xml:"Husbandry,omitempty"`

	NotificationState *TranslatedEnumTypeOfEnumEquidNotificationState `xml:"NotificationState,omitempty"`

	TypeOfUse         *TranslatedEnumTypeOfEnumEquidTypeOfUse `xml:"TypeOfUse,omitempty"`

	Gender            *TranslatedEnumTypeOfEnumGender `xml:"Gender,omitempty"`

	Breed             *TranslatedEnumTypeOfEnumEquidBreed `xml:"Breed,omitempty"`

	WithersClass      *TranslatedEnumTypeOfEnumEquidWithersClass `xml:"WithersClass,omitempty"`

	IsPassPresent     bool `xml:"IsPassPresent,omitempty"`

	ColorFreeText     string `xml:"ColorFreeText,omitempty"`

	BirthDate         time.Time `xml:"BirthDate,omitempty"`

	DeathDate         time.Time `xml:"DeathDate,omitempty"`

	ArrivalDate       time.Time `xml:"ArrivalDate,omitempty"`

	LeavingDate       time.Time `xml:"LeavingDate,omitempty"`
}

type HusbandryResult struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryResult"`

	TVDNumber string `xml:"TVDNumber,omitempty"`

	Name      string `xml:"Name,omitempty"`

	Street    string `xml:"Street,omitempty"`

	PostCode  string `xml:"PostCode,omitempty"`

	City      string `xml:"City,omitempty"`
}

type EquidLivestockRequest struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidLivestockRequest"`

	TVDNumber      int32 `xml:"TVDNumber,omitempty"`

	SearchDateFrom time.Time `xml:"SearchDateFrom,omitempty"`

	SearchDateTo   time.Time `xml:"SearchDateTo,omitempty"`
}

type GetEquidLivestockResult struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetEquidLivestockResult"`

	Result    *ProcessingResult `xml:"Result,omitempty"`

	EquidList *ArrayOfEquidItem `xml:"EquidList,omitempty"`
}

type EquidRelocationRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidRelocationRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`

	EventDate             time.Time `xml:"EventDate,omitempty"`

	LocationChangeType    *EnumEquidLocationChangeType `xml:"LocationChangeType,omitempty"`

	NewTVDNumber          int32 `xml:"NewTVDNumber,omitempty"`

	OriginTVDNumber       int32 `xml:"OriginTVDNumber,omitempty"`

	ImportOrExportCountry *EnumCountry `xml:"ImportOrExportCountry,omitempty"`
}

type WriteNotificationResult struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteNotificationResult"`

	Result  *ProcessingResult `xml:"Result,omitempty"`
}

type EquidAcquireOwnershipRequest struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidAcquireOwnershipRequest"`

	OnBehalfOfAgateNumber  string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                   string `xml:"Ueln,omitempty"`

	EventDate              time.Time `xml:"EventDate,omitempty"`

	ActualOwnerAgateNumber string `xml:"ActualOwnerAgateNumber,omitempty"`
}

type EquidCedeOwnershipRequest struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidCedeOwnershipRequest"`

	OnBehalfOfAgateNumber       string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                        string `xml:"Ueln,omitempty"`

	EventDate                   time.Time `xml:"EventDate,omitempty"`

	IsNewEquidOwnerLivingAbroad bool `xml:"IsNewEquidOwnerLivingAbroad,omitempty"`

	CededToPersonAgateNumber    string `xml:"CededToPersonAgateNumber,omitempty"`
}

type EquidNewChipRequest struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidNewChipRequest"`

	OnBehalfOfAgateNumber  string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                   string `xml:"Ueln,omitempty"`

	ChipNumber             string `xml:"ChipNumber,omitempty"`

	ImplementationDate     time.Time `xml:"ImplementationDate,omitempty"`

	ImplementationLocation string `xml:"ImplementationLocation,omitempty"`
}

type EquidWithersClassRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidWithersClassRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`

	WithersClass          *EnumEquidWithersClass `xml:"WithersClass,omitempty"`
}

type EquidDetailRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidDetailRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	UserRole              *EnumRole `xml:"UserRole,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`
}

type EquidDetailResult struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidDetailResult"`

	Result          *ProcessingResult `xml:"Result,omitempty"`

	EquidDetailData *EquidDetailData `xml:"EquidDetailData,omitempty"`
}

type EquidDetailData struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidDetailData"`

	EquidBasicData                  *EquidBasicData `xml:"EquidBasicData,omitempty"`

	EquidAdditonalInformationData   *EquidAdditonalInformationData `xml:"EquidAdditonalInformationData,omitempty"`

	EquidOwnershipHistoryDataList   *ArrayOfEquidOwnershipHistoryData `xml:"EquidOwnershipHistoryDataList,omitempty"`

	EquidHusbandryHistoryDataList   *ArrayOfEquidHusbandryHistoryData `xml:"EquidHusbandryHistoryDataList,omitempty"`

	EquidDescendantDataList         *ArrayOfEquidDescendantData `xml:"EquidDescendantDataList,omitempty"`

	EquidPassDataList               *ArrayOfEquidPassData `xml:"EquidPassDataList,omitempty"`

	EquidUelnDataList               *ArrayOfEquidUelnData `xml:"EquidUelnDataList,omitempty"`

	EquidChipDataList               *ArrayOfEquidChipData `xml:"EquidChipDataList,omitempty"`

	EquidMembershipDataList         *ArrayOfEquidMembershipData `xml:"EquidMembershipDataList,omitempty"`

	EquidRudimentaryDescriptionData *EquidRudimentaryDescriptionData `xml:"EquidRudimentaryDescriptionData,omitempty"`
}

type EquidBasicData struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidBasicData"`

	Name              string `xml:"Name,omitempty"`

	SportBreedingName string `xml:"SportBreedingName,omitempty"`

	BirthDate         time.Time `xml:"BirthDate,omitempty"`

	HerdBookNumber    string `xml:"HerdBookNumber,omitempty"`

	PlaceOfBirth      string `xml:"PlaceOfBirth,omitempty"`

	ColorFreeText     string `xml:"ColorFreeText,omitempty"`

	EquidColor        *EnumEquidColor `xml:"EquidColor,omitempty"`

	EquidGender       *TranslatedEnumTypeOfEnumGender `xml:"EquidGender,omitempty"`

	EquidSpecies      *TranslatedEnumTypeOfEnumEquidSpecies `xml:"EquidSpecies,omitempty"`

	EquidBreed        *TranslatedEnumTypeOfEnumEquidBreed `xml:"EquidBreed,omitempty"`

	UelnMother        string `xml:"UelnMother,omitempty"`

	UelnGeneticMother string `xml:"UelnGeneticMother,omitempty"`
}

type EquidAdditonalInformationData struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidAdditonalInformationData"`

	IsCastrated                       bool `xml:"IsCastrated,omitempty"`

	EquidTypeOfUsage                  *TranslatedEnumTypeOfEnumEquidTypeOfUse `xml:"EquidTypeOfUsage,omitempty"`

	DeathDate                         time.Time `xml:"DeathDate,omitempty"`

	IsPassIssuerPermittedToChangeData bool `xml:"IsPassIssuerPermittedToChangeData,omitempty"`

	EquidWithersClass                 *TranslatedEnumTypeOfEnumEquidWithersClass `xml:"EquidWithersClass,omitempty"`

	EquidNotificationState            *TranslatedEnumTypeOfEnumEquidNotificationState `xml:"EquidNotificationState,omitempty"`

	EquidPassIssuingState             *TranslatedEnumTypeOfEnumEquidPassIssuingState `xml:"EquidPassIssuingState,omitempty"`
}

type ArrayOfEquidOwnershipHistoryData struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidOwnershipHistoryData"`

	EquidOwnershipHistoryData []*EquidOwnershipHistoryData `xml:"EquidOwnershipHistoryData,omitempty"`
}

type EquidOwnershipHistoryData struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidOwnershipHistoryData"`

	FromDate              time.Time `xml:"FromDate,omitempty"`

	ToDate                time.Time `xml:"ToDate,omitempty"`

	Name                  string `xml:"Name,omitempty"`

	StreetAndStreetNumber string `xml:"StreetAndStreetNumber,omitempty"`

	Postcode              string `xml:"Postcode,omitempty"`

	City                  string `xml:"City,omitempty"`

	AgateNumber           string `xml:"AgateNumber,omitempty"`

	AGIS                  int32 `xml:"AGIS,omitempty"`
}

type ArrayOfEquidHusbandryHistoryData struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidHusbandryHistoryData"`

	EquidHusbandryHistoryData []*EquidHusbandryHistoryData `xml:"EquidHusbandryHistoryData,omitempty"`
}

type EquidHusbandryHistoryData struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidHusbandryHistoryData"`

	FromDate              time.Time `xml:"FromDate,omitempty"`

	ToDate                time.Time `xml:"ToDate,omitempty"`

	Name                  string `xml:"Name,omitempty"`

	StreetAndStreetNumber string `xml:"StreetAndStreetNumber,omitempty"`

	Postcode              string `xml:"Postcode,omitempty"`

	City                  string `xml:"City,omitempty"`

	TVDNumber             string `xml:"TVDNumber,omitempty"`

	AnimalHusbandryType   *TranslatedEnumTypeOfEnumAnimalHusbandryType `xml:"AnimalHusbandryType,omitempty"`

	NotificationType      *TranslatedEnumTypeOfEnumEquidNotificationType `xml:"NotificationType,omitempty"`
}

type ArrayOfEquidDescendantData struct {
	XMLName             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidDescendantData"`

	EquidDescendantData []*EquidDescendantData `xml:"EquidDescendantData,omitempty"`
}

type EquidDescendantData struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidDescendantData"`

	Name             string `xml:"Name,omitempty"`

	Ueln             string `xml:"Ueln,omitempty"`

	BirthDate        time.Time `xml:"BirthDate,omitempty"`

	Gender           *TranslatedEnumTypeOfEnumGender `xml:"Gender,omitempty"`

	EquidBreed       *TranslatedEnumTypeOfEnumEquidBreed `xml:"EquidBreed,omitempty"`

	EquidTypeOfUsage *TranslatedEnumTypeOfEnumEquidTypeOfUse `xml:"EquidTypeOfUsage,omitempty"`
}

type ArrayOfEquidPassData struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidPassData"`

	EquidPassData []*EquidPassData `xml:"EquidPassData,omitempty"`
}

type EquidPassData struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidPassData"`

	PassType         *TranslatedEnumTypeOfEnumEquidPassType `xml:"PassType,omitempty"`

	IssueDate        time.Time `xml:"IssueDate,omitempty"`

	AgateNumber      string `xml:"AgateNumber,omitempty"`

	PassportIssuer   string `xml:"PassportIssuer,omitempty"`

	NotificationDate time.Time `xml:"NotificationDate,omitempty"`
}

type ArrayOfEquidUelnData struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidUelnData"`

	EquidUelnData []*EquidUelnData `xml:"EquidUelnData,omitempty"`
}

type EquidUelnData struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidUelnData"`

	Ueln             string `xml:"Ueln,omitempty"`

	NotificationDate time.Time `xml:"NotificationDate,omitempty"`

	AgateNumber      string `xml:"AgateNumber,omitempty"`

	FirstName        string `xml:"FirstName,omitempty"`

	LastName         string `xml:"LastName,omitempty"`
}

type ArrayOfEquidChipData struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidChipData"`

	EquidChipData []*EquidChipData `xml:"EquidChipData,omitempty"`
}

type EquidChipData struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidChipData"`

	ImplementationDateTime time.Time `xml:"ImplementationDateTime,omitempty"`

	ChipNumber             string `xml:"ChipNumber,omitempty"`

	ImplementationLocation string `xml:"ImplementationLocation,omitempty"`

	AgateNumber            string `xml:"AgateNumber,omitempty"`

	FirstName              string `xml:"FirstName,omitempty"`

	LastName               string `xml:"LastName,omitempty"`
}

type ArrayOfEquidMembershipData struct {
	XMLName             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidMembershipData"`

	EquidMembershipData []*EquidMembershipData `xml:"EquidMembershipData,omitempty"`
}

type EquidMembershipData struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidMembershipData"`

	MembershipOrganisationAgateNumber string `xml:"MembershipOrganisationAgateNumber,omitempty"`

	MembershipOrganisationName        string `xml:"MembershipOrganisationName,omitempty"`
}

type EquidRudimentaryDescriptionData struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidRudimentaryDescriptionData"`

	HasWhiteOnHead          bool `xml:"HasWhiteOnHead,omitempty"`

	HasWhiteOnLegFrontLeft  bool `xml:"HasWhiteOnLegFrontLeft,omitempty"`

	HasWhiteOnLegFrontRight bool `xml:"HasWhiteOnLegFrontRight,omitempty"`

	HasWhiteOnLegBackLeft   bool `xml:"HasWhiteOnLegBackLeft,omitempty"`

	HasWhiteOnLegBackRight  bool `xml:"HasWhiteOnLegBackRight,omitempty"`
}

type EquidCastrationRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidCastrationRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`

	EventDate             time.Time `xml:"EventDate,omitempty"`
}

type EquidSlaughterRequest struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidSlaughterRequest"`

	TVDNumber                      int32 `xml:"TVDNumber,omitempty"`

	Ueln                           string `xml:"Ueln,omitempty"`

	ChipNumber                     string `xml:"ChipNumber,omitempty"`

	EventDate                      time.Time `xml:"EventDate,omitempty"`

	OriginAnimalHusbandryTVDNumber int32 `xml:"OriginAnimalHusbandryTVDNumber,omitempty"`

	SlaughterInitiatorTVDNumber    int32 `xml:"SlaughterInitiatorTVDNumber,omitempty"`
}

type EquidDeathRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidDeathRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`

	DeathNotificationType *EnumEquidDeathNotificationType `xml:"DeathNotificationType,omitempty"`

	EventDate             time.Time `xml:"EventDate,omitempty"`
}

type EquidImportRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidImportRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`

	EventDate             time.Time `xml:"EventDate,omitempty"`

	ChipNumber            string `xml:"ChipNumber,omitempty"`

	HerdBookNumber        string `xml:"HerdBookNumber,omitempty"`

	Country               *EnumCountry `xml:"Country,omitempty"`

	TVDNumber             int32 `xml:"TVDNumber,omitempty"`

	Name                  string `xml:"Name,omitempty"`

	SportBreedingName     string `xml:"SportBreedingName,omitempty"`

	BirthDate             time.Time `xml:"BirthDate,omitempty"`

	EquidSpecies          *EnumEquidSpecies `xml:"EquidSpecies,omitempty"`

	EquidBreed            *EnumEquidBreed `xml:"EquidBreed,omitempty"`

	EquidColor            string `xml:"EquidColor,omitempty"`

	Gender                *EnumGender `xml:"Gender,omitempty"`

	IsCastrated           bool `xml:"IsCastrated,omitempty"`

	EquidTypeOfUse        *EnumEquidTypeOfUse `xml:"EquidTypeOfUse,omitempty"`

	IsPassAvailable       bool `xml:"IsPassAvailable,omitempty"`

	UelnGeneticMother     string `xml:"UelnGeneticMother,omitempty"`
}

type EquidInitialResult struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidInitialResult"`

	Ueln    string `xml:"Ueln,omitempty"`

	Result  *ProcessingResult `xml:"Result,omitempty"`
}

type EquidBirthRequest struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidBirthRequest"`

	OnBehalfOfAgateNumber       string `xml:"OnBehalfOfAgateNumber,omitempty"`

	BirthDate                   time.Time `xml:"BirthDate,omitempty"`

	HerdBookNumber              string `xml:"HerdBookNumber,omitempty"`

	TVDNumber                   int32 `xml:"TVDNumber,omitempty"`

	IsMultipleBirth             bool `xml:"IsMultipleBirth,omitempty"`

	Name                        string `xml:"Name,omitempty"`

	EquidSpecies                *EnumEquidSpecies `xml:"EquidSpecies,omitempty"`

	EquidBreed                  *EnumEquidBreed `xml:"EquidBreed,omitempty"`

	EquidColor                  *EnumEquidColor `xml:"EquidColor,omitempty"`

	Gender                      *EnumGender `xml:"Gender,omitempty"`

	UelnMother                  string `xml:"UelnMother,omitempty"`

	UelnGeneticMother           string `xml:"UelnGeneticMother,omitempty"`

	EquidRudimentaryDescription *EquidRudimentaryDescriptionData `xml:"EquidRudimentaryDescription,omitempty"`

	IsPassAvailable             bool `xml:"IsPassAvailable,omitempty"`
}

type EquidPassIssuerPermissionRequest struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidPassIssuerPermissionRequest"`

	OnBehalfOfAgateNumber    string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                     string `xml:"Ueln,omitempty"`

	IsEquidPassIssuerAllowed bool `xml:"IsEquidPassIssuerAllowed,omitempty"`
}

type EquidOrderBasePassRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidOrderBasePassRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`

	EquidPassOrderType    *EnumEquidPassOrderType `xml:"EquidPassOrderType,omitempty"`
}

type EquidPassIssuedRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidPassIssuedRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`

	EquidPassType         *EnumEquidPassType `xml:"EquidPassType,omitempty"`

	EventDate             time.Time `xml:"EventDate,omitempty"`
}

type EquidMembershipOrganisationRequest struct {
	XMLName                            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidMembershipOrganisationRequest"`

	OnBehalfOfAgateNumber              string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                               string `xml:"Ueln,omitempty"`

	MembershipOrganisationAgateNumbers *ArrayOfstring `xml:"MembershipOrganisationAgateNumbers,omitempty"`
}

type EquidBasicDataRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidBasicDataRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`

	EquidChangeBasicData  *EquidChangeBasicData `xml:"EquidChangeBasicData,omitempty"`
}

type EquidChangeBasicData struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidChangeBasicData"`

	Name              string `xml:"Name,omitempty"`

	SportBreedingName string `xml:"SportBreedingName,omitempty"`

	HerdBookNumber    string `xml:"HerdBookNumber,omitempty"`

	BirthDate         time.Time `xml:"BirthDate,omitempty"`

	PlaceOfBirth      string `xml:"PlaceOfBirth,omitempty"`

	ColorFreeText     string `xml:"ColorFreeText,omitempty"`

	EquidColor        *EnumEquidColor `xml:"EquidColor,omitempty"`

	EquidGender       *EnumGender `xml:"EquidGender,omitempty"`

	EquidSpecies      *EnumEquidSpecies `xml:"EquidSpecies,omitempty"`

	EquidBreed        *EnumEquidBreed `xml:"EquidBreed,omitempty"`

	UelnMother        string `xml:"UelnMother,omitempty"`

	UelnGeneticMother string `xml:"UelnGeneticMother,omitempty"`
}

type EquidSearchRequest struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidSearchRequest"`

	OnBehalfOfAgateNumber    string `xml:"OnBehalfOfAgateNumber,omitempty"`

	UserRole                 *EnumRole `xml:"UserRole,omitempty"`

	Ueln                     string `xml:"Ueln,omitempty"`

	ChipNumber               string `xml:"ChipNumber,omitempty"`

	AnimalHusbandryTVDNumber int32 `xml:"AnimalHusbandryTVDNumber,omitempty"`

	Name                     string `xml:"Name,omitempty"`

	HerdBookNumber           string `xml:"HerdBookNumber,omitempty"`

	OwnerAgateNumber         string `xml:"OwnerAgateNumber,omitempty"`

	SportBreedingName        string `xml:"SportBreedingName,omitempty"`

	BirthDateFrom            time.Time `xml:"BirthDateFrom,omitempty"`

	BirthDateTo              time.Time `xml:"BirthDateTo,omitempty"`

	EquidBreed               *EnumEquidBreed `xml:"EquidBreed,omitempty"`

	ShowOnlyLivingEquids     bool `xml:"ShowOnlyLivingEquids,omitempty"`
}

type EquidSearchResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidSearchResult"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	SearchResults *ArrayOfEquidSearchResultData `xml:"SearchResults,omitempty"`
}

type ArrayOfEquidSearchResultData struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidSearchResultData"`

	EquidSearchResultData []*EquidSearchResultData `xml:"EquidSearchResultData,omitempty"`
}

type EquidSearchResultData struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidSearchResultData"`

	Name              string `xml:"Name,omitempty"`

	Ueln              string `xml:"Ueln,omitempty"`

	ChipNumber        string `xml:"ChipNumber,omitempty"`

	OwnerAgateNumber  string `xml:"OwnerAgateNumber,omitempty"`

	HerdBookNumber    string `xml:"HerdBookNumber,omitempty"`

	TVDNumber         string `xml:"TVDNumber,omitempty"`

	BirthDate         time.Time `xml:"BirthDate,omitempty"`

	SportBreedingName string `xml:"SportBreedingName,omitempty"`

	DeathDate         time.Time `xml:"DeathDate,omitempty"`

	EquidTypeOfUse    *TranslatedEnumTypeOfEnumEquidTypeOfUse `xml:"EquidTypeOfUse,omitempty"`
}

type EquidImportAfterExportRequest struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidImportAfterExportRequest"`

	OnBehalfOfAgateNumber string `xml:"OnBehalfOfAgateNumber,omitempty"`

	Ueln                  string `xml:"Ueln,omitempty"`

	EventDate             time.Time `xml:"EventDate,omitempty"`

	ImportCountry         *EnumCountry `xml:"ImportCountry,omitempty"`

	TVDNumber             int32 `xml:"TVDNumber,omitempty"`
}

type EquidSlaughterListRequest struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidSlaughterListRequest"`

	TVDNumber int32 `xml:"TVDNumber,omitempty"`

	DateFrom  time.Time `xml:"DateFrom,omitempty"`

	DateTo    time.Time `xml:"DateTo,omitempty"`
}

type EquidSlaughterListResult struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidSlaughterListResult"`

	Result                     *ProcessingResult `xml:"Result,omitempty"`

	EquidSlaughterListDataList *ArrayOfEquidSlaughterListData `xml:"EquidSlaughterListDataList,omitempty"`
}

type ArrayOfEquidSlaughterListData struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidSlaughterListData"`

	EquidSlaughterListData []*EquidSlaughterListData `xml:"EquidSlaughterListData,omitempty"`
}

type EquidSlaughterListData struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidSlaughterListData"`

	Name                              string `xml:"Name,omitempty"`

	Ueln                              string `xml:"Ueln,omitempty"`

	ChipNumber                        string `xml:"ChipNumber,omitempty"`

	OriginAnimalHusbandryTVDNumber    int32 `xml:"OriginAnimalHusbandryTVDNumber,omitempty"`

	BirthDate                         time.Time `xml:"BirthDate,omitempty"`

	SlaughterDate                     time.Time `xml:"SlaughterDate,omitempty"`

	HasDisposalFee                    bool `xml:"HasDisposalFee,omitempty"`

	AnimalHusbandrySlaughterInitiator int32 `xml:"AnimalHusbandrySlaughterInitiator,omitempty"`
}

type EquidNotificationsForMembershipOrganisationRequest struct {
	XMLName                           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidNotificationsForMembershipOrganisationRequest"`

	MembershipOrganisationAgateNumber string `xml:"MembershipOrganisationAgateNumber,omitempty"`

	DateFrom                          time.Time `xml:"DateFrom,omitempty"`

	DateTo                            time.Time `xml:"DateTo,omitempty"`
}

type EquidsWithNotificationsOfMOResult struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidsWithNotificationsOfMOResult"`

	Result    *ProcessingResult `xml:"Result,omitempty"`

	EquidList *ArrayOfEquidItemMO `xml:"EquidList,omitempty"`
}

type ArrayOfEquidItemMO struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidItemMO"`

	EquidItemMO []*EquidItemMO `xml:"EquidItemMO,omitempty"`
}

type EquidItemMO struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidItemMO"`

	EquidUelnItemList *ArrayOfEquidUelnItem `xml:"EquidUelnItemList,omitempty"`

	NotificationList  *ArrayOfEquidNotificationItem `xml:"NotificationList,omitempty"`
}

type ArrayOfEquidUelnItem struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidUelnItem"`

	EquidUelnItem []*EquidUelnItem `xml:"EquidUelnItem,omitempty"`
}

type EquidUelnItem struct {
	XMLName  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidUelnItem"`

	Ueln     string `xml:"Ueln,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`
}

type ArrayOfEquidNotificationItem struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfEquidNotificationItem"`

	EquidNotificationItem []*EquidNotificationItem `xml:"EquidNotificationItem,omitempty"`
}

type EquidNotificationItem struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EquidNotificationItem"`

	NotificationType *TranslatedEnumTypeOfEnumEquidNotificationType `xml:"NotificationType,omitempty"`

	NotificationDate time.Time `xml:"NotificationDate,omitempty"`

	EventDate        time.Time `xml:"EventDate,omitempty"`

	DeletionDate     time.Time `xml:"DeletionDate,omitempty"`
}

type PersonListResult struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonListResult"`

	Result          *ProcessingResult `xml:"Result,omitempty"`

	PersonDataItems *PersonDataArray `xml:"PersonDataItems,omitempty"`
}

type PersonDataArray struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonDataArray"`

	PersonDataItem []*PersonResult `xml:"PersonDataItem,omitempty"`
}

type PersonResult struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonResult"`

	AgateNumber           string `xml:"AgateNumber,omitempty"`

	LastName              string `xml:"LastName,omitempty"`

	FirstName             string `xml:"FirstName,omitempty"`

	Street                string `xml:"Street,omitempty"`

	PostCode              string `xml:"PostCode,omitempty"`

	City                  string `xml:"City,omitempty"`

	EmailAddress          string `xml:"EmailAddress,omitempty"`

	PhoneNumbers          *StringArray `xml:"PhoneNumbers,omitempty"`

	PreferredLanguageLCID int32 `xml:"PreferredLanguageLCID,omitempty"`
}

type StringArray struct {
	XMLName    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 StringArray"`

	StringItem []string `xml:"StringItem,omitempty"`
}

type FarmManagerResult struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 FarmManagerResult"`

	AgateNumber           string `xml:"AgateNumber,omitempty"`

	Title                 string `xml:"Title,omitempty"`

	LastName              string `xml:"LastName,omitempty"`

	FirstName             string `xml:"FirstName,omitempty"`

	Street                string `xml:"Street,omitempty"`

	PostCode              string `xml:"PostCode,omitempty"`

	City                  string `xml:"City,omitempty"`

	EmailAddress          string `xml:"EmailAddress,omitempty"`

	PhoneNumbers          *StringArray `xml:"PhoneNumbers,omitempty"`

	PreferredLanguageLCID int32 `xml:"PreferredLanguageLCID,omitempty"`

	Result                *ProcessingResult `xml:"Result,omitempty"`
}

type AnimalHusbandryAddressResult struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 AnimalHusbandryAddressResult"`

	Result                   *ProcessingResult `xml:"Result,omitempty"`

	PostData                 *HusbandryResult `xml:"PostData,omitempty"`

	IsActive                 bool `xml:"IsActive,omitempty"`

	MunicipalityName         string `xml:"MunicipalityName,omitempty"`

	CantonShortname          string `xml:"CantonShortname,omitempty"`

	CoordinateX              int32 `xml:"CoordinateX,omitempty"`

	CoordinateY              int32 `xml:"CoordinateY,omitempty"`

	Altitude                 int32 `xml:"Altitude,omitempty"`

	CantonAnimalHusbandryKey string `xml:"CantonAnimalHusbandryKey,omitempty"`

	CantonPersonKey          string `xml:"CantonPersonKey,omitempty"`

	BurNumber                string `xml:"BurNumber,omitempty"`

	AnimalHusbandryType      int32 `xml:"AnimalHusbandryType,omitempty"`

	MunicipalityNumber       int32 `xml:"MunicipalityNumber,omitempty"`

	TypeOfUse                int32 `xml:"TypeOfUse,omitempty"`
}

type AnimalHusbandryAddressResultV2 struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 AnimalHusbandryAddressResultV2"`

	Result                   *ProcessingResult `xml:"Result,omitempty"`

	PostData                 *HusbandryResult `xml:"PostData,omitempty"`

	IsActive                 bool `xml:"IsActive,omitempty"`

	MunicipalityName         string `xml:"MunicipalityName,omitempty"`

	CantonShortname          string `xml:"CantonShortname,omitempty"`

	CoordinateX              int32 `xml:"CoordinateX,omitempty"`

	CoordinateY              int32 `xml:"CoordinateY,omitempty"`

	Altitude                 int32 `xml:"Altitude,omitempty"`

	CantonAnimalHusbandryKey string `xml:"CantonAnimalHusbandryKey,omitempty"`

	CantonPersonKey          string `xml:"CantonPersonKey,omitempty"`

	BurNumber                string `xml:"BurNumber,omitempty"`

	AnimalHusbandryType      int32 `xml:"AnimalHusbandryType,omitempty"`

	MunicipalityNumber       int32 `xml:"MunicipalityNumber,omitempty"`

	TypeOfUse                int32 `xml:"TypeOfUse,omitempty"`

	IsMountain               bool `xml:"IsMountain,omitempty"`
}

type CodesResult struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CodesResult"`

	Result  *ProcessingResult `xml:"Result,omitempty"`

	Codes   *CodeArray `xml:"Codes,omitempty"`
}

type CodeArray struct {
	XMLName  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CodeArray"`

	CodeItem []*Code `xml:"CodeItem,omitempty"`
}

type Code struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 Code"`

	Key     int32 `xml:"Key,omitempty"`

	Text_de string `xml:"Text_de,omitempty"`

	Text_fr string `xml:"Text_fr,omitempty"`

	Text_it string `xml:"Text_it,omitempty"`
}

type PersonAddressResult struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonAddressResult"`

	Result      *ProcessingResult `xml:"Result,omitempty"`

	Title       string `xml:"Title,omitempty"`

	PostAddress *PersonResult `xml:"PostAddress,omitempty"`
}

type PersonIdentifiersResult struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PersonIdentifiersResult"`

	Result          *ProcessingResult `xml:"Result,omitempty"`

	CantonPersonKey string `xml:"CantonPersonKey,omitempty"`

	ID_AGIS         int32 `xml:"ID_AGIS,omitempty"`
}

type GeneraResult struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GeneraResult"`

	Result  *ProcessingResult `xml:"Result,omitempty"`

	Genera  *IntArray `xml:"Genera,omitempty"`
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

	EarTagNumber    string `xml:"EarTagNumber,omitempty"`

	EventDate       time.Time `xml:"EventDate,omitempty"`

	TVDNumberOrigin int32 `xml:"TVDNumberOrigin,omitempty"`
}

type WriteCattleBatchNotificationResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteCattleBatchNotificationResult"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	Resultdetails *CattleNotificationResultArray `xml:"Resultdetails,omitempty"`
}

type CattleNotificationResultArray struct {
	XMLName                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleNotificationResultArray"`

	CattleNotificationResultItem []*CattleNotificationResult `xml:"CattleNotificationResultItem,omitempty"`
}

type CattleNotificationResult struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleNotificationResult"`

	EarTagNumber   string `xml:"EarTagNumber,omitempty"`

	NotificationID int32 `xml:"NotificationID,omitempty"`

	Result         *ProcessingResult `xml:"Result,omitempty"`
}

type CattleNotificationRequest struct {
	XMLName      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleNotificationRequest"`

	*BaseRequest

	TVDNumber    int32 `xml:"TVDNumber,omitempty"`

	EarTagNumber string `xml:"EarTagNumber,omitempty"`

	EventDate    time.Time `xml:"EventDate,omitempty"`
}

type CattleNotificationWithCountryRequest struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleNotificationWithCountryRequest"`

	*CattleNotificationRequest

	Country *EnumCountry `xml:"Country,omitempty"`
}

type CattleBirthData struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleBirthData"`

	EarTagNumber              string `xml:"EarTagNumber,omitempty"`

	IsMultipleBirth           bool `xml:"IsMultipleBirth,omitempty"`

	RaceColor                 int32 `xml:"RaceColor,omitempty"`

	Gender                    int32 `xml:"Gender,omitempty"`

	BirthDate                 time.Time `xml:"BirthDate,omitempty"`

	Race                      int32 `xml:"Race,omitempty"`

	EarTagNumberFather        string `xml:"EarTagNumberFather,omitempty"`

	EarTagNumberMother        string `xml:"EarTagNumberMother,omitempty"`

	BreedingOrganisation      int32 `xml:"BreedingOrganisation,omitempty"`

	Name                      string `xml:"Name,omitempty"`

	BirthProcess              int32 `xml:"BirthProcess,omitempty"`

	IsCastrated               bool `xml:"IsCastrated,omitempty"`

	InseminationDate          time.Time `xml:"InseminationDate,omitempty"`

	EarTagNumberGeneticMother string `xml:"EarTagNumberGeneticMother,omitempty"`

	BirthWeight               int32 `xml:"BirthWeight,omitempty"`

	IsPassportDesired         bool `xml:"IsPassportDesired,omitempty"`
}

type CattleDaystayDataArray struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDaystayDataArray"`

	CattleDaystayDataItem []*CattleDaystayData `xml:"CattleDaystayDataItem,omitempty"`
}

type CattleDaystayData struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDaystayData"`

	EarTagNumber    string `xml:"EarTagNumber,omitempty"`

	EventDate       time.Time `xml:"EventDate,omitempty"`

	TVDNumberOrigin int32 `xml:"TVDNumberOrigin,omitempty"`
}

type CattleLeavingDataArray struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLeavingDataArray"`

	CattleLeavingDataItem []*CattleLeavingData `xml:"CattleLeavingDataItem,omitempty"`
}

type CattleLeavingData struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLeavingData"`

	EarTagNumber                string `xml:"EarTagNumber,omitempty"`

	EventDate                   time.Time `xml:"EventDate,omitempty"`

	LeavingReason               int32 `xml:"LeavingReason,omitempty"`

	MainLeavingReasonBreeding   int32 `xml:"MainLeavingReasonBreeding,omitempty"`

	SecondLeavingReasonBreeding int32 `xml:"SecondLeavingReasonBreeding,omitempty"`
}

type HusbandryMembershipResult struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryMembershipResult"`

	Result      *ProcessingResult `xml:"Result,omitempty"`

	Memberships *HusbandryMembershipArray `xml:"Memberships,omitempty"`
}

type HusbandryMembershipArray struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryMembershipArray"`

	HusbandryMembershipItem []*HusbandryMembership `xml:"HusbandryMembershipItem,omitempty"`
}

type HusbandryMembership struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryMembership"`

	ID_PRG      int32 `xml:"ID_PRG,omitempty"`

	AgateNumber string `xml:"AgateNumber,omitempty"`

	LastName    string `xml:"LastName,omitempty"`

	FirstName   string `xml:"FirstName,omitempty"`

	Role        string `xml:"Role,omitempty"`

	Genus       int32 `xml:"Genus,omitempty"`
}

type EarTagOrderResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EarTagOrderResult"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	Resultdetails *EarTagOrderDataArray `xml:"Resultdetails,omitempty"`
}

type EarTagOrderDataArray struct {
	XMLName             xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EarTagOrderDataArray"`

	EarTagOrderDataItem []*EarTagOrderData `xml:"EarTagOrderDataItem,omitempty"`
}

type EarTagOrderData struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EarTagOrderData"`

	NotificationID   int32 `xml:"NotificationID,omitempty"`

	EarTagType       int32 `xml:"EarTagType,omitempty"`

	Amount           int32 `xml:"Amount,omitempty"`

	IsExpress        bool `xml:"IsExpress,omitempty"`

	OrderStatus      int32 `xml:"OrderStatus,omitempty"`

	OrderStatusDate  time.Time `xml:"OrderStatusDate,omitempty"`

	EarTagNumberFrom string `xml:"EarTagNumberFrom,omitempty"`

	EarTagNumberTo   string `xml:"EarTagNumberTo,omitempty"`

	Text1            string `xml:"Text1,omitempty"`

	Text2            string `xml:"Text2,omitempty"`
}

type DeleteAnimalHusbandryMembershipResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DeleteAnimalHusbandryMembershipResult"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	Resultdetails *HusbandryNotificationResultArray `xml:"Resultdetails,omitempty"`
}

type HusbandryNotificationResultArray struct {
	XMLName                         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryNotificationResultArray"`

	HusbandryNotificationResultItem []*HusbandryNotificationResult `xml:"HusbandryNotificationResultItem,omitempty"`
}

type HusbandryNotificationResult struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 HusbandryNotificationResult"`

	TVDNumber      int32 `xml:"TVDNumber,omitempty"`

	NotificationID int32 `xml:"NotificationID,omitempty"`

	Result         *ProcessingResult `xml:"Result,omitempty"`
}

type CattleDeathBirthData struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDeathBirthData"`

	DeathBirthTime            int32 `xml:"DeathBirthTime,omitempty"`

	IsMultipleBirth           bool `xml:"IsMultipleBirth,omitempty"`

	RaceColor                 int32 `xml:"RaceColor,omitempty"`

	Gender                    int32 `xml:"Gender,omitempty"`

	BirthDate                 time.Time `xml:"BirthDate,omitempty"`

	Race                      int32 `xml:"Race,omitempty"`

	EarTagNumberFather        string `xml:"EarTagNumberFather,omitempty"`

	EarTagNumberMother        string `xml:"EarTagNumberMother,omitempty"`

	BreedingOrganisation      int32 `xml:"BreedingOrganisation,omitempty"`

	Name                      string `xml:"Name,omitempty"`

	BirthProcess              int32 `xml:"BirthProcess,omitempty"`

	IsCastrated               bool `xml:"IsCastrated,omitempty"`

	InseminationDate          time.Time `xml:"InseminationDate,omitempty"`

	EarTagNumberGeneticMother string `xml:"EarTagNumberGeneticMother,omitempty"`

	BirthWeight               int32 `xml:"BirthWeight,omitempty"`

	IsPassportDesired         bool `xml:"IsPassportDesired,omitempty"`
}

type CattleDeathBirthNotificationResult struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDeathBirthNotificationResult"`

	NotificationID int32 `xml:"NotificationID,omitempty"`

	Result         *ProcessingResult `xml:"Result,omitempty"`
}

type CattleSlaughterDataArray struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleSlaughterDataArray"`

	CattleSlaughterDataItem []*CattleSlaughterData `xml:"CattleSlaughterDataItem,omitempty"`
}

type CattleSlaughterData struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleSlaughterData"`

	EarTagNumber                string `xml:"EarTagNumber,omitempty"`

	EventDate                   time.Time `xml:"EventDate,omitempty"`

	TVDNumberOrigin             int32 `xml:"TVDNumberOrigin,omitempty"`

	TVDNumberSlaughterInitiator int32 `xml:"TVDNumberSlaughterInitiator,omitempty"`
}

type CattleEarTagsResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleEarTagsResult"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	CattleEarTags *CattleEarTagDataArray `xml:"CattleEarTags,omitempty"`
}

type CattleEarTagDataArray struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleEarTagDataArray"`

	CattleEarTagDataItem []*CattleEarTagData `xml:"CattleEarTagDataItem,omitempty"`
}

type CattleEarTagData struct {
	XMLName      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleEarTagData"`

	EarTagNumber string `xml:"EarTagNumber,omitempty"`

	OrderDate    time.Time `xml:"OrderDate,omitempty"`

	EarTagState  int32 `xml:"EarTagState,omitempty"`
}

type WorkingFocusArray struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WorkingFocusArray"`

	WorkingFocusItem []*WorkingFocus `xml:"WorkingFocusItem,omitempty"`
}

type WorkingFocus struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WorkingFocus"`

	WorkingFocusType int32 `xml:"WorkingFocusType,omitempty"`

	TVDNumber        int32 `xml:"TVDNumber,omitempty"`

	MandateGiver     string `xml:"MandateGiver,omitempty"`
}

type CattleMovementsResult struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleMovementsResult"`

	Result             *ProcessingResult `xml:"Result,omitempty"`

	CattleHistoryState int32 `xml:"CattleHistoryState,omitempty"`

	EarTagNumber       string `xml:"EarTagNumber,omitempty"`

	CattleMovements    *CattleMovementDataArray `xml:"CattleMovements,omitempty"`
}

type CattleMovementDataArray struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleMovementDataArray"`

	CattleMovementDataItem []*CattleMovementData `xml:"CattleMovementDataItem,omitempty"`
}

type CattleMovementData struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleMovementData"`

	NotificationID    int32 `xml:"NotificationID,omitempty"`

	EventDate         time.Time `xml:"EventDate,omitempty"`

	NotificationDate  time.Time `xml:"NotificationDate,omitempty"`

	NotificationType  int32 `xml:"NotificationType,omitempty"`

	TVDNumberNotifier int32 `xml:"TVDNumberNotifier,omitempty"`

	TVDNumberOrigin   int32 `xml:"TVDNumberOrigin,omitempty"`
}

type CattleHistoryResult struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleHistoryResult"`

	Result             *ProcessingResult `xml:"Result,omitempty"`

	EarTagNumber       string `xml:"EarTagNumber,omitempty"`

	RemainingQuota     int32 `xml:"RemainingQuota,omitempty"`

	CattleHistoryState int32 `xml:"CattleHistoryState,omitempty"`

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

	ArrivalNotificationType int32 `xml:"ArrivalNotificationType,omitempty"`

	CountryOrigin           string `xml:"CountryOrigin,omitempty"`

	TVDNumberOrigin         int32 `xml:"TVDNumberOrigin,omitempty"`

	TVDNumberStay           int32 `xml:"TVDNumberStay,omitempty"`

	StayAddress             string `xml:"StayAddress,omitempty"`

	LeavingDate             time.Time `xml:"LeavingDate,omitempty"`

	LeavingNotificationDate time.Time `xml:"LeavingNotificationDate,omitempty"`

	LeavingNotificationType int32 `xml:"LeavingNotificationType,omitempty"`

	CattleStayState         int32 `xml:"CattleStayState,omitempty"`
}

type CattleHistoryResultV2 struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleHistoryResultV2"`

	Result             *ProcessingResult `xml:"Result,omitempty"`

	EarTagNumber       string `xml:"EarTagNumber,omitempty"`

	RemainingQuota     int32 `xml:"RemainingQuota,omitempty"`

	CattleHistoryState int32 `xml:"CattleHistoryState,omitempty"`

	CattleStays        *CattleStayDataArrayV2 `xml:"CattleStays,omitempty"`
}

type CattleStayDataArrayV2 struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStayDataArrayV2"`

	CattleStayDataItemV2 []*CattleStayDataV2 `xml:"CattleStayDataItemV2,omitempty"`
}

type CattleStayDataV2 struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStayDataV2"`

	ArrivalDate             time.Time `xml:"ArrivalDate,omitempty"`

	ArrivalNotificationDate time.Time `xml:"ArrivalNotificationDate,omitempty"`

	ArrivalNotificationType int32 `xml:"ArrivalNotificationType,omitempty"`

	CountryOrigin           string `xml:"CountryOrigin,omitempty"`

	TVDNumberOrigin         int32 `xml:"TVDNumberOrigin,omitempty"`

	TVDNumberStay           int32 `xml:"TVDNumberStay,omitempty"`

	AnhName                 string `xml:"AnhName,omitempty"`

	Street                  string `xml:"Street,omitempty"`

	ZipCode                 int32 `xml:"ZipCode,omitempty"`

	City                    string `xml:"City,omitempty"`

	LeavingDate             time.Time `xml:"LeavingDate,omitempty"`

	LeavingNotificationDate time.Time `xml:"LeavingNotificationDate,omitempty"`

	LeavingNotificationType int32 `xml:"LeavingNotificationType,omitempty"`

	CattleStayState         int32 `xml:"CattleStayState,omitempty"`
}

type CattleLivestockResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockResult"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	TVDNumber     int32 `xml:"TVDNumber,omitempty"`

	Resultdetails *CattleLivestockDataArray `xml:"Resultdetails,omitempty"`
}

type CattleLivestockDataArray struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockDataArray"`

	CattleLivestockDataItem []*CattleLivestockData `xml:"CattleLivestockDataItem,omitempty"`
}

type CattleLivestockData struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockData"`

	EarTagNumber       string `xml:"EarTagNumber,omitempty"`

	Name               string `xml:"Name,omitempty"`

	Gender             int32 `xml:"Gender,omitempty"`

	BirthDate          time.Time `xml:"BirthDate,omitempty"`

	Race               int32 `xml:"Race,omitempty"`

	RaceColor          int32 `xml:"RaceColor,omitempty"`

	ArrivalDate        time.Time `xml:"ArrivalDate,omitempty"`

	LeavingDate        time.Time `xml:"LeavingDate,omitempty"`

	PendulumHusbandry  int32 `xml:"PendulumHusbandry,omitempty"`

	BvdState           string `xml:"BvdState,omitempty"`

	BtState            string `xml:"BtState,omitempty"`

	CattleTypeOfUse    int32 `xml:"CattleTypeOfUse,omitempty"`

	TypeOfUseDate      time.Time `xml:"TypeOfUseDate,omitempty"`

	FirstCalvingDate   time.Time `xml:"FirstCalvingDate,omitempty"`

	CattleHistoryState int32 `xml:"CattleHistoryState,omitempty"`
}

type CattleLivestockResultV2 struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockResultV2"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	TVDNumber     int32 `xml:"TVDNumber,omitempty"`

	Resultdetails *CattleLivestockDataArrayV2 `xml:"Resultdetails,omitempty"`
}

type CattleLivestockDataArrayV2 struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockDataArrayV2"`

	CattleLivestockDataItem []*CattleLivestockDataV2 `xml:"CattleLivestockDataItem,omitempty"`
}

type CattleLivestockDataV2 struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleLivestockDataV2"`

	EarTagNumber       string `xml:"EarTagNumber,omitempty"`

	Name               string `xml:"Name,omitempty"`

	Gender             int32 `xml:"Gender,omitempty"`

	BirthDate          time.Time `xml:"BirthDate,omitempty"`

	Race               int32 `xml:"Race,omitempty"`

	RaceColor          int32 `xml:"RaceColor,omitempty"`

	ArrivalDate        time.Time `xml:"ArrivalDate,omitempty"`

	LeavingDate        time.Time `xml:"LeavingDate,omitempty"`

	PendulumHusbandry  int32 `xml:"PendulumHusbandry,omitempty"`

	BvdState           string `xml:"BvdState,omitempty"`

	BtState            string `xml:"BtState,omitempty"`

	CattleTypeOfUse    int32 `xml:"CattleTypeOfUse,omitempty"`

	TypeOfUseDate      time.Time `xml:"TypeOfUseDate,omitempty"`

	FirstCalvingDate   time.Time `xml:"FirstCalvingDate,omitempty"`

	CattleHistoryState int32 `xml:"CattleHistoryState,omitempty"`

	LastCalvingDate    time.Time `xml:"LastCalvingDate,omitempty"`
}

type CattleDetailResult struct {
	XMLName      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDetailResult"`

	Result       *ProcessingResult `xml:"Result,omitempty"`

	CattleDetail *CattleDetailData `xml:"CattleDetail,omitempty"`
}

type CattleDetailData struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDetailData"`

	BirthNotificationData *CattleBirthData `xml:"BirthNotificationData,omitempty"`

	LongName              string `xml:"LongName,omitempty"`

	RaceFather            int32 `xml:"RaceFather,omitempty"`

	RaceColorFather       int32 `xml:"RaceColorFather,omitempty"`

	NameFather            string `xml:"NameFather,omitempty"`

	RaceMother            int32 `xml:"RaceMother,omitempty"`

	RaceColorMother       int32 `xml:"RaceColorMother,omitempty"`

	NameMother            string `xml:"NameMother,omitempty"`

	Deformations          *IntArray `xml:"Deformations,omitempty"`

	DeathDate             time.Time `xml:"DeathDate,omitempty"`

	BvdState              string `xml:"BvdState,omitempty"`

	BtState               string `xml:"BtState,omitempty"`

	CattleTypeOfUse       int32 `xml:"CattleTypeOfUse,omitempty"`

	TypeOfUseDate         time.Time `xml:"TypeOfUseDate,omitempty"`

	FirstCalvingDate      time.Time `xml:"FirstCalvingDate,omitempty"`

	CurrentHusbandry      int32 `xml:"CurrentHusbandry,omitempty"`

	AllYearHusbandry      int32 `xml:"AllYearHusbandry,omitempty"`

	PendulumHusbandry     int32 `xml:"PendulumHusbandry,omitempty"`

	Beefiness             string `xml:"Beefiness,omitempty"`

	SlaughterCategory     string `xml:"SlaughterCategory,omitempty"`

	FatTissue             string `xml:"FatTissue,omitempty"`
}

type CattleDetailResultV2 struct {
	XMLName      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleDetailResultV2"`

	Result       *ProcessingResult `xml:"Result,omitempty"`

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

	CattleState *CattleStateData `xml:"CattleState,omitempty"`
}

type CattleStateData struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStateData"`

	EarTagNumber       string `xml:"EarTagNumber,omitempty"`

	Name               string `xml:"Name,omitempty"`

	BvdState           string `xml:"BvdState,omitempty"`

	CattleHistoryState int32 `xml:"CattleHistoryState,omitempty"`
}

type CattleStateExternalResultV2 struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStateExternalResultV2"`

	Result      *ProcessingResult `xml:"Result,omitempty"`

	CattleState *CattleStateDataV2 `xml:"CattleState,omitempty"`
}

type CattleStateDataV2 struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleStateDataV2"`

	EarTagNumber       string `xml:"EarTagNumber,omitempty"`

	Name               string `xml:"Name,omitempty"`

	BvdState           string `xml:"BvdState,omitempty"`

	CattleHistoryState int32 `xml:"CattleHistoryState,omitempty"`

	CattleBirthDate    time.Time `xml:"CattleBirthDate,omitempty"`

	CattleAgeInDays    int32 `xml:"CattleAgeInDays,omitempty"`
}

type NewEarTagOrderResult struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 NewEarTagOrderResult"`

	Result      *ProcessingResult `xml:"Result,omitempty"`

	EarTagOrder *EarTagOrderData `xml:"EarTagOrder,omitempty"`
}

type DisposalContributionResult struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DisposalContributionResult"`

	DisposalContributionState bool `xml:"DisposalContributionState,omitempty"`

	Result                    *ProcessingResult `xml:"Result,omitempty"`
}

type AnimalClassificationData struct {
	XMLName                     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 AnimalClassificationData"`

	MessageID                   int64 `xml:"MessageID,omitempty"`

	EarTagNumber                string `xml:"EarTagNumber,omitempty"`

	TVDNumberOrigin             int32 `xml:"TVDNumberOrigin,omitempty"`

	SlaughterDate               time.Time `xml:"SlaughterDate,omitempty"`

	TVDNumberSlaughterInitiator int32 `xml:"TVDNumberSlaughterInitiator,omitempty"`

	Genus                       int32 `xml:"Genus,omitempty"`

	Weight                      float64 `xml:"Weight,omitempty"`

	ClassifierNumber            int32 `xml:"ClassifierNumber,omitempty"`

	ClassifierEquipmentID       string `xml:"ClassifierEquipmentID,omitempty"`

	ContractorNumberSlaughter   string `xml:"ContractorNumberSlaughter,omitempty"`

	SlaughterCategory           string `xml:"SlaughterCategory,omitempty"`

	Beefiness                   string `xml:"Beefiness,omitempty"`

	FatTissue                   string `xml:"FatTissue,omitempty"`

	MFA                         int32 `xml:"MFA,omitempty"`
}

type NotificationResult struct {
	XMLName        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 NotificationResult"`

	Result         *ProcessingResult `xml:"Result,omitempty"`

	NotificationID int32 `xml:"NotificationID,omitempty"`
}

type AnimalClassificationDataV2 struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 AnimalClassificationDataV2"`

	*AnimalClassificationData

	LValue  float64 `xml:"LValue,omitempty"`
}

type DataAccessResult struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 DataAccessResult"`

	Result  *ProcessingResult `xml:"Result,omitempty"`

	ID_PRG  int32 `xml:"ID_PRG,omitempty"`
}

type CattlePassportResult struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattlePassportResult"`

	Result          *ProcessingResult `xml:"Result,omitempty"`

	CattlePassports *CattlePassportArray `xml:"CattlePassports,omitempty"`
}

type CattlePassportArray struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattlePassportArray"`

	CattlePassportItem []*CattlePassport `xml:"CattlePassportItem,omitempty"`
}

type CattlePassport struct {
	XMLName           xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattlePassport"`

	EarTagNumber      string `xml:"EarTagNumber,omitempty"`

	PassportNumber    int32 `xml:"PassportNumber,omitempty"`

	PassportIssueDate time.Time `xml:"PassportIssueDate,omitempty"`

	Result            *ProcessingResult `xml:"Result,omitempty"`
}

type CattleOffspringResult struct {
	XMLName            xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleOffspringResult"`

	Result             *ProcessingResult `xml:"Result,omitempty"`

	EarTagNumberMother string `xml:"EarTagNumberMother,omitempty"`

	CattleOffsprings   *CattleOffspringDataArray `xml:"CattleOffsprings,omitempty"`
}

type CattleOffspringDataArray struct {
	XMLName                 xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleOffspringDataArray"`

	CattleOffspringDataItem []*CattleOffspringData `xml:"CattleOffspringDataItem,omitempty"`
}

type CattleOffspringData struct {
	XMLName      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 CattleOffspringData"`

	EarTagNumber string `xml:"EarTagNumber,omitempty"`

	BirthDate    time.Time `xml:"BirthDate,omitempty"`

	Gender       int32 `xml:"Gender,omitempty"`

	Name         string `xml:"Name,omitempty"`

	Race         int32 `xml:"Race,omitempty"`
}

type PoultryBarnInNotification struct {
	XMLName                   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PoultryBarnInNotification"`

	Amount                    int32 `xml:"Amount,omitempty"`

	EventDate                 time.Time `xml:"EventDate,omitempty"`

	SourceTvdNumber           int32 `xml:"SourceTvdNumber,omitempty"`

	PoultryUsageReason        *EnumPoultryUsageReason `xml:"PoultryUsageReason,omitempty"`

	AllowSeveralNotifications bool `xml:"AllowSeveralNotifications,omitempty"`
}

type SearchPoultryBarnInNotificationRequest struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SearchPoultryBarnInNotificationRequest"`

	*BaseRequest

	TVDNumber int32 `xml:"TVDNumber,omitempty"`

	FromDate  time.Time `xml:"FromDate,omitempty"`

	ToDate    time.Time `xml:"ToDate,omitempty"`
}

type GetPoultryBarnInNotificationResult struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetPoultryBarnInNotificationResult"`

	Result               *ProcessingResult `xml:"Result,omitempty"`

	ArrivalNotifications *SmallAnimalNotificationDataArrayOfTypeSearchPoultryBarnInNotification `xml:"ArrivalNotifications,omitempty"`
}

type SmallAnimalNotificationDataArrayOfTypeSearchPoultryBarnInNotification struct {
	XMLName                              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SmallAnimalNotificationDataArrayOfTypeSearchPoultryBarnInNotification"`

	SmallAnimalNotificationDataArrayItem []*SearchPoultryBarnInNotification `xml:"SmallAnimalNotificationDataArrayItem,omitempty"`
}

type SearchPoultryBarnInNotification struct {
	XMLName                    xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SearchPoultryBarnInNotification"`

	Amount                     int32 `xml:"Amount,omitempty"`

	EventDate                  time.Time `xml:"EventDate,omitempty"`

	SourceTvdNumber            int32 `xml:"SourceTvdNumber,omitempty"`

	PoultryUsageReason         *TranslatedEnumTypeOfEnumPoultryUsageReason `xml:"PoultryUsageReason,omitempty"`

	NotificationDate           time.Time `xml:"NotificationDate,omitempty"`

	HerdenIdentificationNumber string `xml:"HerdenIdentificationNumber,omitempty"`

	CreatedBy                  string `xml:"CreatedBy,omitempty"`

	DeletedBy                  string `xml:"DeletedBy,omitempty"`

	IsDeleted                  bool `xml:"IsDeleted,omitempty"`
}

type GetAnimalHusbandryDetailRequest struct {
	XMLName   xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryDetailRequest"`

	*BaseRequest

	TVDNumber int32 `xml:"TVDNumber,omitempty"`
}

type GetAnimalHusbandryDetailResult struct {
	XMLName                  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetAnimalHusbandryDetailResult"`

	Result                   *ProcessingResult `xml:"Result,omitempty"`

	PostData                 *HusbandryResult `xml:"PostData,omitempty"`

	IsActive                 bool `xml:"IsActive,omitempty"`

	MunicipalityName         string `xml:"MunicipalityName,omitempty"`

	CantonShortname          string `xml:"CantonShortname,omitempty"`

	CoordinateX              int32 `xml:"CoordinateX,omitempty"`

	CoordinateY              int32 `xml:"CoordinateY,omitempty"`

	Altitude                 int32 `xml:"Altitude,omitempty"`

	CantonAnimalHusbandryKey string `xml:"CantonAnimalHusbandryKey,omitempty"`

	CantonPersonKey          string `xml:"CantonPersonKey,omitempty"`

	BurNumber                string `xml:"BurNumber,omitempty"`

	AnimalHusbandryType      *TranslatedEnumTypeOfEnumAnimalHusbandryType `xml:"AnimalHusbandryType,omitempty"`

	MunicipalityNumber       int32 `xml:"MunicipalityNumber,omitempty"`

	TypeOfUse                *TranslatedEnumTypeOfEnumAnimalHusbandryTypeOfUse `xml:"TypeOfUse,omitempty"`

	IsMountain               bool `xml:"IsMountain,omitempty"`

	Zone                     *TranslatedEnumTypeOfEnumZone `xml:"Zone,omitempty"`

	Area                     *TranslatedEnumTypeOfEnumArea `xml:"Area,omitempty"`
}

type SearchSmallAnimalMovementRequest struct {
	XMLName  xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SearchSmallAnimalMovementRequest"`

	*BaseRequest

	FromDate time.Time `xml:"FromDate,omitempty"`

	ToDate   time.Time `xml:"ToDate,omitempty"`
}

type PigArrivalNotificationResult struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PigArrivalNotificationResult"`

	Result               *ProcessingResult `xml:"Result,omitempty"`

	ArrivalNotifications *SmallAnimalNotificationDataArrayOfTypePigArrivalNotification `xml:"ArrivalNotifications,omitempty"`
}

type SmallAnimalNotificationDataArrayOfTypePigArrivalNotification struct {
	XMLName                              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 SmallAnimalNotificationDataArrayOfTypePigArrivalNotification"`

	SmallAnimalNotificationDataArrayItem []*PigArrivalNotification `xml:"SmallAnimalNotificationDataArrayItem,omitempty"`
}

type PigArrivalNotification struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 PigArrivalNotification"`

	TVDNumber       int32 `xml:"TVDNumber,omitempty"`

	SourceTVDNumber int32 `xml:"SourceTVDNumber,omitempty"`

	Amount          int32 `xml:"Amount,omitempty"`

	DeletionDate    time.Time `xml:"DeletionDate,omitempty"`

	CreationDate    time.Time `xml:"CreationDate,omitempty"`

	EventDate       time.Time `xml:"EventDate,omitempty"`

	PigCategory     *TranslatedEnumTypeOfEnumPigCategory `xml:"PigCategory,omitempty"`
}

type GetMembershipForOrganisationRequest struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetMembershipForOrganisationRequest"`

	*BaseRequest

	TVDNumber   int32 `xml:"TVDNumber,omitempty"`

	AgateNumber string `xml:"AgateNumber,omitempty"`
}

type GetMembershipForOrganisationResult struct {
	XMLName                xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetMembershipForOrganisationResult"`

	Result                 *ProcessingResult `xml:"Result,omitempty"`

	MembershipOrganisation *GenusElements `xml:"MembershipOrganisation,omitempty"`

	BreedingOrganisation   *GenusElements `xml:"BreedingOrganisation,omitempty"`
}

type GenusElements struct {
	XMLName xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GenusElements"`

	Genus   []*TranslatedEnumTypeOfEnumGenus `xml:"Genus,omitempty"`
}

type WriteReplacementBatchEarTagOrderRequest struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 WriteReplacementBatchEarTagOrderRequest"`

	TVDNumber            int32 `xml:"TVDNumber,omitempty"`

	Genus                *EnumGenus `xml:"Genus,omitempty"`

	EartagOrderItemArray *EartagOrderItemArray `xml:"EartagOrderItemArray,omitempty"`
}

type EartagOrderItemArray struct {
	XMLName              xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EartagOrderItemArray"`

	EartagOrderArrayItem []*EartagOrderItem `xml:"EartagOrderArrayItem,omitempty"`
}

type EartagOrderItem struct {
	XMLName          xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 EartagOrderItem"`

	EarTagNumber     string `xml:"EarTagNumber,omitempty"`

	IsExpress        bool `xml:"IsExpress,omitempty"`

	IsLeftSideOrder  bool `xml:"IsLeftSideOrder,omitempty"`

	IsRightSideOrder bool `xml:"IsRightSideOrder,omitempty"`
}

type ReplacementEarTagOrdersResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ReplacementEarTagOrdersResult"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	ResultDetails *ArrayOfReplacementEarTagOrderDataItem `xml:"ResultDetails,omitempty"`
}

type ArrayOfReplacementEarTagOrderDataItem struct {
	XMLName                        xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfReplacementEarTagOrderDataItem"`

	ReplacementEarTagOrderDataItem []*ReplacementEarTagOrderDataItem `xml:"ReplacementEarTagOrderDataItem,omitempty"`
}

type ReplacementEarTagOrderDataItem struct {
	XMLName         xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ReplacementEarTagOrderDataItem"`

	NotificationID  int32 `xml:"NotificationID,omitempty"`

	CombiArticle    int32 `xml:"CombiArticle,omitempty"`

	OrderStatus     *TranslatedEnumTypeOfEnumOrderStatus `xml:"OrderStatus,omitempty"`

	OrderStatusDate time.Time `xml:"OrderStatusDate,omitempty"`

	EarTagNumber    string `xml:"EarTagNumber,omitempty"`

	IsExpress       bool `xml:"IsExpress,omitempty"`

	Text            string `xml:"Text,omitempty"`
}

type GetCattlesPerLeavingDateRequest struct {
	XMLName     xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDateRequest"`

	TVDNumber   int32 `xml:"TVDNumber,omitempty"`

	LeavingDate time.Time `xml:"LeavingDate,omitempty"`
}

type GetCattlesPerLeavingDateResult struct {
	XMLName       xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDateResult"`

	Result        *ProcessingResult `xml:"Result,omitempty"`

	ResultDetails *ArrayOfGetCattlesPerLeavingDateData `xml:"ResultDetails,omitempty"`
}

type ArrayOfGetCattlesPerLeavingDateData struct {
	XMLName                      xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 ArrayOfGetCattlesPerLeavingDateData"`

	GetCattlesPerLeavingDateData []*GetCattlesPerLeavingDateData `xml:"GetCattlesPerLeavingDateData,omitempty"`
}

type GetCattlesPerLeavingDateData struct {
	XMLName               xml.Name `xml:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1 GetCattlesPerLeavingDateData"`

	EarTagNumber          string `xml:"EarTagNumber,omitempty"`

	Name                  string `xml:"Name,omitempty"`

	BirthDate             time.Time `xml:"BirthDate,omitempty"`

	LeavingDate           time.Time `xml:"LeavingDate,omitempty"`

	Gender                *TranslatedEnumTypeOfEnumGender `xml:"Gender,omitempty"`

	Race                  *TranslatedEnumTypeOfEnumCattleRace `xml:"Race,omitempty"`

	TVDNumber             int32 `xml:"TVDNumber,omitempty"`

	HusbandryName         string `xml:"HusbandryName,omitempty"`

	StreetAndStreetNumber string `xml:"StreetAndStreetNumber,omitempty"`

	City                  string `xml:"City,omitempty"`

	PostCode              int32 `xml:"PostCode,omitempty"`
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

	EnumEquidBreedAchal_Tekkiner EnumEquidBreed = "Achal_Tekkiner"

	EnumEquidBreedAchal_Tekkiner_Partbred EnumEquidBreed = "Achal_Tekkiner_Partbred"

	EnumEquidBreedAegidienberger EnumEquidBreed = "Aegidienberger"

	EnumEquidBreedAmerican_Miniature_Horse EnumEquidBreed = "American_Miniature_Horse"

	EnumEquidBreedAmerican_Saddlebred EnumEquidBreed = "American_Saddlebred"

	EnumEquidBreedAndalusier EnumEquidBreed = "Andalusier"

	EnumEquidBreedAngloarab_Vollblut_Vorbuch EnumEquidBreed = "Angloarab_Vollblut_Vorbuch"

	EnumEquidBreedAngloaraber EnumEquidBreed = "Angloaraber"

	EnumEquidBreedAnglo_Araber_Vorbuch EnumEquidBreed = "Anglo_Araber_Vorbuch"

	EnumEquidBreedAngloarabisches_Halbblut EnumEquidBreed = "Angloarabisches_Halbblut"

	EnumEquidBreedAngloarabisches_Halbblut_Vorbuch EnumEquidBreed = "Angloarabisches_Halbblut_Vorbuch"

	EnumEquidBreedAngloarabisches_Vollblut EnumEquidBreed = "Angloarabisches_Vollblut"

	EnumEquidBreedAnglo_Normaenner EnumEquidBreed = "Anglo_Normaenner"

	EnumEquidBreedAppaloosa EnumEquidBreed = "Appaloosa"

	EnumEquidBreedAraber_Berber EnumEquidBreed = "Araber_Berber"

	EnumEquidBreedArabisches_Vollblut EnumEquidBreed = "Arabisches_Vollblut"

	EnumEquidBreedArabo_Friesen EnumEquidBreed = "Arabo_Friesen"

	EnumEquidBreedBaden_Wuerttemberg EnumEquidBreed = "Baden_Wuerttemberg"

	EnumEquidBreedBayerisches_Warmblut EnumEquidBreed = "Bayerisches_Warmblut"

	EnumEquidBreedBelgisches_Warmblut EnumEquidBreed = "Belgisches_Warmblut"

	EnumEquidBreedBerber EnumEquidBreed = "Berber"

	EnumEquidBreedBER_MA EnumEquidBreed = "BER_MA"

	EnumEquidBreedBosniaken EnumEquidBreed = "Bosniaken"

	EnumEquidBreedBrandenburger EnumEquidBreed = "Brandenburger"

	EnumEquidBreedBritish_Sporthorse EnumEquidBreed = "British_Sporthorse"

	EnumEquidBreedBritish_Spotted_Pony EnumEquidBreed = "British_Spotted_Pony"

	EnumEquidBreedCamargue EnumEquidBreed = "Camargue"

	EnumEquidBreedCH_Kleinpferd EnumEquidBreed = "CH_Kleinpferd"

	EnumEquidBreedCH_Sportpony EnumEquidBreed = "CH_Sportpony"

	EnumEquidBreedClydesdale EnumEquidBreed = "Clydesdale"

	EnumEquidBreedCob_Normand EnumEquidBreed = "Cob_Normand"

	EnumEquidBreedConnemara EnumEquidBreed = "Connemara"

	EnumEquidBreedCream_Color_Schweiz EnumEquidBreed = "Cream_Color_Schweiz"

	EnumEquidBreedCreme_Colors EnumEquidBreed = "Creme_Colors"

	EnumEquidBreedCria_Caballar EnumEquidBreed = "Cria_Caballar"

	EnumEquidBreedCriollo EnumEquidBreed = "Criollo"

	EnumEquidBreedDartmoor EnumEquidBreed = "Dartmoor"

	EnumEquidBreedDeutsches_Classic_Pony EnumEquidBreed = "Deutsches_Classic_Pony"

	EnumEquidBreedDeutsches_Reitpony EnumEquidBreed = "Deutsches_Reitpony"

	EnumEquidBreedDeutsches_Warmblut EnumEquidBreed = "Deutsches_Warmblut"

	EnumEquidBreedDuelmener EnumEquidBreed = "Duelmener"

	EnumEquidBreedEnglisches_Vollblut EnumEquidBreed = "Englisches_Vollblut"

	EnumEquidBreedAndere_Eselrasse EnumEquidBreed = "Andere_Eselrasse"

	EnumEquidBreedExmoor EnumEquidBreed = "Exmoor"

	EnumEquidBreedFell_Pony EnumEquidBreed = "Fell_Pony"

	EnumEquidBreedFinnisches_Warmblut EnumEquidBreed = "Finnisches_Warmblut"

	EnumEquidBreedFjord EnumEquidBreed = "Fjord"

	EnumEquidBreedFriese EnumEquidBreed = "Friese"

	EnumEquidBreedHaflinger_Mischrasse EnumEquidBreed = "Haflinger_Mischrasse"

	EnumEquidBreedHannoveraner EnumEquidBreed = "Hannoveraner"

	EnumEquidBreedHessen EnumEquidBreed = "Hessen"

	EnumEquidBreedHighland_Pony EnumEquidBreed = "Highland_Pony"

	EnumEquidBreedHighland_Pony_Carron EnumEquidBreed = "Highland_Pony_Carron"

	EnumEquidBreedHispano EnumEquidBreed = "Hispano"

	EnumEquidBreedHolsteiner EnumEquidBreed = "Holsteiner"

	EnumEquidBreedIrlaender EnumEquidBreed = "Irlaender"

	EnumEquidBreedIslandpferd EnumEquidBreed = "Islandpferd"

	EnumEquidBreedKladruber EnumEquidBreed = "Kladruber"

	EnumEquidBreedKnabstrupper EnumEquidBreed = "Knabstrupper"

	EnumEquidBreedKreuzung_ZVCH EnumEquidBreed = "Kreuzung_ZVCH"

	EnumEquidBreedLettisches_Warmblut EnumEquidBreed = "Lettisches_Warmblut"

	EnumEquidBreedLewitzer EnumEquidBreed = "Lewitzer"

	EnumEquidBreedLewitzschecke EnumEquidBreed = "Lewitzschecke"

	EnumEquidBreedLipizzaner EnumEquidBreed = "Lipizzaner"

	EnumEquidBreedLittauisches_Warmblut EnumEquidBreed = "Littauisches_Warmblut"

	EnumEquidBreedLusitano EnumEquidBreed = "Lusitano"

	EnumEquidBreedLuxemburgisches_Warmblut EnumEquidBreed = "Luxemburgisches_Warmblut"

	EnumEquidBreedMaulesel EnumEquidBreed = "Maulesel"

	EnumEquidBreedMaultier EnumEquidBreed = "Maultier"

	EnumEquidBreedMazedonier_ EnumEquidBreed = "Mazedonier_"

	EnumEquidBreedMazedonier_Partbred EnumEquidBreed = "Mazedonier_Partbred"

	EnumEquidBreedMecklenburger EnumEquidBreed = "Mecklenburger"

	EnumEquidBreedMerens EnumEquidBreed = "Merens"

	EnumEquidBreedMissouri_Foxtrotter EnumEquidBreed = "Missouri_Foxtrotter"

	EnumEquidBreedMorgan EnumEquidBreed = "Morgan"

	EnumEquidBreedNew_Forest EnumEquidBreed = "New_Forest"

	EnumEquidBreedNoriker EnumEquidBreed = "Noriker"

	EnumEquidBreedOesterr_Warmblut EnumEquidBreed = "Oesterr_Warmblut"

	EnumEquidBreedOldenburger_Pferd EnumEquidBreed = "Oldenburger_Pferd"

	EnumEquidBreedOrlow EnumEquidBreed = "Orlow"

	EnumEquidBreedOstfriese EnumEquidBreed = "Ostfriese"

	EnumEquidBreedPaint EnumEquidBreed = "Paint"

	EnumEquidBreedPaint_Horse EnumEquidBreed = "Paint_Horse"

	EnumEquidBreedPalomino EnumEquidBreed = "Palomino"

	EnumEquidBreedPartbred EnumEquidBreed = "Partbred"

	EnumEquidBreedPartbredaraber EnumEquidBreed = "Partbredaraber"

	EnumEquidBreedPaso EnumEquidBreed = "Paso"

	EnumEquidBreedPaso_Colombiano EnumEquidBreed = "Paso_Colombiano"

	EnumEquidBreedPaso_Fino EnumEquidBreed = "Paso_Fino"

	EnumEquidBreedPaso_Peruano EnumEquidBreed = "Paso_Peruano"

	EnumEquidBreedPercheron EnumEquidBreed = "Percheron"

	EnumEquidBreedPinto EnumEquidBreed = "Pinto"

	EnumEquidBreedPleven EnumEquidBreed = "Pleven"

	EnumEquidBreedPolen_Trakehner EnumEquidBreed = "Polen_Trakehner"

	EnumEquidBreedPottok EnumEquidBreed = "Pottok"

	EnumEquidBreedPrzewalski EnumEquidBreed = "Przewalski"

	EnumEquidBreedPura_raza_espanola EnumEquidBreed = "Pura_raza_espanola"

	EnumEquidBreedQuarter_Horse EnumEquidBreed = "Quarter_Horse"

	EnumEquidBreedRheinland EnumEquidBreed = "Rheinland"

	EnumEquidBreedRussisches_Warmblut EnumEquidBreed = "Russisches_Warmblut"

	EnumEquidBreedSachsen EnumEquidBreed = "Sachsen"

	EnumEquidBreedSachsen_Anhaltiner EnumEquidBreed = "Sachsen_Anhaltiner"

	EnumEquidBreedSelle_Francais EnumEquidBreed = "Selle_Francais"

	EnumEquidBreedShagya_Araber EnumEquidBreed = "Shagya_Araber"

	EnumEquidBreedShetlandpony EnumEquidBreed = "Shetlandpony"

	EnumEquidBreedShire_Horse EnumEquidBreed = "Shire_Horse"

	EnumEquidBreedSlowenisches_Warmblut EnumEquidBreed = "Slowenisches_Warmblut"

	EnumEquidBreedSpecial_Color_Schweiz EnumEquidBreed = "Special_Color_Schweiz"

	EnumEquidBreedSporthorse_Brasilien EnumEquidBreed = "Sporthorse_Brasilien"

	EnumEquidBreedSporthorse_Mexico EnumEquidBreed = "Sporthorse_Mexico"

	EnumEquidBreedTennessee_Walking_Horse EnumEquidBreed = "Tennessee_Walking_Horse"

	EnumEquidBreedThueringer EnumEquidBreed = "Thueringer"

	EnumEquidBreedTinker EnumEquidBreed = "Tinker"

	EnumEquidBreedTraber EnumEquidBreed = "Traber"

	EnumEquidBreedTrait_Comtois EnumEquidBreed = "Trait_Comtois"

	EnumEquidBreedTrakehner EnumEquidBreed = "Trakehner"

	EnumEquidBreedWarmblut_Rheinland_Pfalz_Saar EnumEquidBreed = "Warmblut_Rheinland_Pfalz_Saar"

	EnumEquidBreedWelsh EnumEquidBreed = "Welsh"

	EnumEquidBreedWelsh_Cob_Sektion_D_WD EnumEquidBreed = "Welsh_Cob_Sektion_D_WD"

	EnumEquidBreedWelsh_Mountain_Pony_WA EnumEquidBreed = "Welsh_Mountain_Pony_WA"

	EnumEquidBreedWelsh_Partbred_WK EnumEquidBreed = "Welsh_Partbred_WK"

	EnumEquidBreedWelsh_Pony_Cob_Typ_WC EnumEquidBreed = "Welsh_Pony_Cob_Typ_WC"

	EnumEquidBreedWelsh_Riding_Pony_WB EnumEquidBreed = "Welsh_Riding_Pony_WB"

	EnumEquidBreedWestfale EnumEquidBreed = "Westfale"

	EnumEquidBreedWuerttemberger EnumEquidBreed = "Wuerttemberger"

	EnumEquidBreedZangersheide EnumEquidBreed = "Zangersheide"

	EnumEquidBreedZweibruecken EnumEquidBreed = "Zweibruecken"

	EnumEquidBreedFreiberger EnumEquidBreed = "Freiberger"

	EnumEquidBreedHaflinger EnumEquidBreed = "Haflinger"

	EnumEquidBreedSchweizer_Warmblut EnumEquidBreed = "Schweizer_Warmblut"

	EnumEquidBreedAraber EnumEquidBreed = "Araber"

	EnumEquidBreedPony EnumEquidBreed = "Pony"

	EnumEquidBreedWarmblut EnumEquidBreed = "Warmblut"

	EnumEquidBreedVollblut EnumEquidBreed = "Vollblut"

	EnumEquidBreedKaltblut EnumEquidBreed = "Kaltblut"

	EnumEquidBreedKreuzung EnumEquidBreed = "Kreuzung"

	EnumEquidBreedAndere EnumEquidBreed = "Andere"

	EnumEquidBreedHollaendisches_Warmblut_KWPN EnumEquidBreed = "Hollaendisches_Warmblut_KWPN"

	EnumEquidBreedDaenisches_Warmblut EnumEquidBreed = "Daenisches_Warmblut"

	EnumEquidBreedPolnisches_Warmblut EnumEquidBreed = "Polnisches_Warmblut"

	EnumEquidBreedUngarisches_Warmblut EnumEquidBreed = "Ungarisches_Warmblut"

	EnumEquidBreedTschechisches_Warmblut EnumEquidBreed = "Tschechisches_Warmblut"

	EnumEquidBreedItalienisches_Warmblut EnumEquidBreed = "Italienisches_Warmblut"

	EnumEquidBreedCH_Sportpferd EnumEquidBreed = "CH_Sportpferd"

	EnumEquidBreedHalbblut EnumEquidBreed = "Halbblut"

	EnumEquidBreedSchwedisches_Halbblut EnumEquidBreed = "Schwedisches_Halbblut"

	EnumEquidBreedDaenisches_Reitpony EnumEquidBreed = "Daenisches_Reitpony"

	EnumEquidBreedBardigiano EnumEquidBreed = "Bardigiano"

	EnumEquidBreedKabardiner EnumEquidBreed = "Kabardiner"

	EnumEquidBreedArgentinisches_Polopony EnumEquidBreed = "Argentinisches_Polopony"

	EnumEquidBreedUngarisches_Vollblut EnumEquidBreed = "Ungarisches_Vollblut"

	EnumEquidBreedHalf_Saddlebred EnumEquidBreed = "Half_Saddlebred"

	EnumEquidBreedTigerscheck EnumEquidBreed = "Tigerscheck"

	EnumEquidBreedTigerscheck_Shetlandtyp EnumEquidBreed = "Tigerscheck_Shetlandtyp"

	EnumEquidBreedMini_Shetlandpony EnumEquidBreed = "Mini_Shetlandpony"

	EnumEquidBreedIrish_Cob EnumEquidBreed = "Irish_Cob"

	EnumEquidBreedCurly_Horse EnumEquidBreed = "Curly_Horse"

	EnumEquidBreedCruzado_Iberico EnumEquidBreed = "Cruzado_Iberico"

	EnumEquidBreedDales_Pony EnumEquidBreed = "Dales_Pony"

	EnumEquidBreedDeutsches_Partbred_Shetland_Pony EnumEquidBreed = "Deutsches_Partbred_Shetland_Pony"

	EnumEquidBreedDeutsches_Sportpferd EnumEquidBreed = "Deutsches_Sportpferd"

	EnumEquidBreedEnglisches_Reitpony EnumEquidBreed = "Englisches_Reitpony"

	EnumEquidBreedHannoverschers_Halbblut EnumEquidBreed = "Hannoverschers_Halbblut"

	EnumEquidBreedInternationales_Oldenburger_Springpferd EnumEquidBreed = "Internationales_Oldenburger_Springpferd"

	EnumEquidBreedKiger_Mustang EnumEquidBreed = "Kiger_Mustang"

	EnumEquidBreedKleines_Deutsches_Pony EnumEquidBreed = "Kleines_Deutsches_Pony"

	EnumEquidBreedKleines_Deutsches_Reitpferd EnumEquidBreed = "Kleines_Deutsches_Reitpferd"

	EnumEquidBreedLeonharder EnumEquidBreed = "Leonharder"

	EnumEquidBreedLeutstettener_Pferd EnumEquidBreed = "Leutstettener_Pferd"

	EnumEquidBreedMangalarga_Marchadores EnumEquidBreed = "Mangalarga_Marchadores"

	EnumEquidBreedNRW_Reitpferd EnumEquidBreed = "NRW_Reitpferd"

	EnumEquidBreedOstfriese_Alt_Oldenburger EnumEquidBreed = "Ostfriese_Alt_Oldenburger"

	EnumEquidBreedPaso_Iberoamericano EnumEquidBreed = "Paso_Iberoamericano"

	EnumEquidBreedPortugiesisches_Sport_pferd EnumEquidBreed = "Portugiesisches_Sport_pferd"

	EnumEquidBreedRheinisch_Deutsches_Kaltblut EnumEquidBreed = "Rheinisch_Deutsches_Kaltblut"

	EnumEquidBreedRottaler EnumEquidBreed = "Rottaler"

	EnumEquidBreedSchwarzwaelder_Kaltblut EnumEquidBreed = "Schwarzwaelder_Kaltblut"

	EnumEquidBreedSpanisches_Sportpferd EnumEquidBreed = "Spanisches_Sportpferd"

	EnumEquidBreedSueddeutsches_Kaltblut EnumEquidBreed = "Sueddeutsches_Kaltblut"

	EnumEquidBreedWarlander EnumEquidBreed = "Warlander"

	EnumEquidBreedPoitou_Esel EnumEquidBreed = "Poitou_Esel"

	EnumEquidBreedGrand_Noir_du_Berry EnumEquidBreed = "Grand_Noir_du_Berry"

	EnumEquidBreedAndalusischer_Riesenesel EnumEquidBreed = "Andalusischer_Riesenesel"

	EnumEquidBreedKatalanischer_Riesenesel EnumEquidBreed = "Katalanischer_Riesenesel"

	EnumEquidBreedAmerican_Miniatur_Esel EnumEquidBreed = "American_Miniatur_Esel"

	EnumEquidBreedRagusana EnumEquidBreed = "Ragusana"

	EnumEquidBreedAmiata_Esel EnumEquidBreed = "Amiata_Esel"

	EnumEquidBreedMartina_Franca_Esel EnumEquidBreed = "Martina_Franca_Esel"

	EnumEquidBreedContentin_Esel EnumEquidBreed = "Contentin_Esel"

	EnumEquidBreedNormandie_Esel EnumEquidBreed = "Normandie_Esel"

	EnumEquidBreedHausesel EnumEquidBreed = "Hausesel"

	EnumEquidBreedSchweizer_Zuchtpferd EnumEquidBreed = "Schweizer_Zuchtpferd"
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

type EnumEquidColor string

const (
	EnumEquidColorUnclassified EnumEquidColor = "Unclassified"

	EnumEquidColorRappe EnumEquidColor = "Rappe"

	EnumEquidColorBraun EnumEquidColor = "Braun"

	EnumEquidColorFuchs EnumEquidColor = "Fuchs"

	EnumEquidColorSchimmel EnumEquidColor = "Schimmel"

	EnumEquidColorSchecke EnumEquidColor = "Schecke"

	EnumEquidColorGrau EnumEquidColor = "Grau"

	EnumEquidColorOtherColor EnumEquidColor = "OtherColor"
)

type EnumEquidSpecies string

const (
	EnumEquidSpeciesUndefined EnumEquidSpecies = "Undefined"

	EnumEquidSpeciesPferd EnumEquidSpecies = "Pferd"

	EnumEquidSpeciesEsel EnumEquidSpecies = "Esel"

	EnumEquidSpeciesMaultier EnumEquidSpecies = "Maultier"

	EnumEquidSpeciesMaulesel EnumEquidSpecies = "Maulesel"
)

type EnumEquidPassIssuingState string

const (
	EnumEquidPassIssuingStateNotStarted EnumEquidPassIssuingState = "NotStarted"

	EnumEquidPassIssuingStateOrdered EnumEquidPassIssuingState = "Ordered"

	EnumEquidPassIssuingStateOrderedAndIssued EnumEquidPassIssuingState = "OrderedAndIssued"

	EnumEquidPassIssuingStateJustIssued EnumEquidPassIssuingState = "JustIssued"
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

type EnumEquidNotificationType string

const (
	EnumEquidNotificationTypeUndefined EnumEquidNotificationType = "Undefined"

	EnumEquidNotificationTypeBasicDataChange EnumEquidNotificationType = "BasicDataChange"

	EnumEquidNotificationTypeSlaughter EnumEquidNotificationType = "Slaughter"

	EnumEquidNotificationTypeFirstRegistration EnumEquidNotificationType = "FirstRegistration"

	EnumEquidNotificationTypeBirth EnumEquidNotificationType = "Birth"

	EnumEquidNotificationTypeImport EnumEquidNotificationType = "Import"

	EnumEquidNotificationTypeLegacyVerbalDescription EnumEquidNotificationType = "LegacyVerbalDescription"

	EnumEquidNotificationTypeGraphicalDescription EnumEquidNotificationType = "GraphicalDescription"

	EnumEquidNotificationTypeOwnershipCession_National EnumEquidNotificationType = "OwnershipCession_National"

	EnumEquidNotificationTypeOwnershipAcquisition EnumEquidNotificationType = "OwnershipAcquisition"

	EnumEquidNotificationTypeChipImplantation EnumEquidNotificationType = "ChipImplantation"

	EnumEquidNotificationTypePass EnumEquidNotificationType = "Pass"

	EnumEquidNotificationTypeLocationChange EnumEquidNotificationType = "LocationChange"

	EnumEquidNotificationTypeNewUeln EnumEquidNotificationType = "NewUeln"

	EnumEquidNotificationTypeTypeOfUsageChange EnumEquidNotificationType = "TypeOfUsageChange"

	EnumEquidNotificationTypeDataRetrieval EnumEquidNotificationType = "DataRetrieval"

	EnumEquidNotificationTypeEquidMembership EnumEquidNotificationType = "EquidMembership"

	EnumEquidNotificationTypeLocationChangePending EnumEquidNotificationType = "LocationChangePending"

	EnumEquidNotificationTypePassOrder EnumEquidNotificationType = "PassOrder"

	EnumEquidNotificationTypeDecease_Euthanize EnumEquidNotificationType = "Decease_Euthanize"

	EnumEquidNotificationTypeCastration EnumEquidNotificationType = "Castration"

	EnumEquidNotificationTypePassIssuerPermission EnumEquidNotificationType = "PassIssuerPermission"

	EnumEquidNotificationTypeWithers EnumEquidNotificationType = "Withers"

	EnumEquidNotificationTypeOwnershipCession_International EnumEquidNotificationType = "OwnershipCession_International"

	EnumEquidNotificationTypeImportAfterExport EnumEquidNotificationType = "ImportAfterExport"
)

type EnumEquidPassType string

const (
	EnumEquidPassTypeUnknown EnumEquidPassType = "Unknown"

	EnumEquidPassTypeFirstPass EnumEquidPassType = "FirstPass"

	EnumEquidPassTypeSubstitute EnumEquidPassType = "Substitute"

	EnumEquidPassTypeDuplicate EnumEquidPassType = "Duplicate"
)

type EnumEquidDeathNotificationType string

const (
	EnumEquidDeathNotificationTypeDecease EnumEquidDeathNotificationType = "Decease"

	EnumEquidDeathNotificationTypeEuthanize EnumEquidDeathNotificationType = "Euthanize"
)

type EnumEquidPassOrderType string

const (
	EnumEquidPassOrderTypeNot_Migratable EnumEquidPassOrderType = "Not_Migratable"

	EnumEquidPassOrderTypeHardcopy EnumEquidPassOrderType = "Hardcopy"

	EnumEquidPassOrderTypePDF EnumEquidPassOrderType = "PDF"
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

	EnumCattleRaceOriginal_Braunvieh EnumCattleRace = "Original_Braunvieh"

	EnumCattleRaceRed_Holstein EnumCattleRace = "Red_Holstein"

	EnumCattleRaceKreuzung EnumCattleRace = "Kreuzung"

	EnumCattleRaceHinterwaelder EnumCattleRace = "Hinterwaelder"

	EnumCattleRaceCharolais EnumCattleRace = "Charolais"

	EnumCattleRaceLimousin EnumCattleRace = "Limousin"

	EnumCattleRaceWeissblaue_Belgier EnumCattleRace = "Weissblaue_Belgier"

	EnumCattleRaceBlonde_d_Aquitaine EnumCattleRace = "Blonde_d_Aquitaine"

	EnumCattleRaceMaine_Anjou EnumCattleRace = "Maine_Anjou"

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

	EnumCattleRaceHighland_Cattle EnumCattleRace = "Highland_Cattle"

	EnumCattleRaceGalloway EnumCattleRace = "Galloway"

	EnumCattleRaceGuernsey EnumCattleRace = "Guernsey"

	EnumCattleRaceSwiss_Fleckvieh EnumCattleRace = "Swiss_Fleckvieh"

	EnumCattleRaceLuing EnumCattleRace = "Luing"

	EnumCattleRaceKiwicross EnumCattleRace = "Kiwicross"

	EnumCattleRaceNormande EnumCattleRace = "Normande"

	EnumCattleRaceAyrshire EnumCattleRace = "Ayrshire"

	EnumCattleRaceAbondance EnumCattleRace = "Abondance"

	EnumCattleRaceGrauvieh EnumCattleRace = "Grauvieh"

	EnumCattleRaceDexter EnumCattleRace = "Dexter"

	EnumCattleRaceBazadaise EnumCattleRace = "Bazadaise"

	EnumCattleRaceTuxer EnumCattleRace = "Tuxer"

	EnumCattleRaceMurnau_Werdenfelser EnumCattleRace = "Murnau_Werdenfelser"

	EnumCattleRaceDaenische_Rotbunte EnumCattleRace = "Daenische_Rotbunte"

	EnumCattleRaceSchwedische_Rotbunte EnumCattleRace = "Schwedische_Rotbunte"

	EnumCattleRaceNorwegische_Rotbunte EnumCattleRace = "Norwegische_Rotbunte"

	EnumCattleRacePinzgauer EnumCattleRace = "Pinzgauer"

	EnumCattleRaceSimmental EnumCattleRace = "Simmental"

	EnumCattleRaceDahomey EnumCattleRace = "Dahomey"

	EnumCattleRaceTarentaise EnumCattleRace = "Tarentaise"

	EnumCattleRaceVosgienne EnumCattleRace = "Vosgienne"

	EnumCattleRaceTexas_Longhorn EnumCattleRace = "Texas_Longhorn"

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

	EnumCattleRaceRotes_Hoehenvieh EnumCattleRace = "Rotes_Hoehenvieh"

	EnumCattleRaceValdostana EnumCattleRace = "Valdostana"

	EnumCattleRaceZwergzebu EnumCattleRace = "Zwergzebu"

	EnumCattleRaceOther EnumCattleRace = "Other"

	EnumCattleRaceBordelaise EnumCattleRace = "Bordelaise"

	EnumCattleRacePustertaler_Sprinzen EnumCattleRace = "Pustertaler_Sprinzen"

	EnumCattleRaceLowLine EnumCattleRace = "LowLine"

	EnumCattleRaceWelsh_Black EnumCattleRace = "Welsh_Black"
)

type TranslatedEnumTypeOfEnumEquidNotificationState struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidNotificationState"`

	EnumValue            *EnumEquidNotificationState `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidTypeOfUse struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidTypeOfUse"`

	EnumValue            *EnumEquidTypeOfUse `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumGender struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumGender"`

	EnumValue            *EnumGender `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidBreed struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidBreed"`

	EnumValue            *EnumEquidBreed `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidWithersClass struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidWithersClass"`

	EnumValue            *EnumEquidWithersClass `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidSpecies struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidSpecies"`

	EnumValue            *EnumEquidSpecies `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidPassIssuingState struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidPassIssuingState"`

	EnumValue            *EnumEquidPassIssuingState `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumAnimalHusbandryType struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumAnimalHusbandryType"`

	EnumValue            *EnumAnimalHusbandryType `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidNotificationType struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidNotificationType"`

	EnumValue            *EnumEquidNotificationType `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumEquidPassType struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumEquidPassType"`

	EnumValue            *EnumEquidPassType `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumPoultryUsageReason struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumPoultryUsageReason"`

	EnumValue            *EnumPoultryUsageReason `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumAnimalHusbandryTypeOfUse struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumAnimalHusbandryTypeOfUse"`

	EnumValue            *EnumAnimalHusbandryTypeOfUse `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumZone struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumZone"`

	EnumValue            *EnumZone `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumArea struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumArea"`

	EnumValue            *EnumArea `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumPigCategory struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumPigCategory"`

	EnumValue            *EnumPigCategory `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumGenus struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumGenus"`

	EnumValue            *EnumGenus `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumOrderStatus struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumOrderStatus"`

	EnumValue            *EnumOrderStatus `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type TranslatedEnumTypeOfEnumCattleRace struct {
	XMLName              xml.Name `xml:"http://www.identitas.ch/enumerations TranslatedEnumTypeOfEnumCattleRace"`

	EnumValue            *EnumCattleRace `xml:"EnumValue,omitempty"`

	RequestedTranslation string `xml:"RequestedTranslation,omitempty"`
}

type EnumEquidLocationChangeType string

const (
	EnumEquidLocationChangeTypeInSwitzerland EnumEquidLocationChangeType = "InSwitzerland"

	EnumEquidLocationChangeTypeOutOfSwitzerland EnumEquidLocationChangeType = "OutOfSwitzerland"

	EnumEquidLocationChangeTypeIntoSwitzerland EnumEquidLocationChangeType = "IntoSwitzerland"
)

type EnumRole string

const (
	EnumRoleSystem EnumRole = "System"

	EnumRoleAdministrator EnumRole = "Administrator"

	EnumRoleIdentitas EnumRole = "Identitas"

	EnumRoleFarmer EnumRole = "Farmer"

	EnumRoleOwner EnumRole = "Owner"

	EnumRoleIdentityVerifier EnumRole = "IdentityVerifier"

	EnumRolePassportIssuer EnumRole = "PassportIssuer"

	EnumRoleGovernment EnumRole = "Government"

	EnumRoleSlaughterHouse EnumRole = "SlaughterHouse"

	EnumRoleBreedingOrganisation EnumRole = "BreedingOrganisation"

	EnumRoleMemberOrganisation EnumRole = "MemberOrganisation"

	EnumRoleMandateTaker EnumRole = "MandateTaker"

	EnumRoleGuest EnumRole = "Guest"

	EnumRoleChipImplanter EnumRole = "ChipImplanter"

	EnumRoleAccountant EnumRole = "Accountant"

	EnumRoleCattlePassportIssuer EnumRole = "CattlePassportIssuer"

	EnumRoleScientificalDataRetriever EnumRole = "ScientificalDataRetriever"

	EnumRoleSlaughterInitiator EnumRole = "SlaughterInitiator"
)

type ArrayOfstring struct {
	XMLName xml.Name `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays ArrayOfstring"`

	String  []string `xml:"string,omitempty"`
}

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

func (service *AnimalTracingPortType) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *AnimalTracingPortType) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func (service *AnimalTracingPortType) Version(request *Version) (*VersionResponse, error) {
	response := new(VersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

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

func (service *AnimalTracingPortType) GetSmallAnimalSlaughterList(request *GetSmallAnimalSlaughterList) (*GetSmallAnimalSlaughterListResponse, error) {
	response := new(GetSmallAnimalSlaughterListResponse)
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
	err := service.client.Call("", request, response)
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

func (service *AnimalTracingPortType) WriteEquidNewChipNotification(request *WriteEquidNewChipNotification) (*WriteEquidNewChipNotificationResponse, error) {
	response := new(WriteEquidNewChipNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidWithersClassNotification(request *WriteEquidWithersClassNotification) (*WriteEquidWithersClassNotificationResponse, error) {
	response := new(WriteEquidWithersClassNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetEquidDetail(request *GetEquidDetail) (*GetEquidDetailResponse, error) {
	response := new(GetEquidDetailResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidCastrationNotification(request *WriteEquidCastrationNotification) (*WriteEquidCastrationNotificationResponse, error) {
	response := new(WriteEquidCastrationNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidSlaughterNotification(request *WriteEquidSlaughterNotification) (*WriteEquidSlaughterNotificationResponse, error) {
	response := new(WriteEquidSlaughterNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidDeathNotification(request *WriteEquidDeathNotification) (*WriteEquidDeathNotificationResponse, error) {
	response := new(WriteEquidDeathNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidImportNotification(request *WriteEquidImportNotification) (*WriteEquidImportNotificationResponse, error) {
	response := new(WriteEquidImportNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidBirthNotification(request *WriteEquidBirthNotification) (*WriteEquidBirthNotificationResponse, error) {
	response := new(WriteEquidBirthNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidPassIssuerPermissionNotification(request *WriteEquidPassIssuerPermissionNotification) (*WriteEquidPassIssuerPermissionNotificationResponse, error) {
	response := new(WriteEquidPassIssuerPermissionNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidOrderBasePassNotification(request *WriteEquidOrderBasePassNotification) (*WriteEquidOrderBasePassNotificationResponse, error) {
	response := new(WriteEquidOrderBasePassNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidPassIssuedNotification(request *WriteEquidPassIssuedNotification) (*WriteEquidPassIssuedNotificationResponse, error) {
	response := new(WriteEquidPassIssuedNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidMembershipOrganisationsNotification(request *WriteEquidMembershipOrganisationsNotification) (*WriteEquidMembershipOrganisationsNotificationResponse, error) {
	response := new(WriteEquidMembershipOrganisationsNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidTypeOfUseNotification(request *WriteEquidTypeOfUseNotification) (*WriteEquidTypeOfUseNotificationResponse, error) {
	response := new(WriteEquidTypeOfUseNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidBasicDataNotification(request *WriteEquidBasicDataNotification) (*WriteEquidBasicDataNotificationResponse, error) {
	response := new(WriteEquidBasicDataNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) SearchEquid(request *SearchEquid) (*SearchEquidResponse, error) {
	response := new(SearchEquidResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) WriteEquidImportAfterExportNotification(request *WriteEquidImportAfterExportNotification) (*WriteEquidImportAfterExportNotificationResponse, error) {
	response := new(WriteEquidImportAfterExportNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetEquidSlaughterList(request *GetEquidSlaughterList) (*GetEquidSlaughterListResponse, error) {
	response := new(GetEquidSlaughterListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AnimalTracingPortType) GetEquidsWithNotificationsOfMO(request *GetEquidsWithNotificationsOfMO) (*GetEquidsWithNotificationsOfMOResponse, error) {
	response := new(GetEquidsWithNotificationsOfMOResponse)
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
	err := service.client.Call("", request, response)
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
	err := service.client.Call("", request, response)
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

func (service *AnimalTracingPortType) WriteCattleImportAfterExportNotification(request *WriteCattleImportAfterExportNotification) (*WriteCattleImportAfterExportNotificationResponse, error) {
	response := new(WriteCattleImportAfterExportNotificationResponse)
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
	err := service.client.Call("", request, response)
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

func (service *AnimalTracingPortType) GetCattleHistoryV2(request *GetCattleHistoryV2) (*GetCattleHistoryV2Response, error) {
	response := new(GetCattleHistoryV2Response)
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
	err := service.client.Call("", request, response)
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
	err := service.client.Call("", request, response)
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
	err := service.client.Call("", request, response)
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
	XMLName   xml.Name `xml:"soap:Envelope"`
	XmlnsSoap string   `xml:"xmlns:soap,attr,omitempty"`
	XmlnsNs   string   `xml:"xmlns:ns,attr,omitempty"`
	Header    *SOAPHeader
	Body      SOAPBody
}

type SOAPHeader struct {
	XMLName  xml.Name `xml:"soap:Header"`
	Xmlnswsa string   `xml:"xmlns:wsa,attr,omitempty"`

	Action   string `xml:"wsa:Action,omitempty"`
	To       string `xml:"wsa:To,omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"soap:Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPResponseEnvelope struct {
	XMLName xml.Name  `xml:"Envelope"`
	Body    SOAPResponseBody
}

type SOAPResponseBody struct {
	XMLName xml.Name `xml:"Body"`
	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ soap:Fault"`

	Code    string `xml:"faultcode,omitempty"`
	String  string `xml:"faultstring,omitempty"`
	Actor   string `xml:"faultactor,omitempty"`
	Detail  string `xml:"detail,omitempty"`
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName        xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse      string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token          *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id        string `xml:"wsu:Id,attr,omitempty"`

	Username  *WSSUsername `xml:",omitempty"`
	Password  *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data      string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data      string `xml:",chardata"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tls     bool
	auth    *BasicAuth
	headers []interface{}
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1 << letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n - 1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: "UsernameToken-" + randStringBytesMaskImprSrc(9)}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
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

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	envelope.XmlnsSoap = "http://www.w3.org/2003/05/soap-envelope"
	envelope.XmlnsNs = "http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1"


	//if s.headers != nil && len(s.headers) > 0 {
	//	soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
	//	copy(soapHeader.Items, s.headers)
	//	envelope.Header = soapHeader
	//}


	envelope.Header = &SOAPHeader{
		Xmlnswsa:"http://www.w3.org/2005/08/addressing",
		Action: "http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1/AnimalTracingPortType/Version",
		To: "https://ws.wbf.admin.ch/Livestock/AnimalTracing/1",
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "gowsdl/0.1")
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

	log.Println(string(rawbody))
	respEnvelope := new(SOAPResponseEnvelope)
	respEnvelope.Body = SOAPResponseBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		beego.Debug("Error:", err)
		return fault
	}
	return nil
}
