package mappers

func GetStringPointer(value interface{}) *string {
	if value == nil {
		return nil
	}
	str, ok := value.(*string)
	if !ok {
		return nil
	}
	return str
}
