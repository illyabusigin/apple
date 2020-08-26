# Apple

This repository contains tools that help make building Apple apps easier. 

[`entitlements`](https://pkg.go.dev/github.com/illyabusigin/apple/entitlements?tab=doc "API documentation") package
-------------------------------------------------------------------------------------------

The `entitlements` package providers a builder for declaring and generating your App.entitlements file for your Xcode project.

Features include:
- Functional approach
- Extensible
- String output
- Write to file

See it in action:

```go
package main

import (
	"fmt"
	"log"

	"github.com/illyabusigin/apple/entitlements"
)

func main() {
	e := entitlements.New()
	e.DataProtection.Complete()
	e.APS.Production()

	output, err := e.Build()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
```

**NOTE**: The `entitlements` package is incomplete, missing many entitlement options. More will be added over time. Pull-requets are appreciated.


[`plist`](https://pkg.go.dev/github.com/illyabusigin/apple/plist?tab=doc "API documentation") package
-------------------------------------------------------------------------------------------

The `plist` package provides methods for declaring and generating your Info.plist for your Xcode project. This package is built on the delightful [howett.net/plist](https://github.com/DHowett/go-plist) package.

Features include:
- Functional approach 
- Strongly typed
- Built-in validation with human-readale errors
- Extensible
- Output to a string or file

See it in action:

```go
package main

import (
	"fmt"
	"log"

	"github.com/illyabusigin/apple/plist"
)

func main() {
	plist := plist.New(plist.PlatformIOS)
	plist.Defaults()
	plist.DisplayName("BestApp")
	plist.BundleID("com.best.app")
	plist.AppTransportSecurity(func(s *AppTransportSecurity) {
		s.AllowArbitraryLoads(true)
	})
	plist.Orientations(func(o *Orientations) {
		o.Portrait()
		o.LandscapeLeft()
		o.LandscapeRight()
		o.UpsideDown()
	})
	plist.Privacy(func(p *Privacy) {
		p.Calendar("Let me use your calendar")
	})

	output, err := e.Build()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
```

**NOTE**: The `plist` package is still a work-in-progress and should not be considered ready for production use. Use at your own risk!