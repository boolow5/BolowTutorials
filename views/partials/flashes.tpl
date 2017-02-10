<div class="row column">
  {{if .error}}
  <div class="row column">
    <div class="alert callout">
      {{.error}}
    </div>
  </div>
  {{end}}
  {{if .warning}}
  <div class="row column">
    <div class="warning callout">
      {{.warning}}
    </div>
  </div>
  {{end}}
  {{if .notice}}
  <div class="row column">
    <div class="notice callout">
      {{.notice}}
    </div>
  </div>
  {{end}}
  {{if .success}}
  <div class="row column">
    <div class="success callout">
      {{.success}}
    </div>
  </div>
  {{end}}
</div>
