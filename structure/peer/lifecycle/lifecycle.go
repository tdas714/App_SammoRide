package lifecycle

import "github.com/App-SammoRide/structure/peer"

// InstallChaincodeArgs is the message used as the argument to
// '_lifecycle.InstallChaincode'.
type InstallChaincodeArgs struct {
	ChaincodeInstallPackage []byte
}

func (m *InstallChaincodeArgs) GetChaincodeInstallPackage() []byte {
	if m != nil {
		return m.ChaincodeInstallPackage
	}
	return nil
}

//InstallChaincodeArgs is the message returned by
// '_lifecycle.InstallChaincode'.
type InstallChaincodeResult struct {
	PackageId string
	Label     string
}

func (m *InstallChaincodeResult) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}
func (m *InstallChaincodeResult) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// QueryInstalledChaincodeArgs is the message used as arguments
// '_lifecycle.QueryInstalledChaincode'
type QueryInstalledChaincodeArgs struct {
	PackageId string
}

func (m *QueryInstalledChaincodeArgs) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

// QueryInstalledChaincodeResult is the message returned by
// '_lifecycle.QueryInstalledChaincode'
type QueryInstalledChaincodeResult struct {
	PackageId  string
	Label      string
	References map[string]*QueryInstalledChaincodeResult_References
}

func (m *QueryInstalledChaincodeResult) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}
func (m *QueryInstalledChaincodeResult) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}
func (m *QueryInstalledChaincodeResult) GetReferences() map[string]*QueryInstalledChaincodeResult_References {
	if m != nil {
		return m.References
	}
	return nil
}

type QueryInstalledChaincodeResult_References struct {
	Chaincodes []*QueryInstalledChaincodeResult_Chaincode
}

func (m *QueryInstalledChaincodeResult_References) GetChaincodes() []*QueryInstalledChaincodeResult_Chaincode {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

type QueryInstalledChaincodeResult_Chaincode struct {
	Name    string
	Version string
}

func (m *QueryInstalledChaincodeResult_Chaincode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *QueryInstalledChaincodeResult_Chaincode) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// GetInstalledChaincodePackageArgs is the message used as the argument to
// '_lifecycle.GetInstalledChaincodePackage'.
type GetInstalledChaincodePackageArgs struct {
	PackageId string
}

func (m *GetInstalledChaincodePackageArgs) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

// GetInstalledChaincodePackageResult is the message returned by
// '_lifecycle.GetInstalledChaincodePackage'.
type GetInstalledChaincodePackageResult struct {
	ChaincodeInstallPackage []byte
}

func (m *GetInstalledChaincodePackageResult) GetChaincodeInstallPackage() []byte {
	if m != nil {
		return m.ChaincodeInstallPackage
	}
	return nil
}

type QueryInstalledChaincodesArgs struct{}

// QueryInstalledChaincodesResult is the message returned by
// '_lifecycle.QueryInstalledChaincodes'.  It returns a list of installed
// chaincodes, including a map of channel name to chaincode name and version
// pairs of chaincode definitions that reference this chaincode package.
type QueryInstalledChaincodesResult struct {
	InstalledChaincodes []*QueryInstalledChaincodesResult_InstalledChaincode
}

func (m *QueryInstalledChaincodesResult) GetInstalledChaincodes() []*QueryInstalledChaincodesResult_InstalledChaincode {
	if m != nil {
		return m.InstalledChaincodes
	}
	return nil
}

type QueryInstalledChaincodesResult_InstalledChaincode struct {
	PackageId  string
	Label      string
	References map[string]*QueryInstalledChaincodesResult_References
}

func (m *QueryInstalledChaincodesResult_InstalledChaincode) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}
func (m *QueryInstalledChaincodesResult_InstalledChaincode) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}
func (m *QueryInstalledChaincodesResult_InstalledChaincode) GetReferences() map[string]*QueryInstalledChaincodesResult_References {
	if m != nil {
		return m.References
	}
	return nil
}

type QueryInstalledChaincodesResult_References struct {
	Chaincodes []*QueryInstalledChaincodesResult_Chaincode
}

func (m *QueryInstalledChaincodesResult_References) GetChaincodes() []*QueryInstalledChaincodesResult_Chaincode {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

type QueryInstalledChaincodesResult_Chaincode struct {
	Name    string
	Version string
}

func (m *QueryInstalledChaincodesResult_Chaincode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *QueryInstalledChaincodesResult_Chaincode) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// ApproveChaincodeDefinitionForMyOrgArgs is the message used as arguments to
// `_lifecycle.ApproveChaincodeDefinitionForMyOrg`.
type ApproveChaincodeDefinitionForMyOrgArgs struct {
	Sequence            int64
	Name                string
	Version             string
	EndorsementPlugin   string
	ValidationPlugin    string
	ValidationParameter []byte
	Collections         *peer.CollectionConfigPackage
	InitRequired        bool
	Source              *ChaincodeSource
}

func (m *ApproveChaincodeDefinitionForMyOrgArgs) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}
func (m *ApproveChaincodeDefinitionForMyOrgArgs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *ApproveChaincodeDefinitionForMyOrgArgs) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}
func (m *ApproveChaincodeDefinitionForMyOrgArgs) GetEndorsementPlugin() string {
	if m != nil {
		return m.EndorsementPlugin
	}
	return ""
}
func (m *ApproveChaincodeDefinitionForMyOrgArgs) GetValidationPlugin() string {
	if m != nil {
		return m.ValidationPlugin
	}
	return ""
}
func (m *ApproveChaincodeDefinitionForMyOrgArgs) GetValidationParameter() []byte {
	if m != nil {
		return m.ValidationParameter
	}
	return nil
}
func (m *ApproveChaincodeDefinitionForMyOrgArgs) GetCollections() *peer.CollectionConfigPackage {
	if m != nil {
		return m.Collections
	}
	return nil
}
func (m *ApproveChaincodeDefinitionForMyOrgArgs) GetInitRequired() bool {
	if m != nil {
		return m.InitRequired
	}
	return false
}
func (m *ApproveChaincodeDefinitionForMyOrgArgs) GetSource() *ChaincodeSource {
	if m != nil {
		return m.Source
	}
	return nil
}

