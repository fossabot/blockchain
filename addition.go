// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

// ValidateBitcoinAddress bitcoin address validator.
func ValidateBitcoinAddress(address string) bool {
	return validateBitcoinAddress(address) != -1
}

// ValidateBitcoinXpub bitcoin address validator.
func ValidateBitcoinXpub(xpub string) bool {
	return validateBitcoinXpub(xpub) != -1
}

func (c *Client) checkAddress(address string) error {
	if !ValidateBitcoinAddress(address) {
		return c.err4(ErrAIW, address)
	}

	return nil
}

func (c *Client) checkAddresses(addresses []string) (e error) {
	if len(addresses) == 0 {
		return c.err(ErrNAP)
	}

	for _, address := range removeDuplicates(addresses) {
		if !ValidateBitcoinAddress(address) && !ValidateBitcoinXpub(address) {
			return c.err4(ErrAIW, address)
		}
	}

	return nil
}

func removeDuplicates(elements []string) (result []string) {
	encountered := map[string]bool{}
	for _, v := range elements {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return
}

// addressesForTestings list of addresses to test and get the maximum
// number of addresses to be checked at the same time.
var addressesForTestings = []string{
	"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
	"12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX",
	"1HLoD9E4SDFFPDiYfNYnkBLQ85Y51J3Zb1",
	"1FvzCLoTPGANNjWoUo6jUGuAG3wg1w4YjR",
	"15ubicBBWFnvoZLT7GiU2qxjRaKJPdkDMG",
	"1JfbZRwdDHKZmuiZgYArJZhcuuzuw2HuMu",
	"1GkQmKAmHtNfnD3LHhTkewJxKHVSta4m2a",
	"16LoW7y83wtawMg5XmT4M3Q7EdjjUmenjM",
	"1J6PYEzr4CUoGbnXrELyHszoTSz3wCsCaj",
	"12cbQLTFMXRnSzktFkuoG3eHoMeFtpTu3S",
	"15yN7NPEpu82sHhB6TzCW5z5aXoamiKeGy",
	"1dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9",
	"1PYELM7jXHy5HhatbXGXfRpGrgMMxmpobu",
	"17abzUBJr7cnqfnxnmznn8W38s9f9EoXiq",
	"1DMGtVnRrgZaji7C9noZS3a1QtoaAN2uRG",
	"1CYG7y3fukVLdobqgUtbknwWKUZ5p1HVmV",
	"16kktFTqsruEfPPphW4YgjktRF28iT8Dby",
	"1LPBetDzQ3cYwqQepg4teFwR7FnR1TkMCM",
	"1DJkjSqW9cX9XWdU71WX3Aw6s6Mk4C3TtN",
	"1P9VmZogiic8d5ZUVZofrdtzXgtpbG9fop",
	"15ubjFzmWVvj3TqcpJ1bSsb8joJ6gF6dZa",
	"1Fi7o3BKMcT82NVtnMRNqsj8aE5CWbAo4z",
	"1DmKBaveG8iQA7nTGpRQ1bf8cJ7zqSMCjJ",
	"1AqC4PhwYf7QAyGBhThcyQCKHJyyyLyAwc",
	"1JXLFv719ec3bzTXaSq7vqRFS634LErtJu",
	"15VF3MsCzjHmFQ3wK3SMrTEBTmFY8zhJnU",
	"1KDnp5D9sEUNGFY2z4JiKhwfbm7zqYp4Bz",
	"1G43MvhzCqRz1ctsQUmgU4LgLuSVdfU557",
	"18a3JX7ZZvu5CdaqY33RxzXB3ytEjspsGq",
	"1GnYgH4V4kHdYEdHwAczRHXwqxdY7xars1",
	"17x23dNjXJLzGMev6R63uyRhMWP1VHawKc",
	"1PHB5i7JMEZCKvcjYSQXPbi5oSK8DoJucS",
	"1Bw7RG9a19SjCNszmXQZhBwe1gEj4Vb2JZ",
	"1MKRkcXG7cUb83EGNjK5TSHcKgVjVMTou5",
	"1BTsxjF9rXtFvUZ2UFditbeUpohGgKCxUt",
	"16bo8EmUnLJAmtDNavgcs1BQ5rU4YcR9UC",
	"1NHHWnZKTetjZZfAbVjTS9EHQafdH6xn8a",
	"1DpJP16GXSht3hJ5DZDumKcwAFqQQVkffN",
	"1DpSDwF12wX5ogkS2HXabT7iWtCkAqJD9k",
	"1KzvBTUbdwNBXiTkzr1msFUtPf7Vu2zLiu",
	"1AQ2m6GH78oLgZYCdueUb4Zjxg1L6BkHZM",
	"1DEwdHYmo8q6AhSHG7UgxEjttNMFdw9e7u",
	"1DnzNm7QC479kSydvzmgb16X2c5GwRVzs1",
	"1Bhq4Uz1Cg8vZLWy5uT3X3Ea3XDSNvkvCM",
	"1HjE5zRg4eYm7jCHB9oSH5WbyZdH1GE2Ns",
	"1D2zWLwXkdgFw5VtcYVtBhZuq5j4qwiRQP",
	"1BEKzoSq5kTRNhL9ezh8KJ5vt2NzigQV2K",
	"136g5f3pBidF6q8udL5GGb66CzG7AzBNwv",
	"1FaUNstQc4FUena9XodmbfHgoMD9yKpT4d",
	"18K6LJ9wFKSkYDA3dspswRPWiPgNZKvN7d",
	"17iyRRXBHJKbv5DKPPkttWewm3CHdNPGQd",
	"192kkZSRQHrS4dnBoPGgoYB6WeC1HXyYZe",
	"17xQEk6df1Zyzvzs3oLPU9k5BZ85uGxC19",
	"1h2Znru8Y624tuVnu1UgUujNACTzaFABz",
	"1Kc4RPywLu8bK7JvTQJY1MWHo1S4jqQfzM",
	"18p2otBZ33Ytgk3afUrHsmiPrFLNPiM4G5",
	"1281ZBYNe7qJGMsGgHL2YgUEEFwCdXQwBq",
	"19knNYsDH25AsBGhhoaxnmUKrn9WkVMSUi",
	"14UyXKWBcVHYcqSpxYcmR9JLZu6pArwDrd",
	"1FERXfs9RYGhhLe2yBsSqYfSqtmUEBUSTc",
	"16kf2cAY8V7iZjX21Pm8znMJxqa2TYcqQ5",
	"1AiE8NiskCXRovzHdVzAjydQ9V8Dv7Hiy8",
	"12GBk2eiu3AEVvNYTwWYnUz49ecRjJ1CR6",
	"16jZuw5nwn3ZgzWYrRvhi8NZW2b1NRk6S7",
	"197x9cjGHGq9ah4AAwduRbu48J9PLiwcFD",
	"1UZhhzWo85osGzNrs1BVjoE3FP8ea2umX",
	"1PeWcJrkpHMnUs3t5GvuwcHh1oBHYNoDwG",
	"17uWtnxcJdpH7pEXXmr1SLrohVxe18B3fD",
	"1GPrHBrbFjzGVTwA4P4y5QnaDhd3GCYU5x",
	"1JgpRwCt8rm29g3y4FkVKWWck8qEeKgfZv",
	"1NnYa2jL24hLgXBbk3TAXHANUQEzXNnSHg",
	"19CAnt15TpQrJ9gjmdZZsLcwRj8E1mkqPP",
	"1K2UaAY9jfz7wUxadbD3ivBLGqrcF4ai8M",
	"1Ky8dP7oR1cBeg1MzkrgHAeHAHyn92DCar",
	"1FnbdYntfohuZ1EhZ7f9oiT2R5sDsZBohL",
	"14U5EYTN54agAngQu92D9gESvHYfKw8EqA",
	"19X7soRFwUCRX8aUc3NrRZddHgpLmFmpzc",
	"1H2nuVj2EnEVoYSw3opGjVxJdDqnehiYWN",
	"12ti6ZDtjJTFnfSXJD1yZeojFdc1u4nJcS",
	"1BwWdLV5wbnZvSYfNA8zaEMqEDDjvA99wX",
	"1KmbQCNdfCDkEB6yWw5iLjBTGJ5MYptqAL",
	"1NHVxANujgiDvPxNLQ1cc5B8AEisnx2Yz6",
	"1DmJE9Z1Teye3JmxhoDUcmnMmubCg8dNtF",
	"1HWP2q1RjpRbjm2K5ZPFnnfbnpNjN7XJWj",
	"135uQWV3vn5JNgQrWcLd6gWr9S2ybj2BB7",
	"15eCuQ4Kd25VDhtv9Pko6DL3tt1ACa3tR8",
	"12jiJuWynwFcqbNyoay5Sy9pmJn9dsFRgh",
	"133xmQyVGHnyQwr2pYNiHUMeDUS3mYVsB7",
	"1KSHc1tmsUhS9f1TD6RHR8Kmwg9Zv8WhCt",
	"1KMYVBvfMAUZH4gpRknxvKxbqATMAfuu4x",
	"1BnhHuk62o9tPuG5JerexTX3MhnR8sobLm",
	"1DRzyGtGkXYnFz9n9qxnTFbGpYBv4FNmYM",
	"1LNV5xnjneJwXc6jN8X2co586gjiSz6asS",
	"1PHG1AoZJ5yyg6MV2fkqYCuDDik2NTfzfm",
	"18GyZ216oMhpCbZ7JkKZyT8x68v2a8HuNA",
	"12XPHPCGYz1WgRhquiAfVeAyjZ7Gbdpih3",
	"1Bq1TLyuvZLvu73eueduUMhBL2QfYDeoNt",
	"1yNErqkG6GFaQTkRjrp38nqpVFn7QJnhV",
	"16cAVR3SQbNzu8KZtGdo8cG1iueWpcngxz",
	"13A1W4jLPP75pzvn2qJ5KyyqG3qPSpb9jM",
	"13CiNuEMKASJBvGdupqaoRs2MqDNhAqmce",
}
