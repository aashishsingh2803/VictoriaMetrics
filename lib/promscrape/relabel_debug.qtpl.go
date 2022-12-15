// Code generated by qtc from "relabel_debug.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line lib/promscrape/relabel_debug.qtpl:1
package promscrape

//line lib/promscrape/relabel_debug.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promrelabel"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promutils"
)

//line lib/promscrape/relabel_debug.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/promscrape/relabel_debug.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/promscrape/relabel_debug.qtpl:8
func StreamMetricRelabelDebugSteps(qw422016 *qt422016.Writer, dss []promrelabel.DebugStep, metric, relabelConfigs string, err error) {
//line lib/promscrape/relabel_debug.qtpl:8
	qw422016.N().S(`<!DOCTYPE html><html lang="en"><head>`)
//line lib/promscrape/relabel_debug.qtpl:12
	streamcommonHeader(qw422016)
//line lib/promscrape/relabel_debug.qtpl:12
	qw422016.N().S(`<title>Metric relabel debug</title></head><body>`)
//line lib/promscrape/relabel_debug.qtpl:16
	streamnavbar(qw422016)
//line lib/promscrape/relabel_debug.qtpl:16
	qw422016.N().S(`<div class="container-fluid"><a href="https://docs.victoriametrics.com/relabeling.html" target="_blank">Relabeling docs</a>`)
//line lib/promscrape/relabel_debug.qtpl:18
	qw422016.N().S(` `)
//line lib/promscrape/relabel_debug.qtpl:18
	qw422016.N().S(`<a href="target-relabel-debug">Target relabel debug</a><br>`)
//line lib/promscrape/relabel_debug.qtpl:21
	if err != nil {
//line lib/promscrape/relabel_debug.qtpl:22
		streamerrorNotification(qw422016, err)
//line lib/promscrape/relabel_debug.qtpl:23
	}
//line lib/promscrape/relabel_debug.qtpl:23
	qw422016.N().S(`<div class="m-3"><form method="POST">`)
//line lib/promscrape/relabel_debug.qtpl:27
	streamrelabelDebugFormInputs(qw422016, metric, relabelConfigs)
//line lib/promscrape/relabel_debug.qtpl:27
	qw422016.N().S(`<input type="submit" value="Submit" class="btn btn-primary m-1" /></form></div><div class="row"><main class="col-12">`)
//line lib/promscrape/relabel_debug.qtpl:35
	streamrelabelDebugSteps(qw422016, dss)
//line lib/promscrape/relabel_debug.qtpl:35
	qw422016.N().S(`</main></div></div></body></html>`)
//line lib/promscrape/relabel_debug.qtpl:41
}

//line lib/promscrape/relabel_debug.qtpl:41
func WriteMetricRelabelDebugSteps(qq422016 qtio422016.Writer, dss []promrelabel.DebugStep, metric, relabelConfigs string, err error) {
//line lib/promscrape/relabel_debug.qtpl:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/relabel_debug.qtpl:41
	StreamMetricRelabelDebugSteps(qw422016, dss, metric, relabelConfigs, err)
//line lib/promscrape/relabel_debug.qtpl:41
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/relabel_debug.qtpl:41
}

//line lib/promscrape/relabel_debug.qtpl:41
func MetricRelabelDebugSteps(dss []promrelabel.DebugStep, metric, relabelConfigs string, err error) string {
//line lib/promscrape/relabel_debug.qtpl:41
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/relabel_debug.qtpl:41
	WriteMetricRelabelDebugSteps(qb422016, dss, metric, relabelConfigs, err)
//line lib/promscrape/relabel_debug.qtpl:41
	qs422016 := string(qb422016.B)
//line lib/promscrape/relabel_debug.qtpl:41
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/relabel_debug.qtpl:41
	return qs422016
//line lib/promscrape/relabel_debug.qtpl:41
}

