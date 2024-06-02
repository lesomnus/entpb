edition = "2023";

{{ if ne "" .PbPackage -}}
package {{ .PbPackage }};
{{- end }}

import "{{ print .PbDir "/public.proto" }}";

option features.field_presence = IMPLICIT;
{{ if ne "" .GoPackage -}}
option go_package = "{{ .GoPackage }}";
{{- end }}
