package middleware

type logInterface interface{
	Init(name string)
	AddMessage(key, value string)
	Print()
}
