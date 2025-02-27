/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package golang

import (
	"time"

	"github.com/nuclio/nuclio-sdk-go"
)

// this is used for running a standalone processor during development
func builtInHandler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	context.Logger.InfoWith("Got event",
		"URL", event.GetURL(),
		"Path", event.GetPath(),
		"Type", event.GetType(),
		"TypeVersion", event.GetTypeVersion(),
		"Version", event.GetVersion(),
		"Source", event.GetTriggerInfo().GetKind(),
		"ID", event.GetID(),
		"Time", event.GetTimestamp().String(),
		"Headers", event.GetHeaders(),
		"ContentType", event.GetContentType(),
		"Body", string(event.GetBody()))

	time.Sleep(3 * time.Second)

	return "Built in handler called", nil
}
