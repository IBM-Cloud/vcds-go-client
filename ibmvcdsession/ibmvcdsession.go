package ibmvcdsession

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/IBM-Cloud/vcds-go-client/vcds/client"
	"github.com/IBM-Cloud/vcds-go-client/vcds/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	vcdsAPIEndpoint = "api.vmware-solutions.cloud.ibm.com"
)

type IBMVCDSession struct {
	IAMToken    string
	IMSToken    string
	VCDS        *client.IBMCloudForVmwareSolutions
	Timeout     time.Duration
	UserAccount string
	Region      string
	Zone        string
}

func powerJSONConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		/*t := reflect.TypeOf(data)
		if data != nil && t.Kind() == reflect.Ptr {
			v := reflect.Indirect(reflect.ValueOf(data))
			if t.Elem().Kind() == reflect.String {
				buf := new(bytes.Buffer)
				_, err := buf.ReadFrom(reader)
				if err != nil {
					return err
				}
				b := buf.Bytes()
				v.SetString(string(b))
				return nil
			}
		}*/
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(reader)
		if err != nil {
			return err
		}
		b := buf.Bytes()
		if b != nil {
			dec := json.NewDecoder(bytes.NewReader(b))
			dec.UseNumber() // preserve number formats
			err = dec.Decode(data)
		}
		if string(b) == "null" || err != nil {
			errorRecord, _ := data.(*models.Error)
			log.Printf("The errorrecord is %s ", errorRecord)
			return nil
		}
		return err
	})
}

func New(iamtoken, region string, debug bool, timeout time.Duration, useraccount string, zone string) (*IBMVCDSession, error) {
	session := &IBMVCDSession{
		IAMToken:    iamtoken,
		UserAccount: useraccount,
		Region:      region,
		Zone:        zone,
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}

	log.Printf("the apiendpoint url for power is %s", vcdsAPIEndpoint)
	transport := httptransport.New(vcdsAPIEndpoint, "/", []string{"https"})
	if debug {
		transport.Debug = debug
	}
	transport.Consumers[runtime.JSONMime] = powerJSONConsumer()
	session.VCDS = client.New(transport, nil)
	session.Timeout = timeout
	return session, nil
}
