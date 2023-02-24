//go:build !ios && !cgo

package chrome

import (
	"errors"
	"os/exec"
	"strings"
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
	out, err := exec.Command(`/usr/bin/security`, `find-generic-password`, `-s`, service,`-a`, account, `-w`, `Chrome`).Output()
	if err != nil {
		return nil, err
	}
	s.KeyringPasswordBytes = []byte(strings.TrimSpace(string(out)))
	keyringPasswordMap.set(kpmKey, s.KeyringPasswordBytes)

	return s.KeyringPasswordBytes, nil
}