//line lib/promscrape/relabel_debug.qtpl:43
func StreamTargetRelabelDebugSteps(qw422016 *qt422016.Writer, targetID string, dss []promrelabel.DebugStep, metric, relabelConfigs string, err error) {
//line lib/promscrape/relabel_debug.qtpl:43
	qw422016.N().S(`<!DOCTYPE html><html lang="en"><head>`)
//line lib/promscrape/relabel_debug.qtpl:47
	streamcommonHeader(qw422016)
//line lib/promscrape/relabel_debug.qtpl:47
	qw422016.N().S(`<title>Target relabel debug</title></head><body>`)
//line lib/promscrape/relabel_debug.qtpl:51
	streamnavbar(qw422016)
//line lib/promscrape/relabel_debug.qtpl:51
	qw422016.N().S(`<div class="container-fluid"><a href="https://docs.victoriametrics.com/relabeling.html" target="_blank">Relabeling docs</a>`)
//line lib/promscrape/relabel_debug.qtpl:53
	qw422016.N().S(` `)
//line lib/promscrape/relabel_debug.qtpl:53
	qw422016.N().S(`<a href="metric-relabel-debug">Metric relabel debug</a><br/>`)
//line lib/promscrape/relabel_debug.qtpl:56
	if err != nil {
//line lib/promscrape/relabel_debug.qtpl:57
		streamerrorNotification(qw422016, err)
//line lib/promscrape/relabel_debug.qtpl:58
	}
//line lib/promscrape/relabel_debug.qtpl:58
	qw422016.N().S(`<div class="m-3"><form method="POST">`)
//line lib/promscrape/relabel_debug.qtpl:62
	streamrelabelDebugFormInputs(qw422016, metric, relabelConfigs)
//line lib/promscrape/relabel_debug.qtpl:62
	qw422016.N().S(`<input type="hidden" name="id" value="`)
//line lib/promscrape/relabel_debug.qtpl:64
	qw422016.E().S(targetID)
//line lib/promscrape/relabel_debug.qtpl:64
	qw422016.N().S(`" /><input type="submit" value="Submit" class="btn btn-primary m-1" />`)
//line lib/promscrape/relabel_debug.qtpl:67
	if targetID != "" {
//line lib/promscrape/relabel_debug.qtpl:67
		qw422016.N().S(`<button type="button" onclick="location.href='target-relabel-debug?id=`)
//line lib/promscrape/relabel_debug.qtpl:68
		qw422016.E().S(targetID)
//line lib/promscrape/relabel_debug.qtpl:68
		qw422016.N().S(`'" class="btn btn-secondary m-1">Reset</button>`)
//line lib/promscrape/relabel_debug.qtpl:69
	}
//line lib/promscrape/relabel_debug.qtpl:69
	qw422016.N().S(`</form></div><div class="row"><main class="col-12">`)
//line lib/promscrape/relabel_debug.qtpl:75
	streamrelabelDebugSteps(qw422016, dss)
//line lib/promscrape/relabel_debug.qtpl:75
	qw422016.N().S(`</main></div></div></body></html>`)
//line lib/promscrape/relabel_debug.qtpl:81
}

//line lib/promscrape/relabel_debug.qtpl:81
func WriteTargetRelabelDebugSteps(qq422016 qtio422016.Writer, targetID string, dss []promrelabel.DebugStep, metric, relabelConfigs string, err error) {
//line lib/promscrape/relabel_debug.qtpl:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/relabel_debug.qtpl:81
	StreamTargetRelabelDebugSteps(qw422016, targetID, dss, metric, relabelConfigs, err)
//line lib/promscrape/relabel_debug.qtpl:81
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/relabel_debug.qtpl:81
}

//line lib/promscrape/relabel_debug.qtpl:81
func TargetRelabelDebugSteps(targetID string, dss []promrelabel.DebugStep, metric, relabelConfigs string, err error) string {
//line lib/promscrape/relabel_debug.qtpl:81
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/relabel_debug.qtpl:81
	WriteTargetRelabelDebugSteps(qb422016, targetID, dss, metric, relabelConfigs, err)
//line lib/promscrape/relabel_debug.qtpl:81
	qs422016 := string(qb422016.B)
//line lib/promscrape/relabel_debug.qtpl:81
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/relabel_debug.qtpl:81
	return qs422016
//line lib/promscrape/relabel_debug.qtpl:81
}

