## httptest.NewRequest
# How to include path with value example

Summary:
```
...
r, err := http.NewRequest(method, path, body)
r.SetPathValue("key", "value")
...
```

Complete example in main_test.go