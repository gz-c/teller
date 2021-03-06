package exchange

import "github.com/skycoin/teller/src/daemon"

// Client provides helper apis to interact with exchange service
type Client struct {
	s *Service
}

// NewClient creates exchange client
func NewClient(s *Service) *Client {
	return &Client{s: s}
}

// BindAddress binds deposit btc address with skycoin address, and
// add the btc address to scan service, when detect deposit coin
// to the btc address, will send specific skycoin to the binded
// skycoin address
func (ec *Client) BindAddress(btcAddr, skyAddr string) error {
	return ec.s.addDepositInfo(btcAddr, skyAddr)
}

// GetDepositStatuses returns deamon.DepositStatus array of given skycoin address
func (ec *Client) GetDepositStatuses(skyAddr string) ([]daemon.DepositStatus, error) {
	return ec.s.getDepositStatuses(skyAddr)
}