//line lib/promscrape/relabel_debug.qtpl:83
func streamrelabelDebugFormInputs(qw422016 *qt422016.Writer, metric, relabelConfigs string) {
//line lib/promscrape/relabel_debug.qtpl:83
	qw422016.N().S(`<div>Relabel configs:<br/><textarea name="relabel_configs" style="width: 100%; height: 15em" class="m-1">`)
//line lib/promscrape/relabel_debug.qtpl:86
	qw422016.E().S(relabelConfigs)
//line lib/promscrape/relabel_debug.qtpl:86
	qw422016.N().S(`</textarea></div><div>Labels:<br/><textarea name="metric" style="width: 100%; height: 5em" class="m-1">`)
//line lib/promscrape/relabel_debug.qtpl:91
	qw422016.E().S(metric)
//line lib/promscrape/relabel_debug.qtpl:91
	qw422016.N().S(`</textarea></div>`)
//line lib/promscrape/relabel_debug.qtpl:93
}

//line lib/promscrape/relabel_debug.qtpl:93
func writerelabelDebugFormInputs(qq422016 qtio422016.Writer, metric, relabelConfigs string) {
//line lib/promscrape/relabel_debug.qtpl:93
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/relabel_debug.qtpl:93
	streamrelabelDebugFormInputs(qw422016, metric, relabelConfigs)
//line lib/promscrape/relabel_debug.qtpl:93
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/relabel_debug.qtpl:93
}

//line lib/promscrape/relabel_debug.qtpl:93
func relabelDebugFormInputs(metric, relabelConfigs string) string {
//line lib/promscrape/relabel_debug.qtpl:93
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/relabel_debug.qtpl:93
	writerelabelDebugFormInputs(qb422016, metric, relabelConfigs)
//line lib/promscrape/relabel_debug.qtpl:93
	qs422016 := string(qb422016.B)
//line lib/promscrape/relabel_debug.qtpl:93
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/relabel_debug.qtpl:93
	return qs422016
//line lib/promscrape/relabel_debug.qtpl:93
}

//line lib/promscrape/relabel_debug.qtpl:95
func streamrelabelDebugSteps(qw422016 *qt422016.Writer, dss []promrelabel.DebugStep) {
//line lib/promscrape/relabel_debug.qtpl:96
	if len(dss) > 0 {
//line lib/promscrape/relabel_debug.qtpl:96
		qw422016.N().S(`<div class="m-3"><b>Original labels:</b> <samp>`)
//line lib/promscrape/relabel_debug.qtpl:98
		streammustFormatLabels(qw422016, dss[0].In)
//line lib/promscrape/relabel_debug.qtpl:98
		qw422016.N().S(`</samp></div>`)
//line lib/promscrape/relabel_debug.qtpl:100
	}
//line lib/promscrape/relabel_debug.qtpl:100
	qw422016.N().S(`<table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col" style="width: 5%">Step</th><th scope="col" style="width: 25%">Relabeling Rule</th><th scope="col" style="width: 35%">Input Labels</th><th scope="col" stile="width: 35%">Output labels</a></tr></thead><tbody>`)
//line lib/promscrape/relabel_debug.qtpl:111
	for i, ds := range dss {
//line lib/promscrape/relabel_debug.qtpl:113
		inLabels := promutils.MustNewLabelsFromString(ds.In)
		outLabels := promutils.MustNewLabelsFromString(ds.Out)
		changedLabels := getChangedLabelNames(inLabels, outLabels)

//line lib/promscrape/relabel_debug.qtpl:116
		qw422016.N().S(`<tr><td>`)
//line lib/promscrape/relabel_debug.qtpl:118
		qw422016.N().D(i)
//line lib/promscrape/relabel_debug.qtpl:118
		qw422016.N().S(`</td><td><b><pre class="m-2">`)
//line lib/promscrape/relabel_debug.qtpl:119
		qw422016.E().S(ds.Rule)
//line lib/promscrape/relabel_debug.qtpl:119
		qw422016.N().S(`</pre></b></td><td><div class="m-2" style="font-size: 0.9em" title="deleted and updated labels highlighted in red">`)
//line lib/promscrape/relabel_debug.qtpl:122
		streamlabelsWithHighlight(qw422016, inLabels, changedLabels, "red")
//line lib/promscrape/relabel_debug.qtpl:122
		qw422016.N().S(`</div></td><td><div class="m-2" style="font-size: 0.9em" title="added and updated labels highlighted in blue">`)
//line lib/promscrape/relabel_debug.qtpl:127
		streamlabelsWithHighlight(qw422016, outLabels, changedLabels, "blue")
//line lib/promscrape/relabel_debug.qtpl:127
		qw422016.N().S(`</div></td></tr>`)
//line lib/promscrape/relabel_debug.qtpl:131
	}
//line lib/promscrape/relabel_debug.qtpl:131
	qw422016.N().S(`</tbody></table>`)
//line lib/promscrape/relabel_debug.qtpl:134
	if len(dss) > 0 {
//line lib/promscrape/relabel_debug.qtpl:134
		qw422016.N().S(`<div class="m-3"><b>Resulting labels:</b> <samp>`)
//line lib/promscrape/relabel_debug.qtpl:136
		streammustFormatLabels(qw422016, dss[len(dss)-1].Out)
//line lib/promscrape/relabel_debug.qtpl:136
		qw422016.N().S(`</samp></div>`)
//line lib/promscrape/relabel_debug.qtpl:138
	}
//line lib/promscrape/relabel_debug.qtpl:139
}

