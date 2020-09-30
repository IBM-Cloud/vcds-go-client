package instance

import (
	"fmt"
	"github.com/IBM-Cloud/vcds-go-client/ibmvcdsession"
	"github.com/IBM-Cloud/vcds-go-client/vcds/client/utils"
	"github.com/IBM-Cloud/vcds-go-client/vcds/models"
	"log"
	"time"
)

type IBMVCDUtils struct {
	session *ibmvcdsession.IBMVCDSession
}

func NewIBMVCDUtilsClient(sess *ibmvcdsession.IBMVCDSession) *IBMVCDUtils {
	return &IBMVCDUtils{
		session: sess,
	}
}

//input iamtoken,datacenter,transactionid,timeout
func (f *IBMVCDUtils) GetV1ExistingNetwork(iamtoken, dc, transactionid string, timeout time.Duration) (models.ArrayExistingVlans, string, error) {
	log.Printf("calling the getv1network ")
	params := utils.NewControllersPublicV1ProxyListExistingNetworkDcParamsWithTimeout(timeout).WithAuthorization(iamtoken).WithDatacenter(dc).WithXGlobalTransactionID(&transactionid)

	resp, err := f.session.VCDS.Utils.ControllersPublicV1ProxyListExistingNetworkDc(params)

	if err != nil {
		log.Printf("failed to get the dc for this customer")
		return nil, transactionid, fmt.Errorf("failed to get the dc")
	}
	return resp.Payload, resp.XGlobalTransactionID, nil
}

//Get the list of supported disk types as a JSON object, and the key name is "disk_types".

func (f *IBMVCDUtils) GetV1ExistingDisks(transactionid string, timeout time.Duration) (*models.DefDiskTypes, string, error) {
	log.Printf("calling the getv1existingdisks ")
	params := utils.NewControllersPublicV2ProxyListDiskTypesParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid)

	resp, err := f.session.VCDS.Utils.ControllersPublicV2ProxyListDiskTypes(params)

	if err != nil {
		log.Printf("failed to get the dc for this customer")
		return nil, transactionid, fmt.Errorf("failed to get the dc")
	}
	return resp.Payload, resp.XGlobalTransactionID, nil
}

//Get the list of supported data center locations as a JSON object, and the key name is "locations".

func (f *IBMVCDUtils) GetV1ExistingLocations(transactionid string, timeout time.Duration) (*models.DefLocations, string, error) {
	log.Printf("calling the getv1existinglocations ")
	params := utils.NewControllersPublicV2ProxyListLocationsParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid)

	resp, err := f.session.VCDS.Utils.ControllersPublicV2ProxyListLocations(params)

	if err != nil {
		log.Printf("failed to get the locations for this customer")
		return nil, transactionid, fmt.Errorf("failed to get the locations %v", err)
	}
	return resp.Payload, resp.XGlobalTransactionID, nil
}

// Get the list of supported data center locations as a JSON object, and the key name is "mzr_locations".

func (f *IBMVCDUtils) GetV1ExistingMzrLocations(transactionid string, timeout time.Duration) (*models.DefMzrLocations, error) {
	log.Printf("calling the getv1existingmzrlocations ")
	params := utils.NewControllersPublicV2ProxyListMzrLocationsParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid)

	resp, err := f.session.VCDS.Utils.ControllersPublicV2ProxyListMzrLocations(params)

	if err != nil {
		log.Printf("failed to get the mzr locations for this customer")
		return nil, fmt.Errorf("failed to get the mzr locations %v", err)
	}
	return resp.Payload, nil
}

// Get the list of supported RAM types as a JSON object, and the key name is "ram_types".

func (f *IBMVCDUtils) GetV1ExistingRam(transactionid string, timeout time.Duration) (*models.DefRAMTypes, error) {
	log.Printf("calling the getv1existingmzrlocations ")
	params := utils.NewControllersPublicV2ProxyListRAMTypesParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid)

	resp, err := f.session.VCDS.Utils.ControllersPublicV2ProxyListRAMTypes(params)

	if err != nil {
		log.Printf("failed to get the ram types  for this customer")
		return nil, fmt.Errorf("failed to get the ram types %v", err)
	}
	return resp.Payload, nil
}

// Get the list of supported server types per location as a JSON object, and the key name is the given location value.

func (f *IBMVCDUtils) GetV1ExistingServerTypes(transactionid string, timeout time.Duration) (*models.DefServerTypes, error) {
	log.Printf("calling the getv1existingmzrlocations ")
	params := utils.NewControllersPublicV2ProxyListServerTypesParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid)

	resp, err := f.session.VCDS.Utils.ControllersPublicV2ProxyListServerTypes(params)

	if err != nil {
		log.Printf("failed to get the server types  for this customer")
		return nil, fmt.Errorf("failed to get the server types %v", err)
	}
	return resp.Payload, nil
}

//Get the list of supported shared storage tiers as a JSON object, and the key name is "shared_storage_tiers".

func (f *IBMVCDUtils) GetV1ExistingSharedStorageTypes(transactionid string, timeout time.Duration) (*models.DefSharedStorageTiers, error) {
	log.Printf("calling the getv1sharedstoragetypes ")
	params := utils.NewControllersPublicV2ProxyListSharedStorageTiersParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid)

	resp, err := f.session.VCDS.Utils.ControllersPublicV2ProxyListSharedStorageTiers(params)

	if err != nil {
		log.Printf("failed to get the shared storage types for this customer")
		return nil, fmt.Errorf("failed to get the shared storage types %v", err)
	}
	return resp.Payload, nil
}

// Get the list of available hardware templates as a JSON object, and the key name is "templates".

func (f *IBMVCDUtils) GetV1ExistingTemplates(transactionid string, timeout time.Duration) (*models.DefTemplates, error) {
	log.Printf("calling the getv1existingtemplates ")
	params := utils.NewControllersPublicV2ProxyListTemplatesParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid)

	resp, err := f.session.VCDS.Utils.ControllersPublicV2ProxyListTemplates(params)

	if err != nil {
		log.Printf("failed to get the templates for this customer")
		return nil, fmt.Errorf("failed to get the templates %v", err)
	}
	return resp.Payload, nil
}

//Get the list of supported vSphere versions as a JSON object, and the key name is "vsphere_versions".

func (f *IBMVCDUtils) GetV1ExistingVsphereVersions(transactionid string, timeout time.Duration) (*models.DefVsphereVersions, error) {
	log.Printf("calling the getv1existingvsphereversions ")
	params := utils.NewControllersPublicV2ProxyListVsphereVersionsParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid)

	resp, err := f.session.VCDS.Utils.ControllersPublicV2ProxyListVsphereVersions(params)

	if err != nil {
		log.Printf("failed to get the vsphere versions for this customer")
		return nil, fmt.Errorf("failed to get the vsphere versions %v", err)
	}
	return resp.Payload, nil
}
