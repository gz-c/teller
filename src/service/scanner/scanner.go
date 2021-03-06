package scanner

// Scanner provids helper apis to interact with scan service
type Scanner struct {
	s *ScanService
}

// NewScanner creates scanner instance
func NewScanner(s *ScanService) *Scanner {
	return &Scanner{s: s}
}

// AddDepositAddress add new deposit address to scan service
func (s *Scanner) AddDepositAddress(addr string) error {
	return s.s.AddDepositAddress(addr)
}

// GetDepositValue returns deposit value channel
func (s *Scanner) GetDepositValue() <-chan DepositValue {
	return s.s.depositC
}
