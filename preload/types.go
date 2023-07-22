package preload

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"fmt"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// RuleSetID unique id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-RuleSetId
type RuleSetID string

// String returns the RuleSetID as string value.
func (t RuleSetID) String() string {
	return string(t)
}

// RuleSet corresponds to SpeculationRuleSet.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-RuleSet
type RuleSet struct {
	ID            RuleSetID         `json:"id"`
	LoaderID      cdp.LoaderID      `json:"loaderId"`                // Identifies a document which the rule set is associated with.
	SourceText    string            `json:"sourceText"`              // Source text of JSON representing the rule set. If it comes from <script> tag, it is the textContent of the node. Note that it is a JSON for valid case.  See also: - https://wicg.github.io/nav-speculation/speculation-rules.html - https://github.com/WICG/nav-speculation/blob/main/triggers.md
	BackendNodeID cdp.BackendNodeID `json:"backendNodeId,omitempty"` // A speculation rule set is either added through an inline <script> tag or through an external resource via the 'Speculation-Rules' HTTP header. For the first case, we include the BackendNodeId of the relevant <script> tag. For the second case, we include the external URL where the rule set was loaded from, and also RequestId if Network domain is enabled.  See also: - https://wicg.github.io/nav-speculation/speculation-rules.html#speculation-rules-script - https://wicg.github.io/nav-speculation/speculation-rules.html#speculation-rules-header
	URL           string            `json:"url,omitempty"`
	RequestID     network.RequestID `json:"requestId,omitempty"`
	ErrorType     RuleSetErrorType  `json:"errorType,omitempty"` // Error information errorMessage is null iff errorType is null.
}

// RuleSetErrorType [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-RuleSetErrorType
type RuleSetErrorType string

// String returns the RuleSetErrorType as string value.
func (t RuleSetErrorType) String() string {
	return string(t)
}

