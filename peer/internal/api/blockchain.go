package api

import (
	"encoding/json"
	"net/http"
)

type WalletPassphraseJSONRequest struct {
	Wallet_Address string  `json:"walletName"`
	Time           float64 `json:"timeUnlock"`
}

func walletPassphrase(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			// For JSON content type, decode the JSON into a struct
			var payload WalletPassphraseJSONRequest
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&payload); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				writeStatusUpdate(w, "Cannot marshal payload in Go object. Does the payload have the correct body structure?")
				return
			}

		default:
			w.WriteHeader(http.StatusBadRequest)
			writeStatusUpdate(w, "Request must have the content header set as application/json")
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Only POST requests will be handled.")
		return
	}
}

type SendToAddressRequest struct {
	Wallet_Address string  `json:"walletAddress"`
	Amount         float64 `json:"amount"`
}
type SendToAddressResponse struct {
	Hash string `json:"hash"`
}

func sendToAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			var payload SendToAddressRequest
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&payload); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				writeStatusUpdate(w, "Cannot marshal payload in Go object. Does the payload have the correct body structure?")
				return
			}
			// Check if wallet is unlocked

		default:
			w.WriteHeader(http.StatusBadRequest)
			writeStatusUpdate(w, "Request must have the content header set as application/json")
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Only POST requests will be handled.")
		return
	}
}

type GenerateJSONRequest struct {
	Blocks int `json:"blocks"`
}
type GenerateJSONResponse struct {
	Hash string `json:"hash"`
}

func generateCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			// For JSON content type, decode the JSON into a struct
			var payload GenerateJSONRequest
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&payload); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				writeStatusUpdate(w, "Cannot marshal payload in Go object. Does the payload have the correct body structure?")
				return
			}

		default:
			w.WriteHeader(http.StatusBadRequest)
			writeStatusUpdate(w, "Request must have the content header set as application/json")
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Only POST requests will be handled.")
		return
	}
}

