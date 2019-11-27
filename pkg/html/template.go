package html

import (
	"html/template"
"time"


"github.com/fielmann-ag/ops-version-monitor/pkg/version"
)

var page = template.Must(template.New("page").Parse(pageTmpl))
var pageTmpl = `
<html>
<head>
</head>
	<h3>Latest sync at {{ .Date }}:</h3>
	
	<table border="1" cellpadding="10px" cellspacing="0" style="border: 1px solid black;">
		<tr>
			<th>Name</th>
			<th>Currently deployed</th>
			<th>Latest version</th>
		</tr>
	{{ range $version := .Versions }}
		<tr>
			<td valign="top">{{ $version.Name }}</td>
			<td valign="top">{{ $version.Current }}</td>
			<td valign="top">{{ $version.Latest }}</td>
		</tr>
	{{ end }}
	</table>
</body>
</html>
`

type pageParams struct {
	Versions []version.Version
	Date     time.Time
}
