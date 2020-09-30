package instance

import (
	"fmt"
	"github.com/IBM-Cloud/vcds-go-client/ibmvcdsession"
	"github.com/IBM-Cloud/vcds-go-client/vcds/client/users_and_accounts"
	"github.com/IBM-Cloud/vcds-go-client/vcds/models"
	"log"
	"time"
)

type IBMVCDCredentials struct {
	session *ibmvcdsession.IBMVCDSession
}

func NewIBMVCDCredentialClient(sess *ibmvcdsession.IBMVCDSession) *IBMVCDCredentials {
	return &IBMVCDCredentials{
		session: sess,
	}
}

//input transactionid,timeout
func (f *IBMVCDCredentials) GetV1Credentials(iamtoken, transactionid string, timeout time.Duration) (*models.AccountInfraCredentials, string, error) {
	log.Printf("calling the Get V1 Service Catalog ")
	params := users_and_accounts.NewControllersPublicV1ProxySetSlAccountKeyBssaccountParamsWithTimeout(timeout).WithAuthorization(iamtoken).WithXGlobalTransactionID(&transactionid)
	resp, err := f.session.VCDS.UsersAndAccounts.ControllersPublicV1ProxySetSlAccountKeyBssaccount(params)
	log.Printf("the Accountname  in the credentials  is %s", resp.Payload.AccountName)
	if err != nil {
		log.Printf("failed to get the credentials for this customer")
		return nil, transactionid, fmt.Errorf("failed to get the user and accounts  %v", err)
	}
	return resp.Payload, resp.XGlobalTransactionID, nil
}
