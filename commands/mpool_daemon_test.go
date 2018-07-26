package commands

import (
	"strings"
	"testing"

	"gx/ipfs/QmYVNvtQkeZ6AKSwDrjQTs432QtL6umrrK41EBq3cu7iSP/go-cid"

	"github.com/stretchr/testify/assert"

	"sync"

	tf "github.com/filecoin-project/go-filecoin/testhelpers/testfiles"
)

func TestMpool(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	wtf := tf.WalletFilePath()

	t.Run("return all messages", func(t *testing.T) {
		t.Parallel()
		d := NewDaemon(t, WalletFile(wtf), WalletAddr(testAddress1)).Start()
		defer d.ShutdownSuccess()

		d.RunSuccess("message", "send",
			"--from", testAddress1,
			"--value=10", testAddress3,
		)

		out := d.RunSuccess("mpool")
		c := strings.Trim(out.ReadStdout(), "\n")
		ci, err := cid.Decode(c)
		assert.NoError(err)
		assert.NotNil(ci)
	})

	t.Run("wait for enough messages", func(t *testing.T) {
		t.Parallel()
		d := NewDaemon(t, WalletFile(wtf), WalletAddr(testAddress1)).Start()
		defer d.ShutdownSuccess()

		wg := sync.WaitGroup{}
		wg.Add(1)

		complete := false
		go func() {
			out := d.RunSuccess("mpool", "--wait-for-count=3")
			complete = true
			wg.Done()
			c := strings.Split(strings.Trim(out.ReadStdout(), "\n"), "\n")
			assert.Equal(3, len(c))
		}()

		d.RunSuccess("message", "send",
			"--from", testAddress1,
			"--value=10", testAddress2,
		)

		assert.False(complete)

		d.RunSuccess("message", "send",
			"--from", testAddress1,
			"--value=10", testAddress2,
		)

		assert.False(complete)

		d.RunSuccess("message", "send",
			"--from", testAddress1,
			"--value=10", testAddress2,
		)

		wg.Wait()

		assert.True(complete)
	})
}
