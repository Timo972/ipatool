package unicorn

import "fmt"

const (
	ucCtlTCGBufferSize = 13
	ucCtlIOWrite       = 1
	ucCtlArgumentCount = 1
	tcgBufferSize      = 64 << 20
)

func configureEngine(engine *Engine) error {
	// Unicorn's default translation buffer is far larger than the guest code
	// needs. Windows' VirtualAlloc compatibility hook commits it eagerly, and
	// elsewhere it is reserved per engine, which memory-limited hosts and
	// containers can refuse ("Could not allocate dynamic translator buffer").
	// Unicorn flushes the buffer when it fills, so the bound costs no
	// correctness.
	control := uint32(ucCtlTCGBufferSize |
		ucCtlArgumentCount<<26 |
		ucCtlIOWrite<<30)
	if err := engine.err(engine.api.ctl(engine.handle, control, tcgBufferSize)); err != nil {
		return fmt.Errorf("configure Unicorn translation buffer: %w", err)
	}

	return nil
}
