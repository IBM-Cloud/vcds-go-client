package instance

import (
	"fmt"
	"github.com/IBM-Cloud/vcds-go-client/ibmvcdsession"
	"github.com/IBM-Cloud/vcds-go-client/vcds/client/services"
	"github.com/IBM-Cloud/vcds-go-client/vcds/models"
	"log"
	"time"
)

type IBMVCDServices struct {
	session *ibmvcdsession.IBMVCDSession
}

func NewIBMVCDServicesClient(sess *ibmvcdsession.IBMVCDSession) *IBMVCDServices {
	return &IBMVCDServices{
		session: sess,
	}
}

//input transactionid,timeout
func (f *IBMVCDServices) GetV1ServiceCatalog(transactionid string, timeout time.Duration) (*models.ServiceCatalog, string, error) {
	log.Printf("calling the Get V1 Service Catalog ")
	params := services.NewControllersPublicV1ProxyGetServiceCatalogListParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid)
	resp, err := f.session.VCDS.Services.ControllersPublicV1ProxyGetServiceCatalogList(params)
	log.Printf("printing the service catalog elements count is %d ", resp.Payload.ServicesCatalog)

	if err != nil {
		log.Printf("failed to get the service catalog for this customer")
		return nil, transactionid, fmt.Errorf("failed to get the service catalog %v", err)
	}
	return resp.Payload, resp.XGlobalTransactionID, nil
}

//iamtoken,transactionid,serviceid

func (f *IBMVCDServices) GetV1ServiceInfo(iamtoken, transactionid, serviceid string, timeout time.Duration) (*models.ServiceCatalogDetails, string, error) {
	log.Printf("calling the service catalog details ")
	params := services.NewControllersPublicV1ProxyGetServiceCatalogDetailsParamsWithTimeout(timeout).WithXGlobalTransactionID(&transactionid).WithAuthorization(iamtoken).WithServiceID(serviceid)

	resp, err := f.session.VCDS.Services.ControllersPublicV1ProxyGetServiceCatalogDetails(params)

	if err != nil {
		log.Printf("failed to get service catalog details")
		return nil, transactionid, fmt.Errorf("failed to get the service catalog details %v", err)
	}
	return resp.Payload, resp.XGlobalTransactionID, nil
}
