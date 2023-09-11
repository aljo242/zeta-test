package types

const (
	// ModuleName defines the module name
	ModuleName = "sync"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sync"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	HeaderKey      = "Header/value/"
	HeaderCountKey = "Header/count/"
	AdminKey       = "Admin"
)
