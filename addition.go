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
		return c.setError(ErrAIW, nil, nil, &address)
	}

	return nil
}

func (c *Client) checkAddresses(addresses []string) (e error) {
	if len(addresses) == 0 {
		return c.setErrorOne(ErrNAP)
	}

	for _, address := range removeDuplicates(addresses) {
		if !ValidateBitcoinAddress(address) && !ValidateBitcoinXpub(address) {
			return c.setError(ErrAIW, nil, nil, &address)
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
	"1HFtrEq1YGKxFKdq5nKyd7xm6NrAKcT7zf",
	"1B4HeenWZfzcyHSpHeZGQZtSVUi18wmd2B",
	"1MBvf6WPKwekXLRxsjjudqUXWJDYUy4qzB",
	"1Bjrramo5EisLKeUXjw9GcUcfXcVnpi6vK",
	"1T2TbagADcZFN433zk1HjYuTiXTPs62Yd",
	"13t2rxB1FiZ8rSR3jVVNAmJN7hxFCq6pjp",
	"1PN2MiaGhnR65x41pbUWbsfCB62hqG8KSG",
	"1Cef65T8Cqi47KJvc4vYkxuZ9ymBxQfhQf",
	"1CbjjM7LDBvmtdvyyZbbb8hp3B8gwUEaio",
	"1HCeeU957J4NTXDer2fGvDsb7mVjU9TtLb",
	"1JNWvsrDf8mdH82Y7xp8yACDfm6TpVwZeM",
	"19K4cNVYVyNiwZ5xkzjW9ZtMb8XvBS2LkT",
	"1NG4uecvtRyRx9eWxUaw6PVLPKMfDE7nz4",
	"14w8FgPQhPRuP4wnVpER9p9Cb7kYQWkzwY",
	"1MF2sw6Zan5p4ViuGdxRoNr48M5rFvHkqn",
	"18NmcVxBKf5tZXuV5KijPTgbdguGohnBCV",
	"1N24Lobvo9BN4iGrm4SRhCUwpwd273dSfq",
	"1AUR9CVmm5wH8xc2EufyDpuA8HEZcSamhY",
	"1FeGetWU2tR2QSrxnpRwHGXGcxzhN6zQza",
	"1B7CyZF8e6TYzhNBSHy8yYuTRJNpMtNChg",
	"1nQ5KafzA4DWGa327n8JnZ29q4PKg5Y7m",
	"1CT9huFgxMFveRvzZ7zPPJNoaMm2Fo64VH",
	"1M7U7kbW9CY7znW6sSRzycgPCxr5LkJy2r",
	"1JRiK5gbEjTEAdFU2kYoQJRCMZgVCQgbpQ",
	"1PwcmHdjVFfQgXpd6V4BGzRY7yLBvGb4eK",
	"15kHBARv1ZaCrmyExVUn6Uh7wnJVXveF8A",
	"135S6YvALYCydS6R82hWi967f3xUnWVQ3S",
	"1BF5nbuKF3UTYCQTpxuH62UA8WJWAqQKo5",
	"1PZnnfPq3EpcGZZyiX8WCbEhm9amzpjTYr",
	"13aK6boaLbWkyLsDBmNZnUGhzh4iL2vQky",
	"1GPhCZpM72YrpkZVqQECpT6eKprKFEwkSH",
	"1PpU2yx6CSruakyYjckH8eUJWAH1XUQ7ZM",
	"13pS5WuihmpE7WMQAZkFdvfvVo1NDp7R7V",
	"1EDV8iD4cDZbKxZu6UGAPW5i2ytQmR5xoQ",
	"1MMFzwh3r5fxNwpTZvm6iZReX7RUwZxnaY",
	"1FdgjwkXbQc1gpHjpHyZNY85vdyQyFfBbg",
	"1HRRURSfxj8NntfzcPT4ajz7uBefLuwLmB",
	"1G6RMmurphMJg6wyobQSzC2az9zZ3VipXE",
	"1ArDoL42A6ZZdudpogut2NUGcwxhRj42S8",
	"12dRaqF6c7n8rvbhCWsQysKmNKPgDJfepK",
	"17B3kpD6d9SSaoKATBgGqDn25dFnSvjcVE",
	"1DQCdPGY1MUFwPdJ2bNziT4RfDFusBddyS",
	"1Jd2omwPAZq3egSS8HsxFMxhWmKdonG9Lz",
	"1CwoMxSXitseAKdNgfaHtzaUB7mUhcrmMp",
	"18xph1cecYkJwD81bHSKZ3SrUsPqinU2bA",
	"1Pqrx3NMF9VmMW6ikLZST3veBkUmY6be15",
	"15EMUiFD7XoHpzCA4sErNmqCT4krsQzgKx",
	"18xzfpkj8cNz8NdmU7eumVjpeDfQSJ1CHR",
	"1N8SwzBqNQccJQBMsVe5ukvSVhf6vxHVSm",
	"1DLfzP5v2KegZK88rsNj8z1YKSb4VzBfsu",
	"17KG5gGwKu1DNNaMdrngw3UdZUh4ZhacNT",
	"1B1urSvRRSFjupa5THHsF8FNxTarDVexYg",
	"1QHTV856dWWwzSQN738tvEneYgsagQdYpM",
	"1AQgsAVd49Tk9kxgsdDRFMwLjyD997p6dJ",
	"1J9Wdx5F57jgNYSw7vaP3oCMWHnjaD3MEq",
	"1CtABfwubzxuoHacNEZ6FJCBumtTbz71AA",
	"1H5canZg3AwU44Fe65HG2siti9mFJnkc8",
	"1KuX1qeF7qimFiurBDxQUktAgNE3qNA8zx",
	"1ax2fn3tdpH1R3A2yNFtp8Fn7Z7rt9DnY",
	"1ELLM7xHcUC1Z2BnDpGQ2Eg2ZXZeGSDkbU",
	"1HgeuTkMfQ3xh6hPo9ZZKvHom8AAXdfA4X",
	"1Fdby4yp1iUnf8biiVAE14X4XKV6ZKzReR",
	"1Lh4fvU1NAEdZTuTqoe4c3cPMv3sgkFWbM",
	"1Eof6XM7AKnBdpTt3c92sVb1LaLqH3cnqC",
	"1BQRxmgPsyWWv6GD64kUDxtNq3K2RGc9KL",
	"1Fr33GRhy7zN8eXkWDf31icUBRUmoicnJz",
	"1NZz1UN616LfyX87zGRnVifiWRSN5xm9ht",
	"13ph5zPCBLeZcPph9FBZKeeyDjvU2tvcMY",
	"1PSSGeFHDnKNxiEyFrD1wcEaHr9hrQDDWc",
	"1JhxuRHf7FiJYb8Ymc9czohwm9cJn2WaQW",
	"19d4S14KCEDXi1Xyt9x98feD9SxWPXYYNm",
	"1LNKBdzwXmisa67qtYCofsZWwgZW2wewmn",
	"1FWcb8Ddw3BXearNaWugikvEfK7gViRde7",
	"1917k1LUodPbbjJenczkVaGyJpSj8CqYTC",
	"1H5BMx61bxfseTU9jUPZgit5B9D3AVviX5",
	"1PA4XBfKLjACMQUUupBEDzwhSsLRKriAkr",
	"128KyT3PMB3WLcZVu48asfaTwtb6tZE2dm",
	"1HPgEWbEzFC6FWbcEJtasvZBaW5YTdmfMC",
	"1Pudc88gyFynBVZccRJeYyEV7ZnjfXnfKn",
	"1JSW4QekxPokWWU4hcRwrheZbZKSkFz9oc",
	"1Fr947YZyEWZd2JPcvDJbsYN6Po5gXRyau",
	"1642o19pahkkdwHuPSx6uxk9HMmC243Bu7",
	"1PNF7mJ2VrgZCMSfUjiWxDteNBHEm15kMg",
	"1H4Urq49poTRfszssMDT12dMnoDxSBEs7e",
	"16RGBFEHfNqjzN6UQ95imqtM91PBcQUMCF",
	"1FDMwEo8qNa9icVcooBUoGvA6NriePtJJ3",
	"176n9Cx6LKri6YUxFNBr7aUzb2TbLzJbsM",
	"1EPeRrfFgn8Dm7mbfaFHUNUpD5UnKkeSNQ",
	"1MNHWu6JeCrtFyjEPcR6bJSrfQ3kJ3Hyin",
	"1QEDtTS36hrcJuv6SvhqoUvVrfMLdkaN7C",
	"15HWA4SmQ9zm9iYGyxPK7r1oDugPXeWwDk",
	"1NEAg9cRUBUmTV8FG5ed33nPiXwEKaCxra",
	"12RyzPinkywNAN1L78tpBS14QTFQCBhuLE",
	"1Kj7zDjmTcjadztrU1LWT1erGuyToScwex",
	"1Dw5XyQk2pFDNLQUsc2Eknq5ikQAwfRqsR",
	"1F22VKzwQWf2LoSziz78Zc5e7owfrHDRCy",
	"1AL11Yfw2ZP9C4wGiWf99jwZyTjWWY4JVn",
}
