package share

import (
	"os"
)

const (
	Pakfile   = "Pakfile"
	Paklock   = "Pakfile.lock"
	Pakbranch = "pak"
	Paktag    = "_pak_latest_"
)

var Gopath = os.Getenv("GOPATH")

type PakInfo struct {
	Packages []string
}

type PaklockInfo map[string]string

type PakOption struct {
	PakMeter       []string // used for containing spcified packages
	UsePakfileLock bool
	Force          bool
}

type GetOption struct {
	Force    bool
	Checksum string
}

// Notes:
// Containing branch named pak does not mean that pkg is managed by pak.
// Containing tag named _pak_latest_ means this pkg is managed by pak, but
// still can't make sure the pkg is on the pak branch or it's status is wanted
// by Pakfile or Pakfile.lock.

type PkgProxy interface {
	Fetch() error
	NewBranch(string) error
	NewTag(string, string) error
	RemoveTag(string) error
	Pak(string) (string, error)
	Unpak() error
	// Report() error

	IsClean() (bool, error)
	ContainsRemoteBranch() (bool, error)
	ContainsPakbranch() (bool, error)
	ContainsPaktag() (bool, error)
	GetChecksum(string) (string, error)
	GetHeadChecksum() (string, error)
	GetHeadRefName() (string, error)
	GetPaktagRef() string
	GetPakbranchRef() string
	GetRemoteBranch() string
}

type PkgProxyBuilder struct {
	IsTracking func(name string) (bool, error)
	NewVCS func(name, remote, branch string) PkgProxy
}

var PkgProxyList = []PkgProxyBuilder{}
func RegisterPkgProxy(newBuilder PkgProxyBuilder) {
	PkgProxyList = append(PkgProxyList, newBuilder)
}