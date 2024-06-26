package models

type Loan struct {
	// LoanTypeCode      string  `json:"Loan_Type_Code" validate:"max=50"` : Cancel on R3
	LoanAmount    float64 `json:"Loan_Amount" validate:"omitempty,number"`
	LoanPurpose   string  `json:"Loan_Purpose" validate:"max=1024"`
	ProductCode   string  `json:"Product_Code" validate:"max=50"`
	MarketingCode string  `json:"Marketing_Code" validate:"max=50"`
	// InstallmentPeriod int     `json:"Installment_Period" validate:"omitempty,numeric"` : Cancel on R3
	InstallmentAmount float64 `json:"Installment_Amount" validate:"omitempty,number"`
	/* R3 */
	ProductType             string  `json:"Product_Type" validate:"max=50"`
	ProductSubType          string  `json:"Product_Sub_Type" validate:"max=50"`
	ExistingCreditLimit     float64 `json:"Existing_Credit_Limit" validate:"omitempty,number"`
	ChangingCreditLimit     float64 `json:"Changing_Credit_Limit" validate:"omitempty,number"`
	ArrangementPurpose_Code string  `json:"Arrangement_Purpose_Code" validate:"max=3"`
	LoanAccountNumber       string  `json:"Loan_Account_Number" validate:"max=20"`
	Refinance               bool    `json:"Refinance"`
}

type Applicant struct {
	ApplicantType   string `json:"Applicant_Type" validate:"omitempty,oneof=B C G S"`
	ThaiName        string `json:"Thai_Name" validate:"max=255"`
	ThaiSurname     string `json:"Thai_Surname" validate:"max=255"`
	CitizenIdNumber string `json:"Citizen_Id_Number" validate:"max=20"`
	ThaiMiddleName  string `json:"Thai_Middle_Name" validate:"max=255"`
	ThaiMaidenName  string `json:"Thai_Maiden_Name" validate:"max=255"`
	EngName         string `json:"Eng_Name" validate:"max=255"`
	EngSurname      string `json:"Eng_Surname" validate:"max=255"`
	CIFNumber       string `json:"CIF_Number" validate:"max=10"`
	Sex             string `json:"Sex" validate:"omitempty,oneof=M F"`
	DateOfBirth     string `json:"Date_Of_Birth,omitempty" validate:"dateFormatYYYYMMDD,omitempty"`
	Age             int    `json:"Age"`
	Nationality     string `json:"Nationality" validate:"max=2"`
	MaritalStatus   string `json:"Marital_Status" validate:"max=2"`
	EducationLevel  string `json:"Education_Level" validate:"max=1"`
	// CitizenIdExp             string `json:"Citizen_Id_Exp,omitempty" validate:"dateFormatYYYYMMDD,omitempty"` : Cancel on R3
	PassportID     string `json:"Passport_ID" validate:"max=20"`
	CountryOfIssue string `json:"Country_Of_Issue" validate:"max=2"`
	// ExpiryDate               string `json:"Expiry_Date,omitempty" validate:"dateFormatYYYYMMDD",omitempty` : Cancel on R3
	// TaxIdCardNumber          string `json:"Tax_Id_Card_Number" validate:"max=20"` : Cancel on R3
	MobilePhone string `json:"Mobile_Phone" validate:"max=30"`
	Email       string `json:"Email" validate:"max=50"`
	// ResidentType             string `json:"Resident_Type" validate:"max=1"` : Cancel on R3
	RelationShipWithBorrower string `json:"RelationShip_with_Borrower" validate:"max=1"`
	KYCCDDStatus             string `json:"KYC_CDD_Status" validate:"max=1"`
	// KTBGrade                 string `json:"KTB_Grade" validate:"max=10"` : Cancel on R3
	// LivingPlacePeriodMonths  string `json:"Living_Place_Period_Months" validate:"max=3"` : Cancel on R3
	// Staff                    string `json:"Staff" validate:"max=255"`
	TUEFID string `json:"Tuef_Id" validate:"max=100"`
	// NcbConsentType           string `json:"Ncb_Consent_Type" validate:"max=255"` : Cancel on R3
	// NcbConsentAcceptanceCode string `json:"Ncb_Consent_Acceptance_Code" validate:"max=255"` : Cancel on R3
}

