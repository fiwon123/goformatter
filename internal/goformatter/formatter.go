package formatter

// internal/goformatter/formatter.go

func Process(filepath string, dirOutput string, isChecking bool, idDryRun bool, isPretty bool, isInPlace bool, fileType string) {
	// 	print_msg("-------------------------------------")
	// 	print_msg("Verifying filepath...")
	// 	if file_type != "" {
	// 		formatter_class = get_formatter(file_type)

	// 		if formatter_class == None {
	// 			print_error("Invalid parameter file_type: " + file_type)
	// 		}
	// 	} else {
	// 		formatter_class = get_formatter(get_extension(filepath)[1:])

	// 		if formatter_class == None {
	// 			print_error("Invalid input file: " + filepath)
	// 		}
	// 	}

	// 	print_msg("Verification passed.")
	// formatter:
	// 	BaseFormatter = formatter_class(filepath)

	// 	if isChecking {
	// 		formatter.check()
	// 	}
	// 	if isDryRun {
	// 		formatter.dry_run()
	// 	}
	// 	if !isChecking && !isDryRun {
	// 		formatter.format(dir_output, in_place)
	// 	}
}
