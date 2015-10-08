package main

import "fmt"

func ExampleBasicConverter() {
	raw := `
type Ha struct {
	ID          int64
	Name        string
	Amount      float32
	Version     int64
	Alive       bool
}
`
	fmt.Println(Convert(raw))

	// Output:
	// $Ha = array();
	// $Ha['ID'] = (int) $data['Ha_ID'];
	// $Ha['Name'] = (string) $data['Ha_Name'];
	// $Ha['Amount'] = (float) $data['Ha_Amount'];
	// $Ha['Version'] = (int) $data['Ha_Version'];
	// $Ha['Alive'] = (bool) $data['Ha_Alive'];
}

func ExampleNothingToDo() {
	raw := `type Ha struct {
`
	fmt.Println(Convert(raw))

	// Output:
	// (╯°□°）╯︵ ┻━┻
}
