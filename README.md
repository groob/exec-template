Execute arbitrary Go template using some JSON input.

# Example

```
./exec-template \
    -json='{"Name": "groob", "Status": "active"}' \
    -template='{{.Name}} is {{.Status}}.' \
    -string
```

Use the `-string` flag to specify wether the template is a path to a file or a string parameter like above.

# Install

If you have Go installed:

```
go get github.com/groob/exec-template
```

Otherwise, you can download the binary from the [releases](https://github.com/groob/exec-template/releases) page.



# Usage

```
  -json string
       	json input to use in template
  -string
       	template is a string, not a file
  -template string
       	path to template
```

Go template tutorial: https://golang.org/pkg/text/template/
