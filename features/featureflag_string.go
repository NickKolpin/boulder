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
	_ = x[CAAValidationMethods-6]
	_ = x[CAAAccountURI-7]
	_ = x[EnforceMultiVA-8]
	_ = x[MultiVAFullResults-9]
	_ = x[MandatoryPOSTAsGET-10]
	_ = x[AllowV1Registration-11]
	_ = x[V1DisableNewValidations-12]
	_ = x[StoreRevokerInfo-13]
	_ = x[RestrictRSAKeySizes-14]
	_ = x[FasterNewOrdersRateLimit-15]
	_ = x[ECDSAForAll-16]
	_ = x[ServeRenewalInfo-17]
}

const _FeatureFlag_name = "unusedPrecertificateRevocationStripDefaultSchemePortNonCFSSLSignerStoreIssuerInfoStreamlineOrderAndAuthzsCAAValidationMethodsCAAAccountURIEnforceMultiVAMultiVAFullResultsMandatoryPOSTAsGETAllowV1RegistrationV1DisableNewValidationsStoreRevokerInfoRestrictRSAKeySizesFasterNewOrdersRateLimitECDSAForAllServeRenewalInfo"

var _FeatureFlag_index = [...]uint16{0, 6, 30, 52, 66, 81, 105, 125, 138, 152, 170, 188, 207, 230, 246, 265, 289, 300, 316}

func (i FeatureFlag) String() string {
	if i < 0 || i >= FeatureFlag(len(_FeatureFlag_index)-1) {
		return "FeatureFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FeatureFlag_name[_FeatureFlag_index[i]:_FeatureFlag_index[i+1]]
}