type Business struct {
	ApplicantType   string `json:"Applicant_Type" validate:"omitempty,oneof=B C G S"`
	ThaiName        string `json:"Thai_Name" validate:"max=255"`
	ThaiSurname     string `json:"Thai_Surname" validate:"max=255"`
	CitizenIdNumber string `json:"Citizen_Id_Number" validate:"max=20"`
	BusinessType    string `json:"Business_Type" validate:"max=5"`
	// RegistrationType     string  `json:"Registration_Type" validate:"max=2"` : Cancel on R3
	TaxIDNumber      string `json:"Tax_ID_Number" validate:"max=20"`
	RegistrationDate string `json:"Registration_Date,omitempty" validate:"omitempty,dateFormatYYYYMMDD"`
	// RegisteredCapital    float64 `json:"Registered_Capital,omitempty" validate:"omitempty,number"` : Cancel on R3
	DurationOfBusiness int    `json:"Duration_Of_Business" validate:"omitempty,numeric"`
	BusinessName       string `json:"Business_Name" validate:"max=255"`
	BusinessNameEng    string `json:"Business_Name_Eng" validate:"max=255"`
	// SupervisingOfficer   string  `json:"Supervising_Officer" validate:"max=255"` : Cancel on R3
	// TypeOfBusiness       string  `json:"Type_Of_Business" validate:"max=2"` : Cancel on R3
	// OwnerBranch          string  `json:"Owner_Branch" validate:"max=6"` : Cancel on R3
	AnnualSales          float64 `json:"Annual_Sales" validate:"omitempty,number"`
	WebsiteName          string  `json:"Website_Name" validate:"max=255"`
	BusinessPurposeCode1 string  `json:"Business_Purpose_Code_1" validate:"max=3"`
	//R3 Business_Purpose_Code_3
	KTBCustomerCode   string `json:"KTB_Customer_Code" validate:"max=10"`
	InvolvePartyGroup string `json:"Involve_Party_Group" validate:"max=50"`
	// DebtLevel                 string  `json:"Debt_Level" validate:"max=50"` : Cancel on R3
	KYCCDDStatus        string `json:"KYC_CDD_Status" validate:"max=50"`
	NetFixedAsset       string `json:"Net_Fixed_Asset" validate:"max=20"`
	TotalNoOfEmployees  string `json:"Total_No_Of_Employees" validate:"max=8"`
	RegistrationCountry string `json:"Registration_Country" validate:"max=2"`
	BusinessStatus      string `json:"Business_Status" validate:"max=50"`
	// YearBecomingLoanCustomer  string  `json:"Year_Becoming_Loan_Customer" validate:"max=2"` : Cancel on R3
	// YearBecomingOtherCustomer string  `json:"Year_Becoming_Other_Customer" validate:"max=2"` : Cancel on R3
	// CustomerGroupName         string  `json:"Customer_Group_Name" validate:"max=50"` : Cancel on R3
	// TotalDebit                float64 `json:"Total_Debit" validate:"number,omitempty"` : Cancel on R3
	// TotalCredit               float64 `json:"Total_Credit" validate:"number,omitempty"` : Cancel on R3
	// AverageDebit              float64 `json:"Average_Debit" validate:"number,omitempty"` : Cancel on R3
	// AverageCredit             float64 `json:"Average_Credit" validate:"number,omitempty"` : Cancel on R3
}

