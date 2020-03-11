package clitest

import (
	"github.com/cosmos/cosmos-sdk/tests"
	"github.com/desmos-labs/desmos/x/reactions"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDesmosCLIPostsReactions(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Create a post
	success, _, sterr := f.TxPostsCreate("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"message", true, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// __________________________________________________________________________________
	// add-reaction

	// Add a reaction
	success, _, sterr = f.TxPostsAddReaction(1, "👍", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the reaction is added
	storedPost := f.QueryPost(1)
	require.Len(t, storedPost.PostReactions, 1)
	require.Equal(t, storedPost.PostReactions[0], reactions.NewPostReaction("👍", fooAddr))

	// Test --dry-run
	success, _, _ = f.TxPostsAddReaction(1, "😊", fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsAddReaction(1, "👎", fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPost = f.QueryPost(1)
	require.Len(t, storedPost.PostReactions, 1)

	// __________________________________________________________________________________
	// remove-reaction

	// Remove a reaction
	success, _, sterr = f.TxPostsRemoveReaction(1, "👍", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the reaction has been removed
	storedPost = f.QueryPost(1)
	require.Empty(t, storedPost.PostReactions)

	// Test --dry-run
	success, _, _ = f.TxPostsRemoveReaction(1, "😊", fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr = f.TxPostsRemoveReaction(1, "👎", fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg = unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPost = f.QueryPost(1)
	require.Empty(t, storedPost.PostReactions)

	f.Cleanup()
}
