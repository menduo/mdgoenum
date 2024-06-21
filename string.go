// Created by @menduo @ 2024/6/21
package mdgoenum

import (
	"encoding/json"
	"fmt"
	"sync"
)

// StringMember is an enum member, a specific value bound to a variable.
type StringMember struct {
	value string
	desc  string
}

// stringMemberJSON is a JSON representation of a StringMember.
type stringMemberJSON struct {
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type StringMemberMapType map[string]StringMember
type StringDescMapType map[string]string

// GetValue returns the value of the member.
func (sm *StringMember) GetValue() string {
	return sm.value
}

// GetDesc returns the description of the member.
func (sm *StringMember) GetDesc() string {
	return sm.desc
}

// MarshalJSON marshals a StringMember to JSON.
func (sm *StringMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(stringMemberJSON{
		Value: sm.value,
		Desc:  sm.desc,
	})
}

// UnmarshalJSON unmarshals a StringMember from JSON.
func (sm *StringMember) UnmarshalJSON(data []byte) error {
	var aux stringMemberJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	sm.value = aux.Value
	sm.desc = aux.Desc

	return nil
}

// StringEnum is a collection of enum members.
type StringEnum struct {
	mu               sync.RWMutex
	options          *optionType
	members          []StringMember
	v2map            StringMemberMapType
	v2descCacheData  StringDescMapType
	v2descCacheValid bool
}

// NewStringEnum constructs a new StringEnum.
func NewStringEnum(opFuncs ...OpsFuncType) *StringEnum {
	v2map := make(StringMemberMapType)
	members := make([]StringMember, 0)
	options := newOptionWithOpts(opFuncs...)
	return &StringEnum{
		options:          options,
		members:          members,
		v2map:            v2map,
		v2descCacheData:  make(StringDescMapType),
		v2descCacheValid: false,
	}
}

// TypeName is a string representation of the wrapped type.
func (e *StringEnum) TypeName() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return fmt.Sprintf("%T", e)
}

// Add adds a new member to the enum.
func (e *StringEnum) Add(value string, desc string) (*StringMember, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if ov, ok := e.v2map[value]; ok {
		return nil, fmt.Errorf("`%s`: duplicate member `%v`, desc: %s", e.TypeName(), ov.GetValue(), ov.GetDesc())
	}
	member := StringMember{value: value, desc: desc}

	// expand the members slice if necessary
	if len(e.members) == cap(e.members) {
		newMembers := make([]StringMember, len(e.members), 2*len(e.members)+1)
		copy(newMembers, e.members)
		e.members = newMembers
	}

	e.members = append(e.members, member)
	e.v2map[value] = member
	e.v2descCacheValid = false // set cache invalid

	return &member, nil
}

// MustAdd is like Add but panics if the value already exists.
func (e *StringEnum) MustAdd(value string, desc string) *StringMember {
	member, err := e.Add(value, desc)
	if err != nil {
		panic(err)
	}
	return member
}

// Get returns a member by its value.
func (e *StringEnum) Get(value string) (*StringMember, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if member, ok := e.v2map[value]; ok {
		return &member, nil
	}
	return nil, fmt.Errorf("`%s` member `%v` not found", e.TypeName(), value)
}

// MustGet is like Get but panics if the value doesn't exist.
func (e *StringEnum) MustGet(value string) *StringMember {
	member, err := e.Get(value)
	if err != nil {
		panic(err)
	}
	return member
}

// Len returns how many members the enum has.
func (e *StringEnum) Len() int {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return len(e.members)
}

// Contains returns true if the enum has the given member.
func (e *StringEnum) Contains(value string) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()

	_, ok := e.v2map[value]
	return ok
}

// Members returns all the members of the enum.
func (e *StringEnum) Members() []StringMember {
	e.mu.RLock()
	defer e.mu.RUnlock()

	members := make([]StringMember, len(e.members))
	copy(members, e.members)

	return e.members
}

// ToMemberMap returns a map of the enum members, with the value as the key and the member as the value.
func (e *StringEnum) ToMemberMap() StringMemberMapType {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return e.v2map
}

// ToDescMap returns a map of the enum members, with the value as the key and the description as the value.
func (e *StringEnum) ToDescMap() StringDescMapType {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if !e.v2descCacheValid {
		m := make(StringDescMapType, len(e.members))
		for _, v := range e.members {
			m[v.GetValue()] = v.GetDesc()
		}

		e.v2descCacheData = m
		e.v2descCacheValid = true
	}

	return e.v2descCacheData
}

// IsEmpty returns true if the enum doesn't have any members.
func (e *StringEnum) IsEmpty() bool {
	return e.Len() == 0
}