type GetBalanceJSONResponse struct {
	Balance int `json:"balance"`
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//Need to execute command
		fileInfoResp := GetBalanceJSONResponse{
			Balance: 0,
		}
		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

// Revenue represents revenue data for a specific time period
type Revenue struct {
	Time   string `json:"time"`
	Income string `json:"income"`
	Expand string `json:"expand"`
}

// RevenueResponse represents the response from the /revenue endpoint
type RevenueResponse struct {
	ID       string    `json:"_id"`
	WalletID string    `json:"wallet_id"`
	Revenue  []Revenue `json:"revenue"`
}

func getRevenue(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy revenue data
		revenueData := []Revenue{
			{
				Time:   "2023-11",
				Income: "20",
				Expand: "23",
			},
		}

		// Create response object
		fileInfoResp := RevenueResponse{
			ID:       "65680d250505420b42427a82",
			WalletID: "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			Revenue:  revenueData,
		}
		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

// Transaction represents a single transaction
type Transaction struct {
	TransactionID string  `json:"transaction_id"`
	Reason        string  `json:"reason"`
	Status        string  `json:"status"`
	Time          string  `json:"time"`
	Amount        float64 `json:"amount"`
}

// Response represents the response from the API
type getTransactionResponse struct {
	ID           string        `json:"_id"`
	WalletID     string        `json:"wallet_id"`
	Transactions []Transaction `json:"transaction"`
}

func getLatestTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy transaction data
		transactions := []Transaction{
			{
				TransactionID: "59a53ee428a643e940546c5ccfc5663e",
				Reason:        "Dota2_OnePunchGodModeMenu.exe",
				Status:        "pending",
				Time:          "2023-11-30T04:20:20.244Z",
				Amount:        0.000012323,
			},
		}

		// Create response object
		fileInfoResp := getTransactionResponse{
			ID:           "65680d250505420b42427a82",
			WalletID:     "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			Transactions: transactions,
		}
		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

func getCompleteTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy transaction data
		transactions := []Transaction{
			{
				TransactionID: "59a53ee428a643e940546c5ccfc5663e",
				Reason:        "Dota2_OnePunchGodModeMenu.exe",
				Status:        "pending",
				Time:          "2023-11-30T04:20:20.244Z",
				Amount:        0.000012323,
			},
		}

		// Create response object
		fileInfoResp := getTransactionResponse{
			ID:           "65680d250505420b42427a82",
			WalletID:     "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			Transactions: transactions,
		}
		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

// TransferResponse represents the response from the /transfer endpoint
type TransferResponse struct {
	ID        string  `json:"_id"`
	WalletID  string  `json:"wallet_id"`
	Amount    float64 `json:"amount"`
	ReceiveID string  `json:"receive_id"`
	Reason    string  `json:"reason"`
}

func postTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy data
		fileInfoResp := TransferResponse{
			ID:        "65680d250505420b42427a82",
			WalletID:  "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			Amount:    20.0002,
			ReceiveID: "XvPyWFABdGQByyxvN8cWCgDpDX613hgriw",
			Reason:    "Transaction test",
		}

		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

// StatsResponse represents the response from the /transaction endpoint
type StatsResponse struct {
	ID            string `json:"_id"`
	PubKey        string `json:"pub_key"`
	UploadSpeed   string `json:"upload_speed"`
	DownloadSpeed string `json:"download_speed"`
}

func getSpeed(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy data
		fileInfoResp := StatsResponse{
			ID:            "65680d250505420b42427a82",
			PubKey:        "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			UploadSpeed:   "30",
			DownloadSpeed: "2000",
		}

		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

// TypesResponse represents the response from the /types endpoint
type TypesResponse struct {
	ID             string           `json:"_id"`
	PubKey         string           `json:"pub_key"`
	FileTypes      []string         `json:"filetype"`
	FileTypeNumber []map[string]int `json:"filetypeNumber"`
}

func getFileType(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy data
		fileInfoResp := TypesResponse{
			ID:        "65680d250505420b42427a82",
			PubKey:    "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			FileTypes: []string{"exe", "pdf", "mp4", "mp3", "jpg"},
			FileTypeNumber: []map[string]int{
				{"exe": 22, "pdf": 10, "mp4": 9, "mp3": 4, "jpg": 1},
			},
		}

		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

// Activity represents the activity data for a specific date
type ActivityBlock struct {
	Date     string `json:"date"`
	Download string `json:"download"`
	Upload   string `json:"upload"`
}

// ActivityResponse represents the response from the /activity endpoint
type ActivityResponse struct {
	ID         string          `json:"_id"`
	PubKey     string          `json:"pub_key"`
	Activities []ActivityBlock `json:"activities"`
}

func getActivity(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy data
		activities := []ActivityBlock{
			{
				Date:     "Mar30",
				Download: "10",
				Upload:   "2",
			},
		}

		// Create response object
		fileInfoResp := ActivityResponse{
			ID:         "65680d250505420b42427a82",
			PubKey:     "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			Activities: activities,
		}

		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

// Device represents a mining device
type Device struct {
	DeviceID      string `json:"device_id"`
	DeviceName    string `json:"device_name"`
	HashPower     string `json:"hash_power"`
	Status        string `json:"status"`
	Power         string `json:"power"`
	Profitability string `json:"profitability"`
}

// DeviceListResponse represents the response from the /device_list endpoint
type DeviceListResponse struct {
	ID      string   `json:"_id"`
	PubKey  string   `json:"pub_key"`
	Devices []Device `json:"devices"`
}

func getDeviceList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy data
		devices := []Device{
			{
				DeviceID:      "65680d250505420b42427a82",
				DeviceName:    "GeForce RTX 4090",
				HashPower:     "37.56",
				Status:        "Mining",
				Power:         "30",
				Profitability: "0.5523342",
			},
		}
		// Create response object
		fileInfoResp := DeviceListResponse{
			ID:      "65680d250505420b42427a82",
			PubKey:  "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			Devices: devices,
		}

		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

type PutDeviceResponse struct {
	ID            string `json:"_id"`
	PubKey        string `json:"pub_key"`
	DeviceID      string `json:"device_id"`
	DeviceName    string `json:"device_name"`
	HashPower     string `json:"hash_power"`
	Status        string `json:"status"`
	Power         string `json:"power"`
	Profitability string `json:"profitability"`
}

func putDevice(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy data
		device := Device{
			DeviceID:      "65680d250505420b42427a82",
			DeviceName:    "GeForce RTX 4090",
			HashPower:     "37.56",
			Status:        "Mining",
			Power:         "30",
			Profitability: "0.5523342",
		}

		// Create response object
		fileInfoResp := PutDeviceResponse{
			ID:            "65680d250505420b42427a82",
			PubKey:        "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			DeviceID:      device.DeviceID,
			DeviceName:    device.DeviceName,
			HashPower:     device.HashPower,
			Status:        "Mining",
			Power:         device.Power,
			Profitability: device.Profitability,
		}

		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}

type UnpaidBalanceResponse struct {
	WalletID      string `json:"wallet_id"`
	UnpaidBalance string `json:"unpaidBalance"`
}

func getUnpaidBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Dummy data
		unpaidBalance := "100.00"

		// Create response object
		fileInfoResp := UnpaidBalanceResponse{
			WalletID:      "13hgriwdGXvPyWFABDX6QByyxvN8cWCgDp",
			UnpaidBalance: unpaidBalance,
		}

		jsonData, err := json.Marshal(fileInfoResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeStatusUpdate(w, "Failed to convert JSON Data into a string")
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeStatusUpdate(w, "Request must have the content header set as application/json")
		return
	}
}
