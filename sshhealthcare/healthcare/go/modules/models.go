package modules

// Define the healthcare smart contract
type HealthcareContract struct {
}

// Define the patient structure
type Patient struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	BloodType   string `json:"bloodType"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	RecordIDs []string `json:"recordIDs"`
	Read      []string `json:"readAccess`
	Write     []string `json:"wrtieAccess`
}

// Define the doctor structure
type Doctor struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Specialty   string `json:"specialty"`
	PhoneNumber string `json:"phoneNumber"`
}


// Define the medical record structure
type MedicalRecord struct {
    ID               string `json:"id"`
    PatientID        string `json:"patientID"`
    DoctorID         string `json:"doctorID"`
    Date             string `json:"date"`
    Prescription     string `json:"prescription"`
}

// Define the lab technician structure
type LabTechnician struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
}

// Define the lab insuarance structure
type InsuranceCompany struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Company          string `json:"company"`
}