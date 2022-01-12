package maas

type HardwareInfo struct {
	SystemVendor             string `json:"system_vendor"`
	SystemProduct            string `json:"system_product"`
	SystemFamily             string `json:"system_family"`
	SystemVersion            string `json:"system_version"`
	SystemSku                string `json:"system_sku"`
	SystemSerial             string `json:"system_serial"`
	CpuModel                 string `json:"cpu_model"`
	MainboardVendor          string `json:"mainboard_vendor"`
	MainboardProduct         string `json:"mainboard_product"`
	MainboardSerial          string `json:"mainboard_serial"`
	MainboardVersion         string `json:"mainboard_version"`
	MainboardFirmwareVendor  string `json:"mainboard_firmware_vendor"`
	MainboardFirmwareDate    string `json:"mainboard_firmware_date"`
	MainboardFirmwareVersion string `json:"mainboard_firmware_version"`
	ChassisVendor            string `json:"chassis_vendor"`
	ChassisType              string `json:"chassis_type"`
	ChassisSerial            string `json:"chassis_serial"`
	ChassisVersion           string `json:"chassis_version"`
}
