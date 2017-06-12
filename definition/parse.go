package definition

func (d *AppDefinition) Parse() {
	for k, def := range d.Integrations.Panels {
		if def.Hook.AppID == "" {
			def.Hook.AppID = d.AppID
		}
		if def.Hook.VendorID == "" {
			def.Hook.VendorID = d.VendorID
		}
		d.Integrations.Panels[k] = def
	}
	for k, def := range d.Integrations.HeaderMenuItems {
		if def.Hook.AppID == "" {
			def.Hook.AppID = d.AppID
		}
		if def.Hook.VendorID == "" {
			def.Hook.VendorID = d.VendorID
		}
		d.Integrations.HeaderMenuItems[k] = def
	}
	for k, def := range d.Integrations.PageMenuItems {
		if def.Hook.AppID == "" {
			def.Hook.AppID = d.AppID
		}
		if def.Hook.VendorID == "" {
			def.Hook.VendorID = d.VendorID
		}
		d.Integrations.PageMenuItems[k] = def
	}
	for k, def := range d.Integrations.HeaderActions {
		if def.Hook.AppID == "" {
			def.Hook.AppID = d.AppID
		}
		if def.Hook.VendorID == "" {
			def.Hook.VendorID = d.VendorID
		}
		d.Integrations.HeaderActions [k] = def
	}
	for k, def := range d.Integrations.PageActions {
		if def.Hook.AppID == "" {
			def.Hook.AppID = d.AppID
		}
		if def.Hook.VendorID == "" {
			def.Hook.VendorID = d.VendorID
		}
		d.Integrations.PageActions [k] = def
	}
}