//line lib/promscrape/relabel_debug.qtpl:139
func writerelabelDebugSteps(qq422016 qtio422016.Writer, dss []promrelabel.DebugStep) {
//line lib/promscrape/relabel_debug.qtpl:139
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/relabel_debug.qtpl:139
	streamrelabelDebugSteps(qw422016, dss)
//line lib/promscrape/relabel_debug.qtpl:139
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/relabel_debug.qtpl:139
}

//line lib/promscrape/relabel_debug.qtpl:139
func relabelDebugSteps(dss []promrelabel.DebugStep) string {
//line lib/promscrape/relabel_debug.qtpl:139
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/relabel_debug.qtpl:139
	writerelabelDebugSteps(qb422016, dss)
//line lib/promscrape/relabel_debug.qtpl:139
	qs422016 := string(qb422016.B)
//line lib/promscrape/relabel_debug.qtpl:139
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/relabel_debug.qtpl:139
	return qs422016
//line lib/promscrape/relabel_debug.qtpl:139
}

//line lib/promscrape/relabel_debug.qtpl:141
func streamlabelsWithHighlight(qw422016 *qt422016.Writer, labels *promutils.Labels, highlight map[string]struct{}, color string) {
//line lib/promscrape/relabel_debug.qtpl:143
	labelsList := labels.GetLabels()
	metricName := ""
	for i, label := range labelsList {
		if label.Name == "__name__" {
			metricName = label.Value
			labelsList = append(labelsList[:i], labelsList[i+1:]...)
			break
		}
	}

//line lib/promscrape/relabel_debug.qtpl:153
	if metricName != "" {
//line lib/promscrape/relabel_debug.qtpl:154
		if _, ok := highlight["__name__"]; ok {
//line lib/promscrape/relabel_debug.qtpl:154
			qw422016.N().S(`<span style="font-weight:bold;color:`)
//line lib/promscrape/relabel_debug.qtpl:155
			qw422016.E().S(color)
//line lib/promscrape/relabel_debug.qtpl:155
			qw422016.N().S(`">`)
//line lib/promscrape/relabel_debug.qtpl:155
			qw422016.E().S(metricName)
//line lib/promscrape/relabel_debug.qtpl:155
			qw422016.N().S(`</span>`)
//line lib/promscrape/relabel_debug.qtpl:156
		} else {
//line lib/promscrape/relabel_debug.qtpl:157
			qw422016.E().S(metricName)
//line lib/promscrape/relabel_debug.qtpl:158
		}
//line lib/promscrape/relabel_debug.qtpl:159
		if len(labelsList) == 0 {
//line lib/promscrape/relabel_debug.qtpl:159
			return
//line lib/promscrape/relabel_debug.qtpl:159
		}
//line lib/promscrape/relabel_debug.qtpl:160
	}
//line lib/promscrape/relabel_debug.qtpl:160
	qw422016.N().S(`{`)
//line lib/promscrape/relabel_debug.qtpl:162
	for i, label := range labelsList {
//line lib/promscrape/relabel_debug.qtpl:163
		if _, ok := highlight[label.Name]; ok {
//line lib/promscrape/relabel_debug.qtpl:163
			qw422016.N().S(`<span style="font-weight:bold;color:`)
//line lib/promscrape/relabel_debug.qtpl:164
			qw422016.E().S(color)
//line lib/promscrape/relabel_debug.qtpl:164
			qw422016.N().S(`">`)
//line lib/promscrape/relabel_debug.qtpl:164
			qw422016.E().S(label.Name)
//line lib/promscrape/relabel_debug.qtpl:164
			qw422016.N().S(`=`)
//line lib/promscrape/relabel_debug.qtpl:164
			qw422016.E().Q(label.Value)
//line lib/promscrape/relabel_debug.qtpl:164
			qw422016.N().S(`</span>`)
//line lib/promscrape/relabel_debug.qtpl:165
		} else {
//line lib/promscrape/relabel_debug.qtpl:166
			qw422016.E().S(label.Name)
//line lib/promscrape/relabel_debug.qtpl:166
			qw422016.N().S(`=`)
//line lib/promscrape/relabel_debug.qtpl:166
			qw422016.E().Q(label.Value)
//line lib/promscrape/relabel_debug.qtpl:167
		}
//line lib/promscrape/relabel_debug.qtpl:168
		if i < len(labelsList)-1 {
//line lib/promscrape/relabel_debug.qtpl:168
			qw422016.N().S(`,`)
//line lib/promscrape/relabel_debug.qtpl:168
			qw422016.N().S(` `)
//line lib/promscrape/relabel_debug.qtpl:168
		}
//line lib/promscrape/relabel_debug.qtpl:169
	}
//line lib/promscrape/relabel_debug.qtpl:169
	qw422016.N().S(`}`)
//line lib/promscrape/relabel_debug.qtpl:171
}

