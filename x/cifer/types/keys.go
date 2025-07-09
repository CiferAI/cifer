package types

const (
	// ModuleName defines the module name
	ModuleName = "cifer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cifer"
)

var (
	ParamsKey = []byte("p_cifer")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
