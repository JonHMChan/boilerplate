package core

import ()

// Declare every setting here with a pointer to a SettingStruct
type SettingsStruct struct {
}

// For each setting, add correct attributes for the current value
func (settings *SettingsStruct) Init() {
}

type SettingStruct struct {
	defaultValue bool // What the default value should be when in progress. You should NEVER use this attribute in a conditional.
	readyValue   bool // Whether it is production ready. If false, will always be inaccessible to all users on production.
}

func (setting *SettingStruct) Enabled() bool {
	if Config.Prod() && !setting.readyValue {
		return false
	}
	return setting.defaultValue
}
