package unit

// // recovers from the error and asserts error with the expected one
// func RecoverError(expectedError *error_handling.CustomError, t *testing.T) {
// 	err := recover().(*error_handling.CustomError)
// 	assert.Equal(t, expectedError, err)
// }

// // check that the status is the same and the error message is not empty
// func RecoverNotStandardError(expectedStatusCode int, t *testing.T) {
// 	err := recover().(*error_handling.CustomError)

// 	assert.Equal(t, expectedStatusCode, err.Status())
// 	assert.NotEqual(t, "", err.Error())
// }
