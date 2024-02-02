package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// SmartContract provides functions for managing a Farmer
type SmartContract struct {
	contractapi.Contract
}

// Farmer describes basic details of a farmer
type Farmer struct {
	Aadhar string `json:"aadhar"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	GST  string `json:"gst"`
}
type Crop struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	LastUpdate string `json:"last_update"` // ISO timestamp
}

type Bidder struct {
	CompanyName string `json:"company_name"`
	BidPrice    int    `json:"bid_price"`
	Timestamp   string `json:"timestamp"` // ISO timestamp
}

type Bid struct {
	CropID      string   `json:"crop_id"`
	CropName    string   `json:"crop_name"`
	FarmerAadhar string  `json:"farmer_aadhar"`
	BasePrice   int      `json:"base_price"`
	Date        string   `json:"date"`     // ISO timestamp
	Expiry      string   `json:"expiry"`   // ISO timestamp
	Bidders     []Bidder `json:"bidders"`
}


// InitLedger adds a base set of farmers to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	farmers := []Farmer{
		{ID: "farmer1", Name: "Ravi Kumar", Aadhar: "123456789012"},
		{ID: "farmer2", Name: "Sunita Sharma", Aadhar: "234567890123"},
		{ID: "farmer3", Name: "Amit Patel", Aadhar: "345678901234"},
		{ID: "farmer4", Name: "Geeta Verma", Aadhar: "456789012345"},
	}

	for _, farmer := range farmers {
		farmerJSON, err := json.Marshal(farmer)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(farmer.ID, farmerJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateFarmer adds a new farmer to the world state
func (s *SmartContract) CreateFarmer(ctx contractapi.TransactionContextInterface, id string, name string, aadhar string) error {
	exists, err := s.FarmerExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the farmer %s already exists", id)
	}

	farmer := Farmer{
		ID:     id,
		Name:   name,
		Aadhar: aadhar,
	}
	farmerJSON, err := json.Marshal(farmer)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, farmerJSON)
}

// ReadFarmer returns the farmer stored in the world state with given id
func (s *SmartContract) ReadFarmer(ctx contractapi.TransactionContextInterface, id string) (*Farmer, error) {
	farmerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if farmerJSON == nil {
		return nil, fmt.Errorf("the farmer %s does not exist", id)
	}

	var farmer Farmer
	err = json.Unmarshal(farmerJSON, &farmer)
	if err != nil {
		return nil, err
	}

	return &farmer, nil
}

// UpdateFarmer updates an existing farmer in the world state
func (s *SmartContract) UpdateFarmer(ctx contractapi.TransactionContextInterface, id string, name string, aadhar string) error {
	exists, err := s.FarmerExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the farmer %s does not exist", id)
	}

	farmer := Farmer{
		ID:     id,
		Name:   name,
		Aadhar: aadhar,
	}
	farmerJSON, err := json.Marshal(farmer)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, farmerJSON)
}

// DeleteFarmer deletes a farmer from the world state
func (s *SmartContract) DeleteFarmer(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.FarmerExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the farmer %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// FarmerExists checks if a farmer with given ID exists
func (s *SmartContract) FarmerExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	farmerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return farmerJSON != nil, nil
}

// GetAllFarmers returns all farmers in the world state
func (s *SmartContract) GetAllFarmers(ctx contractapi.TransactionContextInterface) ([]*Farmer, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var farmers []*Farmer
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var farmer Farmer
		err = json.Unmarshal(queryResponse.Value, &farmer)
		if err != nil {
			return nil, err
		}
		farmers = append(farmers, &farmer)
	}

	return farmers, nil
}

// CreateCompany adds a new company to the world state
func (s *SmartContract) CreateCompany(ctx contractapi.TransactionContextInterface, id string, name string, gst string) error {
	exists, err := s.CompanyExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the company %s already exists", id)
	}

	company := Company{
		ID:   id,
		Name: name,
		GST:  gst,
	}
	companyJSON, err := json.Marshal(company)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, companyJSON)
}

// ReadCompany returns a company from the world state
func (s *SmartContract) ReadCompany(ctx contractapi.TransactionContextInterface, id string) (*Company, error) {
	companyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if companyJSON == nil {
		return nil, fmt.Errorf("the company %s does not exist", id)
	}

	var company Company
	err = json.Unmarshal(companyJSON, &company)
	if err != nil {
		return nil, err
	}

	return &company, nil
}

// UpdateCompany updates an existing company
func (s *SmartContract) UpdateCompany(ctx contractapi.TransactionContextInterface, id string, name string, gst string) error {
	exists, err := s.CompanyExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the company %s does not exist", id)
	}

	company := Company{
		ID:   id,
		Name: name,
		GST:  gst,
	}
	companyJSON, err := json.Marshal(company)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, companyJSON)
}

// DeleteCompany deletes a company from the world state
func (s *SmartContract) DeleteCompany(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.CompanyExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the company %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// CompanyExists checks if a company exists in the world state
func (s *SmartContract) CompanyExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	companyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return companyJSON != nil, nil
}

// GetAllCompanies returns all companies in the world state
func (s *SmartContract) GetAllCompanies(ctx contractapi.TransactionContextInterface) ([]*Company, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var companies []*Company
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var company Company
		err = json.Unmarshal(queryResponse.Value, &company)
		if err != nil {
			continue // skip if it's not a Company
		}

		// Optional: You could add a prefix to distinguish keys like "company1", "farmer1"
		// and filter here if needed
		companies = append(companies, &company)
	}

	return companies, nil
}


func (s *SmartContract) CreateCrop(ctx contractapi.TransactionContextInterface, id string, name string, price int, lastUpdate string) error {
	crop := Crop{
		ID:         id,
		Name:       name,
		Price:      price,
		LastUpdate: lastUpdate,
	}
	cropJSON, err := json.Marshal(crop)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, cropJSON)
}

func (s *SmartContract) GetCrop(ctx contractapi.TransactionContextInterface, id string) (*Crop, error) {
	cropJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}
	if cropJSON == nil {
		return nil, fmt.Errorf("crop %s not found", id)
	}
	var crop Crop
	err = json.Unmarshal(cropJSON, &crop)
	if err != nil {
		return nil, err
	}
	return &crop, nil
}

func (s *SmartContract) GetAllCrops(ctx contractapi.TransactionContextInterface) ([]*Crop, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var crops []*Crop
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var crop Crop
		if err := json.Unmarshal(queryResponse.Value, &crop); err == nil && crop.Name != "" {
			crops = append(crops, &crop)
		}
	}
	return crops, nil
}

// --- Bid Functions ---

func (s *SmartContract) CreateBid(ctx contractapi.TransactionContextInterface, cropId string, cropName string, farmerAadhar string, basePrice int, date string, expiry string) error {
	bid := Bid{
		CropID:       cropId,
		CropName:     cropName,
		FarmerAadhar: farmerAadhar,
		BasePrice:    basePrice,
		Date:         date,
		Expiry:       expiry,
		Bidders:      []Bidder{},
	}
	bidJSON, err := json.Marshal(bid)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState("BID_"+cropId, bidJSON)
}

func (s *SmartContract) GetBid(ctx contractapi.TransactionContextInterface, cropId string) (*Bid, error) {
	bidJSON, err := ctx.GetStub().GetState("BID_" + cropId)
	if err != nil {
		return nil, err
	}
	if bidJSON == nil {
		return nil, fmt.Errorf("bid for crop %s not found", cropId)
	}
	var bid Bid
	err = json.Unmarshal(bidJSON, &bid)
	if err != nil {
		return nil, err
	}
	return &bid, nil
}

func (s *SmartContract) GetAllBids(ctx contractapi.TransactionContextInterface) ([]*Bid, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var bids []*Bid
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var bid Bid
		if err := json.Unmarshal(queryResponse.Value, &bid); err == nil && bid.CropID != "" {
			bids = append(bids, &bid)
		}
	}
	return bids, nil
}

func (s *SmartContract) MakeBid(ctx contractapi.TransactionContextInterface, cropId string, companyName string, bidPrice int, timestamp string) error {
	bid, err := s.GetBid(ctx, cropId)
	if err != nil {
		return err
	}

	newBidder := Bidder{
		CompanyName: companyName,
		BidPrice:    bidPrice,
		Timestamp:   timestamp,
	}
	bid.Bidders = append(bid.Bidders, newBidder)

	bidJSON, err := json.Marshal(bid)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState("BID_"+cropId, bidJSON)
}