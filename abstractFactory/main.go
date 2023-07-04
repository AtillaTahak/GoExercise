package main


func main(){
	var factory abstractFactory

	platform := "mac"
	if platform == "mac"{
		factory = macFactory{}
	}else if platform == "windows"{
		factory = windowsFactory{}
	}

	button := factory.createButton()
	button.printButton()

	label := factory.createLabel()
	label.printLabel()
	
}