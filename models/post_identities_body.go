// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostIdentitiesBody post identities body
// swagger:model post_identities_body
type PostIdentitiesBody struct {

	// The id of the connection of the secondary account being linked. This is optional and may be useful in the case of having more than a database connection for the 'auth0' provider.
	ConnectionID string `json:"connection_id,omitempty"`

	// The JWT of the secondary account being linked. If sending this parameter, the 'provider', 'user_id' and 'connection_id' params are invalid.
	LinkWith *string `json:"link_with,omitempty"`

	// The type of identity provider of the secondary account being linked.
	// Enum: [ad adfs amazon dropbox bitbucket aol auth0-adldap auth0-oidc auth0 baidu bitly box custom daccount dwolla email evernote-sandbox evernote exact facebook fitbit flickr github google-apps google-oauth2 guardian instagram ip linkedin miicard oauth1 oauth2 office365 paypal paypal-sandbox pingfederate planningcenter renren salesforce-community salesforce-sandbox salesforce samlp sharepoint shopify sms soundcloud thecity-sandbox thecity thirtysevensignals twitter untappd vkontakte waad weibo windowslive wordpress yahoo yammer yandex]
	Provider *string `json:"provider,omitempty"`

	// The user_id of the secondary account being linked.
	UserID *string `json:"user_id,omitempty"`
}

// Validate validates this post identities body
func (m *PostIdentitiesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postIdentitiesBodyTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ad","adfs","amazon","dropbox","bitbucket","aol","auth0-adldap","auth0-oidc","auth0","baidu","bitly","box","custom","daccount","dwolla","email","evernote-sandbox","evernote","exact","facebook","fitbit","flickr","github","google-apps","google-oauth2","guardian","instagram","ip","linkedin","miicard","oauth1","oauth2","office365","paypal","paypal-sandbox","pingfederate","planningcenter","renren","salesforce-community","salesforce-sandbox","salesforce","samlp","sharepoint","shopify","sms","soundcloud","thecity-sandbox","thecity","thirtysevensignals","twitter","untappd","vkontakte","waad","weibo","windowslive","wordpress","yahoo","yammer","yandex"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postIdentitiesBodyTypeProviderPropEnum = append(postIdentitiesBodyTypeProviderPropEnum, v)
	}
}

