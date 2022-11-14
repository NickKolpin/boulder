// Code generated by "stringer -type=FeatureFlag"; DO NOT EDIT.

package features

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[unused-0]
	_ = x[PrecertificateRevocation-1]
	_ = x[StripDefaultSchemePort-2]
	_ = x[NonCFSSLSigner-3]
	_ = x[StoreIssuerInfo-4]
	_ = x[StreamlineOrderAndAuthzs-5]
	_ = x[V1DisableNewValidations-6]
	_ = x[ExpirationMailerDontLookTwice-7]
	_ = x[OldTLSInbound-8]
	_ = x[OldTLSOutbound-9]
	_ = x[ROCSPStage1-10]
	_ = x[ROCSPStage2-11]
	_ = x[ROCSPStage3-12]
	_ = x[GetAuthzReadOnly-13]
	_ = x[GetAuthzUseIndex-14]
	_ = x[CheckFailedAuthorizationsFirst-15]
	_ = x[FasterNewOrdersRateLimit-16]
	_ = x[AllowV1Registration-17]
	_ = x[RestrictRSAKeySizes-18]
	_ = x[AllowReRevocation-19]
	_ = x[MozRevocationReasons-20]
	_ = x[SHA1CSRs-21]
	_ = x[RejectDuplicateCSRExtensions-22]
	_ = x[CAAValidationMethods-23]
	_ = x[CAAAccountURI-24]
	_ = x[EnforceMultiVA-25]
	_ = x[MultiVAFullResults-26]
	_ = x[MandatoryPOSTAsGET-27]
	_ = x[StoreRevokerInfo-28]
	_ = x[ECDSAForAll-29]
	_ = x[ServeRenewalInfo-30]
	_ = x[AllowUnrecognizedFeatures-31]
	_ = x[ROCSPStage6-32]
	_ = x[ROCSPStage7-33]
}

const _FeatureFlag_name = "unusedPrecertificateRevocationStripDefaultSchemePortNonCFSSLSignerStoreIssuerInfoStreamlineOrderAndAuthzsV1DisableNewValidationsExpirationMailerDontLookTwiceOldTLSInboundOldTLSOutboundROCSPStage1ROCSPStage2ROCSPStage3GetAuthzReadOnlyGetAuthzUseIndexCheckFailedAuthorizationsFirstFasterNewOrdersRateLimitAllowV1RegistrationRestrictRSAKeySizesAllowReRevocationMozRevocationReasonsSHA1CSRsRejectDuplicateCSRExtensionsCAAValidationMethodsCAAAccountURIEnforceMultiVAMultiVAFullResultsMandatoryPOSTAsGETStoreRevokerInfoECDSAForAllServeRenewalInfoAllowUnrecognizedFeaturesROCSPStage6ROCSPStage7"

var _FeatureFlag_index = [...]uint16{0, 6, 30, 52, 66, 81, 105, 128, 157, 170, 184, 195, 206, 217, 233, 249, 279, 303, 322, 341, 358, 378, 386, 414, 434, 447, 461, 479, 497, 513, 524, 540, 565, 576, 587}

func (i FeatureFlag) String() string {
	if i < 0 || i >= FeatureFlag(len(_FeatureFlag_index)-1) {
		return "FeatureFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FeatureFlag_name[_FeatureFlag_index[i]:_FeatureFlag_index[i+1]]
}