// RuleSetErrorType values.
const (
	RuleSetErrorTypeSourceIsNotJSONObject RuleSetErrorType = "SourceIsNotJsonObject"
	RuleSetErrorTypeInvalidRulesSkipped   RuleSetErrorType = "InvalidRulesSkipped"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t RuleSetErrorType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t RuleSetErrorType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *RuleSetErrorType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch RuleSetErrorType(v) {
	case RuleSetErrorTypeSourceIsNotJSONObject:
		*t = RuleSetErrorTypeSourceIsNotJSONObject
	case RuleSetErrorTypeInvalidRulesSkipped:
		*t = RuleSetErrorTypeInvalidRulesSkipped

	default:
		in.AddError(fmt.Errorf("unknown RuleSetErrorType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *RuleSetErrorType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SpeculationAction the type of preloading attempted. It corresponds to
// mojom::SpeculationAction (although PrefetchWithSubresources is omitted as it
// isn't being used by clients).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-SpeculationAction
type SpeculationAction string

// String returns the SpeculationAction as string value.
func (t SpeculationAction) String() string {
	return string(t)
}

// SpeculationAction values.
const (
	SpeculationActionPrefetch  SpeculationAction = "Prefetch"
	SpeculationActionPrerender SpeculationAction = "Prerender"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SpeculationAction) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SpeculationAction) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SpeculationAction) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch SpeculationAction(v) {
	case SpeculationActionPrefetch:
		*t = SpeculationActionPrefetch
	case SpeculationActionPrerender:
		*t = SpeculationActionPrerender

	default:
		in.AddError(fmt.Errorf("unknown SpeculationAction value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SpeculationAction) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SpeculationTargetHint corresponds to mojom::SpeculationTargetHint. See
// https://github.com/WICG/nav-speculation/blob/main/triggers.md#window-name-targeting-hints.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-SpeculationTargetHint
type SpeculationTargetHint string

// String returns the SpeculationTargetHint as string value.
func (t SpeculationTargetHint) String() string {
	return string(t)
}

// SpeculationTargetHint values.
const (
	SpeculationTargetHintBlank SpeculationTargetHint = "Blank"
	SpeculationTargetHintSelf  SpeculationTargetHint = "Self"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SpeculationTargetHint) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SpeculationTargetHint) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SpeculationTargetHint) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch SpeculationTargetHint(v) {
	case SpeculationTargetHintBlank:
		*t = SpeculationTargetHintBlank
	case SpeculationTargetHintSelf:
		*t = SpeculationTargetHintSelf

	default:
		in.AddError(fmt.Errorf("unknown SpeculationTargetHint value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SpeculationTargetHint) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// IngAttemptKey a key that identifies a preloading attempt. The url used is
// the url specified by the trigger (i.e. the initial URL), and not the final
// url that is navigated to. For example, prerendering allows same-origin main
// frame navigations during the attempt, but the attempt is still keyed with the
// initial URL.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-PreloadingAttemptKey
type IngAttemptKey struct {
	LoaderID   cdp.LoaderID          `json:"loaderId"`
	Action     SpeculationAction     `json:"action"`
	URL        string                `json:"url"`
	TargetHint SpeculationTargetHint `json:"targetHint,omitempty"`
}

// IngAttemptSource lists sources for a preloading attempt, specifically the
// ids of rule sets that had a speculation rule that triggered the attempt, and
// the BackendNodeIds of <a href> or <area href> elements that triggered the
// attempt (in the case of attempts triggered by a document rule). It is
// possible for multiple rule sets and links to trigger a single attempt.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-PreloadingAttemptSource
type IngAttemptSource struct {
	Key        *IngAttemptKey      `json:"key"`
	RuleSetIDs []RuleSetID         `json:"ruleSetIds"`
	NodeIDs    []cdp.BackendNodeID `json:"nodeIds"`
}

// PrerenderFinalStatus list of FinalStatus reasons for Prerender2.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-PrerenderFinalStatus
type PrerenderFinalStatus string

// String returns the PrerenderFinalStatus as string value.
func (t PrerenderFinalStatus) String() string {
	return string(t)
}

// PrerenderFinalStatus values.
const (
	PrerenderFinalStatusActivated                                                  PrerenderFinalStatus = "Activated"
	PrerenderFinalStatusDestroyed                                                  PrerenderFinalStatus = "Destroyed"
	PrerenderFinalStatusLowEndDevice                                               PrerenderFinalStatus = "LowEndDevice"
	PrerenderFinalStatusInvalidSchemeRedirect                                      PrerenderFinalStatus = "InvalidSchemeRedirect"
	PrerenderFinalStatusInvalidSchemeNavigation                                    PrerenderFinalStatus = "InvalidSchemeNavigation"
	PrerenderFinalStatusInProgressNavigation                                       PrerenderFinalStatus = "InProgressNavigation"
	PrerenderFinalStatusNavigationRequestBlockedByCsp                              PrerenderFinalStatus = "NavigationRequestBlockedByCsp"
	PrerenderFinalStatusMainFrameNavigation                                        PrerenderFinalStatus = "MainFrameNavigation"
	PrerenderFinalStatusMojoBinderPolicy                                           PrerenderFinalStatus = "MojoBinderPolicy"
	PrerenderFinalStatusRendererProcessCrashed                                     PrerenderFinalStatus = "RendererProcessCrashed"
	PrerenderFinalStatusRendererProcessKilled                                      PrerenderFinalStatus = "RendererProcessKilled"
	PrerenderFinalStatusDownload                                                   PrerenderFinalStatus = "Download"
	PrerenderFinalStatusTriggerDestroyed                                           PrerenderFinalStatus = "TriggerDestroyed"
	PrerenderFinalStatusNavigationNotCommitted                                     PrerenderFinalStatus = "NavigationNotCommitted"
	PrerenderFinalStatusNavigationBadHTTPStatus                                    PrerenderFinalStatus = "NavigationBadHttpStatus"
	PrerenderFinalStatusClientCertRequested                                        PrerenderFinalStatus = "ClientCertRequested"
	PrerenderFinalStatusNavigationRequestNetworkError                              PrerenderFinalStatus = "NavigationRequestNetworkError"
	PrerenderFinalStatusMaxNumOfRunningPrerendersExceeded                          PrerenderFinalStatus = "MaxNumOfRunningPrerendersExceeded"
	PrerenderFinalStatusCancelAllHostsForTesting                                   PrerenderFinalStatus = "CancelAllHostsForTesting"
	PrerenderFinalStatusDidFailLoad                                                PrerenderFinalStatus = "DidFailLoad"
	PrerenderFinalStatusStop                                                       PrerenderFinalStatus = "Stop"
	PrerenderFinalStatusSslCertificateError                                        PrerenderFinalStatus = "SslCertificateError"
	PrerenderFinalStatusLoginAuthRequested                                         PrerenderFinalStatus = "LoginAuthRequested"
	PrerenderFinalStatusUaChangeRequiresReload                                     PrerenderFinalStatus = "UaChangeRequiresReload"
	PrerenderFinalStatusBlockedByClient                                            PrerenderFinalStatus = "BlockedByClient"
	PrerenderFinalStatusAudioOutputDeviceRequested                                 PrerenderFinalStatus = "AudioOutputDeviceRequested"
	PrerenderFinalStatusMixedContent                                               PrerenderFinalStatus = "MixedContent"
	PrerenderFinalStatusTriggerBackgrounded                                        PrerenderFinalStatus = "TriggerBackgrounded"
	PrerenderFinalStatusMemoryLimitExceeded                                        PrerenderFinalStatus = "MemoryLimitExceeded"
	PrerenderFinalStatusFailToGetMemoryUsage                                       PrerenderFinalStatus = "FailToGetMemoryUsage"
	PrerenderFinalStatusDataSaverEnabled                                           PrerenderFinalStatus = "DataSaverEnabled"
	PrerenderFinalStatusHasEffectiveURL                                            PrerenderFinalStatus = "HasEffectiveUrl"
	PrerenderFinalStatusActivatedBeforeStarted                                     PrerenderFinalStatus = "ActivatedBeforeStarted"
	PrerenderFinalStatusInactivePageRestriction                                    PrerenderFinalStatus = "InactivePageRestriction"
	PrerenderFinalStatusStartFailed                                                PrerenderFinalStatus = "StartFailed"
	PrerenderFinalStatusTimeoutBackgrounded                                        PrerenderFinalStatus = "TimeoutBackgrounded"
	PrerenderFinalStatusCrossSiteRedirectInInitialNavigation                       PrerenderFinalStatus = "CrossSiteRedirectInInitialNavigation"
	PrerenderFinalStatusCrossSiteNavigationInInitialNavigation                     PrerenderFinalStatus = "CrossSiteNavigationInInitialNavigation"
	PrerenderFinalStatusSameSiteCrossOriginRedirectNotOptInInInitialNavigation     PrerenderFinalStatus = "SameSiteCrossOriginRedirectNotOptInInInitialNavigation"
	PrerenderFinalStatusSameSiteCrossOriginNavigationNotOptInInInitialNavigation   PrerenderFinalStatus = "SameSiteCrossOriginNavigationNotOptInInInitialNavigation"
	PrerenderFinalStatusActivationNavigationParameterMismatch                      PrerenderFinalStatus = "ActivationNavigationParameterMismatch"
	PrerenderFinalStatusActivatedInBackground                                      PrerenderFinalStatus = "ActivatedInBackground"
	PrerenderFinalStatusEmbedderHostDisallowed                                     PrerenderFinalStatus = "EmbedderHostDisallowed"
	PrerenderFinalStatusActivationNavigationDestroyedBeforeSuccess                 PrerenderFinalStatus = "ActivationNavigationDestroyedBeforeSuccess"
	PrerenderFinalStatusTabClosedByUserGesture                                     PrerenderFinalStatus = "TabClosedByUserGesture"
	PrerenderFinalStatusTabClosedWithoutUserGesture                                PrerenderFinalStatus = "TabClosedWithoutUserGesture"
	PrerenderFinalStatusPrimaryMainFrameRendererProcessCrashed                     PrerenderFinalStatus = "PrimaryMainFrameRendererProcessCrashed"
	PrerenderFinalStatusPrimaryMainFrameRendererProcessKilled                      PrerenderFinalStatus = "PrimaryMainFrameRendererProcessKilled"
	PrerenderFinalStatusActivationFramePolicyNotCompatible                         PrerenderFinalStatus = "ActivationFramePolicyNotCompatible"
	PrerenderFinalStatusPreloadingDisabled                                         PrerenderFinalStatus = "PreloadingDisabled"
	PrerenderFinalStatusBatterySaverEnabled                                        PrerenderFinalStatus = "BatterySaverEnabled"
	PrerenderFinalStatusActivatedDuringMainFrameNavigation                         PrerenderFinalStatus = "ActivatedDuringMainFrameNavigation"
	PrerenderFinalStatusPreloadingUnsupportedByWebContents                         PrerenderFinalStatus = "PreloadingUnsupportedByWebContents"
	PrerenderFinalStatusCrossSiteRedirectInMainFrameNavigation                     PrerenderFinalStatus = "CrossSiteRedirectInMainFrameNavigation"
	PrerenderFinalStatusCrossSiteNavigationInMainFrameNavigation                   PrerenderFinalStatus = "CrossSiteNavigationInMainFrameNavigation"
	PrerenderFinalStatusSameSiteCrossOriginRedirectNotOptInInMainFrameNavigation   PrerenderFinalStatus = "SameSiteCrossOriginRedirectNotOptInInMainFrameNavigation"
	PrerenderFinalStatusSameSiteCrossOriginNavigationNotOptInInMainFrameNavigation PrerenderFinalStatus = "SameSiteCrossOriginNavigationNotOptInInMainFrameNavigation"
	PrerenderFinalStatusMemoryPressureOnTrigger                                    PrerenderFinalStatus = "MemoryPressureOnTrigger"
	PrerenderFinalStatusMemoryPressureAfterTriggered                               PrerenderFinalStatus = "MemoryPressureAfterTriggered"
	PrerenderFinalStatusPrerenderingDisabledByDevTools                             PrerenderFinalStatus = "PrerenderingDisabledByDevTools"
	PrerenderFinalStatusResourceLoadBlockedByClient                                PrerenderFinalStatus = "ResourceLoadBlockedByClient"
	PrerenderFinalStatusSpeculationRuleRemoved                                     PrerenderFinalStatus = "SpeculationRuleRemoved"
	PrerenderFinalStatusActivatedWithAuxiliaryBrowsingContexts                     PrerenderFinalStatus = "ActivatedWithAuxiliaryBrowsingContexts"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PrerenderFinalStatus) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PrerenderFinalStatus) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PrerenderFinalStatus) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch PrerenderFinalStatus(v) {
	case PrerenderFinalStatusActivated:
		*t = PrerenderFinalStatusActivated
	case PrerenderFinalStatusDestroyed:
		*t = PrerenderFinalStatusDestroyed
	case PrerenderFinalStatusLowEndDevice:
		*t = PrerenderFinalStatusLowEndDevice
	case PrerenderFinalStatusInvalidSchemeRedirect:
		*t = PrerenderFinalStatusInvalidSchemeRedirect
	case PrerenderFinalStatusInvalidSchemeNavigation:
		*t = PrerenderFinalStatusInvalidSchemeNavigation
	case PrerenderFinalStatusInProgressNavigation:
		*t = PrerenderFinalStatusInProgressNavigation
	case PrerenderFinalStatusNavigationRequestBlockedByCsp:
		*t = PrerenderFinalStatusNavigationRequestBlockedByCsp
	case PrerenderFinalStatusMainFrameNavigation:
		*t = PrerenderFinalStatusMainFrameNavigation
	case PrerenderFinalStatusMojoBinderPolicy:
		*t = PrerenderFinalStatusMojoBinderPolicy
	case PrerenderFinalStatusRendererProcessCrashed:
		*t = PrerenderFinalStatusRendererProcessCrashed
	case PrerenderFinalStatusRendererProcessKilled:
		*t = PrerenderFinalStatusRendererProcessKilled
	case PrerenderFinalStatusDownload:
		*t = PrerenderFinalStatusDownload
	case PrerenderFinalStatusTriggerDestroyed:
		*t = PrerenderFinalStatusTriggerDestroyed
	case PrerenderFinalStatusNavigationNotCommitted:
		*t = PrerenderFinalStatusNavigationNotCommitted
	case PrerenderFinalStatusNavigationBadHTTPStatus:
		*t = PrerenderFinalStatusNavigationBadHTTPStatus
	case PrerenderFinalStatusClientCertRequested:
		*t = PrerenderFinalStatusClientCertRequested
	case PrerenderFinalStatusNavigationRequestNetworkError:
		*t = PrerenderFinalStatusNavigationRequestNetworkError
	case PrerenderFinalStatusMaxNumOfRunningPrerendersExceeded:
		*t = PrerenderFinalStatusMaxNumOfRunningPrerendersExceeded
	case PrerenderFinalStatusCancelAllHostsForTesting:
		*t = PrerenderFinalStatusCancelAllHostsForTesting
	case PrerenderFinalStatusDidFailLoad:
		*t = PrerenderFinalStatusDidFailLoad
	case PrerenderFinalStatusStop:
		*t = PrerenderFinalStatusStop
	case PrerenderFinalStatusSslCertificateError:
		*t = PrerenderFinalStatusSslCertificateError
	case PrerenderFinalStatusLoginAuthRequested:
		*t = PrerenderFinalStatusLoginAuthRequested
	case PrerenderFinalStatusUaChangeRequiresReload:
		*t = PrerenderFinalStatusUaChangeRequiresReload
	case PrerenderFinalStatusBlockedByClient:
		*t = PrerenderFinalStatusBlockedByClient
	case PrerenderFinalStatusAudioOutputDeviceRequested:
		*t = PrerenderFinalStatusAudioOutputDeviceRequested
	case PrerenderFinalStatusMixedContent:
		*t = PrerenderFinalStatusMixedContent
	case PrerenderFinalStatusTriggerBackgrounded:
		*t = PrerenderFinalStatusTriggerBackgrounded
	case PrerenderFinalStatusMemoryLimitExceeded:
		*t = PrerenderFinalStatusMemoryLimitExceeded
	case PrerenderFinalStatusFailToGetMemoryUsage:
		*t = PrerenderFinalStatusFailToGetMemoryUsage
	case PrerenderFinalStatusDataSaverEnabled:
		*t = PrerenderFinalStatusDataSaverEnabled
	case PrerenderFinalStatusHasEffectiveURL:
		*t = PrerenderFinalStatusHasEffectiveURL
	case PrerenderFinalStatusActivatedBeforeStarted:
		*t = PrerenderFinalStatusActivatedBeforeStarted
	case PrerenderFinalStatusInactivePageRestriction:
		*t = PrerenderFinalStatusInactivePageRestriction
	case PrerenderFinalStatusStartFailed:
		*t = PrerenderFinalStatusStartFailed
	case PrerenderFinalStatusTimeoutBackgrounded:
		*t = PrerenderFinalStatusTimeoutBackgrounded
	case PrerenderFinalStatusCrossSiteRedirectInInitialNavigation:
		*t = PrerenderFinalStatusCrossSiteRedirectInInitialNavigation
	case PrerenderFinalStatusCrossSiteNavigationInInitialNavigation:
		*t = PrerenderFinalStatusCrossSiteNavigationInInitialNavigation
	case PrerenderFinalStatusSameSiteCrossOriginRedirectNotOptInInInitialNavigation:
		*t = PrerenderFinalStatusSameSiteCrossOriginRedirectNotOptInInInitialNavigation
	case PrerenderFinalStatusSameSiteCrossOriginNavigationNotOptInInInitialNavigation:
		*t = PrerenderFinalStatusSameSiteCrossOriginNavigationNotOptInInInitialNavigation
	case PrerenderFinalStatusActivationNavigationParameterMismatch:
		*t = PrerenderFinalStatusActivationNavigationParameterMismatch
	case PrerenderFinalStatusActivatedInBackground:
		*t = PrerenderFinalStatusActivatedInBackground
	case PrerenderFinalStatusEmbedderHostDisallowed:
		*t = PrerenderFinalStatusEmbedderHostDisallowed
	case PrerenderFinalStatusActivationNavigationDestroyedBeforeSuccess:
		*t = PrerenderFinalStatusActivationNavigationDestroyedBeforeSuccess
	case PrerenderFinalStatusTabClosedByUserGesture:
		*t = PrerenderFinalStatusTabClosedByUserGesture
	case PrerenderFinalStatusTabClosedWithoutUserGesture:
		*t = PrerenderFinalStatusTabClosedWithoutUserGesture
	case PrerenderFinalStatusPrimaryMainFrameRendererProcessCrashed:
		*t = PrerenderFinalStatusPrimaryMainFrameRendererProcessCrashed
	case PrerenderFinalStatusPrimaryMainFrameRendererProcessKilled:
		*t = PrerenderFinalStatusPrimaryMainFrameRendererProcessKilled
	case PrerenderFinalStatusActivationFramePolicyNotCompatible:
		*t = PrerenderFinalStatusActivationFramePolicyNotCompatible
	case PrerenderFinalStatusPreloadingDisabled:
		*t = PrerenderFinalStatusPreloadingDisabled
	case PrerenderFinalStatusBatterySaverEnabled:
		*t = PrerenderFinalStatusBatterySaverEnabled
	case PrerenderFinalStatusActivatedDuringMainFrameNavigation:
		*t = PrerenderFinalStatusActivatedDuringMainFrameNavigation
	case PrerenderFinalStatusPreloadingUnsupportedByWebContents:
		*t = PrerenderFinalStatusPreloadingUnsupportedByWebContents
	case PrerenderFinalStatusCrossSiteRedirectInMainFrameNavigation:
		*t = PrerenderFinalStatusCrossSiteRedirectInMainFrameNavigation
	case PrerenderFinalStatusCrossSiteNavigationInMainFrameNavigation:
		*t = PrerenderFinalStatusCrossSiteNavigationInMainFrameNavigation
	case PrerenderFinalStatusSameSiteCrossOriginRedirectNotOptInInMainFrameNavigation:
		*t = PrerenderFinalStatusSameSiteCrossOriginRedirectNotOptInInMainFrameNavigation
	case PrerenderFinalStatusSameSiteCrossOriginNavigationNotOptInInMainFrameNavigation:
		*t = PrerenderFinalStatusSameSiteCrossOriginNavigationNotOptInInMainFrameNavigation
	case PrerenderFinalStatusMemoryPressureOnTrigger:
		*t = PrerenderFinalStatusMemoryPressureOnTrigger
	case PrerenderFinalStatusMemoryPressureAfterTriggered:
		*t = PrerenderFinalStatusMemoryPressureAfterTriggered
	case PrerenderFinalStatusPrerenderingDisabledByDevTools:
		*t = PrerenderFinalStatusPrerenderingDisabledByDevTools
	case PrerenderFinalStatusResourceLoadBlockedByClient:
		*t = PrerenderFinalStatusResourceLoadBlockedByClient
	case PrerenderFinalStatusSpeculationRuleRemoved:
		*t = PrerenderFinalStatusSpeculationRuleRemoved
	case PrerenderFinalStatusActivatedWithAuxiliaryBrowsingContexts:
		*t = PrerenderFinalStatusActivatedWithAuxiliaryBrowsingContexts

	default:
		in.AddError(fmt.Errorf("unknown PrerenderFinalStatus value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PrerenderFinalStatus) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// IngStatus preloading status values, see also PreloadingTriggeringOutcome.
// This status is shared by prefetchStatusUpdated and prerenderStatusUpdated.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-PreloadingStatus
type IngStatus string

// String returns the IngStatus as string value.
func (t IngStatus) String() string {
	return string(t)
}

// IngStatus values.
const (
	IngStatusPending      IngStatus = "Pending"
	IngStatusRunning      IngStatus = "Running"
	IngStatusReady        IngStatus = "Ready"
	IngStatusSuccess      IngStatus = "Success"
	IngStatusFailure      IngStatus = "Failure"
	IngStatusNotSupported IngStatus = "NotSupported"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t IngStatus) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t IngStatus) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *IngStatus) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch IngStatus(v) {
	case IngStatusPending:
		*t = IngStatusPending
	case IngStatusRunning:
		*t = IngStatusRunning
	case IngStatusReady:
		*t = IngStatusReady
	case IngStatusSuccess:
		*t = IngStatusSuccess
	case IngStatusFailure:
		*t = IngStatusFailure
	case IngStatusNotSupported:
		*t = IngStatusNotSupported

	default:
		in.AddError(fmt.Errorf("unknown IngStatus value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *IngStatus) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// PrefetchStatus tODO(https://crbug.com/1384419): revisit the list of
// PrefetchStatus and filter out the ones that aren't necessary to the
// developers.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#type-PrefetchStatus
type PrefetchStatus string

// String returns the PrefetchStatus as string value.
func (t PrefetchStatus) String() string {
	return string(t)
}

// PrefetchStatus values.
const (
	PrefetchStatusPrefetchAllowed                                             PrefetchStatus = "PrefetchAllowed"
	PrefetchStatusPrefetchFailedIneligibleRedirect                            PrefetchStatus = "PrefetchFailedIneligibleRedirect"
	PrefetchStatusPrefetchFailedInvalidRedirect                               PrefetchStatus = "PrefetchFailedInvalidRedirect"
	PrefetchStatusPrefetchFailedMIMENotSupported                              PrefetchStatus = "PrefetchFailedMIMENotSupported"
	PrefetchStatusPrefetchFailedNetError                                      PrefetchStatus = "PrefetchFailedNetError"
	PrefetchStatusPrefetchFailedNon2xX                                        PrefetchStatus = "PrefetchFailedNon2XX"
	PrefetchStatusPrefetchFailedPerPageLimitExceeded                          PrefetchStatus = "PrefetchFailedPerPageLimitExceeded"
	PrefetchStatusPrefetchEvicted                                             PrefetchStatus = "PrefetchEvicted"
	PrefetchStatusPrefetchHeldback                                            PrefetchStatus = "PrefetchHeldback"
	PrefetchStatusPrefetchIneligibleRetryAfter                                PrefetchStatus = "PrefetchIneligibleRetryAfter"
	PrefetchStatusPrefetchIsPrivacyDecoy                                      PrefetchStatus = "PrefetchIsPrivacyDecoy"
	PrefetchStatusPrefetchIsStale                                             PrefetchStatus = "PrefetchIsStale"
	PrefetchStatusPrefetchNotEligibleBrowserContextOffTheRecord               PrefetchStatus = "PrefetchNotEligibleBrowserContextOffTheRecord"
	PrefetchStatusPrefetchNotEligibleDataSaverEnabled                         PrefetchStatus = "PrefetchNotEligibleDataSaverEnabled"
	PrefetchStatusPrefetchNotEligibleExistingProxy                            PrefetchStatus = "PrefetchNotEligibleExistingProxy"
	PrefetchStatusPrefetchNotEligibleHostIsNonUnique                          PrefetchStatus = "PrefetchNotEligibleHostIsNonUnique"
	PrefetchStatusPrefetchNotEligibleNonDefaultStoragePartition               PrefetchStatus = "PrefetchNotEligibleNonDefaultStoragePartition"
	PrefetchStatusPrefetchNotEligibleSameSiteCrossOriginPrefetchRequiredProxy PrefetchStatus = "PrefetchNotEligibleSameSiteCrossOriginPrefetchRequiredProxy"
	PrefetchStatusPrefetchNotEligibleSchemeIsNotHTTPS                         PrefetchStatus = "PrefetchNotEligibleSchemeIsNotHttps"
	PrefetchStatusPrefetchNotEligibleUserHasCookies                           PrefetchStatus = "PrefetchNotEligibleUserHasCookies"
	PrefetchStatusPrefetchNotEligibleUserHasServiceWorker                     PrefetchStatus = "PrefetchNotEligibleUserHasServiceWorker"
	PrefetchStatusPrefetchNotEligibleBatterySaverEnabled                      PrefetchStatus = "PrefetchNotEligibleBatterySaverEnabled"
	PrefetchStatusPrefetchNotEligiblePreloadingDisabled                       PrefetchStatus = "PrefetchNotEligiblePreloadingDisabled"
	PrefetchStatusPrefetchNotFinishedInTime                                   PrefetchStatus = "PrefetchNotFinishedInTime"
	PrefetchStatusPrefetchNotStarted                                          PrefetchStatus = "PrefetchNotStarted"
	PrefetchStatusPrefetchNotUsedCookiesChanged                               PrefetchStatus = "PrefetchNotUsedCookiesChanged"
	PrefetchStatusPrefetchProxyNotAvailable                                   PrefetchStatus = "PrefetchProxyNotAvailable"
	PrefetchStatusPrefetchResponseUsed                                        PrefetchStatus = "PrefetchResponseUsed"
	PrefetchStatusPrefetchSuccessfulButNotUsed                                PrefetchStatus = "PrefetchSuccessfulButNotUsed"
	PrefetchStatusPrefetchNotUsedProbeFailed                                  PrefetchStatus = "PrefetchNotUsedProbeFailed"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PrefetchStatus) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PrefetchStatus) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PrefetchStatus) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch PrefetchStatus(v) {
	case PrefetchStatusPrefetchAllowed:
		*t = PrefetchStatusPrefetchAllowed
	case PrefetchStatusPrefetchFailedIneligibleRedirect:
		*t = PrefetchStatusPrefetchFailedIneligibleRedirect
	case PrefetchStatusPrefetchFailedInvalidRedirect:
		*t = PrefetchStatusPrefetchFailedInvalidRedirect
	case PrefetchStatusPrefetchFailedMIMENotSupported:
		*t = PrefetchStatusPrefetchFailedMIMENotSupported
	case PrefetchStatusPrefetchFailedNetError:
		*t = PrefetchStatusPrefetchFailedNetError
	case PrefetchStatusPrefetchFailedNon2xX:
		*t = PrefetchStatusPrefetchFailedNon2xX
	case PrefetchStatusPrefetchFailedPerPageLimitExceeded:
		*t = PrefetchStatusPrefetchFailedPerPageLimitExceeded
	case PrefetchStatusPrefetchEvicted:
		*t = PrefetchStatusPrefetchEvicted
	case PrefetchStatusPrefetchHeldback:
		*t = PrefetchStatusPrefetchHeldback
	case PrefetchStatusPrefetchIneligibleRetryAfter:
		*t = PrefetchStatusPrefetchIneligibleRetryAfter
	case PrefetchStatusPrefetchIsPrivacyDecoy:
		*t = PrefetchStatusPrefetchIsPrivacyDecoy
	case PrefetchStatusPrefetchIsStale:
		*t = PrefetchStatusPrefetchIsStale
	case PrefetchStatusPrefetchNotEligibleBrowserContextOffTheRecord:
		*t = PrefetchStatusPrefetchNotEligibleBrowserContextOffTheRecord
	case PrefetchStatusPrefetchNotEligibleDataSaverEnabled:
		*t = PrefetchStatusPrefetchNotEligibleDataSaverEnabled
	case PrefetchStatusPrefetchNotEligibleExistingProxy:
		*t = PrefetchStatusPrefetchNotEligibleExistingProxy
	case PrefetchStatusPrefetchNotEligibleHostIsNonUnique:
		*t = PrefetchStatusPrefetchNotEligibleHostIsNonUnique
	case PrefetchStatusPrefetchNotEligibleNonDefaultStoragePartition:
		*t = PrefetchStatusPrefetchNotEligibleNonDefaultStoragePartition
	case PrefetchStatusPrefetchNotEligibleSameSiteCrossOriginPrefetchRequiredProxy:
		*t = PrefetchStatusPrefetchNotEligibleSameSiteCrossOriginPrefetchRequiredProxy
	case PrefetchStatusPrefetchNotEligibleSchemeIsNotHTTPS:
		*t = PrefetchStatusPrefetchNotEligibleSchemeIsNotHTTPS
	case PrefetchStatusPrefetchNotEligibleUserHasCookies:
		*t = PrefetchStatusPrefetchNotEligibleUserHasCookies
	case PrefetchStatusPrefetchNotEligibleUserHasServiceWorker:
		*t = PrefetchStatusPrefetchNotEligibleUserHasServiceWorker
	case PrefetchStatusPrefetchNotEligibleBatterySaverEnabled:
		*t = PrefetchStatusPrefetchNotEligibleBatterySaverEnabled
	case PrefetchStatusPrefetchNotEligiblePreloadingDisabled:
		*t = PrefetchStatusPrefetchNotEligiblePreloadingDisabled
	case PrefetchStatusPrefetchNotFinishedInTime:
		*t = PrefetchStatusPrefetchNotFinishedInTime
	case PrefetchStatusPrefetchNotStarted:
		*t = PrefetchStatusPrefetchNotStarted
	case PrefetchStatusPrefetchNotUsedCookiesChanged:
		*t = PrefetchStatusPrefetchNotUsedCookiesChanged
	case PrefetchStatusPrefetchProxyNotAvailable:
		*t = PrefetchStatusPrefetchProxyNotAvailable
	case PrefetchStatusPrefetchResponseUsed:
		*t = PrefetchStatusPrefetchResponseUsed
	case PrefetchStatusPrefetchSuccessfulButNotUsed:
		*t = PrefetchStatusPrefetchSuccessfulButNotUsed
	case PrefetchStatusPrefetchNotUsedProbeFailed:
		*t = PrefetchStatusPrefetchNotUsedProbeFailed

	default:
		in.AddError(fmt.Errorf("unknown PrefetchStatus value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PrefetchStatus) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
