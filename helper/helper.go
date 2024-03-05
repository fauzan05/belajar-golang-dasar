package helper

var Application = "Aplikasi Versi 1.0.0"
var application = "Aplikasi Versi 1.0.0"

// method ini tidak bisa diakses di package manapun karena diawali dengan huruf kecil
func sayGoodbye(name string) string {
	return "Goodbye " + name
}

// jika method diawali huruf kapital, maka bisa diakses di package manapun
func SayHello(name string) string {
	return "Hello " + name
}
