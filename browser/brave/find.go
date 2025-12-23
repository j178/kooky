package brave

import (
	"github.com/j178/kooky"
	"github.com/j178/kooky/internal/chrome"
	"github.com/j178/kooky/internal/chrome/find"
	"github.com/j178/kooky/internal/cookies"
)

type braveFinder struct{}

var _ kooky.CookieStoreFinder = (*braveFinder)(nil)

func init() {
	kooky.RegisterFinder(`brave`, &braveFinder{})
}

func (f *braveFinder) FindCookieStores() ([]kooky.CookieStore, error) {
	files, err := find.FindBraveCookieStoreFiles()
	if err != nil {
		return nil, err
	}

	var ret []kooky.CookieStore
	for _, file := range files {
		ret = append(
			ret,
			&cookies.CookieJar{
				CookieStore: &chrome.CookieStore{
					DefaultCookieStore: cookies.DefaultCookieStore{
						BrowserStr:           file.Browser,
						ProfileStr:           file.Profile,
						OSStr:                file.OS,
						IsDefaultProfileBool: file.IsDefaultProfile,
						FileNameStr:          file.Path,
					},
				},
			},
		)
	}
	return ret, nil
}
