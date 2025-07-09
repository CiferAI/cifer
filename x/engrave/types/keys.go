package types

const (
	// ModuleName defines the module name
	ModuleName = "engrave"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_engrave"
)

var (
	ParamsKey = []byte("p_engrave")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