type ChaincodeSource struct {
	// Types that are valid to be assigned to Type:
	//	*ChaincodeSource_Unavailable_
	//	*ChaincodeSource_LocalPackage
	Type isChaincodeSource_Type
}

type isChaincodeSource_Type interface {
	isChaincodeSource_Type()
}

type ChaincodeSource_Unavailable_ struct {
	Unavailable *ChaincodeSource_Unavailable
}

type ChaincodeSource_LocalPackage struct {
	LocalPackage *ChaincodeSource_Local
}

func (*ChaincodeSource_Unavailable_) isChaincodeSource_Type() {}
func (*ChaincodeSource_LocalPackage) isChaincodeSource_Type() {}
func (m *ChaincodeSource) GetType() isChaincodeSource_Type {
	if m != nil {
		return m.Type
	}
	return nil
}
func (m *ChaincodeSource) GetUnavailable() *ChaincodeSource_Unavailable {
	if x, ok := m.GetType().(*ChaincodeSource_Unavailable_); ok {
		return x.Unavailable
	}
	return nil
}
func (m *ChaincodeSource) GetLocalPackage() *ChaincodeSource_Local {
	if x, ok := m.GetType().(*ChaincodeSource_LocalPackage); ok {
		return x.LocalPackage
	}
	return nil
}

type ChaincodeSource_Unavailable struct{}

type ChaincodeSource_Local struct {
	PackageId string
}

func (m *ChaincodeSource_Local) GetPackageId() string {
	if m != nil {
		return m.PackageId
	}
	return ""
}

// ApproveChaincodeDefinitionForMyOrgResult is the message returned by
// `_lifecycle.ApproveChaincodeDefinitionForMyOrg`. Currently it returns
// nothing, but may be extended in the future.
type ApproveChaincodeDefinitionForMyOrgResult struct{}

// CommitChaincodeDefinitionArgs is the message used as arguments to
// `_lifecycle.CommitChaincodeDefinition`.
type CommitChaincodeDefinitionArgs struct {
	Sequence            int64
	Name                string
	Version             string
	EndorsementPlugin   string
	ValidationPlugin    string
	ValidationParameter []byte
	Collections         *peer.CollectionConfigPackage
	InitRequired        bool
}

func (m *CommitChaincodeDefinitionArgs) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}
func (m *CommitChaincodeDefinitionArgs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *CommitChaincodeDefinitionArgs) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}
func (m *CommitChaincodeDefinitionArgs) GetEndorsementPlugin() string {
	if m != nil {
		return m.EndorsementPlugin
	}
	return ""
}
func (m *CommitChaincodeDefinitionArgs) GetValidationPlugin() string {
	if m != nil {
		return m.ValidationPlugin
	}
	return ""
}
func (m *CommitChaincodeDefinitionArgs) GetValidationParameter() []byte {
	if m != nil {
		return m.ValidationParameter
	}
	return nil
}
func (m *CommitChaincodeDefinitionArgs) GetCollections() *peer.CollectionConfigPackage {
	if m != nil {
		return m.Collections
	}
	return nil
}
func (m *CommitChaincodeDefinitionArgs) GetInitRequired() bool {
	if m != nil {
		return m.InitRequired
	}
	return false
}

type CommitChaincodeDefinitionResult struct{}

type QueryChaincodeDefinitionsResult_ChaincodeDefinition struct {
	Name                string
	Sequence            int64
	Version             string
	EndorsementPlugin   string
	ValidationPlugin    string
	ValidationParameter []byte
	Collections         *peer.CollectionConfigPackage
	InitRequired        bool
}

func (m *QueryChaincodeDefinitionsResult_ChaincodeDefinition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *QueryChaincodeDefinitionsResult_ChaincodeDefinition) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}
func (m *QueryChaincodeDefinitionsResult_ChaincodeDefinition) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}
func (m *QueryChaincodeDefinitionsResult_ChaincodeDefinition) GetEndorsementPlugin() string {
	if m != nil {
		return m.EndorsementPlugin
	}
	return ""
}
func (m *QueryChaincodeDefinitionsResult_ChaincodeDefinition) GetValidationPlugin() string {
	if m != nil {
		return m.ValidationPlugin
	}
	return ""
}
func (m *QueryChaincodeDefinitionsResult_ChaincodeDefinition) GetValidationParameter() []byte {
	if m != nil {
		return m.ValidationParameter
	}
	return nil
}
func (m *QueryChaincodeDefinitionsResult_ChaincodeDefinition) GetCollections() *peer.CollectionConfigPackage {
	if m != nil {
		return m.Collections
	}
	return nil
}
func (m *QueryChaincodeDefinitionsResult_ChaincodeDefinition) GetInitRequired() bool {
	if m != nil {
		return m.InitRequired
	}
	return false
}
