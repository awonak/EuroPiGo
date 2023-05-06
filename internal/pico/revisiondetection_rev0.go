//go:build pico && (revision0 || europiproto || europiprototype)
// +build pico
// +build revision0 europiproto europiprototype

package pico

func init() {
	rev1AsRev0 = true
}
