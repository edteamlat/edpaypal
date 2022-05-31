package paypal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/google/uuid"

	"github.com/edteamlat/edpaypal/models"
)

const (
	ExpectedVerification = "SUCCESS"
	ExpectedStatus       = "completed"
)

const (
	EventTypeProduct      = "PAYMENT.CAPTURE.COMPLETED"
	EventTypeSubscription = "PAYMENT.SALE.COMPLETED"
)

type UseCase struct {
	useCaseOrder        Order
	useCaseSubscription Subscription
	useCaseInvoice      Invoice
}

func New(o Order, s Subscription, i Invoice) UseCase {
	return UseCase{
		useCaseOrder:        o,
		useCaseSubscription: s,
		useCaseInvoice:      i,
	}
}

func (uc UseCase) ProcessRequest(headers http.Header, body []byte) error {
	payPalRequest := parsePayPalRequest(headers, body)
	err := validate(&payPalRequest)
	if err != nil {
		return err
	}

	eventType, err := jsonparser.GetString(body, "event_type")
	if err != nil {
		return err
	}

	return uc.processPayment(eventType, &payPalRequest, body)
}

func (uc UseCase) processPayment(eventType string, request *models.PayPalRequest, body []byte) error {
	switch eventType {
	case EventTypeProduct:
		return uc.saleProduct(request, body)
	case EventTypeSubscription:
		return uc.saleSubscription(request, body)
	}

	log.Printf("the event type %q is not processed", eventType)

	return nil
}

func (uc UseCase) saleProduct(request *models.PayPalRequest, body []byte) error {
	var err error

	request.ID, err = jsonparser.GetString(body, "id")
	if err != nil {
		return err
	}

	request.ResourceID, err = jsonparser.GetString(body, "resource", "id")
	if err != nil {
		return err
	}

	request.Status, err = jsonparser.GetString(body, "resource", "status")
	if err != nil {
		return err
	}

	request.Custom, err = jsonparser.GetString(body, "resource", "custom_id")
	if err != nil {
		return err
	}

	request.Price, err = jsonparser.GetString(body, "resource", "amount", "value")
	if err != nil {
		return err
	}

	order, err := uc.useCaseOrder.ByID(uuid.MustParse(request.Custom))
	if err != nil {
		return err
	}

	value, err := strconv.ParseFloat(request.Price, 64)
	if err != nil {
		return err
	}

	if order.Price != value {
		return fmt.Errorf("el valor recibido: %0.2f, es diferente al valor esperado %0.2f", value, order.Price)
	}

	if !strings.EqualFold(request.Status, ExpectedStatus) {
		return fmt.Errorf("el estado de la transacción: %q no es el estado esperado: %q", request.Status, ExpectedStatus)
	}

	return uc.useCaseInvoice.Create(&order, uuid.Nil)
}

func (uc UseCase) saleSubscription(request *models.PayPalRequest, body []byte) error {
	var err error

	request.ID, err = jsonparser.GetString(body, "id")
	if err != nil {
		return err
	}

	request.ResourceID, err = jsonparser.GetString(body, "resource", "id")
	if err != nil {
		return err
	}

	request.Status, err = jsonparser.GetString(body, "resource", "state")
	if err != nil {
		return err
	}

	request.Custom, err = jsonparser.GetString(body, "resource", "custom")
	if err != nil {
		return err
	}

	request.Price, err = jsonparser.GetString(body, "resource", "amount", "total")
	if err != nil {
		return err
	}

	order, err := uc.useCaseOrder.ByID(uuid.MustParse(request.Custom))
	if err != nil {
		return err
	}

	value, err := strconv.ParseFloat(request.Price, 64)
	if err != nil {
		return err
	}

	if order.Price != value {
		return fmt.Errorf("el valor recibido: %0.2f, es diferente al valor esperado %0.2f", value, order.Price)
	}

	if !strings.EqualFold(request.Status, ExpectedStatus) {
		return fmt.Errorf("el estado de la transacción: %q no es el estado esperado: %q", request.Status, ExpectedStatus)
	}

	subscription := models.Subscription{
		Email:    order.Email,
		TypeSubs: order.TypeSubs,
	}

	err = uc.useCaseSubscription.Create(&subscription)
	if err != nil {
		return err
	}

	return uc.useCaseInvoice.Create(&order, subscription.ID)
}

func parsePayPalRequest(headers http.Header, body []byte) models.PayPalRequest {
	return models.PayPalRequest{
		AuthAlgo:         headers.Get("Paypal-Auth-Algo"),
		CertURL:          headers.Get("Paypal-Cert-Url"),
		TransmissionID:   headers.Get("Paypal-Transmission-Id"),
		TransmissionSig:  headers.Get("Paypal-Transmission-Sig"),
		TransmissionTime: headers.Get("Paypal-Transmission-Time"),
		WebhookEvent:     body,
		WebhookID:        os.Getenv("WEBHOOK_ID"),
	}
}

func validate(p *models.PayPalRequest) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, os.Getenv("VALIDATION_URL"), bytes.NewReader(data))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(os.Getenv("CLIENT_ID"), os.Getenv("SECRET_ID"))

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("PayPal response with status code %d, body: %s", response.StatusCode, string(body))
	}

	bodyMap := make(map[string]string)
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return err
	}

	if bodyMap["verification_status"] != ExpectedVerification {
		return fmt.Errorf("verification status is %s", bodyMap["verification_status"])
	}

	return nil
}