const (

	// PostIdentitiesBodyProviderAd captures enum value "ad"
	PostIdentitiesBodyProviderAd string = "ad"

	// PostIdentitiesBodyProviderAdfs captures enum value "adfs"
	PostIdentitiesBodyProviderAdfs string = "adfs"

	// PostIdentitiesBodyProviderAmazon captures enum value "amazon"
	PostIdentitiesBodyProviderAmazon string = "amazon"

	// PostIdentitiesBodyProviderDropbox captures enum value "dropbox"
	PostIdentitiesBodyProviderDropbox string = "dropbox"

	// PostIdentitiesBodyProviderBitbucket captures enum value "bitbucket"
	PostIdentitiesBodyProviderBitbucket string = "bitbucket"

	// PostIdentitiesBodyProviderAol captures enum value "aol"
	PostIdentitiesBodyProviderAol string = "aol"

	// PostIdentitiesBodyProviderAuth0Adldap captures enum value "auth0-adldap"
	PostIdentitiesBodyProviderAuth0Adldap string = "auth0-adldap"

	// PostIdentitiesBodyProviderAuth0Oidc captures enum value "auth0-oidc"
	PostIdentitiesBodyProviderAuth0Oidc string = "auth0-oidc"

	// PostIdentitiesBodyProviderAuth0 captures enum value "auth0"
	PostIdentitiesBodyProviderAuth0 string = "auth0"

	// PostIdentitiesBodyProviderBaidu captures enum value "baidu"
	PostIdentitiesBodyProviderBaidu string = "baidu"

	// PostIdentitiesBodyProviderBitly captures enum value "bitly"
	PostIdentitiesBodyProviderBitly string = "bitly"

	// PostIdentitiesBodyProviderBox captures enum value "box"
	PostIdentitiesBodyProviderBox string = "box"

	// PostIdentitiesBodyProviderCustom captures enum value "custom"
	PostIdentitiesBodyProviderCustom string = "custom"

	// PostIdentitiesBodyProviderDaccount captures enum value "daccount"
	PostIdentitiesBodyProviderDaccount string = "daccount"

	// PostIdentitiesBodyProviderDwolla captures enum value "dwolla"
	PostIdentitiesBodyProviderDwolla string = "dwolla"

	// PostIdentitiesBodyProviderEmail captures enum value "email"
	PostIdentitiesBodyProviderEmail string = "email"

	// PostIdentitiesBodyProviderEvernoteSandbox captures enum value "evernote-sandbox"
	PostIdentitiesBodyProviderEvernoteSandbox string = "evernote-sandbox"

	// PostIdentitiesBodyProviderEvernote captures enum value "evernote"
	PostIdentitiesBodyProviderEvernote string = "evernote"

	// PostIdentitiesBodyProviderExact captures enum value "exact"
	PostIdentitiesBodyProviderExact string = "exact"

	// PostIdentitiesBodyProviderFacebook captures enum value "facebook"
	PostIdentitiesBodyProviderFacebook string = "facebook"

	// PostIdentitiesBodyProviderFitbit captures enum value "fitbit"
	PostIdentitiesBodyProviderFitbit string = "fitbit"

	// PostIdentitiesBodyProviderFlickr captures enum value "flickr"
	PostIdentitiesBodyProviderFlickr string = "flickr"

	// PostIdentitiesBodyProviderGithub captures enum value "github"
	PostIdentitiesBodyProviderGithub string = "github"

	// PostIdentitiesBodyProviderGoogleApps captures enum value "google-apps"
	PostIdentitiesBodyProviderGoogleApps string = "google-apps"

	// PostIdentitiesBodyProviderGoogleOauth2 captures enum value "google-oauth2"
	PostIdentitiesBodyProviderGoogleOauth2 string = "google-oauth2"

	// PostIdentitiesBodyProviderGuardian captures enum value "guardian"
	PostIdentitiesBodyProviderGuardian string = "guardian"

	// PostIdentitiesBodyProviderInstagram captures enum value "instagram"
	PostIdentitiesBodyProviderInstagram string = "instagram"

	// PostIdentitiesBodyProviderIP captures enum value "ip"
	PostIdentitiesBodyProviderIP string = "ip"

	// PostIdentitiesBodyProviderLinkedin captures enum value "linkedin"
	PostIdentitiesBodyProviderLinkedin string = "linkedin"

	// PostIdentitiesBodyProviderMiicard captures enum value "miicard"
	PostIdentitiesBodyProviderMiicard string = "miicard"

	// PostIdentitiesBodyProviderOauth1 captures enum value "oauth1"
	PostIdentitiesBodyProviderOauth1 string = "oauth1"

	// PostIdentitiesBodyProviderOauth2 captures enum value "oauth2"
	PostIdentitiesBodyProviderOauth2 string = "oauth2"

	// PostIdentitiesBodyProviderOffice365 captures enum value "office365"
	PostIdentitiesBodyProviderOffice365 string = "office365"

	// PostIdentitiesBodyProviderPaypal captures enum value "paypal"
	PostIdentitiesBodyProviderPaypal string = "paypal"

	// PostIdentitiesBodyProviderPaypalSandbox captures enum value "paypal-sandbox"
	PostIdentitiesBodyProviderPaypalSandbox string = "paypal-sandbox"

	// PostIdentitiesBodyProviderPingfederate captures enum value "pingfederate"
	PostIdentitiesBodyProviderPingfederate string = "pingfederate"

	// PostIdentitiesBodyProviderPlanningcenter captures enum value "planningcenter"
	PostIdentitiesBodyProviderPlanningcenter string = "planningcenter"

	// PostIdentitiesBodyProviderRenren captures enum value "renren"
	PostIdentitiesBodyProviderRenren string = "renren"

	// PostIdentitiesBodyProviderSalesforceCommunity captures enum value "salesforce-community"
	PostIdentitiesBodyProviderSalesforceCommunity string = "salesforce-community"

	// PostIdentitiesBodyProviderSalesforceSandbox captures enum value "salesforce-sandbox"
	PostIdentitiesBodyProviderSalesforceSandbox string = "salesforce-sandbox"

	// PostIdentitiesBodyProviderSalesforce captures enum value "salesforce"
	PostIdentitiesBodyProviderSalesforce string = "salesforce"

	// PostIdentitiesBodyProviderSamlp captures enum value "samlp"
	PostIdentitiesBodyProviderSamlp string = "samlp"

	// PostIdentitiesBodyProviderSharepoint captures enum value "sharepoint"
	PostIdentitiesBodyProviderSharepoint string = "sharepoint"

	// PostIdentitiesBodyProviderShopify captures enum value "shopify"
	PostIdentitiesBodyProviderShopify string = "shopify"

	// PostIdentitiesBodyProviderSms captures enum value "sms"
	PostIdentitiesBodyProviderSms string = "sms"

	// PostIdentitiesBodyProviderSoundcloud captures enum value "soundcloud"
	PostIdentitiesBodyProviderSoundcloud string = "soundcloud"

	// PostIdentitiesBodyProviderThecitySandbox captures enum value "thecity-sandbox"
	PostIdentitiesBodyProviderThecitySandbox string = "thecity-sandbox"

	// PostIdentitiesBodyProviderThecity captures enum value "thecity"
	PostIdentitiesBodyProviderThecity string = "thecity"

	// PostIdentitiesBodyProviderThirtysevensignals captures enum value "thirtysevensignals"
	PostIdentitiesBodyProviderThirtysevensignals string = "thirtysevensignals"

	// PostIdentitiesBodyProviderTwitter captures enum value "twitter"
	PostIdentitiesBodyProviderTwitter string = "twitter"

	// PostIdentitiesBodyProviderUntappd captures enum value "untappd"
	PostIdentitiesBodyProviderUntappd string = "untappd"

	// PostIdentitiesBodyProviderVkontakte captures enum value "vkontakte"
	PostIdentitiesBodyProviderVkontakte string = "vkontakte"

	// PostIdentitiesBodyProviderWaad captures enum value "waad"
	PostIdentitiesBodyProviderWaad string = "waad"

	// PostIdentitiesBodyProviderWeibo captures enum value "weibo"
	PostIdentitiesBodyProviderWeibo string = "weibo"

	// PostIdentitiesBodyProviderWindowslive captures enum value "windowslive"
	PostIdentitiesBodyProviderWindowslive string = "windowslive"

	// PostIdentitiesBodyProviderWordpress captures enum value "wordpress"
	PostIdentitiesBodyProviderWordpress string = "wordpress"

	// PostIdentitiesBodyProviderYahoo captures enum value "yahoo"
	PostIdentitiesBodyProviderYahoo string = "yahoo"

	// PostIdentitiesBodyProviderYammer captures enum value "yammer"
	PostIdentitiesBodyProviderYammer string = "yammer"

	// PostIdentitiesBodyProviderYandex captures enum value "yandex"
	PostIdentitiesBodyProviderYandex string = "yandex"
)

// prop value enum
func (m *PostIdentitiesBody) validateProviderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postIdentitiesBodyTypeProviderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostIdentitiesBody) validateProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("provider", "body", *m.Provider); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostIdentitiesBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostIdentitiesBody) UnmarshalBinary(b []byte) error {
	var res PostIdentitiesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
