package lifecycle

// StateMetadata describes the keys in a namespace.  It is necessary because
// in collections, range scans are not possible during transactions which
// write.  Therefore we must track the keys in our namespace ourselves.
type StateMetadata struct {
	Datatype string
	Fields   []string
}

func (m *StateMetadata) GetDatatype() string {
	if m != nil {
		return m.Datatype
	}
	return ""
}

func (m *StateMetadata) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

// StateData encodes a particular field of a datatype
type StateData struct {
	// Types that are valid to be assigned to Type:
	//	*StateData_Int64
	//	*StateData_Bytes
	//	*StateData_String_
	Type isStateData_Type
}

type isStateData_Type interface {
	isStateData_Type()
}

type StateData_Int64 struct {
	Int64 int64
}

type StateData_Bytes struct {
	Bytes []byte
}

type StateData_String_ struct {
	String_ string
}

func (*StateData_Int64) isStateData_Type() {}

func (*StateData_Bytes) isStateData_Type() {}

func (*StateData_String_) isStateData_Type() {}

func (m *StateData) GetType() isStateData_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *StateData) GetInt64() int64 {
	if x, ok := m.GetType().(*StateData_Int64); ok {
		return x.Int64
	}
	return 0
}

func (m *StateData) GetBytes() []byte {
	if x, ok := m.GetType().(*StateData_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (m *StateData) GetString_() string {
	if x, ok := m.GetType().(*StateData_String_); ok {
		return x.String_
	}
	return ""
}
