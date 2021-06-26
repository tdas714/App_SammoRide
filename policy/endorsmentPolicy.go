package policy

import "time"

type EndorsmentPolicy struct {
	NumberOfEndorsers  int
	NumberOfSignatures int
	ProposalTimeDiff   time.Duration
}

type WritersPolicy struct {
	CertificateValidation []bool
	DriverCheck           string
}

func GetEndorsmentPolicy() *EndorsmentPolicy {
	endor := EndorsmentPolicy{NumberOfEndorsers: 3, NumberOfSignatures: 2,
		ProposalTimeDiff: 10 * time.Second}
	return &endor
}

func GetWritersPolicy() *WritersPolicy {
	writerp := WritersPolicy{[]bool{true, true}, "Driver"}
	return &writerp
}
