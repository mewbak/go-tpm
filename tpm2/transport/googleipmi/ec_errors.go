package googleipmi

import (
	"fmt"
)

// decodeError returns the given error code as an error.
// Returns:
// * nil if the result code was EC_RES_SUCCESS
// * an ecError if the error was within the EC error space
// * an extendedEcError otherwise
func decodeError(code uint16) error {
	if code == success {
		return nil
	}
	if ecErr := ecError(code); ecErr <= lastECError {
		return ecErr
	}
	return extendedEcError(code)
}

// ecError represents an EC error code within the general EC error space.
type ecError uint16

const (
	// success is EC_RES_SUCCESS, and is intentionally not exported or defined as an error (or an ecError).
	success uint16 = 0
	// errInvalidCommand is EC_RES_INVALID_COMMAND
	errInvalidCommand ecError = 1
	// errGenericError is EC_RES_ERROR
	errGenericError ecError = 2
	// errInvalidParam is EC_RES_INVALID_PARAM
	errInvalidParam ecError = 3
	// errAccessDenied is EC_RES_ACCESS_DENIED
	errAccessDenied ecError = 4
	// errInvalidResponse is EC_RES_INVALID_RESPONSE
	errInvalidResponse ecError = 5
	// errInvalidVersion is EC_RES_INVALID_VERSION
	errInvalidVersion ecError = 6
	// errInvalidChecksum is EC_RES_INVALID_CHECKSUM
	errInvalidChecksum ecError = 7
	// errInProgress is EC_RES_IN_PROGRESS
	errInProgress ecError = 8
	// errUnavailable is EC_RES_UNAVAILABLE
	errUnavailable ecError = 9
	// errTimeout is EC_RES_TIMEOUT
	errTimeout ecError = 10
	// errOverflow is EC_RES_OVERFLOW
	errOverflow ecError = 11
	// errInvalidHeader is EC_RES_INVALID_HEADER
	errInvalidHeader ecError = 12
	// errRequestTruncated is EC_RES_REQUEST_TRUNCATED
	errRequestTruncated ecError = 13
	// errResponseTooBig is EC_RES_RESPONSE_TOO_BIG
	errResponseTooBig ecError = 14
	// errBusError is EC_RES_BUS_ERROR
	errBusError ecError = 15
	// errBusy is EC_RES_BUSY
	errBusy ecError = 16
	// errInvalidHeaderVersion is EC_RES_INVALID_HEADER_VERSION
	errInvalidHeaderVersion ecError = 17
	// errInvalidHeaderCRC is EC_RES_INVALID_HEADER_CRC
	errInvalidHeaderCRC ecError = 18
	// errInvalidDataCRC is EC_RES_INVALID_DATA_CRC
	errInvalidDataCRC ecError = 19
	// errDupUnavailable is EC_RES_DUP_UNAVAILABLE
	errDupUnavailable ecError = 20
	// lastECError is the highest-valued error in the ec namespace.
	lastECError = errDupUnavailable
)

// Error returns the string representation of the ecError.
func (code ecError) Error() string {
	switch code {
	case errInvalidCommand:
		return "EC_RES_INVALID_COMMAND"
	case errGenericError:
		return "EC_RES_ERROR"
	case errInvalidParam:
		return "EC_RES_INVALID_PARAM"
	case errAccessDenied:
		return "EC_RES_ACCESS_DENIED"
	case errInvalidResponse:
		return "EC_RES_INVALID_RESPONSE"
	case errInvalidVersion:
		return "EC_RES_INVALID_VERSION"
	case errInvalidChecksum:
		return "EC_RES_INVALID_CHECKSUM"
	case errInProgress:
		return "EC_RES_IN_PROGRESS"
	case errUnavailable:
		return "EC_RES_UNAVAILABLE"
	case errTimeout:
		return "EC_RES_TIMEOUT"
	case errOverflow:
		return "EC_RES_OVERFLOW"
	case errInvalidHeader:
		return "EC_RES_INVALID_HEADER"
	case errRequestTruncated:
		return "EC_RES_REQUEST_TRUNCATED"
	case errResponseTooBig:
		return "EC_RES_RESPONSE_TOO_BIG"
	case errBusError:
		return "EC_RES_BUS_ERROR"
	case errBusy:
		return "EC_RES_BUSY"
	case errInvalidHeaderVersion:
		return "EC_RES_INVALID_HEADER_VERSION"
	case errInvalidHeaderCRC:
		return "EC_RES_INVALID_HEADER_CRC"
	case errInvalidDataCRC:
		return "EC_RES_INVALID_DATA_CRC"
	case errDupUnavailable:
		return "EC_RES_DUP_UNAVAILABLE"
	default:
		return fmt.Sprintf("unknown EC error code (%x)", uint16(code))
	}
}

// extendedEcError represents an error code returned from an EC that is outside
// the EC error space. Packages that send EC commands may cast errors of this
// type into their own error space.
type extendedEcError uint16

// Error returns the string representation of the extendedEcError.
func (code extendedEcError) Error() string {
	return fmt.Sprintf("extended EC error code (%x)", uint16(code))
}
