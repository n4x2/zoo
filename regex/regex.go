// Copyright 2015 Dean Karn. All rights reserved.
// Use of this source code is governed by a MIT license that can be found
// at https://github.com/go-playground/validator/blob/master/LICENSE

// Package regex provides regular expressions.
package regex

import "regexp"

const (
	alphaPattern                 = "^[a-zA-Z]+$"
	alphaDashPattern             = "^[a-zA-Z_-]+$"
	alphaNumericPattern          = "^[a-zA-Z0-9]+$"
	alphaUnicodePattern          = "^[\\p{L}]+$"
	alphaUnicodeNumericPattern   = "^[\\p{L}\\p{N}]+$"
	numericPattern               = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberPattern                = "^[0-9]+$"
	hexadecimalPattern           = "^(0[xX])?[0-9a-fA-F]+$"
	hexColorPattern              = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$"
	rgbPattern                   = "^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$"
	rgbaPattern                  = "^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	hslPattern                   = "^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$"
	hslaPattern                  = "^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	emailPattern                 = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	e164Pattern                  = "^\\+[1-9]?[0-9]{7,14}$"
	base64Pattern                = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	base64URLPattern             = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	base64RawURLPattern          = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2,4})$"
	isbn10Pattern                = "^(?:[0-9]{9}X|[0-9]{10})$"
	isbn13Pattern                = "^(?:(?:97(?:8|9))[0-9]{10})$"
	uuid3Pattern                 = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uuid4Pattern                 = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uuid5Pattern                 = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uuidPattern                  = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uuid3RFC4122Pattern          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uuid4RFC4122Pattern          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uuid5RFC4122Pattern          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uuidRFC4122Pattern           = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	ulidPattern                  = "^[A-HJKMNP-TV-Z0-9]{26}$"
	md4Pattern                   = "^[0-9a-f]{32}$"
	md5Pattern                   = "^[0-9a-f]{32}$"
	sha256Pattern                = "^[0-9a-f]{64}$"
	sha384Pattern                = "^[0-9a-f]{96}$"
	sha512Pattern                = "^[0-9a-f]{128}$"
	ripemd128Pattern             = "^[0-9a-f]{32}$"
	ripemd160Pattern             = "^[0-9a-f]{40}$"
	tiger128Pattern              = "^[0-9a-f]{32}$"
	tiger160Pattern              = "^[0-9a-f]{40}$"
	tiger192Pattern              = "^[0-9a-f]{48}$"
	asciiPattern                 = "^[\x00-\x7F]*$"
	printableASCIIPattern        = "^[\x20-\x7E]*$"
	multibytePattern             = "[^\x00-\x7F]"
	dataURIPattern               = `^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`
	latitudePattern              = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longitudePattern             = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	ssnPattern                   = `^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`
	hostnamePatternRFC952        = `^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`                                                                   // https://tools.ietf.org/html/rfc952
	hostnamePatternRFC1123       = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?$`                                 // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	fqdnPatternRFC1123           = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$` // same as hostnameStringRFC1123 but must contain a non numerical TLD (possibly ending with '.')
	btcAddressPattern            = `^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`                                                                             // bitcoin address
	btcAddressUpperPatternBech32 = `^BC1[02-9AC-HJ-NP-Z]{7,76}$`                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	btcAddressLowerPatternBech32 = `^bc1[02-9ac-hj-np-z]{7,76}$`                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	ethAddressPattern            = `^0x[0-9a-fA-F]{40}$`
	ethAddressUpperPattern       = `^0x[0-9A-F]{40}$`
	ethAddressLowerPattern       = `^0x[0-9a-f]{40}$`
	urlEncodedPattern            = `^(?:[^%]|%[0-9A-Fa-f]{2})*$`
	htmlEncodedPattern           = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	htmlPattern                  = `<[/]?([a-zA-Z]+).*?>`
	jwtPattern                   = "^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$"
	splitParamsPattern           = `'[^']*'|\S+`
	bicPattern                   = `^[A-Za-z]{6}[A-Za-z0-9]{2}([A-Za-z0-9]{3})?$`
	semverPattern                = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$` // numbered capture groups https://semver.org/
	dnsPatternRFC1035Label       = "^[a-z]([-a-z0-9]*[a-z0-9]){0,62}$"
	cvePattern                   = `^CVE-(1999|2\d{3})-(0[^0]\d{2}|0\d[^0]\d{1}|0\d{2}[^0]|[1-9]{1}\d{3,})$` // CVE Format Id https://cve.mitre.org/cve/identifiers/syntaxchange.html
	mongodbPattern               = "^[a-f\\d]{24}$"
	cronPattern                  = `(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})`
	spicedbIDPattern             = `^(([a-zA-Z0-9/_|\-=+]{1,})|\*)$`
	spicedbPermissionPattern     = "^([a-z][a-z0-9_]{1,62}[a-z0-9])?$"
	spicedbTypePattern           = "^([a-z][a-z0-9_]{1,61}[a-z0-9]/)?[a-z][a-z0-9_]{1,62}[a-z0-9]$"
)

var (
	Alpha                 = regexp.MustCompile(alphaPattern)
	AlphaDash             = regexp.MustCompile(alphaDashPattern)
	AlphaNumeric          = regexp.MustCompile(alphaNumericPattern)
	AlphaUnicode          = regexp.MustCompile(alphaUnicodePattern)
	AlphaUnicodeNumeric   = regexp.MustCompile(alphaUnicodeNumericPattern)
	Numeric               = regexp.MustCompile(numericPattern)
	Number                = regexp.MustCompile(numberPattern)
	Hexadecimal           = regexp.MustCompile(hexadecimalPattern)
	HexColor              = regexp.MustCompile(hexColorPattern)
	RGB                   = regexp.MustCompile(rgbPattern)
	RGBA                  = regexp.MustCompile(rgbaPattern)
	HSL                   = regexp.MustCompile(hslPattern)
	HSLA                  = regexp.MustCompile(hslaPattern)
	E164                  = regexp.MustCompile(e164Pattern)
	Email                 = regexp.MustCompile(emailPattern)
	Base64                = regexp.MustCompile(base64Pattern)
	Base64URL             = regexp.MustCompile(base64URLPattern)
	Base64RawURL          = regexp.MustCompile(base64RawURLPattern)
	ISBN10                = regexp.MustCompile(isbn10Pattern)
	ISBN13                = regexp.MustCompile(isbn13Pattern)
	UUID3                 = regexp.MustCompile(uuid3Pattern)
	UUID4                 = regexp.MustCompile(uuid4Pattern)
	UUID5                 = regexp.MustCompile(uuid5Pattern)
	UUID                  = regexp.MustCompile(uuidPattern)
	UUID3RFC4122          = regexp.MustCompile(uuid3RFC4122Pattern)
	UUID4RFC4122          = regexp.MustCompile(uuid4RFC4122Pattern)
	UUID5RFC4122          = regexp.MustCompile(uuid5RFC4122Pattern)
	UUIDRFC4122           = regexp.MustCompile(uuidRFC4122Pattern)
	ULID                  = regexp.MustCompile(ulidPattern)
	MD4                   = regexp.MustCompile(md4Pattern)
	MD5                   = regexp.MustCompile(md5Pattern)
	SHA256                = regexp.MustCompile(sha256Pattern)
	SHA384                = regexp.MustCompile(sha384Pattern)
	SHA512                = regexp.MustCompile(sha512Pattern)
	RipeMD128             = regexp.MustCompile(ripemd128Pattern)
	RipeMD160             = regexp.MustCompile(ripemd160Pattern)
	Tiger128              = regexp.MustCompile(tiger128Pattern)
	Tiger160              = regexp.MustCompile(tiger160Pattern)
	Tiger192              = regexp.MustCompile(tiger192Pattern)
	ASCII                 = regexp.MustCompile(asciiPattern)
	PrintableASCII        = regexp.MustCompile(printableASCIIPattern)
	Multibyte             = regexp.MustCompile(multibytePattern)
	DataURI               = regexp.MustCompile(dataURIPattern)
	Latitude              = regexp.MustCompile(latitudePattern)
	Longitude             = regexp.MustCompile(longitudePattern)
	SSN                   = regexp.MustCompile(ssnPattern)
	HostnameRFC952        = regexp.MustCompile(hostnamePatternRFC952)
	HostnameRFC1123       = regexp.MustCompile(hostnamePatternRFC1123)
	FQDNRFC1123           = regexp.MustCompile(fqdnPatternRFC1123)
	BTCAddress            = regexp.MustCompile(btcAddressPattern)
	BTCUpperAddressBech32 = regexp.MustCompile(btcAddressUpperPatternBech32)
	BTCLowerAddressBech32 = regexp.MustCompile(btcAddressLowerPatternBech32)
	ETHAddress            = regexp.MustCompile(ethAddressPattern)
	URLEncoded            = regexp.MustCompile(urlEncodedPattern)
	HTMLEncoded           = regexp.MustCompile(htmlEncodedPattern)
	HTML                  = regexp.MustCompile(htmlPattern)
	JWT                   = regexp.MustCompile(jwtPattern)
	SplitParams           = regexp.MustCompile(splitParamsPattern)
	BIC                   = regexp.MustCompile(bicPattern)
	Semver                = regexp.MustCompile(semverPattern)
	DNSRFC1035Label       = regexp.MustCompile(dnsPatternRFC1035Label)
	CVE                   = regexp.MustCompile(cvePattern)
	MongoDB               = regexp.MustCompile(mongodbPattern)
	Cron                  = regexp.MustCompile(cronPattern)
	SpicedbID             = regexp.MustCompile(spicedbIDPattern)
	SpicedbPermission     = regexp.MustCompile(spicedbPermissionPattern)
	SpicedbType           = regexp.MustCompile(spicedbTypePattern)
)
