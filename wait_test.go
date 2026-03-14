package gorod

/*
import (
	"testing"
	"time"

	"github.com/go-rod/rod"
	"github.com/stretchr/testify/require"
)

func TestExampleLogin(t *testing.T) {
	page := rod.New().MustConnect().MustPage("https://example.com/login")
	defer page.Close()

	require.NoError(t, WaitVisible(page, "#username", 5*time.Second))
	page.MustElement("#username").MustInput("john")

	require.NoError(t, WaitVisible(page, "#password", 5*time.Second))
	page.MustElement("#password").MustInput("secret")

	require.NoError(t, WaitClick(page, "button[type=submit]", 5*time.Second))
	page.MustWaitLoad()

	require.Contains(t, page.MustInfo().URL, "/dashboard")
}
*/
