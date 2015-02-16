package main

type AdminOrg struct {
	Description string `xml:"Description"`
	FullName    string `xml:"FullName"`
	//IsEnabled     			bool   	`xml:"HostedZoneConfig>Comment"`
	IsEnabled    bool `xml:"IsEnabled"`
	Settings     Settingsx
	Xmlns        string `xml:"xmlns,attr"`
	Xmlnsjaxb    string `xml:"xmlns:jaxb,attr"`
	Xmlnsmeta    string `xml:"xmlns:meta,attr"`
	Xmlnsxsi     string `xml:"xmlns:xsi,attr"`
	XsiSchemaLoc string `xml:"xsi:schemaLocation,attr"`
	Name         string `xml:"name,attr"`
}

func (a *AdminOrg) SetName(name string) {
	a.Name = name
}
func (a AdminOrg) GetName() string {
	return a.Name
}

func (a *AdminOrg) SetDescription(description string) {
	a.Description = description
}
func (a AdminOrg) GetDescription() string {
	return a.Description
}

func (a *AdminOrg) SetFullName(fullName string) {
	a.FullName = fullName
}
func (a AdminOrg) GetFullName() string {
	return a.FullName
}

func (a *AdminOrg) SetIsEnabled(isEnabled bool) {
	a.IsEnabled = isEnabled
}
func (a AdminOrg) GetIsEnabled() bool {
	return a.IsEnabled
}

type Settingsx struct {
	OrgGeneralSettings OrgGeneralSettingsX
	OrgLdapSettings    OrgLdapSettingsX
	OrgEmailSettings   OrgEmailSettingsX
}

type OrgGeneralSettingsX struct {
	CanPublishCatalogs       bool
	DeployedVMQuota          int64
	StoredVmQuota            int64
	UseServerBootSequence    bool
	DelayAfterPowerOnSeconds int64
}

func (g *OrgGeneralSettingsX) SetCanPublishCatalogs(CanPublishCatalogs bool) {

	g.CanPublishCatalogs = CanPublishCatalogs
}

func (g *OrgGeneralSettingsX) SetDeployedVMQuota(DeployedVMQuota int64) {

	g.DeployedVMQuota = DeployedVMQuota
}
func (g *OrgGeneralSettingsX) SetStoredVmQuota(StoredVmQuota int64) {

	g.StoredVmQuota = StoredVmQuota
}
func (g *OrgGeneralSettingsX) SetUseServerBootSequence(UseServerBootSequence bool) {

	g.UseServerBootSequence = UseServerBootSequence
}
func (g *OrgGeneralSettingsX) SetDelayAfterPowerOnSeconds(DelayAfterPowerOnSeconds int64) {

	g.DelayAfterPowerOnSeconds = DelayAfterPowerOnSeconds
}

type OrgEmailSettingsX struct {
	IsDefaultSmtpServer     bool
	IsDefaultOrgEmail       bool
	FromEmailAddress        string
	DefaultSubjectPrefix    string
	IsAlertEmailToAllAdmins bool
}

func (e *OrgEmailSettingsX) SetIsDefaultSmtpServer(IsDefaultSmtpServer bool) {
	e.IsDefaultSmtpServer = IsDefaultSmtpServer

}

func (e *OrgEmailSettingsX) SetIsDefaultOrgEmail(IsDefaultOrgEmail bool) {
	e.IsDefaultOrgEmail = IsDefaultOrgEmail

}

func (e *OrgEmailSettingsX) SetFromEmailAddress(FromEmailAddress string) {
	e.FromEmailAddress = FromEmailAddress

}

func (e *OrgEmailSettingsX) SetDefaultSubjectPrefix(DefaultSubjectPrefix string) {
	e.DefaultSubjectPrefix = DefaultSubjectPrefix

}

func (e *OrgEmailSettingsX) SetIsAlertEmailToAllAdmins(IsAlertEmailToAllAdmins bool) {
	e.IsAlertEmailToAllAdmins = IsAlertEmailToAllAdmins

}

type OrgLdapSettingsX struct {
	OrgLdapMode           string
	CustomUsersOu         string
	CustomOrgLdapSettings CustomOrgLdapSettingsX
}

func (l *OrgLdapSettingsX) SetOrgLdapMode(OrgLdapMode string) {
	l.OrgLdapMode = OrgLdapMode
}

func (l *OrgLdapSettingsX) SetCustomUsersOu(CustomUsersOu string) {
	l.CustomUsersOu = CustomUsersOu
}

type CustomOrgLdapSettingsX struct {
	HostName                 string
	Port                     string
	SearchBase               string
	UserName                 string
	Password                 string
	AuthenticationMechanism  string
	GroupSearchBase          string
	IsGroupSearchBaseEnabled bool
	ConnectorType            string
	UserAttributes           UserAttributesX
	GroupAttributes          GroupAttributesX
}

func (c *CustomOrgLdapSettingsX) SetHostName(HostName string) {
	c.HostName = HostName
}

func (c *CustomOrgLdapSettingsX) SetPort(Port string) {
	c.Port = Port
}

func (c *CustomOrgLdapSettingsX) SetSearchBase(SearchBase string) {
	c.SearchBase = SearchBase
}