type Collateral struct {
	ApplicantType       string  `json:"Applicant_Type" validate:"omitempty,oneof=B C G S"`
	ThaiName            string  `json:"Thai_Name" validate:"max=255"`
	ThaiSurname         string  `json:"Thai_Surname" validate:"max=255"`
	CitizenIdNumber     string  `json:"Citizen_Id_Number" validate:"max=20"`
	CollateralNumber    string  `json:"Collateral_Number" validate:"max=20"`
	TypeOfCollateral    string  `json:"Type_Of_Collateral" validate:"max=2"`
	CollateralCondition string  `json:"Collateral_Condition" validate:"max=3"`
	SubCollateralType   string  `json:"Sub_Collateral_Type" validate:"max=2"`
	AmountOfMortgage    float64 `json:"Amount_Of_Mortgage" validate:"number,omitempty"`
	AppraisalPrice      float64 `json:"Appraisal_Price" validate:"number,omitempty"`
	// AppraisalBy               string  `json:"Appraisal_By" validate:"max=1"` : Canceled on R3
	AppraisalAssessmentAgency string `json:"Appraisal_Assessment_Agency" validate:"max=8"`
	// ProjectCode               string  `json:"Project_Code" validate:"max=11"` : Canceled on R3
	// ThaiProjectName           string  `json:"Thai_Project_Name" validate:"max=255"` : Canceled on R3
	// AssessmentNumber          string  `json:"Assessment_Number" validate:"max=18"` : Canceled on R3
	DateOfLastAppraisal string `json:"Date_Of_Last_Appraisal,omitempty" validate:"dateFormatYYYYMMDD"`
	// LicenseNumber             string  `json:"License_Number" validate:"max=10"` : Canceled on R3
	// SerialNumber              string  `json:"Serial_Number" validate:"max=40"` : Canceled on R3
	LandNumber string  `json:"Land_Number" validate:"max=10"`
	Area1Rai   float64 `json:"Area_1_Rai" validate:"omitempty,number"`
	Area2Ngan  float64 `json:"Area_2_Ngan" validate:"omitempty,number"`
	Area3SqWa  float64 `json:"Area_3_Sq_Wa" validate:"omitempty,number"`
	// Location                  string  `json:"Location" validate:"max=150"` : Canceled on R3
	CollateralSeller string `json:"Collateral_Seller" validate:"max=14"`
	// TotalUsableArea           float64 `json:"Total_Usable_Area" validate:"omitempty,number"` : Canceled on R3
	// BuildingName              string  `json:"Building_Name" validate:"max=255"` : Canceled on R3
	// SuiteNo                   string  `json:"Suite_No" validate:"max=50"` : Canceled on R3
	// Floor                     string  `json:"Floor" validate:"max=15"` : Canceled on R3
	// Number                    string  `json:"Number" validate:"max=30"` : Canceled on R3
	// Alley                     string  `json:"Alley" validate:"max=100"` : Canceled on R3
	// Road                      string  `json:"Road" validate:"max=100"` : Canceled on R3
	SubDistrict string `json:"Sub_district" validate:"max=20"`
	District    string `json:"District" validate:"max=20"`
	Province    string `json:"Province" validate:"max=20"`
	// R3 Region
	Region    string `json:"Region" validate:"max=20"`
	Latitude  string `json:"Latitude" validate:"max=20"`
	Longitude string `json:"Longitude" validate:"max=20"`
	// AgeOfBuilding    int    `json:"Age_of_Building" validate:"omitempty,numeric"` : Canceled on R3
	// DeveloperCompany string `json:"Developer_Company" validate:"max=255"` : Canceled on R3
}

type Employment struct {
	ApplicantType   string `json:"Applicant_Type" validate:"omitempty,oneof= B C G S"`
	ThaiName        string `json:"Thai_Name" validate:"max=255"`
	ThaiSurname     string `json:"Thai_Surname" validate:"max=255"`
	CitizenIdNumber string `json:"Citizen_Id_Number" validate:"max=20"`
	// CompanyName        string  `json:"Company_Name" validate:"max=100"` : Cancel on R3
	OccupationType string `json:"Occupation_Type" validate:"max=4"`
	Occupation     string `json:"Occupation" validate:"max=2"`
	// Position           string  `json:"Position" validate:"max=2"` : Cancel on R3
	// StartDate          string  `json:"Start_Date,omitempty" validate:"dateFormatYYYYMMDD,omitempty"` : Cancel on R3
	// WorkingPeriod      int     `json:"Working_Period" validate:"omitempty,numeric"` : Cancel on R3
	MonthlyIncome float64 `json:"Monthly_Income" validate:"omitempty,number"`
	// TotalMonthlyIncome float64 `json:"Total_Monthly_Income" validate:"omitempty,number"` : Cancel on R3
	// AvgMonthlyIncome   float64 `json:"Avg_Monthly_Income" validate:"omitempty,number"` : Cancel on R3
	// ProofOfIncome      string  `json:"Proof_of_Income" validate:"max=2"` : Cancel on R3
}