//line lib/promscrape/relabel_debug.qtpl:171
func writelabelsWithHighlight(qq422016 qtio422016.Writer, labels *promutils.Labels, highlight map[string]struct{}, color string) {
//line lib/promscrape/relabel_debug.qtpl:171
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/relabel_debug.qtpl:171
	streamlabelsWithHighlight(qw422016, labels, highlight, color)
//line lib/promscrape/relabel_debug.qtpl:171
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/relabel_debug.qtpl:171
}

//line lib/promscrape/relabel_debug.qtpl:171
func labelsWithHighlight(labels *promutils.Labels, highlight map[string]struct{}, color string) string {
//line lib/promscrape/relabel_debug.qtpl:171
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/relabel_debug.qtpl:171
	writelabelsWithHighlight(qb422016, labels, highlight, color)
//line lib/promscrape/relabel_debug.qtpl:171
	qs422016 := string(qb422016.B)
//line lib/promscrape/relabel_debug.qtpl:171
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/relabel_debug.qtpl:171
	return qs422016
//line lib/promscrape/relabel_debug.qtpl:171
}

//line lib/promscrape/relabel_debug.qtpl:173
func streammustFormatLabels(qw422016 *qt422016.Writer, s string) {
//line lib/promscrape/relabel_debug.qtpl:174
	labels := promutils.MustNewLabelsFromString(s)

//line lib/promscrape/relabel_debug.qtpl:175
	streamlabelsWithHighlight(qw422016, labels, nil, "")
//line lib/promscrape/relabel_debug.qtpl:176
}

//line lib/promscrape/relabel_debug.qtpl:176
func writemustFormatLabels(qq422016 qtio422016.Writer, s string) {
//line lib/promscrape/relabel_debug.qtpl:176
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/relabel_debug.qtpl:176
	streammustFormatLabels(qw422016, s)
//line lib/promscrape/relabel_debug.qtpl:176
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/relabel_debug.qtpl:176
}

//line lib/promscrape/relabel_debug.qtpl:176
func mustFormatLabels(s string) string {
//line lib/promscrape/relabel_debug.qtpl:176
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/relabel_debug.qtpl:176
	writemustFormatLabels(qb422016, s)
//line lib/promscrape/relabel_debug.qtpl:176
	qs422016 := string(qb422016.B)
//line lib/promscrape/relabel_debug.qtpl:176
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/relabel_debug.qtpl:176
	return qs422016
//line lib/promscrape/relabel_debug.qtpl:176
}
