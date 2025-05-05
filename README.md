# 🔍 About
A formatter CLI tool using go language

## ✨ Features
Simple CLI interface.

## 📦 Installation
```bash
go install
```

## 🛠 Usage
Use this following command to formmat any file available in the package:
```bash
goformatter -f
```

Use this command to list all availables parameters:
```bash
goformatter --help
```

## 📁 Project Structure
goformatter/
├── cmd/                
│   └── root.go           
├── internal/                      
|   ├── goformatter/
│   ├── testdata/
│   ├── baseFormatter.go
│   ├── registry.go
│   ├── goformatter_test.go
│   ├── formatter.go
│   ├── jsonFormatter.go
│   ├── tomlFormatter.go
│   ├── yamlFormatter.go
│   ├── logger/
│   │   └── logger.go
│   └── utils/
│       └── utils.go
├── main.go                 # entry point
├── go.mod                  
├── go.sum
├── LICENSE     
└── README.md

## 🧪 Running Tests
```bash
go test
```

## 🧾 License
This project is licensed under the MIT License - see the LICENSE file for details.