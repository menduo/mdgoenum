// Created by @menduo @ 2024/6/21
package mdgoenum

import (
	"encoding/json"
	"fmt"
	"sync"
)

// IntMember is an enum member, a specific value bound to a variable.
type IntMember struct {
	value int
	desc  string
}

// IntMemberJSON is a JSON representation of a IntMember.
type IntMemberJSON struct {
	Value int    `json:"value"`
	Desc  string `json:"desc"`
}

type IntMemberMapType map[int]IntMember
type IntDescMapType map[int]string

// GetValue returns the value of the member.
func (sm *IntMember) GetValue() int {
	return sm.value
}

// GetDesc returns the description of the member.
func (sm *IntMember) GetDesc() string {
	return sm.desc
}

// MarshalJSON marshals a IntMember to JSON.
func (sm *IntMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(IntMemberJSON{
		Value: sm.value,
		Desc:  sm.desc,
	})
}

// UnmarshalJSON unmarshals a IntMember from JSON.
func (sm *IntMember) UnmarshalJSON(data []byte) error {
	var aux IntMemberJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	sm.value = aux.Value
	sm.desc = aux.Desc

	return nil
}

// IntEnum is a collection of enum members.
type IntEnum struct {
	mu               sync.RWMutex
	options          *optionType
	members          []IntMember
	v2map            IntMemberMapType
	v2descCacheData  IntDescMapType
	v2descCacheValid bool
}

// NewIntEnum constructs a new IntEnum.
func NewIntEnum(opFuncs ...OpsFuncType) *IntEnum {
	v2map := make(IntMemberMapType)
	members := make([]IntMember, 0)
	options := newOptionWithOpts(opFuncs...)
	return &IntEnum{
		options:          options,
		members:          members,
		v2map:            v2map,
		v2descCacheData:  make(IntDescMapType),
		v2descCacheValid: false,
	}
}

// TypeName is a string representation of the wrapped type.
func (e *IntEnum) TypeName() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return fmt.Sprintf("%T", e)
}

// Add adds a new member to the enum.
func (e *IntEnum) Add(value int, desc string) (*IntMember, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if ov, ok := e.v2map[value]; ok {
		return nil, fmt.Errorf("`%s`: duplicate member `%v`, desc: %s", e.TypeName(), ov.GetValue(), ov.GetDesc())
	}
	member := IntMember{value: value, desc: desc}

	// expand the members slice if necessary
	if len(e.members) == cap(e.members) {
		newMembers := make([]IntMember, len(e.members), 2*len(e.members)+1)
		copy(newMembers, e.members)
		e.members = newMembers
	}

	e.members = append(e.members, member)
	e.v2map[value] = member
	e.v2descCacheValid = false // set cache invalid

	return &member, nil
}

// MustAdd is like Add but panics if the value already exists.
func (e *IntEnum) MustAdd(value int, desc string) *IntMember {
	member, err := e.Add(value, desc)
	if err != nil {
		panic(err)
	}
	return member
}

// Get returns a member by its value.
func (e *IntEnum) Get(value int) (*IntMember, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if member, ok := e.v2map[value]; ok {
		return &member, nil
	}
	return nil, fmt.Errorf("`%s` member `%v` not found", e.TypeName(), value)
}

// MustGet is like Get but panics if the value doesn't exist.
func (e *IntEnum) MustGet(value int) *IntMember {
	member, err := e.Get(value)
	if err != nil {
		panic(err)
	}
	return member
}

// Len returns how many members the enum has.
func (e *IntEnum) Len() int {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return len(e.members)
}

// Contains returns true if the enum has the given member.
func (e *IntEnum) Contains(value int) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()

	_, ok := e.v2map[value]
	return ok
}

// Members returns all the members of the enum.
func (e *IntEnum) Members() []IntMember {
	e.mu.RLock()
	defer e.mu.RUnlock()

	members := make([]IntMember, len(e.members))
	copy(members, e.members)

	return e.members
}

// ToMemberMap returns a map of the enum members, with the value as the key and the member as the value.
func (e *IntEnum) ToMemberMap() IntMemberMapType {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return e.v2map
}

// ToDescMap returns a map of the enum members, with the value as the key and the description as the value.
func (e *IntEnum) ToDescMap() IntDescMapType {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if !e.v2descCacheValid {
		m := make(IntDescMapType, len(e.members))
		for _, v := range e.members {
			m[v.GetValue()] = v.GetDesc()
		}

		e.v2descCacheData = m
		e.v2descCacheValid = true
	}

	return e.v2descCacheData
}

// IsEmpty returns true if the enum doesn't have any members.
func (e *IntEnum) IsEmpty() bool {
	return e.Len() == 0
}
