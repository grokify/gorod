module github.com/grokify/gorod

go 1.26.0

require (
	github.com/go-rod/rod v0.116.2
	github.com/grokify/mogo v0.74.5
	github.com/jessevdk/go-flags v1.6.1
	github.com/spf13/cobra v1.10.2
)

require (
	github.com/huandu/xstrings v1.5.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/ysmood/fetchup v0.5.3 // indirect
	github.com/ysmood/goob v0.4.0 // indirect
	github.com/ysmood/got v0.42.3 // indirect
	github.com/ysmood/gson v0.7.3 // indirect
	github.com/ysmood/leakless v0.9.0 // indirect
	golang.org/x/exp v0.0.0-20260508232706-74f9aab9d74a // indirect
	golang.org/x/net v0.54.0 // indirect
	golang.org/x/sys v0.44.0 // indirect
	golang.org/x/text v0.37.0 // indirect
)

// Pin fetchup to v0.2.3 for compatibility with go-rod/rod v0.116.2.
// The fetchup API changed in v0.3+ breaking rod's launcher package.
// Remove this replace directive when upgrading rod to a version that
// supports newer fetchup releases.
replace github.com/ysmood/fetchup => github.com/ysmood/fetchup v0.2.3