func (c *CustomOrgLdapSettingsX) SetUserName(UserName string) {
	c.UserName = UserName
}

func (c *CustomOrgLdapSettingsX) SetAuthenticationMechanism(AuthenticationMechanism string) {
	c.AuthenticationMechanism = AuthenticationMechanism
}
func (c *CustomOrgLdapSettingsX) SetGroupSearchBase(GroupSearchBase string) {
	c.GroupSearchBase = GroupSearchBase
}
func (c *CustomOrgLdapSettingsX) SetIsGroupSearchBaseEnabled(IsGroupSearchBaseEnabled bool) {
	c.IsGroupSearchBaseEnabled = IsGroupSearchBaseEnabled
}

type UserAttributesX struct {
	ObjectClass               string
	ObjectIdentifier          string
	UserName                  string
	Email                     string
	FullName                  string
	GivenName                 string
	Surname                   string
	Telephone                 string
	GroupMembershipIdentifier string
}

func (u *UserAttributesX) SetObjectClass(ObjectClass string) {
	u.ObjectClass = ObjectClass
}
func (u *UserAttributesX) SetObjectIdentifier(ObjectIdentifier string) {
	u.ObjectIdentifier = ObjectIdentifier
}
func (u *UserAttributesX) SetUserName(UserName string) {
	u.UserName = UserName
}
func (u *UserAttributesX) SetEmail(Email string) {
	u.Email = Email
}
func (u *UserAttributesX) SetFullName(FullName string) {
	u.FullName = FullName
}
func (u *UserAttributesX) SetGivenName(GivenName string) {
	u.GivenName = GivenName
}
func (u *UserAttributesX) SetSurname(Surname string) {
	u.Surname = Surname
}
func (u *UserAttributesX) SetTelephone(Telephone string) {
	u.Telephone = Telephone
}
func (u *UserAttributesX) SetGroupMembershipIdentifier(GroupMembershipIdentifier string) {
	u.GroupMembershipIdentifier = GroupMembershipIdentifier
}

type GroupAttributesX struct {
	ObjectClass          string
	ObjectIdentifier     string
	GroupName            string
	Membership           string
	MembershipIdentifier string
}

func (c *GroupAttributesX) SetObjectClass(ObjectClass string) {
	c.ObjectClass = ObjectClass
}
func (c *GroupAttributesX) SetObjectIdentifier(ObjectIdentifier string) {
	c.ObjectIdentifier = ObjectIdentifier
}
func (c *GroupAttributesX) SetGroupName(GroupName string) {
	c.GroupName = GroupName
}
func (c *GroupAttributesX) SetMembership(Membership string) {
	c.Membership = Membership
}
func (c *GroupAttributesX) SetMembershipIdentifier(MembershipIdentifier string) {
	c.MembershipIdentifier = MembershipIdentifier
}

// *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *
//      JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON  JSON
// *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *

type CustomOrgLdapSettings struct {
	HostName                 string
	Port                     string
	SearchBase               string
	UserName                 string
	Password                 string
	AuthenticationMechanism  string
	GroupSearchBase          string
	IsGroupSearchBaseEnabled bool
	ConnectorType            string
	UserAttributes           UserAttributes
	GroupAttributes          GroupAttributes
}

type UserAttributes struct {
	ObjectClass               string
	ObjectIdentifier          string
	UserName                  string
	Email                     string
	FullName                  string
	GivenName                 string
	Surname                   string
	Telephone                 string
	GroupMembershipIdentifier string
}

type GroupAttributes struct {
	ObjectClass          string
	ObjectIdentifier     string
	GroupName            string
	Membership           string
	MembershipIdentifier string
}

type OrgEmailSettings struct {
	IsDefaultSmtpServer     bool
	IsDefaultOrgEmail       bool
	FromEmailAddress        string
	DefaultSubjectPrefix    string
	IsAlertEmailToAllAdmins bool
}

type OrgLdapSettings struct {
	OrgLdapMode           string
	CustomUsersOu         string
	CustomOrgLdapSettings CustomOrgLdapSettings
}

type OrgGeneralSettings struct {
	CanPublishCatalogs       bool  `json:"CanPublishCatalogs"`
	DeployedVMQuota          int64 `json:"DeployedVMQuota"`
	DelayAfterPowerOnSeconds int64 `json:"DelayAfterPowerOnSeconds"`
	StoredVmQuota            int64 `json:"StoredVmQuota"`
	UseServerBootSequence    bool  `json:"UseServerBootSequence"`
}

type Settings struct {
	OrgGeneralSettings OrgGeneralSettings `json:"OrgGeneralSettings"`
	OrgLdapSettings    OrgLdapSettings    `json:"OrgLdapSettings"`
	OrgEmailSettings   OrgEmailSettings   `json:"OrgEmailSettings"`
}

type OrgType struct {
	Description string `json:"Description"`
	FullName    string `json:"FullName"`
	IsEnabled   bool   `json:"IsEnabled,string"`
}

type ConfigType struct {
	Name     string   `json:"Name"`
	Org      OrgType  `json:"Org"`
	Settings Settings `json:"Settings"`
}

type JsonObject struct {
	ConfigType []ConfigType
}

func (j *JsonObject) AllConfigs() (config []ConfigType) {

	return j.ConfigType
}
