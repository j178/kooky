//go:build darwin && cgo

package chrome

import (
	"errors"
	"fmt"

	"github.com/keybase/go-keychain"
)

// getKeyringPassword retrieves the Chrome Safe Storage password,
// caching it for future calls.
func (s *CookieStore) getKeyringPassword(useSaved bool) ([]byte, error) {
	if s == nil {
		return nil, errors.New(`cookie store is nil`)
	}
	if useSaved && s.KeyringPasswordBytes != nil {
		return s.KeyringPasswordBytes, nil
	}

	kpmKey := `keychain_` + s.BrowserStr
	if useSaved {
		if kpw, ok := keyringPasswordMap.get(kpmKey); ok {
			return kpw, nil
		}
	}
	var service, account string
	switch s.BrowserStr {
	case `chrome`:
		service = `Chrome Safe Storage`
		account = `Chrome`
	case `chromium`:
		service = `Chromium Safe Storage`
		account = `Chromium`
	case `edge`:
		service = `Microsoft Edge Safe Storage`
		account = `Microsoft Edge`
	default:
		return nil, fmt.Errorf(`unknown browser: %s`, s.BrowserStr)
	}
	password, err := keychain.GetGenericPassword(service, account, "", "")
	if err != nil {
		return nil, fmt.Errorf("error reading '%s' keychain password: %w", service, err)
	}
	s.KeyringPasswordBytes = password
	keyringPasswordMap.set(kpmKey, password)

	return s.KeyringPasswordBytes, nil
}
