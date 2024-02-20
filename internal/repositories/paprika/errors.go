package paprika

import "errors"

/*
	Known error responses
	=====================

	{
	   "type":"payment_required",
	   "hard_limit":"25k requests per month",
	   "soft_limit":"60 requests per hour",
	   "error":"Payment Required - this request could not be processed; request limits were reached. If you want to be able to process this API request please upgrade your plan at https://coinpaprika.com/api/",
	   "block_duration":"1h"
	}

	{"error":"id not found"}
*/

var (
	ErrCoinIDNotFound = errors.New("id not found")
	ErrLimitData      = errors.New("Payment Required - this request could not be processed; request limits were reached. If you want to be able to process this API request please upgrade your plan at https://coinpaprika.com/api/")
	ErrUnknown        = errors.New("unknown paprika error")
)
