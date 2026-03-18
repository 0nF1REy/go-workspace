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
		// 01_hello_world
		{"Hello World", "01_hello_world"},

		// 02_pointer
		{"Pointer", "02_pointer"},

		// 03_album_info
		{"Album Info", "03_album_info"},

		// 04_loops
		{"Loops", "04_loops"},

		// 05_arrays
		{"Arrays", "05_arrays"},

		// 06_maps
		{"Maps", "06_maps"},

		// 07_check_even_odd
		{"Check Even/Odd", "07_check_even_odd"},

		// 08_structs
		{"Structs Value", "08_structs/01_value"},
		{"Structs Pointer", "08_structs/02_pointer"},
		{"Structs Pointer and Slice", "08_structs/03_pointer_and_slice"},
		{"Structs JSON", "08_structs/04_json"},

		// 09_interfaces
		{"Interfaces Basic", "09_interfaces/01_basic"},
		{"Interfaces Polymorphism", "09_interfaces/02_polymorphism"},
		{"Interfaces Type Assertion", "09_interfaces/03_type_assertion"},
		{"Interfaces Type Switch", "09_interfaces/04_type_switch"},
		{"Interfaces Empty", "09_interfaces/05_empty"},

		// 10_user_welcome
		{"User Welcome", "10_user_welcome"},

		// 11_auth
		{"Auth", "11_auth"},

		// 12_error
		{"Error", "12_error"},

		// 13_variadic_functions
		{"Variadic Functions", "13_variadic_functions"},

		// 14_concurrency
		{"Concurrency Goroutines", "14_concurrency/01_goroutines"},
		{"Concurrency Starvation", "14_concurrency/02_goroutine_starvation"},

		// 15_channels
		{"Channels Hello", "15_channels/01_hello_channel"},
		{"Channels Forever Block", "15_channels/02_forever_block"},
		{"Channels Select Timeout", "15_channels/03_select_timeout"},
		{"Channels Infinite Queue", "15_channels/04_infinite_queue"},
		{"Channels Multi Worker", "15_channels/05_multi_worker"},

		// 16_generics
		{"Generics", "16_generics"},

		// 17_context
		{"Context Client", "17_context/client"},
		{"Context Hotel", "17_context/hotel"},
		{"Context Server", "17_context/server"},

		// 18_web_scraper
		{"Web Scraper Wikipedia", "18_web_scraper/01_wikipedia"},
		{"Web Scraper Oxylabs Single Page", "18_web_scraper/02_oxylabs/01_single_page"},
		{"Web Scraper Oxylabs Multi Page", "18_web_scraper/02_oxylabs/02_multi_page"},
		{"Web Scraper Oxylabs Auto Pagination", "18_web_scraper/02_oxylabs/03_auto_pagination"},

		// 19_tcp_port_scanner
		{"TCP Port Scanner Sequential", "19_tcp_port_scanner/01_sequential"},
		{"TCP Port Scanner Concurrent", "19_tcp_port_scanner/02_concurrent"},
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
