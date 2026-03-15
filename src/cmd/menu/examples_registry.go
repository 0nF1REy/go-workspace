package menu

import (
	"fmt"
	"os"
	"path/filepath"
)

type ExampleInfo struct {
	Name     string
	MainPath string
}

func GetExamples() []ExampleInfo {
	examples := []struct {
		Name   string
		Folder string
	}{
		{"Hello World", "01_hello_world"},
		{"Pointer", "02_pointer"},
		{"Album Info", "03_album_info"},
		{"Loops", "04_loops"},
		{"Arrays", "05_arrays"},
		{"Maps", "06_maps"},
		{"Check Even/Odd", "07_check_even_odd"},

		{"Structs Value", "08_structs/01_value"},
		{"Structs Pointer", "08_structs/02_pointer"},
		{"Structs Pointer and Slice", "08_structs/03_pointer_and_slice"},
		{"Structs JSON", "08_structs/04_json"},

		{"Interfaces Basic", "09_interfaces/01_basic"},
		{"Interfaces Polymorphism", "09_interfaces/02_polymorphism"},
		{"Interfaces Type Assertion", "09_interfaces/03_type_assertion"},
		{"Interfaces Type Switch", "09_interfaces/04_type_switch"},
		{"Interfaces Empty", "09_interfaces/05_empty"},

		{"User Welcome", "10_user_welcome"},
		{"Auth", "11_auth"},
		{"Error", "12_error"},
		{"Variadic Functions", "13_variadic_functions"},

		{"Concurrency Goroutines", "14_concurrency/01_goroutines"},
		{"Concurrency Starvation", "14_concurrency/02_goroutine_starvation"},

		{"Channels Hello", "15_channels/01_hello_channel"},
		{"Channels Forever Block", "15_channels/02_forever_block"},
		{"Channels Select Timeout", "15_channels/03_select_timeout"},
		{"Channels Infinite Queue", "15_channels/04_infinite_queue"},
		{"Channels Multi Worker", "15_channels/05_multi_worker"},
	}

	root := os.Getenv("EXAMPLES_ROOT")

	if root == "" {
		fmt.Println("EXAMPLES_ROOT não definido. Configure no .env.")
		return nil
	}

	var result []ExampleInfo
	for _, ex := range examples {
		mainPath := filepath.Join(root, ex.Folder, "main.go")
		result = append(result, ExampleInfo{
			Name:     ex.Name,
			MainPath: mainPath,
		})
	}
	return result
}
