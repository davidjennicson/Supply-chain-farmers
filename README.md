


# Blockchain App for Farmers, Companies, Crops & Bids

This is a basic smart contract (chaincode) for managing farmers, companies, crops, and bidding on crops using Hyperledger Fabric. It's written in Go and stores everything on the blockchain.

---

##  What can this chaincode do?

### 1.  Farmer
- Add a farmer
- Get a farmer's details
- Get all farmers
- Update a farmer
- Delete a farmer

### 2.  Company
- Add a company
- Get a company's details
- Get all companies
- Update a company
- Delete a company

### 3.  Crop
- Add a crop with name, price, ID, and last updated time
- Get a crop
- Get all crops
- Update or delete a crop

### 4.  Bids
- Create a bid for a crop
- Let companies place their bid on it
- View bid details
- Get all bids

Each bid keeps track of:
- Crop info
- Farmer’s Aadhar
- Base price
- Date and expiry
- List of bidders (company name, bid price, time)

---

##  Sample Commands (using peer CLI)

### Add a Farmer
```bash
peer chaincode invoke -C mychannel -n basic -c '{"function":"CreateFarmer","Args":["farmer1","Ravi Kumar","123456789012"]}'
```

### Add a Company
```bash
peer chaincode invoke -C mychannel -n basic -c '{"function":"CreateCompany","Args":["company1","AgroCorp","LIC1234"]}'
```

### Add a Crop
```bash
peer chaincode invoke -C mychannel -n basic -c '{"function":"CreateCrop","Args":["crop1","Wheat","2000","2025-04-08T10:00:00Z"]}'
```

### Create a Bid
```bash
peer chaincode invoke -C mychannel -n basic -c '{"function":"CreateBid","Args":["bid1","Wheat","crop1","123456789012","2000","2025-04-08","2025-04-12"]}'
```

### Place a Bid (Make Bid)
```bash
peer chaincode invoke -C mychannel -n basic -c '{"function":"MakeBid","Args":["bid1","AgroCorp","2100","2025-04-08T12:00:00Z"]}'
```

### Get a Bid
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"GetBid","Args":["bid1"]}'
```

---

## 📦 Files
- `smartcontract.go`: main chaincode file
- `README.md`: you're reading it 🙂

