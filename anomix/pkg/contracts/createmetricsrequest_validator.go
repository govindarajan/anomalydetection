// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// Autogenerated by rest_tool -t CreateMetricsRequest -m  validator
package contracts

// Validate checks if a contract is valid or not and return with appropriate message
func (_receiver_ *CreateMetricsRequest) Validate() *Error {
	if _receiver_.Request != nil {
		if typ, ok := interface{}(_receiver_.Request).(Validator); ok {
			if err := typ.Validate(); err != nil {
				return err
			}
		}
	}
	if _receiver_.Anomaly != nil {
		if typ, ok := interface{}(_receiver_.Anomaly).(Validator); ok {
			if err := typ.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}