package shared

// variabel const untuk error response
const (
	//variabel const untuk error pada autentikasi
	RespSuccess = "0200"
)

// variabel const untuk error response
const (
	//variabel const untuk error pada autentikasi
	ErrCodeCredentialsRequired = "0101"
	ErrCodeErrorSaveCredential = "0102"

	//variabel const untuk error pada user
	ErrCodeFailCreateUser = "0201"
	ErrCodeFailUpdateUser = "0202"

	//variabel const untuk error pada submission
	ErrCodeFailCreateSubmission = "0301"

	//variabel const untuk error pada partner
	ErrCodeFailUpdatePartner = "0402"
	ErrCodeFailDeletePartner = "0403"

	//other errors
	Unauthorized             = "0401"
	ErrCodeServerError       = "0500"
	ErrCodeTimeout           = "0501"
	ErrCodeErrorValidation   = "0502"
	ErrFileLarge             = "0503"
	ErrInvalidFile           = "0504"
	ErrFileEmpty             = "0505"
	ErrUnavailableService    = "0506"
	ErrInvalidParam          = "0507"
	ErrDataNotExist          = "0509"
	ErrUnexpectedError       = "0510"
	ErrInvalidAccount        = "0511"
	ErrInactiveAccount       = "0512"
	ErrTimeOut               = "0513"
	ErrInvalidRequestFamily  = "0514"
	ErrInvalidFieldFormat    = "0515"
	ErrInvalidFieldMandatory = "0516"
	ErrCodeEmptyParameter    = "0517"
	ErrDocStatus             = "0518"
	ErrInvalidMinFieldFormat = "0519"
)
