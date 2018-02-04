package TransmissionRPC

import (
	"fmt"
)

/*
	Torrent Action Requests
	https://trac.transmissionbt.com/browser/tags/2.92/extras/rpc-spec.txt?rev=14714#L71
*/

type torrentActionIDsParam struct {
	IDs []int64 `json:"ids,omitempty"`
}

type torrentActionHashesParam struct {
	IDs []string `json:"ids,omitempty"`
}

type torrentActionRecentlyActiveParam struct {
	IDs string `json:"ids"`
}

// TorrentStartIDs starts torrent(s) which id is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentStartIDs(ids []int64) (err error) {
	if err = c.rpcCall("torrent-start", &torrentActionIDsParam{IDs: ids}, nil); err != nil {
		err = fmt.Errorf("'torrent-start' rpc method failed: %v", err)
	}
	return
}

// TorrentStartHashes starts torrent(s) which hash is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentStartHashes(hashes []string) (err error) {
	if err = c.rpcCall("torrent-start", &torrentActionHashesParam{IDs: hashes}, nil); err != nil {
		err = fmt.Errorf("'torrent-start' rpc method failed: %v", err)
	}
	return
}

// TorrentStartRecentlyActive starts torrent(s) which have been recently active.
func (c *Controller) TorrentStartRecentlyActive() (err error) {
	if err = c.rpcCall("torrent-start", &torrentActionRecentlyActiveParam{IDs: "recently-active"}, nil); err != nil {
		err = fmt.Errorf("'torrent-start' rpc method failed: %v", err)
	}
	return
}

// TorrentStartNowIDs start-nows torrent(s) which id is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentStartNowIDs(ids []int64) (err error) {
	if err = c.rpcCall("torrent-start-now", &torrentActionIDsParam{IDs: ids}, nil); err != nil {
		err = fmt.Errorf("'torrent-start-now' rpc method failed: %v", err)
	}
	return
}

// TorrentStartNowHashes start-nows torrent(s) which hash is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentStartNowHashes(hashes []string) (err error) {
	if err = c.rpcCall("torrent-start-now", &torrentActionHashesParam{IDs: hashes}, nil); err != nil {
		err = fmt.Errorf("'torrent-start-now' rpc method failed: %v", err)
	}
	return
}

// TorrentStartNowRecentlyActive start-nows torrent(s) which have been recently active.
func (c *Controller) TorrentStartNowRecentlyActive() (err error) {
	if err = c.rpcCall("torrent-start-now", &torrentActionRecentlyActiveParam{IDs: "recently-active"}, nil); err != nil {
		err = fmt.Errorf("'torrent-start-now' rpc method failed: %v", err)
	}
	return
}

// TorrentStopIDs stops torrent(s) which id is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentStopIDs(ids []int64) (err error) {
	if err = c.rpcCall("torrent-stop", &torrentActionIDsParam{IDs: ids}, nil); err != nil {
		err = fmt.Errorf("'torrent-stop' rpc method failed: %v", err)
	}
	return
}

// TorrentStopHashes stops torrent(s) which hash is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentStopHashes(hashes []string) (err error) {
	if err = c.rpcCall("torrent-stop", &torrentActionHashesParam{IDs: hashes}, nil); err != nil {
		err = fmt.Errorf("'torrent-stop' rpc method failed: %v", err)
	}
	return
}

// TorrentStopRecentlyActive stops torrent(s) which have been recently active.
func (c *Controller) TorrentStopRecentlyActive() (err error) {
	if err = c.rpcCall("torrent-stop", &torrentActionRecentlyActiveParam{IDs: "recently-active"}, nil); err != nil {
		err = fmt.Errorf("'torrent-stop' rpc method failed: %v", err)
	}
	return
}

// TorrentVerifyIDs verifys torrent(s) which id is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentVerifyIDs(ids []int64) (err error) {
	if err = c.rpcCall("torrent-verify", &torrentActionIDsParam{IDs: ids}, nil); err != nil {
		err = fmt.Errorf("'torrent-verify' rpc method failed: %v", err)
	}
	return
}

// TorrentVerifyHashes verifys torrent(s) which hash is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentVerifyHashes(hashes []string) (err error) {
	if err = c.rpcCall("torrent-verify", &torrentActionHashesParam{IDs: hashes}, nil); err != nil {
		err = fmt.Errorf("'torrent-verify' rpc method failed: %v", err)
	}
	return
}

// TorrentVerifyRecentlyActive verifys torrent(s) which have been recently active.
func (c *Controller) TorrentVerifyRecentlyActive() (err error) {
	if err = c.rpcCall("torrent-verify", &torrentActionRecentlyActiveParam{IDs: "recently-active"}, nil); err != nil {
		err = fmt.Errorf("'torrent-verify' rpc method failed: %v", err)
	}
	return
}

// TorrentReannounceIDs reannounces torrent(s) which id is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentReannounceIDs(ids []int64) (err error) {
	if err = c.rpcCall("torrent-reannounce", &torrentActionIDsParam{IDs: ids}, nil); err != nil {
		err = fmt.Errorf("'torrent-reannounce' rpc method failed: %v", err)
	}
	return
}

// TorrentReannounceHashes reannounces torrent(s) which hash is in the provided slice.
// Can be one, can be several, can be all (if slice is empty or nil).
func (c *Controller) TorrentReannounceHashes(hashes []string) (err error) {
	if err = c.rpcCall("torrent-reannounce", &torrentActionHashesParam{IDs: hashes}, nil); err != nil {
		err = fmt.Errorf("'torrent-reannounce' rpc method failed: %v", err)
	}
	return
}

// TorrentReannounceRecentlyActive reannounces torrent(s) which have been recently active.
func (c *Controller) TorrentReannounceRecentlyActive() (err error) {
	if err = c.rpcCall("torrent-reannounce", &torrentActionRecentlyActiveParam{IDs: "recently-active"}, nil); err != nil {
		err = fmt.Errorf("'torrent-reannounce' rpc method failed: %v", err)
	}
	return
}