type Address struct {
	ApplicantType   string `json:"Applicant_Type" validate:"omitempty,oneof=B C G S"`
	ThaiName        string `json:"Thai_Name" validate:"max=255"`
	ThaiSurname     string `json:"Thai_Surname" validate:"max=255"`
	CitizenIdNumber string `json:"Citizen_Id_Number" validate:"max=20"`
	AddressType     string `json:"Address_Type" validate:"max=20"`
	Address1        string `json:"Address_1" validate:"max=255"`
	Address2        string `json:"Address_2" validate:"max=255"`
	Address3        string `json:"Address_3" validate:"max=255"`
	Address4        string `json:"Address_4" validate:"max=255"`
	Address5        string `json:"Address_5" validate:"max=255"`
	Province        string `json:"Province" validate:"max=3"`
	District        string `json:"District" validate:"max=3"`
	SubDistrict     string `json:"Sub_District" validate:"max=4"`
	ZipCode         string `json:"Zip_Code" validate:"max=10"`
	Country         string `json:"Country" validate:"max=2"`
	// Telephone       string `json:"Telephone" validate:"max=30"` : Cancel on R3
	// YearsOfResidence int    `json:"Years_Of_Residence" validate:"omitempty,numeric"` : Cancel on R3
	Longitude float64 `json:"Longitude" validate:"max=20"`
	Latitude  float64 `json:"Latitude" validate:"max=20"`
}

type BCLRequest struct {
	ApplicationNumber string `json:"Application_Number" validate:"required,max=50"`
	// Date is custom date
	ApplicationDate string `json:"Application_Date" validate:"required"`
	SystemSource    string `json:"System_Source" validate:"omitempty,required,max=255"`
	// SequenceNumber    int          `json:"Sequence_Number" validate:"omitempty,numeric"` : canceled on r3
	LoanAmount float64 `json:"Loan_Amount" validate:"omitempty,number"`
	// Channel           string       `json:"Channel" validate:"max=50"`
	ProductType       string `json:"Product_Type" validate:"max=50"`
	ApplicationType   string `json:"Application_Type" validate:"max=50"`
	ApplicationStatus string `json:"Application_Status" validate:"max=50"`
	SaleCode          string `json:"Sale_Code" validate:"max=10"`
	RequestCreator    string `json:"Request_Creator" validate:"max=10"`
	DecisionReason    string `json:"Decision_Reason" validate:"max=1024"`
	DecisionEmployee  string `json:"Decision_Employee" validate:"max=50"`
	DecisionDate      string `json:"Decision_Date,omitempty" validate:"omitempty,dateFormatYYYYMMDD"`
	// ReqNumOtherSys    string       `json:"Req_Num_Other_Sys" validate:"max=50"` : canceled on r3
	// StaffName     string       `json:"Staff_Name" validate:"max=255"` : canceled on r3
	// StaffID       string       `json:"Staff_ID" validate:"max=10"` : canceled on r3
	// StaffPosition string       `json:"Staff_Position" validate:"max=255"` : canceled on r3
	// StaffRole     string       `json:"Staff_Role" validate:"max=255"` : canceled on r3
	BorrowerType      string       `json:"Borrower_Type" validate:"max=1"`
	TotalLoanAmount   string       `json:"Total_Loan_Amount" validate:"max=20"`
	BranchSubDistrict string       `json:"Branch_Sub_District" validate:"max=20"`
	BranchDistrict    string       `json:"Branch_District" validate:"max=20"`
	BranchProvince    string       `json:"Branch_Province" validate:"max=20"`
	BranchRegion      string       `json:"Branch_Region" validate:"max=20"`
	BranchLatitude    float64      `json:"Branch_Latitude" validate:"omitempty,number"`
	BranchLongitude   float64      `json:"Branch_Longitude" validate:"omitempty,number"`
	Loans             []Loan       `json:"Loans" validate:"dive"`
	Applicants        []Applicant  `json:"Applicants" validate:"dive"`
	Addresses         []Address    `json:"Addresses" validate:"dive"`
	Employments       []Employment `json:"Employments" validate:"dive"`
	Businesses        []Business   `json:"Businesses" validate:"dive"`
	Collaterals       []Collateral `json:"Collaterals" validate:"dive"`
}
