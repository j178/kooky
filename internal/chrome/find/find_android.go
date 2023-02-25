//go:build android

package find

import "errors"

func chromeRoots() ([]string, error) {
	// https://chromium.googlesource.com/chromium/src.git/+/62.0.3202.58/docs/user_data_dir.md#android
	var ret = []string{
		`/data/user/0/com.android.chrome/app_chrome`, // TODO check
	}
	return ret, nil
}

func chromiumRoots() ([]string, error) {
	return ret, errors.New(`not implemented`)
}

func edgeRoots() ([]string, error) {
	return nil, errors.New(`not implemented`)
}
