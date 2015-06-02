package hiro

// DefaultTemplate is the default iglo template used by hiro
var DefaultTemplate = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{.Name}}</title>
		<meta name="description" content="{{trim .Description "\n"}}">
		<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css">
		<style>
			tt, pre, code { font-family: Consolas, "Liberation Mono", Courier, monospace; background-color: transparent !important; }
			pre.prettyprint { border: 0px !important; background-color: #fff; margin-bottom: -0.5em; }
			.panel-heading h2 { margin-top: 0.5em; }
			.bg-default { background-color: #F8F8F8; }
			.snippet { background: #F8F8; list-style: none; display: none; }
			.snippet-toggle { margin-top: -0.3em; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="row">
				<div class="col-md-12">
					<div class="page-header">
						<h1>{{.Name}}</h1>
						<h2><small>{{.Description | markdownize}}</small></h2>
					</div>
				</div>
			</div>
			<div class="row" style="margin-bottom: 2em;">
				<div class="col-md-12">{{template "NavResourceGroups" .ResourceGroups}}</div>
			</div>
			<div class="row">
				<div class="col-xs-12 col-sm-12 col-md-12">
					<div class="tab-content">{{template "ResourceGroups" .ResourceGroups}}</div>
				</div>
			</div>
		</div>
		<script src="//cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js"></script>
		<script src="//netdna.bootstrapcdn.com/bootstrap/3.3.1/js/bootstrap.min.js"></script>
		<script src="//google-code-prettify.googlecode.com/svn/loader/run_prettify.js"></script>
		<script>
			jQuery(function($) {
 				$('#group-tab a[data-toggle="tab"]').on("click", function(e) {
 					window.location.hash = $(this).attr("href");
 				});

 				if(window.location.hash){
 					$('#group-tab a:first').tab('show');
 					$('#group-tab a[href$="'+ window.location.hash +'"]').tab("show");
 				} else {
 					$('#group-tab a:first').tab('show');
 				}

				$('.snippet-toggle').on("click", function(e) {
					e.preventDefault();
					var target = $(this).data('target');
					$(this).parent().next().toggle();
					if ($(this).text() == "SHOW") {
						$(this).text("HIDE");
					} else {
						$(this).text("SHOW");
					}
				});
			});
		</script>
	</body>
</html>
{{define "Responses"}}
	{{range .}}
		<li class="list-group-item bg-default response">
			<strong>Response <code>{{.Name}}</code></strong>
			<a href="javascript:;" class="pull-right btn btn-link btn-sm snippet-toggle">SHOW</a>
		</li>
		<li class="list-group-item snippet">
			{{if .Headers}}
				{{range $index := .Headers}}<code>&lt; {{.Name}}: {{.Value}}</code><br>{{end}}
			{{end}}
			{{if .Body}}
				<pre class="prettyprint">{{.Body}}</pre>
			{{end}}
		</li>
	{{end}}
{{end}}
{{define "Requests"}}
	{{range .}}
		<li class="list-group-item bg-default response">
			<strong>Requests</strong>
			<a href="javascript:;" class="pull-right btn btn-link btn-sm snippet-toggle">SHOW</a>
		</li>
		<li class="list-group-item snippet">
			{{if .Headers}}
				{{range $index := .Headers}}<code>&gt; {{.Name}}: {{.Value}}</code><br>{{end}}
			{{end}}
			{{if .Body}}
				<pre class="prettyprint">{{.Body}}</pre>
			{{end}}
		</li>
	{{end}}
{{end}}
{{define "Examples"}}
	{{range .}}
		{{template "Requests" .Requests}}
		{{template "Responses" .Responses}}
	{{end}}
{{end}}
{{define "Parameters"}}
<dl class="dl-horizontal">
	{{range $index := .}}
		<dt>{{.Name}}</dt>
		<dd>
		{{if .Required}}
			<strong>(required)</strong>
		{{end}}
		<code>{{.Type}}</code> {{.Description}}
		</dd>
	{{end}}
</dl>
{{end}}
{{define "Resources"}}
{{range .}}
	{{$UriTemplate := .UriTemplate}}
	{{$Parameters := .Parameters}}
	{{range .Actions}}
	<div class="panel panel-info">
		<div class="panel-heading">
			<span class="btn btn-{{.Method | labelize}}">{{.Method}}</span>
			<code>{{$UriTemplate}}</code>
		</div>
		<div class="panel-body">
			{{.Description | markdownize}}
		</div>
		<ul class="list-group">
			{{if $Parameters}}
				<li class="list-group-item bg-default"><strong>Parameters</strong></li>
				<li class="list-group-item">{{template "Parameters" $Parameters}}</li>
			{{end}}
			{{if .Examples}}{{template "Examples" .Examples}}{{end}}
		</ul>
	</div>
	{{end}}
{{end}}
{{end}}
{{define "ResourceGroups"}}
{{range .}}
	<div class="tab-pane" id="{{.Name | dasherize}}">
		<div class="panel panel-default">
			<div class="panel-heading">
				<h2 id="{{.Name | dasherize}}">{{.Name}}</h2>
			</div>
			<div class="panel-body">
				<p class="lead"><small>{{.Description}}</small></p>
				{{template "Resources" .Resources}}
			</div>
		</div>
	</div>
{{end}}
{{end}}
{{define "NavResourceGroups"}}
<ul class="nav nav-pills" id="group-tab">
	{{range .}}
		<li><a href="#{{.Name | dasherize}}" data-toggle="tab"><strong>{{.Name}}</strong></a></li>
	{{end}}
</ul>
{{end}}`
