package main

type abstractFactory interface{
	createButton() Button
	createLabel() Label
}

type macFactory struct{
}

func (m macFactory) createButton() Button{
	return MacButton{}
}

func (m macFactory) createLabel() Label{
	return MacLabel{}
}

type windowsFactory struct{
}

func (w windowsFactory) createButton() Button{
	return WindowsButton{}
}

func (w windowsFactory) createLabel() Label{
	return WindowsLabel{}
}
