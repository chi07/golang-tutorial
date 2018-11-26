package utils

import "github.com/astaxie/beego/validation"

var validate = validation.Validation{}

func init() {
	validation.SetDefaultMessage(map[string]string {
		"Required":     "Không được để trống",
		"Min":          "Gía trị nhỏ nhất là %d",
		"Max":          "Giá trị lớn nhất là %d",
		"Range":        "Vùng giá trị phải thuộc %d đến %d",
		"MinSize":      "Kích thước nhỏ nhất %d",
		"MaxSize":      "Kích thước lớn nhất %d",
		"Length":       "Chiều dài %d",
		"Alpha":        "Phải là các chữ cái",
		"Numeric":      "Phải là ký tự số hợp lệ",
		"AlphaNumeric": "Phải là các chữ cái hoặc số",
		"Match":        "Phải khớp với сhuỗi %s",
		"NoMatch":      "Phải không trùng khớp với %s",
		"AlphaDash":    "Phải là các chữ caí, số, hoặc dâú gạch dưới (-_)",
		"Email":        "Địa chỉ email phải hợp lệ",
		"IP":           "Địa chỉ IP phải hợp lệ",
		"Base64":       "Phải là các kí tự mã 86 hợp lệ",
		"Mobile":       "Số điện thoại di động phải hợp lệ",
		"Tel":          "Số máy phải hợp lệ",
		"Phone":        "Số điện thoại di động hoặc số liên lạc phải hợp lệ",
		"ZipCode":      "Mã Zipcode phải hợp lệ",
	})
}

func GetValidator() validation.Validation {
	return validate
}

