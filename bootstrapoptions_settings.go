package europi

// EnableVirtualFileSystem sets the enable flag for the virtual file system (backed by flash memory)
func EnableVirtualFileSystem(enable bool) BootstrapOption {
	return func(o *bootstrapConfig) error {
		o.vfsEnable = enable
		return nil
	}
}
