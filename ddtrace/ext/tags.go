// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016 Datadog, Inc.

// Package ext contains a set of Datadog-specific constants. Most of them are used
// for setting span metadata.
package ext

const (
	// TargetHost sets the target host address.
	// Legacy: Kept for backwards compatibility. Use NetworkDestinationName for hostname
	// and NetworkDestinationIP for IP addresses
	TargetHost = "out.host"

	// NetworkDestinationName is the remote hostname or similar where the outbound connection is being made to.
	NetworkDestinationName = "network.destination.name"

	// NetworkDestinationIP is the remote address where the outbound connection is being made to.
	NetworkDestinationIP = "network.destination.ip"

	// NetworkClientIP is the client IP address.
	NetworkClientIP = "network.client.ip"

	// TargetPort sets the target host port.
	// Legacy: Kept for backwards compatability. Use NetworkDestinationPort instead.
	TargetPort = "out.port"

	// TargetDB sets the target db.
	TargetDB = "out.db"

	// NetworkDestinationPort is the remote port number of the outbound connection.
	NetworkDestinationPort = "network.destination.port"

	// SQLType sets the sql type tag.
	SQLType = "sql"

	// SQLQuery sets the sql query tag on a span.
	SQLQuery = "sql.query"

	// HTTPMethod specifies the HTTP method used in a span.
	HTTPMethod = "http.method"

	// HTTPCode sets the HTTP status code as a tag.
	HTTPCode = "http.status_code"

	// HTTPRoute is the route value of the HTTP request.
	HTTPRoute = "http.route"

	// HTTPURL sets the HTTP URL for a span.
	HTTPURL = "http.url"

	// HTTPUserAgent is the user agent header value of the HTTP request.
	HTTPUserAgent = "http.useragent"

	// HTTPClientIP sets the HTTP client IP tag.
	HTTPClientIP = "http.client_ip"

	// HTTPRequestHeaders sets the HTTP request headers partial tag
	// This tag is meant to be composed, i.e http.request.headers.headerX, http.request.headers.headerY, etc...
	// See https://docs.datadoghq.com/tracing/trace_collection/tracing_naming_convention/#http-requests
	HTTPRequestHeaders = "http.request.headers"

	// SpanName is a pseudo-key for setting a span's operation name by means of
	// a tag. It is mostly here to facilitate vendor-agnostic frameworks like Opentracing
	// and OpenCensus.
	SpanName = "span.name"

	// SpanType defines the Span type (web, db, cache).
	SpanType = "span.type"

	// ServiceName defines the Service name for this Span.
	ServiceName = "service.name"

	// Version is a tag that specifies the current application version.
	Version = "version"

	// ResourceName defines the Resource name for the Span.
	ResourceName = "resource.name"

	// Error specifies the error tag. It's value is usually of type "error".
	Error = "error"

	// ErrorMsg specifies the error message.
	ErrorMsg = "error.message"

	// ErrorType specifies the error type.
	ErrorType = "error.type"

	// ErrorStack specifies the stack dump.
	ErrorStack = "error.stack"

	// ErrorDetails holds details about an error which implements a formatter.
	ErrorDetails = "error.details"

	// Environment specifies the environment to use with a trace.
	Environment = "env"

	// EventSampleRate specifies the rate at which this span will be sampled
	// as an APM event.
	EventSampleRate = "_dd1.sr.eausr"

	// AnalyticsEvent specifies whether the span should be recorded as a Trace
	// Search & Analytics event.
	AnalyticsEvent = "analytics.event"

	// ManualKeep is a tag which specifies that the trace to which this span
	// belongs to should be kept when set to true.
	ManualKeep = "manual.keep"

	// ManualDrop is a tag which specifies that the trace to which this span
	// belongs to should be dropped when set to true.
	ManualDrop = "manual.drop"

	// RuntimeID is a tag that contains a unique id for this process.
	RuntimeID = "runtime-id"

	// Component defines library integration the span originated from.
	Component = "component"

	// SpanKind defines the kind of span based on Otel requirements (client, server, producer, consumer).
	SpanKind = "span.kind"

	// MapSpanStart is used by Span.AsMap to store the span start.
	MapSpanStart = "_ddtrace.span_start"

	// MapSpanDuration is used by Span.AsMap to store the span duration.
	MapSpanDuration = "_ddtrace.span_duration"

	// MapSpanSpanID is used by Span.AsMap to store the span id.
	MapSpanID = "_ddtrace.span_id"

	// MapSpanTraceID is used by Span.AsMap to store the span trace id.
	MapSpanTraceID = "_ddtrace.span_traceid"

	// MapSpanParentID is used by Span.AsMap to store the span parent id.
	MapSpanParentID = "_ddtrace.span_parentid"

	// MapSpanError is used by Span.AsMap to store the span error value.
	MapSpanError = "_ddtrace.span_error"

	// MapSpanEvents is used by Span.AsMap to store the spanEvents value.
	MapSpanEvents = "_ddtrace.span_events"
)
