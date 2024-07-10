package awesomesearchqueries_test

import (
	"testing"

	awesomesearchqueries "github.com/projectdiscovery/awesome-search-queries"
	"github.com/stretchr/testify/require"
)

// === this test should not break ===
// === cvemap dependent on this ===

func TestLoadShodanQueries(t *testing.T) {
	queries, err := awesomesearchqueries.LoadShodanQueries("./QUERIES.json")
	require.NoError(t, err)
	require.Greater(t, len(queries), 1)
}
