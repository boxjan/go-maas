package maas

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestReloadRegex(t *testing.T) {
	Convey("check all compile regex", t, func() {
		for k, rgx := range IgnoreTrailUrlSlashRegexPreCompile {
			tS := strings.ReplaceAll(strings.ReplaceAll(k, "{", ""), "}", "")
			So(rgx.FindStringIndex(tS), ShouldNotBeNil)
		}
	})
}

func BenchmarkTestRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, rgx := range IgnoreTrailUrlSlashRegexPreCompile {
			rgx.FindStringIndex(k)
		}
	}
}

func TestEnsureTrailingSlash(t *testing.T) {
	Convey("test should add / for url", t, func() {
		So(
			EnsureTrailingSlash("/MAAS/api/2.0/account"),
			ShouldEqual,
			"/MAAS/api/2.0/account/")

		So(
			EnsureTrailingSlash("/MAAS/api/2.0/nodes/{system_id}/blockdevices/{device_id}/partition/{id}"),
			ShouldEqual,
			"/MAAS/api/2.0/nodes/{system_id}/blockdevices/{device_id}/partition/{id}")

		So(
			EnsureTrailingSlash("/MAAS/api/2.0/nodes/{system_id}/blockdevices/{device_id}/partitions"),
			ShouldEqual,
			"/MAAS/api/2.0/nodes/{system_id}/blockdevices/{device_id}/partitions/")

		So(
			EnsureTrailingSlash("/MAAS/api/2.0/fannetworks/"),
			ShouldEqual,
			"/MAAS/api/2.0/fannetworks/")
	})
}

func TestJoinURLs(t *testing.T) {
	Convey("join urls ", t, func() {
		So(JoinURLs("http://example.com/base/", "/szot"),
			ShouldEqual, "http://example.com/base/szot")
	})
}
