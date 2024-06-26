package handler

import (
	"dgl-kafka-publisher/models"
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
)

type MockupIntutionHandler struct {
}

func NewMockupIntutionHandler() *MockupIntutionHandler {
	return &MockupIntutionHandler{}
}

func (h *MockupIntutionHandler) PostToInstution(ctx *fiber.Ctx) error {
	var body interface{}

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	data, err := os.ReadFile("mockup/resp_intuition.json")
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	var text interface{}
	err = json.Unmarshal(data, &text)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(body)
}

func (h *MockupIntutionHandler) PostToInstutionV2(ctx *fiber.Ctx) error {
	var body interface{}

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(body)

}

func (h *MockupIntutionHandler) GenerateBCLMock(ctx *fiber.Ctx) error {
	m := new(models.BCLRequest)

	m = &models.BCLRequest{
		ApplicationNumber: "123456",
		ApplicationDate:   "2021-01-01",
		SystemSource:      "DGL",
		LoanAmount:        1000000,
		ProductType:       "-",
		ApplicationType:   "-",
		ApplicationStatus: "-",
		SaleCode:          "-",
		RequestCreator:    "",
		DecisionReason:    "",
		DecisionEmployee:  "",
		DecisionDate:      "",
		BorrowerType:      "",
		TotalLoanAmount:   "",
		BranchSubDistrict: "",
		BranchDistrict:    "",
		BranchProvince:    "",
		BranchRegion:      "",
		BranchLatitude:    0,
		BranchLongitude:   0,
		Loans: []models.Loan{
			{
				LoanAmount:              1000000,
				LoanPurpose:             "-",
				ProductCode:             "-",
				MarketingCode:           "-",
				InstallmentAmount:       10000,
				ProductType:             "-",
				ProductSubType:          "-",
				ExistingCreditLimit:     0,
				ChangingCreditLimit:     0,
				ArrangementPurpose_Code: "-",
				LoanAccountNumber:       "-",
				Refinance:               false,
			},
		},
		Applicants: []models.Applicant{
			{
				ApplicantType:            "-",
				ThaiName:                 "-",
				EngName:                  "-",
				ThaiSurname:              "-",
				EngSurname:               "-",
				CitizenIdNumber:          "-",
				ThaiMiddleName:           "-",
				ThaiMaidenName:           "-",
				CIFNumber:                "-",
				Sex:                      "-",
				DateOfBirth:              "-",
				Age:                      0,
				Nationality:              "-",
				MaritalStatus:            "-",
				EducationLevel:           "-",
				PassportID:               "-",
				CountryOfIssue:           "-",
				MobilePhone:              "-",
				Email:                    "-",
				RelationShipWithBorrower: "-",
				KYCCDDStatus:             "-",
				TUEFID:                   "-",
			},
		},
		Addresses: []models.Address{
			{
				ApplicantType:   "-",
				ThaiName:        "-",
				ThaiSurname:     "-",
				CitizenIdNumber: "-",
				AddressType:     "-",
				Address1:        "-",
				Address2:        "-",
				Address3:        "-",
				Address4:        "-",
				Address5:        "-",
				Province:        "-",
				District:        "-",
				SubDistrict:     "-",
				ZipCode:         "-",
				Country:         "-",
				Longitude:       0.0,
				Latitude:        0.0,
			},
		},
		Employments: []models.Employment{{
			ApplicantType:   "-",
			ThaiName:        "-",
			ThaiSurname:     "-",
			CitizenIdNumber: "-",
			OccupationType:  "-",
			Occupation:      "-",
			MonthlyIncome:   0,
		}},
		Businesses: []models.Business{{
			ApplicantType:        "-",
			ThaiName:             "-",
			ThaiSurname:          "-",
			CitizenIdNumber:      "-",
			BusinessType:         "-",
			TaxIDNumber:          "-",
			RegistrationDate:     "-",
			DurationOfBusiness:   0,
			BusinessName:         "-",
			BusinessNameEng:      "-",
			AnnualSales:          0,
			WebsiteName:          "-",
			BusinessPurposeCode1: "-",
			KTBCustomerCode:      "-",
			InvolvePartyGroup:    "-",
			KYCCDDStatus:         "-",
			NetFixedAsset:        "0",
			TotalNoOfEmployees:   "0",
			RegistrationCountry:  "-",
			BusinessStatus:       "-",
		}},
		Collaterals: []models.Collateral{{
			ApplicantType:             "-",
			ThaiName:                  "-",
			ThaiSurname:               "-",
			CitizenIdNumber:           "-",
			CollateralNumber:          "-",
			TypeOfCollateral:          "-",
			CollateralCondition:       "-",
			SubCollateralType:         "-",
			AmountOfMortgage:          0,
			AppraisalPrice:            0,
			AppraisalAssessmentAgency: "-",
			DateOfLastAppraisal:       "-",
			LandNumber:                "-",
			Area1Rai:                  0,
			Area2Ngan:                 0,
			Area3SqWa:                 0,
			CollateralSeller:          "-",
			SubDistrict:               "-",
			District:                  "-",
			Province:                  "-",
			Region:                    "-",
			Latitude:                  "0",
			Longitude:                 "0",
		}},
	}

	return ctx.JSON(m)
}
