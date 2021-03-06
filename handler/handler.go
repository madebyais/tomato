/* GENERATED FILE - DO NOT EDIT */
/* Rebuild from the cmd/gen/main.go tool */
package handler

import (
	"github.com/DATA-DOG/godog"
	"github.com/alileza/tomato/resource"
)

type Handler struct {
	resource resource.Manager
}

func New(r resource.Manager) func(s *godog.Suite) {
	h := &Handler{r}
	return func(s *godog.Suite) {
		s.Step(`^"([^"]*)" send request to "([^"]*)"$`, h.sendRequestTo)
		s.Step(`^"([^"]*)" send request to "([^"]*)" with body$`, h.sendRequestToWithBody)
		s.Step(`^"([^"]*)" response code should be (\d+)$`, h.responseCodeShouldBe)
		s.Step(`^"([^"]*)" response body should be$`, h.responseBodyShouldBe)
		s.Step(`^set "([^"]*)" response code to (\d+) and response body$`, h.setResponseCodeToAndResponseBody)
		s.Step(`^set "([^"]*)" with path "([^"]*)" response code to (\d+) and response body$`, h.setWithPathResponseCodeToAndResponseBody)
		s.Step(`^set "([^"]*)" table "([^"]*)" list of content$`, h.setTableListOfContent)
		s.Step(`^"([^"]*)" table "([^"]*)" should look like$`, h.tableShouldLookLike)
		s.Step(`^set "([^"]*)" table "([^"]*)" to empty$`, h.setTableToEmpty)
		s.Step(`^publish message to "([^"]*)" target "([^"]*)" with payload$`, h.publishMessageToTargetWithPayload)
		s.Step(`^listen message from "([^"]*)" target "([^"]*)"$`, h.listenMessageFromTarget)
		s.Step(`^message from "([^"]*)" target "([^"]*)" count should be (\d+)$`, h.messageFromTargetCountShouldBe)
		s.Step(`^message from "([^"]*)" target "([^"]*)" should look like$`, h.messageFromTargetShouldLookLike)
	}
}
