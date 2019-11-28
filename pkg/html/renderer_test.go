package html

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	internalTesting "github.com/fielmann-ag/version-monitor/pkg/internal/testing"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

func TestPageRenderer_ServeHTTP(t *testing.T) {
	type fields struct {
		monitor monitor.Monitor
	}
	tests := []struct {
		name     string
		fields   fields
		wantBody string
	}{
		{
			name: "render",
			fields: fields{
				monitor: internalTesting.NewMonitor([]monitor.Version{
					{Name: "x-test", Current: "1.2.3", Latest: "1.4.5"},
					{Name: "a/test", Current: "4.8.3", Latest: "8.0.0"},
					{Name: "l_test", Current: "9.2.4", Latest: "10.1.5"},
				}, internalTesting.Time, nil),
			},
			wantBody: `
<html>
<head>
</head>
<body>
	<div style="margin: 30px">Latest sync at <b>12 Dec 25 15:00 UTC</b></div>
	<table border="1" cellpadding="10px" cellspacing="0" style="border: 1px solid black;">
		<tr>
			<th>Name</th>
			<th>Currently deployed</th>
			<th>Latest version</th>
		</tr>
		<tr>
			<td valign="top">a/test</td>
			<td valign="top">4.8.3</td>
			<td valign="top">8.0.0</td>
		</tr>
		<tr>
			<td valign="top">l_test</td>
			<td valign="top">9.2.4</td>
			<td valign="top">10.1.5</td>
		</tr>
		<tr>
			<td valign="top">x-test</td>
			<td valign="top">1.2.3</td>
			<td valign="top">1.4.5</td>
		</tr>
	</table>
</body>
</html>
`,
		},
		{
			name: "error",
			fields: fields{
				monitor: internalTesting.NewMonitor([]monitor.Version{}, internalTesting.Time, errors.New("something bad happened")),
			},
			wantBody: `failed to render versions template: failed to fetch versions from monitor: something bad happened`,
		},
		{
			name: "empty",
			fields: fields{
				monitor: internalTesting.NewMonitor(nil, internalTesting.Time, nil),
			},
			wantBody: `
<html>
<head>
</head>
<body>
	<div style="margin: 30px">Latest sync at <b>12 Dec 25 15:00 UTC</b></div>
	<table border="1" cellpadding="10px" cellspacing="0" style="border: 1px solid black;">
		<tr>
			<th>Name</th>
			<th>Currently deployed</th>
			<th>Latest version</th>
		</tr>
	</table>
</body>
</html>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderer := NewPageRenderer(tt.fields.monitor, logging.NewTestLogger(t))
			rw := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			renderer.ServeHTTP(rw, r)
			body := rw.Body.String()

			if body != tt.wantBody {
				t.Log("want body:", tt.wantBody)
				t.Log("got body:", body)
				t.Errorf("expected to find body %v but found %v", tt.wantBody, body)
			}
		})
	}
}
